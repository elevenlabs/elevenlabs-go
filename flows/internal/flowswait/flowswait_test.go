// Hand-written (listed in .fernignore): unit tests of the Retry-After parsing and the option
// handling behind the flows CreateAndWait helpers.

package flowswait

import (
	"context"
	"math"
	"net/http"
	"testing"
	"time"

	core "github.com/elevenlabs/elevenlabs-go/core"
	option "github.com/elevenlabs/elevenlabs-go/option"
)

func TestParseRetryAfter(t *testing.T) {
	now := time.Date(2026, 11, 19, 12, 0, 0, 0, time.UTC)
	cases := []struct {
		name   string
		value  string
		want   time.Duration
		wantOK bool
	}{
		{"absent", "", 0, false},
		{"seconds", "5", 5 * time.Second, true},
		{"seconds with whitespace", " 12 ", 12 * time.Second, true},
		{"zero seconds", "0", 0, true},
		{"IMF-fixdate", now.Add(30 * time.Second).Format(http.TimeFormat), 30 * time.Second, true},
		{"RFC 850 date", now.Add(10 * time.Second).Format(time.RFC850), 10 * time.Second, true},
		{"ANSI C date", now.Add(20 * time.Second).Format(time.ANSIC), 20 * time.Second, true},
		{"date in the past", now.Add(-time.Minute).Format(http.TimeFormat), 0, true},
		{"absurdly long number", "99999999999999999999999", time.Duration(math.MaxInt64), true},
		{"negative number", "-5", 0, false},
		{"fractional seconds", "1.5", 0, false},
		{"garbage", "soon", 0, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			header := http.Header{}
			if tc.value != "" {
				header.Set("Retry-After", tc.value)
			}
			got, ok := ParseRetryAfter(header, now)
			if got != tc.want || ok != tc.wantOK {
				t.Fatalf("ParseRetryAfter(%q) = %s, %v; want %s, %v", tc.value, got, ok, tc.want, tc.wantOK)
			}
		})
	}
}

type fixedPacer struct {
	now    time.Time
	jitter float64
}

func (p fixedPacer) Now() time.Time                             { return p.now }
func (p fixedPacer) Jitter() float64                            { return p.jitter }
func (p fixedPacer) Sleep(context.Context, time.Duration) error { return nil }

func TestDelay(t *testing.T) {
	now := time.Date(2026, 11, 19, 12, 0, 0, 0, time.UTC)
	waiter := func(jitter float64) *Waiter {
		return &Waiter{minInterval: 3 * time.Second, pacer: fixedPacer{now: now, jitter: jitter}}
	}
	header := func(value string) http.Header {
		h := http.Header{}
		h.Set("Retry-After", value)
		return h
	}
	cases := []struct {
		name   string
		header http.Header
		jitter float64
		want   time.Duration
	}{
		{"no header uses the minimum", http.Header{}, 0, 3 * time.Second},
		{"unparseable header uses the minimum", header("later"), 0, 3 * time.Second},
		{"Retry-After below the minimum is raised to it", header("1"), 0, 3 * time.Second},
		{"Retry-After above the minimum is honoured", header("8"), 0, 8 * time.Second},
		{"jitter adds up to 20 percent", header("10"), 0.999, 10*time.Second + 1998*time.Millisecond},
		{"an endless Retry-After does not overflow", header("99999999999999999999999"), 0.5, time.Duration(math.MaxInt64)},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := waiter(tc.jitter).delay(tc.header); got != tc.want {
				t.Fatalf("delay = %s, want %s", got, tc.want)
			}
		})
	}
}

func TestBeginAppliesDefaults(t *testing.T) {
	w, err := Begin(context.Background(), nil)
	if err != nil {
		t.Fatalf("Begin(nil): %v", err)
	}
	if w.minInterval != 3*time.Second || w.timeout != 0 {
		t.Fatalf("defaults = %s interval, %s timeout", w.minInterval, w.timeout)
	}
	if _, ok := w.pacer.(realPacer); !ok {
		t.Fatalf("pacer = %T, want the real clock", w.pacer)
	}
}

func TestPollOptionsDropsOnlyBodyProperties(t *testing.T) {
	opts := []option.RequestOption{
		option.WithHTTPHeader(http.Header{"X-Trace": []string{"abc"}}),
		option.WithBodyProperties(map[string]interface{}{"extra": 1}),
		option.WithMaxAttempts(5),
	}
	resolved := core.NewRequestOptions(PollOptions(opts)...)
	if len(resolved.BodyProperties) != 0 {
		t.Fatalf("BodyProperties = %v, want none", resolved.BodyProperties)
	}
	if resolved.HTTPHeader.Get("X-Trace") != "abc" || resolved.MaxAttempts != 5 {
		t.Fatalf("other options lost: %+v", resolved)
	}
	if len(core.NewRequestOptions(opts...).BodyProperties) != 1 {
		t.Fatalf("PollOptions modified the caller's options")
	}
}
