package conversations

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
)

type RunConversationEvaluationsRequest struct {
	// ID of the conversation
	ConversationId string `json:"-" url:"-"`
	// ID of the single evaluation criterion to rerun.
	EvaluationId string               `json:"evaluation_id" xml:"evaluation_id" required:"true"`
	Scope        *types.AnalysisScope `json:"scope,omitempty" xml:"scope,omitempty"`
	// ID of the agent whose evaluation result to rerun, for conversations that transferred between agents. If omitted, a conversation-scoped evaluation reruns for the agent that started the conversation.
	EvaluatedAgentId *string `json:"evaluated_agent_id,omitempty" xml:"evaluated_agent_id,omitempty"`
	// Branch of `evaluated_agent_id` whose evaluation result to rerun, for conversations that visited the agent on more than one branch. If omitted, the result reruns for every branch of that agent.
	EvaluatedBranchId *string `json:"evaluated_branch_id,omitempty" xml:"evaluated_branch_id,omitempty"`
}

func (r *RunConversationEvaluationsRequest) GetConversationId() (value string) {
	if r == nil {
		return
	}
	return r.ConversationId
}

func (r *RunConversationEvaluationsRequest) GetEvaluationId() (value string) {
	if r == nil {
		return
	}
	return r.EvaluationId
}

func (r *RunConversationEvaluationsRequest) GetScope() (value *types.AnalysisScope) {
	if r == nil {
		return
	}
	return r.Scope
}

func (r *RunConversationEvaluationsRequest) GetEvaluatedAgentId() (value *string) {
	if r == nil {
		return
	}
	return r.EvaluatedAgentId
}

func (r *RunConversationEvaluationsRequest) GetEvaluatedBranchId() (value *string) {
	if r == nil {
		return
	}
	return r.EvaluatedBranchId
}

func (r *RunConversationEvaluationsRequest) SetConversationId(conversationId string) {
	r.ConversationId = conversationId
}

func (r *RunConversationEvaluationsRequest) SetEvaluationId(evaluationId string) {
	r.EvaluationId = evaluationId
}

func (r *RunConversationEvaluationsRequest) SetScope(scope *types.AnalysisScope) {
	r.Scope = scope
}

func (r *RunConversationEvaluationsRequest) SetEvaluatedAgentId(evaluatedAgentId *string) {
	r.EvaluatedAgentId = evaluatedAgentId
}

func (r *RunConversationEvaluationsRequest) SetEvaluatedBranchId(evaluatedBranchId *string) {
	r.EvaluatedBranchId = evaluatedBranchId
}

func (r *RunConversationEvaluationsRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (r RunConversationEvaluationsRequest) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RunConversationEvaluationsRequest to string"
	}
	return string(jsonData)
}

func (r *RunConversationEvaluationsRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, r)
}

func (r *RunConversationEvaluationsRequest) MarshalJSON() ([]byte, error) {
	type embed RunConversationEvaluationsRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
