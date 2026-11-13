package types

import (
	"encoding/json"
	"fmt"
)

type OpenAiSessionConfigInputToolChoice struct {
	String                   string
	OpenAiToolChoiceFunction *OpenAiToolChoiceFunction
	typ                      string
}

func (o *OpenAiSessionConfigInputToolChoice) GetString() string {
	if o == nil {
		return ""
	}
	return o.String
}

func (o *OpenAiSessionConfigInputToolChoice) GetOpenAiToolChoiceFunction() *OpenAiToolChoiceFunction {
	if o == nil {
		return nil
	}
	if o.OpenAiToolChoiceFunction == nil {
		return nil
	}
	return o.OpenAiToolChoiceFunction
}

func (o *OpenAiSessionConfigInputToolChoice) UnmarshalJSON(data []byte) error {
	var unionKeys map[string]json.RawMessage
	_ = json.Unmarshal(data, &unionKeys)
	hasLiteral := func(key string, want interface{}) bool {
		raw, ok := unionKeys[key]
		if !ok {
			return false
		}
		var got interface{}
		return json.Unmarshal(raw, &got) == nil && got == want
	}
	matchedLiteral := false
	if hasLiteral("type", "function") {
		matchedLiteral = true
		var valueOpenAiToolChoiceFunctionLiteral *OpenAiToolChoiceFunction
		if err := json.Unmarshal(data, &valueOpenAiToolChoiceFunctionLiteral); err == nil {
			o.typ = "OpenAiToolChoiceFunction"
			o.OpenAiToolChoiceFunction = valueOpenAiToolChoiceFunctionLiteral
			return nil
		}
	}
	if matchedLiteral {
		return fmt.Errorf("%s cannot be deserialized as a %T", data, o)
	}
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		o.typ = "String"
		o.String = valueString
		return nil
	}
	var valueOpenAiToolChoiceFunction *OpenAiToolChoiceFunction
	if err := json.Unmarshal(data, &valueOpenAiToolChoiceFunction); err == nil {
		o.typ = "OpenAiToolChoiceFunction"
		o.OpenAiToolChoiceFunction = valueOpenAiToolChoiceFunction
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, o)
}

func (o OpenAiSessionConfigInputToolChoice) MarshalJSON() ([]byte, error) {
	if o.typ == "String" || o.String != "" {
		return json.Marshal(o.String)
	}
	if o.typ == "OpenAiToolChoiceFunction" || o.OpenAiToolChoiceFunction != nil {
		return json.Marshal(o.OpenAiToolChoiceFunction)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", o)
}

type OpenAiSessionConfigInputToolChoiceVisitor interface {
	VisitString(string) error
	VisitOpenAiToolChoiceFunction(*OpenAiToolChoiceFunction) error
}

func (o *OpenAiSessionConfigInputToolChoice) Accept(visitor OpenAiSessionConfigInputToolChoiceVisitor) error {
	if o.typ == "String" || o.String != "" {
		return visitor.VisitString(o.String)
	}
	if o.typ == "OpenAiToolChoiceFunction" || o.OpenAiToolChoiceFunction != nil {
		return visitor.VisitOpenAiToolChoiceFunction(o.OpenAiToolChoiceFunction)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", o)
}
