package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ListAgentRulesParams struct {
	extraProperties map[string]interface{}
}

func (l *ListAgentRulesParams) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l ListAgentRulesParams) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListAgentRulesParams to string"
	}
	return string(jsonData)
}

func (l *ListAgentRulesParams) UnmarshalJSON(data []byte) error {
	type unmarshaler ListAgentRulesParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListAgentRulesParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l, "smb_tool_type")
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *ListAgentRulesParams) MarshalJSON() ([]byte, error) {
	type embed ListAgentRulesParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
