package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AstOrOperatorNodeOutput struct {
	type_ string
	// Child nodes of the logical operator.
	Children        []*AstNodeOutput `json:"children" xml:"children" required:"true" minItems:"2"`
	extraProperties map[string]interface{}
}

func (a *AstOrOperatorNodeOutput) GetChildren() (value []*AstNodeOutput) {
	if a == nil {
		return
	}
	return a.Children
}

func (a *AstOrOperatorNodeOutput) SetChildren(children []*AstNodeOutput) {
	a.Children = children
}

func (a *AstOrOperatorNodeOutput) Type() string {
	return a.type_
}

func (a *AstOrOperatorNodeOutput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AstOrOperatorNodeOutput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AstOrOperatorNodeOutput to string"
	}
	return string(jsonData)
}

func (a *AstOrOperatorNodeOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AstOrOperatorNodeOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AstOrOperatorNodeOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "or_operator" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", a, "or_operator", *literals.Type)
	}
	a.type_ = "or_operator"
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "type")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AstOrOperatorNodeOutput) MarshalJSON() ([]byte, error) {
	type embed AstOrOperatorNodeOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*a),
		Type:  "or_operator",
	}
	return json.Marshal(marshaler)
}
