package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type OptOutSmsReminderParams struct {
	extraProperties map[string]interface{}
}

func (o *OptOutSmsReminderParams) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OptOutSmsReminderParams) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OptOutSmsReminderParams to string"
	}
	return string(jsonData)
}

func (o *OptOutSmsReminderParams) UnmarshalJSON(data []byte) error {
	type unmarshaler OptOutSmsReminderParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = OptOutSmsReminderParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *o, "smb_tool_type")
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OptOutSmsReminderParams) MarshalJSON() ([]byte, error) {
	type embed OptOutSmsReminderParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
