package types

import (
	"encoding/json"
	"fmt"
)

type SpeechEngineConfigRequestHeadersValue struct {
	String                string
	ConvAiSecretLocator   *ConvAiSecretLocator
	ConvAiDynamicVariable *ConvAiDynamicVariable
	typ                   string
}

func (s *SpeechEngineConfigRequestHeadersValue) GetString() string {
	if s == nil {
		return ""
	}
	return s.String
}

func (s *SpeechEngineConfigRequestHeadersValue) GetConvAiSecretLocator() *ConvAiSecretLocator {
	if s == nil {
		return nil
	}
	if s.ConvAiSecretLocator == nil {
		return nil
	}
	return s.ConvAiSecretLocator
}

func (s *SpeechEngineConfigRequestHeadersValue) GetConvAiDynamicVariable() *ConvAiDynamicVariable {
	if s == nil {
		return nil
	}
	if s.ConvAiDynamicVariable == nil {
		return nil
	}
	return s.ConvAiDynamicVariable
}

func (s *SpeechEngineConfigRequestHeadersValue) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		s.typ = "String"
		s.String = valueString
		return nil
	}
	var valueConvAiSecretLocator *ConvAiSecretLocator
	if err := json.Unmarshal(data, &valueConvAiSecretLocator); err == nil {
		s.typ = "ConvAiSecretLocator"
		s.ConvAiSecretLocator = valueConvAiSecretLocator
		return nil
	}
	var valueConvAiDynamicVariable *ConvAiDynamicVariable
	if err := json.Unmarshal(data, &valueConvAiDynamicVariable); err == nil {
		s.typ = "ConvAiDynamicVariable"
		s.ConvAiDynamicVariable = valueConvAiDynamicVariable
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, s)
}

func (s SpeechEngineConfigRequestHeadersValue) MarshalJSON() ([]byte, error) {
	if s.typ == "String" || s.String != "" {
		return json.Marshal(s.String)
	}
	if s.typ == "ConvAiSecretLocator" || s.ConvAiSecretLocator != nil {
		return json.Marshal(s.ConvAiSecretLocator)
	}
	if s.typ == "ConvAiDynamicVariable" || s.ConvAiDynamicVariable != nil {
		return json.Marshal(s.ConvAiDynamicVariable)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", s)
}

type SpeechEngineConfigRequestHeadersValueVisitor interface {
	VisitString(string) error
	VisitConvAiSecretLocator(*ConvAiSecretLocator) error
	VisitConvAiDynamicVariable(*ConvAiDynamicVariable) error
}

func (s *SpeechEngineConfigRequestHeadersValue) Accept(visitor SpeechEngineConfigRequestHeadersValueVisitor) error {
	if s.typ == "String" || s.String != "" {
		return visitor.VisitString(s.String)
	}
	if s.typ == "ConvAiSecretLocator" || s.ConvAiSecretLocator != nil {
		return visitor.VisitConvAiSecretLocator(s.ConvAiSecretLocator)
	}
	if s.typ == "ConvAiDynamicVariable" || s.ConvAiDynamicVariable != nil {
		return visitor.VisitConvAiDynamicVariable(s.ConvAiDynamicVariable)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", s)
}
