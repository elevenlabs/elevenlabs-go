package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SegmentedJsonExportOptions struct {
	IncludeSpeakers             *bool `json:"include_speakers,omitempty" xml:"include_speakers,omitempty"`
	IncludeTimestamps           *bool `json:"include_timestamps,omitempty" xml:"include_timestamps,omitempty"`
	format                      string
	SegmentOnSilenceLongerThanS *float64 `json:"segment_on_silence_longer_than_s,omitempty" xml:"segment_on_silence_longer_than_s,omitempty"`
	MaxSegmentDurationS         *float64 `json:"max_segment_duration_s,omitempty" xml:"max_segment_duration_s,omitempty"`
	MaxSegmentChars             *int     `json:"max_segment_chars,omitempty" xml:"max_segment_chars,omitempty"`
	extraProperties             map[string]interface{}
}

func (s *SegmentedJsonExportOptions) GetIncludeSpeakers() (value *bool) {
	if s == nil {
		return
	}
	return s.IncludeSpeakers
}

func (s *SegmentedJsonExportOptions) GetIncludeTimestamps() (value *bool) {
	if s == nil {
		return
	}
	return s.IncludeTimestamps
}

func (s *SegmentedJsonExportOptions) GetSegmentOnSilenceLongerThanS() (value *float64) {
	if s == nil {
		return
	}
	return s.SegmentOnSilenceLongerThanS
}

func (s *SegmentedJsonExportOptions) GetMaxSegmentDurationS() (value *float64) {
	if s == nil {
		return
	}
	return s.MaxSegmentDurationS
}

func (s *SegmentedJsonExportOptions) GetMaxSegmentChars() (value *int) {
	if s == nil {
		return
	}
	return s.MaxSegmentChars
}

func (s *SegmentedJsonExportOptions) SetIncludeSpeakers(includeSpeakers *bool) {
	s.IncludeSpeakers = includeSpeakers
}

func (s *SegmentedJsonExportOptions) SetIncludeTimestamps(includeTimestamps *bool) {
	s.IncludeTimestamps = includeTimestamps
}

func (s *SegmentedJsonExportOptions) SetSegmentOnSilenceLongerThanS(segmentOnSilenceLongerThanS *float64) {
	s.SegmentOnSilenceLongerThanS = segmentOnSilenceLongerThanS
}

func (s *SegmentedJsonExportOptions) SetMaxSegmentDurationS(maxSegmentDurationS *float64) {
	s.MaxSegmentDurationS = maxSegmentDurationS
}

func (s *SegmentedJsonExportOptions) SetMaxSegmentChars(maxSegmentChars *int) {
	s.MaxSegmentChars = maxSegmentChars
}

func (s *SegmentedJsonExportOptions) Format() string {
	return s.format
}

func (s *SegmentedJsonExportOptions) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SegmentedJsonExportOptions) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SegmentedJsonExportOptions to string"
	}
	return string(jsonData)
}

func (s *SegmentedJsonExportOptions) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, s); err != nil {
		return err
	}
	var literals struct {
		Format *string `json:"format"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Format != nil && *literals.Format != "segmented_json" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "segmented_json", *literals.Format)
	}
	s.format = "segmented_json"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "format")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SegmentedJsonExportOptions) MarshalJSON() ([]byte, error) {
	type embed SegmentedJsonExportOptions
	var marshaler = struct {
		embed
		Format string `json:"format"`
	}{
		embed:  embed(*s),
		Format: "segmented_json",
	}
	return json.Marshal(marshaler)
}
