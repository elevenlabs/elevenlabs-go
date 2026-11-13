package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AstLessThanOrEqualsOperatorNodeOutput struct {
	type_           string
	Left            *AstNodeOutput `json:"left" xml:"left" required:"true"`
	Right           *AstNodeOutput `json:"right" xml:"right" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AstLessThanOrEqualsOperatorNodeOutput) GetLeft() (value *AstNodeOutput) {
	if a == nil {
		return
	}
	return a.Left
}

func (a *AstLessThanOrEqualsOperatorNodeOutput) GetRight() (value *AstNodeOutput) {
	if a == nil {
		return
	}
	return a.Right
}

func (a *AstLessThanOrEqualsOperatorNodeOutput) SetLeft(left *AstNodeOutput) {
	a.Left = left
}

func (a *AstLessThanOrEqualsOperatorNodeOutput) SetRight(right *AstNodeOutput) {
	a.Right = right
}

func (a *AstLessThanOrEqualsOperatorNodeOutput) Type() string {
	return a.type_
}

func (a *AstLessThanOrEqualsOperatorNodeOutput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AstLessThanOrEqualsOperatorNodeOutput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AstLessThanOrEqualsOperatorNodeOutput to string"
	}
	return string(jsonData)
}

func (a *AstLessThanOrEqualsOperatorNodeOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AstLessThanOrEqualsOperatorNodeOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AstLessThanOrEqualsOperatorNodeOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "lte_operator" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", a, "lte_operator", *literals.Type)
	}
	a.type_ = "lte_operator"
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "type")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AstLessThanOrEqualsOperatorNodeOutput) MarshalJSON() ([]byte, error) {
	type embed AstLessThanOrEqualsOperatorNodeOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*a),
		Type:  "lte_operator",
	}
	return json.Marshal(marshaler)
}
