package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ObjectJsonSchemaPropertyInput struct {
	PropertyKind *ObjectJsonSchemaPropertyInputPropertyKind `json:"property_kind,omitempty" xml:"property_kind,omitempty"`
	Description  *string                                    `json:"description,omitempty" xml:"description,omitempty"`
	// When set, the entire parameter is populated from this dynamic variable at runtime. Mutually exclusive with description (LLM-provided value), constant_value, and is_omitted.
	DynamicVariable *string `json:"dynamic_variable,omitempty" xml:"dynamic_variable,omitempty"`
	// When set, the entire object uses this constant JSON value at runtime. Mutually exclusive with description (LLM-provided object), dynamic_variable, and is_omitted.
	ConstantValue map[string]any `json:"constant_value,omitempty"`
	// If true, this parameter will be completely omitted from the request. Only valid for optional parameters. Mutually exclusive with description, dynamic_variable, and constant_value.
	IsOmitted           *bool                                                    `json:"is_omitted,omitempty" xml:"is_omitted,omitempty"`
	Type                *string                                                  `json:"type,omitempty" xml:"type,omitempty"`
	Required            []string                                                 `json:"required,omitempty" xml:"required,omitempty"`
	Properties          map[string]*ObjectJsonSchemaPropertyInputPropertiesValue `json:"properties,omitempty"`
	RequiredConstraints *RequiredConstraints                                     `json:"required_constraints,omitempty" xml:"required_constraints,omitempty"`
	extraProperties     map[string]interface{}
}

func (o *ObjectJsonSchemaPropertyInput) GetPropertyKind() (value *ObjectJsonSchemaPropertyInputPropertyKind) {
	if o == nil {
		return
	}
	return o.PropertyKind
}

func (o *ObjectJsonSchemaPropertyInput) GetDescription() (value *string) {
	if o == nil {
		return
	}
	return o.Description
}

func (o *ObjectJsonSchemaPropertyInput) GetDynamicVariable() (value *string) {
	if o == nil {
		return
	}
	return o.DynamicVariable
}

func (o *ObjectJsonSchemaPropertyInput) GetConstantValue() (value map[string]any) {
	if o == nil {
		return
	}
	return o.ConstantValue
}

func (o *ObjectJsonSchemaPropertyInput) GetIsOmitted() (value *bool) {
	if o == nil {
		return
	}
	return o.IsOmitted
}

func (o *ObjectJsonSchemaPropertyInput) GetType() (value *string) {
	if o == nil {
		return
	}
	return o.Type
}

func (o *ObjectJsonSchemaPropertyInput) GetRequired() (value []string) {
	if o == nil {
		return
	}
	return o.Required
}

func (o *ObjectJsonSchemaPropertyInput) GetProperties() (value map[string]*ObjectJsonSchemaPropertyInputPropertiesValue) {
	if o == nil {
		return
	}
	return o.Properties
}

func (o *ObjectJsonSchemaPropertyInput) GetRequiredConstraints() (value *RequiredConstraints) {
	if o == nil {
		return
	}
	return o.RequiredConstraints
}

func (o *ObjectJsonSchemaPropertyInput) SetPropertyKind(propertyKind *ObjectJsonSchemaPropertyInputPropertyKind) {
	o.PropertyKind = propertyKind
}

func (o *ObjectJsonSchemaPropertyInput) SetDescription(description *string) {
	o.Description = description
}

func (o *ObjectJsonSchemaPropertyInput) SetDynamicVariable(dynamicVariable *string) {
	o.DynamicVariable = dynamicVariable
}

func (o *ObjectJsonSchemaPropertyInput) SetConstantValue(constantValue map[string]any) {
	o.ConstantValue = constantValue
}

func (o *ObjectJsonSchemaPropertyInput) SetIsOmitted(isOmitted *bool) {
	o.IsOmitted = isOmitted
}

func (o *ObjectJsonSchemaPropertyInput) SetType(type_ *string) {
	o.Type = type_
}

func (o *ObjectJsonSchemaPropertyInput) SetRequired(required []string) {
	o.Required = required
}

func (o *ObjectJsonSchemaPropertyInput) SetProperties(properties map[string]*ObjectJsonSchemaPropertyInputPropertiesValue) {
	o.Properties = properties
}

func (o *ObjectJsonSchemaPropertyInput) SetRequiredConstraints(requiredConstraints *RequiredConstraints) {
	o.RequiredConstraints = requiredConstraints
}

func (o *ObjectJsonSchemaPropertyInput) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o ObjectJsonSchemaPropertyInput) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: ObjectJsonSchemaPropertyInput to string"
	}
	return string(jsonData)
}

func (o *ObjectJsonSchemaPropertyInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, o); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *ObjectJsonSchemaPropertyInput) MarshalJSON() ([]byte, error) {
	type embed ObjectJsonSchemaPropertyInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
