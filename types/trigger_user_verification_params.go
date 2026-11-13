package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type TriggerUserVerificationParams struct {
	extraProperties map[string]interface{}
}

func (t *TriggerUserVerificationParams) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TriggerUserVerificationParams) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TriggerUserVerificationParams to string"
	}
	return string(jsonData)
}

func (t *TriggerUserVerificationParams) UnmarshalJSON(data []byte) error {
	type unmarshaler TriggerUserVerificationParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TriggerUserVerificationParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "smb_tool_type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TriggerUserVerificationParams) MarshalJSON() ([]byte, error) {
	type embed TriggerUserVerificationParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
