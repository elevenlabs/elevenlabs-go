package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type KnowledgeBaseToolConfig struct {
	EnabledStrategies []SearchStrategy `json:"enabled_strategies,omitempty" xml:"enabled_strategies,omitempty"`
	extraProperties   map[string]interface{}
}

func (k *KnowledgeBaseToolConfig) GetEnabledStrategies() (value []SearchStrategy) {
	if k == nil {
		return
	}
	return k.EnabledStrategies
}

func (k *KnowledgeBaseToolConfig) SetEnabledStrategies(enabledStrategies []SearchStrategy) {
	k.EnabledStrategies = enabledStrategies
}

func (k *KnowledgeBaseToolConfig) GetExtraProperties() map[string]interface{} {
	if k == nil {
		return nil
	}
	return k.extraProperties
}

func (k KnowledgeBaseToolConfig) String() string {
	jsonData, err := json.MarshalIndent(k, "", "  ")
	if err != nil {
		return "error converting struct: KnowledgeBaseToolConfig to string"
	}
	return string(jsonData)
}

func (k *KnowledgeBaseToolConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler KnowledgeBaseToolConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*k = KnowledgeBaseToolConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *k, "system_tool_type")
	if err != nil {
		return err
	}
	k.extraProperties = extraProperties
	return nil
}

func (k *KnowledgeBaseToolConfig) MarshalJSON() ([]byte, error) {
	type embed KnowledgeBaseToolConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*k),
	}
	return json.Marshal(marshaler)
}
