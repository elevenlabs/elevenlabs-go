package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PhoneNumberDynamicVariableTransferDestination struct {
	PhoneNumber     string `json:"phone_number" xml:"phone_number" required:"true"`
	extraProperties map[string]interface{}
}

func (p *PhoneNumberDynamicVariableTransferDestination) GetPhoneNumber() (value string) {
	if p == nil {
		return
	}
	return p.PhoneNumber
}

func (p *PhoneNumberDynamicVariableTransferDestination) SetPhoneNumber(phoneNumber string) {
	p.PhoneNumber = phoneNumber
}

func (p *PhoneNumberDynamicVariableTransferDestination) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PhoneNumberDynamicVariableTransferDestination) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PhoneNumberDynamicVariableTransferDestination to string"
	}
	return string(jsonData)
}

func (p *PhoneNumberDynamicVariableTransferDestination) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias PhoneNumberDynamicVariableTransferDestination
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = PhoneNumberDynamicVariableTransferDestination(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *p, "type")
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PhoneNumberDynamicVariableTransferDestination) MarshalJSON() ([]byte, error) {
	type embed PhoneNumberDynamicVariableTransferDestination
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
