package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AstConditionalOperatorNodeInput struct {
	Type            *string       `json:"type,omitempty" xml:"type,omitempty"`
	Condition       *AstNodeInput `json:"condition" xml:"condition" required:"true"`
	TrueExpression  *AstNodeInput `json:"trueExpression" xml:"trueExpression" required:"true"`
	FalseExpression *AstNodeInput `json:"falseExpression" xml:"falseExpression" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AstConditionalOperatorNodeInput) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AstConditionalOperatorNodeInput) GetCondition() (value *AstNodeInput) {
	if a == nil {
		return
	}
	return a.Condition
}

func (a *AstConditionalOperatorNodeInput) GetTrueExpression() (value *AstNodeInput) {
	if a == nil {
		return
	}
	return a.TrueExpression
}

func (a *AstConditionalOperatorNodeInput) GetFalseExpression() (value *AstNodeInput) {
	if a == nil {
		return
	}
	return a.FalseExpression
}

func (a *AstConditionalOperatorNodeInput) SetType(type_ *string) {
	a.Type = type_
}

func (a *AstConditionalOperatorNodeInput) SetCondition(condition *AstNodeInput) {
	a.Condition = condition
}

func (a *AstConditionalOperatorNodeInput) SetTrueExpression(trueExpression *AstNodeInput) {
	a.TrueExpression = trueExpression
}

func (a *AstConditionalOperatorNodeInput) SetFalseExpression(falseExpression *AstNodeInput) {
	a.FalseExpression = falseExpression
}

func (a *AstConditionalOperatorNodeInput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AstConditionalOperatorNodeInput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AstConditionalOperatorNodeInput to string"
	}
	return string(jsonData)
}

func (a *AstConditionalOperatorNodeInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AstConditionalOperatorNodeInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AstConditionalOperatorNodeInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AstConditionalOperatorNodeInput) MarshalJSON() ([]byte, error) {
	type embed AstConditionalOperatorNodeInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
