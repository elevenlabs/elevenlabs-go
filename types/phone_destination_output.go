package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PhoneDestinationOutput struct {
	// Phone number, SIP URI, or the dynamic variable name that resolves to one.
	Destination     string `json:"destination" xml:"destination" required:"true" minLength:"1"`
	type_           string
	extraProperties map[string]interface{}
}

func (p *PhoneDestinationOutput) GetDestination() (value string) {
	if p == nil {
		return
	}
	return p.Destination
}

func (p *PhoneDestinationOutput) SetDestination(destination string) {
	p.Destination = destination
}

func (p *PhoneDestinationOutput) Type() string {
	return p.type_
}

func (p *PhoneDestinationOutput) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PhoneDestinationOutput) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PhoneDestinationOutput to string"
	}
	return string(jsonData)
}

func (p *PhoneDestinationOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias PhoneDestinationOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = PhoneDestinationOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "phone" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", p, "phone", *literals.Type)
	}
	p.type_ = "phone"
	extraProperties, err := internal.ExtractExtraProperties(data, *p, "type")
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PhoneDestinationOutput) MarshalJSON() ([]byte, error) {
	type embed PhoneDestinationOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*p),
		Type:  "phone",
	}
	return json.Marshal(marshaler)
}
