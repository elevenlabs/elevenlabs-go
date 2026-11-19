package types

import (
	"encoding/json"
	"fmt"
)

type ConversationHistoryTranscriptCommonModelOutputToolResultsItem struct {
	ConversationHistoryTranscriptOtherToolsResultCommonModel                       *ConversationHistoryTranscriptOtherToolsResultCommonModel
	ConversationHistoryTranscriptSystemToolResultCommonModelOutput                 *ConversationHistoryTranscriptSystemToolResultCommonModelOutput
	ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput
	ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput              *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput
	typ                                                                            string
}

func (c *ConversationHistoryTranscriptCommonModelOutputToolResultsItem) GetConversationHistoryTranscriptOtherToolsResultCommonModel() *ConversationHistoryTranscriptOtherToolsResultCommonModel {
	if c == nil {
		return nil
	}
	if c.ConversationHistoryTranscriptOtherToolsResultCommonModel == nil {
		return nil
	}
	return c.ConversationHistoryTranscriptOtherToolsResultCommonModel
}

func (c *ConversationHistoryTranscriptCommonModelOutputToolResultsItem) GetConversationHistoryTranscriptSystemToolResultCommonModelOutput() *ConversationHistoryTranscriptSystemToolResultCommonModelOutput {
	if c == nil {
		return nil
	}
	if c.ConversationHistoryTranscriptSystemToolResultCommonModelOutput == nil {
		return nil
	}
	return c.ConversationHistoryTranscriptSystemToolResultCommonModelOutput
}

func (c *ConversationHistoryTranscriptCommonModelOutputToolResultsItem) GetConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput() *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput {
	if c == nil {
		return nil
	}
	if c.ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput == nil {
		return nil
	}
	return c.ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput
}

func (c *ConversationHistoryTranscriptCommonModelOutputToolResultsItem) GetConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput() *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput {
	if c == nil {
		return nil
	}
	if c.ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput == nil {
		return nil
	}
	return c.ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput
}

func (c *ConversationHistoryTranscriptCommonModelOutputToolResultsItem) UnmarshalJSON(data []byte) error {
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
		var valueConversationHistoryTranscriptSystemToolResultCommonModelOutputLiteral *ConversationHistoryTranscriptSystemToolResultCommonModelOutput
		if err := json.Unmarshal(data, &valueConversationHistoryTranscriptSystemToolResultCommonModelOutputLiteral); err == nil {
			c.typ = "ConversationHistoryTranscriptSystemToolResultCommonModelOutput"
			c.ConversationHistoryTranscriptSystemToolResultCommonModelOutput = valueConversationHistoryTranscriptSystemToolResultCommonModelOutputLiteral
			return nil
		}
	}
	if hasLiteral("type", "api_integration_webhook") {
		matchedLiteral = true
		var valueConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutputLiteral *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput
		if err := json.Unmarshal(data, &valueConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutputLiteral); err == nil {
			c.typ = "ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput"
			c.ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput = valueConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutputLiteral
			return nil
		}
	}
	if hasLiteral("type", "workflow") {
		matchedLiteral = true
		var valueConversationHistoryTranscriptWorkflowToolsResultCommonModelOutputLiteral *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput
		if err := json.Unmarshal(data, &valueConversationHistoryTranscriptWorkflowToolsResultCommonModelOutputLiteral); err == nil {
			c.typ = "ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput"
			c.ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput = valueConversationHistoryTranscriptWorkflowToolsResultCommonModelOutputLiteral
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
	var valueConversationHistoryTranscriptSystemToolResultCommonModelOutput *ConversationHistoryTranscriptSystemToolResultCommonModelOutput
	if err := json.Unmarshal(data, &valueConversationHistoryTranscriptSystemToolResultCommonModelOutput); err == nil {
		c.typ = "ConversationHistoryTranscriptSystemToolResultCommonModelOutput"
		c.ConversationHistoryTranscriptSystemToolResultCommonModelOutput = valueConversationHistoryTranscriptSystemToolResultCommonModelOutput
		return nil
	}
	var valueConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput
	if err := json.Unmarshal(data, &valueConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput); err == nil {
		c.typ = "ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput"
		c.ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput = valueConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput
		return nil
	}
	var valueConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput
	if err := json.Unmarshal(data, &valueConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput); err == nil {
		c.typ = "ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput"
		c.ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput = valueConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, c)
}

func (c ConversationHistoryTranscriptCommonModelOutputToolResultsItem) MarshalJSON() ([]byte, error) {
	if c.typ == "ConversationHistoryTranscriptOtherToolsResultCommonModel" || c.ConversationHistoryTranscriptOtherToolsResultCommonModel != nil {
		return json.Marshal(c.ConversationHistoryTranscriptOtherToolsResultCommonModel)
	}
	if c.typ == "ConversationHistoryTranscriptSystemToolResultCommonModelOutput" || c.ConversationHistoryTranscriptSystemToolResultCommonModelOutput != nil {
		return json.Marshal(c.ConversationHistoryTranscriptSystemToolResultCommonModelOutput)
	}
	if c.typ == "ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput" || c.ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput != nil {
		return json.Marshal(c.ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput)
	}
	if c.typ == "ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput" || c.ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput != nil {
		return json.Marshal(c.ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", c)
}

type ConversationHistoryTranscriptCommonModelOutputToolResultsItemVisitor interface {
	VisitConversationHistoryTranscriptOtherToolsResultCommonModel(*ConversationHistoryTranscriptOtherToolsResultCommonModel) error
	VisitConversationHistoryTranscriptSystemToolResultCommonModelOutput(*ConversationHistoryTranscriptSystemToolResultCommonModelOutput) error
	VisitConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput(*ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) error
	VisitConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput(*ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) error
}

func (c *ConversationHistoryTranscriptCommonModelOutputToolResultsItem) Accept(visitor ConversationHistoryTranscriptCommonModelOutputToolResultsItemVisitor) error {
	if c.typ == "ConversationHistoryTranscriptOtherToolsResultCommonModel" || c.ConversationHistoryTranscriptOtherToolsResultCommonModel != nil {
		return visitor.VisitConversationHistoryTranscriptOtherToolsResultCommonModel(c.ConversationHistoryTranscriptOtherToolsResultCommonModel)
	}
	if c.typ == "ConversationHistoryTranscriptSystemToolResultCommonModelOutput" || c.ConversationHistoryTranscriptSystemToolResultCommonModelOutput != nil {
		return visitor.VisitConversationHistoryTranscriptSystemToolResultCommonModelOutput(c.ConversationHistoryTranscriptSystemToolResultCommonModelOutput)
	}
	if c.typ == "ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput" || c.ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput != nil {
		return visitor.VisitConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput(c.ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput)
	}
	if c.typ == "ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput" || c.ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput != nil {
		return visitor.VisitConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput(c.ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", c)
}
