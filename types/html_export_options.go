package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type HtmlExportOptions struct {
	IncludeSpeakers             *bool `json:"include_speakers,omitempty" xml:"include_speakers,omitempty"`
	IncludeTimestamps           *bool `json:"include_timestamps,omitempty" xml:"include_timestamps,omitempty"`
	format                      string
	SegmentOnSilenceLongerThanS *float64 `json:"segment_on_silence_longer_than_s,omitempty" xml:"segment_on_silence_longer_than_s,omitempty"`
	MaxSegmentDurationS         *float64 `json:"max_segment_duration_s,omitempty" xml:"max_segment_duration_s,omitempty"`
	MaxSegmentChars             *int     `json:"max_segment_chars,omitempty" xml:"max_segment_chars,omitempty"`
	extraProperties             map[string]interface{}
}

func (h *HtmlExportOptions) GetIncludeSpeakers() (value *bool) {
	if h == nil {
		return
	}
	return h.IncludeSpeakers
}

func (h *HtmlExportOptions) GetIncludeTimestamps() (value *bool) {
	if h == nil {
		return
	}
	return h.IncludeTimestamps
}

func (h *HtmlExportOptions) GetSegmentOnSilenceLongerThanS() (value *float64) {
	if h == nil {
		return
	}
	return h.SegmentOnSilenceLongerThanS
}

func (h *HtmlExportOptions) GetMaxSegmentDurationS() (value *float64) {
	if h == nil {
		return
	}
	return h.MaxSegmentDurationS
}

func (h *HtmlExportOptions) GetMaxSegmentChars() (value *int) {
	if h == nil {
		return
	}
	return h.MaxSegmentChars
}

func (h *HtmlExportOptions) SetIncludeSpeakers(includeSpeakers *bool) {
	h.IncludeSpeakers = includeSpeakers
}

func (h *HtmlExportOptions) SetIncludeTimestamps(includeTimestamps *bool) {
	h.IncludeTimestamps = includeTimestamps
}

func (h *HtmlExportOptions) SetSegmentOnSilenceLongerThanS(segmentOnSilenceLongerThanS *float64) {
	h.SegmentOnSilenceLongerThanS = segmentOnSilenceLongerThanS
}

func (h *HtmlExportOptions) SetMaxSegmentDurationS(maxSegmentDurationS *float64) {
	h.MaxSegmentDurationS = maxSegmentDurationS
}

func (h *HtmlExportOptions) SetMaxSegmentChars(maxSegmentChars *int) {
	h.MaxSegmentChars = maxSegmentChars
}

func (h *HtmlExportOptions) Format() string {
	return h.format
}

func (h *HtmlExportOptions) GetExtraProperties() map[string]interface{} {
	if h == nil {
		return nil
	}
	return h.extraProperties
}

func (h HtmlExportOptions) String() string {
	jsonData, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		return "error converting struct: HtmlExportOptions to string"
	}
	return string(jsonData)
}

func (h *HtmlExportOptions) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, h); err != nil {
		return err
	}
	var literals struct {
		Format *string `json:"format"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Format != nil && *literals.Format != "html" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", h, "html", *literals.Format)
	}
	h.format = "html"
	extraProperties, err := internal.ExtractExtraProperties(data, *h, "format")
	if err != nil {
		return err
	}
	h.extraProperties = extraProperties
	return nil
}

func (h *HtmlExportOptions) MarshalJSON() ([]byte, error) {
	type embed HtmlExportOptions
	var marshaler = struct {
		embed
		Format string `json:"format"`
	}{
		embed:  embed(*h),
		Format: "html",
	}
	return json.Marshal(marshaler)
}
