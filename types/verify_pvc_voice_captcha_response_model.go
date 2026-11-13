package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type VerifyPvcVoiceCaptchaResponseModel struct {
	// The status of the verify PVC captcha request. If the request was successful, the status will be 'ok'. Otherwise an error message with status 500 will be returned.
	Status          string `json:"status" xml:"status" required:"true"`
	extraProperties map[string]interface{}
}

func (v *VerifyPvcVoiceCaptchaResponseModel) GetStatus() (value string) {
	if v == nil {
		return
	}
	return v.Status
}

func (v *VerifyPvcVoiceCaptchaResponseModel) SetStatus(status string) {
	v.Status = status
}

func (v *VerifyPvcVoiceCaptchaResponseModel) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v VerifyPvcVoiceCaptchaResponseModel) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VerifyPvcVoiceCaptchaResponseModel to string"
	}
	return string(jsonData)
}

func (v *VerifyPvcVoiceCaptchaResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, v); err != nil {
		return err
	}
	type alias VerifyPvcVoiceCaptchaResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*v = VerifyPvcVoiceCaptchaResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	return nil
}

func (v *VerifyPvcVoiceCaptchaResponseModel) MarshalJSON() ([]byte, error) {
	type embed VerifyPvcVoiceCaptchaResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
