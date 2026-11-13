package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ArrayJsonSchemaPropertyOutput struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// When set, the entire parameter is populated from this dynamic variable at runtime. Mutually exclusive with description (LLM-provided value), constant_value, and is_omitted.
	DynamicVariable *string `json:"dynamic_variable,omitempty" xml:"dynamic_variable,omitempty"`
	// When set, the entire array uses this constant value at runtime. Mutually exclusive with description (LLM-provided array), dynamic_variable, and is_omitted.
	ConstantValue []any `json:"constant_value,omitempty" xml:"constant_value,omitempty"`
	// If true, this parameter will be completely omitted from the request. Only valid for optional parameters. Mutually exclusive with description, dynamic_variable, and constant_value.
	IsOmitted *bool   `json:"is_omitted,omitempty" xml:"is_omitted,omitempty"`
	Type      *string `json:"type,omitempty" xml:"type,omitempty"`
	// Schema for array elements.
	Items           *ArrayJsonSchemaPropertyOutputItems `json:"items,omitempty" xml:"items,omitempty"`
	extraProperties map[string]interface{}
}

func (a *ArrayJsonSchemaPropertyOutput) GetDescription() (value *string) {
	if a == nil {
		return
	}
	return a.Description
}

func (a *ArrayJsonSchemaPropertyOutput) GetDynamicVariable() (value *string) {
	if a == nil {
		return
	}
	return a.DynamicVariable
}

func (a *ArrayJsonSchemaPropertyOutput) GetConstantValue() (value []any) {
	if a == nil {
		return
	}
	return a.ConstantValue
}

func (a *ArrayJsonSchemaPropertyOutput) GetIsOmitted() (value *bool) {
	if a == nil {
		return
	}
	return a.IsOmitted
}

func (a *ArrayJsonSchemaPropertyOutput) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *ArrayJsonSchemaPropertyOutput) GetItems() (value *ArrayJsonSchemaPropertyOutputItems) {
	if a == nil {
		return
	}
	return a.Items
}

func (a *ArrayJsonSchemaPropertyOutput) SetDescription(description *string) {
	a.Description = description
}

func (a *ArrayJsonSchemaPropertyOutput) SetDynamicVariable(dynamicVariable *string) {
	a.DynamicVariable = dynamicVariable
}

func (a *ArrayJsonSchemaPropertyOutput) SetConstantValue(constantValue []any) {
	a.ConstantValue = constantValue
}

func (a *ArrayJsonSchemaPropertyOutput) SetIsOmitted(isOmitted *bool) {
	a.IsOmitted = isOmitted
}

func (a *ArrayJsonSchemaPropertyOutput) SetType(type_ *string) {
	a.Type = type_
}

func (a *ArrayJsonSchemaPropertyOutput) SetItems(items *ArrayJsonSchemaPropertyOutputItems) {
	a.Items = items
}

func (a *ArrayJsonSchemaPropertyOutput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a ArrayJsonSchemaPropertyOutput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: ArrayJsonSchemaPropertyOutput to string"
	}
	return string(jsonData)
}

func (a *ArrayJsonSchemaPropertyOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *ArrayJsonSchemaPropertyOutput) MarshalJSON() ([]byte, error) {
	type embed ArrayJsonSchemaPropertyOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
