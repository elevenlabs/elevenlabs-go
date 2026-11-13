package types

import (
	"encoding/json"
	"fmt"
)

type ConversationHistoryTranscriptCommonModelInputToolResultsItem struct {
	ConversationHistoryTranscriptOtherToolsResultCommonModel                      *ConversationHistoryTranscriptOtherToolsResultCommonModel
	ConversationHistoryTranscriptSystemToolResultCommonModelInput                 *ConversationHistoryTranscriptSystemToolResultCommonModelInput
	ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput
	ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput              *ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput
	typ                                                                           string
}

func (c *ConversationHistoryTranscriptCommonModelInputToolResultsItem) GetConversationHistoryTranscriptOtherToolsResultCommonModel() *ConversationHistoryTranscriptOtherToolsResultCommonModel {
	if c == nil {
		return nil
	}
	if c.ConversationHistoryTranscriptOtherToolsResultCommonModel == nil {
		return nil
	}
	return c.ConversationHistoryTranscriptOtherToolsResultCommonModel
}

func (c *ConversationHistoryTranscriptCommonModelInputToolResultsItem) GetConversationHistoryTranscriptSystemToolResultCommonModelInput() *ConversationHistoryTranscriptSystemToolResultCommonModelInput {
	if c == nil {
		return nil
	}
	if c.ConversationHistoryTranscriptSystemToolResultCommonModelInput == nil {
		return nil
	}
	return c.ConversationHistoryTranscriptSystemToolResultCommonModelInput
}

func (c *ConversationHistoryTranscriptCommonModelInputToolResultsItem) GetConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput() *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput {
	if c == nil {
		return nil
	}
	if c.ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput == nil {
		return nil
	}
	return c.ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput
}

func (c *ConversationHistoryTranscriptCommonModelInputToolResultsItem) GetConversationHistoryTranscriptWorkflowToolsResultCommonModelInput() *ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput {
	if c == nil {
		return nil
	}
	if c.ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput == nil {
		return nil
	}
	return c.ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput
}

func (c *ConversationHistoryTranscriptCommonModelInputToolResultsItem) UnmarshalJSON(data []byte) error {
	var unionKeys map[string]json.RawMessage
	_ = json.Unmarshal(data, &unionKeys)
	hasLiteral := func(key string, want interface{}) bool {
		raw, ok := unionKeys[key]
		if !ok {
			return false
		}
		var got interface{}
		return json.Unmarshal(raw, &got) == nil && got == want
	}
	matchedLiteral := false
	if hasLiteral("type", "system") {
		matchedLiteral = true
		var valueConversationHistoryTranscriptSystemToolResultCommonModelInputLiteral *ConversationHistoryTranscriptSystemToolResultCommonModelInput
		if err := json.Unmarshal(data, &valueConversationHistoryTranscriptSystemToolResultCommonModelInputLiteral); err == nil {
			c.typ = "ConversationHistoryTranscriptSystemToolResultCommonModelInput"
			c.ConversationHistoryTranscriptSystemToolResultCommonModelInput = valueConversationHistoryTranscriptSystemToolResultCommonModelInputLiteral
			return nil
		}
	}
	if hasLiteral("type", "api_integration_webhook") {
		matchedLiteral = true
		var valueConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInputLiteral *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput
		if err := json.Unmarshal(data, &valueConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInputLiteral); err == nil {
			c.typ = "ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput"
			c.ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput = valueConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInputLiteral
			return nil
		}
	}
	if hasLiteral("type", "workflow") {
		matchedLiteral = true
		var valueConversationHistoryTranscriptWorkflowToolsResultCommonModelInputLiteral *ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput
		if err := json.Unmarshal(data, &valueConversationHistoryTranscriptWorkflowToolsResultCommonModelInputLiteral); err == nil {
			c.typ = "ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput"
			c.ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput = valueConversationHistoryTranscriptWorkflowToolsResultCommonModelInputLiteral
			return nil
		}
	}
	if matchedLiteral {
		return fmt.Errorf("%s cannot be deserialized as a %T", data, c)
	}
	var valueConversationHistoryTranscriptOtherToolsResultCommonModel *ConversationHistoryTranscriptOtherToolsResultCommonModel
	if err := json.Unmarshal(data, &valueConversationHistoryTranscriptOtherToolsResultCommonModel); err == nil {
		c.typ = "ConversationHistoryTranscriptOtherToolsResultCommonModel"
		c.ConversationHistoryTranscriptOtherToolsResultCommonModel = valueConversationHistoryTranscriptOtherToolsResultCommonModel
		return nil
	}
	var valueConversationHistoryTranscriptSystemToolResultCommonModelInput *ConversationHistoryTranscriptSystemToolResultCommonModelInput
	if err := json.Unmarshal(data, &valueConversationHistoryTranscriptSystemToolResultCommonModelInput); err == nil {
		c.typ = "ConversationHistoryTranscriptSystemToolResultCommonModelInput"
		c.ConversationHistoryTranscriptSystemToolResultCommonModelInput = valueConversationHistoryTranscriptSystemToolResultCommonModelInput
		return nil
	}
	var valueConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput
	if err := json.Unmarshal(data, &valueConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput); err == nil {
		c.typ = "ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput"
		c.ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput = valueConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput
		return nil
	}
	var valueConversationHistoryTranscriptWorkflowToolsResultCommonModelInput *ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput
	if err := json.Unmarshal(data, &valueConversationHistoryTranscriptWorkflowToolsResultCommonModelInput); err == nil {
		c.typ = "ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput"
		c.ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput = valueConversationHistoryTranscriptWorkflowToolsResultCommonModelInput
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, c)
}

