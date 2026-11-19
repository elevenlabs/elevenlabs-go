package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationHistoryRagUsageCommonModel struct {
	UsageCount      int    `json:"usage_count" xml:"usage_count" required:"true"`
	EmbeddingModel  string `json:"embedding_model" xml:"embedding_model" required:"true"`
	extraProperties map[string]interface{}
}

func (c *ConversationHistoryRagUsageCommonModel) GetUsageCount() (value int) {
	if c == nil {
		return
	}
	return c.UsageCount
}

func (c *ConversationHistoryRagUsageCommonModel) GetEmbeddingModel() (value string) {
	if c == nil {
		return
	}
	return c.EmbeddingModel
}

func (c *ConversationHistoryRagUsageCommonModel) SetUsageCount(usageCount int) {
	c.UsageCount = usageCount
}

func (c *ConversationHistoryRagUsageCommonModel) SetEmbeddingModel(embeddingModel string) {
	c.EmbeddingModel = embeddingModel
}

func (c *ConversationHistoryRagUsageCommonModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationHistoryRagUsageCommonModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationHistoryRagUsageCommonModel to string"
	}
	return string(jsonData)
}

func (c *ConversationHistoryRagUsageCommonModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ConversationHistoryRagUsageCommonModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ConversationHistoryRagUsageCommonModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationHistoryRagUsageCommonModel) MarshalJSON() ([]byte, error) {
	type embed ConversationHistoryRagUsageCommonModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
