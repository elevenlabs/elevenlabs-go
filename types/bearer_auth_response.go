package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Response model for bearer auth
type BearerAuthResponse struct {
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

func (b *BearerAuthResponse) GetName() (value string) {
	if b == nil {
		return
	}
	return b.Name
}

func (b *BearerAuthResponse) GetProvider() (value string) {
	if b == nil {
		return
	}
	return b.Provider
}

func (b *BearerAuthResponse) GetId() (value string) {
	if b == nil {
		return
	}
	return b.Id
}

func (b *BearerAuthResponse) GetUsedBy() (value *AuthConnectionDependencies) {
	if b == nil {
		return
	}
	return b.UsedBy
}

func (b *BearerAuthResponse) GetStatus() (value *AuthConnectionStatus) {
	if b == nil {
		return
	}
	return b.Status
}

func (b *BearerAuthResponse) GetStatusDetail() (value *string) {
	if b == nil {
		return
	}
	return b.StatusDetail
}

func (b *BearerAuthResponse) GetStatusUpdatedAt() (value *string) {
	if b == nil {
		return
	}
	return b.StatusUpdatedAt
}

func (b *BearerAuthResponse) SetName(name string) {
	b.Name = name
}

func (b *BearerAuthResponse) SetProvider(provider string) {
	b.Provider = provider
}

func (b *BearerAuthResponse) SetId(id string) {
	b.Id = id
}

func (b *BearerAuthResponse) SetUsedBy(usedBy *AuthConnectionDependencies) {
	b.UsedBy = usedBy
}

func (b *BearerAuthResponse) SetStatus(status *AuthConnectionStatus) {
	b.Status = status
}

func (b *BearerAuthResponse) SetStatusDetail(statusDetail *string) {
	b.StatusDetail = statusDetail
}

func (b *BearerAuthResponse) SetStatusUpdatedAt(statusUpdatedAt *string) {
	b.StatusUpdatedAt = statusUpdatedAt
}

func (b *BearerAuthResponse) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b BearerAuthResponse) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BearerAuthResponse to string"
	}
	return string(jsonData)
}

func (b *BearerAuthResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, b); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *b, "auth_type")
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	return nil
}

func (b *BearerAuthResponse) MarshalJSON() ([]byte, error) {
	type embed BearerAuthResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
