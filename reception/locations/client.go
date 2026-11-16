package locations

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

// List all business locations.
func (c *Client) List(ctx context.Context, request *reception.ListLocationsRequest, opts ...option.RequestOption) ([]*types.ReceptionLocationResponse, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Add a new business location.
func (c *Client) Create(ctx context.Context, request *reception.ReceptionLocationRequest, opts ...option.RequestOption) (*types.ReceptionLocationResponse, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get a single location by ID.
func (c *Client) Get(ctx context.Context, locationId string, request *reception.GetLocationsRequest, opts ...option.RequestOption) (*types.ReceptionLocationResponse, error) {
	response, err := c.WithRawResponse.Get(ctx, locationId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Update an existing location. Only the fields sent are changed.
func (c *Client) Update(ctx context.Context, locationId string, request *reception.ReceptionLocationRequestPatch, opts ...option.RequestOption) (*types.ReceptionLocationResponse, error) {
	response, err := c.WithRawResponse.Update(ctx, locationId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete a location.
func (c *Client) Delete(ctx context.Context, locationId string, request *reception.DeleteLocationsRequest, opts ...option.RequestOption) (map[string]string, error) {
	response, err := c.WithRawResponse.Delete(ctx, locationId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
