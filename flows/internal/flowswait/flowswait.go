// Package flowswait is the polling loop behind the flows CreateAndWait helpers.
//
// Hand-written (listed in .fernignore). The helpers call the generated WithRawResponse Create and
// Get methods, which expose the response headers, so the loop can honour the server's
// Retry-After header. The clock, sleep and jitter sit behind Pacer so the tests run without
// waiting.
package flowswait

import (
	"context"
	"fmt"
	"math"
	"math/rand/v2"
	"net/http"
	"strconv"
	"strings"
	"time"

	flows "github.com/elevenlabs/elevenlabs-go/flows"
	option "github.com/elevenlabs/elevenlabs-go/option"
)

// pollJitter spreads out clients that started together so their status checks don't arrive in
// lockstep.
const pollJitter = 0.2

// Pacer supplies time and randomness to the wait: the real clock in production, a fake in tests.
type Pacer interface {
	// Now is the current time. Elapsed time and Retry-After HTTP dates are measured against it.
	Now() time.Time
	// Sleep waits for d, returning ctx.Err() early if ctx is done first.
	Sleep(ctx context.Context, d time.Duration) error
	// Jitter is a random number in [0, 1).
	Jitter() float64
}

type realPacer struct{}

func (realPacer) Now() time.Time { return time.Now() }

func (realPacer) Jitter() float64 { return rand.Float64() }

func (realPacer) Sleep(ctx context.Context, d time.Duration) error {
	timer := time.NewTimer(d)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

type pacerKey struct{}

// WithPacer returns a context that makes CreateAndWait use p instead of the real clock. Only the
// tests call it; the package is internal, so it is not part of the SDK's API.
func WithPacer(ctx context.Context, p Pacer) context.Context {
	return context.WithValue(ctx, pacerKey{}, p)
}

func pacerFrom(ctx context.Context) Pacer {
	if p, ok := ctx.Value(pacerKey{}).(Pacer); ok {
		return p
	}
	return realPacer{}
}

// Waiter is one wait, started just before the create call so WaitTimeout covers it.
type Waiter struct {
	timeout     time.Duration
	minInterval time.Duration
	pacer       Pacer
	startedAt   time.Time
}

// Begin validates the options (a nil pointer means the defaults) and starts the clock.
func Begin(ctx context.Context, options *flows.WaitOptions) (*Waiter, error) {
	var resolved flows.WaitOptions
	if options != nil {
		resolved = *options
	}
	if resolved.MinPollInterval == 0 {
		resolved.MinPollInterval = flows.DefaultMinPollInterval
	}
	if resolved.MinPollInterval < flows.SmallestMinPollInterval {
		return nil, fmt.Errorf(
			"%w: MinPollInterval must be at least %s, got %s",
			flows.ErrInvalidWaitOptions, flows.SmallestMinPollInterval, resolved.MinPollInterval,
		)
	}
	if resolved.WaitTimeout < 0 {
		return nil, fmt.Errorf(
			"%w: WaitTimeout must not be negative, got %s",
			flows.ErrInvalidWaitOptions, resolved.WaitTimeout,
		)
	}
	pacer := pacerFrom(ctx)
	return &Waiter{
		timeout:     resolved.WaitTimeout,
		minInterval: resolved.MinPollInterval,
		pacer:       pacer,
		startedAt:   pacer.Now(),
	}, nil
}

// Target names what is being waited on, for the timeout and cancellation errors.
type Target struct {
	ID         string
	TemplateID string
}

func (t Target) String() string {
	if t.TemplateID != "" {
		return fmt.Sprintf("template run %s of template %s", t.ID, t.TemplateID)
	}
	return "generation " + t.ID
}

// IsTerminal reports whether status means the work has finished.
func IsTerminal(status string) bool {
	return status == "completed" || status == "failed"
}

// PollOptions returns the caller's options for the status checks: all of them, followed by one
// that clears the body properties, which only the create call takes (a GET has no body).
func PollOptions(opts []option.RequestOption) []option.RequestOption {
	poll := make([]option.RequestOption, 0, len(opts)+1)
	poll = append(poll, opts...)
	return append(poll, option.WithBodyProperties(nil))
}

// Get fetches the current state: the response, its status and the response headers.
type Get[T any] func(ctx context.Context) (T, string, http.Header, error)

// Poll calls get until the status is terminal and returns that response. It starts from the
// create call's status, response and headers; the caller returns early when the create response
// is already terminal.
func Poll[T any](
	ctx context.Context,
	w *Waiter,
	target Target,
	status string,
	lastResponse interface{},
	header http.Header,
	get Get[T],
) (T, error) {
	var zero T
	for {
		delay := w.delay(header)
		if w.timeout > 0 {
			// The last sleep ends on the deadline; the GET after it takes time, so the next pass
			// returns the timeout.
			remaining := w.timeout - w.pacer.Now().Sub(w.startedAt)
			if remaining <= 0 {
				return zero, &flows.WaitTimeoutError{
					ID:           target.ID,
					TemplateID:   target.TemplateID,
					WaitTimeout:  w.timeout,
					LastStatus:   status,
					LastResponse: lastResponse,
				}
			}
			delay = min(delay, remaining)
		}
		if err := w.pacer.Sleep(ctx, delay); err != nil {
			return zero, fmt.Errorf("flows: stopped waiting for %s: %w", target, err)
		}
		response, responseStatus, responseHeader, err := get(ctx)
		if err != nil {
			return zero, err
		}
		if IsTerminal(responseStatus) {
			return response, nil
		}
		status, lastResponse, header = responseStatus, response, responseHeader
	}
}

// delay is the wait before the next status check: the server's Retry-After, never below the
// minimum interval, plus 0-20% jitter.
func (w *Waiter) delay(header http.Header) time.Duration {
	base := w.minInterval
	if retryAfter, ok := ParseRetryAfter(header, w.pacer.Now()); ok && retryAfter > base {
		base = retryAfter
	}
	scaled := float64(base) * (1 + w.pacer.Jitter()*pollJitter)
	if scaled >= math.MaxInt64 {
		return time.Duration(math.MaxInt64)
	}
	return time.Duration(scaled)
}

// ParseRetryAfter reads a Retry-After header given as whole seconds or an HTTP date (RFC 9110).
// It reports false when the header is absent or unparseable. A date in the past is a wait of zero.
// Kept here rather than shared with the generated retrier, whose private helpers can change on
// any regeneration.
func ParseRetryAfter(header http.Header, now time.Time) (time.Duration, bool) {
	value := strings.TrimSpace(header.Get("Retry-After"))
	if value == "" {
		return 0, false
	}
	if isDigits(value) {
		seconds, err := strconv.ParseInt(value, 10, 64)
		if err != nil || seconds > int64(math.MaxInt64/time.Second) {
			// All digits, so this only fails on absurd lengths; treat those as "wait forever".
			return time.Duration(math.MaxInt64), true
		}
		return time.Duration(seconds) * time.Second, true
	}
	when, err := http.ParseTime(value)
	if err != nil {
		return 0, false
	}
	return max(when.Sub(now), 0), true
}

func isDigits(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}
