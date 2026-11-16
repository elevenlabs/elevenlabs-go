package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PhoneDynamicVariableDestinationOutput struct {
	// Phone number, SIP URI, or the dynamic variable name that resolves to one.
	Destination     string `json:"destination" xml:"destination" required:"true" minLength:"1"`
	type_           string
	extraProperties map[string]interface{}
}

func (p *PhoneDynamicVariableDestinationOutput) GetDestination() (value string) {
	if p == nil {
		return
	}
	return p.Destination
}

func (p *PhoneDynamicVariableDestinationOutput) SetDestination(destination string) {
	p.Destination = destination
}

func (p *PhoneDynamicVariableDestinationOutput) Type() string {
	return p.type_
}

func (p *PhoneDynamicVariableDestinationOutput) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PhoneDynamicVariableDestinationOutput) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PhoneDynamicVariableDestinationOutput to string"
	}
	return string(jsonData)
}

func (p *PhoneDynamicVariableDestinationOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias PhoneDynamicVariableDestinationOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = PhoneDynamicVariableDestinationOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "phone_dynamic_variable" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", p, "phone_dynamic_variable", *literals.Type)
	}
	p.type_ = "phone_dynamic_variable"
	extraProperties, err := internal.ExtractExtraProperties(data, *p, "type")
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PhoneDynamicVariableDestinationOutput) MarshalJSON() ([]byte, error) {
	type embed PhoneDynamicVariableDestinationOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*p),
		Type:  "phone_dynamic_variable",
	}
	return json.Marshal(marshaler)
}
