package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AstLessThanOrEqualsOperatorNodeInput struct {
	Type            *string       `json:"type,omitempty" xml:"type,omitempty"`
	Left            *AstNodeInput `json:"left" xml:"left" required:"true"`
	Right           *AstNodeInput `json:"right" xml:"right" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AstLessThanOrEqualsOperatorNodeInput) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AstLessThanOrEqualsOperatorNodeInput) GetLeft() (value *AstNodeInput) {
	if a == nil {
		return
	}
	return a.Left
}

func (a *AstLessThanOrEqualsOperatorNodeInput) GetRight() (value *AstNodeInput) {
	if a == nil {
		return
	}
	return a.Right
}

func (a *AstLessThanOrEqualsOperatorNodeInput) SetType(type_ *string) {
	a.Type = type_
}

func (a *AstLessThanOrEqualsOperatorNodeInput) SetLeft(left *AstNodeInput) {
	a.Left = left
}

func (a *AstLessThanOrEqualsOperatorNodeInput) SetRight(right *AstNodeInput) {
	a.Right = right
}

func (a *AstLessThanOrEqualsOperatorNodeInput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AstLessThanOrEqualsOperatorNodeInput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AstLessThanOrEqualsOperatorNodeInput to string"
	}
	return string(jsonData)
}

func (a *AstLessThanOrEqualsOperatorNodeInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AstLessThanOrEqualsOperatorNodeInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AstLessThanOrEqualsOperatorNodeInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AstLessThanOrEqualsOperatorNodeInput) MarshalJSON() ([]byte, error) {
	type embed AstLessThanOrEqualsOperatorNodeInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
