// Hand-written (listed in .fernignore): end-to-end tests of the CreateAndWait helpers on
// client.Flows.Image, .Video, .TextToSpeech and .Templates.Runs against an httptest server. A fake
// pacer (flowswait.WithPacer) records the waits instead of sleeping, except in the tests that
// exercise the real clock and context cancellation.

package flows_test

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	elevenlabs "github.com/elevenlabs/elevenlabs-go"
	core "github.com/elevenlabs/elevenlabs-go/core"
	flows "github.com/elevenlabs/elevenlabs-go/flows"
	flowsclient "github.com/elevenlabs/elevenlabs-go/flows/client"
	flowswait "github.com/elevenlabs/elevenlabs-go/flows/internal/flowswait"
	templates "github.com/elevenlabs/elevenlabs-go/flows/templates"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

// fakePacer advances a virtual clock by each requested sleep instead of sleeping.
type fakePacer struct {
	mu     sync.Mutex
	now    time.Time
	jitter float64
	sleeps []time.Duration
}

func newFakePacer() *fakePacer {
	return &fakePacer{now: time.Date(2026, 11, 19, 12, 0, 0, 0, time.UTC)}
}

func (p *fakePacer) Now() time.Time {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.now
}

func (p *fakePacer) Jitter() float64 { return p.jitter }

func (p *fakePacer) Sleep(ctx context.Context, d time.Duration) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	p.sleeps = append(p.sleeps, d)
	p.now = p.now.Add(d)
	return nil
}

func (p *fakePacer) recorded() []time.Duration {
	p.mu.Lock()
	defer p.mu.Unlock()
	return append([]time.Duration(nil), p.sleeps...)
}

// reply is one scripted response.
type reply struct {
	status int
	header map[string]string
	body   string
}

type recordedRequest struct {
	method string
	path   string
	header http.Header
	body   string
}

// server answers the first request (the create) with create, then each status check with the next
// entry of polls, repeating the last one once they run out.
type server struct {
	*httptest.Server
	mu       sync.Mutex
	requests []recordedRequest
}

func newServer(t *testing.T, create reply, polls ...reply) *server {
	t.Helper()
	s := &server{}
	s.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		s.mu.Lock()
		s.requests = append(s.requests, recordedRequest{r.Method, r.URL.Path, r.Header.Clone(), string(body)})
		n := len(s.requests)
		s.mu.Unlock()
		next := create
		if n > 1 {
			if len(polls) == 0 {
				t.Errorf("unexpected status check %s %s", r.Method, r.URL.Path)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			next = polls[min(n-2, len(polls)-1)]
		}
		for k, v := range next.header {
			w.Header().Set(k, v)
		}
		w.Header().Set("Content-Type", "application/json")
		if next.status == 0 {
			next.status = http.StatusOK
		}
		w.WriteHeader(next.status)
		_, _ = io.WriteString(w, next.body)
	}))
	t.Cleanup(s.Close)
	return s
}

func (s *server) recorded() []recordedRequest {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]recordedRequest(nil), s.requests...)
}

func (s *server) client() *flowsclient.Client {
	return flowsclient.NewClient(core.NewRequestOptions(
		option.WithBaseURL(s.URL),
		option.WithApiKey(elevenlabs.String("test-key")),
	))
}

func imageRequest() *types.ImageGenerationRequest {
	return &types.ImageGenerationRequest{
		ModelId:   "gpt-image-2",
		GptImage2: &types.GptImage2Request{Prompt: "A lighthouse at dawn"},
	}
}

const (
	createdGeneration    = `{"id":"gen_1","status":"pending"}`
	generatingGeneration = `{"id":"gen_1","status":"generating"}`
	completedGeneration  = `{"id":"gen_1","status":"completed","content_url":"https://example.com/out.png","content_mime_type":"image/png"}`
	failedGeneration     = `{"id":"gen_1","status":"failed","failure_reason":"moderated","error_message":"Blocked by moderation"}`
)

func run(status string) string {
	return `{"id":"run_1","template_id":"tpl_1","version_id":"ver_1","status":"` + status + `","outputs":{}}`
}

func assertSleeps(t *testing.T, got []time.Duration, want ...time.Duration) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("sleeps = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("sleeps = %v, want %v", got, want)
		}
	}
}

