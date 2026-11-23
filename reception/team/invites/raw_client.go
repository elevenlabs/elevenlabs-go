package invites

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	team "github.com/elevenlabs/elevenlabs-go/reception/team"
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

// List the workspace's pending invites.
func (r *RawClient) List(ctx context.Context, request *team.ListInvitesRequest, opts ...option.RequestOption) (*core.Response[[]*types.ReceptionTeamInviteResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		internal.ResolveEnvironmentBaseURL(options.Environment, "Base"),
		r.baseURL,
		internal.ResolveEnvironmentBaseURL(r.options.Environment, "Base"),
		"https://api.elevenlabs.io",
	)
	endpointURL := baseURL + "/v1/reception/team/invites"
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	if request == nil {
		request = &team.ListInvitesRequest{}
	}
	internal.SetHeader(headers, "xi-workspace-id", request.XiWorkspaceId)
	var response []*types.ReceptionTeamInviteResponse
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
	return &core.Response[[]*types.ReceptionTeamInviteResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       response,
	}, nil
}

// Email an invite to join the workspace. A pending invite takes a seat, so this is refused once members and invites reach the plan's team limit.
func (r *RawClient) Create(ctx context.Context, request *team.CreateTeamInviteRequest, opts ...option.RequestOption) (*core.Response[map[string]string], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		internal.ResolveEnvironmentBaseURL(options.Environment, "Base"),
		r.baseURL,
		internal.ResolveEnvironmentBaseURL(r.options.Environment, "Base"),
		"https://api.elevenlabs.io",
	)
	endpointURL := baseURL + "/v1/reception/team/invites"
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	headers.Add("Content-Type", "application/json")
	if request == nil {
		request = &team.CreateTeamInviteRequest{}
	}
	internal.SetHeader(headers, "xi-workspace-id", request.XiWorkspaceId)
	var response map[string]string
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
	return &core.Response[map[string]string]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       response,
	}, nil
}

// Cancel a pending invite, freeing its seat.
func (r *RawClient) Cancel(ctx context.Context, request *team.BodySmbCancelTeamInvite, opts ...option.RequestOption) (*core.Response[map[string]string], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		internal.ResolveEnvironmentBaseURL(options.Environment, "Base"),
		r.baseURL,
		internal.ResolveEnvironmentBaseURL(r.options.Environment, "Base"),
		"https://api.elevenlabs.io",
	)
	endpointURL := baseURL + "/v1/reception/team/invites"
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	headers.Add("Content-Type", "application/json")
	if request == nil {
		request = &team.BodySmbCancelTeamInvite{}
	}
	internal.SetHeader(headers, "xi-workspace-id", request.XiWorkspaceId)
	var response map[string]string
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
			Request:         request,
			Response:        &response,
		},
	)
	if err != nil {
		return nil, err
	}
	return &core.Response[map[string]string]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       response,
	}, nil
}
