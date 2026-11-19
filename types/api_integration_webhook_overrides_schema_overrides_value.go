package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ApiIntegrationWebhookOverridesSchemaOverridesValue struct {
	Source          string
	Constant        *ConstantSchemaOverride
	DynamicVariable *DynamicVariableSchemaOverride
	Llm             *LlmSchemaOverride
	Omit            *OmitSchemaOverride
}

func (a *ApiIntegrationWebhookOverridesSchemaOverridesValue) GetConstant() *ConstantSchemaOverride {
	if a == nil {
		return nil
	}
	if a.Constant == nil {
		return nil
	}
	return a.Constant
}

func (a *ApiIntegrationWebhookOverridesSchemaOverridesValue) GetDynamicVariable() *DynamicVariableSchemaOverride {
	if a == nil {
		return nil
	}
	if a.DynamicVariable == nil {
		return nil
	}
	return a.DynamicVariable
}

func (a *ApiIntegrationWebhookOverridesSchemaOverridesValue) GetLlm() *LlmSchemaOverride {
	if a == nil {
		return nil
	}
	if a.Llm == nil {
		return nil
	}
	return a.Llm
}

func (a *ApiIntegrationWebhookOverridesSchemaOverridesValue) GetOmit() *OmitSchemaOverride {
	if a == nil {
		return nil
	}
	if a.Omit == nil {
		return nil
	}
	return a.Omit
}

func (a *ApiIntegrationWebhookOverridesSchemaOverridesValue) GetSource() string {
	if a == nil {
		return ""
	}
	return a.Source
}

func (a *ApiIntegrationWebhookOverridesSchemaOverridesValue) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Source string `json:"source"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	a.Source = unmarshaler.Source
	if unmarshaler.Source == "" {
		return fmt.Errorf("%T did not include discriminant source", a)
	}
	switch unmarshaler.Source {
	case "constant":
		value := new(ConstantSchemaOverride)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.Constant = value
	case "dynamic_variable":
		value := new(DynamicVariableSchemaOverride)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.DynamicVariable = value
	case "llm":
		value := new(LlmSchemaOverride)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.Llm = value
	case "omit":
		value := new(OmitSchemaOverride)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.Omit = value
	}
	return nil
}

func (a ApiIntegrationWebhookOverridesSchemaOverridesValue) MarshalJSON() ([]byte, error) {
	if err := a.validate(); err != nil {
		return nil, err
	}
	if a.Constant != nil {
		return internal.MarshalJSONWithDiscriminant(a.Constant, "source", "constant")
	}
	if a.DynamicVariable != nil {
		return internal.MarshalJSONWithDiscriminant(a.DynamicVariable, "source", "dynamic_variable")
	}
	if a.Llm != nil {
		return internal.MarshalJSONWithDiscriminant(a.Llm, "source", "llm")
	}
	if a.Omit != nil {
		return internal.MarshalJSONWithDiscriminant(a.Omit, "source", "omit")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", a)
}

type ApiIntegrationWebhookOverridesSchemaOverridesValueVisitor interface {
	VisitConstant(*ConstantSchemaOverride) error
	VisitDynamicVariable(*DynamicVariableSchemaOverride) error
	VisitLlm(*LlmSchemaOverride) error
	VisitOmit(*OmitSchemaOverride) error
}

func (a *ApiIntegrationWebhookOverridesSchemaOverridesValue) Accept(visitor ApiIntegrationWebhookOverridesSchemaOverridesValueVisitor) error {
	if a.Constant != nil {
		return visitor.VisitConstant(a.Constant)
	}
	if a.DynamicVariable != nil {
		return visitor.VisitDynamicVariable(a.DynamicVariable)
	}
	if a.Llm != nil {
		return visitor.VisitLlm(a.Llm)
	}
	if a.Omit != nil {
		return visitor.VisitOmit(a.Omit)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", a)
}

func (a *ApiIntegrationWebhookOverridesSchemaOverridesValue) validate() error {
	if a == nil {
		return fmt.Errorf("type %T is nil", a)
	}
	var fields []string
	if a.Constant != nil {
		fields = append(fields, "constant")
	}
	if a.DynamicVariable != nil {
		fields = append(fields, "dynamic_variable")
	}
	if a.Llm != nil {
		fields = append(fields, "llm")
	}
	if a.Omit != nil {
		fields = append(fields, "omit")
	}
	if len(fields) == 0 {
		if a.Source != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", a, a.Source)
		}
		return fmt.Errorf("type %T is empty", a)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", a, fields)
	}
	if a.Source != "" {
		field := fields[0]
		if a.Source != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				a,
				a.Source,
				a,
			)
		}
	}
	return nil
}
