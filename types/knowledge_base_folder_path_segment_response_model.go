package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type KnowledgeBaseFolderPathSegmentResponseModel struct {
	Id              string  `json:"id" xml:"id" required:"true"`
	Name            *string `json:"name" xml:"name" required:"true"`
	extraProperties map[string]interface{}
}

func (k *KnowledgeBaseFolderPathSegmentResponseModel) GetId() (value string) {
	if k == nil {
		return
	}
	return k.Id
}

func (k *KnowledgeBaseFolderPathSegmentResponseModel) GetName() (value *string) {
	if k == nil {
		return
	}
	return k.Name
}

func (k *KnowledgeBaseFolderPathSegmentResponseModel) SetId(id string) {
	k.Id = id
}

func (k *KnowledgeBaseFolderPathSegmentResponseModel) SetName(name *string) {
	k.Name = name
}

func (k *KnowledgeBaseFolderPathSegmentResponseModel) GetExtraProperties() map[string]interface{} {
	if k == nil {
		return nil
	}
	return k.extraProperties
}

func (k KnowledgeBaseFolderPathSegmentResponseModel) String() string {
	jsonData, err := json.MarshalIndent(k, "", "  ")
	if err != nil {
		return "error converting struct: KnowledgeBaseFolderPathSegmentResponseModel to string"
	}
	return string(jsonData)
}

func (k *KnowledgeBaseFolderPathSegmentResponseModel) UnmarshalJSON(data []byte) error {
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

func (k *KnowledgeBaseFolderPathSegmentResponseModel) MarshalJSON() ([]byte, error) {
	type embed KnowledgeBaseFolderPathSegmentResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*k),
	}
	return json.Marshal(marshaler)
}
