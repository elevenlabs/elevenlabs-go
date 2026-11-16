package types

import (
	"encoding/json"
	"fmt"
)

// Schema for array elements.
type ArrayJsonSchemaPropertyOutputItems struct {
	LiteralJsonSchemaProperty      *LiteralJsonSchemaProperty
	ObjectJsonSchemaPropertyOutput *ObjectJsonSchemaPropertyOutput
	ArrayJsonSchemaPropertyOutput  *ArrayJsonSchemaPropertyOutput
	typ                            string
}

func (a *ArrayJsonSchemaPropertyOutputItems) GetLiteralJsonSchemaProperty() *LiteralJsonSchemaProperty {
	if a == nil {
		return nil
	}
	if a.LiteralJsonSchemaProperty == nil {
		return nil
	}
	return a.LiteralJsonSchemaProperty
}

func (a *ArrayJsonSchemaPropertyOutputItems) GetObjectJsonSchemaPropertyOutput() *ObjectJsonSchemaPropertyOutput {
	if a == nil {
		return nil
	}
	if a.ObjectJsonSchemaPropertyOutput == nil {
		return nil
	}
	return a.ObjectJsonSchemaPropertyOutput
}

func (a *ArrayJsonSchemaPropertyOutputItems) GetArrayJsonSchemaPropertyOutput() *ArrayJsonSchemaPropertyOutput {
	if a == nil {
		return nil
	}
	if a.ArrayJsonSchemaPropertyOutput == nil {
		return nil
	}
	return a.ArrayJsonSchemaPropertyOutput
}

func (a *ArrayJsonSchemaPropertyOutputItems) UnmarshalJSON(data []byte) error {
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
		var valueObjectJsonSchemaPropertyOutputLiteral *ObjectJsonSchemaPropertyOutput
		if err := json.Unmarshal(data, &valueObjectJsonSchemaPropertyOutputLiteral); err == nil {
			a.typ = "ObjectJsonSchemaPropertyOutput"
			a.ObjectJsonSchemaPropertyOutput = valueObjectJsonSchemaPropertyOutputLiteral
			return nil
		}
	}
	if hasLiteral("type", "array") {
		matchedLiteral = true
		var valueArrayJsonSchemaPropertyOutputLiteral *ArrayJsonSchemaPropertyOutput
		if err := json.Unmarshal(data, &valueArrayJsonSchemaPropertyOutputLiteral); err == nil {
			a.typ = "ArrayJsonSchemaPropertyOutput"
			a.ArrayJsonSchemaPropertyOutput = valueArrayJsonSchemaPropertyOutputLiteral
			return nil
		}
	}
	if matchedLiteral {
		return fmt.Errorf("%s cannot be deserialized as a %T", data, a)
	}
	var valueLiteralJsonSchemaProperty *LiteralJsonSchemaProperty
	if err := json.Unmarshal(data, &valueLiteralJsonSchemaProperty); err == nil {
		a.typ = "LiteralJsonSchemaProperty"
		a.LiteralJsonSchemaProperty = valueLiteralJsonSchemaProperty
		return nil
	}
	if hasAnyKey("description", "dynamic_variable", "constant_value", "is_omitted", "type", "required", "properties", "required_constraints") {
		var valueObjectJsonSchemaPropertyOutput *ObjectJsonSchemaPropertyOutput
		if err := json.Unmarshal(data, &valueObjectJsonSchemaPropertyOutput); err == nil {
			a.typ = "ObjectJsonSchemaPropertyOutput"
			a.ObjectJsonSchemaPropertyOutput = valueObjectJsonSchemaPropertyOutput
			return nil
		}
	}
	if hasAnyKey("description", "dynamic_variable", "constant_value", "is_omitted", "type", "items") {
		var valueArrayJsonSchemaPropertyOutput *ArrayJsonSchemaPropertyOutput
		if err := json.Unmarshal(data, &valueArrayJsonSchemaPropertyOutput); err == nil {
			a.typ = "ArrayJsonSchemaPropertyOutput"
			a.ArrayJsonSchemaPropertyOutput = valueArrayJsonSchemaPropertyOutput
			return nil
		}
	}

	// Retried without their key guards, for a payload that matched none of them.
	var valueObjectJsonSchemaPropertyOutputFallback *ObjectJsonSchemaPropertyOutput
	if err := json.Unmarshal(data, &valueObjectJsonSchemaPropertyOutputFallback); err == nil {
		a.typ = "ObjectJsonSchemaPropertyOutput"
		a.ObjectJsonSchemaPropertyOutput = valueObjectJsonSchemaPropertyOutputFallback
		return nil
	}
	var valueArrayJsonSchemaPropertyOutputFallback *ArrayJsonSchemaPropertyOutput
	if err := json.Unmarshal(data, &valueArrayJsonSchemaPropertyOutputFallback); err == nil {
		a.typ = "ArrayJsonSchemaPropertyOutput"
		a.ArrayJsonSchemaPropertyOutput = valueArrayJsonSchemaPropertyOutputFallback
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, a)
}

func (a ArrayJsonSchemaPropertyOutputItems) MarshalJSON() ([]byte, error) {
	if a.typ == "LiteralJsonSchemaProperty" || a.LiteralJsonSchemaProperty != nil {
		return json.Marshal(a.LiteralJsonSchemaProperty)
	}
	if a.typ == "ObjectJsonSchemaPropertyOutput" || a.ObjectJsonSchemaPropertyOutput != nil {
		return json.Marshal(a.ObjectJsonSchemaPropertyOutput)
	}
	if a.typ == "ArrayJsonSchemaPropertyOutput" || a.ArrayJsonSchemaPropertyOutput != nil {
		return json.Marshal(a.ArrayJsonSchemaPropertyOutput)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", a)
}

type ArrayJsonSchemaPropertyOutputItemsVisitor interface {
	VisitLiteralJsonSchemaProperty(*LiteralJsonSchemaProperty) error
	VisitObjectJsonSchemaPropertyOutput(*ObjectJsonSchemaPropertyOutput) error
	VisitArrayJsonSchemaPropertyOutput(*ArrayJsonSchemaPropertyOutput) error
}

func (a *ArrayJsonSchemaPropertyOutputItems) Accept(visitor ArrayJsonSchemaPropertyOutputItemsVisitor) error {
	if a.typ == "LiteralJsonSchemaProperty" || a.LiteralJsonSchemaProperty != nil {
		return visitor.VisitLiteralJsonSchemaProperty(a.LiteralJsonSchemaProperty)
	}
	if a.typ == "ObjectJsonSchemaPropertyOutput" || a.ObjectJsonSchemaPropertyOutput != nil {
		return visitor.VisitObjectJsonSchemaPropertyOutput(a.ObjectJsonSchemaPropertyOutput)
	}
	if a.typ == "ArrayJsonSchemaPropertyOutput" || a.ArrayJsonSchemaPropertyOutput != nil {
		return visitor.VisitArrayJsonSchemaPropertyOutput(a.ArrayJsonSchemaPropertyOutput)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", a)
}
