package authconnections

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
	workspace "github.com/elevenlabs/elevenlabs-go/workspace"
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

// Get all auth connections for the workspace
func (c *Client) List(ctx context.Context, opts ...option.RequestOption) (*types.ListAuthConnectionsResponse, error) {
	response, err := c.WithRawResponse.List(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Create a new OAuth2 auth connection for the workspace
func (c *Client) Create(ctx context.Context, request *workspace.CreateAuthConnectionsRequestBody, opts ...option.RequestOption) (*workspace.CreateAuthConnectionsResponse, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Update an auth connection
func (c *Client) Update(ctx context.Context, authConnectionId string, request *workspace.UpdateAuthConnectionsRequestBody, opts ...option.RequestOption) (*workspace.UpdateAuthConnectionsResponse, error) {
	response, err := c.WithRawResponse.Update(ctx, authConnectionId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete an auth connection
func (c *Client) Delete(ctx context.Context, authConnectionId string, opts ...option.RequestOption) (any, error) {
	response, err := c.WithRawResponse.Delete(ctx, authConnectionId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
