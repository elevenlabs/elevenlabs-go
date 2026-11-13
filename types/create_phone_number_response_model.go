package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CreatePhoneNumberResponseModel struct {
	// Phone entity ID
	PhoneNumberId   string `json:"phone_number_id" xml:"phone_number_id" required:"true"`
	extraProperties map[string]interface{}
}

func (c *CreatePhoneNumberResponseModel) GetPhoneNumberId() (value string) {
	if c == nil {
		return
	}
	return c.PhoneNumberId
}

func (c *CreatePhoneNumberResponseModel) SetPhoneNumberId(phoneNumberId string) {
	c.PhoneNumberId = phoneNumberId
}

func (c *CreatePhoneNumberResponseModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreatePhoneNumberResponseModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreatePhoneNumberResponseModel to string"
	}
	return string(jsonData)
}

func (c *CreatePhoneNumberResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreatePhoneNumberResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreatePhoneNumberResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreatePhoneNumberResponseModel) MarshalJSON() ([]byte, error) {
	type embed CreatePhoneNumberResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
