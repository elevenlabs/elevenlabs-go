package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DependentPhoneNumberIdentifier struct {
	PhoneNumberId   string            `json:"phone_number_id" xml:"phone_number_id" required:"true"`
	PhoneNumber     string            `json:"phone_number" xml:"phone_number" required:"true"`
	Label           string            `json:"label" xml:"label" required:"true"`
	Provider        TelephonyProvider `json:"provider" xml:"provider" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DependentPhoneNumberIdentifier) GetPhoneNumberId() (value string) {
	if d == nil {
		return
	}
	return d.PhoneNumberId
}

func (d *DependentPhoneNumberIdentifier) GetPhoneNumber() (value string) {
	if d == nil {
		return
	}
	return d.PhoneNumber
}

func (d *DependentPhoneNumberIdentifier) GetLabel() (value string) {
	if d == nil {
		return
	}
	return d.Label
}

func (d *DependentPhoneNumberIdentifier) GetProvider() (value TelephonyProvider) {
	if d == nil {
		return
	}
	return d.Provider
}

func (d *DependentPhoneNumberIdentifier) SetPhoneNumberId(phoneNumberId string) {
	d.PhoneNumberId = phoneNumberId
}

func (d *DependentPhoneNumberIdentifier) SetPhoneNumber(phoneNumber string) {
	d.PhoneNumber = phoneNumber
}

func (d *DependentPhoneNumberIdentifier) SetLabel(label string) {
	d.Label = label
}

func (d *DependentPhoneNumberIdentifier) SetProvider(provider TelephonyProvider) {
	d.Provider = provider
}

func (d *DependentPhoneNumberIdentifier) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DependentPhoneNumberIdentifier) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DependentPhoneNumberIdentifier to string"
	}
	return string(jsonData)
}

func (d *DependentPhoneNumberIdentifier) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DependentPhoneNumberIdentifier
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DependentPhoneNumberIdentifier(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DependentPhoneNumberIdentifier) MarshalJSON() ([]byte, error) {
	type embed DependentPhoneNumberIdentifier
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
