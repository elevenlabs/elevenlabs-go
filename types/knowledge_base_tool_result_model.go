package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type KnowledgeBaseToolResultModel struct {
	Status          *KnowledgeBaseToolStatus `json:"status,omitempty" xml:"status,omitempty"`
	ChunkCount      *int                     `json:"chunk_count,omitempty" xml:"chunk_count,omitempty"`
	Message         *string                  `json:"message,omitempty" xml:"message,omitempty"`
	extraProperties map[string]interface{}
}

func (k *KnowledgeBaseToolResultModel) GetStatus() (value *KnowledgeBaseToolStatus) {
	if k == nil {
		return
	}
	return k.Status
}

func (k *KnowledgeBaseToolResultModel) GetChunkCount() (value *int) {
	if k == nil {
		return
	}
	return k.ChunkCount
}

func (k *KnowledgeBaseToolResultModel) GetMessage() (value *string) {
	if k == nil {
		return
	}
	return k.Message
}

func (k *KnowledgeBaseToolResultModel) SetStatus(status *KnowledgeBaseToolStatus) {
	k.Status = status
}

func (k *KnowledgeBaseToolResultModel) SetChunkCount(chunkCount *int) {
	k.ChunkCount = chunkCount
}

func (k *KnowledgeBaseToolResultModel) SetMessage(message *string) {
	k.Message = message
}

func (k *KnowledgeBaseToolResultModel) GetExtraProperties() map[string]interface{} {
	if k == nil {
		return nil
	}
	return k.extraProperties
}

func (k KnowledgeBaseToolResultModel) String() string {
	jsonData, err := json.MarshalIndent(k, "", "  ")
	if err != nil {
		return "error converting struct: KnowledgeBaseToolResultModel to string"
	}
	return string(jsonData)
}

func (k *KnowledgeBaseToolResultModel) UnmarshalJSON(data []byte) error {
	type unmarshaler KnowledgeBaseToolResultModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*k = KnowledgeBaseToolResultModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *k, "result_type")
	if err != nil {
		return err
	}
	k.extraProperties = extraProperties
	return nil
}

func (k *KnowledgeBaseToolResultModel) MarshalJSON() ([]byte, error) {
	type embed KnowledgeBaseToolResultModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*k),
	}
	return json.Marshal(marshaler)
}
