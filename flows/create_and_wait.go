// Hand-written (listed in .fernignore): the options and errors of the CreateAndWait helpers on
// client.Flows.Image, .Video, .TextToSpeech and .Templates.Runs. The helpers themselves live next to
// the generated clients (create_and_wait.go in each package) and share the polling loop in
// flows/internal/flowswait. Mirrors create_and_wait in the Python and Rust SDKs, createAndWait in
// the TypeScript SDK and CreateAndWaitAsync in the .NET SDK.

package flows

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// DefaultMinPollInterval is the shortest wait between status checks when WaitOptions leaves
// MinPollInterval unset. It matches the server's shortest Retry-After (3 seconds for image, audio
// and text), so the SDK polls no faster than the server would pace it when the header is missing.
const DefaultMinPollInterval = 3 * time.Second

// SmallestMinPollInterval is the smallest MinPollInterval accepted. Below this, a missing server
// hint would turn the wait into a tight loop of status checks.
const SmallestMinPollInterval = 500 * time.Millisecond

// ErrInvalidWaitOptions is wrapped by the error CreateAndWait returns when its WaitOptions are
// rejected. Nothing has been sent when it is returned.
var ErrInvalidWaitOptions = errors.New("flows: invalid wait options")

// WaitOptions sets how long CreateAndWait waits and how often it checks. The zero value (and a nil
// *WaitOptions) waits indefinitely, checking no more often than every DefaultMinPollInterval.
type WaitOptions struct {
	// WaitTimeout bounds the whole wait. When it runs out, CreateAndWait returns a
	// *WaitTimeoutError carrying the id, so you can resume with Get. Zero waits indefinitely;
	// negative values are rejected. Not to be confused with a per-request timeout on the HTTP
	// client or a deadline on ctx, which stop the wait with the context's error instead.
	WaitTimeout time.Duration
	// MinPollInterval is the shortest wait between status checks, and the wait used when the
	// server sends no Retry-After header. Zero means DefaultMinPollInterval; otherwise it must be
	// at least SmallestMinPollInterval.
	MinPollInterval time.Duration
}

// WaitTimeoutError is returned by CreateAndWait when a generation or template run is still
// unfinished after WaitOptions.WaitTimeout. Reach it with errors.As.
//
// The work keeps running on the server. Resume waiting by calling Get with ID (and TemplateID for
// a template run), or let a webhook deliver the result.
type WaitTimeoutError struct {
	// ID is the generation id, or the run id for a template run.
	ID string
	// TemplateID is the template id for a template run, and empty for a generation.
	TemplateID string
	// WaitTimeout is the WaitOptions.WaitTimeout that ran out.
	WaitTimeout time.Duration
	// LastStatus is the status of the last response received, such as "pending" or "generating".
	LastStatus string
	// LastResponse is the last response received: the create response
	// (*types.MediaGenerationCreateResponse or *types.TemplateRunResponse) if no status check
	// ran, otherwise the last *types.MediaGenerationResponse or *types.TemplateRunResponse.
	LastResponse interface{}
}

func (e *WaitTimeoutError) Error() string {
	after := strconv.FormatFloat(e.WaitTimeout.Seconds(), 'f', -1, 64) + "s"
	if e.TemplateID != "" {
		return fmt.Sprintf(
			"flows: template run %s of template %s is still %s after %s; call Templates.Runs.Get with the template id and run id to keep waiting",
			e.ID, e.TemplateID, e.LastStatus, after,
		)
	}
	return fmt.Sprintf(
		"flows: generation %s is still %s after %s; call Get with this id to keep waiting",
		e.ID, e.LastStatus, after,
	)
}
