package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WebhookToolApiSchemaConfigInput struct {
	// Headers that should be included in the request
	RequestHeaders map[string]*WebhookToolApiSchemaConfigInputRequestHeadersValue `json:"request_headers,omitempty"`
	// The URL that the webhook will be sent to. May include path parameters, e.g. https://example.com/agents/{agent_id}
	Url string `json:"url" xml:"url" required:"true"`
	// The HTTP method to use for the webhook
	Method *WebhookToolApiSchemaConfigInputMethod `json:"method,omitempty" xml:"method,omitempty"`
	// Schema for path parameters, if any. The keys should match the placeholders in the URL.
	PathParamsSchema map[string]*LiteralJsonSchemaProperty `json:"path_params_schema,omitempty"`
	// Schema for any query params, if any. These will be added to end of the URL as query params. Note: properties in a query param must all be literal types
	QueryParamsSchema *QueryParamsJsonSchemaInput `json:"query_params_schema,omitempty" xml:"query_params_schema,omitempty"`
	// Schema for the body parameters, if any. Used for POST/PATCH/PUT requests. The schema should be an object which will be sent as the json body
	RequestBodySchema *ObjectJsonSchemaPropertyInput `json:"request_body_schema,omitempty" xml:"request_body_schema,omitempty"`
	// Schema describing the expected response body structure. For documentation only; not surfaced to the LLM.
	ResponseBodySchema *ObjectJsonSchemaPropertyInput `json:"response_body_schema,omitempty" xml:"response_body_schema,omitempty"`
	// Optional allow-list filter applied to the response before the LLM sees it, so large responses don't pollute the context. Defaults to the full response.
	ResponseFilter *ResponseFilter `json:"response_filter,omitempty" xml:"response_filter,omitempty"`
	// Content type for the request body. Only applies to POST/PUT/PATCH requests.
	ContentType *WebhookToolApiSchemaConfigInputContentType `json:"content_type,omitempty" xml:"content_type,omitempty"`
	// URL placeholders resolved from the auth connection (e.g. secrets injected via UrlSecretAuthConnection) rather than from path_params_schema.
	AuthResolvedParams []string `json:"auth_resolved_params,omitempty" xml:"auth_resolved_params,omitempty" uniqueItems:"true"`
	// Optional auth connection to use for authentication with this webhook
	AuthConnection  *WebhookToolApiSchemaConfigInputAuthConnection `json:"auth_connection,omitempty" xml:"auth_connection,omitempty"`
	extraProperties map[string]interface{}
}

func (w *WebhookToolApiSchemaConfigInput) GetRequestHeaders() (value map[string]*WebhookToolApiSchemaConfigInputRequestHeadersValue) {
	if w == nil {
		return
	}
	return w.RequestHeaders
}

func (w *WebhookToolApiSchemaConfigInput) GetUrl() (value string) {
	if w == nil {
		return
	}
	return w.Url
}

func (w *WebhookToolApiSchemaConfigInput) GetMethod() (value *WebhookToolApiSchemaConfigInputMethod) {
	if w == nil {
		return
	}
	return w.Method
}

func (w *WebhookToolApiSchemaConfigInput) GetPathParamsSchema() (value map[string]*LiteralJsonSchemaProperty) {
	if w == nil {
		return
	}
	return w.PathParamsSchema
}

func (w *WebhookToolApiSchemaConfigInput) GetQueryParamsSchema() (value *QueryParamsJsonSchemaInput) {
	if w == nil {
		return
	}
	return w.QueryParamsSchema
}

func (w *WebhookToolApiSchemaConfigInput) GetRequestBodySchema() (value *ObjectJsonSchemaPropertyInput) {
	if w == nil {
		return
	}
	return w.RequestBodySchema
}

func (w *WebhookToolApiSchemaConfigInput) GetResponseBodySchema() (value *ObjectJsonSchemaPropertyInput) {
	if w == nil {
		return
	}
	return w.ResponseBodySchema
}

func (w *WebhookToolApiSchemaConfigInput) GetResponseFilter() (value *ResponseFilter) {
	if w == nil {
		return
	}
	return w.ResponseFilter
}

func (w *WebhookToolApiSchemaConfigInput) GetContentType() (value *WebhookToolApiSchemaConfigInputContentType) {
	if w == nil {
		return
	}
	return w.ContentType
}

func (w *WebhookToolApiSchemaConfigInput) GetAuthResolvedParams() (value []string) {
	if w == nil {
		return
	}
	return w.AuthResolvedParams
}

func (w *WebhookToolApiSchemaConfigInput) GetAuthConnection() (value *WebhookToolApiSchemaConfigInputAuthConnection) {
	if w == nil {
		return
	}
	return w.AuthConnection
}

func (w *WebhookToolApiSchemaConfigInput) SetRequestHeaders(requestHeaders map[string]*WebhookToolApiSchemaConfigInputRequestHeadersValue) {
	w.RequestHeaders = requestHeaders
}

func (w *WebhookToolApiSchemaConfigInput) SetUrl(url string) {
	w.Url = url
}

func (w *WebhookToolApiSchemaConfigInput) SetMethod(method *WebhookToolApiSchemaConfigInputMethod) {
	w.Method = method
}

func (w *WebhookToolApiSchemaConfigInput) SetPathParamsSchema(pathParamsSchema map[string]*LiteralJsonSchemaProperty) {
	w.PathParamsSchema = pathParamsSchema
}

func (w *WebhookToolApiSchemaConfigInput) SetQueryParamsSchema(queryParamsSchema *QueryParamsJsonSchemaInput) {
	w.QueryParamsSchema = queryParamsSchema
}

func (w *WebhookToolApiSchemaConfigInput) SetRequestBodySchema(requestBodySchema *ObjectJsonSchemaPropertyInput) {
	w.RequestBodySchema = requestBodySchema
}

func (w *WebhookToolApiSchemaConfigInput) SetResponseBodySchema(responseBodySchema *ObjectJsonSchemaPropertyInput) {
	w.ResponseBodySchema = responseBodySchema
}

func (w *WebhookToolApiSchemaConfigInput) SetResponseFilter(responseFilter *ResponseFilter) {
	w.ResponseFilter = responseFilter
}

func (w *WebhookToolApiSchemaConfigInput) SetContentType(contentType *WebhookToolApiSchemaConfigInputContentType) {
	w.ContentType = contentType
}

func (w *WebhookToolApiSchemaConfigInput) SetAuthResolvedParams(authResolvedParams []string) {
	w.AuthResolvedParams = authResolvedParams
}

func (w *WebhookToolApiSchemaConfigInput) SetAuthConnection(authConnection *WebhookToolApiSchemaConfigInputAuthConnection) {
	w.AuthConnection = authConnection
}

func (w *WebhookToolApiSchemaConfigInput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WebhookToolApiSchemaConfigInput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WebhookToolApiSchemaConfigInput to string"
	}
	return string(jsonData)
}

func (w *WebhookToolApiSchemaConfigInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, w); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WebhookToolApiSchemaConfigInput) MarshalJSON() ([]byte, error) {
	type embed WebhookToolApiSchemaConfigInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
