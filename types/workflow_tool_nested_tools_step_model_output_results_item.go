package types

import (
	"encoding/json"
	"fmt"
)

type WorkflowToolNestedToolsStepModelOutputResultsItem struct {
	ConversationHistoryTranscriptOtherToolsResultCommonModel                       *ConversationHistoryTranscriptOtherToolsResultCommonModel
	ConversationHistoryTranscriptSystemToolResultCommonModelOutput                 *ConversationHistoryTranscriptSystemToolResultCommonModelOutput
	ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput
	ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput              *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput
	typ                                                                            string
}

func (w *WorkflowToolNestedToolsStepModelOutputResultsItem) GetConversationHistoryTranscriptOtherToolsResultCommonModel() *ConversationHistoryTranscriptOtherToolsResultCommonModel {
	if w == nil {
		return nil
	}
	if w.ConversationHistoryTranscriptOtherToolsResultCommonModel == nil {
		return nil
	}
	return w.ConversationHistoryTranscriptOtherToolsResultCommonModel
}

func (w *WorkflowToolNestedToolsStepModelOutputResultsItem) GetConversationHistoryTranscriptSystemToolResultCommonModelOutput() *ConversationHistoryTranscriptSystemToolResultCommonModelOutput {
	if w == nil {
		return nil
	}
	if w.ConversationHistoryTranscriptSystemToolResultCommonModelOutput == nil {
		return nil
	}
	return w.ConversationHistoryTranscriptSystemToolResultCommonModelOutput
}

func (w *WorkflowToolNestedToolsStepModelOutputResultsItem) GetConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput() *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput {
	if w == nil {
		return nil
	}
	if w.ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput == nil {
		return nil
	}
	return w.ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput
}

func (w *WorkflowToolNestedToolsStepModelOutputResultsItem) GetConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput() *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput {
	if w == nil {
		return nil
	}
	if w.ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput == nil {
		return nil
	}
	return w.ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput
}

func (w *WorkflowToolNestedToolsStepModelOutputResultsItem) UnmarshalJSON(data []byte) error {
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
			w.typ = "ConversationHistoryTranscriptSystemToolResultCommonModelOutput"
			w.ConversationHistoryTranscriptSystemToolResultCommonModelOutput = valueConversationHistoryTranscriptSystemToolResultCommonModelOutputLiteral
			return nil
		}
	}
	if hasLiteral("type", "api_integration_webhook") {
		matchedLiteral = true
		var valueConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutputLiteral *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput
		if err := json.Unmarshal(data, &valueConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutputLiteral); err == nil {
			w.typ = "ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput"
			w.ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput = valueConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutputLiteral
			return nil
		}
	}
	if hasLiteral("type", "workflow") {
		matchedLiteral = true
		var valueConversationHistoryTranscriptWorkflowToolsResultCommonModelOutputLiteral *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput
		if err := json.Unmarshal(data, &valueConversationHistoryTranscriptWorkflowToolsResultCommonModelOutputLiteral); err == nil {
			w.typ = "ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput"
			w.ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput = valueConversationHistoryTranscriptWorkflowToolsResultCommonModelOutputLiteral
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
	var valueConversationHistoryTranscriptSystemToolResultCommonModelOutput *ConversationHistoryTranscriptSystemToolResultCommonModelOutput
	if err := json.Unmarshal(data, &valueConversationHistoryTranscriptSystemToolResultCommonModelOutput); err == nil {
		w.typ = "ConversationHistoryTranscriptSystemToolResultCommonModelOutput"
		w.ConversationHistoryTranscriptSystemToolResultCommonModelOutput = valueConversationHistoryTranscriptSystemToolResultCommonModelOutput
		return nil
	}
	var valueConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput *ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput
	if err := json.Unmarshal(data, &valueConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput); err == nil {
		w.typ = "ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput"
		w.ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput = valueConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput
		return nil
	}
	var valueConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput
	if err := json.Unmarshal(data, &valueConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput); err == nil {
		w.typ = "ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput"
		w.ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput = valueConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, w)
}

func (w WorkflowToolNestedToolsStepModelOutputResultsItem) MarshalJSON() ([]byte, error) {
	if w.typ == "ConversationHistoryTranscriptOtherToolsResultCommonModel" || w.ConversationHistoryTranscriptOtherToolsResultCommonModel != nil {
		return json.Marshal(w.ConversationHistoryTranscriptOtherToolsResultCommonModel)
	}
	if w.typ == "ConversationHistoryTranscriptSystemToolResultCommonModelOutput" || w.ConversationHistoryTranscriptSystemToolResultCommonModelOutput != nil {
		return json.Marshal(w.ConversationHistoryTranscriptSystemToolResultCommonModelOutput)
	}
	if w.typ == "ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput" || w.ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput != nil {
		return json.Marshal(w.ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput)
	}
	if w.typ == "ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput" || w.ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput != nil {
		return json.Marshal(w.ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", w)
}

type WorkflowToolNestedToolsStepModelOutputResultsItemVisitor interface {
	VisitConversationHistoryTranscriptOtherToolsResultCommonModel(*ConversationHistoryTranscriptOtherToolsResultCommonModel) error
	VisitConversationHistoryTranscriptSystemToolResultCommonModelOutput(*ConversationHistoryTranscriptSystemToolResultCommonModelOutput) error
	VisitConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput(*ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput) error
	VisitConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput(*ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) error
}

func (w *WorkflowToolNestedToolsStepModelOutputResultsItem) Accept(visitor WorkflowToolNestedToolsStepModelOutputResultsItemVisitor) error {
	if w.typ == "ConversationHistoryTranscriptOtherToolsResultCommonModel" || w.ConversationHistoryTranscriptOtherToolsResultCommonModel != nil {
		return visitor.VisitConversationHistoryTranscriptOtherToolsResultCommonModel(w.ConversationHistoryTranscriptOtherToolsResultCommonModel)
	}
	if w.typ == "ConversationHistoryTranscriptSystemToolResultCommonModelOutput" || w.ConversationHistoryTranscriptSystemToolResultCommonModelOutput != nil {
		return visitor.VisitConversationHistoryTranscriptSystemToolResultCommonModelOutput(w.ConversationHistoryTranscriptSystemToolResultCommonModelOutput)
	}
	if w.typ == "ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput" || w.ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput != nil {
		return visitor.VisitConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput(w.ConversationHistoryTranscriptApiIntegrationWebhookToolsResultCommonModelOutput)
	}
	if w.typ == "ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput" || w.ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput != nil {
		return visitor.VisitConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput(w.ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", w)
}
