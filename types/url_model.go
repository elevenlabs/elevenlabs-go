package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// OCSF URL object.
//
// Spec: https://schema.ocsf.io/1.6.0/objects/url
type UrlModel struct {
	// Full URL string
	UrlString *string `json:"url_string,omitempty" xml:"url_string,omitempty"`
	// URL scheme (e.g., https)
	Scheme *string `json:"scheme,omitempty" xml:"scheme,omitempty"`
	// URL hostname
	Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty"`
	// URL port
	Port *int `json:"port,omitempty" xml:"port,omitempty"`
	// URL path
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// URL query string
	QueryString     *string `json:"query_string,omitempty" xml:"query_string,omitempty"`
	extraProperties map[string]interface{}
}

func (u *UrlModel) GetUrlString() (value *string) {
	if u == nil {
		return
	}
	return u.UrlString
}

func (u *UrlModel) GetScheme() (value *string) {
	if u == nil {
		return
	}
	return u.Scheme
}

func (u *UrlModel) GetHostname() (value *string) {
	if u == nil {
		return
	}
	return u.Hostname
}

func (u *UrlModel) GetPort() (value *int) {
	if u == nil {
		return
	}
	return u.Port
}

func (u *UrlModel) GetPath() (value *string) {
	if u == nil {
		return
	}
	return u.Path
}

func (u *UrlModel) GetQueryString() (value *string) {
	if u == nil {
		return
	}
	return u.QueryString
}

func (u *UrlModel) SetUrlString(urlString *string) {
	u.UrlString = urlString
}

func (u *UrlModel) SetScheme(scheme *string) {
	u.Scheme = scheme
}

func (u *UrlModel) SetHostname(hostname *string) {
	u.Hostname = hostname
}

func (u *UrlModel) SetPort(port *int) {
	u.Port = port
}

func (u *UrlModel) SetPath(path *string) {
	u.Path = path
}

func (u *UrlModel) SetQueryString(queryString *string) {
	u.QueryString = queryString
}

func (u *UrlModel) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UrlModel) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UrlModel to string"
	}
	return string(jsonData)
}

func (u *UrlModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, u); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UrlModel) MarshalJSON() ([]byte, error) {
	type embed UrlModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
