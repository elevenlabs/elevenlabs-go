package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SegmentSubtitleFrame struct {
	StartTime       float64  `json:"start_time" xml:"start_time" required:"true"`
	EndTime         float64  `json:"end_time" xml:"end_time" required:"true"`
	Lines           []string `json:"lines" xml:"lines" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SegmentSubtitleFrame) GetStartTime() (value float64) {
	if s == nil {
		return
	}
	return s.StartTime
}

func (s *SegmentSubtitleFrame) GetEndTime() (value float64) {
	if s == nil {
		return
	}
	return s.EndTime
}

func (s *SegmentSubtitleFrame) GetLines() (value []string) {
	if s == nil {
		return
	}
	return s.Lines
}

func (s *SegmentSubtitleFrame) SetStartTime(startTime float64) {
	s.StartTime = startTime
}

func (s *SegmentSubtitleFrame) SetEndTime(endTime float64) {
	s.EndTime = endTime
}

func (s *SegmentSubtitleFrame) SetLines(lines []string) {
	s.Lines = lines
}

func (s *SegmentSubtitleFrame) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SegmentSubtitleFrame) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SegmentSubtitleFrame to string"
	}
	return string(jsonData)
}

func (s *SegmentSubtitleFrame) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SegmentSubtitleFrame
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SegmentSubtitleFrame(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SegmentSubtitleFrame) MarshalJSON() ([]byte, error) {
	type embed SegmentSubtitleFrame
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
