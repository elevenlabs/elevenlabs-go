package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CustomLlm struct {
	// The URL of the Chat Completions compatible endpoint
	Url string `json:"url" xml:"url" required:"true"`
	// The model ID to be used if URL serves multiple models
	ModelId *string `json:"model_id,omitempty" xml:"model_id,omitempty"`
	// The API key for authentication. Either a workspace secret reference {'secret_id': '...'} or an environment variable reference {'env_var_label': '...'}.
	ApiKey *CustomLlmApiKey `json:"api_key,omitempty" xml:"api_key,omitempty"`
	// Optional workspace auth connection for authentication. Only auth connections that produce an Authorization Bearer token are supported; Basic auth, mTLS, custom header, and URL secret auth connections are not supported.
	AuthConnection *CustomLlmAuthConnection `json:"auth_connection,omitempty" xml:"auth_connection,omitempty"`
	// Optional mTLS auth connection. When set, its client certificate is used for the TLS handshake with the LLM endpoint and, when auth_connection is OAuth2 client credentials or OAuth2 JWT, with its token endpoint too. Application auth stays on auth_connection or api_key. auth_connection must then be OAuth2 client credentials, OAuth2 JWT or bearer.
	MtlsAuthConnection *CustomLlmMtlsAuthConnection `json:"mtls_auth_connection,omitempty" xml:"mtls_auth_connection,omitempty"`
	// Headers that should be included in the request
	RequestHeaders map[string]*CustomLlmRequestHeadersValue `json:"request_headers,omitempty"`
	// The API version to use for the request
	ApiVersion      *string           `json:"api_version,omitempty" xml:"api_version,omitempty"`
	ApiType         *CustomLlmapiType `json:"api_type,omitempty" xml:"api_type,omitempty"`
	extraProperties map[string]interface{}
}

func (c *CustomLlm) GetUrl() (value string) {
	if c == nil {
		return
	}
	return c.Url
}

func (c *CustomLlm) GetModelId() (value *string) {
	if c == nil {
		return
	}
	return c.ModelId
}

func (c *CustomLlm) GetApiKey() (value *CustomLlmApiKey) {
	if c == nil {
		return
	}
	return c.ApiKey
}

func (c *CustomLlm) GetAuthConnection() (value *CustomLlmAuthConnection) {
	if c == nil {
		return
	}
	return c.AuthConnection
}

func (c *CustomLlm) GetMtlsAuthConnection() (value *CustomLlmMtlsAuthConnection) {
	if c == nil {
		return
	}
	return c.MtlsAuthConnection
}

func (c *CustomLlm) GetRequestHeaders() (value map[string]*CustomLlmRequestHeadersValue) {
	if c == nil {
		return
	}
	return c.RequestHeaders
}

func (c *CustomLlm) GetApiVersion() (value *string) {
	if c == nil {
		return
	}
	return c.ApiVersion
}

func (c *CustomLlm) GetApiType() (value *CustomLlmapiType) {
	if c == nil {
		return
	}
	return c.ApiType
}

func (c *CustomLlm) SetUrl(url string) {
	c.Url = url
}

func (c *CustomLlm) SetModelId(modelId *string) {
	c.ModelId = modelId
}

func (c *CustomLlm) SetApiKey(apiKey *CustomLlmApiKey) {
	c.ApiKey = apiKey
}

func (c *CustomLlm) SetAuthConnection(authConnection *CustomLlmAuthConnection) {
	c.AuthConnection = authConnection
}

func (c *CustomLlm) SetMtlsAuthConnection(mtlsAuthConnection *CustomLlmMtlsAuthConnection) {
	c.MtlsAuthConnection = mtlsAuthConnection
}

func (c *CustomLlm) SetRequestHeaders(requestHeaders map[string]*CustomLlmRequestHeadersValue) {
	c.RequestHeaders = requestHeaders
}

func (c *CustomLlm) SetApiVersion(apiVersion *string) {
	c.ApiVersion = apiVersion
}

func (c *CustomLlm) SetApiType(apiType *CustomLlmapiType) {
	c.ApiType = apiType
}

func (c *CustomLlm) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CustomLlm) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CustomLlm to string"
	}
	return string(jsonData)
}

func (c *CustomLlm) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CustomLlm) MarshalJSON() ([]byte, error) {
	type embed CustomLlm
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
