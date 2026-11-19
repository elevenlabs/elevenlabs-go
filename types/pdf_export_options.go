package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PdfExportOptions struct {
	IncludeSpeakers             *bool `json:"include_speakers,omitempty" xml:"include_speakers,omitempty"`
	IncludeTimestamps           *bool `json:"include_timestamps,omitempty" xml:"include_timestamps,omitempty"`
	format                      string
	SegmentOnSilenceLongerThanS *float64 `json:"segment_on_silence_longer_than_s,omitempty" xml:"segment_on_silence_longer_than_s,omitempty"`
	MaxSegmentDurationS         *float64 `json:"max_segment_duration_s,omitempty" xml:"max_segment_duration_s,omitempty"`
	MaxSegmentChars             *int     `json:"max_segment_chars,omitempty" xml:"max_segment_chars,omitempty"`
	extraProperties             map[string]interface{}
}

func (p *PdfExportOptions) GetIncludeSpeakers() (value *bool) {
	if p == nil {
		return
	}
	return p.IncludeSpeakers
}

func (p *PdfExportOptions) GetIncludeTimestamps() (value *bool) {
	if p == nil {
		return
	}
	return p.IncludeTimestamps
}

func (p *PdfExportOptions) GetSegmentOnSilenceLongerThanS() (value *float64) {
	if p == nil {
		return
	}
	return p.SegmentOnSilenceLongerThanS
}

func (p *PdfExportOptions) GetMaxSegmentDurationS() (value *float64) {
	if p == nil {
		return
	}
	return p.MaxSegmentDurationS
}

func (p *PdfExportOptions) GetMaxSegmentChars() (value *int) {
	if p == nil {
		return
	}
	return p.MaxSegmentChars
}

func (p *PdfExportOptions) SetIncludeSpeakers(includeSpeakers *bool) {
	p.IncludeSpeakers = includeSpeakers
}

func (p *PdfExportOptions) SetIncludeTimestamps(includeTimestamps *bool) {
	p.IncludeTimestamps = includeTimestamps
}

func (p *PdfExportOptions) SetSegmentOnSilenceLongerThanS(segmentOnSilenceLongerThanS *float64) {
	p.SegmentOnSilenceLongerThanS = segmentOnSilenceLongerThanS
}

func (p *PdfExportOptions) SetMaxSegmentDurationS(maxSegmentDurationS *float64) {
	p.MaxSegmentDurationS = maxSegmentDurationS
}

func (p *PdfExportOptions) SetMaxSegmentChars(maxSegmentChars *int) {
	p.MaxSegmentChars = maxSegmentChars
}

func (p *PdfExportOptions) Format() string {
	return p.format
}

func (p *PdfExportOptions) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PdfExportOptions) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PdfExportOptions to string"
	}
	return string(jsonData)
}

func (p *PdfExportOptions) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, p); err != nil {
		return err
	}
	var literals struct {
		Format *string `json:"format"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Format != nil && *literals.Format != "pdf" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", p, "pdf", *literals.Format)
	}
	p.format = "pdf"
	extraProperties, err := internal.ExtractExtraProperties(data, *p, "format")
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PdfExportOptions) MarshalJSON() ([]byte, error) {
	type embed PdfExportOptions
	var marshaler = struct {
		embed
		Format string `json:"format"`
	}{
		embed:  embed(*p),
		Format: "pdf",
	}
	return json.Marshal(marshaler)
}
