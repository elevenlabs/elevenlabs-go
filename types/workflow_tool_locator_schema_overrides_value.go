package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type WorkflowToolLocatorSchemaOverridesValue struct {
	Source          string
	Constant        *ConstantSchemaOverride
	DynamicVariable *DynamicVariableSchemaOverride
	Llm             *LlmSchemaOverride
	Omit            *OmitSchemaOverride
}

func (w *WorkflowToolLocatorSchemaOverridesValue) GetConstant() *ConstantSchemaOverride {
	if w == nil {
		return nil
	}
	if w.Constant == nil {
		return nil
	}
	return w.Constant
}

func (w *WorkflowToolLocatorSchemaOverridesValue) GetDynamicVariable() *DynamicVariableSchemaOverride {
	if w == nil {
		return nil
	}
	if w.DynamicVariable == nil {
		return nil
	}
	return w.DynamicVariable
}

func (w *WorkflowToolLocatorSchemaOverridesValue) GetLlm() *LlmSchemaOverride {
	if w == nil {
		return nil
	}
	if w.Llm == nil {
		return nil
	}
	return w.Llm
}

func (w *WorkflowToolLocatorSchemaOverridesValue) GetOmit() *OmitSchemaOverride {
	if w == nil {
		return nil
	}
	if w.Omit == nil {
		return nil
	}
	return w.Omit
}

func (w *WorkflowToolLocatorSchemaOverridesValue) GetSource() string {
	if w == nil {
		return ""
	}
	return w.Source
}

func (w *WorkflowToolLocatorSchemaOverridesValue) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Source string `json:"source"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	w.Source = unmarshaler.Source
	if unmarshaler.Source == "" {
		return fmt.Errorf("%T did not include discriminant source", w)
	}
	switch unmarshaler.Source {
	case "constant":
		value := new(ConstantSchemaOverride)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.Constant = value
	case "dynamic_variable":
		value := new(DynamicVariableSchemaOverride)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.DynamicVariable = value
	case "llm":
		value := new(LlmSchemaOverride)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.Llm = value
	case "omit":
		value := new(OmitSchemaOverride)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.Omit = value
	}
	return nil
}

func (w WorkflowToolLocatorSchemaOverridesValue) MarshalJSON() ([]byte, error) {
	if err := w.validate(); err != nil {
		return nil, err
	}
	if w.Constant != nil {
		return internal.MarshalJSONWithDiscriminant(w.Constant, "source", "constant")
	}
	if w.DynamicVariable != nil {
		return internal.MarshalJSONWithDiscriminant(w.DynamicVariable, "source", "dynamic_variable")
	}
	if w.Llm != nil {
		return internal.MarshalJSONWithDiscriminant(w.Llm, "source", "llm")
	}
	if w.Omit != nil {
		return internal.MarshalJSONWithDiscriminant(w.Omit, "source", "omit")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", w)
}

type WorkflowToolLocatorSchemaOverridesValueVisitor interface {
	VisitConstant(*ConstantSchemaOverride) error
	VisitDynamicVariable(*DynamicVariableSchemaOverride) error
	VisitLlm(*LlmSchemaOverride) error
	VisitOmit(*OmitSchemaOverride) error
}

func (w *WorkflowToolLocatorSchemaOverridesValue) Accept(visitor WorkflowToolLocatorSchemaOverridesValueVisitor) error {
	if w.Constant != nil {
		return visitor.VisitConstant(w.Constant)
	}
	if w.DynamicVariable != nil {
		return visitor.VisitDynamicVariable(w.DynamicVariable)
	}
	if w.Llm != nil {
		return visitor.VisitLlm(w.Llm)
	}
	if w.Omit != nil {
		return visitor.VisitOmit(w.Omit)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", w)
}

func (w *WorkflowToolLocatorSchemaOverridesValue) validate() error {
	if w == nil {
		return fmt.Errorf("type %T is nil", w)
	}
	var fields []string
	if w.Constant != nil {
		fields = append(fields, "constant")
	}
	if w.DynamicVariable != nil {
		fields = append(fields, "dynamic_variable")
	}
	if w.Llm != nil {
		fields = append(fields, "llm")
	}
	if w.Omit != nil {
		fields = append(fields, "omit")
	}
	if len(fields) == 0 {
		if w.Source != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", w, w.Source)
		}
		return fmt.Errorf("type %T is empty", w)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", w, fields)
	}
	if w.Source != "" {
		field := fields[0]
		if w.Source != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				w,
				w.Source,
				w,
			)
		}
	}
	return nil
}