func TestImagePollsUntilCompletedHonouringRetryAfterSeconds(t *testing.T) {
	srv := newServer(t,
		reply{header: map[string]string{"Retry-After": "5"}, body: createdGeneration},
		reply{header: map[string]string{"Retry-After": "7"}, body: generatingGeneration},
		reply{body: completedGeneration},
	)
	pacer := newFakePacer()
	ctx := flowswait.WithPacer(context.Background(), pacer)

	result, err := srv.client().Image.CreateAndWait(ctx, imageRequest(), nil)
	if err != nil {
		t.Fatalf("CreateAndWait: %v", err)
	}
	completed := result.GetCompleted()
	if result.Status != "completed" || completed == nil || completed.ContentUrl != "https://example.com/out.png" {
		t.Fatalf("result = %+v, want the completed generation", result)
	}
	assertSleeps(t, pacer.recorded(), 5*time.Second, 7*time.Second)

	requests := srv.recorded()
	if len(requests) != 3 {
		t.Fatalf("got %d requests, want 3", len(requests))
	}
	if requests[0].method != http.MethodPost || requests[0].path != "/v1/flows/image" {
		t.Fatalf("create = %s %s", requests[0].method, requests[0].path)
	}
	var created map[string]interface{}
	if err := json.Unmarshal([]byte(requests[0].body), &created); err != nil || created["model_id"] != "gpt-image-2" {
		t.Fatalf("create body = %s", requests[0].body)
	}
	for _, r := range requests[1:] {
		if r.method != http.MethodGet || r.path != "/v1/flows/image/gen_1" {
			t.Fatalf("status check = %s %s", r.method, r.path)
		}
	}
}

func TestMissingRetryAfterWaitsTheDefaultMinPollInterval(t *testing.T) {
	srv := newServer(t, reply{body: createdGeneration}, reply{body: completedGeneration})
	pacer := newFakePacer()

	if _, err := srv.client().Image.CreateAndWait(flowswait.WithPacer(context.Background(), pacer), imageRequest(), &flows.WaitOptions{}); err != nil {
		t.Fatalf("CreateAndWait: %v", err)
	}
	assertSleeps(t, pacer.recorded(), flows.DefaultMinPollInterval)
}

func TestFailedGenerationIsReturnedNotAnError(t *testing.T) {
	srv := newServer(t, reply{body: createdGeneration}, reply{body: failedGeneration})
	ctx := flowswait.WithPacer(context.Background(), newFakePacer())

	result, err := srv.client().Image.CreateAndWait(ctx, imageRequest(), nil)
	if err != nil {
		t.Fatalf("CreateAndWait: %v", err)
	}
	failed := result.GetFailed()
	if result.Status != "failed" || failed == nil {
		t.Fatalf("result = %+v, want the failed generation", result)
	}
	if failed.FailureReason != types.MediaGenerationFailedResponseFailureReasonModerated || failed.ErrorMessage != "Blocked by moderation" {
		t.Fatalf("failure = %+v", failed)
	}
}

func TestVideoAndTextToSpeechUseTheirOwnEndpoints(t *testing.T) {
	t.Run("video", func(t *testing.T) {
		srv := newServer(t, reply{body: createdGeneration}, reply{body: completedGeneration})
		request := &types.VideoGenerationRequest{ModelId: "veo-3.1-generate-001", Veo31Generate001: &types.Veo31Request{Prompt: "Waves"}}
		result, err := srv.client().Video.CreateAndWait(flowswait.WithPacer(context.Background(), newFakePacer()), request, nil)
		if err != nil || result.Status != "completed" {
			t.Fatalf("CreateAndWait = %+v, %v", result, err)
		}
		requests := srv.recorded()
		if requests[0].path != "/v1/flows/video" || requests[1].path != "/v1/flows/video/gen_1" {
			t.Fatalf("paths = %s, %s", requests[0].path, requests[1].path)
		}
	})
	t.Run("text-to-speech", func(t *testing.T) {
		srv := newServer(t, reply{body: createdGeneration}, reply{body: completedGeneration})
		request := &types.TextToSpeechGenerationRequest{
			ModelId:  "eleven_v3",
			ElevenV3: &types.ElevenV3Request{Text: "Hello", Voice: "JBFqnCBsd6RMkjVDRZzb"},
		}
		result, err := srv.client().TextToSpeech.CreateAndWait(flowswait.WithPacer(context.Background(), newFakePacer()), request, nil)
		if err != nil || result.Status != "completed" {
			t.Fatalf("CreateAndWait = %+v, %v", result, err)
		}
		requests := srv.recorded()
		if requests[0].path != "/v1/flows/text-to-speech" || requests[1].path != "/v1/flows/text-to-speech/gen_1" {
			t.Fatalf("paths = %s, %s", requests[0].path, requests[1].path)
		}
	})
}

