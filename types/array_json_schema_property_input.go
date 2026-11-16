package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ArrayJsonSchemaPropertyInput struct {
	PropertyKind *ObjectJsonSchemaPropertyInputPropertyKind `json:"property_kind,omitempty" xml:"property_kind,omitempty"`
	Description  *string                                    `json:"description,omitempty" xml:"description,omitempty"`
	// When set, the entire parameter is populated from this dynamic variable at runtime. Mutually exclusive with description (LLM-provided value), constant_value, and is_omitted.
	DynamicVariable *string `json:"dynamic_variable,omitempty" xml:"dynamic_variable,omitempty"`
	// When set, the entire array uses this constant value at runtime. Mutually exclusive with description (LLM-provided array), dynamic_variable, and is_omitted.
	ConstantValue []any `json:"constant_value,omitempty" xml:"constant_value,omitempty"`
	// If true, this parameter will be completely omitted from the request. Only valid for optional parameters. Mutually exclusive with description, dynamic_variable, and constant_value.
	IsOmitted *bool   `json:"is_omitted,omitempty" xml:"is_omitted,omitempty"`
	Type      *string `json:"type,omitempty" xml:"type,omitempty"`
	// Schema for array elements.
	Items           *ArrayJsonSchemaPropertyInputItems `json:"items,omitempty" xml:"items,omitempty"`
	extraProperties map[string]interface{}
}

func (a *ArrayJsonSchemaPropertyInput) GetPropertyKind() (value *ObjectJsonSchemaPropertyInputPropertyKind) {
	if a == nil {
		return
	}
	return a.PropertyKind
}

func (a *ArrayJsonSchemaPropertyInput) GetDescription() (value *string) {
	if a == nil {
		return
	}
	return a.Description
}

func (a *ArrayJsonSchemaPropertyInput) GetDynamicVariable() (value *string) {
	if a == nil {
		return
	}
	return a.DynamicVariable
}

func (a *ArrayJsonSchemaPropertyInput) GetConstantValue() (value []any) {
	if a == nil {
		return
	}
	return a.ConstantValue
}

func (a *ArrayJsonSchemaPropertyInput) GetIsOmitted() (value *bool) {
	if a == nil {
		return
	}
	return a.IsOmitted
}

func (a *ArrayJsonSchemaPropertyInput) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *ArrayJsonSchemaPropertyInput) GetItems() (value *ArrayJsonSchemaPropertyInputItems) {
	if a == nil {
		return
	}
	return a.Items
}

func (a *ArrayJsonSchemaPropertyInput) SetPropertyKind(propertyKind *ObjectJsonSchemaPropertyInputPropertyKind) {
	a.PropertyKind = propertyKind
}

func (a *ArrayJsonSchemaPropertyInput) SetDescription(description *string) {
	a.Description = description
}

func (a *ArrayJsonSchemaPropertyInput) SetDynamicVariable(dynamicVariable *string) {
	a.DynamicVariable = dynamicVariable
}

func (a *ArrayJsonSchemaPropertyInput) SetConstantValue(constantValue []any) {
	a.ConstantValue = constantValue
}

func (a *ArrayJsonSchemaPropertyInput) SetIsOmitted(isOmitted *bool) {
	a.IsOmitted = isOmitted
}

func (a *ArrayJsonSchemaPropertyInput) SetType(type_ *string) {
	a.Type = type_
}

func (a *ArrayJsonSchemaPropertyInput) SetItems(items *ArrayJsonSchemaPropertyInputItems) {
	a.Items = items
}

func (a *ArrayJsonSchemaPropertyInput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a ArrayJsonSchemaPropertyInput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: ArrayJsonSchemaPropertyInput to string"
	}
	return string(jsonData)
}

func (a *ArrayJsonSchemaPropertyInput) UnmarshalJSON(data []byte) error {
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

func (a *ArrayJsonSchemaPropertyInput) MarshalJSON() ([]byte, error) {
	type embed ArrayJsonSchemaPropertyInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
