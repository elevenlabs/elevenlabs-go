package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WebhookToolApiSchemaConfigOutput struct {
	// Headers that should be included in the request
	RequestHeaders map[string]*WebhookToolApiSchemaConfigOutputRequestHeadersValue `json:"request_headers,omitempty"`
	// The URL that the webhook will be sent to. May include path parameters, e.g. https://example.com/agents/{agent_id}
	Url string `json:"url" xml:"url" required:"true"`
	// The HTTP method to use for the webhook
	Method *WebhookToolApiSchemaConfigInputMethod `json:"method,omitempty" xml:"method,omitempty"`
	// Schema for path parameters, if any. The keys should match the placeholders in the URL.
	PathParamsSchema map[string]*LiteralJsonSchemaProperty `json:"path_params_schema,omitempty"`
	// Schema for any query params, if any. These will be added to end of the URL as query params. Note: properties in a query param must all be literal types
	QueryParamsSchema *QueryParamsJsonSchemaOutput `json:"query_params_schema,omitempty" xml:"query_params_schema,omitempty"`
	// Schema for the body parameters, if any. Used for POST/PATCH/PUT requests. The schema should be an object which will be sent as the json body
	RequestBodySchema *ObjectJsonSchemaPropertyOutput `json:"request_body_schema,omitempty" xml:"request_body_schema,omitempty"`
	// Schema describing the expected response body structure. For documentation only; not surfaced to the LLM.
	ResponseBodySchema *ObjectJsonSchemaPropertyOutput `json:"response_body_schema,omitempty" xml:"response_body_schema,omitempty"`
	// Optional allow-list filter applied to the response before the LLM sees it, so large responses don't pollute the context. Defaults to the full response.
	ResponseFilter *ResponseFilter `json:"response_filter,omitempty" xml:"response_filter,omitempty"`
	// Content type for the request body. Only applies to POST/PUT/PATCH requests.
	ContentType *WebhookToolApiSchemaConfigOutputContentType `json:"content_type,omitempty" xml:"content_type,omitempty"`
	// URL placeholders resolved from the auth connection (e.g. secrets injected via UrlSecretAuthConnection) rather than from path_params_schema.
	AuthResolvedParams []string `json:"auth_resolved_params,omitempty" xml:"auth_resolved_params,omitempty"`
	// Optional auth connection to use for authentication with this webhook
	AuthConnection  *WebhookToolApiSchemaConfigOutputAuthConnection `json:"auth_connection,omitempty" xml:"auth_connection,omitempty"`
	extraProperties map[string]interface{}
}

func (w *WebhookToolApiSchemaConfigOutput) GetRequestHeaders() (value map[string]*WebhookToolApiSchemaConfigOutputRequestHeadersValue) {
	if w == nil {
		return
	}
	return w.RequestHeaders
}

func (w *WebhookToolApiSchemaConfigOutput) GetUrl() (value string) {
	if w == nil {
		return
	}
	return w.Url
}

func (w *WebhookToolApiSchemaConfigOutput) GetMethod() (value *WebhookToolApiSchemaConfigInputMethod) {
	if w == nil {
		return
	}
	return w.Method
}

func (w *WebhookToolApiSchemaConfigOutput) GetPathParamsSchema() (value map[string]*LiteralJsonSchemaProperty) {
	if w == nil {
		return
	}
	return w.PathParamsSchema
}

func (w *WebhookToolApiSchemaConfigOutput) GetQueryParamsSchema() (value *QueryParamsJsonSchemaOutput) {
	if w == nil {
		return
	}
	return w.QueryParamsSchema
}

func (w *WebhookToolApiSchemaConfigOutput) GetRequestBodySchema() (value *ObjectJsonSchemaPropertyOutput) {
	if w == nil {
		return
	}
	return w.RequestBodySchema
}

func (w *WebhookToolApiSchemaConfigOutput) GetResponseBodySchema() (value *ObjectJsonSchemaPropertyOutput) {
	if w == nil {
		return
	}
	return w.ResponseBodySchema
}

func (w *WebhookToolApiSchemaConfigOutput) GetResponseFilter() (value *ResponseFilter) {
	if w == nil {
		return
	}
	return w.ResponseFilter
}

func (w *WebhookToolApiSchemaConfigOutput) GetContentType() (value *WebhookToolApiSchemaConfigOutputContentType) {
	if w == nil {
		return
	}
	return w.ContentType
}

func (w *WebhookToolApiSchemaConfigOutput) GetAuthResolvedParams() (value []string) {
	if w == nil {
		return
	}
	return w.AuthResolvedParams
}

func (w *WebhookToolApiSchemaConfigOutput) GetAuthConnection() (value *WebhookToolApiSchemaConfigOutputAuthConnection) {
	if w == nil {
		return
	}
	return w.AuthConnection
}

func (w *WebhookToolApiSchemaConfigOutput) SetRequestHeaders(requestHeaders map[string]*WebhookToolApiSchemaConfigOutputRequestHeadersValue) {
	w.RequestHeaders = requestHeaders
}

func (w *WebhookToolApiSchemaConfigOutput) SetUrl(url string) {
	w.Url = url
}

func (w *WebhookToolApiSchemaConfigOutput) SetMethod(method *WebhookToolApiSchemaConfigInputMethod) {
	w.Method = method
}

func (w *WebhookToolApiSchemaConfigOutput) SetPathParamsSchema(pathParamsSchema map[string]*LiteralJsonSchemaProperty) {
	w.PathParamsSchema = pathParamsSchema
}

func (w *WebhookToolApiSchemaConfigOutput) SetQueryParamsSchema(queryParamsSchema *QueryParamsJsonSchemaOutput) {
	w.QueryParamsSchema = queryParamsSchema
}

func (w *WebhookToolApiSchemaConfigOutput) SetRequestBodySchema(requestBodySchema *ObjectJsonSchemaPropertyOutput) {
	w.RequestBodySchema = requestBodySchema
}

func (w *WebhookToolApiSchemaConfigOutput) SetResponseBodySchema(responseBodySchema *ObjectJsonSchemaPropertyOutput) {
	w.ResponseBodySchema = responseBodySchema
}

func (w *WebhookToolApiSchemaConfigOutput) SetResponseFilter(responseFilter *ResponseFilter) {
	w.ResponseFilter = responseFilter
}

func (w *WebhookToolApiSchemaConfigOutput) SetContentType(contentType *WebhookToolApiSchemaConfigOutputContentType) {
	w.ContentType = contentType
}

func (w *WebhookToolApiSchemaConfigOutput) SetAuthResolvedParams(authResolvedParams []string) {
	w.AuthResolvedParams = authResolvedParams
}

func (w *WebhookToolApiSchemaConfigOutput) SetAuthConnection(authConnection *WebhookToolApiSchemaConfigOutputAuthConnection) {
	w.AuthConnection = authConnection
}

func (w *WebhookToolApiSchemaConfigOutput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WebhookToolApiSchemaConfigOutput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WebhookToolApiSchemaConfigOutput to string"
	}
	return string(jsonData)
}

func (w *WebhookToolApiSchemaConfigOutput) UnmarshalJSON(data []byte) error {
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

func (w *WebhookToolApiSchemaConfigOutput) MarshalJSON() ([]byte, error) {
	type embed WebhookToolApiSchemaConfigOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
