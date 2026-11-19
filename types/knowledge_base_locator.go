package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type KnowledgeBaseLocator struct {
	Type KnowledgeBaseDocumentType `json:"type" xml:"type" required:"true"`
	// The name of the knowledge base
	Name string `json:"name" xml:"name" required:"true"`
	// The ID of the knowledge base
	Id              string                 `json:"id" xml:"id" required:"true"`
	UsageMode       *DocumentUsageModeEnum `json:"usage_mode,omitempty" xml:"usage_mode,omitempty"`
	extraProperties map[string]interface{}
}

func (k *KnowledgeBaseLocator) GetType() (value KnowledgeBaseDocumentType) {
	if k == nil {
		return
	}
	return k.Type
}

func (k *KnowledgeBaseLocator) GetName() (value string) {
	if k == nil {
		return
	}
	return k.Name
}

func (k *KnowledgeBaseLocator) GetId() (value string) {
	if k == nil {
		return
	}
	return k.Id
}

func (k *KnowledgeBaseLocator) GetUsageMode() (value *DocumentUsageModeEnum) {
	if k == nil {
		return
	}
	return k.UsageMode
}

func (k *KnowledgeBaseLocator) SetType(type_ KnowledgeBaseDocumentType) {
	k.Type = type_
}

func (k *KnowledgeBaseLocator) SetName(name string) {
	k.Name = name
}

func (k *KnowledgeBaseLocator) SetId(id string) {
	k.Id = id
}

func (k *KnowledgeBaseLocator) SetUsageMode(usageMode *DocumentUsageModeEnum) {
	k.UsageMode = usageMode
}

func (k *KnowledgeBaseLocator) GetExtraProperties() map[string]interface{} {
	if k == nil {
		return nil
	}
	return k.extraProperties
}

func (k KnowledgeBaseLocator) String() string {
	jsonData, err := json.MarshalIndent(k, "", "  ")
	if err != nil {
		return "error converting struct: KnowledgeBaseLocator to string"
	}
	return string(jsonData)
}

func (k *KnowledgeBaseLocator) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, k); err != nil {
		return err
	}
	type alias KnowledgeBaseLocator
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*k = KnowledgeBaseLocator(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *k)
	if err != nil {
		return err
	}
	k.extraProperties = extraProperties
	return nil
}

func (k *KnowledgeBaseLocator) MarshalJSON() ([]byte, error) {
	type embed KnowledgeBaseLocator
	var marshaler = struct {
		embed
	}{
		embed: embed(*k),
	}
	return json.Marshal(marshaler)
}
