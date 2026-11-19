package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ListAgentReferencesParams struct {
	extraProperties map[string]interface{}
}

func (l *ListAgentReferencesParams) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l ListAgentReferencesParams) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListAgentReferencesParams to string"
	}
	return string(jsonData)
}

func (l *ListAgentReferencesParams) UnmarshalJSON(data []byte) error {
	type unmarshaler ListAgentReferencesParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListAgentReferencesParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l, "smb_tool_type")
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *ListAgentReferencesParams) MarshalJSON() ([]byte, error) {
	type embed ListAgentReferencesParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
