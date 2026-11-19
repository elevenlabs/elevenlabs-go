package attention

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

// List problems in the workspace's outbound campaigns that need you to act, such as an agent or phone number that can no longer be used, newest first. A campaign with an open item holds back the calls it affects.
func (c *Client) List(ctx context.Context, request *outboundcampaigns.ListAttentionRequest, opts ...option.RequestOption) (*types.AttentionPageResponse, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Place the calls a retryable attention item holds back, once its cause is fixed. Resumes the campaign.
func (c *Client) Retry(ctx context.Context, attentionId string, opts ...option.RequestOption) (*types.AttentionView, error) {
	response, err := c.WithRawResponse.Retry(ctx, attentionId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Close an attention item that can't be retried, recording how it was handled.
func (c *Client) Resolve(ctx context.Context, attentionId string, request *outboundcampaigns.AttentionResolutionRequest, opts ...option.RequestOption) (*types.AttentionView, error) {
	response, err := c.WithRawResponse.Resolve(ctx, attentionId, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
