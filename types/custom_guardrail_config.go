package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Single custom guardrail configuration
type CustomGuardrailConfig struct {
	IsEnabled *bool `json:"is_enabled,omitempty" xml:"is_enabled,omitempty"`
	// User-facing name for this guardrail
	Name string `json:"name" xml:"name" required:"true" maxLength:"500" minLength:"1"`
	// Instruction describing what to block, e.g. 'don't talk about politics'
	Prompt        string                  `json:"prompt" xml:"prompt" required:"true" maxLength:"10000" minLength:"1"`
	ExecutionMode *GuardrailExecutionMode `json:"execution_mode,omitempty" xml:"execution_mode,omitempty"`
	// LLM model to use for custom guardrail evaluation
	Model *CustomGuardrailConfigModel `json:"model,omitempty" xml:"model,omitempty"`
	// Request-time processing tier for a model.
	//
	// STANDARD: default price and throughput.
	// FAST: pays a per-token price premium for higher throughput / lower latency.
	// See OpenAI's fast mode: https://openai.com/api-fast-mode/ and, for Gemini, Vertex
	// Priority PayGo: https://docs.cloud.google.com/gemini-enterprise-agent-platform/models/priority-paygo
	ServiceTier *LlmServiceTier `json:"service_tier,omitempty" xml:"service_tier,omitempty"`
	// How much recent history the guardrail sees before the reply it evaluates, counted in user messages (the agent replies between them are included too). The guardrail always gets a single <conversation_history> transcript ending in the evaluated reply, marked 'AGENT [current reply]:'. 0 (default) adds no prior history (just that line); 1 adds the latest user message onward.
	HistoryMessageCount *int                                `json:"history_message_count,omitempty" xml:"history_message_count,omitempty" min:"0" max:"100"`
	TriggerAction       *CustomGuardrailConfigTriggerAction `json:"trigger_action,omitempty" xml:"trigger_action,omitempty"`
	// Evaluate once against the complete non-TTS response instead of cumulative partials. Requires blocking mode.
	EvaluateFullResponseOnly *bool `json:"evaluate_full_response_only,omitempty" xml:"evaluate_full_response_only,omitempty"`
	extraProperties          map[string]interface{}
}

func (c *CustomGuardrailConfig) GetIsEnabled() (value *bool) {
	if c == nil {
		return
	}
	return c.IsEnabled
}

func (c *CustomGuardrailConfig) GetName() (value string) {
	if c == nil {
		return
	}
	return c.Name
}

func (c *CustomGuardrailConfig) GetPrompt() (value string) {
	if c == nil {
		return
	}
	return c.Prompt
}

func (c *CustomGuardrailConfig) GetExecutionMode() (value *GuardrailExecutionMode) {
	if c == nil {
		return
	}
	return c.ExecutionMode
}

func (c *CustomGuardrailConfig) GetModel() (value *CustomGuardrailConfigModel) {
	if c == nil {
		return
	}
	return c.Model
}

func (c *CustomGuardrailConfig) GetServiceTier() (value *LlmServiceTier) {
	if c == nil {
		return
	}
	return c.ServiceTier
}

func (c *CustomGuardrailConfig) GetHistoryMessageCount() (value *int) {
	if c == nil {
		return
	}
	return c.HistoryMessageCount
}

func (c *CustomGuardrailConfig) GetTriggerAction() (value *CustomGuardrailConfigTriggerAction) {
	if c == nil {
		return
	}
	return c.TriggerAction
}

func (c *CustomGuardrailConfig) GetEvaluateFullResponseOnly() (value *bool) {
	if c == nil {
		return
	}
	return c.EvaluateFullResponseOnly
}

func (c *CustomGuardrailConfig) SetIsEnabled(isEnabled *bool) {
	c.IsEnabled = isEnabled
}

func (c *CustomGuardrailConfig) SetName(name string) {
	c.Name = name
}

func (c *CustomGuardrailConfig) SetPrompt(prompt string) {
	c.Prompt = prompt
}

func (c *CustomGuardrailConfig) SetExecutionMode(executionMode *GuardrailExecutionMode) {
	c.ExecutionMode = executionMode
}

func (c *CustomGuardrailConfig) SetModel(model *CustomGuardrailConfigModel) {
	c.Model = model
}

func (c *CustomGuardrailConfig) SetServiceTier(serviceTier *LlmServiceTier) {
	c.ServiceTier = serviceTier
}

func (c *CustomGuardrailConfig) SetHistoryMessageCount(historyMessageCount *int) {
	c.HistoryMessageCount = historyMessageCount
}

func (c *CustomGuardrailConfig) SetTriggerAction(triggerAction *CustomGuardrailConfigTriggerAction) {
	c.TriggerAction = triggerAction
}

func (c *CustomGuardrailConfig) SetEvaluateFullResponseOnly(evaluateFullResponseOnly *bool) {
	c.EvaluateFullResponseOnly = evaluateFullResponseOnly
}

func (c *CustomGuardrailConfig) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CustomGuardrailConfig) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CustomGuardrailConfig to string"
	}
	return string(jsonData)
}

func (c *CustomGuardrailConfig) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CustomGuardrailConfig
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CustomGuardrailConfig(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CustomGuardrailConfig) MarshalJSON() ([]byte, error) {
	type embed CustomGuardrailConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
