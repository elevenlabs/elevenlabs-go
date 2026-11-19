package settings

import (
	context "context"
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

// Gets the default settings for voices. "similarity_boost" corresponds to"Clarity + Similarity Enhancement" in the web app and "stability" corresponds to "Stability" slider in the web app.
func (r *RawClient) GetDefault(ctx context.Context, opts ...option.RequestOption) (*core.Response[*types.VoiceSettings], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.elevenlabs.io",
	)
	endpointURL := baseURL + "/v1/voices/settings/default"
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	var response types.VoiceSettings
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
	return &core.Response[*types.VoiceSettings]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}

// Returns the settings for a specific voice. "similarity_boost" corresponds to"Clarity + Similarity Enhancement" in the web app and "stability" corresponds to "Stability" slider in the web app.
func (r *RawClient) Get(ctx context.Context, voiceId string, opts ...option.RequestOption) (*core.Response[*types.VoiceSettings], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.elevenlabs.io",
	)
	endpointURL := internal.EncodeURL(baseURL+"/v1/voices/%v/settings", voiceId)
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	var response types.VoiceSettings
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
	return &core.Response[*types.VoiceSettings]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}

// Edit your settings for a specific voice. "similarity_boost" corresponds to "Clarity + Similarity Enhancement" in the web app and "stability" corresponds to "Stability" slider in the web app.
func (r *RawClient) Update(ctx context.Context, voiceId string, request *types.VoiceSettings, opts ...option.RequestOption) (*core.Response[*types.EditVoiceSettingsResponseModel], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.elevenlabs.io",
	)
	endpointURL := internal.EncodeURL(baseURL+"/v1/voices/%v/settings/edit", voiceId)
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	headers.Add("Content-Type", "application/json")
	var response types.EditVoiceSettingsResponseModel
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
	return &core.Response[*types.EditVoiceSettingsResponseModel]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}
