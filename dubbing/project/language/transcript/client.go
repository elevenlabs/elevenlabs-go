package transcript

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	language "github.com/elevenlabs/elevenlabs-go/dubbing/project/language"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

type Client struct {
	WithRawResponse *RawClient

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse: NewRawClient(options),
		options:         options,
		baseURL:         options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:         options.HTTPClient,
				MaxAttempts:    options.MaxAttempts,
				DisableRetries: options.DisableRetries,
			},
		),
	}
}

// A language target's transcript: source segments with their translations. Available once the target has produced an output. Returns a conflict while the target is still on its first dub, since it has no translations to return yet.
func (c *Client) Get(ctx context.Context, projectId string, languageId string, opts ...option.RequestOption) (*types.DubbingTargetTranscriptResponse, error) {
	response, err := c.WithRawResponse.Get(ctx, projectId, languageId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Enterprise only. Edit a segment's translation for a language target. Omitted fields are left unchanged; an explicit null clears the field. Bumps the target's `revision` and marks it `stale` if it had already completed. The source transcript and the project's other languages are untouched, and no audio changes until you regenerate the target.
func (c *Client) UpdateSegment(ctx context.Context, projectId string, languageId string, segmentId string, request *types.DubbingTargetSegmentUpdateRequest, opts ...option.RequestOption) (*types.DubbingTargetSegmentUpdateResponse, error) {
	response, err := c.WithRawResponse.UpdateSegment(ctx, projectId, languageId, segmentId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Enterprise only. Edit several segments' translations for a language target in one atomic request: every edit applies or none does. Bumps the target's `revision` and marks it `stale` if it had already completed. The source transcript and the project's other languages are untouched, and no audio changes until you regenerate the target.
func (c *Client) UpdateSegments(ctx context.Context, projectId string, languageId string, request *language.DubbingBulkTargetSegmentUpdateRequest, opts ...option.RequestOption) (*types.DubbingBulkTargetSegmentUpdateResponse, error) {
	response, err := c.WithRawResponse.UpdateSegments(ctx, projectId, languageId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Enterprise only. Re-dub a target from its edited transcript, re-synthesizing only the edited regions (charged like a generation, less the free-regeneration allowance). Accepted asynchronously: the target returns to `processing` and sends a `dubbing_language_completed` event to the project's `webhook_ids` when the re-dub lands, carrying the new output URLs. Returns a conflict when the target has no edits to apply — nothing is dispatched and nothing is charged.
func (c *Client) Regenerate(ctx context.Context, projectId string, languageId string, opts ...option.RequestOption) (*types.DubbingRegenerateResponse, error) {
	response, err := c.WithRawResponse.Regenerate(ctx, projectId, languageId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
