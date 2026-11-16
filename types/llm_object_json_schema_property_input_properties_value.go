package types

import (
	"encoding/json"
	"fmt"
)

type LlmObjectJsonSchemaPropertyInputPropertiesValue struct {
	LlmLiteralJsonSchemaProperty     *LlmLiteralJsonSchemaProperty
	LlmObjectJsonSchemaPropertyInput *LlmObjectJsonSchemaPropertyInput
	LlmArrayJsonSchemaPropertyInput  *LlmArrayJsonSchemaPropertyInput
	LlmAnyOfJsonSchemaPropertyInput  *LlmAnyOfJsonSchemaPropertyInput
	typ                              string
}

func (l *LlmObjectJsonSchemaPropertyInputPropertiesValue) GetLlmLiteralJsonSchemaProperty() *LlmLiteralJsonSchemaProperty {
	if l == nil {
		return nil
	}
	if l.LlmLiteralJsonSchemaProperty == nil {
		return nil
	}
	return l.LlmLiteralJsonSchemaProperty
}

func (l *LlmObjectJsonSchemaPropertyInputPropertiesValue) GetLlmObjectJsonSchemaPropertyInput() *LlmObjectJsonSchemaPropertyInput {
	if l == nil {
		return nil
	}
	if l.LlmObjectJsonSchemaPropertyInput == nil {
		return nil
	}
	return l.LlmObjectJsonSchemaPropertyInput
}

func (l *LlmObjectJsonSchemaPropertyInputPropertiesValue) GetLlmArrayJsonSchemaPropertyInput() *LlmArrayJsonSchemaPropertyInput {
	if l == nil {
		return nil
	}
	if l.LlmArrayJsonSchemaPropertyInput == nil {
		return nil
	}
	return l.LlmArrayJsonSchemaPropertyInput
}

func (l *LlmObjectJsonSchemaPropertyInputPropertiesValue) GetLlmAnyOfJsonSchemaPropertyInput() *LlmAnyOfJsonSchemaPropertyInput {
	if l == nil {
		return nil
	}
	if l.LlmAnyOfJsonSchemaPropertyInput == nil {
		return nil
	}
	return l.LlmAnyOfJsonSchemaPropertyInput
}

func (l *LlmObjectJsonSchemaPropertyInputPropertiesValue) UnmarshalJSON(data []byte) error {
	var unionKeys map[string]json.RawMessage
	_ = json.Unmarshal(data, &unionKeys)
	hasAnyKey := func(keys ...string) bool {
		for _, key := range keys {
			if _, ok := unionKeys[key]; ok {
				return true
			}
		}
		return false
	}
	hasLiteral := func(key string, want interface{}) bool {
		raw, ok := unionKeys[key]
		if !ok {
			return false
		}
		var got interface{}
		return json.Unmarshal(raw, &got) == nil && got == want
	}
	matchedLiteral := false
	if hasLiteral("type", "object") {
		matchedLiteral = true
		var valueLlmObjectJsonSchemaPropertyInputLiteral *LlmObjectJsonSchemaPropertyInput
		if err := json.Unmarshal(data, &valueLlmObjectJsonSchemaPropertyInputLiteral); err == nil {
			l.typ = "LlmObjectJsonSchemaPropertyInput"
			l.LlmObjectJsonSchemaPropertyInput = valueLlmObjectJsonSchemaPropertyInputLiteral
			return nil
		}
	}
	if hasLiteral("type", "array") {
		matchedLiteral = true
		var valueLlmArrayJsonSchemaPropertyInputLiteral *LlmArrayJsonSchemaPropertyInput
		if err := json.Unmarshal(data, &valueLlmArrayJsonSchemaPropertyInputLiteral); err == nil {
			l.typ = "LlmArrayJsonSchemaPropertyInput"
			l.LlmArrayJsonSchemaPropertyInput = valueLlmArrayJsonSchemaPropertyInputLiteral
			return nil
		}
	}
	if matchedLiteral {
		return fmt.Errorf("%s cannot be deserialized as a %T", data, l)
	}
	var valueLlmLiteralJsonSchemaProperty *LlmLiteralJsonSchemaProperty
	if err := json.Unmarshal(data, &valueLlmLiteralJsonSchemaProperty); err == nil {
		l.typ = "LlmLiteralJsonSchemaProperty"
		l.LlmLiteralJsonSchemaProperty = valueLlmLiteralJsonSchemaProperty
		return nil
	}
	if hasAnyKey("type", "required", "description", "properties") {
		var valueLlmObjectJsonSchemaPropertyInput *LlmObjectJsonSchemaPropertyInput
		if err := json.Unmarshal(data, &valueLlmObjectJsonSchemaPropertyInput); err == nil {
			l.typ = "LlmObjectJsonSchemaPropertyInput"
			l.LlmObjectJsonSchemaPropertyInput = valueLlmObjectJsonSchemaPropertyInput
			return nil
		}
	}
	var valueLlmArrayJsonSchemaPropertyInput *LlmArrayJsonSchemaPropertyInput
	if err := json.Unmarshal(data, &valueLlmArrayJsonSchemaPropertyInput); err == nil {
		l.typ = "LlmArrayJsonSchemaPropertyInput"
		l.LlmArrayJsonSchemaPropertyInput = valueLlmArrayJsonSchemaPropertyInput
		return nil
	}
	var valueLlmAnyOfJsonSchemaPropertyInput *LlmAnyOfJsonSchemaPropertyInput
	if err := json.Unmarshal(data, &valueLlmAnyOfJsonSchemaPropertyInput); err == nil {
		l.typ = "LlmAnyOfJsonSchemaPropertyInput"
		l.LlmAnyOfJsonSchemaPropertyInput = valueLlmAnyOfJsonSchemaPropertyInput
		return nil
	}

	// Retried without their key guards, for a payload that matched none of them.
	var valueLlmObjectJsonSchemaPropertyInputFallback *LlmObjectJsonSchemaPropertyInput
	if err := json.Unmarshal(data, &valueLlmObjectJsonSchemaPropertyInputFallback); err == nil {
		l.typ = "LlmObjectJsonSchemaPropertyInput"
		l.LlmObjectJsonSchemaPropertyInput = valueLlmObjectJsonSchemaPropertyInputFallback
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, l)
}

