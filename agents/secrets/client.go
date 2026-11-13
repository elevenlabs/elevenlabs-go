package secrets

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

// Get all workspace secrets for the user
func (c *Client) List(ctx context.Context, request *agents.ListSecretsRequest, opts ...option.RequestOption) (*types.GetWorkspaceSecretsResponseModel, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Create a new secret for the workspace
func (c *Client) Create(ctx context.Context, request *agents.PostWorkspaceSecretRequest, opts ...option.RequestOption) (*types.PostWorkspaceSecretResponseModel, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get a workspace secret by ID
func (c *Client) Get(ctx context.Context, secretId string, opts ...option.RequestOption) (*types.ConvAiWorkspaceStoredSecretConfig, error) {
	response, err := c.WithRawResponse.Get(ctx, secretId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Update an existing secret for the workspace
func (c *Client) Update(ctx context.Context, secretId string, request *agents.PatchWorkspaceSecretRequest, opts ...option.RequestOption) (*types.PostWorkspaceSecretResponseModel, error) {
	response, err := c.WithRawResponse.Update(ctx, secretId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete a workspace secret if it's not in use
func (c *Client) Delete(ctx context.Context, secretId string, opts ...option.RequestOption) error {
	_, err := c.WithRawResponse.Delete(ctx, secretId, opts...)
	return err
}

// Get paginated list of resources that depend on a specific secret, filtered by resource type.
func (c *Client) GetDependencies(ctx context.Context, secretId string, resourceType *types.SecretDependencyResourceType, request *agents.GetDependenciesSecretsRequest, opts ...option.RequestOption) (*types.GetSecretDependenciesResponseModel, error) {
	response, err := c.WithRawResponse.GetDependencies(ctx, secretId, resourceType, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
