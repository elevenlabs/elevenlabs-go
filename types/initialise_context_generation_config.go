package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type InitialiseContextGenerationConfig struct {
	ChunkLengthSchedule []int `json:"chunk_length_schedule,omitempty" xml:"chunk_length_schedule,omitempty"`
	extraProperties     map[string]interface{}
}

func (i *InitialiseContextGenerationConfig) GetChunkLengthSchedule() (value []int) {
	if i == nil {
		return
	}
	return i.ChunkLengthSchedule
}

func (i *InitialiseContextGenerationConfig) SetChunkLengthSchedule(chunkLengthSchedule []int) {
	i.ChunkLengthSchedule = chunkLengthSchedule
}

func (i *InitialiseContextGenerationConfig) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i InitialiseContextGenerationConfig) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InitialiseContextGenerationConfig to string"
	}
	return string(jsonData)
}

func (i *InitialiseContextGenerationConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler InitialiseContextGenerationConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*i = InitialiseContextGenerationConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	return nil
}

func (i *InitialiseContextGenerationConfig) MarshalJSON() ([]byte, error) {
	type embed InitialiseContextGenerationConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*i),
	}
	return json.Marshal(marshaler)
}
