package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ChunkingOverride struct {
	// Maximum chunk size in tokens. Null uses the inherited value.
	MaxChunkSize *int `json:"max_chunk_size" xml:"max_chunk_size" required:"true"`
	// Chunk overlap as a percentage of the chunk size. Null uses the inherited value.
	OverlapPercentage *int `json:"overlap_percentage" xml:"overlap_percentage" required:"true"`
	// Custom text separators used when splitting. Null uses the inherited value.
	CustomSeparators []string `json:"custom_separators" xml:"custom_separators" required:"true"`
	extraProperties  map[string]interface{}
}

func (c *ChunkingOverride) GetMaxChunkSize() (value *int) {
	if c == nil {
		return
	}
	return c.MaxChunkSize
}

func (c *ChunkingOverride) GetOverlapPercentage() (value *int) {
	if c == nil {
		return
	}
	return c.OverlapPercentage
}

func (c *ChunkingOverride) GetCustomSeparators() (value []string) {
	if c == nil {
		return
	}
	return c.CustomSeparators
}

func (c *ChunkingOverride) SetMaxChunkSize(maxChunkSize *int) {
	c.MaxChunkSize = maxChunkSize
}

func (c *ChunkingOverride) SetOverlapPercentage(overlapPercentage *int) {
	c.OverlapPercentage = overlapPercentage
}

func (c *ChunkingOverride) SetCustomSeparators(customSeparators []string) {
	c.CustomSeparators = customSeparators
}

func (c *ChunkingOverride) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ChunkingOverride) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ChunkingOverride to string"
	}
	return string(jsonData)
}

func (c *ChunkingOverride) UnmarshalJSON(data []byte) error {
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

func (c *ChunkingOverride) MarshalJSON() ([]byte, error) {
	type embed ChunkingOverride
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
