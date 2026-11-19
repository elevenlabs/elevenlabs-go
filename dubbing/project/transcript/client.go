package transcript

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	project "github.com/elevenlabs/elevenlabs-go/dubbing/project"
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

// The project's source transcript, as editable segments. Available once the project is `ready`.
func (c *Client) Get(ctx context.Context, projectId string, opts ...option.RequestOption) (*types.DubbingSourceTranscriptResponse, error) {
	response, err := c.WithRawResponse.Get(ctx, projectId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Enterprise only. Edit a source segment's text, speaker, or timing. Omitted fields are left unchanged. Bumps the project's `revision`, discards the affected translations in every language target, and marks any target that had already completed `stale`. No audio changes until you regenerate a target.
func (c *Client) UpdateSegment(ctx context.Context, projectId string, segmentId string, request *types.DubbingSegmentUpdateRequest, opts ...option.RequestOption) (*types.DubbingSourceSegmentUpdateResponse, error) {
	response, err := c.WithRawResponse.UpdateSegment(ctx, projectId, segmentId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Enterprise only. Remove a source segment from the transcript so it is no longer dubbed. Bumps the project's `revision`, discards the affected translations in every language target, and marks any target that had already completed `stale`. No audio changes until you regenerate a target.
func (c *Client) DeleteSegment(ctx context.Context, projectId string, segmentId string, opts ...option.RequestOption) (*types.DubbingTranscriptRevisionResponse, error) {
	response, err := c.WithRawResponse.DeleteSegment(ctx, projectId, segmentId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Enterprise only. Edit several source segments' text, speaker, or timing in one atomic request: every edit applies or none does. Bumps the project's `revision`, discards the affected translations in every language target, and marks any target that had already completed `stale`. No audio changes until you regenerate a target.
func (c *Client) UpdateSegments(ctx context.Context, projectId string, request *project.DubbingBulkSegmentUpdateRequest, opts ...option.RequestOption) (*types.DubbingBulkSourceSegmentUpdateResponse, error) {
	response, err := c.WithRawResponse.UpdateSegments(ctx, projectId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Enterprise only. Add a new source segment to the transcript. Its span must lie within the source media, last between 0.1 and 25 seconds, and not overlap another segment by the same speaker. Bumps the project's `revision`, discards the affected translations in every language target, and marks any target that had already completed `stale`. No audio changes until you regenerate a target.
func (c *Client) CreateSegment(ctx context.Context, projectId string, request *project.DubbingSegmentCreateRequest, opts ...option.RequestOption) (*types.DubbingSourceSegmentUpdateResponse, error) {
	response, err := c.WithRawResponse.CreateSegment(ctx, projectId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
