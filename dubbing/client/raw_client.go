package client

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

// List the dubs you have access to.
func (r *RawClient) List(ctx context.Context, request *types.ListDubbingRequest, opts ...option.RequestOption) (*core.Response[*types.DubbingMetadataPageResponseModel], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.elevenlabs.io",
	)
	endpointURL := baseURL + "/v1/dubbing"
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
		request = &types.ListDubbingRequest{}
	}
	var response types.DubbingMetadataPageResponseModel
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
	return &core.Response[*types.DubbingMetadataPageResponseModel]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}

// Dubs a provided audio or video file into given language.
func (r *RawClient) Create(ctx context.Context, request *types.DubRequest, opts ...option.RequestOption) (*core.Response[*types.DoDubbingResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.elevenlabs.io",
	)
	endpointURL := baseURL + "/v1/dubbing"
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	writer := internal.NewMultipartWriter()
	if err := writer.WriteStruct(request); err != nil {
		return nil, err
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}
	headers.Set("Content-Type", writer.ContentType())
	var response types.DoDubbingResponse
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
			Request:         writer.Buffer(),
			Response:        &response,
		},
	)
	if err != nil {
		return nil, err
	}
	return &core.Response[*types.DoDubbingResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}

// Returns metadata about a dubbing project, including whether it's still in progress or not
func (r *RawClient) Get(ctx context.Context, dubbingId string, opts ...option.RequestOption) (*core.Response[*types.DubbingMetadataResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.elevenlabs.io",
	)
	endpointURL := internal.EncodeURL(baseURL+"/v1/dubbing/%v", dubbingId)
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	var response types.DubbingMetadataResponse
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
	return &core.Response[*types.DubbingMetadataResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}

// Deletes a dubbing project.
func (r *RawClient) Delete(ctx context.Context, dubbingId string, opts ...option.RequestOption) (*core.Response[*types.DeleteDubbingResponseModel], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.elevenlabs.io",
	)
	endpointURL := internal.EncodeURL(baseURL+"/v1/dubbing/%v", dubbingId)
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	var response types.DeleteDubbingResponseModel
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
	return &core.Response[*types.DeleteDubbingResponseModel]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}
