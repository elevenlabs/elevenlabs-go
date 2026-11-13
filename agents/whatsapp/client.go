package whatsapp

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

// Make an outbound call via WhatsApp
func (c *Client) OutboundCall(ctx context.Context, request *agents.BodyMakeAnOutboundCallViaWhatsAppV1ConvaiWhatsappOutboundCallPost, opts ...option.RequestOption) (*types.WhatsAppOutboundCallResponse, error) {
	response, err := c.WithRawResponse.OutboundCall(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Send an outbound message via WhatsApp
func (c *Client) OutboundMessage(ctx context.Context, request *agents.BodySendAnOutboundMessageViaWhatsAppV1ConvaiWhatsappOutboundMessagePost, opts ...option.RequestOption) (*types.WhatsAppOutboundMessageResponse, error) {
	response, err := c.WithRawResponse.OutboundMessage(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