func (c ConversationHistoryTranscriptCommonModelInputToolResultsItem) MarshalJSON() ([]byte, error) {
	if c.typ == "ConversationHistoryTranscriptOtherToolsResultCommonModel" || c.ConversationHistoryTranscriptOtherToolsResultCommonModel != nil {
		return json.Marshal(c.ConversationHistoryTranscriptOtherToolsResultCommonModel)
	}
	if c.typ == "ConversationHistoryTranscriptSystemToolResultCommonModelInput" || c.ConversationHistoryTranscriptSystemToolResultCommonModelInput != nil {
		return json.Marshal(c.ConversationHistoryTranscriptSystemToolResultCommonModelInput)
	}
	if c.typ == "ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput" || c.ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput != nil {
		return json.Marshal(c.ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput)
	}
	if c.typ == "ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput" || c.ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput != nil {
		return json.Marshal(c.ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", c)
}

type ConversationHistoryTranscriptCommonModelInputToolResultsItemVisitor interface {
	VisitConversationHistoryTranscriptOtherToolsResultCommonModel(*ConversationHistoryTranscriptOtherToolsResultCommonModel) error
	VisitConversationHistoryTranscriptSystemToolResultCommonModelInput(*ConversationHistoryTranscriptSystemToolResultCommonModelInput) error
	VisitConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput(*ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) error
	VisitConversationHistoryTranscriptWorkflowToolsResultCommonModelInput(*ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput) error
}

func (c *ConversationHistoryTranscriptCommonModelInputToolResultsItem) Accept(visitor ConversationHistoryTranscriptCommonModelInputToolResultsItemVisitor) error {
	if c.typ == "ConversationHistoryTranscriptOtherToolsResultCommonModel" || c.ConversationHistoryTranscriptOtherToolsResultCommonModel != nil {
		return visitor.VisitConversationHistoryTranscriptOtherToolsResultCommonModel(c.ConversationHistoryTranscriptOtherToolsResultCommonModel)
	}
	if c.typ == "ConversationHistoryTranscriptSystemToolResultCommonModelInput" || c.ConversationHistoryTranscriptSystemToolResultCommonModelInput != nil {
		return visitor.VisitConversationHistoryTranscriptSystemToolResultCommonModelInput(c.ConversationHistoryTranscriptSystemToolResultCommonModelInput)
	}
	if c.typ == "ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput" || c.ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput != nil {
		return visitor.VisitConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput(c.ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput)
	}
	if c.typ == "ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput" || c.ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput != nil {
		return visitor.VisitConversationHistoryTranscriptWorkflowToolsResultCommonModelInput(c.ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", c)
}
