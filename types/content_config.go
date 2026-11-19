package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ContentConfig struct {
	Sexual                     *ContentThresholdGuardrail `json:"sexual,omitempty" xml:"sexual,omitempty"`
	Violence                   *ContentThresholdGuardrail `json:"violence,omitempty" xml:"violence,omitempty"`
	Harassment                 *ContentThresholdGuardrail `json:"harassment,omitempty" xml:"harassment,omitempty"`
	SelfHarm                   *ContentThresholdGuardrail `json:"self_harm,omitempty" xml:"self_harm,omitempty"`
	Profanity                  *ContentThresholdGuardrail `json:"profanity,omitempty" xml:"profanity,omitempty"`
	ReligionOrPolitics         *ContentThresholdGuardrail `json:"religion_or_politics,omitempty" xml:"religion_or_politics,omitempty"`
	MedicalAndLegalInformation *ContentThresholdGuardrail `json:"medical_and_legal_information,omitempty" xml:"medical_and_legal_information,omitempty"`
	extraProperties            map[string]interface{}
}

func (c *ContentConfig) GetSexual() (value *ContentThresholdGuardrail) {
	if c == nil {
		return
	}
	return c.Sexual
}

func (c *ContentConfig) GetViolence() (value *ContentThresholdGuardrail) {
	if c == nil {
		return
	}
	return c.Violence
}

func (c *ContentConfig) GetHarassment() (value *ContentThresholdGuardrail) {
	if c == nil {
		return
	}
	return c.Harassment
}

func (c *ContentConfig) GetSelfHarm() (value *ContentThresholdGuardrail) {
	if c == nil {
		return
	}
	return c.SelfHarm
}

func (c *ContentConfig) GetProfanity() (value *ContentThresholdGuardrail) {
	if c == nil {
		return
	}
	return c.Profanity
}

func (c *ContentConfig) GetReligionOrPolitics() (value *ContentThresholdGuardrail) {
	if c == nil {
		return
	}
	return c.ReligionOrPolitics
}

func (c *ContentConfig) GetMedicalAndLegalInformation() (value *ContentThresholdGuardrail) {
	if c == nil {
		return
	}
	return c.MedicalAndLegalInformation
}

func (c *ContentConfig) SetSexual(sexual *ContentThresholdGuardrail) {
	c.Sexual = sexual
}

func (c *ContentConfig) SetViolence(violence *ContentThresholdGuardrail) {
	c.Violence = violence
}

func (c *ContentConfig) SetHarassment(harassment *ContentThresholdGuardrail) {
	c.Harassment = harassment
}

func (c *ContentConfig) SetSelfHarm(selfHarm *ContentThresholdGuardrail) {
	c.SelfHarm = selfHarm
}

func (c *ContentConfig) SetProfanity(profanity *ContentThresholdGuardrail) {
	c.Profanity = profanity
}

func (c *ContentConfig) SetReligionOrPolitics(religionOrPolitics *ContentThresholdGuardrail) {
	c.ReligionOrPolitics = religionOrPolitics
}

func (c *ContentConfig) SetMedicalAndLegalInformation(medicalAndLegalInformation *ContentThresholdGuardrail) {
	c.MedicalAndLegalInformation = medicalAndLegalInformation
}

func (c *ContentConfig) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ContentConfig) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContentConfig to string"
	}
	return string(jsonData)
}

func (c *ContentConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler ContentConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ContentConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ContentConfig) MarshalJSON() ([]byte, error) {
	type embed ContentConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
