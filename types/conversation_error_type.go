package types

import (
	"fmt"
)

// Stored in conversation history. Never exposed to customers directly.
type ConversationErrorType string

const (
	ConversationErrorTypeUnknown                 ConversationErrorType = "unknown"
	ConversationErrorTypeCallInitializationError ConversationErrorType = "call_initialization_error"
	ConversationErrorTypeLineBusy                ConversationErrorType = "line_busy"
	ConversationErrorTypeNoAnswer                ConversationErrorType = "no_answer"
	ConversationErrorTypeCallRejected            ConversationErrorType = "call_rejected"
	ConversationErrorTypeBlockedByUser           ConversationErrorType = "blocked_by_user"
	ConversationErrorTypeAgentConfigurationError ConversationErrorType = "agent_configuration_error"
	ConversationErrorTypeInvalidClientRequest    ConversationErrorType = "invalid_client_request"
	ConversationErrorTypePermissionError         ConversationErrorType = "permission_error"
	ConversationErrorTypeEntitlementExceeded     ConversationErrorType = "entitlement_exceeded"
	ConversationErrorTypeClientDisconnected      ConversationErrorType = "client_disconnected"
	ConversationErrorTypeLlmError                ConversationErrorType = "llm_error"
	ConversationErrorTypeSpeechError             ConversationErrorType = "speech_error"
	ConversationErrorTypeToolError               ConversationErrorType = "tool_error"
	ConversationErrorTypeIntegrationError        ConversationErrorType = "integration_error"
	ConversationErrorTypeDependencyError         ConversationErrorType = "dependency_error"
	ConversationErrorTypeGuardrailTriggered      ConversationErrorType = "guardrail_triggered"
	ConversationErrorTypeSafetyViolation         ConversationErrorType = "safety_violation"
	ConversationErrorTypeMaxDurationExceeded     ConversationErrorType = "max_duration_exceeded"
	ConversationErrorTypeWorkflowEndFailure      ConversationErrorType = "workflow_end_failure"
	ConversationErrorTypeInternalError           ConversationErrorType = "internal_error"
	ConversationErrorTypePostProcessingError     ConversationErrorType = "post_processing_error"
)

func NewConversationErrorTypeFromString(s string) (ConversationErrorType, error) {
	switch s {
	case "unknown":
		return ConversationErrorTypeUnknown, nil
	case "call_initialization_error":
		return ConversationErrorTypeCallInitializationError, nil
	case "line_busy":
		return ConversationErrorTypeLineBusy, nil
	case "no_answer":
		return ConversationErrorTypeNoAnswer, nil
	case "call_rejected":
		return ConversationErrorTypeCallRejected, nil
	case "blocked_by_user":
		return ConversationErrorTypeBlockedByUser, nil
	case "agent_configuration_error":
		return ConversationErrorTypeAgentConfigurationError, nil
	case "invalid_client_request":
		return ConversationErrorTypeInvalidClientRequest, nil
	case "permission_error":
		return ConversationErrorTypePermissionError, nil
	case "entitlement_exceeded":
		return ConversationErrorTypeEntitlementExceeded, nil
	case "client_disconnected":
		return ConversationErrorTypeClientDisconnected, nil
	case "llm_error":
		return ConversationErrorTypeLlmError, nil
	case "speech_error":
		return ConversationErrorTypeSpeechError, nil
	case "tool_error":
		return ConversationErrorTypeToolError, nil
	case "integration_error":
		return ConversationErrorTypeIntegrationError, nil
	case "dependency_error":
		return ConversationErrorTypeDependencyError, nil
	case "guardrail_triggered":
		return ConversationErrorTypeGuardrailTriggered, nil
	case "safety_violation":
		return ConversationErrorTypeSafetyViolation, nil
	case "max_duration_exceeded":
		return ConversationErrorTypeMaxDurationExceeded, nil
	case "workflow_end_failure":
		return ConversationErrorTypeWorkflowEndFailure, nil
	case "internal_error":
		return ConversationErrorTypeInternalError, nil
	case "post_processing_error":
		return ConversationErrorTypePostProcessingError, nil
	}
	var t ConversationErrorType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ConversationErrorType) Ptr() *ConversationErrorType {
	return &e
}
