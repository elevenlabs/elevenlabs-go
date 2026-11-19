package client

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	reception "github.com/elevenlabs/elevenlabs-go/reception"
	stock "github.com/elevenlabs/elevenlabs-go/reception/products/stock"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

type Client struct {
	WithRawResponse *RawClient

	Stock *stock.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse: NewRawClient(options),
		Stock:           stock.NewClient(options),
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

// List the workspace's products, optionally filtered by group or active state.
func (c *Client) List(ctx context.Context, request *reception.ListProductsRequest, opts ...option.RequestOption) ([]*types.ReceptionProduct, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Create a product the receptionist can quote and take orders for. A product created with `track_quantity` starts with nothing on hand; add stock with the stock endpoints.
func (c *Client) Create(ctx context.Context, request *reception.ReceptionCreateProductRequest, opts ...option.RequestOption) (*types.ReceptionProduct, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get a product by ID.
func (c *Client) Get(ctx context.Context, productId string, request *reception.GetProductsRequest, opts ...option.RequestOption) (*types.ReceptionProduct, error) {
	response, err := c.WithRawResponse.Get(ctx, productId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Update a product's catalog fields. Only the fields sent are changed. On-hand quantities are not part of this call; use the stock endpoints.
func (c *Client) Update(ctx context.Context, productId string, request *reception.ReceptionUpdateProductRequestPatch, opts ...option.RequestOption) (*types.ReceptionProduct, error) {
	response, err := c.WithRawResponse.Update(ctx, productId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete a product.
func (c *Client) Delete(ctx context.Context, productId string, request *reception.DeleteProductsRequest, opts ...option.RequestOption) (map[string]string, error) {
	response, err := c.WithRawResponse.Delete(ctx, productId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
