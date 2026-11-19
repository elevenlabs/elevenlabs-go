package types

import (
	"encoding/json"
	"fmt"
)

type QueryParamsJsonSchemaInputPropertiesValue struct {
	LiteralJsonSchemaProperty         *LiteralJsonSchemaProperty
	QueryArrayJsonSchemaPropertyInput *QueryArrayJsonSchemaPropertyInput
	typ                               string
}

func (q *QueryParamsJsonSchemaInputPropertiesValue) GetLiteralJsonSchemaProperty() *LiteralJsonSchemaProperty {
	if q == nil {
		return nil
	}
	if q.LiteralJsonSchemaProperty == nil {
		return nil
	}
	return q.LiteralJsonSchemaProperty
}

func (q *QueryParamsJsonSchemaInputPropertiesValue) GetQueryArrayJsonSchemaPropertyInput() *QueryArrayJsonSchemaPropertyInput {
	if q == nil {
		return nil
	}
	if q.QueryArrayJsonSchemaPropertyInput == nil {
		return nil
	}
	return q.QueryArrayJsonSchemaPropertyInput
}

func (q *QueryParamsJsonSchemaInputPropertiesValue) UnmarshalJSON(data []byte) error {
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
	if hasLiteral("type", "array") {
		matchedLiteral = true
		var valueQueryArrayJsonSchemaPropertyInputLiteral *QueryArrayJsonSchemaPropertyInput
		if err := json.Unmarshal(data, &valueQueryArrayJsonSchemaPropertyInputLiteral); err == nil {
			q.typ = "QueryArrayJsonSchemaPropertyInput"
			q.QueryArrayJsonSchemaPropertyInput = valueQueryArrayJsonSchemaPropertyInputLiteral
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
	var valueQueryArrayJsonSchemaPropertyInput *QueryArrayJsonSchemaPropertyInput
	if err := json.Unmarshal(data, &valueQueryArrayJsonSchemaPropertyInput); err == nil {
		q.typ = "QueryArrayJsonSchemaPropertyInput"
		q.QueryArrayJsonSchemaPropertyInput = valueQueryArrayJsonSchemaPropertyInput
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, q)
}

func (q QueryParamsJsonSchemaInputPropertiesValue) MarshalJSON() ([]byte, error) {
	if q.typ == "LiteralJsonSchemaProperty" || q.LiteralJsonSchemaProperty != nil {
		return json.Marshal(q.LiteralJsonSchemaProperty)
	}
	if q.typ == "QueryArrayJsonSchemaPropertyInput" || q.QueryArrayJsonSchemaPropertyInput != nil {
		return json.Marshal(q.QueryArrayJsonSchemaPropertyInput)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", q)
}

type QueryParamsJsonSchemaInputPropertiesValueVisitor interface {
	VisitLiteralJsonSchemaProperty(*LiteralJsonSchemaProperty) error
	VisitQueryArrayJsonSchemaPropertyInput(*QueryArrayJsonSchemaPropertyInput) error
}

func (q *QueryParamsJsonSchemaInputPropertiesValue) Accept(visitor QueryParamsJsonSchemaInputPropertiesValueVisitor) error {
	if q.typ == "LiteralJsonSchemaProperty" || q.LiteralJsonSchemaProperty != nil {
		return visitor.VisitLiteralJsonSchemaProperty(q.LiteralJsonSchemaProperty)
	}
	if q.typ == "QueryArrayJsonSchemaPropertyInput" || q.QueryArrayJsonSchemaPropertyInput != nil {
		return visitor.VisitQueryArrayJsonSchemaPropertyInput(q.QueryArrayJsonSchemaPropertyInput)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", q)
}
