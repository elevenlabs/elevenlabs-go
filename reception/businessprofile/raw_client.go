package businessprofile

import (
	context "context"
	core "github.com/elevenlabs/elevenlabs-go/core"
	internal "github.com/elevenlabs/elevenlabs-go/internal"
	option "github.com/elevenlabs/elevenlabs-go/option"
	reception "github.com/elevenlabs/elevenlabs-go/reception"
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

// Get the business profile the receptionists answer from, with the current time in the business's timezone.
func (r *RawClient) Get(ctx context.Context, request *reception.GetBusinessProfileRequest, opts ...option.RequestOption) (*core.Response[*types.ReceptionAgentInfoResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.elevenlabs.io",
	)
	endpointURL := baseURL + "/v1/reception/agent-info"
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	if request == nil {
		request = &reception.GetBusinessProfileRequest{}
	}
	internal.SetHeader(headers, "xi-workspace-id", request.XiWorkspaceId)
	var response types.ReceptionAgentInfoResponse
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
	return &core.Response[*types.ReceptionAgentInfoResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}

// Update the business profile. Every receptionist is updated to match.
func (r *RawClient) Update(ctx context.Context, request *reception.ReceptionUpdateAgentInfoRequest, opts ...option.RequestOption) (*core.Response[*types.ReceptionAgentInfoResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.elevenlabs.io",
	)
	endpointURL := baseURL + "/v1/reception/agent-info"
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	headers.Add("Content-Type", "application/json")
	if request == nil {
		request = &reception.ReceptionUpdateAgentInfoRequest{}
	}
	internal.SetHeader(headers, "xi-workspace-id", request.XiWorkspaceId)
	var response types.ReceptionAgentInfoResponse
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
	return &core.Response[*types.ReceptionAgentInfoResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}
