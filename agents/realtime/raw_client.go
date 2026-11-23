package realtime

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

// List the realtime conversation groups the current user can view.
func (r *RawClient) ListConversationGroups(ctx context.Context, request *agents.ListConversationGroupsRealtimeRequest, opts ...option.RequestOption) (*core.Response[*types.RealtimeConversationGroupsResponseModel], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		internal.ResolveEnvironmentBaseURL(options.Environment, "Base"),
		r.baseURL,
		internal.ResolveEnvironmentBaseURL(r.options.Environment, "Base"),
		"https://api.elevenlabs.io",
	)
	endpointURL := baseURL + "/v1/convai/realtime/conversation-groups"
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
		request = &agents.ListConversationGroupsRealtimeRequest{}
	}
	var response types.RealtimeConversationGroupsResponseModel
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
	return &core.Response[*types.RealtimeConversationGroupsResponseModel]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}

// Create a realtime conversation group owned by the current user.
func (r *RawClient) CreateConversationGroup(ctx context.Context, request *agents.CreateRealtimeConversationGroupRequestModel, opts ...option.RequestOption) (*core.Response[*types.RealtimeConversationGroupResponseModel], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		internal.ResolveEnvironmentBaseURL(options.Environment, "Base"),
		r.baseURL,
		internal.ResolveEnvironmentBaseURL(r.options.Environment, "Base"),
		"https://api.elevenlabs.io",
	)
	endpointURL := baseURL + "/v1/convai/realtime/conversation-groups"
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	headers.Add("Content-Type", "application/json")
	if request == nil {
		request = &agents.CreateRealtimeConversationGroupRequestModel{}
	}
	var response types.RealtimeConversationGroupResponseModel
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
	return &core.Response[*types.RealtimeConversationGroupResponseModel]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}

// Get a realtime conversation group by id.
func (r *RawClient) GetConversationGroup(ctx context.Context, conversationGroupId string, opts ...option.RequestOption) (*core.Response[*types.RealtimeConversationGroupResponseModel], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		internal.ResolveEnvironmentBaseURL(options.Environment, "Base"),
		r.baseURL,
		internal.ResolveEnvironmentBaseURL(r.options.Environment, "Base"),
		"https://api.elevenlabs.io",
	)
	endpointURL := internal.EncodeURL(baseURL+"/v1/convai/realtime/conversation-groups/%v", conversationGroupId)
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	var response types.RealtimeConversationGroupResponseModel
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
	return &core.Response[*types.RealtimeConversationGroupResponseModel]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}

// Rename or re-describe a realtime conversation group.
func (r *RawClient) UpdateConversationGroup(ctx context.Context, conversationGroupId string, request *agents.UpdateRealtimeConversationGroupRequestModel, opts ...option.RequestOption) (*core.Response[*types.RealtimeConversationGroupResponseModel], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		internal.ResolveEnvironmentBaseURL(options.Environment, "Base"),
		r.baseURL,
		internal.ResolveEnvironmentBaseURL(r.options.Environment, "Base"),
		"https://api.elevenlabs.io",
	)
	endpointURL := internal.EncodeURL(baseURL+"/v1/convai/realtime/conversation-groups/%v", conversationGroupId)
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	headers.Add("Content-Type", "application/json")
	var response types.RealtimeConversationGroupResponseModel
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
	return &core.Response[*types.RealtimeConversationGroupResponseModel]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}

// Delete a realtime conversation group. Every conversation in it becomes ungrouped, so access granted through the group is revoked.
func (r *RawClient) DeleteConversationGroup(ctx context.Context, conversationGroupId string, opts ...option.RequestOption) (*core.Response[any], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		internal.ResolveEnvironmentBaseURL(options.Environment, "Base"),
		r.baseURL,
		internal.ResolveEnvironmentBaseURL(r.options.Environment, "Base"),
		"https://api.elevenlabs.io",
	)
	endpointURL := internal.EncodeURL(baseURL+"/v1/convai/realtime/conversation-groups/%v", conversationGroupId)
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	var response any
	raw, err := r.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodDelete,
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
	return &core.Response[any]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       response,
	}, nil
}

// Place a realtime conversation in a conversation group, moving it out of its current group if it has one.
func (r *RawClient) SetConversationGroup(ctx context.Context, conversationId string, request *agents.SetConversationGroupRequestModel, opts ...option.RequestOption) (*core.Response[any], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		internal.ResolveEnvironmentBaseURL(options.Environment, "Base"),
		r.baseURL,
		internal.ResolveEnvironmentBaseURL(r.options.Environment, "Base"),
		"https://api.elevenlabs.io",
	)
	endpointURL := internal.EncodeURL(baseURL+"/v1/convai/realtime/conversations/%v/conversation-group", conversationId)
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	headers.Add("Content-Type", "application/json")
	var response any
	raw, err := r.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPut,
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
	return &core.Response[any]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       response,
	}, nil
}

// Remove a realtime conversation from its conversation group.
func (r *RawClient) ClearConversationGroup(ctx context.Context, conversationId string, opts ...option.RequestOption) (*core.Response[any], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		internal.ResolveEnvironmentBaseURL(options.Environment, "Base"),
		r.baseURL,
		internal.ResolveEnvironmentBaseURL(r.options.Environment, "Base"),
		"https://api.elevenlabs.io",
	)
	endpointURL := internal.EncodeURL(baseURL+"/v1/convai/realtime/conversations/%v/conversation-group", conversationId)
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	var response any
	raw, err := r.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodDelete,
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
	return &core.Response[any]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       response,
	}, nil
}

// Create an ephemeral client secret for authenticating browser-based Realtime API sessions.
func (r *RawClient) CreateClientSecret(ctx context.Context, request *agents.CreateClientSecretRequest, opts ...option.RequestOption) (*core.Response[*types.ClientSecretResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		internal.ResolveEnvironmentBaseURL(options.Environment, "Base"),
		r.baseURL,
		internal.ResolveEnvironmentBaseURL(r.options.Environment, "Base"),
		"https://api.elevenlabs.io",
	)
	endpointURL := baseURL + "/v1/convai/realtime/client_secrets"
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	headers.Add("Content-Type", "application/json")
	if request == nil {
		request = &agents.CreateClientSecretRequest{}
	}
	var response types.ClientSecretResponse
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
	return &core.Response[*types.ClientSecretResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}
