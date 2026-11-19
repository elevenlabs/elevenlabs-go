package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AstGreaterThanOperatorNodeInput struct {
	Type            *string       `json:"type,omitempty" xml:"type,omitempty"`
	Left            *AstNodeInput `json:"left" xml:"left" required:"true"`
	Right           *AstNodeInput `json:"right" xml:"right" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AstGreaterThanOperatorNodeInput) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AstGreaterThanOperatorNodeInput) GetLeft() (value *AstNodeInput) {
	if a == nil {
		return
	}
	return a.Left
}

func (a *AstGreaterThanOperatorNodeInput) GetRight() (value *AstNodeInput) {
	if a == nil {
		return
	}
	return a.Right
}

func (a *AstGreaterThanOperatorNodeInput) SetType(type_ *string) {
	a.Type = type_
}

func (a *AstGreaterThanOperatorNodeInput) SetLeft(left *AstNodeInput) {
	a.Left = left
}

func (a *AstGreaterThanOperatorNodeInput) SetRight(right *AstNodeInput) {
	a.Right = right
}

func (a *AstGreaterThanOperatorNodeInput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AstGreaterThanOperatorNodeInput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AstGreaterThanOperatorNodeInput to string"
	}
	return string(jsonData)
}

func (a *AstGreaterThanOperatorNodeInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AstGreaterThanOperatorNodeInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AstGreaterThanOperatorNodeInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AstGreaterThanOperatorNodeInput) MarshalJSON() ([]byte, error) {
	type embed AstGreaterThanOperatorNodeInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
