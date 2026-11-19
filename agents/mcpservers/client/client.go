package client

import (
	context "context"
	agents "github.com/elevenlabs/elevenlabs-go/agents"
	toolapprovals "github.com/elevenlabs/elevenlabs-go/agents/mcpservers/toolapprovals"
	toolconfigs "github.com/elevenlabs/elevenlabs-go/agents/mcpservers/toolconfigs"
	tools "github.com/elevenlabs/elevenlabs-go/agents/mcpservers/tools"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

type Client struct {
	WithRawResponse *RawClient

	Tools         *tools.Client
	ToolApprovals *toolapprovals.Client
	ToolConfigs   *toolconfigs.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse: NewRawClient(options),
		Tools:           tools.NewClient(options),
		ToolApprovals:   toolapprovals.NewClient(options),
		ToolConfigs:     toolconfigs.NewClient(options),
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

// Retrieve all MCP server configurations available in the workspace.
func (c *Client) List(ctx context.Context, opts ...option.RequestOption) (*types.McpServersResponseModel, error) {
	response, err := c.WithRawResponse.List(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Create a new MCP server configuration in the workspace.
func (c *Client) Create(ctx context.Context, request *agents.McpServerRequestModel, opts ...option.RequestOption) (*types.McpServerResponseModel, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Retrieve a specific MCP server configuration from the workspace.
func (c *Client) Get(ctx context.Context, mcpServerId string, opts ...option.RequestOption) (*types.McpServerResponseModel, error) {
	response, err := c.WithRawResponse.Get(ctx, mcpServerId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Update the configuration settings for an MCP server.
func (c *Client) Update(ctx context.Context, mcpServerId string, request *agents.McpServerConfigUpdateRequestModel, opts ...option.RequestOption) (*types.McpServerResponseModel, error) {
	response, err := c.WithRawResponse.Update(ctx, mcpServerId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete a specific MCP server configuration from the workspace.
func (c *Client) Delete(ctx context.Context, mcpServerId string, opts ...option.RequestOption) (any, error) {
	response, err := c.WithRawResponse.Delete(ctx, mcpServerId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
