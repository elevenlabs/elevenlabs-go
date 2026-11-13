package types

import (
	"encoding/json"
	"fmt"
)

type OpenAiSessionConfigOutputToolChoice struct {
	String                   string
	OpenAiToolChoiceFunction *OpenAiToolChoiceFunction
	typ                      string
}

func (openAiSessionConfigOutputToolChoice *OpenAiSessionConfigOutputToolChoice) GetString() string {
	if openAiSessionConfigOutputToolChoice == nil {
		return ""
	}
	return openAiSessionConfigOutputToolChoice.String
}

func (openAiSessionConfigOutputToolChoice *OpenAiSessionConfigOutputToolChoice) GetOpenAiToolChoiceFunction() *OpenAiToolChoiceFunction {
	if openAiSessionConfigOutputToolChoice == nil {
		return nil
	}
	if openAiSessionConfigOutputToolChoice.OpenAiToolChoiceFunction == nil {
		return nil
	}
	return openAiSessionConfigOutputToolChoice.OpenAiToolChoiceFunction
}

func (openAiSessionConfigOutputToolChoice *OpenAiSessionConfigOutputToolChoice) UnmarshalJSON(data []byte) error {
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
			openAiSessionConfigOutputToolChoice.typ = "OpenAiToolChoiceFunction"
			openAiSessionConfigOutputToolChoice.OpenAiToolChoiceFunction = valueOpenAiToolChoiceFunctionLiteral
			return nil
		}
	}
	if matchedLiteral {
		return fmt.Errorf("%s cannot be deserialized as a %T", data, openAiSessionConfigOutputToolChoice)
	}
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		openAiSessionConfigOutputToolChoice.typ = "String"
		openAiSessionConfigOutputToolChoice.String = valueString
		return nil
	}
	var valueOpenAiToolChoiceFunction *OpenAiToolChoiceFunction
	if err := json.Unmarshal(data, &valueOpenAiToolChoiceFunction); err == nil {
		openAiSessionConfigOutputToolChoice.typ = "OpenAiToolChoiceFunction"
		openAiSessionConfigOutputToolChoice.OpenAiToolChoiceFunction = valueOpenAiToolChoiceFunction
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, openAiSessionConfigOutputToolChoice)
}

func (openAiSessionConfigOutputToolChoice OpenAiSessionConfigOutputToolChoice) MarshalJSON() ([]byte, error) {
	if openAiSessionConfigOutputToolChoice.typ == "String" || openAiSessionConfigOutputToolChoice.String != "" {
		return json.Marshal(openAiSessionConfigOutputToolChoice.String)
	}
	if openAiSessionConfigOutputToolChoice.typ == "OpenAiToolChoiceFunction" || openAiSessionConfigOutputToolChoice.OpenAiToolChoiceFunction != nil {
		return json.Marshal(openAiSessionConfigOutputToolChoice.OpenAiToolChoiceFunction)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", openAiSessionConfigOutputToolChoice)
}

type OpenAiSessionConfigOutputToolChoiceVisitor interface {
	VisitString(string) error
	VisitOpenAiToolChoiceFunction(*OpenAiToolChoiceFunction) error
}

func (openAiSessionConfigOutputToolChoice *OpenAiSessionConfigOutputToolChoice) Accept(visitor OpenAiSessionConfigOutputToolChoiceVisitor) error {
	if openAiSessionConfigOutputToolChoice.typ == "String" || openAiSessionConfigOutputToolChoice.String != "" {
		return visitor.VisitString(openAiSessionConfigOutputToolChoice.String)
	}
	if openAiSessionConfigOutputToolChoice.typ == "OpenAiToolChoiceFunction" || openAiSessionConfigOutputToolChoice.OpenAiToolChoiceFunction != nil {
		return visitor.VisitOpenAiToolChoiceFunction(openAiSessionConfigOutputToolChoice.OpenAiToolChoiceFunction)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", openAiSessionConfigOutputToolChoice)
}
