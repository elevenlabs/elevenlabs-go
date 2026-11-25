package types

import (
	"encoding/json"
	"fmt"
)

// Schema for array elements.
type ArrayJsonSchemaPropertyInputItems struct {
	LiteralJsonSchemaProperty     *LiteralJsonSchemaProperty
	ObjectJsonSchemaPropertyInput *ObjectJsonSchemaPropertyInput
	ArrayJsonSchemaPropertyInput  *ArrayJsonSchemaPropertyInput
	AnyOfJsonSchemaPropertyInput  *AnyOfJsonSchemaPropertyInput
	typ                           string
}

func (a *ArrayJsonSchemaPropertyInputItems) GetLiteralJsonSchemaProperty() *LiteralJsonSchemaProperty {
	if a == nil {
		return nil
	}
	if a.LiteralJsonSchemaProperty == nil {
		return nil
	}
	return a.LiteralJsonSchemaProperty
}

func (a *ArrayJsonSchemaPropertyInputItems) GetObjectJsonSchemaPropertyInput() *ObjectJsonSchemaPropertyInput {
	if a == nil {
		return nil
	}
	if a.ObjectJsonSchemaPropertyInput == nil {
		return nil
	}
	return a.ObjectJsonSchemaPropertyInput
}

func (a *ArrayJsonSchemaPropertyInputItems) GetArrayJsonSchemaPropertyInput() *ArrayJsonSchemaPropertyInput {
	if a == nil {
		return nil
	}
	if a.ArrayJsonSchemaPropertyInput == nil {
		return nil
	}
	return a.ArrayJsonSchemaPropertyInput
}

func (a *ArrayJsonSchemaPropertyInputItems) GetAnyOfJsonSchemaPropertyInput() *AnyOfJsonSchemaPropertyInput {
	if a == nil {
		return nil
	}
	if a.AnyOfJsonSchemaPropertyInput == nil {
		return nil
	}
	return a.AnyOfJsonSchemaPropertyInput
}

