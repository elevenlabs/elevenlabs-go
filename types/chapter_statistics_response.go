package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ChapterStatisticsResponse struct {
	// The number of unconverted characters.
	CharactersUnconverted int `json:"characters_unconverted" xml:"characters_unconverted" required:"true"`
	// The number of converted characters.
	CharactersConverted int `json:"characters_converted" xml:"characters_converted" required:"true"`
	// The number of converted paragraphs.
	ParagraphsConverted int `json:"paragraphs_converted" xml:"paragraphs_converted" required:"true"`
	// The number of unconverted paragraphs.
	ParagraphsUnconverted int `json:"paragraphs_unconverted" xml:"paragraphs_unconverted" required:"true"`
	// The number of credits needed to convert the remaining paragraphs.
	CreditsNeededToConvert *int `json:"credits_needed_to_convert,omitempty" xml:"credits_needed_to_convert,omitempty"`
	// Per-voice breakdown of character counts.
	VoiceStatistics []*VoiceStatisticsResponseModel `json:"voice_statistics,omitempty" xml:"voice_statistics,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ChapterStatisticsResponse) GetCharactersUnconverted() (value int) {
	if c == nil {
		return
	}
	return c.CharactersUnconverted
}

func (c *ChapterStatisticsResponse) GetCharactersConverted() (value int) {
	if c == nil {
		return
	}
	return c.CharactersConverted
}

func (c *ChapterStatisticsResponse) GetParagraphsConverted() (value int) {
	if c == nil {
		return
	}
	return c.ParagraphsConverted
}

func (c *ChapterStatisticsResponse) GetParagraphsUnconverted() (value int) {
	if c == nil {
		return
	}
	return c.ParagraphsUnconverted
}

func (c *ChapterStatisticsResponse) GetCreditsNeededToConvert() (value *int) {
	if c == nil {
		return
	}
	return c.CreditsNeededToConvert
}

func (c *ChapterStatisticsResponse) GetVoiceStatistics() (value []*VoiceStatisticsResponseModel) {
	if c == nil {
		return
	}
	return c.VoiceStatistics
}

func (c *ChapterStatisticsResponse) SetCharactersUnconverted(charactersUnconverted int) {
	c.CharactersUnconverted = charactersUnconverted
}

func (c *ChapterStatisticsResponse) SetCharactersConverted(charactersConverted int) {
	c.CharactersConverted = charactersConverted
}

func (c *ChapterStatisticsResponse) SetParagraphsConverted(paragraphsConverted int) {
	c.ParagraphsConverted = paragraphsConverted
}

func (c *ChapterStatisticsResponse) SetParagraphsUnconverted(paragraphsUnconverted int) {
	c.ParagraphsUnconverted = paragraphsUnconverted
}

func (c *ChapterStatisticsResponse) SetCreditsNeededToConvert(creditsNeededToConvert *int) {
	c.CreditsNeededToConvert = creditsNeededToConvert
}

func (c *ChapterStatisticsResponse) SetVoiceStatistics(voiceStatistics []*VoiceStatisticsResponseModel) {
	c.VoiceStatistics = voiceStatistics
}

func (c *ChapterStatisticsResponse) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ChapterStatisticsResponse) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ChapterStatisticsResponse to string"
	}
	return string(jsonData)
}

func (c *ChapterStatisticsResponse) UnmarshalJSON(data []byte) error {
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

func (c *ChapterStatisticsResponse) MarshalJSON() ([]byte, error) {
	type embed ChapterStatisticsResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
