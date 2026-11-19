package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Schema property for literal JSON types. IMPORTANT: Only ONE of the following fields can be set: description (LLM provides value), dynamic_variable (value from variable), is_system_provided (system provides value), constant_value (fixed value), or is_omitted (parameter is omitted). These are mutually exclusive.
type LiteralJsonSchemaProperty struct {
	Type *LiteralJsonSchemaPropertyType `json:"type" xml:"type" required:"true"`
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
	ConstantValue *LiteralJsonSchemaPropertyConstantValue `json:"constant_value,omitempty" xml:"constant_value,omitempty"`
	// If true, this parameter will be completely omitted from the request. Only valid for optional parameters. Mutually exclusive with description, dynamic_variable, is_system_provided, and constant_value.
	IsOmitted       *bool `json:"is_omitted,omitempty" xml:"is_omitted,omitempty"`
	extraProperties map[string]interface{}
}

func (l *LiteralJsonSchemaProperty) GetType() (value *LiteralJsonSchemaPropertyType) {
	if l == nil {
		return
	}
	return l.Type
}

func (l *LiteralJsonSchemaProperty) GetDescription() (value *string) {
	if l == nil {
		return
	}
	return l.Description
}

func (l *LiteralJsonSchemaProperty) GetEnum() (value []string) {
	if l == nil {
		return
	}
	return l.Enum
}

func (l *LiteralJsonSchemaProperty) GetIsSystemProvided() (value *bool) {
	if l == nil {
		return
	}
	return l.IsSystemProvided
}

func (l *LiteralJsonSchemaProperty) GetDynamicVariable() (value *string) {
	if l == nil {
		return
	}
	return l.DynamicVariable
}

func (l *LiteralJsonSchemaProperty) GetAllowedValues() (value *AllowedValues) {
	if l == nil {
		return
	}
	return l.AllowedValues
}

func (l *LiteralJsonSchemaProperty) GetAllowedValuesDynamicVariable() (value *string) {
	if l == nil {
		return
	}
	return l.AllowedValuesDynamicVariable
}

func (l *LiteralJsonSchemaProperty) GetConstantValue() (value *LiteralJsonSchemaPropertyConstantValue) {
	if l == nil {
		return
	}
	return l.ConstantValue
}

func (l *LiteralJsonSchemaProperty) GetIsOmitted() (value *bool) {
	if l == nil {
		return
	}
	return l.IsOmitted
}

func (l *LiteralJsonSchemaProperty) SetType(type_ *LiteralJsonSchemaPropertyType) {
	l.Type = type_
}

func (l *LiteralJsonSchemaProperty) SetDescription(description *string) {
	l.Description = description
}

func (l *LiteralJsonSchemaProperty) SetEnum(enum []string) {
	l.Enum = enum
}

func (l *LiteralJsonSchemaProperty) SetIsSystemProvided(isSystemProvided *bool) {
	l.IsSystemProvided = isSystemProvided
}

func (l *LiteralJsonSchemaProperty) SetDynamicVariable(dynamicVariable *string) {
	l.DynamicVariable = dynamicVariable
}

func (l *LiteralJsonSchemaProperty) SetAllowedValues(allowedValues *AllowedValues) {
	l.AllowedValues = allowedValues
}

func (l *LiteralJsonSchemaProperty) SetAllowedValuesDynamicVariable(allowedValuesDynamicVariable *string) {
	l.AllowedValuesDynamicVariable = allowedValuesDynamicVariable
}

func (l *LiteralJsonSchemaProperty) SetConstantValue(constantValue *LiteralJsonSchemaPropertyConstantValue) {
	l.ConstantValue = constantValue
}

func (l *LiteralJsonSchemaProperty) SetIsOmitted(isOmitted *bool) {
	l.IsOmitted = isOmitted
}

func (l *LiteralJsonSchemaProperty) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LiteralJsonSchemaProperty) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LiteralJsonSchemaProperty to string"
	}
	return string(jsonData)
}

func (l *LiteralJsonSchemaProperty) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, l); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *LiteralJsonSchemaProperty) MarshalJSON() ([]byte, error) {
	type embed LiteralJsonSchemaProperty
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
