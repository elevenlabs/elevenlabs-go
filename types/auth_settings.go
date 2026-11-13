package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AuthSettings struct {
	// If set to true, starting a conversation with an agent will require a signed token
	EnableAuth *bool `json:"enable_auth,omitempty" xml:"enable_auth,omitempty"`
	// A list of hosts that are allowed to start conversations with the agent
	Allowlist []*AllowlistItem `json:"allowlist,omitempty" xml:"allowlist,omitempty"`
	// When enabled, connections with no origin header will be rejected. If the allowlist is empty, this option has no effect.
	RequireOriginHeader *bool `json:"require_origin_header,omitempty" xml:"require_origin_header,omitempty"`
	// A shareable token that can be used to start a conversation with the agent
	ShareableToken  *string `json:"shareable_token,omitempty" xml:"shareable_token,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AuthSettings) GetEnableAuth() (value *bool) {
	if a == nil {
		return
	}
	return a.EnableAuth
}

func (a *AuthSettings) GetAllowlist() (value []*AllowlistItem) {
	if a == nil {
		return
	}
	return a.Allowlist
}

func (a *AuthSettings) GetRequireOriginHeader() (value *bool) {
	if a == nil {
		return
	}
	return a.RequireOriginHeader
}

func (a *AuthSettings) GetShareableToken() (value *string) {
	if a == nil {
		return
	}
	return a.ShareableToken
}

func (a *AuthSettings) SetEnableAuth(enableAuth *bool) {
	a.EnableAuth = enableAuth
}

func (a *AuthSettings) SetAllowlist(allowlist []*AllowlistItem) {
	a.Allowlist = allowlist
}

func (a *AuthSettings) SetRequireOriginHeader(requireOriginHeader *bool) {
	a.RequireOriginHeader = requireOriginHeader
}

func (a *AuthSettings) SetShareableToken(shareableToken *string) {
	a.ShareableToken = shareableToken
}

func (a *AuthSettings) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AuthSettings) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AuthSettings to string"
	}
	return string(jsonData)
}

func (a *AuthSettings) UnmarshalJSON(data []byte) error {
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

func (a *AuthSettings) MarshalJSON() ([]byte, error) {
	type embed AuthSettings
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
