package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkflowExpressionConditionModelOutput struct {
	// Optional human-readable label for the condition used throughout the UI.
	Label           *string `json:"label" xml:"label" required:"true"`
	type_           string
	Expression      *AstNodeOutput `json:"expression" xml:"expression" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WorkflowExpressionConditionModelOutput) GetLabel() (value *string) {
	if w == nil {
		return
	}
	return w.Label
}

func (w *WorkflowExpressionConditionModelOutput) GetExpression() (value *AstNodeOutput) {
	if w == nil {
		return
	}
	return w.Expression
}

func (w *WorkflowExpressionConditionModelOutput) SetLabel(label *string) {
	w.Label = label
}

func (w *WorkflowExpressionConditionModelOutput) SetExpression(expression *AstNodeOutput) {
	w.Expression = expression
}

func (w *WorkflowExpressionConditionModelOutput) Type() string {
	return w.type_
}

func (w *WorkflowExpressionConditionModelOutput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowExpressionConditionModelOutput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowExpressionConditionModelOutput to string"
	}
	return string(jsonData)
}

func (w *WorkflowExpressionConditionModelOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, w); err != nil {
		return err
	}
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "expression" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", w, "expression", *literals.Type)
	}
	w.type_ = "expression"
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowExpressionConditionModelOutput) MarshalJSON() ([]byte, error) {
	type embed WorkflowExpressionConditionModelOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*w),
		Type:  "expression",
	}
	return json.Marshal(marshaler)
}
