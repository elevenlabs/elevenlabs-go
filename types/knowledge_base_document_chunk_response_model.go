package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type KnowledgeBaseDocumentChunkResponseModel struct {
	Id              string `json:"id" xml:"id" required:"true"`
	Name            string `json:"name" xml:"name" required:"true"`
	Content         string `json:"content" xml:"content" required:"true"`
	extraProperties map[string]interface{}
}

func (k *KnowledgeBaseDocumentChunkResponseModel) GetId() (value string) {
	if k == nil {
		return
	}
	return k.Id
}

func (k *KnowledgeBaseDocumentChunkResponseModel) GetName() (value string) {
	if k == nil {
		return
	}
	return k.Name
}

func (k *KnowledgeBaseDocumentChunkResponseModel) GetContent() (value string) {
	if k == nil {
		return
	}
	return k.Content
}

func (k *KnowledgeBaseDocumentChunkResponseModel) SetId(id string) {
	k.Id = id
}

func (k *KnowledgeBaseDocumentChunkResponseModel) SetName(name string) {
	k.Name = name
}

func (k *KnowledgeBaseDocumentChunkResponseModel) SetContent(content string) {
	k.Content = content
}

func (k *KnowledgeBaseDocumentChunkResponseModel) GetExtraProperties() map[string]interface{} {
	if k == nil {
		return nil
	}
	return k.extraProperties
}

func (k KnowledgeBaseDocumentChunkResponseModel) String() string {
	jsonData, err := json.MarshalIndent(k, "", "  ")
	if err != nil {
		return "error converting struct: KnowledgeBaseDocumentChunkResponseModel to string"
	}
	return string(jsonData)
}

func (k *KnowledgeBaseDocumentChunkResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, k); err != nil {
		return err
	}
	type alias KnowledgeBaseDocumentChunkResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*k = KnowledgeBaseDocumentChunkResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *k)
	if err != nil {
		return err
	}
	k.extraProperties = extraProperties
	return nil
}

func (k *KnowledgeBaseDocumentChunkResponseModel) MarshalJSON() ([]byte, error) {
	type embed KnowledgeBaseDocumentChunkResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*k),
	}
	return json.Marshal(marshaler)
}
