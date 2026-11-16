package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type VoiceSharingModerationCheckResponseModel struct {
	// The date the moderation check was made in Unix time.
	DateCheckedUnix *int `json:"date_checked_unix,omitempty" xml:"date_checked_unix,omitempty"`
	// The name value of the voice.
	NameValue *string `json:"name_value,omitempty" xml:"name_value,omitempty"`
	// Whether the name check was successful.
	NameCheck *bool `json:"name_check,omitempty" xml:"name_check,omitempty"`
	// The description value of the voice.
	DescriptionValue *string `json:"description_value,omitempty" xml:"description_value,omitempty"`
	// Whether the description check was successful.
	DescriptionCheck *bool `json:"description_check,omitempty" xml:"description_check,omitempty"`
	// A list of sample IDs.
	SampleIds []string `json:"sample_ids,omitempty" xml:"sample_ids,omitempty"`
	// A list of sample checks.
	SampleChecks []float64 `json:"sample_checks,omitempty" xml:"sample_checks,omitempty"`
	// A list of captcha IDs.
	CaptchaIds []string `json:"captcha_ids,omitempty" xml:"captcha_ids,omitempty"`
	// A list of CAPTCHA check values.
	CaptchaChecks   []float64 `json:"captcha_checks,omitempty" xml:"captcha_checks,omitempty"`
	extraProperties map[string]interface{}
}

func (v *VoiceSharingModerationCheckResponseModel) GetDateCheckedUnix() (value *int) {
	if v == nil {
		return
	}
	return v.DateCheckedUnix
}

func (v *VoiceSharingModerationCheckResponseModel) GetNameValue() (value *string) {
	if v == nil {
		return
	}
	return v.NameValue
}

func (v *VoiceSharingModerationCheckResponseModel) GetNameCheck() (value *bool) {
	if v == nil {
		return
	}
	return v.NameCheck
}

func (v *VoiceSharingModerationCheckResponseModel) GetDescriptionValue() (value *string) {
	if v == nil {
		return
	}
	return v.DescriptionValue
}

func (v *VoiceSharingModerationCheckResponseModel) GetDescriptionCheck() (value *bool) {
	if v == nil {
		return
	}
	return v.DescriptionCheck
}

func (v *VoiceSharingModerationCheckResponseModel) GetSampleIds() (value []string) {
	if v == nil {
		return
	}
	return v.SampleIds
}

func (v *VoiceSharingModerationCheckResponseModel) GetSampleChecks() (value []float64) {
	if v == nil {
		return
	}
	return v.SampleChecks
}

func (v *VoiceSharingModerationCheckResponseModel) GetCaptchaIds() (value []string) {
	if v == nil {
		return
	}
	return v.CaptchaIds
}

func (v *VoiceSharingModerationCheckResponseModel) GetCaptchaChecks() (value []float64) {
	if v == nil {
		return
	}
	return v.CaptchaChecks
}

func (v *VoiceSharingModerationCheckResponseModel) SetDateCheckedUnix(dateCheckedUnix *int) {
	v.DateCheckedUnix = dateCheckedUnix
}

func (v *VoiceSharingModerationCheckResponseModel) SetNameValue(nameValue *string) {
	v.NameValue = nameValue
}

func (v *VoiceSharingModerationCheckResponseModel) SetNameCheck(nameCheck *bool) {
	v.NameCheck = nameCheck
}

func (v *VoiceSharingModerationCheckResponseModel) SetDescriptionValue(descriptionValue *string) {
	v.DescriptionValue = descriptionValue
}

func (v *VoiceSharingModerationCheckResponseModel) SetDescriptionCheck(descriptionCheck *bool) {
	v.DescriptionCheck = descriptionCheck
}

func (v *VoiceSharingModerationCheckResponseModel) SetSampleIds(sampleIds []string) {
	v.SampleIds = sampleIds
}

func (v *VoiceSharingModerationCheckResponseModel) SetSampleChecks(sampleChecks []float64) {
	v.SampleChecks = sampleChecks
}

func (v *VoiceSharingModerationCheckResponseModel) SetCaptchaIds(captchaIds []string) {
	v.CaptchaIds = captchaIds
}

func (v *VoiceSharingModerationCheckResponseModel) SetCaptchaChecks(captchaChecks []float64) {
	v.CaptchaChecks = captchaChecks
}

func (v *VoiceSharingModerationCheckResponseModel) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v VoiceSharingModerationCheckResponseModel) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VoiceSharingModerationCheckResponseModel to string"
	}
	return string(jsonData)
}

func (v *VoiceSharingModerationCheckResponseModel) UnmarshalJSON(data []byte) error {
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

func (v *VoiceSharingModerationCheckResponseModel) MarshalJSON() ([]byte, error) {
	type embed VoiceSharingModerationCheckResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
