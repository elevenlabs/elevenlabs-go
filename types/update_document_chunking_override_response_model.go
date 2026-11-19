package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type UpdateDocumentChunkingOverrideResponseModel struct {
	// ID of the document or folder whose chunking override was updated.
	Id string `json:"id" xml:"id" required:"true"`
	// The resulting chunking override. Null means the entity inherits from ancestors or the embedding model default.
	ChunkingOverride *ChunkingOverride `json:"chunking_override" xml:"chunking_override" required:"true"`
	// Whether the effective configuration changed and background reindexing was queued.
	ReindexQueued   bool `json:"reindex_queued" xml:"reindex_queued" required:"true"`
	extraProperties map[string]interface{}
}

func (u *UpdateDocumentChunkingOverrideResponseModel) GetId() (value string) {
	if u == nil {
		return
	}
	return u.Id
}

func (u *UpdateDocumentChunkingOverrideResponseModel) GetChunkingOverride() (value *ChunkingOverride) {
	if u == nil {
		return
	}
	return u.ChunkingOverride
}

func (u *UpdateDocumentChunkingOverrideResponseModel) GetReindexQueued() (value bool) {
	if u == nil {
		return
	}
	return u.ReindexQueued
}

func (u *UpdateDocumentChunkingOverrideResponseModel) SetId(id string) {
	u.Id = id
}

func (u *UpdateDocumentChunkingOverrideResponseModel) SetChunkingOverride(chunkingOverride *ChunkingOverride) {
	u.ChunkingOverride = chunkingOverride
}

func (u *UpdateDocumentChunkingOverrideResponseModel) SetReindexQueued(reindexQueued bool) {
	u.ReindexQueued = reindexQueued
}

func (u *UpdateDocumentChunkingOverrideResponseModel) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UpdateDocumentChunkingOverrideResponseModel) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateDocumentChunkingOverrideResponseModel to string"
	}
	return string(jsonData)
}

func (u *UpdateDocumentChunkingOverrideResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, u); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UpdateDocumentChunkingOverrideResponseModel) MarshalJSON() ([]byte, error) {
	type embed UpdateDocumentChunkingOverrideResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
