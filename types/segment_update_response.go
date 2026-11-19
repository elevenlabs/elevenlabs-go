package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SegmentUpdateResponse struct {
	Version         int `json:"version" xml:"version" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SegmentUpdateResponse) GetVersion() (value int) {
	if s == nil {
		return
	}
	return s.Version
}

func (s *SegmentUpdateResponse) SetVersion(version int) {
	s.Version = version
}

func (s *SegmentUpdateResponse) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SegmentUpdateResponse) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SegmentUpdateResponse to string"
	}
	return string(jsonData)
}

func (s *SegmentUpdateResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SegmentUpdateResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SegmentUpdateResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SegmentUpdateResponse) MarshalJSON() ([]byte, error) {
	type embed SegmentUpdateResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
