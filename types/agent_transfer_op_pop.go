package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type AgentTransferOpPop struct {
	Type            *string `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AgentTransferOpPop) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AgentTransferOpPop) SetType(type_ *string) {
	a.Type = type_
}

func (a *AgentTransferOpPop) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentTransferOpPop) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentTransferOpPop to string"
	}
	return string(jsonData)
}

func (a *AgentTransferOpPop) UnmarshalJSON(data []byte) error {
	type unmarshaler AgentTransferOpPop
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AgentTransferOpPop(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentTransferOpPop) MarshalJSON() ([]byte, error) {
	type embed AgentTransferOpPop
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
