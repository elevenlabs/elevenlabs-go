package billing

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

// List the invoices charged to the reseller's card, newest first, one per workspace invoice
func (c *Client) ListInvoices(ctx context.Context, request *reception.ListInvoicesBillingRequest, opts ...option.RequestOption) (*types.ResellerInvoicesResponse, error) {
	response, err := c.WithRawResponse.ListInvoices(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get one invoice charged to the reseller's card, broken down into the workspace invoice's lines
func (c *Client) GetInvoice(ctx context.Context, invoiceId string, opts ...option.RequestOption) (*types.ResellerInvoiceDetail, error) {
	response, err := c.WithRawResponse.GetInvoice(ctx, invoiceId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// List the next workspace renewals the reseller's card will be charged for, soonest first
func (c *Client) ListUpcomingPayments(ctx context.Context, opts ...option.RequestOption) (*types.ResellerUpcomingPaymentsResponse, error) {
	response, err := c.WithRawResponse.ListUpcomingPayments(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
