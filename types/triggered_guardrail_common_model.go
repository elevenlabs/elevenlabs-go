package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TriggeredGuardrailCommonModel struct {
	GuardrailType GuardrailType `json:"guardrail_type" xml:"guardrail_type" required:"true"`
	GuardrailName *string       `json:"guardrail_name,omitempty" xml:"guardrail_name,omitempty"`
	// `blocking` if the agent's response waited for the guardrail, `streaming` if the guardrail ran while the response was sent.
	ExecutionMode *GuardrailExecutionMode `json:"execution_mode,omitempty" xml:"execution_mode,omitempty"`
	// Why the guardrail triggered: the LLM's rationale for custom guardrails, or the score and threshold for moderation and content guardrails.
	GuardrailDetails *TriggeredGuardrailCommonModelGuardrailDetails `json:"guardrail_details,omitempty" xml:"guardrail_details,omitempty"`
	extraProperties  map[string]interface{}
}

func (t *TriggeredGuardrailCommonModel) GetGuardrailType() (value GuardrailType) {
	if t == nil {
		return
	}
	return t.GuardrailType
}

func (t *TriggeredGuardrailCommonModel) GetGuardrailName() (value *string) {
	if t == nil {
		return
	}
	return t.GuardrailName
}

func (t *TriggeredGuardrailCommonModel) GetExecutionMode() (value *GuardrailExecutionMode) {
	if t == nil {
		return
	}
	return t.ExecutionMode
}

func (t *TriggeredGuardrailCommonModel) GetGuardrailDetails() (value *TriggeredGuardrailCommonModelGuardrailDetails) {
	if t == nil {
		return
	}
	return t.GuardrailDetails
}

func (t *TriggeredGuardrailCommonModel) SetGuardrailType(guardrailType GuardrailType) {
	t.GuardrailType = guardrailType
}

func (t *TriggeredGuardrailCommonModel) SetGuardrailName(guardrailName *string) {
	t.GuardrailName = guardrailName
}

func (t *TriggeredGuardrailCommonModel) SetExecutionMode(executionMode *GuardrailExecutionMode) {
	t.ExecutionMode = executionMode
}

func (t *TriggeredGuardrailCommonModel) SetGuardrailDetails(guardrailDetails *TriggeredGuardrailCommonModelGuardrailDetails) {
	t.GuardrailDetails = guardrailDetails
}

func (t *TriggeredGuardrailCommonModel) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TriggeredGuardrailCommonModel) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TriggeredGuardrailCommonModel to string"
	}
	return string(jsonData)
}

func (t *TriggeredGuardrailCommonModel) UnmarshalJSON(data []byte) error {
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

func (t *TriggeredGuardrailCommonModel) MarshalJSON() ([]byte, error) {
	type embed TriggeredGuardrailCommonModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
