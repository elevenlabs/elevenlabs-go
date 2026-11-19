package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PhoneDestinationInput struct {
	// Phone number, SIP URI, or the dynamic variable name that resolves to one.
	Destination     string `json:"destination" xml:"destination" required:"true" minLength:"1"`
	extraProperties map[string]interface{}
}

func (p *PhoneDestinationInput) GetDestination() (value string) {
	if p == nil {
		return
	}
	return p.Destination
}

func (p *PhoneDestinationInput) SetDestination(destination string) {
	p.Destination = destination
}

func (p *PhoneDestinationInput) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PhoneDestinationInput) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PhoneDestinationInput to string"
	}
	return string(jsonData)
}

func (p *PhoneDestinationInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias PhoneDestinationInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = PhoneDestinationInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *p, "type")
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PhoneDestinationInput) MarshalJSON() ([]byte, error) {
	type embed PhoneDestinationInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
