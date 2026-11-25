package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AnyOfJsonSchemaPropertyOutput struct {
	// The object shapes the value can take.
	AnyOf []*ObjectJsonSchemaPropertyOutput `json:"any_of" xml:"any_of" required:"true" minItems:"2"`
	// The property every branch declares as a required string with a single enum value; that value names the branch.
	Discriminator string `json:"discriminator" xml:"discriminator" required:"true" minLength:"1"`
	// What the value is, for the LLM.
	Description     *string `json:"description,omitempty" xml:"description,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AnyOfJsonSchemaPropertyOutput) GetAnyOf() (value []*ObjectJsonSchemaPropertyOutput) {
	if a == nil {
		return
	}
	return a.AnyOf
}

func (a *AnyOfJsonSchemaPropertyOutput) GetDiscriminator() (value string) {
	if a == nil {
		return
	}
	return a.Discriminator
}

func (a *AnyOfJsonSchemaPropertyOutput) GetDescription() (value *string) {
	if a == nil {
		return
	}
	return a.Description
}

func (a *AnyOfJsonSchemaPropertyOutput) SetAnyOf(anyOf []*ObjectJsonSchemaPropertyOutput) {
	a.AnyOf = anyOf
}

func (a *AnyOfJsonSchemaPropertyOutput) SetDiscriminator(discriminator string) {
	a.Discriminator = discriminator
}

func (a *AnyOfJsonSchemaPropertyOutput) SetDescription(description *string) {
	a.Description = description
}

func (a *AnyOfJsonSchemaPropertyOutput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AnyOfJsonSchemaPropertyOutput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AnyOfJsonSchemaPropertyOutput to string"
	}
	return string(jsonData)
}

func (a *AnyOfJsonSchemaPropertyOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AnyOfJsonSchemaPropertyOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AnyOfJsonSchemaPropertyOutput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AnyOfJsonSchemaPropertyOutput) MarshalJSON() ([]byte, error) {
	type embed AnyOfJsonSchemaPropertyOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
