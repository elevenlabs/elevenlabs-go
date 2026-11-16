package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Platform limit with usage stored in Firestore. Example: credit usage tracked in Firestore.
type StoredUsagePlatformLimit struct {
	Limit *int `json:"limit,omitempty" xml:"limit,omitempty"`
	// Current usage
	Usage           *int `json:"usage,omitempty" xml:"usage,omitempty"`
	extraProperties map[string]interface{}
}

func (s *StoredUsagePlatformLimit) GetLimit() (value *int) {
	if s == nil {
		return
	}
	return s.Limit
}

func (s *StoredUsagePlatformLimit) GetUsage() (value *int) {
	if s == nil {
		return
	}
	return s.Usage
}

func (s *StoredUsagePlatformLimit) SetLimit(limit *int) {
	s.Limit = limit
}

func (s *StoredUsagePlatformLimit) SetUsage(usage *int) {
	s.Usage = usage
}

func (s *StoredUsagePlatformLimit) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s StoredUsagePlatformLimit) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: StoredUsagePlatformLimit to string"
	}
	return string(jsonData)
}

func (s *StoredUsagePlatformLimit) UnmarshalJSON(data []byte) error {
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

func (s *StoredUsagePlatformLimit) MarshalJSON() ([]byte, error) {
	type embed StoredUsagePlatformLimit
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
