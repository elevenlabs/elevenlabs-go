// Hand-written (listed in .fernignore): adds CreateAndWait to the generated video client, so it is
// reachable as client.Flows.Video.CreateAndWait. Go lets a package's methods live in any of its
// files, so the generated client.go is untouched. The polling loop is in flows/internal/flowswait.

package video

import (
	context "context"
	http "net/http"

	flows "github.com/elevenlabs/elevenlabs-go/flows"
	flowswait "github.com/elevenlabs/elevenlabs-go/flows/internal/flowswait"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

// CreateAndWait starts a video generation and waits until it is completed or failed,
// then returns the final generation.
//
// Between status checks it waits as long as the server's Retry-After header suggests (whole
// seconds or an HTTP date), never less than wait.MinPollInterval, plus up to 20% jitter. This is
// meant for scripts and quickstarts: in production, set Webhook on the request and handle the
// webhook event instead of holding a process open.
//
// The result's Status is always "completed" or "failed". A failed generation is returned with a
// nil error, not as an error: check GetFailed() (FailureReason, ErrorMessage) or GetCompleted()
// (ContentUrl) on the result.
//
// A nil wait uses the defaults: no wait timeout and a DefaultMinPollInterval floor. When
// wait.WaitTimeout runs out the error is a *flows.WaitTimeoutError carrying the generation id, for
// Get. Cancelling ctx stops the wait with an error wrapping ctx.Err(). Invalid wait options return
// an error wrapping flows.ErrInvalidWaitOptions before any request is sent. API errors from the
// create call or a status check are returned as they are.
//
// opts apply to the create call and to every status check, except option.WithBodyProperties,
// which only the create call gets.
func (c *Client) CreateAndWait(
	ctx context.Context,
	request *types.VideoGenerationRequest,
	wait *flows.WaitOptions,
	opts ...option.RequestOption,
) (*types.MediaGenerationResponse, error) {
	waiter, err := flowswait.Begin(ctx, wait)
	if err != nil {
		return nil, err
	}
	created, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	generationID := created.Body.GetId()
	pollOpts := flowswait.PollOptions(opts)
	return flowswait.Poll(
		ctx,
		waiter,
		flowswait.Target{ID: generationID},
		created.Body.Status(),
		created.Body,
		created.Header,
		func(ctx context.Context) (*types.MediaGenerationResponse, string, http.Header, error) {
			response, err := c.WithRawResponse.Get(ctx, generationID, pollOpts...)
			if err != nil {
				return nil, "", nil, err
			}
			return response.Body, response.Body.GetStatus(), response.Header, nil
		},
	)
}
