package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SongSection struct {
	// The name of the section. Must be between 1 and 100 characters.
	SectionName string `json:"section_name" xml:"section_name" required:"true" maxLength:"100" minLength:"1"`
	// The styles and musical directions that should be present in this section. Use English language for best result.
	PositiveLocalStyles []string `json:"positive_local_styles" xml:"positive_local_styles" required:"true" maxItems:"50"`
	// The styles and musical directions that should not be present in this section. Use English language for best result.
	NegativeLocalStyles []string `json:"negative_local_styles" xml:"negative_local_styles" required:"true" maxItems:"50"`
	// The duration of the section in milliseconds. Must be between 3000ms and 120000ms.
	DurationMs int `json:"duration_ms" xml:"duration_ms" required:"true" min:"3000" max:"120000"`
	// The lyrics of the section. Max 30 lines per section and max 200 characters per line.
	Lines []string `json:"lines" xml:"lines" required:"true" maxItems:"30"`
	// Optional source to extract the section from. Used for inpainting.
	SourceFrom      *SectionSource `json:"source_from,omitempty" xml:"source_from,omitempty"`
	extraProperties map[string]interface{}
}

func (s *SongSection) GetSectionName() (value string) {
	if s == nil {
		return
	}
	return s.SectionName
}

func (s *SongSection) GetPositiveLocalStyles() (value []string) {
	if s == nil {
		return
	}
	return s.PositiveLocalStyles
}

func (s *SongSection) GetNegativeLocalStyles() (value []string) {
	if s == nil {
		return
	}
	return s.NegativeLocalStyles
}

func (s *SongSection) GetDurationMs() (value int) {
	if s == nil {
		return
	}
	return s.DurationMs
}

func (s *SongSection) GetLines() (value []string) {
	if s == nil {
		return
	}
	return s.Lines
}

func (s *SongSection) GetSourceFrom() (value *SectionSource) {
	if s == nil {
		return
	}
	return s.SourceFrom
}

func (s *SongSection) SetSectionName(sectionName string) {
	s.SectionName = sectionName
}

func (s *SongSection) SetPositiveLocalStyles(positiveLocalStyles []string) {
	s.PositiveLocalStyles = positiveLocalStyles
}

func (s *SongSection) SetNegativeLocalStyles(negativeLocalStyles []string) {
	s.NegativeLocalStyles = negativeLocalStyles
}

func (s *SongSection) SetDurationMs(durationMs int) {
	s.DurationMs = durationMs
}

func (s *SongSection) SetLines(lines []string) {
	s.Lines = lines
}

func (s *SongSection) SetSourceFrom(sourceFrom *SectionSource) {
	s.SourceFrom = sourceFrom
}

func (s *SongSection) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SongSection) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SongSection to string"
	}
	return string(jsonData)
}

func (s *SongSection) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, s); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SongSection) MarshalJSON() ([]byte, error) {
	type embed SongSection
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
