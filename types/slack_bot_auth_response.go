package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Response model for the internal Slack BYO bot auth connection.
type SlackBotAuthResponse struct {
	Name     string                      `json:"name" xml:"name" required:"true"`
	Provider *string                     `json:"provider,omitempty" xml:"provider,omitempty"`
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

func (s *SlackBotAuthResponse) GetName() (value string) {
	if s == nil {
		return
	}
	return s.Name
}

func (s *SlackBotAuthResponse) GetProvider() (value *string) {
	if s == nil {
		return
	}
	return s.Provider
}

func (s *SlackBotAuthResponse) GetId() (value string) {
	if s == nil {
		return
	}
	return s.Id
}

func (s *SlackBotAuthResponse) GetUsedBy() (value *AuthConnectionDependencies) {
	if s == nil {
		return
	}
	return s.UsedBy
}

func (s *SlackBotAuthResponse) GetStatus() (value *AuthConnectionStatus) {
	if s == nil {
		return
	}
	return s.Status
}

func (s *SlackBotAuthResponse) GetStatusDetail() (value *string) {
	if s == nil {
		return
	}
	return s.StatusDetail
}

func (s *SlackBotAuthResponse) GetStatusUpdatedAt() (value *string) {
	if s == nil {
		return
	}
	return s.StatusUpdatedAt
}

func (s *SlackBotAuthResponse) SetName(name string) {
	s.Name = name
}

func (s *SlackBotAuthResponse) SetProvider(provider *string) {
	s.Provider = provider
}

func (s *SlackBotAuthResponse) SetId(id string) {
	s.Id = id
}

func (s *SlackBotAuthResponse) SetUsedBy(usedBy *AuthConnectionDependencies) {
	s.UsedBy = usedBy
}

func (s *SlackBotAuthResponse) SetStatus(status *AuthConnectionStatus) {
	s.Status = status
}

func (s *SlackBotAuthResponse) SetStatusDetail(statusDetail *string) {
	s.StatusDetail = statusDetail
}

func (s *SlackBotAuthResponse) SetStatusUpdatedAt(statusUpdatedAt *string) {
	s.StatusUpdatedAt = statusUpdatedAt
}

func (s *SlackBotAuthResponse) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SlackBotAuthResponse) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SlackBotAuthResponse to string"
	}
	return string(jsonData)
}

func (s *SlackBotAuthResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, s); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "auth_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SlackBotAuthResponse) MarshalJSON() ([]byte, error) {
	type embed SlackBotAuthResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
