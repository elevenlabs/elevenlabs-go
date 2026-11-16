package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type KnowledgeBaseRagChunkModel struct {
	ChunkId         string `json:"chunk_id" xml:"chunk_id" required:"true"`
	DocumentId      string `json:"document_id" xml:"document_id" required:"true"`
	Content         string `json:"content" xml:"content" required:"true"`
	extraProperties map[string]interface{}
}

func (k *KnowledgeBaseRagChunkModel) GetChunkId() (value string) {
	if k == nil {
		return
	}
	return k.ChunkId
}

func (k *KnowledgeBaseRagChunkModel) GetDocumentId() (value string) {
	if k == nil {
		return
	}
	return k.DocumentId
}

func (k *KnowledgeBaseRagChunkModel) GetContent() (value string) {
	if k == nil {
		return
	}
	return k.Content
}

func (k *KnowledgeBaseRagChunkModel) SetChunkId(chunkId string) {
	k.ChunkId = chunkId
}

func (k *KnowledgeBaseRagChunkModel) SetDocumentId(documentId string) {
	k.DocumentId = documentId
}

func (k *KnowledgeBaseRagChunkModel) SetContent(content string) {
	k.Content = content
}

func (k *KnowledgeBaseRagChunkModel) GetExtraProperties() map[string]interface{} {
	if k == nil {
		return nil
	}
	return k.extraProperties
}

func (k KnowledgeBaseRagChunkModel) String() string {
	jsonData, err := json.MarshalIndent(k, "", "  ")
	if err != nil {
		return "error converting struct: KnowledgeBaseRagChunkModel to string"
	}
	return string(jsonData)
}

func (k *KnowledgeBaseRagChunkModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, k); err != nil {
		return err
	}
	type alias KnowledgeBaseRagChunkModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*k = KnowledgeBaseRagChunkModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *k)
	if err != nil {
		return err
	}
	k.extraProperties = extraProperties
	return nil
}

func (k *KnowledgeBaseRagChunkModel) MarshalJSON() ([]byte, error) {
	type embed KnowledgeBaseRagChunkModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*k),
	}
	return json.Marshal(marshaler)
}