func TestRetryAfterAcceptsAnHTTPDate(t *testing.T) {
	pacer := newFakePacer()
	retryAt := pacer.Now().Add(9 * time.Second).Format(http.TimeFormat)
	srv := newServer(t,
		reply{header: map[string]string{"Retry-After": retryAt}, body: createdGeneration},
		reply{body: completedGeneration},
	)

	if _, err := srv.client().Image.CreateAndWait(flowswait.WithPacer(context.Background(), pacer), imageRequest(), nil); err != nil {
		t.Fatalf("CreateAndWait: %v", err)
	}
	assertSleeps(t, pacer.recorded(), 9*time.Second)
}

func TestMinPollIntervalIsAFloorUnderRetryAfter(t *testing.T) {
	srv := newServer(t,
		reply{header: map[string]string{"Retry-After": "1"}, body: createdGeneration},
		reply{header: map[string]string{"Retry-After": "20"}, body: generatingGeneration},
		reply{body: completedGeneration},
	)
	pacer := newFakePacer()

	_, err := srv.client().Image.CreateAndWait(
		flowswait.WithPacer(context.Background(), pacer),
		imageRequest(),
		&flows.WaitOptions{MinPollInterval: 4 * time.Second},
	)
	if err != nil {
		t.Fatalf("CreateAndWait: %v", err)
	}
	// Retry-After 1 is raised to the 4s floor; Retry-After 20 is above it and honoured.
	assertSleeps(t, pacer.recorded(), 4*time.Second, 20*time.Second)
}

func TestJitterAddsUpToTwentyPercent(t *testing.T) {
	srv := newServer(t,
		reply{header: map[string]string{"Retry-After": "10"}, body: createdGeneration},
		reply{body: completedGeneration},
	)
	pacer := newFakePacer()
	pacer.jitter = 0.5

	if _, err := srv.client().Image.CreateAndWait(flowswait.WithPacer(context.Background(), pacer), imageRequest(), nil); err != nil {
		t.Fatalf("CreateAndWait: %v", err)
	}
	assertSleeps(t, pacer.recorded(), 11*time.Second)
}

func TestWaitTimeoutReturnsTheGenerationID(t *testing.T) {
	srv := newServer(t, reply{body: createdGeneration}, reply{body: generatingGeneration})
	pacer := newFakePacer()

	_, err := srv.client().Image.CreateAndWait(
		flowswait.WithPacer(context.Background(), pacer),
		imageRequest(),
		&flows.WaitOptions{WaitTimeout: 10 * time.Second},
	)
	var timeoutErr *flows.WaitTimeoutError
	if !errors.As(err, &timeoutErr) {
		t.Fatalf("err = %v, want *flows.WaitTimeoutError", err)
	}
	if timeoutErr.ID != "gen_1" || timeoutErr.TemplateID != "" || timeoutErr.LastStatus != "generating" || timeoutErr.WaitTimeout != 10*time.Second {
		t.Fatalf("timeout error = %+v", timeoutErr)
	}
	if last, ok := timeoutErr.LastResponse.(*types.MediaGenerationResponse); !ok || last.Status != "generating" {
		t.Fatalf("LastResponse = %#v", timeoutErr.LastResponse)
	}
	want := "flows: generation gen_1 is still generating after 10s; call Get with this id to keep waiting"
	if err.Error() != want {
		t.Fatalf("message = %q, want %q", err.Error(), want)
	}
	// 3+3+3 and then the last sleep is cut to the 1s left; the check after it uses up the deadline.
	assertSleeps(t, pacer.recorded(), 3*time.Second, 3*time.Second, 3*time.Second, time.Second)
	if got := len(srv.recorded()); got != 5 {
		t.Fatalf("got %d requests, want 1 create and 4 status checks", got)
	}
}

