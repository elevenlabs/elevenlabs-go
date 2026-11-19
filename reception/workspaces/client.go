package workspaces

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	reception "github.com/elevenlabs/elevenlabs-go/reception"
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

// List workspaces attributed to the caller's reseller, with cursor-based pagination
func (c *Client) List(ctx context.Context, request *reception.ListWorkspacesRequest, opts ...option.RequestOption) (*types.ResellerWorkspacesResponse, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Create a workspace for one of your customers, optionally with a tier, a phone number and team invites.
func (c *Client) Create(ctx context.Context, request *reception.ProvisionResellerWorkspaceRequest, opts ...option.RequestOption) (*types.ProvisionResellerWorkspaceResponse, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Finish setting up the workspace: create its receptionist and assistant from the business profile, reusing a receptionist from a setup already started, and connect your phone number, if you supply one. Safe to retry.
func (c *Client) CompleteOnboarding(ctx context.Context, workspaceId string, opts ...option.RequestOption) (*types.ResellerWorkspaceOnboardingResponse, error) {
	response, err := c.WithRawResponse.CompleteOnboarding(ctx, workspaceId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
