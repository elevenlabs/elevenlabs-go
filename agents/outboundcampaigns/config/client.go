package config

import (
	context "context"
	outboundcampaigns "github.com/elevenlabs/elevenlabs-go/agents/outboundcampaigns"
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

// Get the workspace's outbound campaign configuration: the default policy new campaigns take, and per-agent overrides.
func (c *Client) Get(ctx context.Context, opts ...option.RequestOption) (*types.WorkspaceConfigResponse, error) {
	response, err := c.WithRawResponse.Get(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Create or replace the workspace's outbound campaign configuration. Send the `version` you last read as `expected_version`, or omit it to create the first configuration. A configuration changed since you read it is not replaced. Campaigns keep the policy they started with.
func (c *Client) Update(ctx context.Context, request *outboundcampaigns.ReplaceWorkspaceConfigRequest, opts ...option.RequestOption) (*types.WorkspaceConfigResponse, error) {
	response, err := c.WithRawResponse.Update(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
