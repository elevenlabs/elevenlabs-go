package types

import (
	"encoding/json"
	"fmt"
)

// Schema for each value; must be a string, number or boolean.
type QueryArrayJsonSchemaPropertyInputItems struct {
	LiteralJsonSchemaProperty     *LiteralJsonSchemaProperty
	ObjectJsonSchemaPropertyInput *ObjectJsonSchemaPropertyInput
	ArrayJsonSchemaPropertyInput  *ArrayJsonSchemaPropertyInput
	typ                           string
}

func (q *QueryArrayJsonSchemaPropertyInputItems) GetLiteralJsonSchemaProperty() *LiteralJsonSchemaProperty {
	if q == nil {
		return nil
	}
	if q.LiteralJsonSchemaProperty == nil {
		return nil
	}
	return q.LiteralJsonSchemaProperty
}

func (q *QueryArrayJsonSchemaPropertyInputItems) GetObjectJsonSchemaPropertyInput() *ObjectJsonSchemaPropertyInput {
	if q == nil {
		return nil
	}
	if q.ObjectJsonSchemaPropertyInput == nil {
		return nil
	}
	return q.ObjectJsonSchemaPropertyInput
}

func (q *QueryArrayJsonSchemaPropertyInputItems) GetArrayJsonSchemaPropertyInput() *ArrayJsonSchemaPropertyInput {
	if q == nil {
		return nil
	}
	if q.ArrayJsonSchemaPropertyInput == nil {
		return nil
	}
	return q.ArrayJsonSchemaPropertyInput
}

func (q *QueryArrayJsonSchemaPropertyInputItems) UnmarshalJSON(data []byte) error {
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
			q.typ = "ObjectJsonSchemaPropertyInput"
			q.ObjectJsonSchemaPropertyInput = valueObjectJsonSchemaPropertyInputLiteral
			return nil
		}
	}
	if hasLiteral("type", "array") {
		matchedLiteral = true
		var valueArrayJsonSchemaPropertyInputLiteral *ArrayJsonSchemaPropertyInput
		if err := json.Unmarshal(data, &valueArrayJsonSchemaPropertyInputLiteral); err == nil {
			q.typ = "ArrayJsonSchemaPropertyInput"
			q.ArrayJsonSchemaPropertyInput = valueArrayJsonSchemaPropertyInputLiteral
			return nil
		}
	}
	if matchedLiteral {
		return fmt.Errorf("%s cannot be deserialized as a %T", data, q)
	}
	var valueLiteralJsonSchemaProperty *LiteralJsonSchemaProperty
	if err := json.Unmarshal(data, &valueLiteralJsonSchemaProperty); err == nil {
		q.typ = "LiteralJsonSchemaProperty"
		q.LiteralJsonSchemaProperty = valueLiteralJsonSchemaProperty
		return nil
	}
	if hasAnyKey("property_kind", "description", "dynamic_variable", "constant_value", "is_omitted", "type", "required", "properties", "required_constraints") {
		var valueObjectJsonSchemaPropertyInput *ObjectJsonSchemaPropertyInput
		if err := json.Unmarshal(data, &valueObjectJsonSchemaPropertyInput); err == nil {
			q.typ = "ObjectJsonSchemaPropertyInput"
			q.ObjectJsonSchemaPropertyInput = valueObjectJsonSchemaPropertyInput
			return nil
		}
	}
	if hasAnyKey("property_kind", "description", "dynamic_variable", "constant_value", "is_omitted", "type", "items") {
		var valueArrayJsonSchemaPropertyInput *ArrayJsonSchemaPropertyInput
		if err := json.Unmarshal(data, &valueArrayJsonSchemaPropertyInput); err == nil {
			q.typ = "ArrayJsonSchemaPropertyInput"
			q.ArrayJsonSchemaPropertyInput = valueArrayJsonSchemaPropertyInput
			return nil
		}
	}

	// Retried without their key guards, for a payload that matched none of them.
	var valueObjectJsonSchemaPropertyInputFallback *ObjectJsonSchemaPropertyInput
	if err := json.Unmarshal(data, &valueObjectJsonSchemaPropertyInputFallback); err == nil {
		q.typ = "ObjectJsonSchemaPropertyInput"
		q.ObjectJsonSchemaPropertyInput = valueObjectJsonSchemaPropertyInputFallback
		return nil
	}
	var valueArrayJsonSchemaPropertyInputFallback *ArrayJsonSchemaPropertyInput
	if err := json.Unmarshal(data, &valueArrayJsonSchemaPropertyInputFallback); err == nil {
		q.typ = "ArrayJsonSchemaPropertyInput"
		q.ArrayJsonSchemaPropertyInput = valueArrayJsonSchemaPropertyInputFallback
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, q)
}

func (q QueryArrayJsonSchemaPropertyInputItems) MarshalJSON() ([]byte, error) {
	if q.typ == "LiteralJsonSchemaProperty" || q.LiteralJsonSchemaProperty != nil {
		return json.Marshal(q.LiteralJsonSchemaProperty)
	}
	if q.typ == "ObjectJsonSchemaPropertyInput" || q.ObjectJsonSchemaPropertyInput != nil {
		return json.Marshal(q.ObjectJsonSchemaPropertyInput)
	}
	if q.typ == "ArrayJsonSchemaPropertyInput" || q.ArrayJsonSchemaPropertyInput != nil {
		return json.Marshal(q.ArrayJsonSchemaPropertyInput)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", q)
}

type QueryArrayJsonSchemaPropertyInputItemsVisitor interface {
	VisitLiteralJsonSchemaProperty(*LiteralJsonSchemaProperty) error
	VisitObjectJsonSchemaPropertyInput(*ObjectJsonSchemaPropertyInput) error
	VisitArrayJsonSchemaPropertyInput(*ArrayJsonSchemaPropertyInput) error
}

func (q *QueryArrayJsonSchemaPropertyInputItems) Accept(visitor QueryArrayJsonSchemaPropertyInputItemsVisitor) error {
	if q.typ == "LiteralJsonSchemaProperty" || q.LiteralJsonSchemaProperty != nil {
		return visitor.VisitLiteralJsonSchemaProperty(q.LiteralJsonSchemaProperty)
	}
	if q.typ == "ObjectJsonSchemaPropertyInput" || q.ObjectJsonSchemaPropertyInput != nil {
		return visitor.VisitObjectJsonSchemaPropertyInput(q.ObjectJsonSchemaPropertyInput)
	}
	if q.typ == "ArrayJsonSchemaPropertyInput" || q.ArrayJsonSchemaPropertyInput != nil {
		return visitor.VisitArrayJsonSchemaPropertyInput(q.ArrayJsonSchemaPropertyInput)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", q)
}
