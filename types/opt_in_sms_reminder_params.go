package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type OptInSmsReminderParams struct {
	extraProperties map[string]interface{}
}

func (o *OptInSmsReminderParams) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OptInSmsReminderParams) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OptInSmsReminderParams to string"
	}
	return string(jsonData)
}

func (o *OptInSmsReminderParams) UnmarshalJSON(data []byte) error {
	type unmarshaler OptInSmsReminderParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = OptInSmsReminderParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *o, "smb_tool_type")
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OptInSmsReminderParams) MarshalJSON() ([]byte, error) {
	type embed OptInSmsReminderParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
