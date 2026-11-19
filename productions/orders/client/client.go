package client

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	productions "github.com/elevenlabs/elevenlabs-go/productions"
	deliverables "github.com/elevenlabs/elevenlabs-go/productions/orders/deliverables"
	items "github.com/elevenlabs/elevenlabs-go/productions/orders/items"
	languages "github.com/elevenlabs/elevenlabs-go/productions/orders/languages"
	media "github.com/elevenlabs/elevenlabs-go/productions/orders/media"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

type Client struct {
	WithRawResponse *RawClient

	Media        *media.Client
	Items        *items.Client
	Deliverables *deliverables.Client
	Languages    *languages.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse: NewRawClient(options),
		Media:           media.NewClient(options),
		Items:           items.NewClient(options),
		Deliverables:    deliverables.NewClient(options),
		Languages:       languages.NewClient(options),
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

// Lists Productions orders in the workspace. Supports filtering by status and date range, with pagination.
func (c *Client) List(ctx context.Context, request *productions.ListOrdersRequest, opts ...option.RequestOption) (*types.ListOrdersResponse, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Creates a new Productions order in the workspace. The order starts in the open state and can be configured with items before submission.
func (c *Client) Create(ctx context.Context, request *productions.CreateOrderRequest, opts ...option.RequestOption) (*types.CreateOrderResponse, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Retrieves full details for a Productions order.
//
// Quote and pricing information may not be available immediately; if you wish to see the quote before submission, you may need to poll the order details until it is ready.
func (c *Client) Get(ctx context.Context, orderId types.OrderId, opts ...option.RequestOption) (*types.OrderResponse, error) {
	response, err := c.WithRawResponse.Get(ctx, orderId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Updates an open order.
func (c *Client) Update(ctx context.Context, orderId types.OrderId, request *productions.BodyUpdateOrderV1ProductionsOrdersOrderIdPatch, opts ...option.RequestOption) (*types.UpdateOrderResponse, error) {
	response, err := c.WithRawResponse.Update(ctx, orderId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Submits an open order for processing. The order must have at least one item. Once submitted, items can no longer be modified.
//
// Upon submission, the workspace will be charged for the order. The quote is based on information extracted from the uploaded media, such as its duration. The quote may not be available immediately; if you wish to see the quote before submission, you may need to poll the order details until the quote is ready.
func (c *Client) Submit(ctx context.Context, orderId types.OrderId, opts ...option.RequestOption) (*types.SubmitOrderResponse, error) {
	response, err := c.WithRawResponse.Submit(ctx, orderId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
