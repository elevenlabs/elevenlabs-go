package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// List scheduled group sessions for a group service in a date range.
//
// Group services are scheduled in advance (e.g. yoga classes, workshops) and
// callers register against an existing session. Use this for group services;
// use “check_service_availability“ for appointment / rental services.
type ListGroupSessionsParams struct {
	extraProperties map[string]interface{}
}

func (l *ListGroupSessionsParams) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l ListGroupSessionsParams) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListGroupSessionsParams to string"
	}
	return string(jsonData)
}

func (l *ListGroupSessionsParams) UnmarshalJSON(data []byte) error {
	type unmarshaler ListGroupSessionsParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListGroupSessionsParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l, "smb_tool_type")
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *ListGroupSessionsParams) MarshalJSON() ([]byte, error) {
	type embed ListGroupSessionsParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
