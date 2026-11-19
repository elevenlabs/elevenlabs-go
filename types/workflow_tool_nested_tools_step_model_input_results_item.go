package types

import (
	"encoding/json"
	"fmt"
)

type WorkflowToolNestedToolsStepModelInputResultsItem struct {
	ConversationHistoryTranscriptOtherToolsResultCommonModel                      *ConversationHistoryTranscriptOtherToolsResultCommonModel
	ConversationHistoryTranscriptSystemToolResultCommonModelInput                 *ConversationHistoryTranscriptSystemToolResultCommonModelInput
	ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput
	ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput              *ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput
	typ                                                                           string
}

func (w *WorkflowToolNestedToolsStepModelInputResultsItem) GetConversationHistoryTranscriptOtherToolsResultCommonModel() *ConversationHistoryTranscriptOtherToolsResultCommonModel {
	if w == nil {
		return nil
	}
	if w.ConversationHistoryTranscriptOtherToolsResultCommonModel == nil {
		return nil
	}
	return w.ConversationHistoryTranscriptOtherToolsResultCommonModel
}

func (w *WorkflowToolNestedToolsStepModelInputResultsItem) GetConversationHistoryTranscriptSystemToolResultCommonModelInput() *ConversationHistoryTranscriptSystemToolResultCommonModelInput {
	if w == nil {
		return nil
	}
	if w.ConversationHistoryTranscriptSystemToolResultCommonModelInput == nil {
		return nil
	}
	return w.ConversationHistoryTranscriptSystemToolResultCommonModelInput
}

func (w *WorkflowToolNestedToolsStepModelInputResultsItem) GetConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput() *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput {
	if w == nil {
		return nil
	}
	if w.ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput == nil {
		return nil
	}
	return w.ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput
}

func (w *WorkflowToolNestedToolsStepModelInputResultsItem) GetConversationHistoryTranscriptWorkflowToolsResultCommonModelInput() *ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput {
	if w == nil {
		return nil
	}
	if w.ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput == nil {
		return nil
	}
	return w.ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput
}

func (w *WorkflowToolNestedToolsStepModelInputResultsItem) UnmarshalJSON(data []byte) error {
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
			w.typ = "ConversationHistoryTranscriptSystemToolResultCommonModelInput"
			w.ConversationHistoryTranscriptSystemToolResultCommonModelInput = valueConversationHistoryTranscriptSystemToolResultCommonModelInputLiteral
			return nil
		}
	}
	if hasLiteral("type", "api_integration_webhook") {
		matchedLiteral = true
		var valueConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInputLiteral *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput
		if err := json.Unmarshal(data, &valueConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInputLiteral); err == nil {
			w.typ = "ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput"
			w.ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput = valueConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInputLiteral
			return nil
		}
	}
	if hasLiteral("type", "workflow") {
		matchedLiteral = true
		var valueConversationHistoryTranscriptWorkflowToolsResultCommonModelInputLiteral *ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput
		if err := json.Unmarshal(data, &valueConversationHistoryTranscriptWorkflowToolsResultCommonModelInputLiteral); err == nil {
			w.typ = "ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput"
			w.ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput = valueConversationHistoryTranscriptWorkflowToolsResultCommonModelInputLiteral
			return nil
		}
	}
	if matchedLiteral {
		return fmt.Errorf("%s cannot be deserialized as a %T", data, w)
	}
	var valueConversationHistoryTranscriptOtherToolsResultCommonModel *ConversationHistoryTranscriptOtherToolsResultCommonModel
	if err := json.Unmarshal(data, &valueConversationHistoryTranscriptOtherToolsResultCommonModel); err == nil {
		w.typ = "ConversationHistoryTranscriptOtherToolsResultCommonModel"
		w.ConversationHistoryTranscriptOtherToolsResultCommonModel = valueConversationHistoryTranscriptOtherToolsResultCommonModel
		return nil
	}
	var valueConversationHistoryTranscriptSystemToolResultCommonModelInput *ConversationHistoryTranscriptSystemToolResultCommonModelInput
	if err := json.Unmarshal(data, &valueConversationHistoryTranscriptSystemToolResultCommonModelInput); err == nil {
		w.typ = "ConversationHistoryTranscriptSystemToolResultCommonModelInput"
		w.ConversationHistoryTranscriptSystemToolResultCommonModelInput = valueConversationHistoryTranscriptSystemToolResultCommonModelInput
		return nil
	}
	var valueConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput
	if err := json.Unmarshal(data, &valueConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput); err == nil {
		w.typ = "ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput"
		w.ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput = valueConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput
		return nil
	}
	var valueConversationHistoryTranscriptWorkflowToolsResultCommonModelInput *ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput
	if err := json.Unmarshal(data, &valueConversationHistoryTranscriptWorkflowToolsResultCommonModelInput); err == nil {
		w.typ = "ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput"
		w.ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput = valueConversationHistoryTranscriptWorkflowToolsResultCommonModelInput
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, w)
}

func (w WorkflowToolNestedToolsStepModelInputResultsItem) MarshalJSON() ([]byte, error) {
	if w.typ == "ConversationHistoryTranscriptOtherToolsResultCommonModel" || w.ConversationHistoryTranscriptOtherToolsResultCommonModel != nil {
		return json.Marshal(w.ConversationHistoryTranscriptOtherToolsResultCommonModel)
	}
	if w.typ == "ConversationHistoryTranscriptSystemToolResultCommonModelInput" || w.ConversationHistoryTranscriptSystemToolResultCommonModelInput != nil {
		return json.Marshal(w.ConversationHistoryTranscriptSystemToolResultCommonModelInput)
	}
	if w.typ == "ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput" || w.ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput != nil {
		return json.Marshal(w.ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput)
	}
	if w.typ == "ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput" || w.ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput != nil {
		return json.Marshal(w.ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", w)
}

type WorkflowToolNestedToolsStepModelInputResultsItemVisitor interface {
	VisitConversationHistoryTranscriptOtherToolsResultCommonModel(*ConversationHistoryTranscriptOtherToolsResultCommonModel) error
	VisitConversationHistoryTranscriptSystemToolResultCommonModelInput(*ConversationHistoryTranscriptSystemToolResultCommonModelInput) error
	VisitConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput(*ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput) error
	VisitConversationHistoryTranscriptWorkflowToolsResultCommonModelInput(*ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput) error
}

func (w *WorkflowToolNestedToolsStepModelInputResultsItem) Accept(visitor WorkflowToolNestedToolsStepModelInputResultsItemVisitor) error {
	if w.typ == "ConversationHistoryTranscriptOtherToolsResultCommonModel" || w.ConversationHistoryTranscriptOtherToolsResultCommonModel != nil {
		return visitor.VisitConversationHistoryTranscriptOtherToolsResultCommonModel(w.ConversationHistoryTranscriptOtherToolsResultCommonModel)
	}
	if w.typ == "ConversationHistoryTranscriptSystemToolResultCommonModelInput" || w.ConversationHistoryTranscriptSystemToolResultCommonModelInput != nil {
		return visitor.VisitConversationHistoryTranscriptSystemToolResultCommonModelInput(w.ConversationHistoryTranscriptSystemToolResultCommonModelInput)
	}
	if w.typ == "ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput" || w.ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput != nil {
		return visitor.VisitConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput(w.ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelInput)
	}
	if w.typ == "ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput" || w.ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput != nil {
		return visitor.VisitConversationHistoryTranscriptWorkflowToolsResultCommonModelInput(w.ConversationHistoryTranscriptWorkflowToolsResultCommonModelInput)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", w)
}
