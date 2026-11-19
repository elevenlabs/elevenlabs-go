package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SrtExportOptions struct {
	MaxCharactersPerLine        *int  `json:"max_characters_per_line,omitempty" xml:"max_characters_per_line,omitempty"`
	IncludeSpeakers             *bool `json:"include_speakers,omitempty" xml:"include_speakers,omitempty"`
	IncludeTimestamps           *bool `json:"include_timestamps,omitempty" xml:"include_timestamps,omitempty"`
	format                      string
	SegmentOnSilenceLongerThanS *float64 `json:"segment_on_silence_longer_than_s,omitempty" xml:"segment_on_silence_longer_than_s,omitempty"`
	MaxSegmentDurationS         *float64 `json:"max_segment_duration_s,omitempty" xml:"max_segment_duration_s,omitempty"`
	MaxSegmentChars             *int     `json:"max_segment_chars,omitempty" xml:"max_segment_chars,omitempty"`
	extraProperties             map[string]interface{}
}

func (s *SrtExportOptions) GetMaxCharactersPerLine() (value *int) {
	if s == nil {
		return
	}
	return s.MaxCharactersPerLine
}

func (s *SrtExportOptions) GetIncludeSpeakers() (value *bool) {
	if s == nil {
		return
	}
	return s.IncludeSpeakers
}

func (s *SrtExportOptions) GetIncludeTimestamps() (value *bool) {
	if s == nil {
		return
	}
	return s.IncludeTimestamps
}

func (s *SrtExportOptions) GetSegmentOnSilenceLongerThanS() (value *float64) {
	if s == nil {
		return
	}
	return s.SegmentOnSilenceLongerThanS
}

func (s *SrtExportOptions) GetMaxSegmentDurationS() (value *float64) {
	if s == nil {
		return
	}
	return s.MaxSegmentDurationS
}

func (s *SrtExportOptions) GetMaxSegmentChars() (value *int) {
	if s == nil {
		return
	}
	return s.MaxSegmentChars
}

func (s *SrtExportOptions) SetMaxCharactersPerLine(maxCharactersPerLine *int) {
	s.MaxCharactersPerLine = maxCharactersPerLine
}

func (s *SrtExportOptions) SetIncludeSpeakers(includeSpeakers *bool) {
	s.IncludeSpeakers = includeSpeakers
}

func (s *SrtExportOptions) SetIncludeTimestamps(includeTimestamps *bool) {
	s.IncludeTimestamps = includeTimestamps
}

func (s *SrtExportOptions) SetSegmentOnSilenceLongerThanS(segmentOnSilenceLongerThanS *float64) {
	s.SegmentOnSilenceLongerThanS = segmentOnSilenceLongerThanS
}

func (s *SrtExportOptions) SetMaxSegmentDurationS(maxSegmentDurationS *float64) {
	s.MaxSegmentDurationS = maxSegmentDurationS
}

func (s *SrtExportOptions) SetMaxSegmentChars(maxSegmentChars *int) {
	s.MaxSegmentChars = maxSegmentChars
}

func (s *SrtExportOptions) Format() string {
	return s.format
}

func (s *SrtExportOptions) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SrtExportOptions) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SrtExportOptions to string"
	}
	return string(jsonData)
}

func (s *SrtExportOptions) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, s); err != nil {
		return err
	}
	var literals struct {
		Format *string `json:"format"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Format != nil && *literals.Format != "srt" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "srt", *literals.Format)
	}
	s.format = "srt"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "format")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SrtExportOptions) MarshalJSON() ([]byte, error) {
	type embed SrtExportOptions
	var marshaler = struct {
		embed
		Format string `json:"format"`
	}{
		embed:  embed(*s),
		Format: "srt",
	}
	return json.Marshal(marshaler)
}
