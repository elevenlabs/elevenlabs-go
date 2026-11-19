package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TestFromConversationMetadataInput struct {
	ConversationId     string                                           `json:"conversation_id" xml:"conversation_id" required:"true"`
	AgentId            string                                           `json:"agent_id" xml:"agent_id" required:"true"`
	BranchId           *string                                          `json:"branch_id,omitempty" xml:"branch_id,omitempty"`
	WorkflowNodeId     *string                                          `json:"workflow_node_id,omitempty" xml:"workflow_node_id,omitempty"`
	OriginalAgentReply []*ConversationHistoryTranscriptCommonModelInput `json:"original_agent_reply,omitempty" xml:"original_agent_reply,omitempty"`
	extraProperties    map[string]interface{}
}

func (t *TestFromConversationMetadataInput) GetConversationId() (value string) {
	if t == nil {
		return
	}
	return t.ConversationId
}

func (t *TestFromConversationMetadataInput) GetAgentId() (value string) {
	if t == nil {
		return
	}
	return t.AgentId
}

func (t *TestFromConversationMetadataInput) GetBranchId() (value *string) {
	if t == nil {
		return
	}
	return t.BranchId
}

func (t *TestFromConversationMetadataInput) GetWorkflowNodeId() (value *string) {
	if t == nil {
		return
	}
	return t.WorkflowNodeId
}

func (t *TestFromConversationMetadataInput) GetOriginalAgentReply() (value []*ConversationHistoryTranscriptCommonModelInput) {
	if t == nil {
		return
	}
	return t.OriginalAgentReply
}

func (t *TestFromConversationMetadataInput) SetConversationId(conversationId string) {
	t.ConversationId = conversationId
}

func (t *TestFromConversationMetadataInput) SetAgentId(agentId string) {
	t.AgentId = agentId
}

func (t *TestFromConversationMetadataInput) SetBranchId(branchId *string) {
	t.BranchId = branchId
}

func (t *TestFromConversationMetadataInput) SetWorkflowNodeId(workflowNodeId *string) {
	t.WorkflowNodeId = workflowNodeId
}

func (t *TestFromConversationMetadataInput) SetOriginalAgentReply(originalAgentReply []*ConversationHistoryTranscriptCommonModelInput) {
	t.OriginalAgentReply = originalAgentReply
}

func (t *TestFromConversationMetadataInput) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TestFromConversationMetadataInput) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TestFromConversationMetadataInput to string"
	}
	return string(jsonData)
}

func (t *TestFromConversationMetadataInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TestFromConversationMetadataInput) MarshalJSON() ([]byte, error) {
	type embed TestFromConversationMetadataInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
