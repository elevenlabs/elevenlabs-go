package livecount

import (
	context "context"
	analytics "github.com/elevenlabs/elevenlabs-go/agents/analytics"
	core "github.com/elevenlabs/elevenlabs-go/core"
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

// Get the live count of the ongoing conversations.
func (c *Client) Get(ctx context.Context, request *analytics.GetLiveCountRequest, opts ...option.RequestOption) (*types.GetLiveCountResponse, error) {
	response, err := c.WithRawResponse.Get(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
