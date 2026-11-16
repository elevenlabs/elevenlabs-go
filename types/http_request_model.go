package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// HTTP request details.
//
// Spec: https://schema.ocsf.io/1.6.0/objects/http_request
type HttpRequestModel struct {
	// HTTP method (GET, POST, etc.)
	HttpMethod string `json:"http_method" xml:"http_method" required:"true"`
	// OCSF URL object.
	//
	// Spec: https://schema.ocsf.io/1.6.0/objects/url
	Url *UrlModel `json:"url" xml:"url" required:"true"`
	// User agent string
	UserAgent *string `json:"user_agent,omitempty" xml:"user_agent,omitempty"`
	// X-Forwarded-For header as a list
	XForwardedFor   []string `json:"x_forwarded_for,omitempty" xml:"x_forwarded_for,omitempty"`
	extraProperties map[string]interface{}
}

func (h *HttpRequestModel) GetHttpMethod() (value string) {
	if h == nil {
		return
	}
	return h.HttpMethod
}

func (h *HttpRequestModel) GetUrl() (value *UrlModel) {
	if h == nil {
		return
	}
	return h.Url
}

func (h *HttpRequestModel) GetUserAgent() (value *string) {
	if h == nil {
		return
	}
	return h.UserAgent
}

func (h *HttpRequestModel) GetXForwardedFor() (value []string) {
	if h == nil {
		return
	}
	return h.XForwardedFor
}

func (h *HttpRequestModel) SetHttpMethod(httpMethod string) {
	h.HttpMethod = httpMethod
}

func (h *HttpRequestModel) SetUrl(url *UrlModel) {
	h.Url = url
}

func (h *HttpRequestModel) SetUserAgent(userAgent *string) {
	h.UserAgent = userAgent
}

func (h *HttpRequestModel) SetXForwardedFor(xForwardedFor []string) {
	h.XForwardedFor = xForwardedFor
}

func (h *HttpRequestModel) GetExtraProperties() map[string]interface{} {
	if h == nil {
		return nil
	}
	return h.extraProperties
}

func (h HttpRequestModel) String() string {
	jsonData, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		return "error converting struct: HttpRequestModel to string"
	}
	return string(jsonData)
}

func (h *HttpRequestModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, h); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *h)
	if err != nil {
		return err
	}
	h.extraProperties = extraProperties
	return nil
}

func (h *HttpRequestModel) MarshalJSON() ([]byte, error) {
	type embed HttpRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*h),
	}
	return json.Marshal(marshaler)
}
