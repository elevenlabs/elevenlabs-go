package environmentvariables

import (
	context "context"
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

// List all environment variables for the workspace with optional filtering
func (c *Client) List(ctx context.Context, request *types.ListEnvironmentVariablesRequest, opts ...option.RequestOption) (*types.EnvironmentVariablesListResponse, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Create a new environment variable for the workspace
func (c *Client) Create(ctx context.Context, createEnvironmentVariablesRequestBody *types.CreateEnvironmentVariablesRequestBody, opts ...option.RequestOption) (*types.EnvironmentVariableResponse, error) {
	response, err := c.WithRawResponse.Create(ctx, createEnvironmentVariablesRequestBody, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get a specific environment variable by ID
func (c *Client) Get(ctx context.Context, envVarId string, opts ...option.RequestOption) (*types.EnvironmentVariableResponse, error) {
	response, err := c.WithRawResponse.Get(ctx, envVarId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Replace an environment variable's values. Use null to remove its value for an environment (except production); delete an environment with DELETE /v1/convai/environments/{environment}.
func (c *Client) Update(ctx context.Context, envVarId string, request *types.UpdateEnvironmentVariableRequest, opts ...option.RequestOption) (*types.EnvironmentVariableResponse, error) {
	response, err := c.WithRawResponse.Update(ctx, envVarId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete an environment variable and its values in all environments. The delete is rejected while the latest version of an unarchived agent branch, a tool, an MCP server or the workspace conversation initiation webhook uses the variable. Archived branches and earlier versions are not checked.
func (c *Client) Delete(ctx context.Context, envVarId string, opts ...option.RequestOption) error {
	_, err := c.WithRawResponse.Delete(ctx, envVarId, opts...)
	return err
}
