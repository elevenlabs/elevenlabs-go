package webhooks

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

// List all webhooks for a workspace
func (c *Client) List(ctx context.Context, request *types.ListWebhooksRequest, opts ...option.RequestOption) (*types.WorkspaceWebhookListResponseModel, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Create a new webhook for the workspace with the specified authentication type.
func (c *Client) Create(ctx context.Context, request *types.BodyCreateWorkspaceWebhookV1WorkspaceWebhooksPost, opts ...option.RequestOption) (*types.WorkspaceCreateWebhookResponseModel, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Update the specified workspace webhook
func (c *Client) Update(ctx context.Context, webhookId string, request *types.BodyUpdateWorkspaceWebhookV1WorkspaceWebhooksWebhookIdPatch, opts ...option.RequestOption) (*types.PatchWorkspaceWebhookResponseModel, error) {
	response, err := c.WithRawResponse.Update(ctx, webhookId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete the specified workspace webhook
func (c *Client) Delete(ctx context.Context, webhookId string, opts ...option.RequestOption) (*types.DeleteWorkspaceWebhookResponseModel, error) {
	response, err := c.WithRawResponse.Delete(ctx, webhookId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
