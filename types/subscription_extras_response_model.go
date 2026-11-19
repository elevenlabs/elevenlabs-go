package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SubscriptionExtrasResponseModel struct {
	// The concurrency of the user.
	Concurrency int `json:"concurrency" xml:"concurrency" required:"true"`
	// The Convai concurrency of the user.
	ConvaiConcurrency int `json:"convai_concurrency" xml:"convai_concurrency" required:"true"`
	// The Music concurrency of the user on enterprise plan.
	EnterpriseMusicConcurrency *int `json:"enterprise_music_concurrency,omitempty" xml:"enterprise_music_concurrency,omitempty"`
	// Custom total finetunes limit for the user.
	MusicFinetunesTotalLimit *int `json:"music_finetunes_total_limit,omitempty" xml:"music_finetunes_total_limit,omitempty"`
	// Custom monthly finetunes limit for the user.
	MusicFinetunesMonthlyLimit *int `json:"music_finetunes_monthly_limit,omitempty" xml:"music_finetunes_monthly_limit,omitempty"`
	// Custom finetunes concurrency limit for the user.
	MusicFinetunesConcurrencyLimit *int `json:"music_finetunes_concurrency_limit,omitempty" xml:"music_finetunes_concurrency_limit,omitempty"`
	// The Convai characters per minute of the user. This field is deprecated and will always return None.
	ConvaiCharsPerMinute *int `json:"convai_chars_per_minute,omitempty" xml:"convai_chars_per_minute,omitempty"`
	// The Convai ASR characters per minute of the user. This field is deprecated and will always return None.
	ConvaiAsrCharsPerMinute *int `json:"convai_asr_chars_per_minute,omitempty" xml:"convai_asr_chars_per_minute,omitempty"`
	// Whether the user's logging is disabled.
	ForceLoggingDisabled bool `json:"force_logging_disabled" xml:"force_logging_disabled" required:"true"`
	// Whether the user can request manual pro voice verification.
	CanRequestManualProVoiceVerification bool `json:"can_request_manual_pro_voice_verification" xml:"can_request_manual_pro_voice_verification" required:"true"`
	// Whether the user can bypass the voice captcha.
	CanBypassVoiceCaptcha bool                           `json:"can_bypass_voice_captcha" xml:"can_bypass_voice_captcha" required:"true"`
	Moderation            *ModerationStatusResponseModel `json:"moderation" xml:"moderation" required:"true"`
	// The unused characters rolled over from the previous period.
	UnusedCharactersRolledOverFromPreviousPeriod *int `json:"unused_characters_rolled_over_from_previous_period,omitempty" xml:"unused_characters_rolled_over_from_previous_period,omitempty"`
	// The overused characters rolled over from the previous period.
	OverusedCharactersRolledOverFromPreviousPeriod *int `json:"overused_characters_rolled_over_from_previous_period,omitempty" xml:"overused_characters_rolled_over_from_previous_period,omitempty"`
	// Data on how the subscription is being used.
	Usage           *SubscriptionUsageResponseModel `json:"usage,omitempty" xml:"usage,omitempty"`
	extraProperties map[string]interface{}
}

func (s *SubscriptionExtrasResponseModel) GetConcurrency() (value int) {
	if s == nil {
		return
	}
	return s.Concurrency
}

func (s *SubscriptionExtrasResponseModel) GetConvaiConcurrency() (value int) {
	if s == nil {
		return
	}
	return s.ConvaiConcurrency
}

func (s *SubscriptionExtrasResponseModel) GetEnterpriseMusicConcurrency() (value *int) {
	if s == nil {
		return
	}
	return s.EnterpriseMusicConcurrency
}

func (s *SubscriptionExtrasResponseModel) GetMusicFinetunesTotalLimit() (value *int) {
	if s == nil {
		return
	}
	return s.MusicFinetunesTotalLimit
}

func (s *SubscriptionExtrasResponseModel) GetMusicFinetunesMonthlyLimit() (value *int) {
	if s == nil {
		return
	}
	return s.MusicFinetunesMonthlyLimit
}

func (s *SubscriptionExtrasResponseModel) GetMusicFinetunesConcurrencyLimit() (value *int) {
	if s == nil {
		return
	}
	return s.MusicFinetunesConcurrencyLimit
}

func (s *SubscriptionExtrasResponseModel) GetConvaiCharsPerMinute() (value *int) {
	if s == nil {
		return
	}
	return s.ConvaiCharsPerMinute
}

func (s *SubscriptionExtrasResponseModel) GetConvaiAsrCharsPerMinute() (value *int) {
	if s == nil {
		return
	}
	return s.ConvaiAsrCharsPerMinute
}

func (s *SubscriptionExtrasResponseModel) GetForceLoggingDisabled() (value bool) {
	if s == nil {
		return
	}
	return s.ForceLoggingDisabled
}

