package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetDocumentChunkingOverrideResponseModel struct {
	// ID of the document or folder whose chunking settings were read.
	Id string `json:"id" xml:"id" required:"true"`
	// Override stored directly on this entity. Null means every setting is inherited.
	ChunkingOverride          *ChunkingOverride `json:"chunking_override" xml:"chunking_override" required:"true"`
	EffectiveChunkingOverride *ChunkingOverride `json:"effective_chunking_override" xml:"effective_chunking_override" required:"true"`
	InheritedChunkingOverride *ChunkingOverride `json:"inherited_chunking_override" xml:"inherited_chunking_override" required:"true"`
	extraProperties           map[string]interface{}
}

func (g *GetDocumentChunkingOverrideResponseModel) GetId() (value string) {
	if g == nil {
		return
	}
	return g.Id
}

func (g *GetDocumentChunkingOverrideResponseModel) GetChunkingOverride() (value *ChunkingOverride) {
	if g == nil {
		return
	}
	return g.ChunkingOverride
}

func (g *GetDocumentChunkingOverrideResponseModel) GetEffectiveChunkingOverride() (value *ChunkingOverride) {
	if g == nil {
		return
	}
	return g.EffectiveChunkingOverride
}

func (g *GetDocumentChunkingOverrideResponseModel) GetInheritedChunkingOverride() (value *ChunkingOverride) {
	if g == nil {
		return
	}
	return g.InheritedChunkingOverride
}

func (g *GetDocumentChunkingOverrideResponseModel) SetId(id string) {
	g.Id = id
}

func (g *GetDocumentChunkingOverrideResponseModel) SetChunkingOverride(chunkingOverride *ChunkingOverride) {
	g.ChunkingOverride = chunkingOverride
}

func (g *GetDocumentChunkingOverrideResponseModel) SetEffectiveChunkingOverride(effectiveChunkingOverride *ChunkingOverride) {
	g.EffectiveChunkingOverride = effectiveChunkingOverride
}

func (g *GetDocumentChunkingOverrideResponseModel) SetInheritedChunkingOverride(inheritedChunkingOverride *ChunkingOverride) {
	g.InheritedChunkingOverride = inheritedChunkingOverride
}

func (g *GetDocumentChunkingOverrideResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetDocumentChunkingOverrideResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetDocumentChunkingOverrideResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetDocumentChunkingOverrideResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, g); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetDocumentChunkingOverrideResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetDocumentChunkingOverrideResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
