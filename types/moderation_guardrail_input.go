package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ModerationGuardrailInput struct {
	ExecutionMode   *GuardrailExecutionMode `json:"execution_mode,omitempty" xml:"execution_mode,omitempty"`
	Config          *ModerationConfig       `json:"config,omitempty" xml:"config,omitempty"`
	extraProperties map[string]interface{}
}

func (m *ModerationGuardrailInput) GetExecutionMode() (value *GuardrailExecutionMode) {
	if m == nil {
		return
	}
	return m.ExecutionMode
}

func (m *ModerationGuardrailInput) GetConfig() (value *ModerationConfig) {
	if m == nil {
		return
	}
	return m.Config
}

func (m *ModerationGuardrailInput) SetExecutionMode(executionMode *GuardrailExecutionMode) {
	m.ExecutionMode = executionMode
}

func (m *ModerationGuardrailInput) SetConfig(config *ModerationConfig) {
	m.Config = config
}

func (m *ModerationGuardrailInput) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m ModerationGuardrailInput) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: ModerationGuardrailInput to string"
	}
	return string(jsonData)
}

func (m *ModerationGuardrailInput) UnmarshalJSON(data []byte) error {
	type unmarshaler ModerationGuardrailInput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = ModerationGuardrailInput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *ModerationGuardrailInput) MarshalJSON() ([]byte, error) {
	type embed ModerationGuardrailInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
