package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Schema property for literal JSON types. IMPORTANT: Only ONE of the following fields can be set: description (LLM provides value), dynamic_variable (value from variable), is_system_provided (system provides value), constant_value (fixed value), or is_omitted (parameter is omitted). These are mutually exclusive.
type DataCollectionValueSchema struct {
	Type *DataCollectionValueSchemaType `json:"type" xml:"type" required:"true"`
	// The description of the property. When set, the LLM will provide the value based on this description. Mutually exclusive with dynamic_variable, is_system_provided, constant_value, and is_omitted.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// List of allowed string values for string type parameters
	Enum []string `json:"enum,omitempty" xml:"enum,omitempty"`
	// If true, the value will be populated by the system at runtime. Used by API Integration Webhook tools for templating. Mutually exclusive with description, dynamic_variable, constant_value, and is_omitted.
	IsSystemProvided *bool `json:"is_system_provided,omitempty" xml:"is_system_provided,omitempty"`
	// The name of the dynamic variable to use for this property's value. Mutually exclusive with description, is_system_provided, constant_value, and is_omitted.
	DynamicVariable *string `json:"dynamic_variable,omitempty" xml:"dynamic_variable,omitempty"`
	// Server-side rejection guard for an LLM-provided value: the runtime rejects any value outside the permitted set this object names, and the set is not advertised to the LLM as an enum. Only supported when the value source is `description`; combining it with dynamic_variable, is_system_provided, constant_value, or is_omitted is rejected.
	AllowedValues *AllowedValues `json:"allowed_values,omitempty" xml:"allowed_values,omitempty"`
	// DEPRECATED: use `allowed_values` instead. When set, the LLM provides the value but the runtime rejects any value not present in the list held by this dynamic variable (must be a JSON array such as ["ws_alpha", "ws_beta"]). Use to let the LLM pick from a server-verified set (e.g. the IDs the current user is allowed to access). Requires description; mutually exclusive with dynamic_variable, is_system_provided, constant_value, and is_omitted.
	AllowedValuesDynamicVariable *string `json:"allowed_values_dynamic_variable,omitempty" xml:"allowed_values_dynamic_variable,omitempty"`
	// A constant value to use for this property. Mutually exclusive with description, dynamic_variable, is_system_provided, and is_omitted.
	ConstantValue *DataCollectionValueSchemaConstantValue `json:"constant_value,omitempty" xml:"constant_value,omitempty"`
	// If true, this parameter will be completely omitted from the request. Only valid for optional parameters. Mutually exclusive with description, dynamic_variable, is_system_provided, and constant_value.
	IsOmitted       *bool `json:"is_omitted,omitempty" xml:"is_omitted,omitempty"`
	extraProperties map[string]interface{}
}

func (d *DataCollectionValueSchema) GetType() (value *DataCollectionValueSchemaType) {
	if d == nil {
		return
	}
	return d.Type
}

func (d *DataCollectionValueSchema) GetDescription() (value *string) {
	if d == nil {
		return
	}
	return d.Description
}

func (d *DataCollectionValueSchema) GetEnum() (value []string) {
	if d == nil {
		return
	}
	return d.Enum
}

func (d *DataCollectionValueSchema) GetIsSystemProvided() (value *bool) {
	if d == nil {
		return
	}
	return d.IsSystemProvided
}

func (d *DataCollectionValueSchema) GetDynamicVariable() (value *string) {
	if d == nil {
		return
	}
	return d.DynamicVariable
}

func (d *DataCollectionValueSchema) GetAllowedValues() (value *AllowedValues) {
	if d == nil {
		return
	}
	return d.AllowedValues
}

func (d *DataCollectionValueSchema) GetAllowedValuesDynamicVariable() (value *string) {
	if d == nil {
		return
	}
	return d.AllowedValuesDynamicVariable
}

func (d *DataCollectionValueSchema) GetConstantValue() (value *DataCollectionValueSchemaConstantValue) {
	if d == nil {
		return
	}
	return d.ConstantValue
}

func (d *DataCollectionValueSchema) GetIsOmitted() (value *bool) {
	if d == nil {
		return
	}
	return d.IsOmitted
}

func (d *DataCollectionValueSchema) SetType(type_ *DataCollectionValueSchemaType) {
	d.Type = type_
}

func (d *DataCollectionValueSchema) SetDescription(description *string) {
	d.Description = description
}

func (d *DataCollectionValueSchema) SetEnum(enum []string) {
	d.Enum = enum
}

func (d *DataCollectionValueSchema) SetIsSystemProvided(isSystemProvided *bool) {
	d.IsSystemProvided = isSystemProvided
}

func (d *DataCollectionValueSchema) SetDynamicVariable(dynamicVariable *string) {
	d.DynamicVariable = dynamicVariable
}

func (d *DataCollectionValueSchema) SetAllowedValues(allowedValues *AllowedValues) {
	d.AllowedValues = allowedValues
}

func (d *DataCollectionValueSchema) SetAllowedValuesDynamicVariable(allowedValuesDynamicVariable *string) {
	d.AllowedValuesDynamicVariable = allowedValuesDynamicVariable
}

func (d *DataCollectionValueSchema) SetConstantValue(constantValue *DataCollectionValueSchemaConstantValue) {
	d.ConstantValue = constantValue
}

func (d *DataCollectionValueSchema) SetIsOmitted(isOmitted *bool) {
	d.IsOmitted = isOmitted
}

func (d *DataCollectionValueSchema) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DataCollectionValueSchema) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DataCollectionValueSchema to string"
	}
	return string(jsonData)
}

func (d *DataCollectionValueSchema) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, d); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DataCollectionValueSchema) MarshalJSON() ([]byte, error) {
	type embed DataCollectionValueSchema
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
