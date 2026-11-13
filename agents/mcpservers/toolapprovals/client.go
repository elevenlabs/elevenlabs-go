package toolapprovals

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

// Add approval for a specific MCP tool when using per-tool approval mode.
func (c *Client) Create(ctx context.Context, mcpServerId string, request *mcpservers.McpToolAddApprovalRequestModel, opts ...option.RequestOption) (*types.McpServerResponseModel, error) {
	response, err := c.WithRawResponse.Create(ctx, mcpServerId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Remove approval for a specific MCP tool when using per-tool approval mode.
func (c *Client) Delete(ctx context.Context, mcpServerId string, toolName string, opts ...option.RequestOption) (*types.McpServerResponseModel, error) {
	response, err := c.WithRawResponse.Delete(ctx, mcpServerId, toolName, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
