package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Execution-related properties for a tool (2025-11-25 only).
type ToolExecution struct {
	TaskSupport     *ToolExecutionTaskSupport `json:"taskSupport,omitempty" xml:"taskSupport,omitempty"`
	extraProperties map[string]interface{}
}

func (t *ToolExecution) GetTaskSupport() (value *ToolExecutionTaskSupport) {
	if t == nil {
		return
	}
	return t.TaskSupport
}

func (t *ToolExecution) SetTaskSupport(taskSupport *ToolExecutionTaskSupport) {
	t.TaskSupport = taskSupport
}

func (t *ToolExecution) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t ToolExecution) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: ToolExecution to string"
	}
	return string(jsonData)
}

func (t *ToolExecution) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *ToolExecution) MarshalJSON() ([]byte, error) {
	type embed ToolExecution
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
