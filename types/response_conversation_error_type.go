package types

import (
	"fmt"
)

// User-facing error types exposed on the public API.
type ResponseConversationErrorType string

const (
	ResponseConversationErrorTypeSystemError             ResponseConversationErrorType = "system_error"
	ResponseConversationErrorTypeCallInitializationError ResponseConversationErrorType = "call_initialization_error"
	ResponseConversationErrorTypeLineBusy                ResponseConversationErrorType = "line_busy"
	ResponseConversationErrorTypeNoAnswer                ResponseConversationErrorType = "no_answer"
	ResponseConversationErrorTypeCallRejected            ResponseConversationErrorType = "call_rejected"
	ResponseConversationErrorTypeBlockedByUser           ResponseConversationErrorType = "blocked_by_user"
	ResponseConversationErrorTypeAgentConfigurationError ResponseConversationErrorType = "agent_configuration_error"
	ResponseConversationErrorTypeInvalidClientRequest    ResponseConversationErrorType = "invalid_client_request"
	ResponseConversationErrorTypePermissionError         ResponseConversationErrorType = "permission_error"
	ResponseConversationErrorTypeEntitlementExceeded     ResponseConversationErrorType = "entitlement_exceeded"
	ResponseConversationErrorTypeClientDisconnected      ResponseConversationErrorType = "client_disconnected"
	ResponseConversationErrorTypeLlmError                ResponseConversationErrorType = "llm_error"
	ResponseConversationErrorTypeSpeechError             ResponseConversationErrorType = "speech_error"
	ResponseConversationErrorTypeToolError               ResponseConversationErrorType = "tool_error"
	ResponseConversationErrorTypeIntegrationError        ResponseConversationErrorType = "integration_error"
	ResponseConversationErrorTypeGuardrailTriggered      ResponseConversationErrorType = "guardrail_triggered"
	ResponseConversationErrorTypeSafetyViolation         ResponseConversationErrorType = "safety_violation"
	ResponseConversationErrorTypeMaxDurationExceeded     ResponseConversationErrorType = "max_duration_exceeded"
	ResponseConversationErrorTypePostProcessingError     ResponseConversationErrorType = "post_processing_error"
)

func NewResponseConversationErrorTypeFromString(s string) (ResponseConversationErrorType, error) {
	switch s {
	case "system_error":
		return ResponseConversationErrorTypeSystemError, nil
	case "call_initialization_error":
		return ResponseConversationErrorTypeCallInitializationError, nil
	case "line_busy":
		return ResponseConversationErrorTypeLineBusy, nil
	case "no_answer":
		return ResponseConversationErrorTypeNoAnswer, nil
	case "call_rejected":
		return ResponseConversationErrorTypeCallRejected, nil
	case "blocked_by_user":
		return ResponseConversationErrorTypeBlockedByUser, nil
	case "agent_configuration_error":
		return ResponseConversationErrorTypeAgentConfigurationError, nil
	case "invalid_client_request":
		return ResponseConversationErrorTypeInvalidClientRequest, nil
	case "permission_error":
		return ResponseConversationErrorTypePermissionError, nil
	case "entitlement_exceeded":
		return ResponseConversationErrorTypeEntitlementExceeded, nil
	case "client_disconnected":
		return ResponseConversationErrorTypeClientDisconnected, nil
	case "llm_error":
		return ResponseConversationErrorTypeLlmError, nil
	case "speech_error":
		return ResponseConversationErrorTypeSpeechError, nil
	case "tool_error":
		return ResponseConversationErrorTypeToolError, nil
	case "integration_error":
		return ResponseConversationErrorTypeIntegrationError, nil
	case "guardrail_triggered":
		return ResponseConversationErrorTypeGuardrailTriggered, nil
	case "safety_violation":
		return ResponseConversationErrorTypeSafetyViolation, nil
	case "max_duration_exceeded":
		return ResponseConversationErrorTypeMaxDurationExceeded, nil
	case "post_processing_error":
		return ResponseConversationErrorTypePostProcessingError, nil
	}
	var t ResponseConversationErrorType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ResponseConversationErrorType) Ptr() *ResponseConversationErrorType {
	return &e
}
