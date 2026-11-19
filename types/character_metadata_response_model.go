package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CharacterMetadataResponseModel struct {
	Description                   *string          `json:"description,omitempty" xml:"description,omitempty"`
	SampleMessage                 *string          `json:"sample_message,omitempty" xml:"sample_message,omitempty"`
	VoiceCreationPromptSuggestion *string          `json:"voice_creation_prompt_suggestion,omitempty" xml:"voice_creation_prompt_suggestion,omitempty"`
	Gender                        *CharacterGender `json:"gender,omitempty" xml:"gender,omitempty"`
	Age                           *CharacterAge    `json:"age,omitempty" xml:"age,omitempty"`
	Accent                        *string          `json:"accent,omitempty" xml:"accent,omitempty"`
	Role                          *CharacterRole   `json:"role,omitempty" xml:"role,omitempty"`
	extraProperties               map[string]interface{}
}

func (c *CharacterMetadataResponseModel) GetDescription() (value *string) {
	if c == nil {
		return
	}
	return c.Description
}

func (c *CharacterMetadataResponseModel) GetSampleMessage() (value *string) {
	if c == nil {
		return
	}
	return c.SampleMessage
}

func (c *CharacterMetadataResponseModel) GetVoiceCreationPromptSuggestion() (value *string) {
	if c == nil {
		return
	}
	return c.VoiceCreationPromptSuggestion
}

func (c *CharacterMetadataResponseModel) GetGender() (value *CharacterGender) {
	if c == nil {
		return
	}
	return c.Gender
}

func (c *CharacterMetadataResponseModel) GetAge() (value *CharacterAge) {
	if c == nil {
		return
	}
	return c.Age
}

func (c *CharacterMetadataResponseModel) GetAccent() (value *string) {
	if c == nil {
		return
	}
	return c.Accent
}

func (c *CharacterMetadataResponseModel) GetRole() (value *CharacterRole) {
	if c == nil {
		return
	}
	return c.Role
}

func (c *CharacterMetadataResponseModel) SetDescription(description *string) {
	c.Description = description
}

func (c *CharacterMetadataResponseModel) SetSampleMessage(sampleMessage *string) {
	c.SampleMessage = sampleMessage
}

func (c *CharacterMetadataResponseModel) SetVoiceCreationPromptSuggestion(voiceCreationPromptSuggestion *string) {
	c.VoiceCreationPromptSuggestion = voiceCreationPromptSuggestion
}

func (c *CharacterMetadataResponseModel) SetGender(gender *CharacterGender) {
	c.Gender = gender
}

func (c *CharacterMetadataResponseModel) SetAge(age *CharacterAge) {
	c.Age = age
}

func (c *CharacterMetadataResponseModel) SetAccent(accent *string) {
	c.Accent = accent
}

func (c *CharacterMetadataResponseModel) SetRole(role *CharacterRole) {
	c.Role = role
}

func (c *CharacterMetadataResponseModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CharacterMetadataResponseModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CharacterMetadataResponseModel to string"
	}
	return string(jsonData)
}

func (c *CharacterMetadataResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CharacterMetadataResponseModel) MarshalJSON() ([]byte, error) {
	type embed CharacterMetadataResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
