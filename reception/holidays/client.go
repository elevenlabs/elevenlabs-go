package holidays

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

// Get all holidays/time-off entries for the current workspace with optional filters
func (c *Client) List(ctx context.Context, request *reception.ListHolidaysRequest, opts ...option.RequestOption) ([]*types.HolidayResponseModel, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Create a new holiday
func (c *Client) Create(ctx context.Context, request *reception.ReceptionCreateHolidayRequest, opts ...option.RequestOption) (*types.HolidayResponseModel, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get a specific holiday by ID
func (c *Client) Get(ctx context.Context, holidayId string, request *reception.GetHolidaysRequest, opts ...option.RequestOption) (*types.HolidayResponseModel, error) {
	response, err := c.WithRawResponse.Get(ctx, holidayId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Update an existing holiday. Only the fields sent are changed.
func (c *Client) Update(ctx context.Context, holidayId string, request *reception.ReceptionUpdateHolidayRequestPatch, opts ...option.RequestOption) (*types.HolidayResponseModel, error) {
	response, err := c.WithRawResponse.Update(ctx, holidayId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete a holiday
func (c *Client) Delete(ctx context.Context, holidayId string, request *reception.DeleteHolidaysRequest, opts ...option.RequestOption) (map[string]string, error) {
	response, err := c.WithRawResponse.Delete(ctx, holidayId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
