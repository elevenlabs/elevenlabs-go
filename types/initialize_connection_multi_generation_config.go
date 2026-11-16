package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type InitializeConnectionMultiGenerationConfig struct {
	ChunkLengthSchedule []int `json:"chunk_length_schedule,omitempty" xml:"chunk_length_schedule,omitempty"`
	extraProperties     map[string]interface{}
}

func (i *InitializeConnectionMultiGenerationConfig) GetChunkLengthSchedule() (value []int) {
	if i == nil {
		return
	}
	return i.ChunkLengthSchedule
}

func (i *InitializeConnectionMultiGenerationConfig) SetChunkLengthSchedule(chunkLengthSchedule []int) {
	i.ChunkLengthSchedule = chunkLengthSchedule
}

func (i *InitializeConnectionMultiGenerationConfig) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i InitializeConnectionMultiGenerationConfig) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InitializeConnectionMultiGenerationConfig to string"
	}
	return string(jsonData)
}

func (i *InitializeConnectionMultiGenerationConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler InitializeConnectionMultiGenerationConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*i = InitializeConnectionMultiGenerationConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	return nil
}

func (i *InitializeConnectionMultiGenerationConfig) MarshalJSON() ([]byte, error) {
	type embed InitializeConnectionMultiGenerationConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*i),
	}
	return json.Marshal(marshaler)
}
