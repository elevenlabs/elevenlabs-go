package staff

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

// Get all staff for the current workspace
func (c *Client) List(ctx context.Context, request *reception.ListStaffRequest, opts ...option.RequestOption) ([]*types.ReceptionStaffResource, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Create a new staff member
func (c *Client) Create(ctx context.Context, request *reception.ReceptionCreateStaffRequest, opts ...option.RequestOption) (*types.ReceptionStaffResource, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get a specific staff member by ID
func (c *Client) Get(ctx context.Context, staffId string, request *reception.GetStaffRequest, opts ...option.RequestOption) (*types.ReceptionStaffResource, error) {
	response, err := c.WithRawResponse.Get(ctx, staffId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Update an existing staff member. Only the fields sent are changed.
func (c *Client) Update(ctx context.Context, staffId string, request *reception.ReceptionUpdateStaffRequestPatch, opts ...option.RequestOption) (*types.ReceptionStaffResource, error) {
	response, err := c.WithRawResponse.Update(ctx, staffId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete a staff member
func (c *Client) Delete(ctx context.Context, staffId string, request *reception.DeleteStaffRequest, opts ...option.RequestOption) (map[string]string, error) {
	response, err := c.WithRawResponse.Delete(ctx, staffId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
