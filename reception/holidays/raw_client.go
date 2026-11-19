package holidays

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

// Get all holidays/time-off entries for the current workspace with optional filters
func (r *RawClient) List(ctx context.Context, request *reception.ListHolidaysRequest, opts ...option.RequestOption) (*core.Response[[]*types.HolidayResponseModel], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.elevenlabs.io",
	)
	endpointURL := baseURL + "/v1/reception/holidays"
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
		request = &reception.ListHolidaysRequest{}
	}
	internal.SetHeader(headers, "xi-workspace-id", request.XiWorkspaceId)
	var response []*types.HolidayResponseModel
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
	return &core.Response[[]*types.HolidayResponseModel]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       response,
	}, nil
}

// Create a new holiday
func (r *RawClient) Create(ctx context.Context, request *reception.ReceptionCreateHolidayRequest, opts ...option.RequestOption) (*core.Response[*types.HolidayResponseModel], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.elevenlabs.io",
	)
	endpointURL := baseURL + "/v1/reception/holidays"
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	headers.Add("Content-Type", "application/json")
	if request == nil {
		request = &reception.ReceptionCreateHolidayRequest{}
	}
	internal.SetHeader(headers, "xi-workspace-id", request.XiWorkspaceId)
	var response types.HolidayResponseModel
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
	return &core.Response[*types.HolidayResponseModel]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}

// Get a specific holiday by ID
func (r *RawClient) Get(ctx context.Context, holidayId string, request *reception.GetHolidaysRequest, opts ...option.RequestOption) (*core.Response[*types.HolidayResponseModel], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.elevenlabs.io",
	)
	endpointURL := internal.EncodeURL(baseURL+"/v1/reception/holidays/%v", holidayId)
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	if request == nil {
		request = &reception.GetHolidaysRequest{}
	}
	internal.SetHeader(headers, "xi-workspace-id", request.XiWorkspaceId)
	var response types.HolidayResponseModel
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
	return &core.Response[*types.HolidayResponseModel]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}

// Update an existing holiday. Only the fields sent are changed.
func (r *RawClient) Update(ctx context.Context, holidayId string, request *reception.ReceptionUpdateHolidayRequestPatch, opts ...option.RequestOption) (*core.Response[*types.HolidayResponseModel], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.elevenlabs.io",
	)
	endpointURL := internal.EncodeURL(baseURL+"/v1/reception/holidays/%v", holidayId)
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	headers.Add("Content-Type", "application/json")
	if request == nil {
		request = &reception.ReceptionUpdateHolidayRequestPatch{}
	}
	internal.SetHeader(headers, "xi-workspace-id", request.XiWorkspaceId)
	var response types.HolidayResponseModel
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
	return &core.Response[*types.HolidayResponseModel]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       &response,
	}, nil
}

// Delete a holiday
func (r *RawClient) Delete(ctx context.Context, holidayId string, request *reception.DeleteHolidaysRequest, opts ...option.RequestOption) (*core.Response[map[string]string], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.elevenlabs.io",
	)
	endpointURL := internal.EncodeURL(baseURL+"/v1/reception/holidays/%v", holidayId)
	headers := internal.MergeHeaders(
		r.options.ToHeader(),
		options.ToHeader(),
	)
	if request == nil {
		request = &reception.DeleteHolidaysRequest{}
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
