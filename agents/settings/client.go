package settings

import (
	context "context"
	agents "github.com/elevenlabs/elevenlabs-go/agents"
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

// Retrieve Convai settings for the workspace
func (c *Client) Get(ctx context.Context, opts ...option.RequestOption) (*types.GetConvAiSettingsResponseModel, error) {
	response, err := c.WithRawResponse.Get(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Update Convai settings for the workspace
func (c *Client) Update(ctx context.Context, request *agents.PatchConvAiSettingsRequest, opts ...option.RequestOption) (*types.GetConvAiSettingsResponseModel, error) {
	response, err := c.WithRawResponse.Update(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
