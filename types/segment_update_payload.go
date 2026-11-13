package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SegmentUpdatePayload struct {
	StartTime       *float64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
	EndTime         *float64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
	Text            *string  `json:"text,omitempty" xml:"text,omitempty"`
	extraProperties map[string]interface{}
}

func (s *SegmentUpdatePayload) GetStartTime() (value *float64) {
	if s == nil {
		return
	}
	return s.StartTime
}

func (s *SegmentUpdatePayload) GetEndTime() (value *float64) {
	if s == nil {
		return
	}
	return s.EndTime
}

func (s *SegmentUpdatePayload) GetText() (value *string) {
	if s == nil {
		return
	}
	return s.Text
}

func (s *SegmentUpdatePayload) SetStartTime(startTime *float64) {
	s.StartTime = startTime
}

func (s *SegmentUpdatePayload) SetEndTime(endTime *float64) {
	s.EndTime = endTime
}

func (s *SegmentUpdatePayload) SetText(text *string) {
	s.Text = text
}

func (s *SegmentUpdatePayload) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SegmentUpdatePayload) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SegmentUpdatePayload to string"
	}
	return string(jsonData)
}

func (s *SegmentUpdatePayload) UnmarshalJSON(data []byte) error {
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

func (s *SegmentUpdatePayload) MarshalJSON() ([]byte, error) {
	type embed SegmentUpdatePayload
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
