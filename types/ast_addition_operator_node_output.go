package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AstAdditionOperatorNodeOutput struct {
	type_           string
	Left            *AstNodeOutput `json:"left" xml:"left" required:"true"`
	Right           *AstNodeOutput `json:"right" xml:"right" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AstAdditionOperatorNodeOutput) GetLeft() (value *AstNodeOutput) {
	if a == nil {
		return
	}
	return a.Left
}

func (a *AstAdditionOperatorNodeOutput) GetRight() (value *AstNodeOutput) {
	if a == nil {
		return
	}
	return a.Right
}

func (a *AstAdditionOperatorNodeOutput) SetLeft(left *AstNodeOutput) {
	a.Left = left
}

func (a *AstAdditionOperatorNodeOutput) SetRight(right *AstNodeOutput) {
	a.Right = right
}

func (a *AstAdditionOperatorNodeOutput) Type() string {
	return a.type_
}

func (a *AstAdditionOperatorNodeOutput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AstAdditionOperatorNodeOutput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AstAdditionOperatorNodeOutput to string"
	}
	return string(jsonData)
}

func (a *AstAdditionOperatorNodeOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AstAdditionOperatorNodeOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AstAdditionOperatorNodeOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "add_operator" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", a, "add_operator", *literals.Type)
	}
	a.type_ = "add_operator"
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "type")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AstAdditionOperatorNodeOutput) MarshalJSON() ([]byte, error) {
	type embed AstAdditionOperatorNodeOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*a),
		Type:  "add_operator",
	}
	return json.Marshal(marshaler)
}
