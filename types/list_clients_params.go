package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// List clients ordered by most recently updated, with an optional limit.
type ListClientsParams struct {
	extraProperties map[string]interface{}
}

func (l *ListClientsParams) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l ListClientsParams) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListClientsParams to string"
	}
	return string(jsonData)
}

func (l *ListClientsParams) UnmarshalJSON(data []byte) error {
	type unmarshaler ListClientsParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListClientsParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l, "smb_tool_type")
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *ListClientsParams) MarshalJSON() ([]byte, error) {
	type embed ListClientsParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
