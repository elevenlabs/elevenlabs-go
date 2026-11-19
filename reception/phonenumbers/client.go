package phonenumbers

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

// List the workspace's phone numbers and the receptionist each one rings. A number with no `customer_facing_agent_id` is held by the workspace but not attached to a receptionist yet.
func (c *Client) List(ctx context.Context, request *reception.ListPhoneNumbersRequest, opts ...option.RequestOption) ([]*types.ReceptionPhoneAssignment, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Give a receptionist a phone number. A number the workspace already holds but has not attached is used first. Otherwise, in the US and Canada, a number with one of the requested `area_codes` is taken from the pool or bought, falling back to any number in the business's country. In countries that require paperwork, the number is issued right away but takes public calls only once the workspace's own registration is approved (see `accepts_public_calls`). Counts toward the plan's phone number limit, and free plans can take only a few new numbers a day. Returns 404 when no number is available and 403 `telephony_managed_by_reseller` when the reseller supplies the workspace's numbers from its own Twilio account or SIP trunk and none is free.
func (c *Client) Create(ctx context.Context, request *reception.CreatePhoneAssignmentRequest, opts ...option.RequestOption) (*types.ReceptionPhoneAssignment, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Move a phone number to a different receptionist in the same workspace, which answers its calls from then on. Returns 409 `phone_number_not_attached` for a number no receptionist answers yet (attach it with `POST /v1/reception/phone-numbers/assignments`).
func (c *Client) Update(ctx context.Context, phoneAssignmentId string, request *reception.ReassignPhoneRequest, opts ...option.RequestOption) (*types.ReceptionPhoneAssignment, error) {
	response, err := c.WithRawResponse.Update(ctx, phoneAssignmentId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Remove a phone number from the workspace. A pool number goes back to the pool it came from. A number the business brought from its own Twilio account, or one issued in a country that requires paperwork, is released at Twilio and cannot be recovered. Returns 403 `telephony_managed_by_reseller` when the reseller supplies the workspace's numbers from its own Twilio account or SIP trunk.
func (c *Client) Delete(ctx context.Context, phoneAssignmentId string, request *reception.DeletePhoneNumbersRequest, opts ...option.RequestOption) error {
	_, err := c.WithRawResponse.Delete(ctx, phoneAssignmentId, request, opts...)
	return err
}
