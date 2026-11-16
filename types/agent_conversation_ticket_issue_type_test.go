package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestAgentConversationTicketIssueTypeVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant AgentConversationTicketIssueType
		wire    string
	}
	cases := []enumCase{
		{variant: AgentConversationTicketIssueTypeKnowledgeGap, wire: "\"knowledge_gap\""},
		{variant: AgentConversationTicketIssueTypeIncorrectInformation, wire: "\"incorrect_information\""},
		{variant: AgentConversationTicketIssueTypeDocumentationGap, wire: "\"documentation_gap\""},
		{variant: AgentConversationTicketIssueTypeProductFeedback, wire: "\"product_feedback\""},
		{variant: AgentConversationTicketIssueTypePlatformBug, wire: "\"platform_bug\""},
		{variant: AgentConversationTicketIssueTypeToolIssue, wire: "\"tool_issue\""},
		{variant: AgentConversationTicketIssueTypeMissingTool, wire: "\"missing_tool\""},
		{variant: AgentConversationTicketIssueTypeUnnecessaryEscalation, wire: "\"unnecessary_escalation\""},
		{variant: AgentConversationTicketIssueTypeWrongAction, wire: "\"wrong_action\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded AgentConversationTicketIssueType
			if err := json.Unmarshal([]byte(tc.wire), &decoded); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if decoded != tc.variant {
				t.Fatalf("decoded mismatch: got %#v want %#v", decoded, tc.variant)
			}
			var want any
			var got any
			if err := json.Unmarshal([]byte(tc.wire), &want); err != nil {
				t.Fatalf("unmarshal want: %v", err)
			}
			if err := json.Unmarshal(encoded, &got); err != nil {
				t.Fatalf("unmarshal got: %v", err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("round-trip mismatch\n got: %#v\nwant: %#v", got, want)
			}
		})
	}
}
