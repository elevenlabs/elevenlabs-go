package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Safety object that has the information of safety evaluations based on used voice.
type SafetyCommonModelOutput struct {
	// Safety evaluation of the agent. Prompt and first message is taken into account.
	// The unsafe reason is provided from the evaluation
	Ivc *SafetyEvaluation `json:"ivc,omitempty" xml:"ivc,omitempty"`
	// Safety evaluation of the agent. Prompt and first message is taken into account.
	// The unsafe reason is provided from the evaluation
	NonIvc          *SafetyEvaluation `json:"non_ivc,omitempty" xml:"non_ivc,omitempty"`
	extraProperties map[string]interface{}
}

func (s *SafetyCommonModelOutput) GetIvc() (value *SafetyEvaluation) {
	if s == nil {
		return
	}
	return s.Ivc
}

func (s *SafetyCommonModelOutput) GetNonIvc() (value *SafetyEvaluation) {
	if s == nil {
		return
	}
	return s.NonIvc
}

func (s *SafetyCommonModelOutput) SetIvc(ivc *SafetyEvaluation) {
	s.Ivc = ivc
}

func (s *SafetyCommonModelOutput) SetNonIvc(nonIvc *SafetyEvaluation) {
	s.NonIvc = nonIvc
}

func (s *SafetyCommonModelOutput) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SafetyCommonModelOutput) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SafetyCommonModelOutput to string"
	}
	return string(jsonData)
}

func (s *SafetyCommonModelOutput) UnmarshalJSON(data []byte) error {
	type unmarshaler SafetyCommonModelOutput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SafetyCommonModelOutput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SafetyCommonModelOutput) MarshalJSON() ([]byte, error) {
	type embed SafetyCommonModelOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
