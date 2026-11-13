package leads

import (
	context "context"
	outboundcampaigns "github.com/elevenlabs/elevenlabs-go/agents/outboundcampaigns"
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

// List the leads of an outbound campaign, one per recipient.
func (c *Client) List(ctx context.Context, campaignId string, request *outboundcampaigns.ListLeadsRequest, opts ...option.RequestOption) (*types.CampaignLeadsResponse, error) {
	response, err := c.WithRawResponse.List(ctx, campaignId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get one lead of an outbound campaign.
func (c *Client) Get(ctx context.Context, campaignId string, leadId string, opts ...option.RequestOption) (*types.LeadDetailResponse, error) {
	response, err := c.WithRawResponse.Get(ctx, campaignId, leadId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Remove a lead from an outbound campaign so it is never called again. Returns `202` when a call already placed may still go through.
func (c *Client) Delete(ctx context.Context, campaignId string, leadId string, opts ...option.RequestOption) (*types.SuppressLeadResponse, error) {
	response, err := c.WithRawResponse.Delete(ctx, campaignId, leadId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
