package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AstEqualsOperatorNodeInput struct {
	Type            *string       `json:"type,omitempty" xml:"type,omitempty"`
	Left            *AstNodeInput `json:"left" xml:"left" required:"true"`
	Right           *AstNodeInput `json:"right" xml:"right" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AstEqualsOperatorNodeInput) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AstEqualsOperatorNodeInput) GetLeft() (value *AstNodeInput) {
	if a == nil {
		return
	}
	return a.Left
}

func (a *AstEqualsOperatorNodeInput) GetRight() (value *AstNodeInput) {
	if a == nil {
		return
	}
	return a.Right
}

func (a *AstEqualsOperatorNodeInput) SetType(type_ *string) {
	a.Type = type_
}

func (a *AstEqualsOperatorNodeInput) SetLeft(left *AstNodeInput) {
	a.Left = left
}

func (a *AstEqualsOperatorNodeInput) SetRight(right *AstNodeInput) {
	a.Right = right
}

func (a *AstEqualsOperatorNodeInput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AstEqualsOperatorNodeInput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AstEqualsOperatorNodeInput to string"
	}
	return string(jsonData)
}

func (a *AstEqualsOperatorNodeInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AstEqualsOperatorNodeInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AstEqualsOperatorNodeInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AstEqualsOperatorNodeInput) MarshalJSON() ([]byte, error) {
	type embed AstEqualsOperatorNodeInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
