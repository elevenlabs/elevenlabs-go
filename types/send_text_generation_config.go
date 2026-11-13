package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type SendTextGenerationConfig struct {
	ChunkLengthSchedule []int `json:"chunk_length_schedule,omitempty" xml:"chunk_length_schedule,omitempty"`
	extraProperties     map[string]interface{}
}

func (s *SendTextGenerationConfig) GetChunkLengthSchedule() (value []int) {
	if s == nil {
		return
	}
	return s.ChunkLengthSchedule
}

func (s *SendTextGenerationConfig) SetChunkLengthSchedule(chunkLengthSchedule []int) {
	s.ChunkLengthSchedule = chunkLengthSchedule
}

func (s *SendTextGenerationConfig) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SendTextGenerationConfig) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SendTextGenerationConfig to string"
	}
	return string(jsonData)
}

func (s *SendTextGenerationConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler SendTextGenerationConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SendTextGenerationConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SendTextGenerationConfig) MarshalJSON() ([]byte, error) {
	type embed SendTextGenerationConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
