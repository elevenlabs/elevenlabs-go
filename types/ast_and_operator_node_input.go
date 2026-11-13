package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type AstAndOperatorNodeInput struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Child nodes of the logical operator.
	Children        []*AstNodeInput `json:"children" xml:"children" required:"true" minItems:"2"`
	extraProperties map[string]interface{}
}

func (a *AstAndOperatorNodeInput) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AstAndOperatorNodeInput) GetChildren() (value []*AstNodeInput) {
	if a == nil {
		return
	}
	return a.Children
}

func (a *AstAndOperatorNodeInput) SetType(type_ *string) {
	a.Type = type_
}

func (a *AstAndOperatorNodeInput) SetChildren(children []*AstNodeInput) {
	a.Children = children
}

func (a *AstAndOperatorNodeInput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AstAndOperatorNodeInput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AstAndOperatorNodeInput to string"
	}
	return string(jsonData)
}

func (a *AstAndOperatorNodeInput) UnmarshalJSON(data []byte) error {
	type unmarshaler AstAndOperatorNodeInput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AstAndOperatorNodeInput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AstAndOperatorNodeInput) MarshalJSON() ([]byte, error) {
	type embed AstAndOperatorNodeInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
