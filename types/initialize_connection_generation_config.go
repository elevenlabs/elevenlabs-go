package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type InitializeConnectionGenerationConfig struct {
	ChunkLengthSchedule []int `json:"chunk_length_schedule,omitempty" xml:"chunk_length_schedule,omitempty"`
	extraProperties     map[string]interface{}
}

func (i *InitializeConnectionGenerationConfig) GetChunkLengthSchedule() (value []int) {
	if i == nil {
		return
	}
	return i.ChunkLengthSchedule
}

func (i *InitializeConnectionGenerationConfig) SetChunkLengthSchedule(chunkLengthSchedule []int) {
	i.ChunkLengthSchedule = chunkLengthSchedule
}

func (i *InitializeConnectionGenerationConfig) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i InitializeConnectionGenerationConfig) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InitializeConnectionGenerationConfig to string"
	}
	return string(jsonData)
}

func (i *InitializeConnectionGenerationConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler InitializeConnectionGenerationConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*i = InitializeConnectionGenerationConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	return nil
}

func (i *InitializeConnectionGenerationConfig) MarshalJSON() ([]byte, error) {
	type embed InitializeConnectionGenerationConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*i),
	}
	return json.Marshal(marshaler)
}
