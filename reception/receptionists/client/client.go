package client

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	reception "github.com/elevenlabs/elevenlabs-go/reception"
	procedures "github.com/elevenlabs/elevenlabs-go/reception/receptionists/procedures"
	rules "github.com/elevenlabs/elevenlabs-go/reception/receptionists/rules"
	transferrules "github.com/elevenlabs/elevenlabs-go/reception/receptionists/transferrules"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

type Client struct {
	WithRawResponse *RawClient

	Rules         *rules.Client
	Procedures    *procedures.Client
	TransferRules *transferrules.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse: NewRawClient(options),
		Rules:           rules.NewClient(options),
		Procedures:      procedures.NewClient(options),
		TransferRules:   transferrules.NewClient(options),
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

// List every customer-facing receptionist for the workspace.
func (c *Client) List(ctx context.Context, request *reception.ListReceptionistsRequest, opts ...option.RequestOption) ([]*types.ReceptionCustomerFacingAgentResponse, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Update a receptionist's display name, config and/or widget appearance.
func (c *Client) Update(ctx context.Context, agentId string, request *reception.ReceptionUpdateCustomerFacingAgentRequest, opts ...option.RequestOption) (*types.ReceptionCustomerFacingAgentResponse, error) {
	response, err := c.WithRawResponse.Update(ctx, agentId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Re-activate a deactivated receptionist. Fails with `quota_exceeded` when the plan has no free receptionist slot.
func (c *Client) Activate(ctx context.Context, agentId string, request *reception.ActivateReceptionistsRequest, opts ...option.RequestOption) (*types.ReceptionCustomerFacingAgentResponse, error) {
	response, err := c.WithRawResponse.Activate(ctx, agentId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
