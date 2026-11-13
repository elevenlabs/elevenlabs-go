package client

import (
	context "context"
	agents "github.com/elevenlabs/elevenlabs-go/agents"
	attention "github.com/elevenlabs/elevenlabs-go/agents/outboundcampaigns/attention"
	calls "github.com/elevenlabs/elevenlabs-go/agents/outboundcampaigns/calls"
	config "github.com/elevenlabs/elevenlabs-go/agents/outboundcampaigns/config"
	leads "github.com/elevenlabs/elevenlabs-go/agents/outboundcampaigns/leads"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

type Client struct {
	WithRawResponse *RawClient

	Config    *config.Client
	Leads     *leads.Client
	Calls     *calls.Client
	Attention *attention.Client

	options *core.RequestOptions
	baseURL string
	caller  *internal.Caller
}

func NewClient(options *core.RequestOptions) *Client {
	return &Client{
		WithRawResponse: NewRawClient(options),
		Config:          config.NewClient(options),
		Leads:           leads.NewClient(options),
		Calls:           calls.NewClient(options),
		Attention:       attention.NewClient(options),
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

// List the workspace's outbound campaigns, newest first.
func (c *Client) List(ctx context.Context, request *agents.ListOutboundCampaignsRequest, opts ...option.RequestOption) (*types.CampaignPageResponse, error) {
	response, err := c.WithRawResponse.List(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Pause every running outbound campaign in the workspace. Paused campaigns place no new calls until resumed.
func (c *Client) PauseAll(ctx context.Context, opts ...option.RequestOption) (*types.PauseAllCampaignsResponse, error) {
	response, err := c.WithRawResponse.PauseAll(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Stop every outbound campaign in the workspace that hasn't finished. Returns `202` while calls already placed are still finishing.
func (c *Client) StopAll(ctx context.Context, opts ...option.RequestOption) (*types.StopAllCampaignsResponse, error) {
	response, err := c.WithRawResponse.StopAll(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Create an outbound campaign that calls or messages each recipient with the agent, within the policy's calling windows, retrying and following up as it allows. Returns `201` for a new campaign. A repeated request, matched by `idempotency_key` or else by its body, returns the existing campaign with `created` false.
func (c *Client) Create(ctx context.Context, request *agents.SubmitCampaignRequest, opts ...option.RequestOption) (*types.SubmitCampaignResponse, error) {
	response, err := c.WithRawResponse.Create(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get an outbound campaign and its lead counts.
func (c *Client) Get(ctx context.Context, campaignId string, opts ...option.RequestOption) (*types.CampaignResponse, error) {
	response, err := c.WithRawResponse.Get(ctx, campaignId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Get where an outbound campaign's leads and calls stand, and when its next call is due.
func (c *Client) GetProgress(ctx context.Context, campaignId string, opts ...option.RequestOption) (*types.CampaignProgressResponse, error) {
	response, err := c.WithRawResponse.GetProgress(ctx, campaignId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Count the callbacks and redirects contacts asked for during an outbound campaign's calls.
func (c *Client) GetFollowUpRequests(ctx context.Context, campaignId string, opts ...option.RequestOption) (*types.CampaignFollowUpRequestsResponse, error) {
	response, err := c.WithRawResponse.GetFollowUpRequests(ctx, campaignId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Pause an outbound campaign. It places no new calls until resumed; calls already placed carry on.
func (c *Client) Pause(ctx context.Context, campaignId string, opts ...option.RequestOption) (*types.CampaignResponse, error) {
	response, err := c.WithRawResponse.Pause(ctx, campaignId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Resume a paused outbound campaign, including calls an attention item held back.
func (c *Client) Resume(ctx context.Context, campaignId string, opts ...option.RequestOption) (*types.CampaignResponse, error) {
	response, err := c.WithRawResponse.Resume(ctx, campaignId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

// Stop an outbound campaign for good: queued calls are cancelled and no lead is called again. Returns `202` with the campaign `stopping` while calls already placed are still finishing.
func (c *Client) Stop(ctx context.Context, campaignId string, opts ...option.RequestOption) (*types.CampaignResponse, error) {
	response, err := c.WithRawResponse.Stop(ctx, campaignId, opts...)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
