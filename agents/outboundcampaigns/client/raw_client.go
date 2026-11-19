package client

import (
	context "context"
	agents "github.com/elevenlabs/elevenlabs-go/agents"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
	http "net/http"
)

type RawClient struct {
	baseURL string
	caller  *internal.Caller
	options *core.RequestOptions
}

func NewRawClient(options *core.RequestOptions) *RawClient {
	return &RawClient{
		options: options,
		baseURL: options.BaseURL,
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
func (r *RawClient) List(ctx context.Context, request *agents.ListOutboundCampaignsRequest, opts ...option.RequestOption) (*core.Response[*types.CampaignPageResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.elevenlabs.io",
	)
	endpointURL := baseURL + "/v1/convai/outbound-campaigns/campaigns"
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	if request == nil {
		request = &agents.ListOutboundCampaignsRequest{}
	}
	var response types.CampaignPageResponse
	raw, err := r.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			DisableRetries:  options.DisableRetries,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
		},
	)
	if err != nil {
		return nil, err
	}
	return &core.Response[*types.CampaignPageResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}

// Pause every running outbound campaign in the workspace. Paused campaigns place no new calls until resumed.
func (r *RawClient) PauseAll(ctx context.Context, opts ...option.RequestOption) (*core.Response[*types.PauseAllCampaignsResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.elevenlabs.io",
	)
	endpointURL := baseURL + "/v1/convai/outbound-campaigns/campaigns/pause-all"
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	var response types.PauseAllCampaignsResponse
	raw, err := r.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			DisableRetries:  options.DisableRetries,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
		},
	)
	if err != nil {
		return nil, err
	}
	return &core.Response[*types.PauseAllCampaignsResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}

// Stop every outbound campaign in the workspace that hasn't finished. Returns `202` while calls already placed are still finishing.
func (r *RawClient) StopAll(ctx context.Context, opts ...option.RequestOption) (*core.Response[*types.StopAllCampaignsResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.elevenlabs.io",
	)
	endpointURL := baseURL + "/v1/convai/outbound-campaigns/campaigns/stop-all"
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	var response types.StopAllCampaignsResponse
	raw, err := r.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			DisableRetries:  options.DisableRetries,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
		},
	)
	if err != nil {
		return nil, err
	}
	return &core.Response[*types.StopAllCampaignsResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}

// Create an outbound campaign that calls or messages each recipient with the agent, within the policy's calling windows, retrying and following up as it allows. Returns `201` for a new campaign. A repeated request, matched by `idempotency_key` or else by its body, returns the existing campaign with `created` false.
func (r *RawClient) Create(ctx context.Context, request *agents.SubmitCampaignRequest, opts ...option.RequestOption) (*core.Response[*types.SubmitCampaignResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.elevenlabs.io",
	)
	endpointURL := baseURL + "/v1/convai/outbound-campaigns/campaigns/submit"
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	headers.Add("Content-Type", "application/json")
	if request == nil {
		request = &agents.SubmitCampaignRequest{}
	}
	var response types.SubmitCampaignResponse
	raw, err := r.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			DisableRetries:  options.DisableRetries,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
		},
	)
	if err != nil {
		return nil, err
	}
	return &core.Response[*types.SubmitCampaignResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}

// Get an outbound campaign and its lead counts.
func (r *RawClient) Get(ctx context.Context, campaignId string, opts ...option.RequestOption) (*core.Response[*types.CampaignResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.elevenlabs.io",
	)
	endpointURL := internal.EncodeURL(baseURL+"/v1/convai/outbound-campaigns/campaigns/%v", campaignId)
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	var response types.CampaignResponse
	raw, err := r.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			DisableRetries:  options.DisableRetries,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
		},
	)
	if err != nil {
		return nil, err
	}
	return &core.Response[*types.CampaignResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}

// Get where an outbound campaign's leads and calls stand, and when its next call is due.
func (r *RawClient) GetProgress(ctx context.Context, campaignId string, opts ...option.RequestOption) (*core.Response[*types.CampaignProgressResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.elevenlabs.io",
	)
	endpointURL := internal.EncodeURL(baseURL+"/v1/convai/outbound-campaigns/campaigns/%v/progress", campaignId)
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	var response types.CampaignProgressResponse
	raw, err := r.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			DisableRetries:  options.DisableRetries,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
		},
	)
	if err != nil {
		return nil, err
	}
	return &core.Response[*types.CampaignProgressResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}

// Count the callbacks and redirects contacts asked for during an outbound campaign's calls.
func (r *RawClient) GetFollowUpRequests(ctx context.Context, campaignId string, opts ...option.RequestOption) (*core.Response[*types.CampaignFollowUpRequestsResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.elevenlabs.io",
	)
	endpointURL := internal.EncodeURL(baseURL+"/v1/convai/outbound-campaigns/campaigns/%v/follow-up-requests", campaignId)
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	var response types.CampaignFollowUpRequestsResponse
	raw, err := r.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			DisableRetries:  options.DisableRetries,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
		},
	)
	if err != nil {
		return nil, err
	}
	return &core.Response[*types.CampaignFollowUpRequestsResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}

// Pause an outbound campaign. It places no new calls until resumed; calls already placed carry on.
func (r *RawClient) Pause(ctx context.Context, campaignId string, opts ...option.RequestOption) (*core.Response[*types.CampaignResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.elevenlabs.io",
	)
	endpointURL := internal.EncodeURL(baseURL+"/v1/convai/outbound-campaigns/campaigns/%v/pause", campaignId)
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	var response types.CampaignResponse
	raw, err := r.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			DisableRetries:  options.DisableRetries,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
		},
	)
	if err != nil {
		return nil, err
	}
	return &core.Response[*types.CampaignResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}

// Resume a paused outbound campaign, including calls an attention item held back.
func (r *RawClient) Resume(ctx context.Context, campaignId string, opts ...option.RequestOption) (*core.Response[*types.CampaignResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.elevenlabs.io",
	)
	endpointURL := internal.EncodeURL(baseURL+"/v1/convai/outbound-campaigns/campaigns/%v/resume", campaignId)
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	var response types.CampaignResponse
	raw, err := r.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			DisableRetries:  options.DisableRetries,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
		},
	)
	if err != nil {
		return nil, err
	}
	return &core.Response[*types.CampaignResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}

// Stop an outbound campaign for good: queued calls are cancelled and no lead is called again. Returns `202` with the campaign `stopping` while calls already placed are still finishing.
func (r *RawClient) Stop(ctx context.Context, campaignId string, opts ...option.RequestOption) (*core.Response[*types.CampaignResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.elevenlabs.io",
	)
	endpointURL := internal.EncodeURL(baseURL+"/v1/convai/outbound-campaigns/campaigns/%v/stop", campaignId)
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	var response types.CampaignResponse
	raw, err := r.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			DisableRetries:  options.DisableRetries,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
		},
	)
	if err != nil {
		return nil, err
	}
	return &core.Response[*types.CampaignResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}
