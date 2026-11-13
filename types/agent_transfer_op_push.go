package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentTransferOpPush struct {
	Type            *string `json:"type,omitempty" xml:"type,omitempty"`
	ReturnNodeId    *string `json:"return_node_id,omitempty" xml:"return_node_id,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AgentTransferOpPush) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AgentTransferOpPush) GetReturnNodeId() (value *string) {
	if a == nil {
		return
	}
	return a.ReturnNodeId
}

func (a *AgentTransferOpPush) SetType(type_ *string) {
	a.Type = type_
}

func (a *AgentTransferOpPush) SetReturnNodeId(returnNodeId *string) {
	a.ReturnNodeId = returnNodeId
}

func (a *AgentTransferOpPush) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentTransferOpPush) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentTransferOpPush to string"
	}
	return string(jsonData)
}

func (a *AgentTransferOpPush) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentTransferOpPush) MarshalJSON() ([]byte, error) {
	type embed AgentTransferOpPush
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
