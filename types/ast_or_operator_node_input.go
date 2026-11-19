package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type AstOrOperatorNodeInput struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Child nodes of the logical operator.
	Children        []*AstNodeInput `json:"children" xml:"children" required:"true" minItems:"2"`
	extraProperties map[string]interface{}
}

func (a *AstOrOperatorNodeInput) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AstOrOperatorNodeInput) GetChildren() (value []*AstNodeInput) {
	if a == nil {
		return
	}
	return a.Children
}

func (a *AstOrOperatorNodeInput) SetType(type_ *string) {
	a.Type = type_
}

func (a *AstOrOperatorNodeInput) SetChildren(children []*AstNodeInput) {
	a.Children = children
}

func (a *AstOrOperatorNodeInput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AstOrOperatorNodeInput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AstOrOperatorNodeInput to string"
	}
	return string(jsonData)
}

func (a *AstOrOperatorNodeInput) UnmarshalJSON(data []byte) error {
	type unmarshaler AstOrOperatorNodeInput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AstOrOperatorNodeInput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AstOrOperatorNodeInput) MarshalJSON() ([]byte, error) {
	type embed AstOrOperatorNodeInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
