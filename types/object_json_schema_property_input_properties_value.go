package types

import (
	"encoding/json"
	"fmt"
)

type ObjectJsonSchemaPropertyInputPropertiesValue struct {
	LiteralJsonSchemaProperty     *LiteralJsonSchemaProperty
	ObjectJsonSchemaPropertyInput *ObjectJsonSchemaPropertyInput
	ArrayJsonSchemaPropertyInput  *ArrayJsonSchemaPropertyInput
	AnyOfJsonSchemaPropertyInput  *AnyOfJsonSchemaPropertyInput
	typ                           string
}

func (o *ObjectJsonSchemaPropertyInputPropertiesValue) GetLiteralJsonSchemaProperty() *LiteralJsonSchemaProperty {
	if o == nil {
		return nil
	}
	if o.LiteralJsonSchemaProperty == nil {
		return nil
	}
	return o.LiteralJsonSchemaProperty
}

func (o *ObjectJsonSchemaPropertyInputPropertiesValue) GetObjectJsonSchemaPropertyInput() *ObjectJsonSchemaPropertyInput {
	if o == nil {
		return nil
	}
	if o.ObjectJsonSchemaPropertyInput == nil {
		return nil
	}
	return o.ObjectJsonSchemaPropertyInput
}

func (o *ObjectJsonSchemaPropertyInputPropertiesValue) GetArrayJsonSchemaPropertyInput() *ArrayJsonSchemaPropertyInput {
	if o == nil {
		return nil
	}
	if o.ArrayJsonSchemaPropertyInput == nil {
		return nil
	}
	return o.ArrayJsonSchemaPropertyInput
}

func (o *ObjectJsonSchemaPropertyInputPropertiesValue) GetAnyOfJsonSchemaPropertyInput() *AnyOfJsonSchemaPropertyInput {
	if o == nil {
		return nil
	}
	if o.AnyOfJsonSchemaPropertyInput == nil {
		return nil
	}
	return o.AnyOfJsonSchemaPropertyInput
}

func (o *ObjectJsonSchemaPropertyInputPropertiesValue) UnmarshalJSON(data []byte) error {
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
			o.typ = "ObjectJsonSchemaPropertyInput"
			o.ObjectJsonSchemaPropertyInput = valueObjectJsonSchemaPropertyInputLiteral
			return nil
		}
	}
	if hasLiteral("type", "array") {
		matchedLiteral = true
		var valueArrayJsonSchemaPropertyInputLiteral *ArrayJsonSchemaPropertyInput
		if err := json.Unmarshal(data, &valueArrayJsonSchemaPropertyInputLiteral); err == nil {
			o.typ = "ArrayJsonSchemaPropertyInput"
			o.ArrayJsonSchemaPropertyInput = valueArrayJsonSchemaPropertyInputLiteral
			return nil
		}
	}
	if matchedLiteral {
		return fmt.Errorf("%s cannot be deserialized as a %T", data, o)
	}
	var valueLiteralJsonSchemaProperty *LiteralJsonSchemaProperty
	if err := json.Unmarshal(data, &valueLiteralJsonSchemaProperty); err == nil {
		o.typ = "LiteralJsonSchemaProperty"
		o.LiteralJsonSchemaProperty = valueLiteralJsonSchemaProperty
		return nil
	}
	if hasAnyKey("property_kind", "description", "dynamic_variable", "constant_value", "is_omitted", "type", "required", "properties", "required_constraints") {
		var valueObjectJsonSchemaPropertyInput *ObjectJsonSchemaPropertyInput
		if err := json.Unmarshal(data, &valueObjectJsonSchemaPropertyInput); err == nil {
			o.typ = "ObjectJsonSchemaPropertyInput"
			o.ObjectJsonSchemaPropertyInput = valueObjectJsonSchemaPropertyInput
			return nil
		}
	}
	if hasAnyKey("property_kind", "description", "dynamic_variable", "constant_value", "is_omitted", "type", "items") {
		var valueArrayJsonSchemaPropertyInput *ArrayJsonSchemaPropertyInput
		if err := json.Unmarshal(data, &valueArrayJsonSchemaPropertyInput); err == nil {
			o.typ = "ArrayJsonSchemaPropertyInput"
			o.ArrayJsonSchemaPropertyInput = valueArrayJsonSchemaPropertyInput
			return nil
		}
	}
	var valueAnyOfJsonSchemaPropertyInput *AnyOfJsonSchemaPropertyInput
	if err := json.Unmarshal(data, &valueAnyOfJsonSchemaPropertyInput); err == nil {
		o.typ = "AnyOfJsonSchemaPropertyInput"
		o.AnyOfJsonSchemaPropertyInput = valueAnyOfJsonSchemaPropertyInput
		return nil
	}

	// Retried without their key guards, for a payload that matched none of them.
	var valueObjectJsonSchemaPropertyInputFallback *ObjectJsonSchemaPropertyInput
	if err := json.Unmarshal(data, &valueObjectJsonSchemaPropertyInputFallback); err == nil {
		o.typ = "ObjectJsonSchemaPropertyInput"
		o.ObjectJsonSchemaPropertyInput = valueObjectJsonSchemaPropertyInputFallback
		return nil
	}
	var valueArrayJsonSchemaPropertyInputFallback *ArrayJsonSchemaPropertyInput
	if err := json.Unmarshal(data, &valueArrayJsonSchemaPropertyInputFallback); err == nil {
		o.typ = "ArrayJsonSchemaPropertyInput"
		o.ArrayJsonSchemaPropertyInput = valueArrayJsonSchemaPropertyInputFallback
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, o)
}

