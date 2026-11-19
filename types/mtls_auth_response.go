package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Response model for mTLS auth connections.
type MtlsAuthResponse struct {
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

func (m *MtlsAuthResponse) GetName() (value string) {
	if m == nil {
		return
	}
	return m.Name
}

func (m *MtlsAuthResponse) GetProvider() (value string) {
	if m == nil {
		return
	}
	return m.Provider
}

func (m *MtlsAuthResponse) GetId() (value string) {
	if m == nil {
		return
	}
	return m.Id
}

func (m *MtlsAuthResponse) GetUsedBy() (value *AuthConnectionDependencies) {
	if m == nil {
		return
	}
	return m.UsedBy
}

func (m *MtlsAuthResponse) GetStatus() (value *AuthConnectionStatus) {
	if m == nil {
		return
	}
	return m.Status
}

func (m *MtlsAuthResponse) GetStatusDetail() (value *string) {
	if m == nil {
		return
	}
	return m.StatusDetail
}

func (m *MtlsAuthResponse) GetStatusUpdatedAt() (value *string) {
	if m == nil {
		return
	}
	return m.StatusUpdatedAt
}

func (m *MtlsAuthResponse) SetName(name string) {
	m.Name = name
}

func (m *MtlsAuthResponse) SetProvider(provider string) {
	m.Provider = provider
}

func (m *MtlsAuthResponse) SetId(id string) {
	m.Id = id
}

func (m *MtlsAuthResponse) SetUsedBy(usedBy *AuthConnectionDependencies) {
	m.UsedBy = usedBy
}

func (m *MtlsAuthResponse) SetStatus(status *AuthConnectionStatus) {
	m.Status = status
}

func (m *MtlsAuthResponse) SetStatusDetail(statusDetail *string) {
	m.StatusDetail = statusDetail
}

func (m *MtlsAuthResponse) SetStatusUpdatedAt(statusUpdatedAt *string) {
	m.StatusUpdatedAt = statusUpdatedAt
}

func (m *MtlsAuthResponse) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m MtlsAuthResponse) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MtlsAuthResponse to string"
	}
	return string(jsonData)
}

func (m *MtlsAuthResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, m); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *m, "auth_type")
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *MtlsAuthResponse) MarshalJSON() ([]byte, error) {
	type embed MtlsAuthResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
