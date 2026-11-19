package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SegmentDeleteResponse struct {
	Version         int `json:"version" xml:"version" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SegmentDeleteResponse) GetVersion() (value int) {
	if s == nil {
		return
	}
	return s.Version
}

func (s *SegmentDeleteResponse) SetVersion(version int) {
	s.Version = version
}

func (s *SegmentDeleteResponse) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SegmentDeleteResponse) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SegmentDeleteResponse to string"
	}
	return string(jsonData)
}

func (s *SegmentDeleteResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SegmentDeleteResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SegmentDeleteResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SegmentDeleteResponse) MarshalJSON() ([]byte, error) {
	type embed SegmentDeleteResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
