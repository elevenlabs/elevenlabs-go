package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Data collection property with optional per-item LLM override for post-call analysis.
//
// TODO: migrate to composition (value_schema: LiteralJsonSchemaProperty + llm) instead of
// inheritance, so this generalizes cleanly to object/array schemas in the future.
type AnalysisProperty struct {
	Type AnalysisPropertyType `json:"type" xml:"type" required:"true"`
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
	ConstantValue *AnalysisPropertyConstantValue `json:"constant_value,omitempty" xml:"constant_value,omitempty"`
	// If true, this parameter will be completely omitted from the request. Only valid for optional parameters. Mutually exclusive with description, dynamic_variable, is_system_provided, and constant_value.
	IsOmitted *bool `json:"is_omitted,omitempty" xml:"is_omitted,omitempty"`
	// The name of this data collection item.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// LLM model to use for this analysis item. If not set, uses agent's analysis_llm default.
	Llm             *Llm `json:"llm,omitempty" xml:"llm,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AnalysisProperty) GetType() (value AnalysisPropertyType) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AnalysisProperty) GetDescription() (value *string) {
	if a == nil {
		return
	}
	return a.Description
}

func (a *AnalysisProperty) GetEnum() (value []string) {
	if a == nil {
		return
	}
	return a.Enum
}

func (a *AnalysisProperty) GetIsSystemProvided() (value *bool) {
	if a == nil {
		return
	}
	return a.IsSystemProvided
}

func (a *AnalysisProperty) GetDynamicVariable() (value *string) {
	if a == nil {
		return
	}
	return a.DynamicVariable
}

func (a *AnalysisProperty) GetAllowedValues() (value *AllowedValues) {
	if a == nil {
		return
	}
	return a.AllowedValues
}

func (a *AnalysisProperty) GetAllowedValuesDynamicVariable() (value *string) {
	if a == nil {
		return
	}
	return a.AllowedValuesDynamicVariable
}

func (a *AnalysisProperty) GetConstantValue() (value *AnalysisPropertyConstantValue) {
	if a == nil {
		return
	}
	return a.ConstantValue
}

func (a *AnalysisProperty) GetIsOmitted() (value *bool) {
	if a == nil {
		return
	}
	return a.IsOmitted
}

func (a *AnalysisProperty) GetName() (value *string) {
	if a == nil {
		return
	}
	return a.Name
}

func (a *AnalysisProperty) GetLlm() (value *Llm) {
	if a == nil {
		return
	}
	return a.Llm
}

func (a *AnalysisProperty) SetType(type_ AnalysisPropertyType) {
	a.Type = type_
}

func (a *AnalysisProperty) SetDescription(description *string) {
	a.Description = description
}

func (a *AnalysisProperty) SetEnum(enum []string) {
	a.Enum = enum
}

func (a *AnalysisProperty) SetIsSystemProvided(isSystemProvided *bool) {
	a.IsSystemProvided = isSystemProvided
}

func (a *AnalysisProperty) SetDynamicVariable(dynamicVariable *string) {
	a.DynamicVariable = dynamicVariable
}

func (a *AnalysisProperty) SetAllowedValues(allowedValues *AllowedValues) {
	a.AllowedValues = allowedValues
}

func (a *AnalysisProperty) SetAllowedValuesDynamicVariable(allowedValuesDynamicVariable *string) {
	a.AllowedValuesDynamicVariable = allowedValuesDynamicVariable
}

func (a *AnalysisProperty) SetConstantValue(constantValue *AnalysisPropertyConstantValue) {
	a.ConstantValue = constantValue
}

func (a *AnalysisProperty) SetIsOmitted(isOmitted *bool) {
	a.IsOmitted = isOmitted
}

func (a *AnalysisProperty) SetName(name *string) {
	a.Name = name
}

func (a *AnalysisProperty) SetLlm(llm *Llm) {
	a.Llm = llm
}

func (a *AnalysisProperty) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AnalysisProperty) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AnalysisProperty to string"
	}
	return string(jsonData)
}

func (a *AnalysisProperty) UnmarshalJSON(data []byte) error {
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

func (a *AnalysisProperty) MarshalJSON() ([]byte, error) {
	type embed AnalysisProperty
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