func (a *ArrayJsonSchemaPropertyInputItems) UnmarshalJSON(data []byte) error {
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
		var valueObjectJsonSchemaPropertyInputLiteral *ObjectJsonSchemaPropertyInput
		if err := json.Unmarshal(data, &valueObjectJsonSchemaPropertyInputLiteral); err == nil {
			a.typ = "ObjectJsonSchemaPropertyInput"
			a.ObjectJsonSchemaPropertyInput = valueObjectJsonSchemaPropertyInputLiteral
			return nil
		}
	}
	if hasLiteral("type", "array") {
		matchedLiteral = true
		var valueArrayJsonSchemaPropertyInputLiteral *ArrayJsonSchemaPropertyInput
		if err := json.Unmarshal(data, &valueArrayJsonSchemaPropertyInputLiteral); err == nil {
			a.typ = "ArrayJsonSchemaPropertyInput"
			a.ArrayJsonSchemaPropertyInput = valueArrayJsonSchemaPropertyInputLiteral
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
	if hasAnyKey("property_kind", "description", "dynamic_variable", "constant_value", "is_omitted", "type", "required", "properties", "required_constraints") {
		var valueObjectJsonSchemaPropertyInput *ObjectJsonSchemaPropertyInput
		if err := json.Unmarshal(data, &valueObjectJsonSchemaPropertyInput); err == nil {
			a.typ = "ObjectJsonSchemaPropertyInput"
			a.ObjectJsonSchemaPropertyInput = valueObjectJsonSchemaPropertyInput
			return nil
		}
	}
	if hasAnyKey("property_kind", "description", "dynamic_variable", "constant_value", "is_omitted", "type", "items") {
		var valueArrayJsonSchemaPropertyInput *ArrayJsonSchemaPropertyInput
		if err := json.Unmarshal(data, &valueArrayJsonSchemaPropertyInput); err == nil {
			a.typ = "ArrayJsonSchemaPropertyInput"
			a.ArrayJsonSchemaPropertyInput = valueArrayJsonSchemaPropertyInput
			return nil
		}
	}
	var valueAnyOfJsonSchemaPropertyInput *AnyOfJsonSchemaPropertyInput
	if err := json.Unmarshal(data, &valueAnyOfJsonSchemaPropertyInput); err == nil {
		a.typ = "AnyOfJsonSchemaPropertyInput"
		a.AnyOfJsonSchemaPropertyInput = valueAnyOfJsonSchemaPropertyInput
		return nil
	}

	// Retried without their key guards, for a payload that matched none of them.
	var valueObjectJsonSchemaPropertyInputFallback *ObjectJsonSchemaPropertyInput
	if err := json.Unmarshal(data, &valueObjectJsonSchemaPropertyInputFallback); err == nil {
		a.typ = "ObjectJsonSchemaPropertyInput"
		a.ObjectJsonSchemaPropertyInput = valueObjectJsonSchemaPropertyInputFallback
		return nil
	}
	var valueArrayJsonSchemaPropertyInputFallback *ArrayJsonSchemaPropertyInput
	if err := json.Unmarshal(data, &valueArrayJsonSchemaPropertyInputFallback); err == nil {
		a.typ = "ArrayJsonSchemaPropertyInput"
		a.ArrayJsonSchemaPropertyInput = valueArrayJsonSchemaPropertyInputFallback
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, a)
}

func (a ArrayJsonSchemaPropertyInputItems) MarshalJSON() ([]byte, error) {
	if a.typ == "LiteralJsonSchemaProperty" || a.LiteralJsonSchemaProperty != nil {
		return json.Marshal(a.LiteralJsonSchemaProperty)
	}
	if a.typ == "ObjectJsonSchemaPropertyInput" || a.ObjectJsonSchemaPropertyInput != nil {
		return json.Marshal(a.ObjectJsonSchemaPropertyInput)
	}
	if a.typ == "ArrayJsonSchemaPropertyInput" || a.ArrayJsonSchemaPropertyInput != nil {
		return json.Marshal(a.ArrayJsonSchemaPropertyInput)
	}
	if a.typ == "AnyOfJsonSchemaPropertyInput" || a.AnyOfJsonSchemaPropertyInput != nil {
		return json.Marshal(a.AnyOfJsonSchemaPropertyInput)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", a)
}

type ArrayJsonSchemaPropertyInputItemsVisitor interface {
	VisitLiteralJsonSchemaProperty(*LiteralJsonSchemaProperty) error
	VisitObjectJsonSchemaPropertyInput(*ObjectJsonSchemaPropertyInput) error
	VisitArrayJsonSchemaPropertyInput(*ArrayJsonSchemaPropertyInput) error
	VisitAnyOfJsonSchemaPropertyInput(*AnyOfJsonSchemaPropertyInput) error
}

func (a *ArrayJsonSchemaPropertyInputItems) Accept(visitor ArrayJsonSchemaPropertyInputItemsVisitor) error {
	if a.typ == "LiteralJsonSchemaProperty" || a.LiteralJsonSchemaProperty != nil {
		return visitor.VisitLiteralJsonSchemaProperty(a.LiteralJsonSchemaProperty)
	}
	if a.typ == "ObjectJsonSchemaPropertyInput" || a.ObjectJsonSchemaPropertyInput != nil {
		return visitor.VisitObjectJsonSchemaPropertyInput(a.ObjectJsonSchemaPropertyInput)
	}
	if a.typ == "ArrayJsonSchemaPropertyInput" || a.ArrayJsonSchemaPropertyInput != nil {
		return visitor.VisitArrayJsonSchemaPropertyInput(a.ArrayJsonSchemaPropertyInput)
	}
	if a.typ == "AnyOfJsonSchemaPropertyInput" || a.AnyOfJsonSchemaPropertyInput != nil {
		return visitor.VisitAnyOfJsonSchemaPropertyInput(a.AnyOfJsonSchemaPropertyInput)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", a)
}
