package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AstMultiplicationOperatorNodeOutput struct {
	type_           string
	Left            *AstNodeOutput `json:"left" xml:"left" required:"true"`
	Right           *AstNodeOutput `json:"right" xml:"right" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AstMultiplicationOperatorNodeOutput) GetLeft() (value *AstNodeOutput) {
	if a == nil {
		return
	}
	return a.Left
}

func (a *AstMultiplicationOperatorNodeOutput) GetRight() (value *AstNodeOutput) {
	if a == nil {
		return
	}
	return a.Right
}

func (a *AstMultiplicationOperatorNodeOutput) SetLeft(left *AstNodeOutput) {
	a.Left = left
}

func (a *AstMultiplicationOperatorNodeOutput) SetRight(right *AstNodeOutput) {
	a.Right = right
}

func (a *AstMultiplicationOperatorNodeOutput) Type() string {
	return a.type_
}

func (a *AstMultiplicationOperatorNodeOutput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AstMultiplicationOperatorNodeOutput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AstMultiplicationOperatorNodeOutput to string"
	}
	return string(jsonData)
}

func (a *AstMultiplicationOperatorNodeOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AstMultiplicationOperatorNodeOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AstMultiplicationOperatorNodeOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "mul_operator" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", a, "mul_operator", *literals.Type)
	}
	a.type_ = "mul_operator"
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "type")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AstMultiplicationOperatorNodeOutput) MarshalJSON() ([]byte, error) {
	type embed AstMultiplicationOperatorNodeOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*a),
		Type:  "mul_operator",
	}
	return json.Marshal(marshaler)
}
