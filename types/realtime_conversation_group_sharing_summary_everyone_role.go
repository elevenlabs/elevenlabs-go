package types

import (
	"fmt"
)

type RealtimeConversationGroupSharingSummaryEveryoneRole string

const (
	RealtimeConversationGroupSharingSummaryEveryoneRoleAdmin     RealtimeConversationGroupSharingSummaryEveryoneRole = "admin"
	RealtimeConversationGroupSharingSummaryEveryoneRoleEditor    RealtimeConversationGroupSharingSummaryEveryoneRole = "editor"
	RealtimeConversationGroupSharingSummaryEveryoneRoleCommenter RealtimeConversationGroupSharingSummaryEveryoneRole = "commenter"
	RealtimeConversationGroupSharingSummaryEveryoneRoleViewer    RealtimeConversationGroupSharingSummaryEveryoneRole = "viewer"
)

func NewRealtimeConversationGroupSharingSummaryEveryoneRoleFromString(s string) (RealtimeConversationGroupSharingSummaryEveryoneRole, error) {
	switch s {
	case "admin":
		return RealtimeConversationGroupSharingSummaryEveryoneRoleAdmin, nil
	case "editor":
		return RealtimeConversationGroupSharingSummaryEveryoneRoleEditor, nil
	case "commenter":
		return RealtimeConversationGroupSharingSummaryEveryoneRoleCommenter, nil
	case "viewer":
		return RealtimeConversationGroupSharingSummaryEveryoneRoleViewer, nil
	}
	var t RealtimeConversationGroupSharingSummaryEveryoneRole
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e RealtimeConversationGroupSharingSummaryEveryoneRole) Ptr() *RealtimeConversationGroupSharingSummaryEveryoneRole {
	return &e
}
