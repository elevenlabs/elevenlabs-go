package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SegmentCreateResponse struct {
	Version         int    `json:"version" xml:"version" required:"true"`
	NewSegment      string `json:"new_segment" xml:"new_segment" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SegmentCreateResponse) GetVersion() (value int) {
	if s == nil {
		return
	}
	return s.Version
}

func (s *SegmentCreateResponse) GetNewSegment() (value string) {
	if s == nil {
		return
	}
	return s.NewSegment
}

func (s *SegmentCreateResponse) SetVersion(version int) {
	s.Version = version
}

func (s *SegmentCreateResponse) SetNewSegment(newSegment string) {
	s.NewSegment = newSegment
}

func (s *SegmentCreateResponse) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SegmentCreateResponse) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SegmentCreateResponse to string"
	}
	return string(jsonData)
}

func (s *SegmentCreateResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SegmentCreateResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SegmentCreateResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SegmentCreateResponse) MarshalJSON() ([]byte, error) {
	type embed SegmentCreateResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
