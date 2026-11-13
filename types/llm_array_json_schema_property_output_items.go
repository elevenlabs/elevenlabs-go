package types

import (
	"encoding/json"
	"fmt"
)

type LlmArrayJsonSchemaPropertyOutputItems struct {
	LlmLiteralJsonSchemaProperty      *LlmLiteralJsonSchemaProperty
	LlmObjectJsonSchemaPropertyOutput *LlmObjectJsonSchemaPropertyOutput
	LlmArrayJsonSchemaPropertyOutput  *LlmArrayJsonSchemaPropertyOutput
	LlmAnyOfJsonSchemaPropertyOutput  LlmAnyOfJsonSchemaPropertyOutput
	typ                               string
}

func (l *LlmArrayJsonSchemaPropertyOutputItems) GetLlmLiteralJsonSchemaProperty() *LlmLiteralJsonSchemaProperty {
	if l == nil {
		return nil
	}
	if l.LlmLiteralJsonSchemaProperty == nil {
		return nil
	}
	return l.LlmLiteralJsonSchemaProperty
}

func (l *LlmArrayJsonSchemaPropertyOutputItems) GetLlmObjectJsonSchemaPropertyOutput() *LlmObjectJsonSchemaPropertyOutput {
	if l == nil {
		return nil
	}
	if l.LlmObjectJsonSchemaPropertyOutput == nil {
		return nil
	}
	return l.LlmObjectJsonSchemaPropertyOutput
}

func (l *LlmArrayJsonSchemaPropertyOutputItems) GetLlmArrayJsonSchemaPropertyOutput() *LlmArrayJsonSchemaPropertyOutput {
	if l == nil {
		return nil
	}
	if l.LlmArrayJsonSchemaPropertyOutput == nil {
		return nil
	}
	return l.LlmArrayJsonSchemaPropertyOutput
}

func (l *LlmArrayJsonSchemaPropertyOutputItems) GetLlmAnyOfJsonSchemaPropertyOutput() LlmAnyOfJsonSchemaPropertyOutput {
	if l == nil {
		return nil
	}
	return l.LlmAnyOfJsonSchemaPropertyOutput
}

func (l *LlmArrayJsonSchemaPropertyOutputItems) UnmarshalJSON(data []byte) error {
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
		var valueLlmObjectJsonSchemaPropertyOutputLiteral *LlmObjectJsonSchemaPropertyOutput
		if err := json.Unmarshal(data, &valueLlmObjectJsonSchemaPropertyOutputLiteral); err == nil {
			l.typ = "LlmObjectJsonSchemaPropertyOutput"
			l.LlmObjectJsonSchemaPropertyOutput = valueLlmObjectJsonSchemaPropertyOutputLiteral
			return nil
		}
	}
	if hasLiteral("type", "array") {
		matchedLiteral = true
		var valueLlmArrayJsonSchemaPropertyOutputLiteral *LlmArrayJsonSchemaPropertyOutput
		if err := json.Unmarshal(data, &valueLlmArrayJsonSchemaPropertyOutputLiteral); err == nil {
			l.typ = "LlmArrayJsonSchemaPropertyOutput"
			l.LlmArrayJsonSchemaPropertyOutput = valueLlmArrayJsonSchemaPropertyOutputLiteral
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
		var valueLlmObjectJsonSchemaPropertyOutput *LlmObjectJsonSchemaPropertyOutput
		if err := json.Unmarshal(data, &valueLlmObjectJsonSchemaPropertyOutput); err == nil {
			l.typ = "LlmObjectJsonSchemaPropertyOutput"
			l.LlmObjectJsonSchemaPropertyOutput = valueLlmObjectJsonSchemaPropertyOutput
			return nil
		}
	}
	var valueLlmArrayJsonSchemaPropertyOutput *LlmArrayJsonSchemaPropertyOutput
	if err := json.Unmarshal(data, &valueLlmArrayJsonSchemaPropertyOutput); err == nil {
		l.typ = "LlmArrayJsonSchemaPropertyOutput"
		l.LlmArrayJsonSchemaPropertyOutput = valueLlmArrayJsonSchemaPropertyOutput
		return nil
	}
	var valueLlmAnyOfJsonSchemaPropertyOutput LlmAnyOfJsonSchemaPropertyOutput
	if err := json.Unmarshal(data, &valueLlmAnyOfJsonSchemaPropertyOutput); err == nil {
		l.typ = "LlmAnyOfJsonSchemaPropertyOutput"
		l.LlmAnyOfJsonSchemaPropertyOutput = valueLlmAnyOfJsonSchemaPropertyOutput
		return nil
	}

	// Retried without their key guards, for a payload that matched none of them.
	var valueLlmObjectJsonSchemaPropertyOutputFallback *LlmObjectJsonSchemaPropertyOutput
	if err := json.Unmarshal(data, &valueLlmObjectJsonSchemaPropertyOutputFallback); err == nil {
		l.typ = "LlmObjectJsonSchemaPropertyOutput"
		l.LlmObjectJsonSchemaPropertyOutput = valueLlmObjectJsonSchemaPropertyOutputFallback
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, l)
}

