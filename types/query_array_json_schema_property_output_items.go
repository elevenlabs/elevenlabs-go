package types

import (
	"encoding/json"
	"fmt"
)

// Schema for each value; must be a string, number or boolean.
type QueryArrayJsonSchemaPropertyOutputItems struct {
	LiteralJsonSchemaProperty      *LiteralJsonSchemaProperty
	ObjectJsonSchemaPropertyOutput *ObjectJsonSchemaPropertyOutput
	ArrayJsonSchemaPropertyOutput  *ArrayJsonSchemaPropertyOutput
	AnyOfJsonSchemaPropertyOutput  *AnyOfJsonSchemaPropertyOutput
	typ                            string
}

func (q *QueryArrayJsonSchemaPropertyOutputItems) GetLiteralJsonSchemaProperty() *LiteralJsonSchemaProperty {
	if q == nil {
		return nil
	}
	if q.LiteralJsonSchemaProperty == nil {
		return nil
	}
	return q.LiteralJsonSchemaProperty
}

func (q *QueryArrayJsonSchemaPropertyOutputItems) GetObjectJsonSchemaPropertyOutput() *ObjectJsonSchemaPropertyOutput {
	if q == nil {
		return nil
	}
	if q.ObjectJsonSchemaPropertyOutput == nil {
		return nil
	}
	return q.ObjectJsonSchemaPropertyOutput
}

func (q *QueryArrayJsonSchemaPropertyOutputItems) GetArrayJsonSchemaPropertyOutput() *ArrayJsonSchemaPropertyOutput {
	if q == nil {
		return nil
	}
	if q.ArrayJsonSchemaPropertyOutput == nil {
		return nil
	}
	return q.ArrayJsonSchemaPropertyOutput
}

func (q *QueryArrayJsonSchemaPropertyOutputItems) GetAnyOfJsonSchemaPropertyOutput() *AnyOfJsonSchemaPropertyOutput {
	if q == nil {
		return nil
	}
	if q.AnyOfJsonSchemaPropertyOutput == nil {
		return nil
	}
	return q.AnyOfJsonSchemaPropertyOutput
}

