package assets

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

// List assets in the workspace, most recently created first.
func (c *Client) List(ctx context.Context, request *types.ListAssetsRequest, opts ...option.RequestOption) (*types.AssetListResponse, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Upload a new asset.
func (c *Client) Create(ctx context.Context, bodyUploadAssetV1AssetsPost *types.BodyUploadAssetV1AssetsPost, opts ...option.RequestOption) (*types.AssetResponse, error) {
	response, err := c.WithRawResponse.Create(ctx, bodyUploadAssetV1AssetsPost, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Retrieve a single asset by ID.
func (c *Client) Get(ctx context.Context, assetId string, opts ...option.RequestOption) (*types.AssetResponse, error) {
	response, err := c.WithRawResponse.Get(ctx, assetId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete an asset by ID.
func (c *Client) Delete(ctx context.Context, assetId string, opts ...option.RequestOption) error {
	_, err := c.WithRawResponse.Delete(ctx, assetId, opts...)
	return err
}
