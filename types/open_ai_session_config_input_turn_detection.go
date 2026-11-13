package types

import (
	"encoding/json"
	"fmt"
)

type OpenAiSessionConfigInputTurnDetection struct {
	OpenAiTurnDetection *OpenAiTurnDetection
	String              string
	typ                 string
}

func (o *OpenAiSessionConfigInputTurnDetection) GetOpenAiTurnDetection() *OpenAiTurnDetection {
	if o == nil {
		return nil
	}
	if o.OpenAiTurnDetection == nil {
		return nil
	}
	return o.OpenAiTurnDetection
}

func (o *OpenAiSessionConfigInputTurnDetection) GetString() string {
	if o == nil {
		return ""
	}
	return o.String
}

func (o *OpenAiSessionConfigInputTurnDetection) UnmarshalJSON(data []byte) error {
	var valueOpenAiTurnDetection *OpenAiTurnDetection
	if err := json.Unmarshal(data, &valueOpenAiTurnDetection); err == nil {
		o.typ = "OpenAiTurnDetection"
		o.OpenAiTurnDetection = valueOpenAiTurnDetection
		return nil
	}
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		o.typ = "String"
		o.String = valueString
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, o)
}

func (o OpenAiSessionConfigInputTurnDetection) MarshalJSON() ([]byte, error) {
	if o.typ == "OpenAiTurnDetection" || o.OpenAiTurnDetection != nil {
		return json.Marshal(o.OpenAiTurnDetection)
	}
	if o.typ == "String" || o.String != "" {
		return json.Marshal(o.String)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", o)
}

type OpenAiSessionConfigInputTurnDetectionVisitor interface {
	VisitOpenAiTurnDetection(*OpenAiTurnDetection) error
	VisitString(string) error
}

func (o *OpenAiSessionConfigInputTurnDetection) Accept(visitor OpenAiSessionConfigInputTurnDetectionVisitor) error {
	if o.typ == "OpenAiTurnDetection" || o.OpenAiTurnDetection != nil {
		return visitor.VisitOpenAiTurnDetection(o.OpenAiTurnDetection)
	}
	if o.typ == "String" || o.String != "" {
		return visitor.VisitString(o.String)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", o)
}
