package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type RagChunkMetadata struct {
	DocumentId      string  `json:"document_id" xml:"document_id" required:"true"`
	ChunkId         string  `json:"chunk_id" xml:"chunk_id" required:"true"`
	VectorDistance  float64 `json:"vector_distance" xml:"vector_distance" required:"true"`
	extraProperties map[string]interface{}
}

func (r *RagChunkMetadata) GetDocumentId() (value string) {
	if r == nil {
		return
	}
	return r.DocumentId
}

func (r *RagChunkMetadata) GetChunkId() (value string) {
	if r == nil {
		return
	}
	return r.ChunkId
}

func (r *RagChunkMetadata) GetVectorDistance() (value float64) {
	if r == nil {
		return
	}
	return r.VectorDistance
}

func (r *RagChunkMetadata) SetDocumentId(documentId string) {
	r.DocumentId = documentId
}

func (r *RagChunkMetadata) SetChunkId(chunkId string) {
	r.ChunkId = chunkId
}

func (r *RagChunkMetadata) SetVectorDistance(vectorDistance float64) {
	r.VectorDistance = vectorDistance
}

func (r *RagChunkMetadata) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r RagChunkMetadata) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RagChunkMetadata to string"
	}
	return string(jsonData)
}

func (r *RagChunkMetadata) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias RagChunkMetadata
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = RagChunkMetadata(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *RagChunkMetadata) MarshalJSON() ([]byte, error) {
	type embed RagChunkMetadata
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
