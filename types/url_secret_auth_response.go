package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type UrlSecretAuthResponse struct {
	Name     string                      `json:"name" xml:"name" required:"true"`
	Provider string                      `json:"provider" xml:"provider" required:"true"`
	Id       string                      `json:"id" xml:"id" required:"true"`
	UsedBy   *AuthConnectionDependencies `json:"used_by,omitempty" xml:"used_by,omitempty"`
	// Single status field shared by every auth type's stored credential.
	//
	// OAuth values (``REFRESH_FAILED``, ``REVOKED``) are written by the OAuth
	// token-manager refresh path. ``CREDENTIAL_INVALID`` is written by the
	// tool execution path when an upstream response matches a credential's
	// ``failure_signatures`` entry (Bearer, Basic auth, etc.).
	Status          *AuthConnectionStatus `json:"status,omitempty" xml:"status,omitempty"`
	StatusDetail    *string               `json:"status_detail,omitempty" xml:"status_detail,omitempty"`
	StatusUpdatedAt *string               `json:"status_updated_at,omitempty" xml:"status_updated_at,omitempty"`
	extraProperties map[string]interface{}
}

func (u *UrlSecretAuthResponse) GetName() (value string) {
	if u == nil {
		return
	}
	return u.Name
}

func (u *UrlSecretAuthResponse) GetProvider() (value string) {
	if u == nil {
		return
	}
	return u.Provider
}

func (u *UrlSecretAuthResponse) GetId() (value string) {
	if u == nil {
		return
	}
	return u.Id
}

func (u *UrlSecretAuthResponse) GetUsedBy() (value *AuthConnectionDependencies) {
	if u == nil {
		return
	}
	return u.UsedBy
}

func (u *UrlSecretAuthResponse) GetStatus() (value *AuthConnectionStatus) {
	if u == nil {
		return
	}
	return u.Status
}

func (u *UrlSecretAuthResponse) GetStatusDetail() (value *string) {
	if u == nil {
		return
	}
	return u.StatusDetail
}

func (u *UrlSecretAuthResponse) GetStatusUpdatedAt() (value *string) {
	if u == nil {
		return
	}
	return u.StatusUpdatedAt
}

func (u *UrlSecretAuthResponse) SetName(name string) {
	u.Name = name
}

func (u *UrlSecretAuthResponse) SetProvider(provider string) {
	u.Provider = provider
}

func (u *UrlSecretAuthResponse) SetId(id string) {
	u.Id = id
}

func (u *UrlSecretAuthResponse) SetUsedBy(usedBy *AuthConnectionDependencies) {
	u.UsedBy = usedBy
}

func (u *UrlSecretAuthResponse) SetStatus(status *AuthConnectionStatus) {
	u.Status = status
}

func (u *UrlSecretAuthResponse) SetStatusDetail(statusDetail *string) {
	u.StatusDetail = statusDetail
}

func (u *UrlSecretAuthResponse) SetStatusUpdatedAt(statusUpdatedAt *string) {
	u.StatusUpdatedAt = statusUpdatedAt
}

func (u *UrlSecretAuthResponse) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UrlSecretAuthResponse) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UrlSecretAuthResponse to string"
	}
	return string(jsonData)
}

func (u *UrlSecretAuthResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, u); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *u, "auth_type")
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UrlSecretAuthResponse) MarshalJSON() ([]byte, error) {
	type embed UrlSecretAuthResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
