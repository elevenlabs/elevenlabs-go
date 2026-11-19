package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AstDivisionOperatorNodeInput struct {
	Type            *string       `json:"type,omitempty" xml:"type,omitempty"`
	Left            *AstNodeInput `json:"left" xml:"left" required:"true"`
	Right           *AstNodeInput `json:"right" xml:"right" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AstDivisionOperatorNodeInput) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AstDivisionOperatorNodeInput) GetLeft() (value *AstNodeInput) {
	if a == nil {
		return
	}
	return a.Left
}

func (a *AstDivisionOperatorNodeInput) GetRight() (value *AstNodeInput) {
	if a == nil {
		return
	}
	return a.Right
}

func (a *AstDivisionOperatorNodeInput) SetType(type_ *string) {
	a.Type = type_
}

func (a *AstDivisionOperatorNodeInput) SetLeft(left *AstNodeInput) {
	a.Left = left
}

func (a *AstDivisionOperatorNodeInput) SetRight(right *AstNodeInput) {
	a.Right = right
}

func (a *AstDivisionOperatorNodeInput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AstDivisionOperatorNodeInput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AstDivisionOperatorNodeInput to string"
	}
	return string(jsonData)
}

func (a *AstDivisionOperatorNodeInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AstDivisionOperatorNodeInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AstDivisionOperatorNodeInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AstDivisionOperatorNodeInput) MarshalJSON() ([]byte, error) {
	type embed AstDivisionOperatorNodeInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
