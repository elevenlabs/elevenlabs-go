package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ChunkingOverrideRequestModel struct {
	// Maximum chunk size in tokens. Null inherits from ancestors or the embedding model default.
	MaxChunkSize *int `json:"max_chunk_size,omitempty" xml:"max_chunk_size,omitempty" min:"200"`
	// Chunk overlap as a percentage of the final chunk size, applied after the chunk size is clamped to the embedding model. Null inherits from ancestors or the embedding model default.
	OverlapPercentage *int `json:"overlap_percentage,omitempty" xml:"overlap_percentage,omitempty" min:"0" max:"25"`
	// Custom text separators used when splitting. Null inherits from ancestors; an empty list clears inherited separators.
	CustomSeparators []string `json:"custom_separators,omitempty" xml:"custom_separators,omitempty" maxItems:"5"`
	extraProperties  map[string]interface{}
}

func (c *ChunkingOverrideRequestModel) GetMaxChunkSize() (value *int) {
	if c == nil {
		return
	}
	return c.MaxChunkSize
}

func (c *ChunkingOverrideRequestModel) GetOverlapPercentage() (value *int) {
	if c == nil {
		return
	}
	return c.OverlapPercentage
}

func (c *ChunkingOverrideRequestModel) GetCustomSeparators() (value []string) {
	if c == nil {
		return
	}
	return c.CustomSeparators
}

func (c *ChunkingOverrideRequestModel) SetMaxChunkSize(maxChunkSize *int) {
	c.MaxChunkSize = maxChunkSize
}

func (c *ChunkingOverrideRequestModel) SetOverlapPercentage(overlapPercentage *int) {
	c.OverlapPercentage = overlapPercentage
}

func (c *ChunkingOverrideRequestModel) SetCustomSeparators(customSeparators []string) {
	c.CustomSeparators = customSeparators
}

func (c *ChunkingOverrideRequestModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ChunkingOverrideRequestModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ChunkingOverrideRequestModel to string"
	}
	return string(jsonData)
}

func (c *ChunkingOverrideRequestModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ChunkingOverrideRequestModel) MarshalJSON() ([]byte, error) {
	type embed ChunkingOverrideRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
