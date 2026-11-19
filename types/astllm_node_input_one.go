package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AstllmNodeInputOne struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The prompt to evaluate to a boolean value. Deprecated. Use a boolean schema instead.
	Prompt          string `json:"prompt" xml:"prompt" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AstllmNodeInputOne) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AstllmNodeInputOne) GetPrompt() (value string) {
	if a == nil {
		return
	}
	return a.Prompt
}

func (a *AstllmNodeInputOne) SetType(type_ *string) {
	a.Type = type_
}

func (a *AstllmNodeInputOne) SetPrompt(prompt string) {
	a.Prompt = prompt
}

func (a *AstllmNodeInputOne) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AstllmNodeInputOne) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AstllmNodeInputOne to string"
	}
	return string(jsonData)
}

func (a *AstllmNodeInputOne) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AstllmNodeInputOne
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AstllmNodeInputOne(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AstllmNodeInputOne) MarshalJSON() ([]byte, error) {
	type embed AstllmNodeInputOne
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
