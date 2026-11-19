package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ValidateUserVerificationCodeParams struct {
	extraProperties map[string]interface{}
}

func (v *ValidateUserVerificationCodeParams) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v ValidateUserVerificationCodeParams) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: ValidateUserVerificationCodeParams to string"
	}
	return string(jsonData)
}

func (v *ValidateUserVerificationCodeParams) UnmarshalJSON(data []byte) error {
	type unmarshaler ValidateUserVerificationCodeParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = ValidateUserVerificationCodeParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *v, "smb_tool_type")
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	return nil
}

func (v *ValidateUserVerificationCodeParams) MarshalJSON() ([]byte, error) {
	type embed ValidateUserVerificationCodeParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
