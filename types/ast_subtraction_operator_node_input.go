package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AstSubtractionOperatorNodeInput struct {
	Type            *string       `json:"type,omitempty" xml:"type,omitempty"`
	Left            *AstNodeInput `json:"left" xml:"left" required:"true"`
	Right           *AstNodeInput `json:"right" xml:"right" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AstSubtractionOperatorNodeInput) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AstSubtractionOperatorNodeInput) GetLeft() (value *AstNodeInput) {
	if a == nil {
		return
	}
	return a.Left
}

func (a *AstSubtractionOperatorNodeInput) GetRight() (value *AstNodeInput) {
	if a == nil {
		return
	}
	return a.Right
}

func (a *AstSubtractionOperatorNodeInput) SetType(type_ *string) {
	a.Type = type_
}

func (a *AstSubtractionOperatorNodeInput) SetLeft(left *AstNodeInput) {
	a.Left = left
}

func (a *AstSubtractionOperatorNodeInput) SetRight(right *AstNodeInput) {
	a.Right = right
}

func (a *AstSubtractionOperatorNodeInput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AstSubtractionOperatorNodeInput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AstSubtractionOperatorNodeInput to string"
	}
	return string(jsonData)
}

func (a *AstSubtractionOperatorNodeInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AstSubtractionOperatorNodeInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AstSubtractionOperatorNodeInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AstSubtractionOperatorNodeInput) MarshalJSON() ([]byte, error) {
	type embed AstSubtractionOperatorNodeInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
