package client

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	audio "github.com/elevenlabs/elevenlabs-go/dubbing/audio"
	projectclient "github.com/elevenlabs/elevenlabs-go/dubbing/project/client"
	transcripts "github.com/elevenlabs/elevenlabs-go/dubbing/transcripts"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

type Client struct {
	WithRawResponse *RawClient

	Project     *projectclient.Client
	Audio       *audio.Client
	Transcripts *transcripts.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse: NewRawClient(options),
		Project:         projectclient.NewClient(options),
		Audio:           audio.NewClient(options),
		Transcripts:     transcripts.NewClient(options),
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

// List the dubs you have access to.
func (c *Client) List(ctx context.Context, request *types.ListDubbingRequest, opts ...option.RequestOption) (*types.DubbingMetadataPageResponseModel, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Dubs a provided audio or video file into given language.
func (c *Client) Create(ctx context.Context, dubRequest *types.DubRequest, opts ...option.RequestOption) (*types.DoDubbingResponse, error) {
	response, err := c.WithRawResponse.Create(ctx, dubRequest, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Returns metadata about a dubbing project, including whether it's still in progress or not
func (c *Client) Get(ctx context.Context, dubbingId string, opts ...option.RequestOption) (*types.DubbingMetadataResponse, error) {
	response, err := c.WithRawResponse.Get(ctx, dubbingId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Deletes a dubbing project.
func (c *Client) Delete(ctx context.Context, dubbingId string, opts ...option.RequestOption) (*types.DeleteDubbingResponseModel, error) {
	response, err := c.WithRawResponse.Delete(ctx, dubbingId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
