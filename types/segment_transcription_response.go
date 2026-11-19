package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SegmentTranscriptionResponse struct {
	Version         int `json:"version" xml:"version" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SegmentTranscriptionResponse) GetVersion() (value int) {
	if s == nil {
		return
	}
	return s.Version
}

func (s *SegmentTranscriptionResponse) SetVersion(version int) {
	s.Version = version
}

func (s *SegmentTranscriptionResponse) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SegmentTranscriptionResponse) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SegmentTranscriptionResponse to string"
	}
	return string(jsonData)
}

func (s *SegmentTranscriptionResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SegmentTranscriptionResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SegmentTranscriptionResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SegmentTranscriptionResponse) MarshalJSON() ([]byte, error) {
	type embed SegmentTranscriptionResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
