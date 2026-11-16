package items

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

// Adds or updates an order item on an open order. Returns the item ID and the quoted price.
func (c *Client) Upsert(ctx context.Context, orderId OrderId, request *orders.BodyUpsertOrderItemV1ProductionsOrdersOrderIdItemsPost, opts ...option.RequestOption) (*types.UpsertOrderItemResponse, error) {
	response, err := c.WithRawResponse.Upsert(ctx, orderId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Removes an order item from an open order.
func (c *Client) Remove(ctx context.Context, orderId OrderId, itemId ItemId, opts ...option.RequestOption) (*types.RemoveOrderItemResponse, error) {
	response, err := c.WithRawResponse.Remove(ctx, orderId, itemId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