func TestTemplateRunPollsUntilCompleted(t *testing.T) {
	srv := newServer(t,
		reply{header: map[string]string{"Retry-After": "6"}, body: run("pending")},
		reply{body: run("generating")},
		reply{body: run("completed")},
	)
	pacer := newFakePacer()

	result, err := srv.client().Templates.Runs.CreateAndWait(
		flowswait.WithPacer(context.Background(), pacer),
		"tpl_1",
		&templates.TemplateRunCreateRequest{Inputs: map[string]*types.TemplateRunInput{"prompt": {String: "A fox"}}},
		nil,
	)
	if err != nil {
		t.Fatalf("CreateAndWait: %v", err)
	}
	if result.Status != types.TemplateRunStatusCompleted || result.Id != "run_1" {
		t.Fatalf("result = %+v", result)
	}
	assertSleeps(t, pacer.recorded(), 6*time.Second, 3*time.Second)
	requests := srv.recorded()
	if requests[0].method != http.MethodPost || requests[0].path != "/v1/flows/templates/tpl_1/runs" {
		t.Fatalf("create = %s %s", requests[0].method, requests[0].path)
	}
	if requests[0].body != `{"inputs":{"prompt":"A fox"}}` {
		t.Fatalf("create body = %s", requests[0].body)
	}
	for _, r := range requests[1:] {
		if r.method != http.MethodGet || r.path != "/v1/flows/templates/tpl_1/runs/run_1" {
			t.Fatalf("status check = %s %s", r.method, r.path)
		}
	}
}

func TestFailedTemplateRunIsReturnedNotAnError(t *testing.T) {
	srv := newServer(t, reply{body: run("pending")}, reply{body: run("failed")})

	result, err := srv.client().Templates.Runs.CreateAndWait(
		flowswait.WithPacer(context.Background(), newFakePacer()),
		"tpl_1",
		&templates.TemplateRunCreateRequest{Inputs: map[string]*types.TemplateRunInput{}},
		nil,
	)
	if err != nil {
		t.Fatalf("CreateAndWait: %v", err)
	}
	if result.Status != types.TemplateRunStatusFailed {
		t.Fatalf("status = %s, want failed", result.Status)
	}
}

func TestTemplateRunAlreadyTerminalOnCreateSkipsPolling(t *testing.T) {
	srv := newServer(t, reply{body: run("completed")})
	pacer := newFakePacer()

	result, err := srv.client().Templates.Runs.CreateAndWait(
		flowswait.WithPacer(context.Background(), pacer),
		"tpl_1",
		&templates.TemplateRunCreateRequest{Inputs: map[string]*types.TemplateRunInput{}},
		nil,
	)
	if err != nil || result.Status != types.TemplateRunStatusCompleted {
		t.Fatalf("CreateAndWait = %+v, %v", result, err)
	}
	if len(pacer.recorded()) != 0 || len(srv.recorded()) != 1 {
		t.Fatalf("sleeps = %v, requests = %d; want no polling", pacer.recorded(), len(srv.recorded()))
	}
}

func TestTemplateRunTimeoutCarriesTheTemplateID(t *testing.T) {
	srv := newServer(t, reply{body: run("pending")}, reply{body: run("generating")})

	_, err := srv.client().Templates.Runs.CreateAndWait(
		flowswait.WithPacer(context.Background(), newFakePacer()),
		"tpl_1",
		&templates.TemplateRunCreateRequest{Inputs: map[string]*types.TemplateRunInput{}},
		&flows.WaitOptions{WaitTimeout: 5 * time.Second},
	)
	var timeoutErr *flows.WaitTimeoutError
	if !errors.As(err, &timeoutErr) {
		t.Fatalf("err = %v, want *flows.WaitTimeoutError", err)
	}
	if timeoutErr.ID != "run_1" || timeoutErr.TemplateID != "tpl_1" || timeoutErr.LastStatus != "generating" {
		t.Fatalf("timeout error = %+v", timeoutErr)
	}
	want := "flows: template run run_1 of template tpl_1 is still generating after 5s; call Templates.Runs.Get with the template id and run id to keep waiting"
	if err.Error() != want {
		t.Fatalf("message = %q, want %q", err.Error(), want)
	}
}

