package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AstLessThanOperatorNodeOutput struct {
	type_           string
	Left            *AstNodeOutput `json:"left" xml:"left" required:"true"`
	Right           *AstNodeOutput `json:"right" xml:"right" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AstLessThanOperatorNodeOutput) GetLeft() (value *AstNodeOutput) {
	if a == nil {
		return
	}
	return a.Left
}

func (a *AstLessThanOperatorNodeOutput) GetRight() (value *AstNodeOutput) {
	if a == nil {
		return
	}
	return a.Right
}

func (a *AstLessThanOperatorNodeOutput) SetLeft(left *AstNodeOutput) {
	a.Left = left
}

func (a *AstLessThanOperatorNodeOutput) SetRight(right *AstNodeOutput) {
	a.Right = right
}

func (a *AstLessThanOperatorNodeOutput) Type() string {
	return a.type_
}

func (a *AstLessThanOperatorNodeOutput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AstLessThanOperatorNodeOutput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AstLessThanOperatorNodeOutput to string"
	}
	return string(jsonData)
}

func (a *AstLessThanOperatorNodeOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AstLessThanOperatorNodeOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AstLessThanOperatorNodeOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "lt_operator" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", a, "lt_operator", *literals.Type)
	}
	a.type_ = "lt_operator"
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "type")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AstLessThanOperatorNodeOutput) MarshalJSON() ([]byte, error) {
	type embed AstLessThanOperatorNodeOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*a),
		Type:  "lt_operator",
	}
	return json.Marshal(marshaler)
}
