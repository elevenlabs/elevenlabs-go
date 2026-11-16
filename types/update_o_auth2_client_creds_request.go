package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type UpdateOAuth2ClientCredsRequest struct {
	AuthType          *string            `json:"auth_type,omitempty" xml:"auth_type,omitempty"`
	Provider          *string            `json:"provider,omitempty" xml:"provider,omitempty"`
	ClientId          *string            `json:"client_id,omitempty" xml:"client_id,omitempty"`
	Scopes            []string           `json:"scopes,omitempty" xml:"scopes,omitempty"`
	ExtraParams       map[string]*string `json:"extra_params,omitempty"`
	BasicAuthInHeader *bool              `json:"basic_auth_in_header,omitempty" xml:"basic_auth_in_header,omitempty"`
	ClientSecret      *string            `json:"client_secret,omitempty" xml:"client_secret,omitempty"`
	CustomHeaders     map[string]*string `json:"custom_headers,omitempty"`
	extraProperties   map[string]interface{}
}

func (u *UpdateOAuth2ClientCredsRequest) GetAuthType() (value *string) {
	if u == nil {
		return
	}
	return u.AuthType
}

func (u *UpdateOAuth2ClientCredsRequest) GetProvider() (value *string) {
	if u == nil {
		return
	}
	return u.Provider
}

func (u *UpdateOAuth2ClientCredsRequest) GetClientId() (value *string) {
	if u == nil {
		return
	}
	return u.ClientId
}

func (u *UpdateOAuth2ClientCredsRequest) GetScopes() (value []string) {
	if u == nil {
		return
	}
	return u.Scopes
}

func (u *UpdateOAuth2ClientCredsRequest) GetExtraParams() (value map[string]*string) {
	if u == nil {
		return
	}
	return u.ExtraParams
}

func (u *UpdateOAuth2ClientCredsRequest) GetBasicAuthInHeader() (value *bool) {
	if u == nil {
		return
	}
	return u.BasicAuthInHeader
}

func (u *UpdateOAuth2ClientCredsRequest) GetClientSecret() (value *string) {
	if u == nil {
		return
	}
	return u.ClientSecret
}

func (u *UpdateOAuth2ClientCredsRequest) GetCustomHeaders() (value map[string]*string) {
	if u == nil {
		return
	}
	return u.CustomHeaders
}

func (u *UpdateOAuth2ClientCredsRequest) SetAuthType(authType *string) {
	u.AuthType = authType
}

func (u *UpdateOAuth2ClientCredsRequest) SetProvider(provider *string) {
	u.Provider = provider
}

func (u *UpdateOAuth2ClientCredsRequest) SetClientId(clientId *string) {
	u.ClientId = clientId
}

func (u *UpdateOAuth2ClientCredsRequest) SetScopes(scopes []string) {
	u.Scopes = scopes
}

func (u *UpdateOAuth2ClientCredsRequest) SetExtraParams(extraParams map[string]*string) {
	u.ExtraParams = extraParams
}

func (u *UpdateOAuth2ClientCredsRequest) SetBasicAuthInHeader(basicAuthInHeader *bool) {
	u.BasicAuthInHeader = basicAuthInHeader
}

func (u *UpdateOAuth2ClientCredsRequest) SetClientSecret(clientSecret *string) {
	u.ClientSecret = clientSecret
}

func (u *UpdateOAuth2ClientCredsRequest) SetCustomHeaders(customHeaders map[string]*string) {
	u.CustomHeaders = customHeaders
}

func (u *UpdateOAuth2ClientCredsRequest) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UpdateOAuth2ClientCredsRequest) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateOAuth2ClientCredsRequest to string"
	}
	return string(jsonData)
}

func (u *UpdateOAuth2ClientCredsRequest) UnmarshalJSON(data []byte) error {
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

func (u *UpdateOAuth2ClientCredsRequest) MarshalJSON() ([]byte, error) {
	type embed UpdateOAuth2ClientCredsRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
