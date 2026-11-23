package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentConversationTicketResponseModel struct {
	AgentqaTicketId string `json:"agentqa_ticket_id" xml:"agentqa_ticket_id" required:"true"`
	WorkspaceId     string `json:"workspace_id" xml:"workspace_id" required:"true"`
	OwnerUserId     string `json:"owner_user_id" xml:"owner_user_id" required:"true"`
	AgentId         string `json:"agent_id" xml:"agent_id" required:"true"`
	NeedsClustering bool   `json:"needs_clustering" xml:"needs_clustering" required:"true"`
	// One-line headline for the agent issue. None only on issues created before titles existed.
	Title *string `json:"title" xml:"title" required:"true"`
	// Always null. Issue types are labels, written as `issue_type:<value>`.
	IssueType         any                                  `json:"issue_type,omitempty" xml:"issue_type,omitempty"`
	Labels            []string                             `json:"labels" xml:"labels" required:"true"`
	ConversationIds   []string                             `json:"conversation_ids" xml:"conversation_ids" required:"true"`
	FirstSeenUnixSecs *int                                 `json:"first_seen_unix_secs" xml:"first_seen_unix_secs" required:"true"`
	LastSeenUnixSecs  *int                                 `json:"last_seen_unix_secs" xml:"last_seen_unix_secs" required:"true"`
	QaComment         *string                              `json:"qa_comment" xml:"qa_comment" required:"true"`
	TicketComments    []*TicketCommentResponseModel        `json:"ticket_comments" xml:"ticket_comments" required:"true"`
	TurnComments      []*TurnCommentResponseModel          `json:"turn_comments" xml:"turn_comments" required:"true"`
	Status            AgentConversationTicketStatus        `json:"status" xml:"status" required:"true"`
	Priority          *AgentConversationTicketPriority     `json:"priority" xml:"priority" required:"true"`
	PriorityChanges   []*TicketPriorityChangeResponseModel `json:"priority_changes" xml:"priority_changes" required:"true"`
	StatusChanges     []*TicketStatusChangeResponseModel   `json:"status_changes" xml:"status_changes" required:"true"`
	// The agent issue this one was merged into, set while its status is 'merged'.
	MergedIntoTicketId *string                       `json:"merged_into_ticket_id,omitempty" xml:"merged_into_ticket_id,omitempty"`
	Source             AgentConversationTicketSource `json:"source" xml:"source" required:"true"`
	AssigneeUserId     *string                       `json:"assignee_user_id" xml:"assignee_user_id" required:"true"`
	CreatedAtUnixSecs  int                           `json:"created_at_unix_secs" xml:"created_at_unix_secs" required:"true"`
	UpdatedAtUnixSecs  int                           `json:"updated_at_unix_secs" xml:"updated_at_unix_secs" required:"true"`
	// Where the list's `search` query matched. Set only on searched lists.
	SearchMatch     *TicketSearchMatchResponseModel `json:"search_match,omitempty" xml:"search_match,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AgentConversationTicketResponseModel) GetAgentqaTicketId() (value string) {
	if a == nil {
		return
	}
	return a.AgentqaTicketId
}

func (a *AgentConversationTicketResponseModel) GetWorkspaceId() (value string) {
	if a == nil {
		return
	}
	return a.WorkspaceId
}

func (a *AgentConversationTicketResponseModel) GetOwnerUserId() (value string) {
	if a == nil {
		return
	}
	return a.OwnerUserId
}

func (a *AgentConversationTicketResponseModel) GetAgentId() (value string) {
	if a == nil {
		return
	}
	return a.AgentId
}

func (a *AgentConversationTicketResponseModel) GetNeedsClustering() (value bool) {
	if a == nil {
		return
	}
	return a.NeedsClustering
}

func (a *AgentConversationTicketResponseModel) GetTitle() (value *string) {
	if a == nil {
		return
	}
	return a.Title
}

func (a *AgentConversationTicketResponseModel) GetIssueType() (value any) {
	if a == nil {
		return
	}
	return a.IssueType
}

func (a *AgentConversationTicketResponseModel) GetLabels() (value []string) {
	if a == nil {
		return
	}
	return a.Labels
}

func (a *AgentConversationTicketResponseModel) GetConversationIds() (value []string) {
	if a == nil {
		return
	}
	return a.ConversationIds
}

func (a *AgentConversationTicketResponseModel) GetFirstSeenUnixSecs() (value *int) {
	if a == nil {
		return
	}
	return a.FirstSeenUnixSecs
}

func (a *AgentConversationTicketResponseModel) GetLastSeenUnixSecs() (value *int) {
	if a == nil {
		return
	}
	return a.LastSeenUnixSecs
}

func (a *AgentConversationTicketResponseModel) GetQaComment() (value *string) {
	if a == nil {
		return
	}
	return a.QaComment
}

func (a *AgentConversationTicketResponseModel) GetTicketComments() (value []*TicketCommentResponseModel) {
	if a == nil {
		return
	}
	return a.TicketComments
}

func (a *AgentConversationTicketResponseModel) GetTurnComments() (value []*TurnCommentResponseModel) {
	if a == nil {
		return
	}
	return a.TurnComments
}

func (a *AgentConversationTicketResponseModel) GetStatus() (value AgentConversationTicketStatus) {
	if a == nil {
		return
	}
	return a.Status
}

func (a *AgentConversationTicketResponseModel) GetPriority() (value *AgentConversationTicketPriority) {
	if a == nil {
		return
	}
	return a.Priority
}

func (a *AgentConversationTicketResponseModel) GetPriorityChanges() (value []*TicketPriorityChangeResponseModel) {
	if a == nil {
		return
	}
	return a.PriorityChanges
}

func (a *AgentConversationTicketResponseModel) GetStatusChanges() (value []*TicketStatusChangeResponseModel) {
	if a == nil {
		return
	}
	return a.StatusChanges
}

func (a *AgentConversationTicketResponseModel) GetMergedIntoTicketId() (value *string) {
	if a == nil {
		return
	}
	return a.MergedIntoTicketId
}

func (a *AgentConversationTicketResponseModel) GetSource() (value AgentConversationTicketSource) {
	if a == nil {
		return
	}
	return a.Source
}

func (a *AgentConversationTicketResponseModel) GetAssigneeUserId() (value *string) {
	if a == nil {
		return
	}
	return a.AssigneeUserId
}

func (a *AgentConversationTicketResponseModel) GetCreatedAtUnixSecs() (value int) {
	if a == nil {
		return
	}
	return a.CreatedAtUnixSecs
}

func (a *AgentConversationTicketResponseModel) GetUpdatedAtUnixSecs() (value int) {
	if a == nil {
		return
	}
	return a.UpdatedAtUnixSecs
}

func (a *AgentConversationTicketResponseModel) GetSearchMatch() (value *TicketSearchMatchResponseModel) {
	if a == nil {
		return
	}
	return a.SearchMatch
}

func (a *AgentConversationTicketResponseModel) SetAgentqaTicketId(agentqaTicketId string) {
	a.AgentqaTicketId = agentqaTicketId
}

func (a *AgentConversationTicketResponseModel) SetWorkspaceId(workspaceId string) {
	a.WorkspaceId = workspaceId
}

func (a *AgentConversationTicketResponseModel) SetOwnerUserId(ownerUserId string) {
	a.OwnerUserId = ownerUserId
}

func (a *AgentConversationTicketResponseModel) SetAgentId(agentId string) {
	a.AgentId = agentId
}

func (a *AgentConversationTicketResponseModel) SetNeedsClustering(needsClustering bool) {
	a.NeedsClustering = needsClustering
}

func (a *AgentConversationTicketResponseModel) SetTitle(title *string) {
	a.Title = title
}

func (a *AgentConversationTicketResponseModel) SetIssueType(issueType any) {
	a.IssueType = issueType
}

func (a *AgentConversationTicketResponseModel) SetLabels(labels []string) {
	a.Labels = labels
}

func (a *AgentConversationTicketResponseModel) SetConversationIds(conversationIds []string) {
	a.ConversationIds = conversationIds
}

func (a *AgentConversationTicketResponseModel) SetFirstSeenUnixSecs(firstSeenUnixSecs *int) {
	a.FirstSeenUnixSecs = firstSeenUnixSecs
}

func (a *AgentConversationTicketResponseModel) SetLastSeenUnixSecs(lastSeenUnixSecs *int) {
	a.LastSeenUnixSecs = lastSeenUnixSecs
}

func (a *AgentConversationTicketResponseModel) SetQaComment(qaComment *string) {
	a.QaComment = qaComment
}

func (a *AgentConversationTicketResponseModel) SetTicketComments(ticketComments []*TicketCommentResponseModel) {
	a.TicketComments = ticketComments
}

func (a *AgentConversationTicketResponseModel) SetTurnComments(turnComments []*TurnCommentResponseModel) {
	a.TurnComments = turnComments
}

func (a *AgentConversationTicketResponseModel) SetStatus(status AgentConversationTicketStatus) {
	a.Status = status
}

func (a *AgentConversationTicketResponseModel) SetPriority(priority *AgentConversationTicketPriority) {
	a.Priority = priority
}

func (a *AgentConversationTicketResponseModel) SetPriorityChanges(priorityChanges []*TicketPriorityChangeResponseModel) {
	a.PriorityChanges = priorityChanges
}

func (a *AgentConversationTicketResponseModel) SetStatusChanges(statusChanges []*TicketStatusChangeResponseModel) {
	a.StatusChanges = statusChanges
}

func (a *AgentConversationTicketResponseModel) SetMergedIntoTicketId(mergedIntoTicketId *string) {
	a.MergedIntoTicketId = mergedIntoTicketId
}

func (a *AgentConversationTicketResponseModel) SetSource(source AgentConversationTicketSource) {
	a.Source = source
}

func (a *AgentConversationTicketResponseModel) SetAssigneeUserId(assigneeUserId *string) {
	a.AssigneeUserId = assigneeUserId
}

func (a *AgentConversationTicketResponseModel) SetCreatedAtUnixSecs(createdAtUnixSecs int) {
	a.CreatedAtUnixSecs = createdAtUnixSecs
}

func (a *AgentConversationTicketResponseModel) SetUpdatedAtUnixSecs(updatedAtUnixSecs int) {
	a.UpdatedAtUnixSecs = updatedAtUnixSecs
}

func (a *AgentConversationTicketResponseModel) SetSearchMatch(searchMatch *TicketSearchMatchResponseModel) {
	a.SearchMatch = searchMatch
}

func (a *AgentConversationTicketResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentConversationTicketResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentConversationTicketResponseModel to string"
	}
	return string(jsonData)
}

func (a *AgentConversationTicketResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentConversationTicketResponseModel) MarshalJSON() ([]byte, error) {
	type embed AgentConversationTicketResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
