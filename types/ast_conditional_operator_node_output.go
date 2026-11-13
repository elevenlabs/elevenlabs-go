package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AstConditionalOperatorNodeOutput struct {
	type_           string
	Condition       *AstNodeOutput `json:"condition" xml:"condition" required:"true"`
	TrueExpression  *AstNodeOutput `json:"trueExpression" xml:"trueExpression" required:"true"`
	FalseExpression *AstNodeOutput `json:"falseExpression" xml:"falseExpression" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AstConditionalOperatorNodeOutput) GetCondition() (value *AstNodeOutput) {
	if a == nil {
		return
	}
	return a.Condition
}

func (a *AstConditionalOperatorNodeOutput) GetTrueExpression() (value *AstNodeOutput) {
	if a == nil {
		return
	}
	return a.TrueExpression
}

func (a *AstConditionalOperatorNodeOutput) GetFalseExpression() (value *AstNodeOutput) {
	if a == nil {
		return
	}
	return a.FalseExpression
}

func (a *AstConditionalOperatorNodeOutput) SetCondition(condition *AstNodeOutput) {
	a.Condition = condition
}

func (a *AstConditionalOperatorNodeOutput) SetTrueExpression(trueExpression *AstNodeOutput) {
	a.TrueExpression = trueExpression
}

func (a *AstConditionalOperatorNodeOutput) SetFalseExpression(falseExpression *AstNodeOutput) {
	a.FalseExpression = falseExpression
}

func (a *AstConditionalOperatorNodeOutput) Type() string {
	return a.type_
}

func (a *AstConditionalOperatorNodeOutput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AstConditionalOperatorNodeOutput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AstConditionalOperatorNodeOutput to string"
	}
	return string(jsonData)
}

func (a *AstConditionalOperatorNodeOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AstConditionalOperatorNodeOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AstConditionalOperatorNodeOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "conditional_operator" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", a, "conditional_operator", *literals.Type)
	}
	a.type_ = "conditional_operator"
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "type")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AstConditionalOperatorNodeOutput) MarshalJSON() ([]byte, error) {
	type embed AstConditionalOperatorNodeOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*a),
		Type:  "conditional_operator",
	}
	return json.Marshal(marshaler)
}
