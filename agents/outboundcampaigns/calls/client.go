package calls

import (
	context "context"
	outboundcampaigns "github.com/elevenlabs/elevenlabs-go/agents/outboundcampaigns"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
	"io"
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

// List the calls of an outbound campaign, newest first: first contacts, retries, callbacks and redirects, whether queued, under way or done.
func (c *Client) List(ctx context.Context, campaignId string, request *outboundcampaigns.ListCallsRequest, opts ...option.RequestOption) (*types.CampaignCallsResponse, error) {
	response, err := c.WithRawResponse.List(ctx, campaignId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Download every call of a campaign, with its lead and conversation results, as CSV.
func (c *Client) Export(ctx context.Context, campaignId string, request *outboundcampaigns.ExportCallsRequest, opts ...option.RequestOption) (io.Reader, error) {
	response, err := c.WithRawResponse.Export(ctx, campaignId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
