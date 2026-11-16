package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ApiIntegrationWebhookOverrides struct {
	SchemaOverrides    map[string]*ApiIntegrationWebhookOverridesSchemaOverridesValue `json:"schema_overrides,omitempty"`
	ResponseFilterMode *ResponseFilterMode                                            `json:"response_filter_mode,omitempty" xml:"response_filter_mode,omitempty"`
	ResponseFilters    []string                                                       `json:"response_filters,omitempty" xml:"response_filters,omitempty"`
	RequestHeaders     map[string]*ApiIntegrationWebhookOverridesRequestHeadersValue  `json:"request_headers,omitempty"`
	extraProperties    map[string]interface{}
}

func (a *ApiIntegrationWebhookOverrides) GetSchemaOverrides() (value map[string]*ApiIntegrationWebhookOverridesSchemaOverridesValue) {
	if a == nil {
		return
	}
	return a.SchemaOverrides
}

func (a *ApiIntegrationWebhookOverrides) GetResponseFilterMode() (value *ResponseFilterMode) {
	if a == nil {
		return
	}
	return a.ResponseFilterMode
}

func (a *ApiIntegrationWebhookOverrides) GetResponseFilters() (value []string) {
	if a == nil {
		return
	}
	return a.ResponseFilters
}

func (a *ApiIntegrationWebhookOverrides) GetRequestHeaders() (value map[string]*ApiIntegrationWebhookOverridesRequestHeadersValue) {
	if a == nil {
		return
	}
	return a.RequestHeaders
}

func (a *ApiIntegrationWebhookOverrides) SetSchemaOverrides(schemaOverrides map[string]*ApiIntegrationWebhookOverridesSchemaOverridesValue) {
	a.SchemaOverrides = schemaOverrides
}

func (a *ApiIntegrationWebhookOverrides) SetResponseFilterMode(responseFilterMode *ResponseFilterMode) {
	a.ResponseFilterMode = responseFilterMode
}

func (a *ApiIntegrationWebhookOverrides) SetResponseFilters(responseFilters []string) {
	a.ResponseFilters = responseFilters
}

func (a *ApiIntegrationWebhookOverrides) SetRequestHeaders(requestHeaders map[string]*ApiIntegrationWebhookOverridesRequestHeadersValue) {
	a.RequestHeaders = requestHeaders
}

func (a *ApiIntegrationWebhookOverrides) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a ApiIntegrationWebhookOverrides) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: ApiIntegrationWebhookOverrides to string"
	}
	return string(jsonData)
}

func (a *ApiIntegrationWebhookOverrides) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *ApiIntegrationWebhookOverrides) MarshalJSON() ([]byte, error) {
	type embed ApiIntegrationWebhookOverrides
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
