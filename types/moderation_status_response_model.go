package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ModerationStatusResponseModel struct {
	// Whether the user is in probation.
	IsInProbation bool `json:"is_in_probation" xml:"is_in_probation" required:"true"`
	// Whether the user's enterprise check nogo voice is enabled.
	EnterpriseCheckNogoVoice bool `json:"enterprise_check_nogo_voice" xml:"enterprise_check_nogo_voice" required:"true"`
	// Whether the user's enterprise check block nogo voice is enabled.
	EnterpriseCheckBlockNogoVoice bool `json:"enterprise_check_block_nogo_voice" xml:"enterprise_check_block_nogo_voice" required:"true"`
	// Whether the user's never live moderate is enabled.
	NeverLiveModerate bool `json:"never_live_moderate" xml:"never_live_moderate" required:"true"`
	// The number of similar voice uploads that have been blocked.
	NogoVoiceSimilarVoiceUploadCount int `json:"nogo_voice_similar_voice_upload_count" xml:"nogo_voice_similar_voice_upload_count" required:"true"`
	// Whether the user's enterprise background moderation is enabled.
	EnterpriseBackgroundModerationEnabled bool `json:"enterprise_background_moderation_enabled" xml:"enterprise_background_moderation_enabled" required:"true"`
	// Whether captcha is required when creating IVCs.
	IsIvcCaptchaRequired bool `json:"is_ivc_captcha_required" xml:"is_ivc_captcha_required" required:"true"`
	// The safety status of the user.
	SafetyStatus *ModerationStatusResponseModelSafetyStatus `json:"safety_status,omitempty" xml:"safety_status,omitempty"`
	// The warning status of the user.
	WarningStatus *ModerationStatusResponseModelWarningStatus `json:"warning_status,omitempty" xml:"warning_status,omitempty"`
	// Whether the user is on the watchlist.
	OnWatchlist     bool `json:"on_watchlist" xml:"on_watchlist" required:"true"`
	extraProperties map[string]interface{}
}

func (m *ModerationStatusResponseModel) GetIsInProbation() (value bool) {
	if m == nil {
		return
	}
	return m.IsInProbation
}

func (m *ModerationStatusResponseModel) GetEnterpriseCheckNogoVoice() (value bool) {
	if m == nil {
		return
	}
	return m.EnterpriseCheckNogoVoice
}

func (m *ModerationStatusResponseModel) GetEnterpriseCheckBlockNogoVoice() (value bool) {
	if m == nil {
		return
	}
	return m.EnterpriseCheckBlockNogoVoice
}

func (m *ModerationStatusResponseModel) GetNeverLiveModerate() (value bool) {
	if m == nil {
		return
	}
	return m.NeverLiveModerate
}

func (m *ModerationStatusResponseModel) GetNogoVoiceSimilarVoiceUploadCount() (value int) {
	if m == nil {
		return
	}
	return m.NogoVoiceSimilarVoiceUploadCount
}

func (m *ModerationStatusResponseModel) GetEnterpriseBackgroundModerationEnabled() (value bool) {
	if m == nil {
		return
	}
	return m.EnterpriseBackgroundModerationEnabled
}

func (m *ModerationStatusResponseModel) GetIsIvcCaptchaRequired() (value bool) {
	if m == nil {
		return
	}
	return m.IsIvcCaptchaRequired
}

func (m *ModerationStatusResponseModel) GetSafetyStatus() (value *ModerationStatusResponseModelSafetyStatus) {
	if m == nil {
		return
	}
	return m.SafetyStatus
}

func (m *ModerationStatusResponseModel) GetWarningStatus() (value *ModerationStatusResponseModelWarningStatus) {
	if m == nil {
		return
	}
	return m.WarningStatus
}

func (m *ModerationStatusResponseModel) GetOnWatchlist() (value bool) {
	if m == nil {
		return
	}
	return m.OnWatchlist
}

func (m *ModerationStatusResponseModel) SetIsInProbation(isInProbation bool) {
	m.IsInProbation = isInProbation
}

func (m *ModerationStatusResponseModel) SetEnterpriseCheckNogoVoice(enterpriseCheckNogoVoice bool) {
	m.EnterpriseCheckNogoVoice = enterpriseCheckNogoVoice
}

func (m *ModerationStatusResponseModel) SetEnterpriseCheckBlockNogoVoice(enterpriseCheckBlockNogoVoice bool) {
	m.EnterpriseCheckBlockNogoVoice = enterpriseCheckBlockNogoVoice
}

func (m *ModerationStatusResponseModel) SetNeverLiveModerate(neverLiveModerate bool) {
	m.NeverLiveModerate = neverLiveModerate
}

func (m *ModerationStatusResponseModel) SetNogoVoiceSimilarVoiceUploadCount(nogoVoiceSimilarVoiceUploadCount int) {
	m.NogoVoiceSimilarVoiceUploadCount = nogoVoiceSimilarVoiceUploadCount
}

func (m *ModerationStatusResponseModel) SetEnterpriseBackgroundModerationEnabled(enterpriseBackgroundModerationEnabled bool) {
	m.EnterpriseBackgroundModerationEnabled = enterpriseBackgroundModerationEnabled
}

func (m *ModerationStatusResponseModel) SetIsIvcCaptchaRequired(isIvcCaptchaRequired bool) {
	m.IsIvcCaptchaRequired = isIvcCaptchaRequired
}

func (m *ModerationStatusResponseModel) SetSafetyStatus(safetyStatus *ModerationStatusResponseModelSafetyStatus) {
	m.SafetyStatus = safetyStatus
}

func (m *ModerationStatusResponseModel) SetWarningStatus(warningStatus *ModerationStatusResponseModelWarningStatus) {
	m.WarningStatus = warningStatus
}

func (m *ModerationStatusResponseModel) SetOnWatchlist(onWatchlist bool) {
	m.OnWatchlist = onWatchlist
}

func (m *ModerationStatusResponseModel) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m ModerationStatusResponseModel) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: ModerationStatusResponseModel to string"
	}
	return string(jsonData)
}

func (m *ModerationStatusResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, m); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *ModerationStatusResponseModel) MarshalJSON() ([]byte, error) {
	type embed ModerationStatusResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