func (q *QueryArrayJsonSchemaPropertyOutputItems) UnmarshalJSON(data []byte) error {
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
			q.typ = "ObjectJsonSchemaPropertyOutput"
			q.ObjectJsonSchemaPropertyOutput = valueObjectJsonSchemaPropertyOutputLiteral
			return nil
		}
	}
	if hasLiteral("type", "array") {
		matchedLiteral = true
		var valueArrayJsonSchemaPropertyOutputLiteral *ArrayJsonSchemaPropertyOutput
		if err := json.Unmarshal(data, &valueArrayJsonSchemaPropertyOutputLiteral); err == nil {
			q.typ = "ArrayJsonSchemaPropertyOutput"
			q.ArrayJsonSchemaPropertyOutput = valueArrayJsonSchemaPropertyOutputLiteral
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
	if hasAnyKey("description", "dynamic_variable", "constant_value", "is_omitted", "type", "required", "properties", "required_constraints") {
		var valueObjectJsonSchemaPropertyOutput *ObjectJsonSchemaPropertyOutput
		if err := json.Unmarshal(data, &valueObjectJsonSchemaPropertyOutput); err == nil {
			q.typ = "ObjectJsonSchemaPropertyOutput"
			q.ObjectJsonSchemaPropertyOutput = valueObjectJsonSchemaPropertyOutput
			return nil
		}
	}
	if hasAnyKey("description", "dynamic_variable", "constant_value", "is_omitted", "type", "items") {
		var valueArrayJsonSchemaPropertyOutput *ArrayJsonSchemaPropertyOutput
		if err := json.Unmarshal(data, &valueArrayJsonSchemaPropertyOutput); err == nil {
			q.typ = "ArrayJsonSchemaPropertyOutput"
			q.ArrayJsonSchemaPropertyOutput = valueArrayJsonSchemaPropertyOutput
			return nil
		}
	}
	var valueAnyOfJsonSchemaPropertyOutput *AnyOfJsonSchemaPropertyOutput
	if err := json.Unmarshal(data, &valueAnyOfJsonSchemaPropertyOutput); err == nil {
		q.typ = "AnyOfJsonSchemaPropertyOutput"
		q.AnyOfJsonSchemaPropertyOutput = valueAnyOfJsonSchemaPropertyOutput
		return nil
	}

	// Retried without their key guards, for a payload that matched none of them.
	var valueObjectJsonSchemaPropertyOutputFallback *ObjectJsonSchemaPropertyOutput
	if err := json.Unmarshal(data, &valueObjectJsonSchemaPropertyOutputFallback); err == nil {
		q.typ = "ObjectJsonSchemaPropertyOutput"
		q.ObjectJsonSchemaPropertyOutput = valueObjectJsonSchemaPropertyOutputFallback
		return nil
	}
	var valueArrayJsonSchemaPropertyOutputFallback *ArrayJsonSchemaPropertyOutput
	if err := json.Unmarshal(data, &valueArrayJsonSchemaPropertyOutputFallback); err == nil {
		q.typ = "ArrayJsonSchemaPropertyOutput"
		q.ArrayJsonSchemaPropertyOutput = valueArrayJsonSchemaPropertyOutputFallback
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, q)
}

func (q QueryArrayJsonSchemaPropertyOutputItems) MarshalJSON() ([]byte, error) {
	if q.typ == "LiteralJsonSchemaProperty" || q.LiteralJsonSchemaProperty != nil {
		return json.Marshal(q.LiteralJsonSchemaProperty)
	}
	if q.typ == "ObjectJsonSchemaPropertyOutput" || q.ObjectJsonSchemaPropertyOutput != nil {
		return json.Marshal(q.ObjectJsonSchemaPropertyOutput)
	}
	if q.typ == "ArrayJsonSchemaPropertyOutput" || q.ArrayJsonSchemaPropertyOutput != nil {
		return json.Marshal(q.ArrayJsonSchemaPropertyOutput)
	}
	if q.typ == "AnyOfJsonSchemaPropertyOutput" || q.AnyOfJsonSchemaPropertyOutput != nil {
		return json.Marshal(q.AnyOfJsonSchemaPropertyOutput)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", q)
}

type QueryArrayJsonSchemaPropertyOutputItemsVisitor interface {
	VisitLiteralJsonSchemaProperty(*LiteralJsonSchemaProperty) error
	VisitObjectJsonSchemaPropertyOutput(*ObjectJsonSchemaPropertyOutput) error
	VisitArrayJsonSchemaPropertyOutput(*ArrayJsonSchemaPropertyOutput) error
	VisitAnyOfJsonSchemaPropertyOutput(*AnyOfJsonSchemaPropertyOutput) error
}

func (q *QueryArrayJsonSchemaPropertyOutputItems) Accept(visitor QueryArrayJsonSchemaPropertyOutputItemsVisitor) error {
	if q.typ == "LiteralJsonSchemaProperty" || q.LiteralJsonSchemaProperty != nil {
		return visitor.VisitLiteralJsonSchemaProperty(q.LiteralJsonSchemaProperty)
	}
	if q.typ == "ObjectJsonSchemaPropertyOutput" || q.ObjectJsonSchemaPropertyOutput != nil {
		return visitor.VisitObjectJsonSchemaPropertyOutput(q.ObjectJsonSchemaPropertyOutput)
	}
	if q.typ == "ArrayJsonSchemaPropertyOutput" || q.ArrayJsonSchemaPropertyOutput != nil {
		return visitor.VisitArrayJsonSchemaPropertyOutput(q.ArrayJsonSchemaPropertyOutput)
	}
	if q.typ == "AnyOfJsonSchemaPropertyOutput" || q.AnyOfJsonSchemaPropertyOutput != nil {
		return visitor.VisitAnyOfJsonSchemaPropertyOutput(q.AnyOfJsonSchemaPropertyOutput)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", q)
}
