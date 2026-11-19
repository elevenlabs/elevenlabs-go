package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AstMultiplicationOperatorNodeInput struct {
	Type            *string       `json:"type,omitempty" xml:"type,omitempty"`
	Left            *AstNodeInput `json:"left" xml:"left" required:"true"`
	Right           *AstNodeInput `json:"right" xml:"right" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AstMultiplicationOperatorNodeInput) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AstMultiplicationOperatorNodeInput) GetLeft() (value *AstNodeInput) {
	if a == nil {
		return
	}
	return a.Left
}

func (a *AstMultiplicationOperatorNodeInput) GetRight() (value *AstNodeInput) {
	if a == nil {
		return
	}
	return a.Right
}

func (a *AstMultiplicationOperatorNodeInput) SetType(type_ *string) {
	a.Type = type_
}

func (a *AstMultiplicationOperatorNodeInput) SetLeft(left *AstNodeInput) {
	a.Left = left
}

func (a *AstMultiplicationOperatorNodeInput) SetRight(right *AstNodeInput) {
	a.Right = right
}

func (a *AstMultiplicationOperatorNodeInput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AstMultiplicationOperatorNodeInput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AstMultiplicationOperatorNodeInput to string"
	}
	return string(jsonData)
}

func (a *AstMultiplicationOperatorNodeInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AstMultiplicationOperatorNodeInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AstMultiplicationOperatorNodeInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AstMultiplicationOperatorNodeInput) MarshalJSON() ([]byte, error) {
	type embed AstMultiplicationOperatorNodeInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
