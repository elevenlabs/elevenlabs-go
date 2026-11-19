package toolconfigs

import (
	context "context"
	mcpservers "github.com/elevenlabs/elevenlabs-go/agents/mcpservers"
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

// Create configuration overrides for a specific MCP tool.
func (c *Client) Create(ctx context.Context, mcpServerId string, request *mcpservers.McpToolConfigOverrideCreateRequestModel, opts ...option.RequestOption) (*types.McpServerResponseModel, error) {
	response, err := c.WithRawResponse.Create(ctx, mcpServerId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Retrieve configuration overrides for a specific MCP tool.
func (c *Client) Get(ctx context.Context, mcpServerId string, toolName string, opts ...option.RequestOption) (*types.McpToolConfigOverrideOutput, error) {
	response, err := c.WithRawResponse.Get(ctx, mcpServerId, toolName, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Update configuration overrides for a specific MCP tool.
func (c *Client) Update(ctx context.Context, mcpServerId string, toolName string, request *mcpservers.McpToolConfigOverrideUpdateRequestModel, opts ...option.RequestOption) (*types.McpServerResponseModel, error) {
	response, err := c.WithRawResponse.Update(ctx, mcpServerId, toolName, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Remove configuration overrides for a specific MCP tool.
func (c *Client) Delete(ctx context.Context, mcpServerId string, toolName string, opts ...option.RequestOption) (*types.McpServerResponseModel, error) {
	response, err := c.WithRawResponse.Delete(ctx, mcpServerId, toolName, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
