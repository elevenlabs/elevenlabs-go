package stock

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	products "github.com/elevenlabs/elevenlabs-go/reception/products"
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

// Set the on-hand quantity of a counted product at one location. Rejected with a 409 if it changed since `expected_on_hand` was read.
func (c *Client) Set(ctx context.Context, productId string, request *products.ReceptionUpdateProductStockRequest, opts ...option.RequestOption) (*types.ReceptionProduct, error) {
	response, err := c.WithRawResponse.Set(ctx, productId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Add to or subtract from the on-hand quantity of a counted product at one location, without overwriting orders taken at the same time.
func (c *Client) Adjust(ctx context.Context, productId string, request *products.ReceptionAdjustProductStockRequest, opts ...option.RequestOption) (*types.ReceptionProduct, error) {
	response, err := c.WithRawResponse.Adjust(ctx, productId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
