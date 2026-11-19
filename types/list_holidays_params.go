package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ListHolidaysParams struct {
	extraProperties map[string]interface{}
}

func (l *ListHolidaysParams) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l ListHolidaysParams) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListHolidaysParams to string"
	}
	return string(jsonData)
}

func (l *ListHolidaysParams) UnmarshalJSON(data []byte) error {
	type unmarshaler ListHolidaysParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListHolidaysParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l, "smb_tool_type")
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *ListHolidaysParams) MarshalJSON() ([]byte, error) {
	type embed ListHolidaysParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
