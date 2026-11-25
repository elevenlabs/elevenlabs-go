package types

import (
	"encoding/json"
	"fmt"
)

type PublishEvent struct {
	UserAudioChunk                   *UserAudioChunk
	Feedback                         *Feedback
	UserMessage                      *UserMessage
	UserActivity                     *UserActivity
	MultimodalMessage                *MultimodalMessage
	ConversationPong                 *ConversationPong
	ClientToolResult                 *ClientToolResult
	McpToolApprovalResult            *McpToolApprovalResult
	ContextualUpdate                 *ContextualUpdate
	ConversationInitiationClientData *ConversationInitiationClientData
	typ                              string
}

func (p *PublishEvent) GetUserAudioChunk() *UserAudioChunk {
	if p == nil {
		return nil
	}
	if p.UserAudioChunk == nil {
		return nil
	}
	return p.UserAudioChunk
}

func (p *PublishEvent) GetFeedback() *Feedback {
	if p == nil {
		return nil
	}
	if p.Feedback == nil {
		return nil
	}
	return p.Feedback
}

func (p *PublishEvent) GetUserMessage() *UserMessage {
	if p == nil {
		return nil
	}
	if p.UserMessage == nil {
		return nil
	}
	return p.UserMessage
}

func (p *PublishEvent) GetUserActivity() *UserActivity {
	if p == nil {
		return nil
	}
	if p.UserActivity == nil {
		return nil
	}
	return p.UserActivity
}

func (p *PublishEvent) GetMultimodalMessage() *MultimodalMessage {
	if p == nil {
		return nil
	}
	if p.MultimodalMessage == nil {
		return nil
	}
	return p.MultimodalMessage
}

func (p *PublishEvent) GetConversationPong() *ConversationPong {
	if p == nil {
		return nil
	}
	if p.ConversationPong == nil {
		return nil
	}
	return p.ConversationPong
}

func (p *PublishEvent) GetClientToolResult() *ClientToolResult {
	if p == nil {
		return nil
	}
	if p.ClientToolResult == nil {
		return nil
	}
	return p.ClientToolResult
}

func (p *PublishEvent) GetMcpToolApprovalResult() *McpToolApprovalResult {
	if p == nil {
		return nil
	}
	if p.McpToolApprovalResult == nil {
		return nil
	}
	return p.McpToolApprovalResult
}

func (p *PublishEvent) GetContextualUpdate() *ContextualUpdate {
	if p == nil {
		return nil
	}
	if p.ContextualUpdate == nil {
		return nil
	}
	return p.ContextualUpdate
}

func (p *PublishEvent) GetConversationInitiationClientData() *ConversationInitiationClientData {
	if p == nil {
		return nil
	}
	if p.ConversationInitiationClientData == nil {
		return nil
	}
	return p.ConversationInitiationClientData
}

