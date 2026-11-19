package services

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

// Get all services for the current workspace with optional filters
func (c *Client) List(ctx context.Context, request *reception.ListServicesRequest, opts ...option.RequestOption) ([]*types.ReceptionService, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Create a new service
func (c *Client) Create(ctx context.Context, request *reception.ReceptionCreateServiceRequest, opts ...option.RequestOption) (*types.ReceptionService, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get a specific service by ID
func (c *Client) Get(ctx context.Context, serviceId string, request *reception.GetServicesRequest, opts ...option.RequestOption) (*types.ReceptionService, error) {
	response, err := c.WithRawResponse.Get(ctx, serviceId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Update an existing service. Only the fields sent are changed.
func (c *Client) Update(ctx context.Context, serviceId string, request *reception.ReceptionUpdateServiceRequestPatch, opts ...option.RequestOption) (*types.ReceptionServiceUpdateResponse, error) {
	response, err := c.WithRawResponse.Update(ctx, serviceId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete a service
func (c *Client) Delete(ctx context.Context, serviceId string, request *reception.DeleteServicesRequest, opts ...option.RequestOption) (map[string]string, error) {
	response, err := c.WithRawResponse.Delete(ctx, serviceId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Save a service update that changes the price of existing bookings. Optionally reprice future bookings.
func (c *Client) ConfirmUpdate(ctx context.Context, serviceId string, request *reception.ReceptionConfirmServiceUpdateRequest, opts ...option.RequestOption) (*types.ReceptionServiceUpdateResponse, error) {
	response, err := c.WithRawResponse.ConfirmUpdate(ctx, serviceId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
