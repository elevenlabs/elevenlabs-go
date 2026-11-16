package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type AgentTransferOpReplace struct {
	Type            *string `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AgentTransferOpReplace) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AgentTransferOpReplace) SetType(type_ *string) {
	a.Type = type_
}

func (a *AgentTransferOpReplace) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentTransferOpReplace) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentTransferOpReplace to string"
	}
	return string(jsonData)
}

func (a *AgentTransferOpReplace) UnmarshalJSON(data []byte) error {
	type unmarshaler AgentTransferOpReplace
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AgentTransferOpReplace(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentTransferOpReplace) MarshalJSON() ([]byte, error) {
	type embed AgentTransferOpReplace
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
