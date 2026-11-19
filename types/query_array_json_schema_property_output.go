package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A query param holding several literals, sent as a repeated key (?a=x&a=y).
type QueryArrayJsonSchemaPropertyOutput struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// When set, the entire parameter is populated from this dynamic variable at runtime. Mutually exclusive with description (LLM-provided value), constant_value, and is_omitted.
	DynamicVariable *string `json:"dynamic_variable,omitempty" xml:"dynamic_variable,omitempty"`
	// When set, the entire array uses this constant value at runtime. Mutually exclusive with description (LLM-provided array), dynamic_variable, and is_omitted.
	ConstantValue []any `json:"constant_value,omitempty" xml:"constant_value,omitempty"`
	// If true, this parameter will be completely omitted from the request. Only valid for optional parameters. Mutually exclusive with description, dynamic_variable, and constant_value.
	IsOmitted *bool `json:"is_omitted,omitempty" xml:"is_omitted,omitempty"`
	type_     string
	// Schema for each value; must be a string, number or boolean.
	Items           *QueryArrayJsonSchemaPropertyOutputItems `json:"items,omitempty" xml:"items,omitempty"`
	extraProperties map[string]interface{}
}

func (q *QueryArrayJsonSchemaPropertyOutput) GetDescription() (value *string) {
	if q == nil {
		return
	}
	return q.Description
}

func (q *QueryArrayJsonSchemaPropertyOutput) GetDynamicVariable() (value *string) {
	if q == nil {
		return
	}
	return q.DynamicVariable
}

func (q *QueryArrayJsonSchemaPropertyOutput) GetConstantValue() (value []any) {
	if q == nil {
		return
	}
	return q.ConstantValue
}

func (q *QueryArrayJsonSchemaPropertyOutput) GetIsOmitted() (value *bool) {
	if q == nil {
		return
	}
	return q.IsOmitted
}

func (q *QueryArrayJsonSchemaPropertyOutput) GetItems() (value *QueryArrayJsonSchemaPropertyOutputItems) {
	if q == nil {
		return
	}
	return q.Items
}

func (q *QueryArrayJsonSchemaPropertyOutput) SetDescription(description *string) {
	q.Description = description
}

func (q *QueryArrayJsonSchemaPropertyOutput) SetDynamicVariable(dynamicVariable *string) {
	q.DynamicVariable = dynamicVariable
}

func (q *QueryArrayJsonSchemaPropertyOutput) SetConstantValue(constantValue []any) {
	q.ConstantValue = constantValue
}

func (q *QueryArrayJsonSchemaPropertyOutput) SetIsOmitted(isOmitted *bool) {
	q.IsOmitted = isOmitted
}

func (q *QueryArrayJsonSchemaPropertyOutput) SetItems(items *QueryArrayJsonSchemaPropertyOutputItems) {
	q.Items = items
}

func (q *QueryArrayJsonSchemaPropertyOutput) Type() string {
	return q.type_
}

func (q *QueryArrayJsonSchemaPropertyOutput) GetExtraProperties() map[string]interface{} {
	if q == nil {
		return nil
	}
	return q.extraProperties
}

func (q QueryArrayJsonSchemaPropertyOutput) String() string {
	jsonData, err := json.MarshalIndent(q, "", "  ")
	if err != nil {
		return "error converting struct: QueryArrayJsonSchemaPropertyOutput to string"
	}
	return string(jsonData)
}

func (q *QueryArrayJsonSchemaPropertyOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, q); err != nil {
		return err
	}
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "array" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", q, "array", *literals.Type)
	}
	q.type_ = "array"
	extraProperties, err := internal.ExtractExtraProperties(data, *q, "type")
	if err != nil {
		return err
	}
	q.extraProperties = extraProperties
	return nil
}

func (q *QueryArrayJsonSchemaPropertyOutput) MarshalJSON() ([]byte, error) {
	type embed QueryArrayJsonSchemaPropertyOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*q),
		Type:  "array",
	}
	return json.Marshal(marshaler)
}
