package twilio

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

// Handle an outbound call via Twilio
func (c *Client) OutboundCall(ctx context.Context, request *agents.BodyHandleAnOutboundCallViaTwilioV1ConvaiTwilioOutboundCallPost, opts ...option.RequestOption) (*types.TwilioOutboundCallResponse, error) {
	response, err := c.WithRawResponse.OutboundCall(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Register a Twilio call and return TwiML to connect the call
func (c *Client) RegisterCall(ctx context.Context, request *agents.BodyRegisterATwilioCallAndReturnTwiMlV1ConvaiTwilioRegisterCallPost, opts ...option.RequestOption) (string, error) {
	response, err := c.WithRawResponse.RegisterCall(ctx, request, opts...)
	if err != nil {
		return "", err
	}
	return response.Body, nil
}
