package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type McpToolConfigOverrideOutputInputOverridesValue struct {
	Source          string
	Constant        *ConstantSchemaOverride
	DynamicVariable *DynamicVariableSchemaOverride
	Llm             *LlmSchemaOverride
	Omit            *OmitSchemaOverride
}

func (m *McpToolConfigOverrideOutputInputOverridesValue) GetConstant() *ConstantSchemaOverride {
	if m == nil {
		return nil
	}
	if m.Constant == nil {
		return nil
	}
	return m.Constant
}

func (m *McpToolConfigOverrideOutputInputOverridesValue) GetDynamicVariable() *DynamicVariableSchemaOverride {
	if m == nil {
		return nil
	}
	if m.DynamicVariable == nil {
		return nil
	}
	return m.DynamicVariable
}

func (m *McpToolConfigOverrideOutputInputOverridesValue) GetLlm() *LlmSchemaOverride {
	if m == nil {
		return nil
	}
	if m.Llm == nil {
		return nil
	}
	return m.Llm
}

func (m *McpToolConfigOverrideOutputInputOverridesValue) GetOmit() *OmitSchemaOverride {
	if m == nil {
		return nil
	}
	if m.Omit == nil {
		return nil
	}
	return m.Omit
}

func (m *McpToolConfigOverrideOutputInputOverridesValue) GetSource() string {
	if m == nil {
		return ""
	}
	return m.Source
}

func (m *McpToolConfigOverrideOutputInputOverridesValue) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Source string `json:"source"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	m.Source = unmarshaler.Source
	if unmarshaler.Source == "" {
		return fmt.Errorf("%T did not include discriminant source", m)
	}
	switch unmarshaler.Source {
	case "constant":
		value := new(ConstantSchemaOverride)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		m.Constant = value
	case "dynamic_variable":
		value := new(DynamicVariableSchemaOverride)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		m.DynamicVariable = value
	case "llm":
		value := new(LlmSchemaOverride)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		m.Llm = value
	case "omit":
		value := new(OmitSchemaOverride)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		m.Omit = value
	}
	return nil
}

func (m McpToolConfigOverrideOutputInputOverridesValue) MarshalJSON() ([]byte, error) {
	if err := m.validate(); err != nil {
		return nil, err
	}
	if m.Constant != nil {
		return internal.MarshalJSONWithDiscriminant(m.Constant, "source", "constant")
	}
	if m.DynamicVariable != nil {
		return internal.MarshalJSONWithDiscriminant(m.DynamicVariable, "source", "dynamic_variable")
	}
	if m.Llm != nil {
		return internal.MarshalJSONWithDiscriminant(m.Llm, "source", "llm")
	}
	if m.Omit != nil {
		return internal.MarshalJSONWithDiscriminant(m.Omit, "source", "omit")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", m)
}

type McpToolConfigOverrideOutputInputOverridesValueVisitor interface {
	VisitConstant(*ConstantSchemaOverride) error
	VisitDynamicVariable(*DynamicVariableSchemaOverride) error
	VisitLlm(*LlmSchemaOverride) error
	VisitOmit(*OmitSchemaOverride) error
}

func (m *McpToolConfigOverrideOutputInputOverridesValue) Accept(visitor McpToolConfigOverrideOutputInputOverridesValueVisitor) error {
	if m.Constant != nil {
		return visitor.VisitConstant(m.Constant)
	}
	if m.DynamicVariable != nil {
		return visitor.VisitDynamicVariable(m.DynamicVariable)
	}
	if m.Llm != nil {
		return visitor.VisitLlm(m.Llm)
	}
	if m.Omit != nil {
		return visitor.VisitOmit(m.Omit)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", m)
}

func (m *McpToolConfigOverrideOutputInputOverridesValue) validate() error {
	if m == nil {
		return fmt.Errorf("type %T is nil", m)
	}
	var fields []string
	if m.Constant != nil {
		fields = append(fields, "constant")
	}
	if m.DynamicVariable != nil {
		fields = append(fields, "dynamic_variable")
	}
	if m.Llm != nil {
		fields = append(fields, "llm")
	}
	if m.Omit != nil {
		fields = append(fields, "omit")
	}
	if len(fields) == 0 {
		if m.Source != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", m, m.Source)
		}
		return fmt.Errorf("type %T is empty", m)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", m, fields)
	}
	if m.Source != "" {
		field := fields[0]
		if m.Source != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				m,
				m.Source,
				m,
			)
		}
	}
	return nil
}
