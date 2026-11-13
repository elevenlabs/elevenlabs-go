package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TxtExportOptions struct {
	MaxCharactersPerLine        *int  `json:"max_characters_per_line,omitempty" xml:"max_characters_per_line,omitempty"`
	IncludeSpeakers             *bool `json:"include_speakers,omitempty" xml:"include_speakers,omitempty"`
	IncludeTimestamps           *bool `json:"include_timestamps,omitempty" xml:"include_timestamps,omitempty"`
	format                      string
	SegmentOnSilenceLongerThanS *float64 `json:"segment_on_silence_longer_than_s,omitempty" xml:"segment_on_silence_longer_than_s,omitempty"`
	MaxSegmentDurationS         *float64 `json:"max_segment_duration_s,omitempty" xml:"max_segment_duration_s,omitempty"`
	MaxSegmentChars             *int     `json:"max_segment_chars,omitempty" xml:"max_segment_chars,omitempty"`
	extraProperties             map[string]interface{}
}

func (t *TxtExportOptions) GetMaxCharactersPerLine() (value *int) {
	if t == nil {
		return
	}
	return t.MaxCharactersPerLine
}

func (t *TxtExportOptions) GetIncludeSpeakers() (value *bool) {
	if t == nil {
		return
	}
	return t.IncludeSpeakers
}

func (t *TxtExportOptions) GetIncludeTimestamps() (value *bool) {
	if t == nil {
		return
	}
	return t.IncludeTimestamps
}

func (t *TxtExportOptions) GetSegmentOnSilenceLongerThanS() (value *float64) {
	if t == nil {
		return
	}
	return t.SegmentOnSilenceLongerThanS
}

func (t *TxtExportOptions) GetMaxSegmentDurationS() (value *float64) {
	if t == nil {
		return
	}
	return t.MaxSegmentDurationS
}

func (t *TxtExportOptions) GetMaxSegmentChars() (value *int) {
	if t == nil {
		return
	}
	return t.MaxSegmentChars
}

func (t *TxtExportOptions) SetMaxCharactersPerLine(maxCharactersPerLine *int) {
	t.MaxCharactersPerLine = maxCharactersPerLine
}

func (t *TxtExportOptions) SetIncludeSpeakers(includeSpeakers *bool) {
	t.IncludeSpeakers = includeSpeakers
}

func (t *TxtExportOptions) SetIncludeTimestamps(includeTimestamps *bool) {
	t.IncludeTimestamps = includeTimestamps
}

func (t *TxtExportOptions) SetSegmentOnSilenceLongerThanS(segmentOnSilenceLongerThanS *float64) {
	t.SegmentOnSilenceLongerThanS = segmentOnSilenceLongerThanS
}

func (t *TxtExportOptions) SetMaxSegmentDurationS(maxSegmentDurationS *float64) {
	t.MaxSegmentDurationS = maxSegmentDurationS
}

func (t *TxtExportOptions) SetMaxSegmentChars(maxSegmentChars *int) {
	t.MaxSegmentChars = maxSegmentChars
}

func (t *TxtExportOptions) Format() string {
	return t.format
}

func (t *TxtExportOptions) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TxtExportOptions) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TxtExportOptions to string"
	}
	return string(jsonData)
}

func (t *TxtExportOptions) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	var literals struct {
		Format *string `json:"format"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Format != nil && *literals.Format != "txt" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "txt", *literals.Format)
	}
	t.format = "txt"
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "format")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TxtExportOptions) MarshalJSON() ([]byte, error) {
	type embed TxtExportOptions
	var marshaler = struct {
		embed
		Format string `json:"format"`
	}{
		embed:  embed(*t),
		Format: "txt",
	}
	return json.Marshal(marshaler)
}
