package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AstDynamicVariableNodeOutput struct {
	type_ string
	// The name of the dynamic variable.
	Name            string `json:"name" xml:"name" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AstDynamicVariableNodeOutput) GetName() (value string) {
	if a == nil {
		return
	}
	return a.Name
}

func (a *AstDynamicVariableNodeOutput) SetName(name string) {
	a.Name = name
}

func (a *AstDynamicVariableNodeOutput) Type() string {
	return a.type_
}

func (a *AstDynamicVariableNodeOutput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AstDynamicVariableNodeOutput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AstDynamicVariableNodeOutput to string"
	}
	return string(jsonData)
}

func (a *AstDynamicVariableNodeOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AstDynamicVariableNodeOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AstDynamicVariableNodeOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "dynamic_variable" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", a, "dynamic_variable", *literals.Type)
	}
	a.type_ = "dynamic_variable"
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "type")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AstDynamicVariableNodeOutput) MarshalJSON() ([]byte, error) {
	type embed AstDynamicVariableNodeOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*a),
		Type:  "dynamic_variable",
	}
	return json.Marshal(marshaler)
}
