package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SegmentMigrationResponse struct {
	Version         int `json:"version" xml:"version" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SegmentMigrationResponse) GetVersion() (value int) {
	if s == nil {
		return
	}
	return s.Version
}

func (s *SegmentMigrationResponse) SetVersion(version int) {
	s.Version = version
}

func (s *SegmentMigrationResponse) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SegmentMigrationResponse) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SegmentMigrationResponse to string"
	}
	return string(jsonData)
}

func (s *SegmentMigrationResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SegmentMigrationResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SegmentMigrationResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SegmentMigrationResponse) MarshalJSON() ([]byte, error) {
	type embed SegmentMigrationResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
