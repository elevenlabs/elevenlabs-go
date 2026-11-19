package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkflowStateUpdateInput struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Name of the dynamic variable to assign.
	VariableName string        `json:"variable_name" xml:"variable_name" required:"true"`
	Expression   *AstNodeInput `json:"expression" xml:"expression" required:"true"`
	// When false, the LLM may omit the values this update relies on. The update is then skipped and the variable keeps its current value.
	Required        *bool `json:"required,omitempty" xml:"required,omitempty"`
	extraProperties map[string]interface{}
}

func (w *WorkflowStateUpdateInput) GetType() (value *string) {
	if w == nil {
		return
	}
	return w.Type
}

func (w *WorkflowStateUpdateInput) GetVariableName() (value string) {
	if w == nil {
		return
	}
	return w.VariableName
}

func (w *WorkflowStateUpdateInput) GetExpression() (value *AstNodeInput) {
	if w == nil {
		return
	}
	return w.Expression
}

func (w *WorkflowStateUpdateInput) GetRequired() (value *bool) {
	if w == nil {
		return
	}
	return w.Required
}

func (w *WorkflowStateUpdateInput) SetType(type_ *string) {
	w.Type = type_
}

func (w *WorkflowStateUpdateInput) SetVariableName(variableName string) {
	w.VariableName = variableName
}

func (w *WorkflowStateUpdateInput) SetExpression(expression *AstNodeInput) {
	w.Expression = expression
}

func (w *WorkflowStateUpdateInput) SetRequired(required *bool) {
	w.Required = required
}

func (w *WorkflowStateUpdateInput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowStateUpdateInput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowStateUpdateInput to string"
	}
	return string(jsonData)
}

func (w *WorkflowStateUpdateInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WorkflowStateUpdateInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WorkflowStateUpdateInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowStateUpdateInput) MarshalJSON() ([]byte, error) {
	type embed WorkflowStateUpdateInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
