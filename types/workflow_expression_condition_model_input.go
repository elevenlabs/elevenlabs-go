package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkflowExpressionConditionModelInput struct {
	// Optional human-readable label for the condition used throughout the UI.
	Label           *string       `json:"label,omitempty" xml:"label,omitempty"`
	Expression      *AstNodeInput `json:"expression" xml:"expression" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WorkflowExpressionConditionModelInput) GetLabel() (value *string) {
	if w == nil {
		return
	}
	return w.Label
}

func (w *WorkflowExpressionConditionModelInput) GetExpression() (value *AstNodeInput) {
	if w == nil {
		return
	}
	return w.Expression
}

func (w *WorkflowExpressionConditionModelInput) SetLabel(label *string) {
	w.Label = label
}

func (w *WorkflowExpressionConditionModelInput) SetExpression(expression *AstNodeInput) {
	w.Expression = expression
}

func (w *WorkflowExpressionConditionModelInput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowExpressionConditionModelInput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowExpressionConditionModelInput to string"
	}
	return string(jsonData)
}

func (w *WorkflowExpressionConditionModelInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, w); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowExpressionConditionModelInput) MarshalJSON() ([]byte, error) {
	type embed WorkflowExpressionConditionModelInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
