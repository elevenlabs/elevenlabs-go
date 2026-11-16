package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Safety evaluation of the agent. Prompt and first message is taken into account.
// The unsafe reason is provided from the evaluation
type SafetyEvaluation struct {
	IsUnsafe            *bool        `json:"is_unsafe,omitempty" xml:"is_unsafe,omitempty"`
	LlmReason           *string      `json:"llm_reason,omitempty" xml:"llm_reason,omitempty"`
	SafetyPromptVersion *int         `json:"safety_prompt_version,omitempty" xml:"safety_prompt_version,omitempty"`
	MatchedRuleId       []SafetyRule `json:"matched_rule_id,omitempty" xml:"matched_rule_id,omitempty"`
	extraProperties     map[string]interface{}
}

func (s *SafetyEvaluation) GetIsUnsafe() (value *bool) {
	if s == nil {
		return
	}
	return s.IsUnsafe
}

func (s *SafetyEvaluation) GetLlmReason() (value *string) {
	if s == nil {
		return
	}
	return s.LlmReason
}

func (s *SafetyEvaluation) GetSafetyPromptVersion() (value *int) {
	if s == nil {
		return
	}
	return s.SafetyPromptVersion
}

func (s *SafetyEvaluation) GetMatchedRuleId() (value []SafetyRule) {
	if s == nil {
		return
	}
	return s.MatchedRuleId
}

func (s *SafetyEvaluation) SetIsUnsafe(isUnsafe *bool) {
	s.IsUnsafe = isUnsafe
}

func (s *SafetyEvaluation) SetLlmReason(llmReason *string) {
	s.LlmReason = llmReason
}

func (s *SafetyEvaluation) SetSafetyPromptVersion(safetyPromptVersion *int) {
	s.SafetyPromptVersion = safetyPromptVersion
}

func (s *SafetyEvaluation) SetMatchedRuleId(matchedRuleId []SafetyRule) {
	s.MatchedRuleId = matchedRuleId
}

func (s *SafetyEvaluation) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SafetyEvaluation) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SafetyEvaluation to string"
	}
	return string(jsonData)
}

func (s *SafetyEvaluation) UnmarshalJSON(data []byte) error {
	type unmarshaler SafetyEvaluation
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SafetyEvaluation(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SafetyEvaluation) MarshalJSON() ([]byte, error) {
	type embed SafetyEvaluation
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
