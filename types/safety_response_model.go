package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type SafetyResponseModel struct {
	IsBlockedIvc           *bool `json:"is_blocked_ivc,omitempty" xml:"is_blocked_ivc,omitempty"`
	IsBlockedNonIvc        *bool `json:"is_blocked_non_ivc,omitempty" xml:"is_blocked_non_ivc,omitempty"`
	IgnoreSafetyEvaluation *bool `json:"ignore_safety_evaluation,omitempty" xml:"ignore_safety_evaluation,omitempty"`
	extraProperties        map[string]interface{}
}

func (s *SafetyResponseModel) GetIsBlockedIvc() (value *bool) {
	if s == nil {
		return
	}
	return s.IsBlockedIvc
}

func (s *SafetyResponseModel) GetIsBlockedNonIvc() (value *bool) {
	if s == nil {
		return
	}
	return s.IsBlockedNonIvc
}

func (s *SafetyResponseModel) GetIgnoreSafetyEvaluation() (value *bool) {
	if s == nil {
		return
	}
	return s.IgnoreSafetyEvaluation
}

func (s *SafetyResponseModel) SetIsBlockedIvc(isBlockedIvc *bool) {
	s.IsBlockedIvc = isBlockedIvc
}

func (s *SafetyResponseModel) SetIsBlockedNonIvc(isBlockedNonIvc *bool) {
	s.IsBlockedNonIvc = isBlockedNonIvc
}

func (s *SafetyResponseModel) SetIgnoreSafetyEvaluation(ignoreSafetyEvaluation *bool) {
	s.IgnoreSafetyEvaluation = ignoreSafetyEvaluation
}

func (s *SafetyResponseModel) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SafetyResponseModel) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SafetyResponseModel to string"
	}
	return string(jsonData)
}

func (s *SafetyResponseModel) UnmarshalJSON(data []byte) error {
	type unmarshaler SafetyResponseModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SafetyResponseModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SafetyResponseModel) MarshalJSON() ([]byte, error) {
	type embed SafetyResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
