package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type LeaveMessageParams struct {
	extraProperties map[string]interface{}
}

func (l *LeaveMessageParams) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LeaveMessageParams) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LeaveMessageParams to string"
	}
	return string(jsonData)
}

func (l *LeaveMessageParams) UnmarshalJSON(data []byte) error {
	type unmarshaler LeaveMessageParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LeaveMessageParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l, "smb_tool_type")
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *LeaveMessageParams) MarshalJSON() ([]byte, error) {
	type embed LeaveMessageParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
