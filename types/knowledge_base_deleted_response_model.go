package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type KnowledgeBaseDeletedResponseModel struct {
	Id              string `json:"id" xml:"id" required:"true"`
	extraProperties map[string]interface{}
}

func (k *KnowledgeBaseDeletedResponseModel) GetId() (value string) {
	if k == nil {
		return
	}
	return k.Id
}

func (k *KnowledgeBaseDeletedResponseModel) SetId(id string) {
	k.Id = id
}

func (k *KnowledgeBaseDeletedResponseModel) GetExtraProperties() map[string]interface{} {
	if k == nil {
		return nil
	}
	return k.extraProperties
}

func (k KnowledgeBaseDeletedResponseModel) String() string {
	jsonData, err := json.MarshalIndent(k, "", "  ")
	if err != nil {
		return "error converting struct: KnowledgeBaseDeletedResponseModel to string"
	}
	return string(jsonData)
}

func (k *KnowledgeBaseDeletedResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, k); err != nil {
		return err
	}
	type alias KnowledgeBaseDeletedResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*k = KnowledgeBaseDeletedResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *k)
	if err != nil {
		return err
	}
	k.extraProperties = extraProperties
	return nil
}

func (k *KnowledgeBaseDeletedResponseModel) MarshalJSON() ([]byte, error) {
	type embed KnowledgeBaseDeletedResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*k),
	}
	return json.Marshal(marshaler)
}
