package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PhoneNumberTransferDestination struct {
	PhoneNumber     string `json:"phone_number" xml:"phone_number" required:"true"`
	extraProperties map[string]interface{}
}

func (p *PhoneNumberTransferDestination) GetPhoneNumber() (value string) {
	if p == nil {
		return
	}
	return p.PhoneNumber
}

func (p *PhoneNumberTransferDestination) SetPhoneNumber(phoneNumber string) {
	p.PhoneNumber = phoneNumber
}

func (p *PhoneNumberTransferDestination) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PhoneNumberTransferDestination) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PhoneNumberTransferDestination to string"
	}
	return string(jsonData)
}

func (p *PhoneNumberTransferDestination) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias PhoneNumberTransferDestination
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = PhoneNumberTransferDestination(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *p, "type")
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PhoneNumberTransferDestination) MarshalJSON() ([]byte, error) {
	type embed PhoneNumberTransferDestination
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
