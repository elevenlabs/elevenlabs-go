package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type SimulationLibrarySettings struct {
	extraProperties map[string]interface{}
}

func (s *SimulationLibrarySettings) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SimulationLibrarySettings) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SimulationLibrarySettings to string"
	}
	return string(jsonData)
}

func (s *SimulationLibrarySettings) UnmarshalJSON(data []byte) error {
	type unmarshaler SimulationLibrarySettings
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SimulationLibrarySettings(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SimulationLibrarySettings) MarshalJSON() ([]byte, error) {
	type embed SimulationLibrarySettings
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
