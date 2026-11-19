package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ConversationHistoryRedactionConfig struct {
	// Whether conversation history redaction is enabled
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The entities to redact from the conversation transcript, audio and analysis. Use top-level types like 'name', 'email_address', or dot notation for specific subtypes like 'name.full_name'.
	Entities []ConfigEntityType `json:"entities,omitempty" xml:"entities,omitempty"`
	// Data collection item IDs whose extracted values are not redacted. Their rationales are still redacted.
	ExcludedDataCollectionIds []string `json:"excluded_data_collection_ids,omitempty" xml:"excluded_data_collection_ids,omitempty" maxItems:"50"`
	extraProperties           map[string]interface{}
}

func (c *ConversationHistoryRedactionConfig) GetEnabled() (value *bool) {
	if c == nil {
		return
	}
	return c.Enabled
}

func (c *ConversationHistoryRedactionConfig) GetEntities() (value []ConfigEntityType) {
	if c == nil {
		return
	}
	return c.Entities
}

func (c *ConversationHistoryRedactionConfig) GetExcludedDataCollectionIds() (value []string) {
	if c == nil {
		return
	}
	return c.ExcludedDataCollectionIds
}

func (c *ConversationHistoryRedactionConfig) SetEnabled(enabled *bool) {
	c.Enabled = enabled
}

func (c *ConversationHistoryRedactionConfig) SetEntities(entities []ConfigEntityType) {
	c.Entities = entities
}

func (c *ConversationHistoryRedactionConfig) SetExcludedDataCollectionIds(excludedDataCollectionIds []string) {
	c.ExcludedDataCollectionIds = excludedDataCollectionIds
}

func (c *ConversationHistoryRedactionConfig) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationHistoryRedactionConfig) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationHistoryRedactionConfig to string"
	}
	return string(jsonData)
}

func (c *ConversationHistoryRedactionConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler ConversationHistoryRedactionConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConversationHistoryRedactionConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationHistoryRedactionConfig) MarshalJSON() ([]byte, error) {
	type embed ConversationHistoryRedactionConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
