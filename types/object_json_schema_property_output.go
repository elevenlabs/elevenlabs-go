package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ObjectJsonSchemaPropertyOutput struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// When set, the entire parameter is populated from this dynamic variable at runtime. Mutually exclusive with description (LLM-provided value), constant_value, and is_omitted.
	DynamicVariable *string `json:"dynamic_variable,omitempty" xml:"dynamic_variable,omitempty"`
	// When set, the entire object uses this constant JSON value at runtime. Mutually exclusive with description (LLM-provided object), dynamic_variable, and is_omitted.
	ConstantValue map[string]any `json:"constant_value,omitempty"`
	// If true, this parameter will be completely omitted from the request. Only valid for optional parameters. Mutually exclusive with description, dynamic_variable, and constant_value.
	IsOmitted           *bool                                                     `json:"is_omitted,omitempty" xml:"is_omitted,omitempty"`
	Type                *string                                                   `json:"type,omitempty" xml:"type,omitempty"`
	Required            []string                                                  `json:"required,omitempty" xml:"required,omitempty"`
	Properties          map[string]*ObjectJsonSchemaPropertyOutputPropertiesValue `json:"properties,omitempty"`
	RequiredConstraints *RequiredConstraints                                      `json:"required_constraints,omitempty" xml:"required_constraints,omitempty"`
	extraProperties     map[string]interface{}
}

func (o *ObjectJsonSchemaPropertyOutput) GetDescription() (value *string) {
	if o == nil {
		return
	}
	return o.Description
}

func (o *ObjectJsonSchemaPropertyOutput) GetDynamicVariable() (value *string) {
	if o == nil {
		return
	}
	return o.DynamicVariable
}

func (o *ObjectJsonSchemaPropertyOutput) GetConstantValue() (value map[string]any) {
	if o == nil {
		return
	}
	return o.ConstantValue
}

func (o *ObjectJsonSchemaPropertyOutput) GetIsOmitted() (value *bool) {
	if o == nil {
		return
	}
	return o.IsOmitted
}

func (o *ObjectJsonSchemaPropertyOutput) GetType() (value *string) {
	if o == nil {
		return
	}
	return o.Type
}

func (o *ObjectJsonSchemaPropertyOutput) GetRequired() (value []string) {
	if o == nil {
		return
	}
	return o.Required
}

func (o *ObjectJsonSchemaPropertyOutput) GetProperties() (value map[string]*ObjectJsonSchemaPropertyOutputPropertiesValue) {
	if o == nil {
		return
	}
	return o.Properties
}

func (o *ObjectJsonSchemaPropertyOutput) GetRequiredConstraints() (value *RequiredConstraints) {
	if o == nil {
		return
	}
	return o.RequiredConstraints
}

func (o *ObjectJsonSchemaPropertyOutput) SetDescription(description *string) {
	o.Description = description
}

func (o *ObjectJsonSchemaPropertyOutput) SetDynamicVariable(dynamicVariable *string) {
	o.DynamicVariable = dynamicVariable
}

func (o *ObjectJsonSchemaPropertyOutput) SetConstantValue(constantValue map[string]any) {
	o.ConstantValue = constantValue
}

func (o *ObjectJsonSchemaPropertyOutput) SetIsOmitted(isOmitted *bool) {
	o.IsOmitted = isOmitted
}

func (o *ObjectJsonSchemaPropertyOutput) SetType(type_ *string) {
	o.Type = type_
}

func (o *ObjectJsonSchemaPropertyOutput) SetRequired(required []string) {
	o.Required = required
}

func (o *ObjectJsonSchemaPropertyOutput) SetProperties(properties map[string]*ObjectJsonSchemaPropertyOutputPropertiesValue) {
	o.Properties = properties
}

func (o *ObjectJsonSchemaPropertyOutput) SetRequiredConstraints(requiredConstraints *RequiredConstraints) {
	o.RequiredConstraints = requiredConstraints
}

func (o *ObjectJsonSchemaPropertyOutput) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o ObjectJsonSchemaPropertyOutput) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: ObjectJsonSchemaPropertyOutput to string"
	}
	return string(jsonData)
}

func (o *ObjectJsonSchemaPropertyOutput) UnmarshalJSON(data []byte) error {
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

func (o *ObjectJsonSchemaPropertyOutput) MarshalJSON() ([]byte, error) {
	type embed ObjectJsonSchemaPropertyOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
