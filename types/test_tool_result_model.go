package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TestToolResultModel struct {
	Status          *string `json:"status,omitempty" xml:"status,omitempty"`
	Reason          *string `json:"reason,omitempty" xml:"reason,omitempty"`
	ToAgent         *string `json:"to_agent,omitempty" xml:"to_agent,omitempty"`
	ToNode          *string `json:"to_node,omitempty" xml:"to_node,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TestToolResultModel) GetStatus() (value *string) {
	if t == nil {
		return
	}
	return t.Status
}

func (t *TestToolResultModel) GetReason() (value *string) {
	if t == nil {
		return
	}
	return t.Reason
}

func (t *TestToolResultModel) GetToAgent() (value *string) {
	if t == nil {
		return
	}
	return t.ToAgent
}

func (t *TestToolResultModel) GetToNode() (value *string) {
	if t == nil {
		return
	}
	return t.ToNode
}

func (t *TestToolResultModel) SetStatus(status *string) {
	t.Status = status
}

func (t *TestToolResultModel) SetReason(reason *string) {
	t.Reason = reason
}

func (t *TestToolResultModel) SetToAgent(toAgent *string) {
	t.ToAgent = toAgent
}

func (t *TestToolResultModel) SetToNode(toNode *string) {
	t.ToNode = toNode
}

func (t *TestToolResultModel) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TestToolResultModel) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TestToolResultModel to string"
	}
	return string(jsonData)
}

func (t *TestToolResultModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "result_type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TestToolResultModel) MarshalJSON() ([]byte, error) {
	type embed TestToolResultModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
