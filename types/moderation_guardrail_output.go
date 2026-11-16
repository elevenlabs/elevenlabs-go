package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ModerationGuardrailOutput struct {
	ExecutionMode   *GuardrailExecutionMode `json:"execution_mode,omitempty" xml:"execution_mode,omitempty"`
	Config          *ModerationConfig       `json:"config,omitempty" xml:"config,omitempty"`
	extraProperties map[string]interface{}
}

func (m *ModerationGuardrailOutput) GetExecutionMode() (value *GuardrailExecutionMode) {
	if m == nil {
		return
	}
	return m.ExecutionMode
}

func (m *ModerationGuardrailOutput) GetConfig() (value *ModerationConfig) {
	if m == nil {
		return
	}
	return m.Config
}

func (m *ModerationGuardrailOutput) SetExecutionMode(executionMode *GuardrailExecutionMode) {
	m.ExecutionMode = executionMode
}

func (m *ModerationGuardrailOutput) SetConfig(config *ModerationConfig) {
	m.Config = config
}

func (m *ModerationGuardrailOutput) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m ModerationGuardrailOutput) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: ModerationGuardrailOutput to string"
	}
	return string(jsonData)
}

func (m *ModerationGuardrailOutput) UnmarshalJSON(data []byte) error {
	type unmarshaler ModerationGuardrailOutput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = ModerationGuardrailOutput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *ModerationGuardrailOutput) MarshalJSON() ([]byte, error) {
	type embed ModerationGuardrailOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
