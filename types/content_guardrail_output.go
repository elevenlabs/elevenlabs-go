package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ContentGuardrailOutput struct {
	ExecutionMode   *GuardrailExecutionMode              `json:"execution_mode,omitempty" xml:"execution_mode,omitempty"`
	Config          *ContentConfig                       `json:"config,omitempty" xml:"config,omitempty"`
	TriggerAction   *ContentGuardrailOutputTriggerAction `json:"trigger_action,omitempty" xml:"trigger_action,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ContentGuardrailOutput) GetExecutionMode() (value *GuardrailExecutionMode) {
	if c == nil {
		return
	}
	return c.ExecutionMode
}

func (c *ContentGuardrailOutput) GetConfig() (value *ContentConfig) {
	if c == nil {
		return
	}
	return c.Config
}

func (c *ContentGuardrailOutput) GetTriggerAction() (value *ContentGuardrailOutputTriggerAction) {
	if c == nil {
		return
	}
	return c.TriggerAction
}

func (c *ContentGuardrailOutput) SetExecutionMode(executionMode *GuardrailExecutionMode) {
	c.ExecutionMode = executionMode
}

func (c *ContentGuardrailOutput) SetConfig(config *ContentConfig) {
	c.Config = config
}

func (c *ContentGuardrailOutput) SetTriggerAction(triggerAction *ContentGuardrailOutputTriggerAction) {
	c.TriggerAction = triggerAction
}

func (c *ContentGuardrailOutput) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ContentGuardrailOutput) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContentGuardrailOutput to string"
	}
	return string(jsonData)
}

func (c *ContentGuardrailOutput) UnmarshalJSON(data []byte) error {
	type unmarshaler ContentGuardrailOutput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ContentGuardrailOutput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ContentGuardrailOutput) MarshalJSON() ([]byte, error) {
	type embed ContentGuardrailOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
