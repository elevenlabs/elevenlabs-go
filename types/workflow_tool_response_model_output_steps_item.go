package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type WorkflowToolResponseModelOutputStepsItem struct {
	Type                  string
	Edge                  *WorkflowToolEdgeStepModel
	MaxIterationsExceeded *WorkflowToolMaxIterationsExceededStepModel
	NestedTools           *WorkflowToolNestedToolsStepModelOutput
}

func (w *WorkflowToolResponseModelOutputStepsItem) GetEdge() *WorkflowToolEdgeStepModel {
	if w == nil {
		return nil
	}
	if w.Edge == nil {
		return nil
	}
	return w.Edge
}

func (w *WorkflowToolResponseModelOutputStepsItem) GetMaxIterationsExceeded() *WorkflowToolMaxIterationsExceededStepModel {
	if w == nil {
		return nil
	}
	if w.MaxIterationsExceeded == nil {
		return nil
	}
	return w.MaxIterationsExceeded
}

func (w *WorkflowToolResponseModelOutputStepsItem) GetNestedTools() *WorkflowToolNestedToolsStepModelOutput {
	if w == nil {
		return nil
	}
	if w.NestedTools == nil {
		return nil
	}
	return w.NestedTools
}

func (w *WorkflowToolResponseModelOutputStepsItem) GetType() string {
	if w == nil {
		return ""
	}
	return w.Type
}

func (w *WorkflowToolResponseModelOutputStepsItem) UnmarshalJSON(data []byte) error {
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
	case "edge":
		value := new(WorkflowToolEdgeStepModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.Edge = value
	case "max_iterations_exceeded":
		value := new(WorkflowToolMaxIterationsExceededStepModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.MaxIterationsExceeded = value
	case "nested_tools":
		value := new(WorkflowToolNestedToolsStepModelOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.NestedTools = value
	}
	return nil
}

func (w WorkflowToolResponseModelOutputStepsItem) MarshalJSON() ([]byte, error) {
	if err := w.validate(); err != nil {
		return nil, err
	}
	if w.Edge != nil {
		return internal.MarshalJSONWithDiscriminant(w.Edge, "type", "edge")
	}
	if w.MaxIterationsExceeded != nil {
		return internal.MarshalJSONWithDiscriminant(w.MaxIterationsExceeded, "type", "max_iterations_exceeded")
	}
	if w.NestedTools != nil {
		return internal.MarshalJSONWithDiscriminant(w.NestedTools, "type", "nested_tools")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", w)
}

type WorkflowToolResponseModelOutputStepsItemVisitor interface {
	VisitEdge(*WorkflowToolEdgeStepModel) error
	VisitMaxIterationsExceeded(*WorkflowToolMaxIterationsExceededStepModel) error
	VisitNestedTools(*WorkflowToolNestedToolsStepModelOutput) error
}

func (w *WorkflowToolResponseModelOutputStepsItem) Accept(visitor WorkflowToolResponseModelOutputStepsItemVisitor) error {
	if w.Edge != nil {
		return visitor.VisitEdge(w.Edge)
	}
	if w.MaxIterationsExceeded != nil {
		return visitor.VisitMaxIterationsExceeded(w.MaxIterationsExceeded)
	}
	if w.NestedTools != nil {
		return visitor.VisitNestedTools(w.NestedTools)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", w)
}

func (w *WorkflowToolResponseModelOutputStepsItem) validate() error {
	if w == nil {
		return fmt.Errorf("type %T is nil", w)
	}
	var fields []string
	if w.Edge != nil {
		fields = append(fields, "edge")
	}
	if w.MaxIterationsExceeded != nil {
		fields = append(fields, "max_iterations_exceeded")
	}
	if w.NestedTools != nil {
		fields = append(fields, "nested_tools")
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
