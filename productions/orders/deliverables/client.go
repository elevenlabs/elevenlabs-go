package deliverables

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

// Retrieves the delivered files for a completed order. Returns an empty list if the order is not yet completed.
func (c *Client) List(ctx context.Context, orderId OrderId, opts ...option.RequestOption) (*types.OrderDeliverablesResponse, error) {
	response, err := c.WithRawResponse.List(ctx, orderId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
