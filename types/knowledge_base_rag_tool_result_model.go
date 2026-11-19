package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type KnowledgeBaseRagToolResultModel struct {
	Status *KnowledgeBaseRagToolStatus `json:"status,omitempty" xml:"status,omitempty"`
	// Number of relevant chunks retrieved
	ChunkCount *int `json:"chunk_count,omitempty" xml:"chunk_count,omitempty"`
	// Human-readable status for the LLM about the search results
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Retrieved chunks; populated only in the rag-result-in-tool-result mode
	Chunks          []*KnowledgeBaseRagChunkModel `json:"chunks,omitempty" xml:"chunks,omitempty"`
	extraProperties map[string]interface{}
}

func (k *KnowledgeBaseRagToolResultModel) GetStatus() (value *KnowledgeBaseRagToolStatus) {
	if k == nil {
		return
	}
	return k.Status
}

func (k *KnowledgeBaseRagToolResultModel) GetChunkCount() (value *int) {
	if k == nil {
		return
	}
	return k.ChunkCount
}

func (k *KnowledgeBaseRagToolResultModel) GetMessage() (value *string) {
	if k == nil {
		return
	}
	return k.Message
}

func (k *KnowledgeBaseRagToolResultModel) GetChunks() (value []*KnowledgeBaseRagChunkModel) {
	if k == nil {
		return
	}
	return k.Chunks
}

func (k *KnowledgeBaseRagToolResultModel) SetStatus(status *KnowledgeBaseRagToolStatus) {
	k.Status = status
}

func (k *KnowledgeBaseRagToolResultModel) SetChunkCount(chunkCount *int) {
	k.ChunkCount = chunkCount
}

func (k *KnowledgeBaseRagToolResultModel) SetMessage(message *string) {
	k.Message = message
}

func (k *KnowledgeBaseRagToolResultModel) SetChunks(chunks []*KnowledgeBaseRagChunkModel) {
	k.Chunks = chunks
}

func (k *KnowledgeBaseRagToolResultModel) GetExtraProperties() map[string]interface{} {
	if k == nil {
		return nil
	}
	return k.extraProperties
}

func (k KnowledgeBaseRagToolResultModel) String() string {
	jsonData, err := json.MarshalIndent(k, "", "  ")
	if err != nil {
		return "error converting struct: KnowledgeBaseRagToolResultModel to string"
	}
	return string(jsonData)
}

func (k *KnowledgeBaseRagToolResultModel) UnmarshalJSON(data []byte) error {
	type unmarshaler KnowledgeBaseRagToolResultModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*k = KnowledgeBaseRagToolResultModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *k, "result_type")
	if err != nil {
		return err
	}
	k.extraProperties = extraProperties
	return nil
}

func (k *KnowledgeBaseRagToolResultModel) MarshalJSON() ([]byte, error) {
	type embed KnowledgeBaseRagToolResultModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*k),
	}
	return json.Marshal(marshaler)
}
