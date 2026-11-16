package tools

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

// Retrieve all tools available for a specific MCP server configuration.
func (c *Client) List(ctx context.Context, mcpServerId string, request *mcpservers.ListToolsRequest, opts ...option.RequestOption) (*types.ListMcpToolsResponseModel, error) {
	response, err := c.WithRawResponse.List(ctx, mcpServerId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
