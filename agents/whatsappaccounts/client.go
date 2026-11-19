package whatsappaccounts

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

// Get a WhatsApp account
func (c *Client) Get(ctx context.Context, phoneNumberId string, opts ...option.RequestOption) (*types.GetWhatsAppAccountResponse, error) {
	response, err := c.WithRawResponse.Get(ctx, phoneNumberId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Update a WhatsApp account
func (c *Client) Update(ctx context.Context, phoneNumberId string, request *agents.UpdateWhatsAppAccountRequest, opts ...option.RequestOption) (any, error) {
	response, err := c.WithRawResponse.Update(ctx, phoneNumberId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete a WhatsApp account
func (c *Client) Delete(ctx context.Context, phoneNumberId string, opts ...option.RequestOption) (any, error) {
	response, err := c.WithRawResponse.Delete(ctx, phoneNumberId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// List all WhatsApp accounts
func (c *Client) List(ctx context.Context, request *agents.ListWhatsappAccountsRequest, opts ...option.RequestOption) (*types.ListWhatsAppAccountsResponse, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
