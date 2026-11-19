// Hand-written (listed in .fernignore): adds CreateAndWait to the generated template runs client,
// so it is reachable as client.Flows.Templates.Runs.CreateAndWait. Go lets a package's methods
// live in any of its files, so the generated client.go is untouched. The polling loop is in
// flows/internal/flowswait.

package runs

import (
	context "context"
	http "net/http"

	flows "github.com/elevenlabs/elevenlabs-go/flows"
	flowswait "github.com/elevenlabs/elevenlabs-go/flows/internal/flowswait"
	templates "github.com/elevenlabs/elevenlabs-go/flows/templates"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

// CreateAndWait starts a run of a flows template and waits until its Status is completed or
// failed, then returns the final run.
//
// Between status checks it waits as long as the server's Retry-After header suggests (whole
// seconds or an HTTP date), never less than wait.MinPollInterval, plus up to 20% jitter. A run
// that is already completed or failed when created is returned without polling. This is meant for
// scripts and quickstarts: in production, set Webhook on the request and handle the
// flows_template_run webhook event instead of holding a process open.
//
// A failed run is returned with a nil error, not as an error: check Status on the result and on
// each of its Outputs.
//
// A nil wait uses the defaults: no wait timeout and a DefaultMinPollInterval floor. When
// wait.WaitTimeout runs out the error is a *flows.WaitTimeoutError carrying the run id and the
// template id, for Get. Cancelling ctx stops the wait with an error wrapping ctx.Err(). Invalid
// wait options return an error wrapping flows.ErrInvalidWaitOptions before any request is sent.
// API errors from the create call or a status check are returned as they are.
//
// opts apply to the create call and to every status check, except option.WithBodyProperties,
// which only the create call gets.
func (c *Client) CreateAndWait(
	ctx context.Context,
	templateId string,
	request *templates.TemplateRunCreateRequest,
	wait *flows.WaitOptions,
	opts ...option.RequestOption,
) (*types.TemplateRunResponse, error) {
	waiter, err := flowswait.Begin(ctx, wait)
	if err != nil {
		return nil, err
	}
	created, err := c.WithRawResponse.Create(ctx, templateId, request, opts...)
	if err != nil {
		return nil, err
	}
	status := string(created.Body.GetStatus())
	if flowswait.IsTerminal(status) {
		return created.Body, nil
	}
	runID := created.Body.GetId()
	pollOpts := flowswait.PollOptions(opts)
	return flowswait.Poll(
		ctx,
		waiter,
		flowswait.Target{ID: runID, TemplateID: templateId},
		status,
		created.Body,
		created.Header,
		func(ctx context.Context) (*types.TemplateRunResponse, string, http.Header, error) {
			response, err := c.WithRawResponse.Get(ctx, templateId, runID, pollOpts...)
			if err != nil {
				return nil, "", nil, err
			}
			return response.Body, string(response.Body.GetStatus()), response.Header, nil
		},
	)
}