func (o ObjectJsonSchemaPropertyInputPropertiesValue) MarshalJSON() ([]byte, error) {
	if o.typ == "LiteralJsonSchemaProperty" || o.LiteralJsonSchemaProperty != nil {
		return json.Marshal(o.LiteralJsonSchemaProperty)
	}
	if o.typ == "ObjectJsonSchemaPropertyInput" || o.ObjectJsonSchemaPropertyInput != nil {
		return json.Marshal(o.ObjectJsonSchemaPropertyInput)
	}
	if o.typ == "ArrayJsonSchemaPropertyInput" || o.ArrayJsonSchemaPropertyInput != nil {
		return json.Marshal(o.ArrayJsonSchemaPropertyInput)
	}
	if o.typ == "AnyOfJsonSchemaPropertyInput" || o.AnyOfJsonSchemaPropertyInput != nil {
		return json.Marshal(o.AnyOfJsonSchemaPropertyInput)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", o)
}

type ObjectJsonSchemaPropertyInputPropertiesValueVisitor interface {
	VisitLiteralJsonSchemaProperty(*LiteralJsonSchemaProperty) error
	VisitObjectJsonSchemaPropertyInput(*ObjectJsonSchemaPropertyInput) error
	VisitArrayJsonSchemaPropertyInput(*ArrayJsonSchemaPropertyInput) error
	VisitAnyOfJsonSchemaPropertyInput(*AnyOfJsonSchemaPropertyInput) error
}

func (o *ObjectJsonSchemaPropertyInputPropertiesValue) Accept(visitor ObjectJsonSchemaPropertyInputPropertiesValueVisitor) error {
	if o.typ == "LiteralJsonSchemaProperty" || o.LiteralJsonSchemaProperty != nil {
		return visitor.VisitLiteralJsonSchemaProperty(o.LiteralJsonSchemaProperty)
	}
	if o.typ == "ObjectJsonSchemaPropertyInput" || o.ObjectJsonSchemaPropertyInput != nil {
		return visitor.VisitObjectJsonSchemaPropertyInput(o.ObjectJsonSchemaPropertyInput)
	}
	if o.typ == "ArrayJsonSchemaPropertyInput" || o.ArrayJsonSchemaPropertyInput != nil {
		return visitor.VisitArrayJsonSchemaPropertyInput(o.ArrayJsonSchemaPropertyInput)
	}
	if o.typ == "AnyOfJsonSchemaPropertyInput" || o.AnyOfJsonSchemaPropertyInput != nil {
		return visitor.VisitAnyOfJsonSchemaPropertyInput(o.AnyOfJsonSchemaPropertyInput)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", o)
}