func (l LlmObjectJsonSchemaPropertyInputPropertiesValue) MarshalJSON() ([]byte, error) {
	if l.typ == "LlmLiteralJsonSchemaProperty" || l.LlmLiteralJsonSchemaProperty != nil {
		return json.Marshal(l.LlmLiteralJsonSchemaProperty)
	}
	if l.typ == "LlmObjectJsonSchemaPropertyInput" || l.LlmObjectJsonSchemaPropertyInput != nil {
		return json.Marshal(l.LlmObjectJsonSchemaPropertyInput)
	}
	if l.typ == "LlmArrayJsonSchemaPropertyInput" || l.LlmArrayJsonSchemaPropertyInput != nil {
		return json.Marshal(l.LlmArrayJsonSchemaPropertyInput)
	}
	if l.typ == "LlmAnyOfJsonSchemaPropertyInput" || l.LlmAnyOfJsonSchemaPropertyInput != nil {
		return json.Marshal(l.LlmAnyOfJsonSchemaPropertyInput)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", l)
}

type LlmObjectJsonSchemaPropertyInputPropertiesValueVisitor interface {
	VisitLlmLiteralJsonSchemaProperty(*LlmLiteralJsonSchemaProperty) error
	VisitLlmObjectJsonSchemaPropertyInput(*LlmObjectJsonSchemaPropertyInput) error
	VisitLlmArrayJsonSchemaPropertyInput(*LlmArrayJsonSchemaPropertyInput) error
	VisitLlmAnyOfJsonSchemaPropertyInput(*LlmAnyOfJsonSchemaPropertyInput) error
}

func (l *LlmObjectJsonSchemaPropertyInputPropertiesValue) Accept(visitor LlmObjectJsonSchemaPropertyInputPropertiesValueVisitor) error {
	if l.typ == "LlmLiteralJsonSchemaProperty" || l.LlmLiteralJsonSchemaProperty != nil {
		return visitor.VisitLlmLiteralJsonSchemaProperty(l.LlmLiteralJsonSchemaProperty)
	}
	if l.typ == "LlmObjectJsonSchemaPropertyInput" || l.LlmObjectJsonSchemaPropertyInput != nil {
		return visitor.VisitLlmObjectJsonSchemaPropertyInput(l.LlmObjectJsonSchemaPropertyInput)
	}
	if l.typ == "LlmArrayJsonSchemaPropertyInput" || l.LlmArrayJsonSchemaPropertyInput != nil {
		return visitor.VisitLlmArrayJsonSchemaPropertyInput(l.LlmArrayJsonSchemaPropertyInput)
	}
	if l.typ == "LlmAnyOfJsonSchemaPropertyInput" || l.LlmAnyOfJsonSchemaPropertyInput != nil {
		return visitor.VisitLlmAnyOfJsonSchemaPropertyInput(l.LlmAnyOfJsonSchemaPropertyInput)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", l)
}
