package types

import (
	"encoding/json"
	"fmt"
)

type OpenAiSessionConfigInputTruncation struct {
	String           string
	StringUnknownMap map[string]any
	typ              string
}

func (o *OpenAiSessionConfigInputTruncation) GetString() string {
	if o == nil {
		return ""
	}
	return o.String
}

func (o *OpenAiSessionConfigInputTruncation) GetStringUnknownMap() map[string]any {
	if o == nil {
		return nil
	}
	return o.StringUnknownMap
}

func (o *OpenAiSessionConfigInputTruncation) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		o.typ = "String"
		o.String = valueString
		return nil
	}
	var valueStringUnknownMap map[string]any
	if err := json.Unmarshal(data, &valueStringUnknownMap); err == nil {
		o.typ = "StringUnknownMap"
		o.StringUnknownMap = valueStringUnknownMap
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, o)
}

func (o OpenAiSessionConfigInputTruncation) MarshalJSON() ([]byte, error) {
	if o.typ == "String" || o.String != "" {
		return json.Marshal(o.String)
	}
	if o.typ == "StringUnknownMap" || o.StringUnknownMap != nil {
		return json.Marshal(o.StringUnknownMap)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", o)
}

type OpenAiSessionConfigInputTruncationVisitor interface {
	VisitString(string) error
	VisitStringUnknownMap(map[string]any) error
}

func (o *OpenAiSessionConfigInputTruncation) Accept(visitor OpenAiSessionConfigInputTruncationVisitor) error {
	if o.typ == "String" || o.String != "" {
		return visitor.VisitString(o.String)
	}
	if o.typ == "StringUnknownMap" || o.StringUnknownMap != nil {
		return visitor.VisitStringUnknownMap(o.StringUnknownMap)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", o)
}
