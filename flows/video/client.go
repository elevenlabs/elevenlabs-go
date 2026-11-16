package video

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	flows "github.com/elevenlabs/elevenlabs-go/flows"
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

// List the video generations created through this API, newest first.
func (c *Client) List(ctx context.Context, request *flows.ListVideoRequest, opts ...option.RequestOption) (*types.MediaGenerationListResponse, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Start a video generation with the selected model.
func (c *Client) Create(ctx context.Context, videoGenerationRequest *types.VideoGenerationRequest, opts ...option.RequestOption) (*types.MediaGenerationCreateResponse, error) {
	response, err := c.WithRawResponse.Create(ctx, videoGenerationRequest, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Retrieve the status of a video generation, and retrieve its output URL once completed.
func (c *Client) Get(ctx context.Context, generationId string, opts ...option.RequestOption) (*types.MediaGenerationResponse, error) {
	response, err := c.WithRawResponse.Get(ctx, generationId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
