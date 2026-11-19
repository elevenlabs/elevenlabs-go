package types

import (
	"encoding/json"
	"fmt"
)

type QueryParamsJsonSchemaOutputPropertiesValue struct {
	LiteralJsonSchemaProperty          *LiteralJsonSchemaProperty
	QueryArrayJsonSchemaPropertyOutput *QueryArrayJsonSchemaPropertyOutput
	typ                                string
}

func (q *QueryParamsJsonSchemaOutputPropertiesValue) GetLiteralJsonSchemaProperty() *LiteralJsonSchemaProperty {
	if q == nil {
		return nil
	}
	if q.LiteralJsonSchemaProperty == nil {
		return nil
	}
	return q.LiteralJsonSchemaProperty
}

func (q *QueryParamsJsonSchemaOutputPropertiesValue) GetQueryArrayJsonSchemaPropertyOutput() *QueryArrayJsonSchemaPropertyOutput {
	if q == nil {
		return nil
	}
	if q.QueryArrayJsonSchemaPropertyOutput == nil {
		return nil
	}
	return q.QueryArrayJsonSchemaPropertyOutput
}

func (q *QueryParamsJsonSchemaOutputPropertiesValue) UnmarshalJSON(data []byte) error {
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
		var valueQueryArrayJsonSchemaPropertyOutputLiteral *QueryArrayJsonSchemaPropertyOutput
		if err := json.Unmarshal(data, &valueQueryArrayJsonSchemaPropertyOutputLiteral); err == nil {
			q.typ = "QueryArrayJsonSchemaPropertyOutput"
			q.QueryArrayJsonSchemaPropertyOutput = valueQueryArrayJsonSchemaPropertyOutputLiteral
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
	var valueQueryArrayJsonSchemaPropertyOutput *QueryArrayJsonSchemaPropertyOutput
	if err := json.Unmarshal(data, &valueQueryArrayJsonSchemaPropertyOutput); err == nil {
		q.typ = "QueryArrayJsonSchemaPropertyOutput"
		q.QueryArrayJsonSchemaPropertyOutput = valueQueryArrayJsonSchemaPropertyOutput
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, q)
}

func (q QueryParamsJsonSchemaOutputPropertiesValue) MarshalJSON() ([]byte, error) {
	if q.typ == "LiteralJsonSchemaProperty" || q.LiteralJsonSchemaProperty != nil {
		return json.Marshal(q.LiteralJsonSchemaProperty)
	}
	if q.typ == "QueryArrayJsonSchemaPropertyOutput" || q.QueryArrayJsonSchemaPropertyOutput != nil {
		return json.Marshal(q.QueryArrayJsonSchemaPropertyOutput)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", q)
}

type QueryParamsJsonSchemaOutputPropertiesValueVisitor interface {
	VisitLiteralJsonSchemaProperty(*LiteralJsonSchemaProperty) error
	VisitQueryArrayJsonSchemaPropertyOutput(*QueryArrayJsonSchemaPropertyOutput) error
}

func (q *QueryParamsJsonSchemaOutputPropertiesValue) Accept(visitor QueryParamsJsonSchemaOutputPropertiesValueVisitor) error {
	if q.typ == "LiteralJsonSchemaProperty" || q.LiteralJsonSchemaProperty != nil {
		return visitor.VisitLiteralJsonSchemaProperty(q.LiteralJsonSchemaProperty)
	}
	if q.typ == "QueryArrayJsonSchemaPropertyOutput" || q.QueryArrayJsonSchemaPropertyOutput != nil {
		return visitor.VisitQueryArrayJsonSchemaPropertyOutput(q.QueryArrayJsonSchemaPropertyOutput)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", q)
}
