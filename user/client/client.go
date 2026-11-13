package client

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
	subscription "github.com/elevenlabs/elevenlabs-go/user/subscription"
)

type Client struct {
	WithRawResponse *RawClient

	Subscription *subscription.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse: NewRawClient(options),
		Subscription:    subscription.NewClient(options),
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

// Gets information about the user
func (c *Client) Get(ctx context.Context, opts ...option.RequestOption) (*types.User, error) {
	response, err := c.WithRawResponse.Get(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
