package exotel

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

// Handle an outbound call via Exotel Connect API
func (c *Client) OutboundCall(ctx context.Context, request *agents.BodyHandleAnOutboundCallViaExotelV1ConvaiExotelOutboundCallPost, opts ...option.RequestOption) (*types.ExotelOutboundCallResponse, error) {
	response, err := c.WithRawResponse.OutboundCall(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
