package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type KnowledgeBaseDocumentChunksResponseModel struct {
	Chunks          []*KnowledgeBaseDocumentChunkResponseModel `json:"chunks" xml:"chunks" required:"true"`
	NextCursor      *string                                    `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	extraProperties map[string]interface{}
}

func (k *KnowledgeBaseDocumentChunksResponseModel) GetChunks() (value []*KnowledgeBaseDocumentChunkResponseModel) {
	if k == nil {
		return
	}
	return k.Chunks
}

func (k *KnowledgeBaseDocumentChunksResponseModel) GetNextCursor() (value *string) {
	if k == nil {
		return
	}
	return k.NextCursor
}

func (k *KnowledgeBaseDocumentChunksResponseModel) SetChunks(chunks []*KnowledgeBaseDocumentChunkResponseModel) {
	k.Chunks = chunks
}

func (k *KnowledgeBaseDocumentChunksResponseModel) SetNextCursor(nextCursor *string) {
	k.NextCursor = nextCursor
}

func (k *KnowledgeBaseDocumentChunksResponseModel) GetExtraProperties() map[string]interface{} {
	if k == nil {
		return nil
	}
	return k.extraProperties
}

func (k KnowledgeBaseDocumentChunksResponseModel) String() string {
	jsonData, err := json.MarshalIndent(k, "", "  ")
	if err != nil {
		return "error converting struct: KnowledgeBaseDocumentChunksResponseModel to string"
	}
	return string(jsonData)
}

func (k *KnowledgeBaseDocumentChunksResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, k); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *k)
	if err != nil {
		return err
	}
	k.extraProperties = extraProperties
	return nil
}

func (k *KnowledgeBaseDocumentChunksResponseModel) MarshalJSON() ([]byte, error) {
	type embed KnowledgeBaseDocumentChunksResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*k),
	}
	return json.Marshal(marshaler)
}
