package client

import (
	context "context"
	agents "github.com/elevenlabs/elevenlabs-go/agents"
	avatar "github.com/elevenlabs/elevenlabs-go/agents/widget/avatar"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

type Client struct {
	WithRawResponse *RawClient

	Avatar *avatar.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse: NewRawClient(options),
		Avatar:          avatar.NewClient(options),
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

// Retrieve the widget configuration for an agent
func (c *Client) Get(ctx context.Context, agentId string, request *agents.GetWidgetRequest, opts ...option.RequestOption) (*types.GetAgentEmbedResponseModel, error) {
	response, err := c.WithRawResponse.Get(ctx, agentId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
