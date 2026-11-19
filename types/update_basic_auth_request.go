package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type UpdateBasicAuthRequest struct {
	AuthType        *string `json:"auth_type,omitempty" xml:"auth_type,omitempty"`
	Provider        *string `json:"provider,omitempty" xml:"provider,omitempty"`
	Username        *string `json:"username,omitempty" xml:"username,omitempty"`
	Password        *string `json:"password,omitempty" xml:"password,omitempty"`
	extraProperties map[string]interface{}
}

func (u *UpdateBasicAuthRequest) GetAuthType() (value *string) {
	if u == nil {
		return
	}
	return u.AuthType
}

func (u *UpdateBasicAuthRequest) GetProvider() (value *string) {
	if u == nil {
		return
	}
	return u.Provider
}

func (u *UpdateBasicAuthRequest) GetUsername() (value *string) {
	if u == nil {
		return
	}
	return u.Username
}

func (u *UpdateBasicAuthRequest) GetPassword() (value *string) {
	if u == nil {
		return
	}
	return u.Password
}

func (u *UpdateBasicAuthRequest) SetAuthType(authType *string) {
	u.AuthType = authType
}

func (u *UpdateBasicAuthRequest) SetProvider(provider *string) {
	u.Provider = provider
}

func (u *UpdateBasicAuthRequest) SetUsername(username *string) {
	u.Username = username
}

func (u *UpdateBasicAuthRequest) SetPassword(password *string) {
	u.Password = password
}

func (u *UpdateBasicAuthRequest) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UpdateBasicAuthRequest) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateBasicAuthRequest to string"
	}
	return string(jsonData)
}

func (u *UpdateBasicAuthRequest) UnmarshalJSON(data []byte) error {
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

func (u *UpdateBasicAuthRequest) MarshalJSON() ([]byte, error) {
	type embed UpdateBasicAuthRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