func (s *SubscriptionExtrasResponseModel) GetCanRequestManualProVoiceVerification() (value bool) {
	if s == nil {
		return
	}
	return s.CanRequestManualProVoiceVerification
}

func (s *SubscriptionExtrasResponseModel) GetCanBypassVoiceCaptcha() (value bool) {
	if s == nil {
		return
	}
	return s.CanBypassVoiceCaptcha
}

func (s *SubscriptionExtrasResponseModel) GetModeration() (value *ModerationStatusResponseModel) {
	if s == nil {
		return
	}
	return s.Moderation
}

func (s *SubscriptionExtrasResponseModel) GetUnusedCharactersRolledOverFromPreviousPeriod() (value *int) {
	if s == nil {
		return
	}
	return s.UnusedCharactersRolledOverFromPreviousPeriod
}

func (s *SubscriptionExtrasResponseModel) GetOverusedCharactersRolledOverFromPreviousPeriod() (value *int) {
	if s == nil {
		return
	}
	return s.OverusedCharactersRolledOverFromPreviousPeriod
}

func (s *SubscriptionExtrasResponseModel) GetUsage() (value *SubscriptionUsageResponseModel) {
	if s == nil {
		return
	}
	return s.Usage
}

func (s *SubscriptionExtrasResponseModel) SetConcurrency(concurrency int) {
	s.Concurrency = concurrency
}

func (s *SubscriptionExtrasResponseModel) SetConvaiConcurrency(convaiConcurrency int) {
	s.ConvaiConcurrency = convaiConcurrency
}

func (s *SubscriptionExtrasResponseModel) SetEnterpriseMusicConcurrency(enterpriseMusicConcurrency *int) {
	s.EnterpriseMusicConcurrency = enterpriseMusicConcurrency
}

func (s *SubscriptionExtrasResponseModel) SetMusicFinetunesTotalLimit(musicFinetunesTotalLimit *int) {
	s.MusicFinetunesTotalLimit = musicFinetunesTotalLimit
}

func (s *SubscriptionExtrasResponseModel) SetMusicFinetunesMonthlyLimit(musicFinetunesMonthlyLimit *int) {
	s.MusicFinetunesMonthlyLimit = musicFinetunesMonthlyLimit
}

func (s *SubscriptionExtrasResponseModel) SetMusicFinetunesConcurrencyLimit(musicFinetunesConcurrencyLimit *int) {
	s.MusicFinetunesConcurrencyLimit = musicFinetunesConcurrencyLimit
}

func (s *SubscriptionExtrasResponseModel) SetConvaiCharsPerMinute(convaiCharsPerMinute *int) {
	s.ConvaiCharsPerMinute = convaiCharsPerMinute
}

func (s *SubscriptionExtrasResponseModel) SetConvaiAsrCharsPerMinute(convaiAsrCharsPerMinute *int) {
	s.ConvaiAsrCharsPerMinute = convaiAsrCharsPerMinute
}

func (s *SubscriptionExtrasResponseModel) SetForceLoggingDisabled(forceLoggingDisabled bool) {
	s.ForceLoggingDisabled = forceLoggingDisabled
}

func (s *SubscriptionExtrasResponseModel) SetCanRequestManualProVoiceVerification(canRequestManualProVoiceVerification bool) {
	s.CanRequestManualProVoiceVerification = canRequestManualProVoiceVerification
}

func (s *SubscriptionExtrasResponseModel) SetCanBypassVoiceCaptcha(canBypassVoiceCaptcha bool) {
	s.CanBypassVoiceCaptcha = canBypassVoiceCaptcha
}

func (s *SubscriptionExtrasResponseModel) SetModeration(moderation *ModerationStatusResponseModel) {
	s.Moderation = moderation
}

func (s *SubscriptionExtrasResponseModel) SetUnusedCharactersRolledOverFromPreviousPeriod(unusedCharactersRolledOverFromPreviousPeriod *int) {
	s.UnusedCharactersRolledOverFromPreviousPeriod = unusedCharactersRolledOverFromPreviousPeriod
}

func (s *SubscriptionExtrasResponseModel) SetOverusedCharactersRolledOverFromPreviousPeriod(overusedCharactersRolledOverFromPreviousPeriod *int) {
	s.OverusedCharactersRolledOverFromPreviousPeriod = overusedCharactersRolledOverFromPreviousPeriod
}

func (s *SubscriptionExtrasResponseModel) SetUsage(usage *SubscriptionUsageResponseModel) {
	s.Usage = usage
}

func (s *SubscriptionExtrasResponseModel) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SubscriptionExtrasResponseModel) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SubscriptionExtrasResponseModel to string"
	}
	return string(jsonData)
}

func (s *SubscriptionExtrasResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, s); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SubscriptionExtrasResponseModel) MarshalJSON() ([]byte, error) {
	type embed SubscriptionExtrasResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