func TestBodyPropertiesGoToTheCreateOnlyAndHeadersToEveryRequest(t *testing.T) {
	srv := newServer(t, reply{body: createdGeneration}, reply{body: generatingGeneration}, reply{body: completedGeneration})

	_, err := srv.client().Image.CreateAndWait(
		flowswait.WithPacer(context.Background(), newFakePacer()),
		imageRequest(),
		nil,
		option.WithHTTPHeader(http.Header{"X-Trace": []string{"abc"}}),
		option.WithBodyProperties(map[string]interface{}{"extra": "only-on-create"}),
	)
	if err != nil {
		t.Fatalf("CreateAndWait: %v", err)
	}
	requests := srv.recorded()
	if len(requests) != 3 {
		t.Fatalf("got %d requests, want 3", len(requests))
	}
	var created map[string]interface{}
	if err := json.Unmarshal([]byte(requests[0].body), &created); err != nil || created["extra"] != "only-on-create" {
		t.Fatalf("create body = %s, want the extra body property", requests[0].body)
	}
	for i, r := range requests {
		if r.header.Get("X-Trace") != "abc" || r.header.Get("xi-api-key") != "test-key" {
			t.Fatalf("request %d headers = %v", i, r.header)
		}
		if i > 0 && r.body != "" {
			t.Fatalf("status check %d sent a body: %q", i, r.body)
		}
	}
}

func TestInvalidWaitOptionsAreRejectedBeforeAnyRequest(t *testing.T) {
	for name, wait := range map[string]*flows.WaitOptions{
		"poll interval below the smallest": {MinPollInterval: 100 * time.Millisecond},
		"negative poll interval":           {MinPollInterval: -time.Second},
		"negative wait timeout":            {WaitTimeout: -time.Second},
	} {
		t.Run(name, func(t *testing.T) {
			srv := newServer(t, reply{body: createdGeneration})
			_, err := srv.client().Image.CreateAndWait(context.Background(), imageRequest(), wait)
			if !errors.Is(err, flows.ErrInvalidWaitOptions) {
				t.Fatalf("err = %v, want flows.ErrInvalidWaitOptions", err)
			}
			if got := len(srv.recorded()); got != 0 {
				t.Fatalf("sent %d requests, want none", got)
			}
		})
	}
}

func TestStatusCheckErrorIsReturned(t *testing.T) {
	srv := newServer(t, reply{body: createdGeneration}, reply{status: http.StatusNotFound, body: `{"detail":"not found"}`})

	_, err := srv.client().Image.CreateAndWait(flowswait.WithPacer(context.Background(), newFakePacer()), imageRequest(), nil)
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) || apiErr.StatusCode != http.StatusNotFound {
		t.Fatalf("err = %v, want a 404 *core.APIError", err)
	}
}

// The remaining tests use the real clock.

func TestRealClockWaitsBetweenStatusChecks(t *testing.T) {
	srv := newServer(t, reply{body: createdGeneration}, reply{body: completedGeneration})

	started := time.Now()
	result, err := srv.client().Image.CreateAndWait(
		context.Background(),
		imageRequest(),
		&flows.WaitOptions{MinPollInterval: flows.SmallestMinPollInterval},
	)
	if err != nil || result.Status != "completed" {
		t.Fatalf("CreateAndWait = %+v, %v", result, err)
	}
	if elapsed := time.Since(started); elapsed < flows.SmallestMinPollInterval {
		t.Fatalf("returned after %s, want at least %s", elapsed, flows.SmallestMinPollInterval)
	}
}

func TestCancellingTheContextStopsTheWait(t *testing.T) {
	srv := newServer(t, reply{body: createdGeneration}, reply{body: generatingGeneration})
	ctx, cancel := context.WithCancel(context.Background())
	time.AfterFunc(50*time.Millisecond, cancel)

	started := time.Now()
	_, err := srv.client().Image.CreateAndWait(ctx, imageRequest(), nil)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("err = %v, want context.Canceled", err)
	}
	if elapsed := time.Since(started); elapsed >= flows.DefaultMinPollInterval {
		t.Fatalf("returned after %s, want well before the %s poll interval", elapsed, flows.DefaultMinPollInterval)
	}
	if got := len(srv.recorded()); got != 1 {
		t.Fatalf("got %d requests, want only the create", got)
	}
}

func TestContextDeadlineStopsTheWait(t *testing.T) {
	srv := newServer(t, reply{body: createdGeneration}, reply{body: generatingGeneration})
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	_, err := srv.client().Image.CreateAndWait(ctx, imageRequest(), nil)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("err = %v, want context.DeadlineExceeded", err)
	}
	var timeoutErr *flows.WaitTimeoutError
	if errors.As(err, &timeoutErr) {
		t.Fatalf("a context deadline must not be reported as a WaitTimeoutError")
	}
}
