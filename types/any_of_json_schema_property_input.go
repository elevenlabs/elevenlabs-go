package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AnyOfJsonSchemaPropertyInput struct {
	// The object shapes the value can take.
	AnyOf []*ObjectJsonSchemaPropertyInput `json:"any_of" xml:"any_of" required:"true" minItems:"2"`
	// The property every branch declares as a required string with a single enum value; that value names the branch.
	Discriminator string `json:"discriminator" xml:"discriminator" required:"true" minLength:"1"`
	// What the value is, for the LLM.
	Description     *string `json:"description,omitempty" xml:"description,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AnyOfJsonSchemaPropertyInput) GetAnyOf() (value []*ObjectJsonSchemaPropertyInput) {
	if a == nil {
		return
	}
	return a.AnyOf
}

func (a *AnyOfJsonSchemaPropertyInput) GetDiscriminator() (value string) {
	if a == nil {
		return
	}
	return a.Discriminator
}

func (a *AnyOfJsonSchemaPropertyInput) GetDescription() (value *string) {
	if a == nil {
		return
	}
	return a.Description
}

func (a *AnyOfJsonSchemaPropertyInput) SetAnyOf(anyOf []*ObjectJsonSchemaPropertyInput) {
	a.AnyOf = anyOf
}

func (a *AnyOfJsonSchemaPropertyInput) SetDiscriminator(discriminator string) {
	a.Discriminator = discriminator
}

func (a *AnyOfJsonSchemaPropertyInput) SetDescription(description *string) {
	a.Description = description
}

func (a *AnyOfJsonSchemaPropertyInput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AnyOfJsonSchemaPropertyInput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AnyOfJsonSchemaPropertyInput to string"
	}
	return string(jsonData)
}

func (a *AnyOfJsonSchemaPropertyInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AnyOfJsonSchemaPropertyInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AnyOfJsonSchemaPropertyInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AnyOfJsonSchemaPropertyInput) MarshalJSON() ([]byte, error) {
	type embed AnyOfJsonSchemaPropertyInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
