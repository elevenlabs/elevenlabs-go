package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type WorkflowEdgeModelOutputBackwardCondition struct {
	Type          string
	Expression    *WorkflowExpressionConditionModelOutput
	Llm           *WorkflowLlmConditionModelOutput
	Result        *WorkflowResultConditionModelOutput
	Unconditional *WorkflowUnconditionalModelOutput
}

func (w *WorkflowEdgeModelOutputBackwardCondition) GetExpression() *WorkflowExpressionConditionModelOutput {
	if w == nil {
		return nil
	}
	if w.Expression == nil {
		return nil
	}
	return w.Expression
}

func (w *WorkflowEdgeModelOutputBackwardCondition) GetLlm() *WorkflowLlmConditionModelOutput {
	if w == nil {
		return nil
	}
	if w.Llm == nil {
		return nil
	}
	return w.Llm
}

func (w *WorkflowEdgeModelOutputBackwardCondition) GetResult() *WorkflowResultConditionModelOutput {
	if w == nil {
		return nil
	}
	if w.Result == nil {
		return nil
	}
	return w.Result
}

func (w *WorkflowEdgeModelOutputBackwardCondition) GetUnconditional() *WorkflowUnconditionalModelOutput {
	if w == nil {
		return nil
	}
	if w.Unconditional == nil {
		return nil
	}
	return w.Unconditional
}

func (w *WorkflowEdgeModelOutputBackwardCondition) GetType() string {
	if w == nil {
		return ""
	}
	return w.Type
}

func (w *WorkflowEdgeModelOutputBackwardCondition) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	w.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", w)
	}
	switch unmarshaler.Type {
	case "expression":
		value := new(WorkflowExpressionConditionModelOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.Expression = value
	case "llm":
		value := new(WorkflowLlmConditionModelOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.Llm = value
	case "result":
		value := new(WorkflowResultConditionModelOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.Result = value
	case "unconditional":
		value := new(WorkflowUnconditionalModelOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.Unconditional = value
	}
	return nil
}

func (w WorkflowEdgeModelOutputBackwardCondition) MarshalJSON() ([]byte, error) {
	if err := w.validate(); err != nil {
		return nil, err
	}
	if w.Expression != nil {
		return internal.MarshalJSONWithDiscriminant(w.Expression, "type", "expression")
	}
	if w.Llm != nil {
		return internal.MarshalJSONWithDiscriminant(w.Llm, "type", "llm")
	}
	if w.Result != nil {
		return internal.MarshalJSONWithDiscriminant(w.Result, "type", "result")
	}
	if w.Unconditional != nil {
		return internal.MarshalJSONWithDiscriminant(w.Unconditional, "type", "unconditional")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", w)
}

type WorkflowEdgeModelOutputBackwardConditionVisitor interface {
	VisitExpression(*WorkflowExpressionConditionModelOutput) error
	VisitLlm(*WorkflowLlmConditionModelOutput) error
	VisitResult(*WorkflowResultConditionModelOutput) error
	VisitUnconditional(*WorkflowUnconditionalModelOutput) error
}

func (w *WorkflowEdgeModelOutputBackwardCondition) Accept(visitor WorkflowEdgeModelOutputBackwardConditionVisitor) error {
	if w.Expression != nil {
		return visitor.VisitExpression(w.Expression)
	}
	if w.Llm != nil {
		return visitor.VisitLlm(w.Llm)
	}
	if w.Result != nil {
		return visitor.VisitResult(w.Result)
	}
	if w.Unconditional != nil {
		return visitor.VisitUnconditional(w.Unconditional)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", w)
}

func (w *WorkflowEdgeModelOutputBackwardCondition) validate() error {
	if w == nil {
		return fmt.Errorf("type %T is nil", w)
	}
	var fields []string
	if w.Expression != nil {
		fields = append(fields, "expression")
	}
	if w.Llm != nil {
		fields = append(fields, "llm")
	}
	if w.Result != nil {
		fields = append(fields, "result")
	}
	if w.Unconditional != nil {
		fields = append(fields, "unconditional")
	}
	if len(fields) == 0 {
		if w.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", w, w.Type)
		}
		return fmt.Errorf("type %T is empty", w)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", w, fields)
	}
	if w.Type != "" {
		field := fields[0]
		if w.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				w,
				w.Type,
				w,
			)
		}
	}
	return nil
}
