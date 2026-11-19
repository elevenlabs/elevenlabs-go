package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type SmbClientAccessConfig struct {
	EnableSecureMode                 *bool `json:"enable_secure_mode,omitempty" xml:"enable_secure_mode,omitempty"`
	IdentityVerificationEnabled      *bool `json:"identity_verification_enabled,omitempty" xml:"identity_verification_enabled,omitempty"`
	IdentityVerificationEmailEnabled *bool `json:"identity_verification_email_enabled,omitempty" xml:"identity_verification_email_enabled,omitempty"`
	IdentityVerificationSmsEnabled   *bool `json:"identity_verification_sms_enabled,omitempty" xml:"identity_verification_sms_enabled,omitempty"`
	extraProperties                  map[string]interface{}
}

func (s *SmbClientAccessConfig) GetEnableSecureMode() (value *bool) {
	if s == nil {
		return
	}
	return s.EnableSecureMode
}

func (s *SmbClientAccessConfig) GetIdentityVerificationEnabled() (value *bool) {
	if s == nil {
		return
	}
	return s.IdentityVerificationEnabled
}

func (s *SmbClientAccessConfig) GetIdentityVerificationEmailEnabled() (value *bool) {
	if s == nil {
		return
	}
	return s.IdentityVerificationEmailEnabled
}

func (s *SmbClientAccessConfig) GetIdentityVerificationSmsEnabled() (value *bool) {
	if s == nil {
		return
	}
	return s.IdentityVerificationSmsEnabled
}

func (s *SmbClientAccessConfig) SetEnableSecureMode(enableSecureMode *bool) {
	s.EnableSecureMode = enableSecureMode
}

func (s *SmbClientAccessConfig) SetIdentityVerificationEnabled(identityVerificationEnabled *bool) {
	s.IdentityVerificationEnabled = identityVerificationEnabled
}

func (s *SmbClientAccessConfig) SetIdentityVerificationEmailEnabled(identityVerificationEmailEnabled *bool) {
	s.IdentityVerificationEmailEnabled = identityVerificationEmailEnabled
}

func (s *SmbClientAccessConfig) SetIdentityVerificationSmsEnabled(identityVerificationSmsEnabled *bool) {
	s.IdentityVerificationSmsEnabled = identityVerificationSmsEnabled
}

func (s *SmbClientAccessConfig) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SmbClientAccessConfig) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SmbClientAccessConfig to string"
	}
	return string(jsonData)
}

func (s *SmbClientAccessConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler SmbClientAccessConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SmbClientAccessConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SmbClientAccessConfig) MarshalJSON() ([]byte, error) {
	type embed SmbClientAccessConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