func (l LlmArrayJsonSchemaPropertyOutputItems) MarshalJSON() ([]byte, error) {
	if l.typ == "LlmLiteralJsonSchemaProperty" || l.LlmLiteralJsonSchemaProperty != nil {
		return json.Marshal(l.LlmLiteralJsonSchemaProperty)
	}
	if l.typ == "LlmObjectJsonSchemaPropertyOutput" || l.LlmObjectJsonSchemaPropertyOutput != nil {
		return json.Marshal(l.LlmObjectJsonSchemaPropertyOutput)
	}
	if l.typ == "LlmArrayJsonSchemaPropertyOutput" || l.LlmArrayJsonSchemaPropertyOutput != nil {
		return json.Marshal(l.LlmArrayJsonSchemaPropertyOutput)
	}
	if l.typ == "LlmAnyOfJsonSchemaPropertyOutput" || l.LlmAnyOfJsonSchemaPropertyOutput != nil {
		return json.Marshal(l.LlmAnyOfJsonSchemaPropertyOutput)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", l)
}

type LlmArrayJsonSchemaPropertyOutputItemsVisitor interface {
	VisitLlmLiteralJsonSchemaProperty(*LlmLiteralJsonSchemaProperty) error
	VisitLlmObjectJsonSchemaPropertyOutput(*LlmObjectJsonSchemaPropertyOutput) error
	VisitLlmArrayJsonSchemaPropertyOutput(*LlmArrayJsonSchemaPropertyOutput) error
	VisitLlmAnyOfJsonSchemaPropertyOutput(LlmAnyOfJsonSchemaPropertyOutput) error
}

func (l *LlmArrayJsonSchemaPropertyOutputItems) Accept(visitor LlmArrayJsonSchemaPropertyOutputItemsVisitor) error {
	if l.typ == "LlmLiteralJsonSchemaProperty" || l.LlmLiteralJsonSchemaProperty != nil {
		return visitor.VisitLlmLiteralJsonSchemaProperty(l.LlmLiteralJsonSchemaProperty)
	}
	if l.typ == "LlmObjectJsonSchemaPropertyOutput" || l.LlmObjectJsonSchemaPropertyOutput != nil {
		return visitor.VisitLlmObjectJsonSchemaPropertyOutput(l.LlmObjectJsonSchemaPropertyOutput)
	}
	if l.typ == "LlmArrayJsonSchemaPropertyOutput" || l.LlmArrayJsonSchemaPropertyOutput != nil {
		return visitor.VisitLlmArrayJsonSchemaPropertyOutput(l.LlmArrayJsonSchemaPropertyOutput)
	}
	if l.typ == "LlmAnyOfJsonSchemaPropertyOutput" || l.LlmAnyOfJsonSchemaPropertyOutput != nil {
		return visitor.VisitLlmAnyOfJsonSchemaPropertyOutput(l.LlmAnyOfJsonSchemaPropertyOutput)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", l)
}
