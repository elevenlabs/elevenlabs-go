package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Per-agent (receptionist) configuration for client identity verification codes.
type ReceptionIdentityVerificationConfig struct {
	// When enabled, the receptionist may only look up or modify a client's appointments and records after that client has completed OTP identity verification in this conversation.
	Enabled  *bool                          `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Channels []ReceptionNotificationChannel `json:"channels,omitempty" xml:"channels,omitempty" uniqueItems:"true"`
	// Twilio phone number ID used to send SMS verification codes.
	SmsPhoneNumberId *string `json:"sms_phone_number_id,omitempty" xml:"sms_phone_number_id,omitempty"`
	extraProperties  map[string]interface{}
}

func (r *ReceptionIdentityVerificationConfig) GetEnabled() (value *bool) {
	if r == nil {
		return
	}
	return r.Enabled
}

func (r *ReceptionIdentityVerificationConfig) GetChannels() (value []ReceptionNotificationChannel) {
	if r == nil {
		return
	}
	return r.Channels
}

func (r *ReceptionIdentityVerificationConfig) GetSmsPhoneNumberId() (value *string) {
	if r == nil {
		return
	}
	return r.SmsPhoneNumberId
}

func (r *ReceptionIdentityVerificationConfig) SetEnabled(enabled *bool) {
	r.Enabled = enabled
}

func (r *ReceptionIdentityVerificationConfig) SetChannels(channels []ReceptionNotificationChannel) {
	r.Channels = channels
}

func (r *ReceptionIdentityVerificationConfig) SetSmsPhoneNumberId(smsPhoneNumberId *string) {
	r.SmsPhoneNumberId = smsPhoneNumberId
}

func (r *ReceptionIdentityVerificationConfig) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReceptionIdentityVerificationConfig) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionIdentityVerificationConfig to string"
	}
	return string(jsonData)
}

func (r *ReceptionIdentityVerificationConfig) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, r); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *ReceptionIdentityVerificationConfig) MarshalJSON() ([]byte, error) {
	type embed ReceptionIdentityVerificationConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
