package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DocxExportOptions struct {
	IncludeSpeakers             *bool `json:"include_speakers,omitempty" xml:"include_speakers,omitempty"`
	IncludeTimestamps           *bool `json:"include_timestamps,omitempty" xml:"include_timestamps,omitempty"`
	format                      string
	SegmentOnSilenceLongerThanS *float64 `json:"segment_on_silence_longer_than_s,omitempty" xml:"segment_on_silence_longer_than_s,omitempty"`
	MaxSegmentDurationS         *float64 `json:"max_segment_duration_s,omitempty" xml:"max_segment_duration_s,omitempty"`
	MaxSegmentChars             *int     `json:"max_segment_chars,omitempty" xml:"max_segment_chars,omitempty"`
	extraProperties             map[string]interface{}
}

func (d *DocxExportOptions) GetIncludeSpeakers() (value *bool) {
	if d == nil {
		return
	}
	return d.IncludeSpeakers
}

func (d *DocxExportOptions) GetIncludeTimestamps() (value *bool) {
	if d == nil {
		return
	}
	return d.IncludeTimestamps
}

func (d *DocxExportOptions) GetSegmentOnSilenceLongerThanS() (value *float64) {
	if d == nil {
		return
	}
	return d.SegmentOnSilenceLongerThanS
}

func (d *DocxExportOptions) GetMaxSegmentDurationS() (value *float64) {
	if d == nil {
		return
	}
	return d.MaxSegmentDurationS
}

func (d *DocxExportOptions) GetMaxSegmentChars() (value *int) {
	if d == nil {
		return
	}
	return d.MaxSegmentChars
}

func (d *DocxExportOptions) SetIncludeSpeakers(includeSpeakers *bool) {
	d.IncludeSpeakers = includeSpeakers
}

func (d *DocxExportOptions) SetIncludeTimestamps(includeTimestamps *bool) {
	d.IncludeTimestamps = includeTimestamps
}

func (d *DocxExportOptions) SetSegmentOnSilenceLongerThanS(segmentOnSilenceLongerThanS *float64) {
	d.SegmentOnSilenceLongerThanS = segmentOnSilenceLongerThanS
}

func (d *DocxExportOptions) SetMaxSegmentDurationS(maxSegmentDurationS *float64) {
	d.MaxSegmentDurationS = maxSegmentDurationS
}

func (d *DocxExportOptions) SetMaxSegmentChars(maxSegmentChars *int) {
	d.MaxSegmentChars = maxSegmentChars
}

func (d *DocxExportOptions) Format() string {
	return d.format
}

func (d *DocxExportOptions) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DocxExportOptions) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DocxExportOptions to string"
	}
	return string(jsonData)
}

func (d *DocxExportOptions) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, d); err != nil {
		return err
	}
	var literals struct {
		Format *string `json:"format"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Format != nil && *literals.Format != "docx" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", d, "docx", *literals.Format)
	}
	d.format = "docx"
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "format")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DocxExportOptions) MarshalJSON() ([]byte, error) {
	type embed DocxExportOptions
	var marshaler = struct {
		embed
		Format string `json:"format"`
	}{
		embed:  embed(*d),
		Format: "docx",
	}
	return json.Marshal(marshaler)
}
