package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkflowStateUpdateOutput struct {
	type_ string
	// Name of the dynamic variable to assign.
	VariableName string         `json:"variable_name" xml:"variable_name" required:"true"`
	Expression   *AstNodeOutput `json:"expression" xml:"expression" required:"true"`
	// When false, the LLM may omit the values this update relies on. The update is then skipped and the variable keeps its current value.
	Required        bool `json:"required" xml:"required" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WorkflowStateUpdateOutput) GetVariableName() (value string) {
	if w == nil {
		return
	}
	return w.VariableName
}

func (w *WorkflowStateUpdateOutput) GetExpression() (value *AstNodeOutput) {
	if w == nil {
		return
	}
	return w.Expression
}

func (w *WorkflowStateUpdateOutput) GetRequired() (value bool) {
	if w == nil {
		return
	}
	return w.Required
}

func (w *WorkflowStateUpdateOutput) SetVariableName(variableName string) {
	w.VariableName = variableName
}

func (w *WorkflowStateUpdateOutput) SetExpression(expression *AstNodeOutput) {
	w.Expression = expression
}

func (w *WorkflowStateUpdateOutput) SetRequired(required bool) {
	w.Required = required
}

func (w *WorkflowStateUpdateOutput) Type() string {
	return w.type_
}

func (w *WorkflowStateUpdateOutput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowStateUpdateOutput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowStateUpdateOutput to string"
	}
	return string(jsonData)
}

func (w *WorkflowStateUpdateOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WorkflowStateUpdateOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WorkflowStateUpdateOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "dynamic_variable" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", w, "dynamic_variable", *literals.Type)
	}
	w.type_ = "dynamic_variable"
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowStateUpdateOutput) MarshalJSON() ([]byte, error) {
	type embed WorkflowStateUpdateOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*w),
		Type:  "dynamic_variable",
	}
	return json.Marshal(marshaler)
}
