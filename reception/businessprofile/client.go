package businessprofile

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

// Get the business profile the receptionists answer from, with the current time in the business's timezone.
func (c *Client) Get(ctx context.Context, request *reception.GetBusinessProfileRequest, opts ...option.RequestOption) (*types.ReceptionAgentInfoResponse, error) {
	response, err := c.WithRawResponse.Get(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Update the business profile. Every receptionist is updated to match.
func (c *Client) Update(ctx context.Context, request *reception.ReceptionUpdateAgentInfoRequest, opts ...option.RequestOption) (*types.ReceptionAgentInfoResponse, error) {
	response, err := c.WithRawResponse.Update(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
