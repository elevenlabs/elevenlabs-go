package mergeproposals

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

// List the proposals for an agent, newest first.
func (r *RawClient) List(ctx context.Context, agentId string, request *agents.ListMergeProposalsRequest, opts ...option.RequestOption) (*core.Response[*types.PaginatedResultAgentMergeProposalResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.elevenlabs.io",
	)
	endpointURL := internal.EncodeURL(baseURL+"/v1/convai/agents/%v/merge-proposals", agentId)
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
		request = &agents.ListMergeProposalsRequest{}
	}
	var response types.PaginatedResultAgentMergeProposalResponse
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
	return &core.Response[*types.PaginatedResultAgentMergeProposalResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}

// Record a request to merge a source branch into a target branch. Anyone with edit access can open one; merging it later is gated on write access to the target branch, so this is how a change reaches a protected branch the author cannot merge into themselves.
func (r *RawClient) Create(ctx context.Context, agentId string, request *agents.BodyCreateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsPost, opts ...option.RequestOption) (*core.Response[*types.CreateAgentMergeProposalResponseModel], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.elevenlabs.io",
	)
	endpointURL := internal.EncodeURL(baseURL+"/v1/convai/agents/%v/merge-proposals", agentId)
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	headers.Add("Content-Type", "application/json")
	var response types.CreateAgentMergeProposalResponseModel
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
	return &core.Response[*types.CreateAgentMergeProposalResponseModel]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}

// Get a single merge_proposal.
func (r *RawClient) Get(ctx context.Context, agentId string, mergeProposalId string, opts ...option.RequestOption) (*core.Response[*types.AgentMergeProposalResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.elevenlabs.io",
	)
	endpointURL := internal.EncodeURL(baseURL+"/v1/convai/agents/%v/merge-proposals/%v", agentId, mergeProposalId)
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	var response types.AgentMergeProposalResponse
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
	return &core.Response[*types.AgentMergeProposalResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}

// Edit an open merge_proposal's title or description, or close it. The author closing it is recorded as withdrawn; anyone else as rejected.
func (r *RawClient) Update(ctx context.Context, agentId string, mergeProposalId string, request *agents.BodyUpdateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdPatch, opts ...option.RequestOption) (*core.Response[*types.AgentMergeProposalResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.elevenlabs.io",
	)
	endpointURL := internal.EncodeURL(baseURL+"/v1/convai/agents/%v/merge-proposals/%v", agentId, mergeProposalId)
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	headers.Add("Content-Type", "application/json")
	var response types.AgentMergeProposalResponse
	raw, err := r.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPatch,
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
	return &core.Response[*types.AgentMergeProposalResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}

// Approve a merge_proposal or request changes on it. A user's latest review replaces their previous one. Non-admins need an approval from another user before the merge is allowed.
func (r *RawClient) SubmitReview(ctx context.Context, agentId string, mergeProposalId string, request *agents.ReviewAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdReviewsPost, opts ...option.RequestOption) (*core.Response[*types.AgentMergeProposalResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.elevenlabs.io",
	)
	endpointURL := internal.EncodeURL(baseURL+"/v1/convai/agents/%v/merge-proposals/%v/reviews", agentId, mergeProposalId)
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	headers.Add("Content-Type", "application/json")
	var response types.AgentMergeProposalResponse
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
	return &core.Response[*types.AgentMergeProposalResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}

// Leave a comment on a merge_proposal without recording a review verdict. Unlike reviews, comments accumulate and can still be added once the merge_proposal is merged or closed.
func (r *RawClient) AddComment(ctx context.Context, agentId string, mergeProposalId string, request *agents.OnAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdCommentsPost, opts ...option.RequestOption) (*core.Response[*types.AgentMergeProposalResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.elevenlabs.io",
	)
	endpointURL := internal.EncodeURL(baseURL+"/v1/convai/agents/%v/merge-proposals/%v/comments", agentId, mergeProposalId)
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	headers.Add("Content-Type", "application/json")
	var response types.AgentMergeProposalResponse
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
	return &core.Response[*types.AgentMergeProposalResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}

// Execute the merge. The caller must have write access to the target branch (admins only, for a protected branch), so this is where a reviewer approves and merges a request opened by someone who could not.
func (r *RawClient) Merge(ctx context.Context, agentId string, mergeProposalId string, request *agents.BodyMergeAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdMergePost, opts ...option.RequestOption) (*core.Response[*types.AgentMergeProposalResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.elevenlabs.io",
	)
	endpointURL := internal.EncodeURL(baseURL+"/v1/convai/agents/%v/merge-proposals/%v/merge", agentId, mergeProposalId)
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	headers.Add("Content-Type", "application/json")
	var response types.AgentMergeProposalResponse
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
	return &core.Response[*types.AgentMergeProposalResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}
