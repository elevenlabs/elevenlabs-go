package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationHistoryBatchCallModel struct {
	BatchCallId          string                           `json:"batch_call_id" xml:"batch_call_id" required:"true"`
	BatchCallRecipientId string                           `json:"batch_call_recipient_id" xml:"batch_call_recipient_id" required:"true"`
	Campaign             *BatchCallingCampaignInformation `json:"campaign,omitempty" xml:"campaign,omitempty"`
	extraProperties      map[string]interface{}
}

func (c *ConversationHistoryBatchCallModel) GetBatchCallId() (value string) {
	if c == nil {
		return
	}
	return c.BatchCallId
}

func (c *ConversationHistoryBatchCallModel) GetBatchCallRecipientId() (value string) {
	if c == nil {
		return
	}
	return c.BatchCallRecipientId
}

func (c *ConversationHistoryBatchCallModel) GetCampaign() (value *BatchCallingCampaignInformation) {
	if c == nil {
		return
	}
	return c.Campaign
}

func (c *ConversationHistoryBatchCallModel) SetBatchCallId(batchCallId string) {
	c.BatchCallId = batchCallId
}

func (c *ConversationHistoryBatchCallModel) SetBatchCallRecipientId(batchCallRecipientId string) {
	c.BatchCallRecipientId = batchCallRecipientId
}

func (c *ConversationHistoryBatchCallModel) SetCampaign(campaign *BatchCallingCampaignInformation) {
	c.Campaign = campaign
}

func (c *ConversationHistoryBatchCallModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationHistoryBatchCallModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationHistoryBatchCallModel to string"
	}
	return string(jsonData)
}

func (c *ConversationHistoryBatchCallModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationHistoryBatchCallModel) MarshalJSON() ([]byte, error) {
	type embed ConversationHistoryBatchCallModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
