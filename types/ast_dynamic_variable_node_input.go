package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AstDynamicVariableNodeInput struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The name of the dynamic variable.
	Name            string `json:"name" xml:"name" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AstDynamicVariableNodeInput) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AstDynamicVariableNodeInput) GetName() (value string) {
	if a == nil {
		return
	}
	return a.Name
}

func (a *AstDynamicVariableNodeInput) SetType(type_ *string) {
	a.Type = type_
}

func (a *AstDynamicVariableNodeInput) SetName(name string) {
	a.Name = name
}

func (a *AstDynamicVariableNodeInput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AstDynamicVariableNodeInput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AstDynamicVariableNodeInput to string"
	}
	return string(jsonData)
}

func (a *AstDynamicVariableNodeInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AstDynamicVariableNodeInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AstDynamicVariableNodeInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AstDynamicVariableNodeInput) MarshalJSON() ([]byte, error) {
	type embed AstDynamicVariableNodeInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