func (p *PublishEvent) UnmarshalJSON(data []byte) error {
	var unionKeys map[string]json.RawMessage
	_ = json.Unmarshal(data, &unionKeys)
	hasAnyKey := func(keys ...string) bool {
		for _, key := range keys {
			if _, ok := unionKeys[key]; ok {
				return true
			}
		}
		return false
	}
	hasLiteral := func(key string, want interface{}) bool {
		raw, ok := unionKeys[key]
		if !ok {
			return false
		}
		var got interface{}
		return json.Unmarshal(raw, &got) == nil && got == want
	}
	matchedLiteral := false
	if hasLiteral("type", "feedback") {
		matchedLiteral = true
		var valueFeedbackLiteral *Feedback
		if err := json.Unmarshal(data, &valueFeedbackLiteral); err == nil {
			p.typ = "Feedback"
			p.Feedback = valueFeedbackLiteral
			return nil
		}
	}
	if hasLiteral("type", "user_message") {
		matchedLiteral = true
		var valueUserMessageLiteral *UserMessage
		if err := json.Unmarshal(data, &valueUserMessageLiteral); err == nil {
			p.typ = "UserMessage"
			p.UserMessage = valueUserMessageLiteral
			return nil
		}
	}
	if hasLiteral("type", "user_activity") {
		matchedLiteral = true
		var valueUserActivityLiteral *UserActivity
		if err := json.Unmarshal(data, &valueUserActivityLiteral); err == nil {
			p.typ = "UserActivity"
			p.UserActivity = valueUserActivityLiteral
			return nil
		}
	}
	if hasLiteral("type", "multimodal_message") {
		matchedLiteral = true
		var valueMultimodalMessageLiteral *MultimodalMessage
		if err := json.Unmarshal(data, &valueMultimodalMessageLiteral); err == nil {
			p.typ = "MultimodalMessage"
			p.MultimodalMessage = valueMultimodalMessageLiteral
			return nil
		}
	}
	if hasLiteral("type", "pong") {
		matchedLiteral = true
		var valueConversationPongLiteral *ConversationPong
		if err := json.Unmarshal(data, &valueConversationPongLiteral); err == nil {
			p.typ = "ConversationPong"
			p.ConversationPong = valueConversationPongLiteral
			return nil
		}
	}
	if hasLiteral("type", "client_tool_result") {
		matchedLiteral = true
		var valueClientToolResultLiteral *ClientToolResult
		if err := json.Unmarshal(data, &valueClientToolResultLiteral); err == nil {
			p.typ = "ClientToolResult"
			p.ClientToolResult = valueClientToolResultLiteral
			return nil
		}
	}
	if hasLiteral("type", "mcp_tool_approval_result") {
		matchedLiteral = true
		var valueMcpToolApprovalResultLiteral *McpToolApprovalResult
		if err := json.Unmarshal(data, &valueMcpToolApprovalResultLiteral); err == nil {
			p.typ = "McpToolApprovalResult"
			p.McpToolApprovalResult = valueMcpToolApprovalResultLiteral
			return nil
		}
	}
	if hasLiteral("type", "contextual_update") {
		matchedLiteral = true
		var valueContextualUpdateLiteral *ContextualUpdate
		if err := json.Unmarshal(data, &valueContextualUpdateLiteral); err == nil {
			p.typ = "ContextualUpdate"
			p.ContextualUpdate = valueContextualUpdateLiteral
			return nil
		}
	}
	if hasLiteral("type", "conversation_initiation_client_data") {
		matchedLiteral = true
		var valueConversationInitiationClientDataLiteral *ConversationInitiationClientData
		if err := json.Unmarshal(data, &valueConversationInitiationClientDataLiteral); err == nil {
			p.typ = "ConversationInitiationClientData"
			p.ConversationInitiationClientData = valueConversationInitiationClientDataLiteral
			return nil
		}
	}
	if matchedLiteral {
		return fmt.Errorf("%s cannot be deserialized as a %T", data, p)
	}
	var valueUserAudioChunk *UserAudioChunk
	if err := json.Unmarshal(data, &valueUserAudioChunk); err == nil {
		p.typ = "UserAudioChunk"
		p.UserAudioChunk = valueUserAudioChunk
		return nil
	}
	var valueFeedback *Feedback
	if err := json.Unmarshal(data, &valueFeedback); err == nil {
		p.typ = "Feedback"
		p.Feedback = valueFeedback
		return nil
	}
	if hasAnyKey("id", "type", "text", "user_identifier", "source_medium") {
		var valueUserMessage *UserMessage
		if err := json.Unmarshal(data, &valueUserMessage); err == nil {
			p.typ = "UserMessage"
			p.UserMessage = valueUserMessage
			return nil
		}
	}
	if hasAnyKey("type") {
		var valueUserActivity *UserActivity
		if err := json.Unmarshal(data, &valueUserActivity); err == nil {
			p.typ = "UserActivity"
			p.UserActivity = valueUserActivity
			return nil
		}
	}
	if hasAnyKey("type", "text", "file", "files") {
		var valueMultimodalMessage *MultimodalMessage
		if err := json.Unmarshal(data, &valueMultimodalMessage); err == nil {
			p.typ = "MultimodalMessage"
			p.MultimodalMessage = valueMultimodalMessage
			return nil
		}
	}
	var valueConversationPong *ConversationPong
	if err := json.Unmarshal(data, &valueConversationPong); err == nil {
		p.typ = "ConversationPong"
		p.ConversationPong = valueConversationPong
		return nil
	}
	var valueClientToolResult *ClientToolResult
	if err := json.Unmarshal(data, &valueClientToolResult); err == nil {
		p.typ = "ClientToolResult"
		p.ClientToolResult = valueClientToolResult
		return nil
	}
	var valueMcpToolApprovalResult *McpToolApprovalResult
	if err := json.Unmarshal(data, &valueMcpToolApprovalResult); err == nil {
		p.typ = "McpToolApprovalResult"
		p.McpToolApprovalResult = valueMcpToolApprovalResult
		return nil
	}
	var valueContextualUpdate *ContextualUpdate
	if err := json.Unmarshal(data, &valueContextualUpdate); err == nil {
		p.typ = "ContextualUpdate"
		p.ContextualUpdate = valueContextualUpdate
		return nil
	}
	if hasAnyKey("conversation_config_override", "custom_llm_extra_body", "user_id", "source_info", "branch_id", "environment", "starting_workflow_node_id", "procedure_ids", "start_procedure_ids", "dynamic_variables", "type") {
		var valueConversationInitiationClientData *ConversationInitiationClientData
		if err := json.Unmarshal(data, &valueConversationInitiationClientData); err == nil {
			p.typ = "ConversationInitiationClientData"
			p.ConversationInitiationClientData = valueConversationInitiationClientData
			return nil
		}
	}

	// Retried without their key guards, for a payload that matched none of them.
	var valueUserMessageFallback *UserMessage
	if err := json.Unmarshal(data, &valueUserMessageFallback); err == nil {
		p.typ = "UserMessage"
		p.UserMessage = valueUserMessageFallback
		return nil
	}
	var valueUserActivityFallback *UserActivity
	if err := json.Unmarshal(data, &valueUserActivityFallback); err == nil {
		p.typ = "UserActivity"
		p.UserActivity = valueUserActivityFallback
		return nil
	}
	var valueMultimodalMessageFallback *MultimodalMessage
	if err := json.Unmarshal(data, &valueMultimodalMessageFallback); err == nil {
		p.typ = "MultimodalMessage"
		p.MultimodalMessage = valueMultimodalMessageFallback
		return nil
	}
	var valueConversationInitiationClientDataFallback *ConversationInitiationClientData
	if err := json.Unmarshal(data, &valueConversationInitiationClientDataFallback); err == nil {
		p.typ = "ConversationInitiationClientData"
		p.ConversationInitiationClientData = valueConversationInitiationClientDataFallback
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, p)
}

