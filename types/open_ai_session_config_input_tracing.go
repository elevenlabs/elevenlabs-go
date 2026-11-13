package types

import (
	"encoding/json"
	"fmt"
)

type OpenAiSessionConfigInputTracing struct {
	AutoStringLiteral          string
	OpenAiTracingConfiguration *OpenAiTracingConfiguration
	typ                        string
}

func (o *OpenAiSessionConfigInputTracing) GetAutoStringLiteral() string {
	if o == nil {
		return ""
	}
	return o.AutoStringLiteral
}

func (o *OpenAiSessionConfigInputTracing) GetOpenAiTracingConfiguration() *OpenAiTracingConfiguration {
	if o == nil {
		return nil
	}
	if o.OpenAiTracingConfiguration == nil {
		return nil
	}
	return o.OpenAiTracingConfiguration
}

func (o *OpenAiSessionConfigInputTracing) UnmarshalJSON(data []byte) error {
	var valueAutoStringLiteral string
	if err := json.Unmarshal(data, &valueAutoStringLiteral); err == nil {
		o.typ = "AutoStringLiteral"
		o.AutoStringLiteral = valueAutoStringLiteral
		if o.AutoStringLiteral != "auto" {
			return fmt.Errorf("unexpected value for literal on type %T; expected auto got %v", o, o.AutoStringLiteral)
		}
		return nil
	}
	var valueOpenAiTracingConfiguration *OpenAiTracingConfiguration
	if err := json.Unmarshal(data, &valueOpenAiTracingConfiguration); err == nil {
		o.typ = "OpenAiTracingConfiguration"
		o.OpenAiTracingConfiguration = valueOpenAiTracingConfiguration
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, o)
}

func (o OpenAiSessionConfigInputTracing) MarshalJSON() ([]byte, error) {
	if o.typ == "AutoStringLiteral" || o.AutoStringLiteral != "" {
		return json.Marshal(o.AutoStringLiteral)
	}
	if o.typ == "OpenAiTracingConfiguration" || o.OpenAiTracingConfiguration != nil {
		return json.Marshal(o.OpenAiTracingConfiguration)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", o)
}

type OpenAiSessionConfigInputTracingVisitor interface {
	VisitAutoStringLiteral(string) error
	VisitOpenAiTracingConfiguration(*OpenAiTracingConfiguration) error
}

func (o *OpenAiSessionConfigInputTracing) Accept(visitor OpenAiSessionConfigInputTracingVisitor) error {
	if o.typ == "AutoStringLiteral" || o.AutoStringLiteral != "" {
		return visitor.VisitAutoStringLiteral(o.AutoStringLiteral)
	}
	if o.typ == "OpenAiTracingConfiguration" || o.OpenAiTracingConfiguration != nil {
		return visitor.VisitOpenAiTracingConfiguration(o.OpenAiTracingConfiguration)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", o)
}

func NewOpenAiSessionConfigInputTracingWithAutoStringLiteral() *OpenAiSessionConfigInputTracing {
	return &OpenAiSessionConfigInputTracing{typ: "AutoStringLiteral", AutoStringLiteral: "auto"}
}
