package types

import (
	"fmt"
)

type ConversationHistoryTranscriptOtherToolsResultCommonModelType string

const (
	ConversationHistoryTranscriptOtherToolsResultCommonModelTypeClient  ConversationHistoryTranscriptOtherToolsResultCommonModelType = "client"
	ConversationHistoryTranscriptOtherToolsResultCommonModelTypeWebhook ConversationHistoryTranscriptOtherToolsResultCommonModelType = "webhook"
	ConversationHistoryTranscriptOtherToolsResultCommonModelTypeMcp     ConversationHistoryTranscriptOtherToolsResultCommonModelType = "mcp"
	ConversationHistoryTranscriptOtherToolsResultCommonModelTypeCode    ConversationHistoryTranscriptOtherToolsResultCommonModelType = "code"
)

func NewConversationHistoryTranscriptOtherToolsResultCommonModelTypeFromString(s string) (ConversationHistoryTranscriptOtherToolsResultCommonModelType, error) {
	switch s {
	case "client":
		return ConversationHistoryTranscriptOtherToolsResultCommonModelTypeClient, nil
	case "webhook":
		return ConversationHistoryTranscriptOtherToolsResultCommonModelTypeWebhook, nil
	case "mcp":
		return ConversationHistoryTranscriptOtherToolsResultCommonModelTypeMcp, nil
	case "code":
		return ConversationHistoryTranscriptOtherToolsResultCommonModelTypeCode, nil
	}
	var t ConversationHistoryTranscriptOtherToolsResultCommonModelType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ConversationHistoryTranscriptOtherToolsResultCommonModelType) Ptr() *ConversationHistoryTranscriptOtherToolsResultCommonModelType {
	return &e
}
