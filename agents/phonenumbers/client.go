package phonenumbers

import (
	context "context"
	agents "github.com/elevenlabs/elevenlabs-go/agents"
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

// Retrieve all Phone Numbers
func (c *Client) List(ctx context.Context, request *agents.ListPhoneNumbersRequest, opts ...option.RequestOption) ([]*agents.ListPhoneNumbersResponseItem, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Import Phone Number from provider configuration (Twilio, Exotel, or SIP trunk)
func (c *Client) Create(ctx context.Context, request *agents.CreatePhoneNumbersRequestBody, opts ...option.RequestOption) (*types.CreatePhoneNumberResponseModel, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Retrieve Phone Number details by ID
func (c *Client) Get(ctx context.Context, phoneNumberId string, opts ...option.RequestOption) (*agents.GetPhoneNumbersResponse, error) {
	response, err := c.WithRawResponse.Get(ctx, phoneNumberId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Update assigned agent of a phone number
func (c *Client) Update(ctx context.Context, phoneNumberId string, request *agents.UpdatePhoneNumberRequest, opts ...option.RequestOption) (*agents.UpdatePhoneNumbersResponse, error) {
	response, err := c.WithRawResponse.Update(ctx, phoneNumberId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Delete Phone Number by ID
func (c *Client) Delete(ctx context.Context, phoneNumberId string, opts ...option.RequestOption) (any, error) {
	response, err := c.WithRawResponse.Delete(ctx, phoneNumberId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Retrieve a page of Phone Numbers
func (c *Client) ListV2(ctx context.Context, request *agents.ListV2PhoneNumbersRequest, opts ...option.RequestOption) (*types.GetPhoneNumbersPageResponseModel, error) {
	response, err := c.WithRawResponse.ListV2(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get SIP messages for a phone number
func (c *Client) GetSipMessages(ctx context.Context, phoneNumberId string, request *agents.GetSipMessagesPhoneNumbersRequest, opts ...option.RequestOption) (*types.GetSipLogMessagesResponse, error) {
	response, err := c.WithRawResponse.GetSipMessages(ctx, phoneNumberId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
