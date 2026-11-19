package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WhatsAppAuthResponse struct {
	Name          string                      `json:"name" xml:"name" required:"true"`
	Provider      *string                     `json:"provider,omitempty" xml:"provider,omitempty"`
	PhoneNumberId string                      `json:"phone_number_id" xml:"phone_number_id" required:"true"`
	Id            string                      `json:"id" xml:"id" required:"true"`
	UsedBy        *AuthConnectionDependencies `json:"used_by,omitempty" xml:"used_by,omitempty"`
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

func (w *WhatsAppAuthResponse) GetName() (value string) {
	if w == nil {
		return
	}
	return w.Name
}

func (w *WhatsAppAuthResponse) GetProvider() (value *string) {
	if w == nil {
		return
	}
	return w.Provider
}

func (w *WhatsAppAuthResponse) GetPhoneNumberId() (value string) {
	if w == nil {
		return
	}
	return w.PhoneNumberId
}

func (w *WhatsAppAuthResponse) GetId() (value string) {
	if w == nil {
		return
	}
	return w.Id
}

func (w *WhatsAppAuthResponse) GetUsedBy() (value *AuthConnectionDependencies) {
	if w == nil {
		return
	}
	return w.UsedBy
}

func (w *WhatsAppAuthResponse) GetStatus() (value *AuthConnectionStatus) {
	if w == nil {
		return
	}
	return w.Status
}

func (w *WhatsAppAuthResponse) GetStatusDetail() (value *string) {
	if w == nil {
		return
	}
	return w.StatusDetail
}

func (w *WhatsAppAuthResponse) GetStatusUpdatedAt() (value *string) {
	if w == nil {
		return
	}
	return w.StatusUpdatedAt
}

func (w *WhatsAppAuthResponse) SetName(name string) {
	w.Name = name
}

func (w *WhatsAppAuthResponse) SetProvider(provider *string) {
	w.Provider = provider
}

func (w *WhatsAppAuthResponse) SetPhoneNumberId(phoneNumberId string) {
	w.PhoneNumberId = phoneNumberId
}

func (w *WhatsAppAuthResponse) SetId(id string) {
	w.Id = id
}

func (w *WhatsAppAuthResponse) SetUsedBy(usedBy *AuthConnectionDependencies) {
	w.UsedBy = usedBy
}

func (w *WhatsAppAuthResponse) SetStatus(status *AuthConnectionStatus) {
	w.Status = status
}

func (w *WhatsAppAuthResponse) SetStatusDetail(statusDetail *string) {
	w.StatusDetail = statusDetail
}

func (w *WhatsAppAuthResponse) SetStatusUpdatedAt(statusUpdatedAt *string) {
	w.StatusUpdatedAt = statusUpdatedAt
}

func (w *WhatsAppAuthResponse) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WhatsAppAuthResponse) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WhatsAppAuthResponse to string"
	}
	return string(jsonData)
}

func (w *WhatsAppAuthResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, w); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "auth_type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WhatsAppAuthResponse) MarshalJSON() ([]byte, error) {
	type embed WhatsAppAuthResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
