package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type CharacterAlignmentResponseModel struct {
	Characters                 []string  `json:"characters" xml:"characters" required:"true"`
	CharacterStartTimesSeconds []float64 `json:"character_start_times_seconds" xml:"character_start_times_seconds" required:"true"`
	CharacterEndTimesSeconds   []float64 `json:"character_end_times_seconds" xml:"character_end_times_seconds" required:"true"`
	extraProperties            map[string]interface{}
}

func (c *CharacterAlignmentResponseModel) GetCharacters() (value []string) {
	if c == nil {
		return
	}
	return c.Characters
}

func (c *CharacterAlignmentResponseModel) GetCharacterStartTimesSeconds() (value []float64) {
	if c == nil {
		return
	}
	return c.CharacterStartTimesSeconds
}

func (c *CharacterAlignmentResponseModel) GetCharacterEndTimesSeconds() (value []float64) {
	if c == nil {
		return
	}
	return c.CharacterEndTimesSeconds
}

func (c *CharacterAlignmentResponseModel) SetCharacters(characters []string) {
	c.Characters = characters
}

func (c *CharacterAlignmentResponseModel) SetCharacterStartTimesSeconds(characterStartTimesSeconds []float64) {
	c.CharacterStartTimesSeconds = characterStartTimesSeconds
}

func (c *CharacterAlignmentResponseModel) SetCharacterEndTimesSeconds(characterEndTimesSeconds []float64) {
	c.CharacterEndTimesSeconds = characterEndTimesSeconds
}

func (c *CharacterAlignmentResponseModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CharacterAlignmentResponseModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CharacterAlignmentResponseModel to string"
	}
	return string(jsonData)
}

func (c *CharacterAlignmentResponseModel) UnmarshalJSON(data []byte) error {
	type unmarshaler CharacterAlignmentResponseModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CharacterAlignmentResponseModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CharacterAlignmentResponseModel) MarshalJSON() ([]byte, error) {
	type embed CharacterAlignmentResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
