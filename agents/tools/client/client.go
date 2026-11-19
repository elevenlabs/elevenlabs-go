package client

import (
	context "context"
	agents "github.com/elevenlabs/elevenlabs-go/agents"
	executions "github.com/elevenlabs/elevenlabs-go/agents/tools/executions"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

type Client struct {
	WithRawResponse *RawClient

	Executions *executions.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse: NewRawClient(options),
		Executions:      executions.NewClient(options),
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

// Get all available tools in the workspace.
func (c *Client) List(ctx context.Context, request *agents.ListToolsRequest, opts ...option.RequestOption) (*types.ToolsResponseModel, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Add a new tool to the available tools in the workspace.
func (c *Client) Create(ctx context.Context, request *types.ToolRequestModel, opts ...option.RequestOption) (*types.ToolResponseModel, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get tool that is available in the workspace.
func (c *Client) Get(ctx context.Context, toolId string, request *agents.GetToolsRequest, opts ...option.RequestOption) (*types.ToolResponseModel, error) {
	response, err := c.WithRawResponse.Get(ctx, toolId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Update tool that is available in the workspace.
func (c *Client) Update(ctx context.Context, toolId string, request *types.ToolRequestModel, opts ...option.RequestOption) (*types.ToolResponseModel, error) {
	response, err := c.WithRawResponse.Update(ctx, toolId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete tool from the workspace.
func (c *Client) Delete(ctx context.Context, toolId string, request *agents.DeleteToolsRequest, opts ...option.RequestOption) (any, error) {
	response, err := c.WithRawResponse.Delete(ctx, toolId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get a list of agents depending on this tool
func (c *Client) GetDependentAgents(ctx context.Context, toolId string, request *agents.GetDependentAgentsToolsRequest, opts ...option.RequestOption) (*types.GetToolDependentAgentsResponseModel, error) {
	response, err := c.WithRawResponse.GetDependentAgents(ctx, toolId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
