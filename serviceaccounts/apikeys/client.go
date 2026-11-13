package apikeys

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	serviceaccounts "github.com/elevenlabs/elevenlabs-go/serviceaccounts"
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

// Get all API keys for a service account
func (c *Client) List(ctx context.Context, serviceAccountUserId string, opts ...option.RequestOption) (*types.WorkspaceApiKeyListResponseModel, error) {
	response, err := c.WithRawResponse.List(ctx, serviceAccountUserId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Create a new API key for a service account
func (c *Client) Create(ctx context.Context, serviceAccountUserId string, request *serviceaccounts.BodyCreateServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysPost, opts ...option.RequestOption) (*types.WorkspaceCreateApiKeyResponseModel, error) {
	response, err := c.WithRawResponse.Create(ctx, serviceAccountUserId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Update an existing API key for a service account
func (c *Client) Update(ctx context.Context, serviceAccountUserId string, apiKeyId string, request *serviceaccounts.EditServiceAccountApiKeyV1ServiceAccountsServiceAccountUserIdApiKeysApiKeyIdPatch, opts ...option.RequestOption) (any, error) {
	response, err := c.WithRawResponse.Update(ctx, serviceAccountUserId, apiKeyId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete an existing API key for a service account
func (c *Client) Delete(ctx context.Context, serviceAccountUserId string, apiKeyId string, opts ...option.RequestOption) (any, error) {
	response, err := c.WithRawResponse.Delete(ctx, serviceAccountUserId, apiKeyId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Disable an existing API key for a service account without changing its other settings. This stays available after a workspace downgrades off a plan that includes service accounts. Turning the key back on still requires that plan.
func (c *Client) Disable(ctx context.Context, serviceAccountUserId string, apiKeyId string, opts ...option.RequestOption) (any, error) {
	response, err := c.WithRawResponse.Disable(ctx, serviceAccountUserId, apiKeyId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
