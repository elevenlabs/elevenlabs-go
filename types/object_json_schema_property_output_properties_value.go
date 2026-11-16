package types

import (
	"encoding/json"
	"fmt"
)

type ObjectJsonSchemaPropertyOutputPropertiesValue struct {
	LiteralJsonSchemaProperty      *LiteralJsonSchemaProperty
	ObjectJsonSchemaPropertyOutput *ObjectJsonSchemaPropertyOutput
	ArrayJsonSchemaPropertyOutput  *ArrayJsonSchemaPropertyOutput
	typ                            string
}

func (o *ObjectJsonSchemaPropertyOutputPropertiesValue) GetLiteralJsonSchemaProperty() *LiteralJsonSchemaProperty {
	if o == nil {
		return nil
	}
	if o.LiteralJsonSchemaProperty == nil {
		return nil
	}
	return o.LiteralJsonSchemaProperty
}

func (o *ObjectJsonSchemaPropertyOutputPropertiesValue) GetObjectJsonSchemaPropertyOutput() *ObjectJsonSchemaPropertyOutput {
	if o == nil {
		return nil
	}
	if o.ObjectJsonSchemaPropertyOutput == nil {
		return nil
	}
	return o.ObjectJsonSchemaPropertyOutput
}

func (o *ObjectJsonSchemaPropertyOutputPropertiesValue) GetArrayJsonSchemaPropertyOutput() *ArrayJsonSchemaPropertyOutput {
	if o == nil {
		return nil
	}
	if o.ArrayJsonSchemaPropertyOutput == nil {
		return nil
	}
	return o.ArrayJsonSchemaPropertyOutput
}

func (o *ObjectJsonSchemaPropertyOutputPropertiesValue) UnmarshalJSON(data []byte) error {
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
			o.typ = "ObjectJsonSchemaPropertyOutput"
			o.ObjectJsonSchemaPropertyOutput = valueObjectJsonSchemaPropertyOutputLiteral
			return nil
		}
	}
	if hasLiteral("type", "array") {
		matchedLiteral = true
		var valueArrayJsonSchemaPropertyOutputLiteral *ArrayJsonSchemaPropertyOutput
		if err := json.Unmarshal(data, &valueArrayJsonSchemaPropertyOutputLiteral); err == nil {
			o.typ = "ArrayJsonSchemaPropertyOutput"
			o.ArrayJsonSchemaPropertyOutput = valueArrayJsonSchemaPropertyOutputLiteral
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
	if hasAnyKey("description", "dynamic_variable", "constant_value", "is_omitted", "type", "required", "properties", "required_constraints") {
		var valueObjectJsonSchemaPropertyOutput *ObjectJsonSchemaPropertyOutput
		if err := json.Unmarshal(data, &valueObjectJsonSchemaPropertyOutput); err == nil {
			o.typ = "ObjectJsonSchemaPropertyOutput"
			o.ObjectJsonSchemaPropertyOutput = valueObjectJsonSchemaPropertyOutput
			return nil
		}
	}
	if hasAnyKey("description", "dynamic_variable", "constant_value", "is_omitted", "type", "items") {
		var valueArrayJsonSchemaPropertyOutput *ArrayJsonSchemaPropertyOutput
		if err := json.Unmarshal(data, &valueArrayJsonSchemaPropertyOutput); err == nil {
			o.typ = "ArrayJsonSchemaPropertyOutput"
			o.ArrayJsonSchemaPropertyOutput = valueArrayJsonSchemaPropertyOutput
			return nil
		}
	}

	// Retried without their key guards, for a payload that matched none of them.
	var valueObjectJsonSchemaPropertyOutputFallback *ObjectJsonSchemaPropertyOutput
	if err := json.Unmarshal(data, &valueObjectJsonSchemaPropertyOutputFallback); err == nil {
		o.typ = "ObjectJsonSchemaPropertyOutput"
		o.ObjectJsonSchemaPropertyOutput = valueObjectJsonSchemaPropertyOutputFallback
		return nil
	}
	var valueArrayJsonSchemaPropertyOutputFallback *ArrayJsonSchemaPropertyOutput
	if err := json.Unmarshal(data, &valueArrayJsonSchemaPropertyOutputFallback); err == nil {
		o.typ = "ArrayJsonSchemaPropertyOutput"
		o.ArrayJsonSchemaPropertyOutput = valueArrayJsonSchemaPropertyOutputFallback
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, o)
}

func (o ObjectJsonSchemaPropertyOutputPropertiesValue) MarshalJSON() ([]byte, error) {
	if o.typ == "LiteralJsonSchemaProperty" || o.LiteralJsonSchemaProperty != nil {
		return json.Marshal(o.LiteralJsonSchemaProperty)
	}
	if o.typ == "ObjectJsonSchemaPropertyOutput" || o.ObjectJsonSchemaPropertyOutput != nil {
		return json.Marshal(o.ObjectJsonSchemaPropertyOutput)
	}
	if o.typ == "ArrayJsonSchemaPropertyOutput" || o.ArrayJsonSchemaPropertyOutput != nil {
		return json.Marshal(o.ArrayJsonSchemaPropertyOutput)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", o)
}

type ObjectJsonSchemaPropertyOutputPropertiesValueVisitor interface {
	VisitLiteralJsonSchemaProperty(*LiteralJsonSchemaProperty) error
	VisitObjectJsonSchemaPropertyOutput(*ObjectJsonSchemaPropertyOutput) error
	VisitArrayJsonSchemaPropertyOutput(*ArrayJsonSchemaPropertyOutput) error
}

func (o *ObjectJsonSchemaPropertyOutputPropertiesValue) Accept(visitor ObjectJsonSchemaPropertyOutputPropertiesValueVisitor) error {
	if o.typ == "LiteralJsonSchemaProperty" || o.LiteralJsonSchemaProperty != nil {
		return visitor.VisitLiteralJsonSchemaProperty(o.LiteralJsonSchemaProperty)
	}
	if o.typ == "ObjectJsonSchemaPropertyOutput" || o.ObjectJsonSchemaPropertyOutput != nil {
		return visitor.VisitObjectJsonSchemaPropertyOutput(o.ObjectJsonSchemaPropertyOutput)
	}
	if o.typ == "ArrayJsonSchemaPropertyOutput" || o.ArrayJsonSchemaPropertyOutput != nil {
		return visitor.VisitArrayJsonSchemaPropertyOutput(o.ArrayJsonSchemaPropertyOutput)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", o)
}
