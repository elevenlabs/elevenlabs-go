package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type KnowledgeBaseRagToolConfig struct {
	extraProperties map[string]interface{}
}

func (k *KnowledgeBaseRagToolConfig) GetExtraProperties() map[string]interface{} {
	if k == nil {
		return nil
	}
	return k.extraProperties
}

func (k KnowledgeBaseRagToolConfig) String() string {
	jsonData, err := json.MarshalIndent(k, "", "  ")
	if err != nil {
		return "error converting struct: KnowledgeBaseRagToolConfig to string"
	}
	return string(jsonData)
}

func (k *KnowledgeBaseRagToolConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler KnowledgeBaseRagToolConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*k = KnowledgeBaseRagToolConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *k, "system_tool_type")
	if err != nil {
		return err
	}
	k.extraProperties = extraProperties
	return nil
}

func (k *KnowledgeBaseRagToolConfig) MarshalJSON() ([]byte, error) {
	type embed KnowledgeBaseRagToolConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*k),
	}
	return json.Marshal(marshaler)
}
