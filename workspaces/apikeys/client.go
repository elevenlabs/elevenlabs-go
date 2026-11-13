package apikeys

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	workspaces "github.com/elevenlabs/elevenlabs-go/workspaces"
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

// Disable the API key used to authenticate this request. Requires the query parameter `api_key_name=self` as an explicit confirmation.
func (c *Client) Disable(ctx context.Context, request *workspaces.DisableApiKeysRequest, opts ...option.RequestOption) (any, error) {
	response, err := c.WithRawResponse.Disable(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
