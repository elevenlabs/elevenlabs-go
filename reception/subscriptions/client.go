package subscriptions

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	reception "github.com/elevenlabs/elevenlabs-go/reception"
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

// Subscribe a reseller-attributed workspace to one of the reseller's tiers
func (c *Client) Assign(ctx context.Context, workspaceId string, request *reception.AssignResellerTierRequest, opts ...option.RequestOption) (map[string]string, error) {
	response, err := c.WithRawResponse.Assign(ctx, workspaceId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Move a reseller-attributed workspace between the reseller's tiers. Upgrades apply immediately; downgrades and lateral moves take effect at the end of the billing period.
func (c *Client) Change(ctx context.Context, workspaceId string, request *reception.ChangeResellerTierRequest, opts ...option.RequestOption) (map[string]string, error) {
	response, err := c.WithRawResponse.Change(ctx, workspaceId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// End a reseller-attributed workspace's subscription at the end of the current billing period
func (c *Client) Unassign(ctx context.Context, workspaceId string, opts ...option.RequestOption) (map[string]string, error) {
	response, err := c.WithRawResponse.Unassign(ctx, workspaceId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Undo a pending unassignment so the subscription renews normally
func (c *Client) Restore(ctx context.Context, workspaceId string, opts ...option.RequestOption) (map[string]string, error) {
	response, err := c.WithRawResponse.Restore(ctx, workspaceId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Drop a scheduled downgrade, leaving the workspace on its current tier
func (c *Client) CancelPendingDowngrade(ctx context.Context, workspaceId string, opts ...option.RequestOption) (map[string]string, error) {
	response, err := c.WithRawResponse.CancelPendingDowngrade(ctx, workspaceId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
