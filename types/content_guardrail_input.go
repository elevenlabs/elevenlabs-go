package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ContentGuardrailInput struct {
	ExecutionMode   *GuardrailExecutionMode             `json:"execution_mode,omitempty" xml:"execution_mode,omitempty"`
	Config          *ContentConfig                      `json:"config,omitempty" xml:"config,omitempty"`
	TriggerAction   *ContentGuardrailInputTriggerAction `json:"trigger_action,omitempty" xml:"trigger_action,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ContentGuardrailInput) GetExecutionMode() (value *GuardrailExecutionMode) {
	if c == nil {
		return
	}
	return c.ExecutionMode
}

func (c *ContentGuardrailInput) GetConfig() (value *ContentConfig) {
	if c == nil {
		return
	}
	return c.Config
}

func (c *ContentGuardrailInput) GetTriggerAction() (value *ContentGuardrailInputTriggerAction) {
	if c == nil {
		return
	}
	return c.TriggerAction
}

func (c *ContentGuardrailInput) SetExecutionMode(executionMode *GuardrailExecutionMode) {
	c.ExecutionMode = executionMode
}

func (c *ContentGuardrailInput) SetConfig(config *ContentConfig) {
	c.Config = config
}

func (c *ContentGuardrailInput) SetTriggerAction(triggerAction *ContentGuardrailInputTriggerAction) {
	c.TriggerAction = triggerAction
}

func (c *ContentGuardrailInput) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ContentGuardrailInput) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContentGuardrailInput to string"
	}
	return string(jsonData)
}

func (c *ContentGuardrailInput) UnmarshalJSON(data []byte) error {
	type unmarshaler ContentGuardrailInput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ContentGuardrailInput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ContentGuardrailInput) MarshalJSON() ([]byte, error) {
	type embed ContentGuardrailInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
