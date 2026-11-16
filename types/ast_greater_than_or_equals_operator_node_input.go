package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AstGreaterThanOrEqualsOperatorNodeInput struct {
	Type            *string       `json:"type,omitempty" xml:"type,omitempty"`
	Left            *AstNodeInput `json:"left" xml:"left" required:"true"`
	Right           *AstNodeInput `json:"right" xml:"right" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AstGreaterThanOrEqualsOperatorNodeInput) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AstGreaterThanOrEqualsOperatorNodeInput) GetLeft() (value *AstNodeInput) {
	if a == nil {
		return
	}
	return a.Left
}

func (a *AstGreaterThanOrEqualsOperatorNodeInput) GetRight() (value *AstNodeInput) {
	if a == nil {
		return
	}
	return a.Right
}

func (a *AstGreaterThanOrEqualsOperatorNodeInput) SetType(type_ *string) {
	a.Type = type_
}

func (a *AstGreaterThanOrEqualsOperatorNodeInput) SetLeft(left *AstNodeInput) {
	a.Left = left
}

func (a *AstGreaterThanOrEqualsOperatorNodeInput) SetRight(right *AstNodeInput) {
	a.Right = right
}

func (a *AstGreaterThanOrEqualsOperatorNodeInput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AstGreaterThanOrEqualsOperatorNodeInput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AstGreaterThanOrEqualsOperatorNodeInput to string"
	}
	return string(jsonData)
}

func (a *AstGreaterThanOrEqualsOperatorNodeInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AstGreaterThanOrEqualsOperatorNodeInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AstGreaterThanOrEqualsOperatorNodeInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AstGreaterThanOrEqualsOperatorNodeInput) MarshalJSON() ([]byte, error) {
	type embed AstGreaterThanOrEqualsOperatorNodeInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
