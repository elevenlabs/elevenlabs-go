package assets

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	reception "github.com/elevenlabs/elevenlabs-go/reception"
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

// List the workspace's bookable assets, such as rooms or equipment.
func (c *Client) List(ctx context.Context, request *reception.ListAssetsRequest, opts ...option.RequestOption) ([]*types.ReceptionAssetResource, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Create a bookable asset, such as a room or a piece of equipment.
func (c *Client) Create(ctx context.Context, request *reception.ReceptionCreateAssetRequest, opts ...option.RequestOption) (*types.ReceptionAssetResource, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get an asset by ID.
func (c *Client) Get(ctx context.Context, assetId string, request *reception.GetAssetsRequest, opts ...option.RequestOption) (*types.ReceptionAssetResource, error) {
	response, err := c.WithRawResponse.Get(ctx, assetId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Update an asset. Only the fields sent are changed.
func (c *Client) Update(ctx context.Context, assetId string, request *reception.ReceptionUpdateAssetRequestPatch, opts ...option.RequestOption) (*types.ReceptionAssetResource, error) {
	response, err := c.WithRawResponse.Update(ctx, assetId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete an asset.
func (c *Client) Delete(ctx context.Context, assetId string, request *reception.DeleteAssetsRequest, opts ...option.RequestOption) (map[string]string, error) {
	response, err := c.WithRawResponse.Delete(ctx, assetId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
