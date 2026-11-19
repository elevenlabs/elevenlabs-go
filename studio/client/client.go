package client

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	projectsclient "github.com/elevenlabs/elevenlabs-go/studio/projects/client"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

type Client struct {
	WithRawResponse *RawClient

	Projects *projectsclient.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse: NewRawClient(options),
		Projects:        projectsclient.NewClient(options),
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

// Create and auto-convert a podcast project. Currently, the LLM cost is covered by us but you will still be charged for the audio generation. In the future, you will be charged for both the LLM and audio generation costs.
func (c *Client) CreatePodcast(ctx context.Context, request *types.BodyCreatePodcastV1StudioPodcastsPost, opts ...option.RequestOption) (*types.PodcastProjectResponseModel, error) {
	response, err := c.WithRawResponse.CreatePodcast(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
