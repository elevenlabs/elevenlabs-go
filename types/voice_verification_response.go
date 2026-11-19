package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type VoiceVerificationResponse struct {
	// Whether the voice requires verification.
	RequiresVerification bool `json:"requires_verification" xml:"requires_verification" required:"true"`
	// Whether the voice has been verified.
	IsVerified bool `json:"is_verified" xml:"is_verified" required:"true"`
	// List of verification failures.
	VerificationFailures []string `json:"verification_failures" xml:"verification_failures" required:"true"`
	// The number of verification attempts.
	VerificationAttemptsCount int `json:"verification_attempts_count" xml:"verification_attempts_count" required:"true"`
	// The language of the voice.
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// Number of times a verification was attempted.
	VerificationAttempts []*VerificationAttemptResponse `json:"verification_attempts,omitempty" xml:"verification_attempts,omitempty"`
	extraProperties      map[string]interface{}
}

func (v *VoiceVerificationResponse) GetRequiresVerification() (value bool) {
	if v == nil {
		return
	}
	return v.RequiresVerification
}

func (v *VoiceVerificationResponse) GetIsVerified() (value bool) {
	if v == nil {
		return
	}
	return v.IsVerified
}

func (v *VoiceVerificationResponse) GetVerificationFailures() (value []string) {
	if v == nil {
		return
	}
	return v.VerificationFailures
}

func (v *VoiceVerificationResponse) GetVerificationAttemptsCount() (value int) {
	if v == nil {
		return
	}
	return v.VerificationAttemptsCount
}

func (v *VoiceVerificationResponse) GetLanguage() (value *string) {
	if v == nil {
		return
	}
	return v.Language
}

func (v *VoiceVerificationResponse) GetVerificationAttempts() (value []*VerificationAttemptResponse) {
	if v == nil {
		return
	}
	return v.VerificationAttempts
}

func (v *VoiceVerificationResponse) SetRequiresVerification(requiresVerification bool) {
	v.RequiresVerification = requiresVerification
}

func (v *VoiceVerificationResponse) SetIsVerified(isVerified bool) {
	v.IsVerified = isVerified
}

func (v *VoiceVerificationResponse) SetVerificationFailures(verificationFailures []string) {
	v.VerificationFailures = verificationFailures
}

func (v *VoiceVerificationResponse) SetVerificationAttemptsCount(verificationAttemptsCount int) {
	v.VerificationAttemptsCount = verificationAttemptsCount
}

func (v *VoiceVerificationResponse) SetLanguage(language *string) {
	v.Language = language
}

func (v *VoiceVerificationResponse) SetVerificationAttempts(verificationAttempts []*VerificationAttemptResponse) {
	v.VerificationAttempts = verificationAttempts
}

func (v *VoiceVerificationResponse) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v VoiceVerificationResponse) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VoiceVerificationResponse to string"
	}
	return string(jsonData)
}

func (v *VoiceVerificationResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, v); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	return nil
}

func (v *VoiceVerificationResponse) MarshalJSON() ([]byte, error) {
	type embed VoiceVerificationResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
