package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SegmentCreatePayload struct {
	StartTime       float64            `json:"start_time" xml:"start_time" required:"true"`
	EndTime         float64            `json:"end_time" xml:"end_time" required:"true"`
	Text            *string            `json:"text,omitempty" xml:"text,omitempty"`
	Translations    map[string]*string `json:"translations,omitempty"`
	extraProperties map[string]interface{}
}

func (s *SegmentCreatePayload) GetStartTime() (value float64) {
	if s == nil {
		return
	}
	return s.StartTime
}

func (s *SegmentCreatePayload) GetEndTime() (value float64) {
	if s == nil {
		return
	}
	return s.EndTime
}

func (s *SegmentCreatePayload) GetText() (value *string) {
	if s == nil {
		return
	}
	return s.Text
}

func (s *SegmentCreatePayload) GetTranslations() (value map[string]*string) {
	if s == nil {
		return
	}
	return s.Translations
}

func (s *SegmentCreatePayload) SetStartTime(startTime float64) {
	s.StartTime = startTime
}

func (s *SegmentCreatePayload) SetEndTime(endTime float64) {
	s.EndTime = endTime
}

func (s *SegmentCreatePayload) SetText(text *string) {
	s.Text = text
}

func (s *SegmentCreatePayload) SetTranslations(translations map[string]*string) {
	s.Translations = translations
}

func (s *SegmentCreatePayload) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SegmentCreatePayload) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SegmentCreatePayload to string"
	}
	return string(jsonData)
}

func (s *SegmentCreatePayload) UnmarshalJSON(data []byte) error {
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

func (s *SegmentCreatePayload) MarshalJSON() ([]byte, error) {
	type embed SegmentCreatePayload
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
