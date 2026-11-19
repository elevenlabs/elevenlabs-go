package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PhoneDynamicVariableDestinationInput struct {
	// Phone number, SIP URI, or the dynamic variable name that resolves to one.
	Destination     string `json:"destination" xml:"destination" required:"true" minLength:"1"`
	extraProperties map[string]interface{}
}

func (p *PhoneDynamicVariableDestinationInput) GetDestination() (value string) {
	if p == nil {
		return
	}
	return p.Destination
}

func (p *PhoneDynamicVariableDestinationInput) SetDestination(destination string) {
	p.Destination = destination
}

func (p *PhoneDynamicVariableDestinationInput) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PhoneDynamicVariableDestinationInput) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PhoneDynamicVariableDestinationInput to string"
	}
	return string(jsonData)
}

func (p *PhoneDynamicVariableDestinationInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias PhoneDynamicVariableDestinationInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = PhoneDynamicVariableDestinationInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *p, "type")
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PhoneDynamicVariableDestinationInput) MarshalJSON() ([]byte, error) {
	type embed PhoneDynamicVariableDestinationInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
