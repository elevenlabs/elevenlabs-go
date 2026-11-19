package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type RagRetrievalInfo struct {
	Chunks          []*RagChunkMetadata `json:"chunks" xml:"chunks" required:"true"`
	EmbeddingModel  EmbeddingModelEnum  `json:"embedding_model" xml:"embedding_model" required:"true"`
	RetrievalQuery  string              `json:"retrieval_query" xml:"retrieval_query" required:"true"`
	RagLatencySecs  float64             `json:"rag_latency_secs" xml:"rag_latency_secs" required:"true"`
	UsedChunkIds    []string            `json:"used_chunk_ids,omitempty" xml:"used_chunk_ids,omitempty"`
	extraProperties map[string]interface{}
}

func (r *RagRetrievalInfo) GetChunks() (value []*RagChunkMetadata) {
	if r == nil {
		return
	}
	return r.Chunks
}

func (r *RagRetrievalInfo) GetEmbeddingModel() (value EmbeddingModelEnum) {
	if r == nil {
		return
	}
	return r.EmbeddingModel
}

func (r *RagRetrievalInfo) GetRetrievalQuery() (value string) {
	if r == nil {
		return
	}
	return r.RetrievalQuery
}

func (r *RagRetrievalInfo) GetRagLatencySecs() (value float64) {
	if r == nil {
		return
	}
	return r.RagLatencySecs
}

func (r *RagRetrievalInfo) GetUsedChunkIds() (value []string) {
	if r == nil {
		return
	}
	return r.UsedChunkIds
}

func (r *RagRetrievalInfo) SetChunks(chunks []*RagChunkMetadata) {
	r.Chunks = chunks
}

func (r *RagRetrievalInfo) SetEmbeddingModel(embeddingModel EmbeddingModelEnum) {
	r.EmbeddingModel = embeddingModel
}

func (r *RagRetrievalInfo) SetRetrievalQuery(retrievalQuery string) {
	r.RetrievalQuery = retrievalQuery
}

func (r *RagRetrievalInfo) SetRagLatencySecs(ragLatencySecs float64) {
	r.RagLatencySecs = ragLatencySecs
}

func (r *RagRetrievalInfo) SetUsedChunkIds(usedChunkIds []string) {
	r.UsedChunkIds = usedChunkIds
}

func (r *RagRetrievalInfo) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r RagRetrievalInfo) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RagRetrievalInfo to string"
	}
	return string(jsonData)
}

func (r *RagRetrievalInfo) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias RagRetrievalInfo
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = RagRetrievalInfo(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *RagRetrievalInfo) MarshalJSON() ([]byte, error) {
	type embed RagRetrievalInfo
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
