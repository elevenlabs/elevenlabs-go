package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type CharacterAlignmentModel struct {
	Characters                 []string  `json:"characters" xml:"characters" required:"true"`
	CharacterStartTimesSeconds []float64 `json:"character_start_times_seconds" xml:"character_start_times_seconds" required:"true"`
	CharacterEndTimesSeconds   []float64 `json:"character_end_times_seconds" xml:"character_end_times_seconds" required:"true"`
	extraProperties            map[string]interface{}
}

func (c *CharacterAlignmentModel) GetCharacters() (value []string) {
	if c == nil {
		return
	}
	return c.Characters
}

func (c *CharacterAlignmentModel) GetCharacterStartTimesSeconds() (value []float64) {
	if c == nil {
		return
	}
	return c.CharacterStartTimesSeconds
}

func (c *CharacterAlignmentModel) GetCharacterEndTimesSeconds() (value []float64) {
	if c == nil {
		return
	}
	return c.CharacterEndTimesSeconds
}

func (c *CharacterAlignmentModel) SetCharacters(characters []string) {
	c.Characters = characters
}

func (c *CharacterAlignmentModel) SetCharacterStartTimesSeconds(characterStartTimesSeconds []float64) {
	c.CharacterStartTimesSeconds = characterStartTimesSeconds
}

func (c *CharacterAlignmentModel) SetCharacterEndTimesSeconds(characterEndTimesSeconds []float64) {
	c.CharacterEndTimesSeconds = characterEndTimesSeconds
}

func (c *CharacterAlignmentModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CharacterAlignmentModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CharacterAlignmentModel to string"
	}
	return string(jsonData)
}

func (c *CharacterAlignmentModel) UnmarshalJSON(data []byte) error {
	type unmarshaler CharacterAlignmentModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CharacterAlignmentModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CharacterAlignmentModel) MarshalJSON() ([]byte, error) {
	type embed CharacterAlignmentModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