func (p PublishEvent) MarshalJSON() ([]byte, error) {
	if p.typ == "UserAudioChunk" || p.UserAudioChunk != nil {
		return json.Marshal(p.UserAudioChunk)
	}
	if p.typ == "Feedback" || p.Feedback != nil {
		return json.Marshal(p.Feedback)
	}
	if p.typ == "UserMessage" || p.UserMessage != nil {
		return json.Marshal(p.UserMessage)
	}
	if p.typ == "UserActivity" || p.UserActivity != nil {
		return json.Marshal(p.UserActivity)
	}
	if p.typ == "MultimodalMessage" || p.MultimodalMessage != nil {
		return json.Marshal(p.MultimodalMessage)
	}
	if p.typ == "ConversationPong" || p.ConversationPong != nil {
		return json.Marshal(p.ConversationPong)
	}
	if p.typ == "ClientToolResult" || p.ClientToolResult != nil {
		return json.Marshal(p.ClientToolResult)
	}
	if p.typ == "McpToolApprovalResult" || p.McpToolApprovalResult != nil {
		return json.Marshal(p.McpToolApprovalResult)
	}
	if p.typ == "ContextualUpdate" || p.ContextualUpdate != nil {
		return json.Marshal(p.ContextualUpdate)
	}
	if p.typ == "ConversationInitiationClientData" || p.ConversationInitiationClientData != nil {
		return json.Marshal(p.ConversationInitiationClientData)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", p)
}

type PublishEventVisitor interface {
	VisitUserAudioChunk(*UserAudioChunk) error
	VisitFeedback(*Feedback) error
	VisitUserMessage(*UserMessage) error
	VisitUserActivity(*UserActivity) error
	VisitMultimodalMessage(*MultimodalMessage) error
	VisitConversationPong(*ConversationPong) error
	VisitClientToolResult(*ClientToolResult) error
	VisitMcpToolApprovalResult(*McpToolApprovalResult) error
	VisitContextualUpdate(*ContextualUpdate) error
	VisitConversationInitiationClientData(*ConversationInitiationClientData) error
}

func (p *PublishEvent) Accept(visitor PublishEventVisitor) error {
	if p.typ == "UserAudioChunk" || p.UserAudioChunk != nil {
		return visitor.VisitUserAudioChunk(p.UserAudioChunk)
	}
	if p.typ == "Feedback" || p.Feedback != nil {
		return visitor.VisitFeedback(p.Feedback)
	}
	if p.typ == "UserMessage" || p.UserMessage != nil {
		return visitor.VisitUserMessage(p.UserMessage)
	}
	if p.typ == "UserActivity" || p.UserActivity != nil {
		return visitor.VisitUserActivity(p.UserActivity)
	}
	if p.typ == "MultimodalMessage" || p.MultimodalMessage != nil {
		return visitor.VisitMultimodalMessage(p.MultimodalMessage)
	}
	if p.typ == "ConversationPong" || p.ConversationPong != nil {
		return visitor.VisitConversationPong(p.ConversationPong)
	}
	if p.typ == "ClientToolResult" || p.ClientToolResult != nil {
		return visitor.VisitClientToolResult(p.ClientToolResult)
	}
	if p.typ == "McpToolApprovalResult" || p.McpToolApprovalResult != nil {
		return visitor.VisitMcpToolApprovalResult(p.McpToolApprovalResult)
	}
	if p.typ == "ContextualUpdate" || p.ContextualUpdate != nil {
		return visitor.VisitContextualUpdate(p.ContextualUpdate)
	}
	if p.typ == "ConversationInitiationClientData" || p.ConversationInitiationClientData != nil {
		return visitor.VisitConversationInitiationClientData(p.ConversationInitiationClientData)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", p)
}
