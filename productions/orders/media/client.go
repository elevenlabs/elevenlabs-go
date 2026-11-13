package media

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	orders "github.com/elevenlabs/elevenlabs-go/productions/orders"
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

// Registers a media file with an order, either by uploading it directly or by providing a URL to fetch it from. Exactly one of `media` or `media_url` must be provided. The registered media can then be referenced when adding order items.
func (c *Client) Register(ctx context.Context, orderId OrderId, request *orders.BodyRegisterMediaV1ProductionsOrdersOrderIdMediaPost, opts ...option.RequestOption) (*types.RegisterMediaResponse, error) {
	response, err := c.WithRawResponse.Register(ctx, orderId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Retrieves metadata and a time-limited download URL for a previously uploaded media file.
func (c *Client) Get(ctx context.Context, orderId OrderId, mediaId MediaId, opts ...option.RequestOption) (*types.OrderMediaResponse, error) {
	response, err := c.WithRawResponse.Get(ctx, orderId, mediaId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
