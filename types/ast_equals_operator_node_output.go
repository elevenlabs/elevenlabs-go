package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AstEqualsOperatorNodeOutput struct {
	type_           string
	Left            *AstNodeOutput `json:"left" xml:"left" required:"true"`
	Right           *AstNodeOutput `json:"right" xml:"right" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AstEqualsOperatorNodeOutput) GetLeft() (value *AstNodeOutput) {
	if a == nil {
		return
	}
	return a.Left
}

func (a *AstEqualsOperatorNodeOutput) GetRight() (value *AstNodeOutput) {
	if a == nil {
		return
	}
	return a.Right
}

func (a *AstEqualsOperatorNodeOutput) SetLeft(left *AstNodeOutput) {
	a.Left = left
}

func (a *AstEqualsOperatorNodeOutput) SetRight(right *AstNodeOutput) {
	a.Right = right
}

func (a *AstEqualsOperatorNodeOutput) Type() string {
	return a.type_
}

func (a *AstEqualsOperatorNodeOutput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AstEqualsOperatorNodeOutput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AstEqualsOperatorNodeOutput to string"
	}
	return string(jsonData)
}

func (a *AstEqualsOperatorNodeOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AstEqualsOperatorNodeOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AstEqualsOperatorNodeOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "eq_operator" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", a, "eq_operator", *literals.Type)
	}
	a.type_ = "eq_operator"
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "type")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AstEqualsOperatorNodeOutput) MarshalJSON() ([]byte, error) {
	type embed AstEqualsOperatorNodeOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*a),
		Type:  "eq_operator",
	}
	return json.Marshal(marshaler)
}
