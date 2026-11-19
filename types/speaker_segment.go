package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SpeakerSegment struct {
	Id              string                    `json:"id" xml:"id" required:"true"`
	StartTime       float64                   `json:"start_time" xml:"start_time" required:"true"`
	EndTime         float64                   `json:"end_time" xml:"end_time" required:"true"`
	Text            string                    `json:"text" xml:"text" required:"true"`
	Subtitles       []*SegmentSubtitleFrame   `json:"subtitles" xml:"subtitles" required:"true"`
	Dubs            map[string]*DubbedSegment `json:"dubs" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SpeakerSegment) GetId() (value string) {
	if s == nil {
		return
	}
	return s.Id
}

func (s *SpeakerSegment) GetStartTime() (value float64) {
	if s == nil {
		return
	}
	return s.StartTime
}

func (s *SpeakerSegment) GetEndTime() (value float64) {
	if s == nil {
		return
	}
	return s.EndTime
}

func (s *SpeakerSegment) GetText() (value string) {
	if s == nil {
		return
	}
	return s.Text
}

func (s *SpeakerSegment) GetSubtitles() (value []*SegmentSubtitleFrame) {
	if s == nil {
		return
	}
	return s.Subtitles
}

func (s *SpeakerSegment) GetDubs() (value map[string]*DubbedSegment) {
	if s == nil {
		return
	}
	return s.Dubs
}

func (s *SpeakerSegment) SetId(id string) {
	s.Id = id
}

func (s *SpeakerSegment) SetStartTime(startTime float64) {
	s.StartTime = startTime
}

func (s *SpeakerSegment) SetEndTime(endTime float64) {
	s.EndTime = endTime
}

func (s *SpeakerSegment) SetText(text string) {
	s.Text = text
}

func (s *SpeakerSegment) SetSubtitles(subtitles []*SegmentSubtitleFrame) {
	s.Subtitles = subtitles
}

func (s *SpeakerSegment) SetDubs(dubs map[string]*DubbedSegment) {
	s.Dubs = dubs
}

func (s *SpeakerSegment) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SpeakerSegment) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SpeakerSegment to string"
	}
	return string(jsonData)
}

func (s *SpeakerSegment) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SpeakerSegment
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SpeakerSegment(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SpeakerSegment) MarshalJSON() ([]byte, error) {
	type embed SpeakerSegment
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
