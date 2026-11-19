package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AstAndOperatorNodeOutput struct {
	type_ string
	// Child nodes of the logical operator.
	Children        []*AstNodeOutput `json:"children" xml:"children" required:"true" minItems:"2"`
	extraProperties map[string]interface{}
}

func (a *AstAndOperatorNodeOutput) GetChildren() (value []*AstNodeOutput) {
	if a == nil {
		return
	}
	return a.Children
}

func (a *AstAndOperatorNodeOutput) SetChildren(children []*AstNodeOutput) {
	a.Children = children
}

func (a *AstAndOperatorNodeOutput) Type() string {
	return a.type_
}

func (a *AstAndOperatorNodeOutput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AstAndOperatorNodeOutput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AstAndOperatorNodeOutput to string"
	}
	return string(jsonData)
}

func (a *AstAndOperatorNodeOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AstAndOperatorNodeOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AstAndOperatorNodeOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "and_operator" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", a, "and_operator", *literals.Type)
	}
	a.type_ = "and_operator"
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "type")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AstAndOperatorNodeOutput) MarshalJSON() ([]byte, error) {
	type embed AstAndOperatorNodeOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*a),
		Type:  "and_operator",
	}
	return json.Marshal(marshaler)
}
