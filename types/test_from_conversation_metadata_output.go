package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TestFromConversationMetadataOutput struct {
	ConversationId     string                                            `json:"conversation_id" xml:"conversation_id" required:"true"`
	AgentId            string                                            `json:"agent_id" xml:"agent_id" required:"true"`
	BranchId           *string                                           `json:"branch_id,omitempty" xml:"branch_id,omitempty"`
	WorkflowNodeId     *string                                           `json:"workflow_node_id,omitempty" xml:"workflow_node_id,omitempty"`
	OriginalAgentReply []*ConversationHistoryTranscriptCommonModelOutput `json:"original_agent_reply,omitempty" xml:"original_agent_reply,omitempty"`
	extraProperties    map[string]interface{}
}

func (t *TestFromConversationMetadataOutput) GetConversationId() (value string) {
	if t == nil {
		return
	}
	return t.ConversationId
}

func (t *TestFromConversationMetadataOutput) GetAgentId() (value string) {
	if t == nil {
		return
	}
	return t.AgentId
}

func (t *TestFromConversationMetadataOutput) GetBranchId() (value *string) {
	if t == nil {
		return
	}
	return t.BranchId
}

func (t *TestFromConversationMetadataOutput) GetWorkflowNodeId() (value *string) {
	if t == nil {
		return
	}
	return t.WorkflowNodeId
}

func (t *TestFromConversationMetadataOutput) GetOriginalAgentReply() (value []*ConversationHistoryTranscriptCommonModelOutput) {
	if t == nil {
		return
	}
	return t.OriginalAgentReply
}

func (t *TestFromConversationMetadataOutput) SetConversationId(conversationId string) {
	t.ConversationId = conversationId
}

func (t *TestFromConversationMetadataOutput) SetAgentId(agentId string) {
	t.AgentId = agentId
}

func (t *TestFromConversationMetadataOutput) SetBranchId(branchId *string) {
	t.BranchId = branchId
}

func (t *TestFromConversationMetadataOutput) SetWorkflowNodeId(workflowNodeId *string) {
	t.WorkflowNodeId = workflowNodeId
}

func (t *TestFromConversationMetadataOutput) SetOriginalAgentReply(originalAgentReply []*ConversationHistoryTranscriptCommonModelOutput) {
	t.OriginalAgentReply = originalAgentReply
}

func (t *TestFromConversationMetadataOutput) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TestFromConversationMetadataOutput) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TestFromConversationMetadataOutput to string"
	}
	return string(jsonData)
}

func (t *TestFromConversationMetadataOutput) UnmarshalJSON(data []byte) error {
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

func (t *TestFromConversationMetadataOutput) MarshalJSON() ([]byte, error) {
	type embed TestFromConversationMetadataOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
