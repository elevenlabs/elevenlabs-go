package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type UpdateBearerAuthRequest struct {
	AuthType        *string `json:"auth_type,omitempty" xml:"auth_type,omitempty"`
	Provider        *string `json:"provider,omitempty" xml:"provider,omitempty"`
	Token           *string `json:"token,omitempty" xml:"token,omitempty"`
	extraProperties map[string]interface{}
}

func (u *UpdateBearerAuthRequest) GetAuthType() (value *string) {
	if u == nil {
		return
	}
	return u.AuthType
}

func (u *UpdateBearerAuthRequest) GetProvider() (value *string) {
	if u == nil {
		return
	}
	return u.Provider
}

func (u *UpdateBearerAuthRequest) GetToken() (value *string) {
	if u == nil {
		return
	}
	return u.Token
}

func (u *UpdateBearerAuthRequest) SetAuthType(authType *string) {
	u.AuthType = authType
}

func (u *UpdateBearerAuthRequest) SetProvider(provider *string) {
	u.Provider = provider
}

func (u *UpdateBearerAuthRequest) SetToken(token *string) {
	u.Token = token
}

func (u *UpdateBearerAuthRequest) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UpdateBearerAuthRequest) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateBearerAuthRequest to string"
	}
	return string(jsonData)
}

func (u *UpdateBearerAuthRequest) UnmarshalJSON(data []byte) error {
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

func (u *UpdateBearerAuthRequest) MarshalJSON() ([]byte, error) {
	type embed UpdateBearerAuthRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
