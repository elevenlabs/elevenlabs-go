package types

import (
	"fmt"
)

type AgentConversationTicketIssueType string

const (
	AgentConversationTicketIssueTypeKnowledgeGap          AgentConversationTicketIssueType = "knowledge_gap"
	AgentConversationTicketIssueTypeIncorrectInformation  AgentConversationTicketIssueType = "incorrect_information"
	AgentConversationTicketIssueTypeDocumentationGap      AgentConversationTicketIssueType = "documentation_gap"
	AgentConversationTicketIssueTypeProductFeedback       AgentConversationTicketIssueType = "product_feedback"
	AgentConversationTicketIssueTypePlatformBug           AgentConversationTicketIssueType = "platform_bug"
	AgentConversationTicketIssueTypeToolIssue             AgentConversationTicketIssueType = "tool_issue"
	AgentConversationTicketIssueTypeMissingTool           AgentConversationTicketIssueType = "missing_tool"
	AgentConversationTicketIssueTypeUnnecessaryEscalation AgentConversationTicketIssueType = "unnecessary_escalation"
	AgentConversationTicketIssueTypeWrongAction           AgentConversationTicketIssueType = "wrong_action"
)

func NewAgentConversationTicketIssueTypeFromString(s string) (AgentConversationTicketIssueType, error) {
	switch s {
	case "knowledge_gap":
		return AgentConversationTicketIssueTypeKnowledgeGap, nil
	case "incorrect_information":
		return AgentConversationTicketIssueTypeIncorrectInformation, nil
	case "documentation_gap":
		return AgentConversationTicketIssueTypeDocumentationGap, nil
	case "product_feedback":
		return AgentConversationTicketIssueTypeProductFeedback, nil
	case "platform_bug":
		return AgentConversationTicketIssueTypePlatformBug, nil
	case "tool_issue":
		return AgentConversationTicketIssueTypeToolIssue, nil
	case "missing_tool":
		return AgentConversationTicketIssueTypeMissingTool, nil
	case "unnecessary_escalation":
		return AgentConversationTicketIssueTypeUnnecessaryEscalation, nil
	case "wrong_action":
		return AgentConversationTicketIssueTypeWrongAction, nil
	}
	var t AgentConversationTicketIssueType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AgentConversationTicketIssueType) Ptr() *AgentConversationTicketIssueType {
	return &e
}
