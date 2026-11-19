package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type KnowledgeBaseToolInfo struct {
	// Search strategies exposed to the model. Must be non-empty.
	EnabledStrategies []SearchStrategy `json:"enabled_strategies,omitempty" xml:"enabled_strategies,omitempty"`
	extraProperties   map[string]interface{}
}

func (k *KnowledgeBaseToolInfo) GetEnabledStrategies() (value []SearchStrategy) {
	if k == nil {
		return
	}
	return k.EnabledStrategies
}

func (k *KnowledgeBaseToolInfo) SetEnabledStrategies(enabledStrategies []SearchStrategy) {
	k.EnabledStrategies = enabledStrategies
}

func (k *KnowledgeBaseToolInfo) GetExtraProperties() map[string]interface{} {
	if k == nil {
		return nil
	}
	return k.extraProperties
}

func (k KnowledgeBaseToolInfo) String() string {
	jsonData, err := json.MarshalIndent(k, "", "  ")
	if err != nil {
		return "error converting struct: KnowledgeBaseToolInfo to string"
	}
	return string(jsonData)
}

func (k *KnowledgeBaseToolInfo) UnmarshalJSON(data []byte) error {
	type unmarshaler KnowledgeBaseToolInfo
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*k = KnowledgeBaseToolInfo(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *k)
	if err != nil {
		return err
	}
	k.extraProperties = extraProperties
	return nil
}

func (k *KnowledgeBaseToolInfo) MarshalJSON() ([]byte, error) {
	type embed KnowledgeBaseToolInfo
	var marshaler = struct {
		embed
	}{
		embed: embed(*k),
	}
	return json.Marshal(marshaler)
}
