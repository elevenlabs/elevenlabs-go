package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Safety object that has the information of safety evaluations based on used voice.
type SafetyCommonModelInput struct {
	// Safety evaluation of the agent. Prompt and first message is taken into account.
	// The unsafe reason is provided from the evaluation
	Ivc *SafetyEvaluation `json:"ivc,omitempty" xml:"ivc,omitempty"`
	// Safety evaluation of the agent. Prompt and first message is taken into account.
	// The unsafe reason is provided from the evaluation
	NonIvc          *SafetyEvaluation `json:"non_ivc,omitempty" xml:"non_ivc,omitempty"`
	extraProperties map[string]interface{}
}

func (s *SafetyCommonModelInput) GetIvc() (value *SafetyEvaluation) {
	if s == nil {
		return
	}
	return s.Ivc
}

func (s *SafetyCommonModelInput) GetNonIvc() (value *SafetyEvaluation) {
	if s == nil {
		return
	}
	return s.NonIvc
}

func (s *SafetyCommonModelInput) SetIvc(ivc *SafetyEvaluation) {
	s.Ivc = ivc
}

func (s *SafetyCommonModelInput) SetNonIvc(nonIvc *SafetyEvaluation) {
	s.NonIvc = nonIvc
}

func (s *SafetyCommonModelInput) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SafetyCommonModelInput) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SafetyCommonModelInput to string"
	}
	return string(jsonData)
}

func (s *SafetyCommonModelInput) UnmarshalJSON(data []byte) error {
	type unmarshaler SafetyCommonModelInput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SafetyCommonModelInput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SafetyCommonModelInput) MarshalJSON() ([]byte, error) {
	type embed SafetyCommonModelInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
