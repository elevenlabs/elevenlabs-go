package agents

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	types "github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

type CreateManualTicketRequestModel struct {
	AgentId string `json:"-" url:"-"`
	// One-line headline shown in the issues list. Defaults to one derived from qa_comment.
	Title *string `json:"title,omitempty" xml:"title,omitempty" maxLength:"120" minLength:"1"`
	// What the agent issue is about, e.g. a follow-up task for the agent.
	QaComment string `json:"qa_comment" xml:"qa_comment" required:"true" maxLength:"10000" minLength:"1"`
	// How urgently the agent issue needs attention.
	Priority *types.AgentConversationTicketPriority `json:"priority,omitempty" xml:"priority,omitempty"`
}

func (c *CreateManualTicketRequestModel) GetAgentId() (value string) {
	if c == nil {
		return
	}
	return c.AgentId
}

func (c *CreateManualTicketRequestModel) GetTitle() (value *string) {
	if c == nil {
		return
	}
	return c.Title
}

func (c *CreateManualTicketRequestModel) GetQaComment() (value string) {
	if c == nil {
		return
	}
	return c.QaComment
}

func (c *CreateManualTicketRequestModel) GetPriority() (value *types.AgentConversationTicketPriority) {
	if c == nil {
		return
	}
	return c.Priority
}

func (c *CreateManualTicketRequestModel) SetAgentId(agentId string) {
	c.AgentId = agentId
}

func (c *CreateManualTicketRequestModel) SetTitle(title *string) {
	c.Title = title
}

func (c *CreateManualTicketRequestModel) SetQaComment(qaComment string) {
	c.QaComment = qaComment
}

func (c *CreateManualTicketRequestModel) SetPriority(priority *types.AgentConversationTicketPriority) {
	c.Priority = priority
}

func (c *CreateManualTicketRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (c CreateManualTicketRequestModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateManualTicketRequestModel to string"
	}
	return string(jsonData)
}

func (c *CreateManualTicketRequestModel) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, c)
}

func (c *CreateManualTicketRequestModel) MarshalJSON() ([]byte, error) {
	type embed CreateManualTicketRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}

type CreateAgentConversationTicketRequestModel struct {
	// Conversation this agent issue is about.
	ConversationId string `json:"conversation_id" xml:"conversation_id" required:"true"`
	// One-line headline shown in the issues list. Defaults to one derived from the comments, falling back to the conversation's summary title. Ignored when the comment is added to the conversation's open issue.
	Title *string `json:"title,omitempty" xml:"title,omitempty" maxLength:"120" minLength:"1"`
	// The problem this agent issue is about, covering the whole conversation rather than a single turn.
	QaComment *string `json:"qa_comment,omitempty" xml:"qa_comment,omitempty" maxLength:"10000" minLength:"1"`
	// Optional turn-level comments on what went wrong.
	TurnComments []*types.TurnCommentRequestModel `json:"turn_comments,omitempty" xml:"turn_comments,omitempty" maxItems:"200"`
	// How urgently the agent issue needs attention. If the conversation already has an open issue, it is raised to this priority when lower.
	Priority *types.AgentConversationTicketPriority `json:"priority,omitempty" xml:"priority,omitempty"`
}

func (c *CreateAgentConversationTicketRequestModel) GetConversationId() (value string) {
	if c == nil {
		return
	}
	return c.ConversationId
}

func (c *CreateAgentConversationTicketRequestModel) GetTitle() (value *string) {
	if c == nil {
		return
	}
	return c.Title
}

func (c *CreateAgentConversationTicketRequestModel) GetQaComment() (value *string) {
	if c == nil {
		return
	}
	return c.QaComment
}

func (c *CreateAgentConversationTicketRequestModel) GetTurnComments() (value []*types.TurnCommentRequestModel) {
	if c == nil {
		return
	}
	return c.TurnComments
}

func (c *CreateAgentConversationTicketRequestModel) GetPriority() (value *types.AgentConversationTicketPriority) {
	if c == nil {
		return
	}
	return c.Priority
}

func (c *CreateAgentConversationTicketRequestModel) SetConversationId(conversationId string) {
	c.ConversationId = conversationId
}

func (c *CreateAgentConversationTicketRequestModel) SetTitle(title *string) {
	c.Title = title
}

func (c *CreateAgentConversationTicketRequestModel) SetQaComment(qaComment *string) {
	c.QaComment = qaComment
}

func (c *CreateAgentConversationTicketRequestModel) SetTurnComments(turnComments []*types.TurnCommentRequestModel) {
	c.TurnComments = turnComments
}

func (c *CreateAgentConversationTicketRequestModel) SetPriority(priority *types.AgentConversationTicketPriority) {
	c.Priority = priority
}

func (c *CreateAgentConversationTicketRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (c CreateAgentConversationTicketRequestModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateAgentConversationTicketRequestModel to string"
	}
	return string(jsonData)
}

func (c *CreateAgentConversationTicketRequestModel) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, c)
}

func (c *CreateAgentConversationTicketRequestModel) MarshalJSON() ([]byte, error) {
	type embed CreateAgentConversationTicketRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}

type PatchAgentConversationTicketRequestModel struct {
	AgentqaTicketId string `json:"-" url:"-"`
	// If provided, updates the agent issue's title. Omit to leave unchanged.
	Title *string `json:"title,omitempty" xml:"title,omitempty" maxLength:"120" minLength:"1"`
	// If provided, updates the agent issue's description. Omit to leave unchanged.
	QaComment *string `json:"qa_comment,omitempty" xml:"qa_comment,omitempty" maxLength:"10000" minLength:"1"`
	// If provided, updates the agent issue's status. Omit to leave unchanged.
	Status *types.AgentConversationTicketStatus `json:"status,omitempty" xml:"status,omitempty"`
	// If provided, updates who is responsible for resolving this agent issue. Must be a workspace member with at least viewer access to the agent. Pass null to unassign. Omit to leave unchanged.
	AssigneeUserId *string `json:"assignee_user_id,omitempty" xml:"assignee_user_id,omitempty"`
	// If provided, updates how urgently the agent issue needs attention. Pass null to clear it. Omit to leave unchanged.
	Priority *types.AgentConversationTicketPriority `json:"priority,omitempty" xml:"priority,omitempty"`
}

func (p *PatchAgentConversationTicketRequestModel) GetAgentqaTicketId() (value string) {
	if p == nil {
		return
	}
	return p.AgentqaTicketId
}

func (p *PatchAgentConversationTicketRequestModel) GetTitle() (value *string) {
	if p == nil {
		return
	}
	return p.Title
}

func (p *PatchAgentConversationTicketRequestModel) GetQaComment() (value *string) {
	if p == nil {
		return
	}
	return p.QaComment
}

func (p *PatchAgentConversationTicketRequestModel) GetStatus() (value *types.AgentConversationTicketStatus) {
	if p == nil {
		return
	}
	return p.Status
}

func (p *PatchAgentConversationTicketRequestModel) GetAssigneeUserId() (value *string) {
	if p == nil {
		return
	}
	return p.AssigneeUserId
}

func (p *PatchAgentConversationTicketRequestModel) GetPriority() (value *types.AgentConversationTicketPriority) {
	if p == nil {
		return
	}
	return p.Priority
}

func (p *PatchAgentConversationTicketRequestModel) SetAgentqaTicketId(agentqaTicketId string) {
	p.AgentqaTicketId = agentqaTicketId
}

func (p *PatchAgentConversationTicketRequestModel) SetTitle(title *string) {
	p.Title = title
}

func (p *PatchAgentConversationTicketRequestModel) SetQaComment(qaComment *string) {
	p.QaComment = qaComment
}

func (p *PatchAgentConversationTicketRequestModel) SetStatus(status *types.AgentConversationTicketStatus) {
	p.Status = status
}

func (p *PatchAgentConversationTicketRequestModel) SetAssigneeUserId(assigneeUserId *string) {
	p.AssigneeUserId = assigneeUserId
}

func (p *PatchAgentConversationTicketRequestModel) SetPriority(priority *types.AgentConversationTicketPriority) {
	p.Priority = priority
}

func (p *PatchAgentConversationTicketRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (p PatchAgentConversationTicketRequestModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PatchAgentConversationTicketRequestModel to string"
	}
	return string(jsonData)
}

func (p *PatchAgentConversationTicketRequestModel) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, p)
}

func (p *PatchAgentConversationTicketRequestModel) MarshalJSON() ([]byte, error) {
	type embed PatchAgentConversationTicketRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}

type AddTicketCommentRequestModel struct {
	AgentqaTicketId string `json:"-" url:"-"`
	// A comment discussing how to resolve the agent issue.
	Comment string `json:"comment" xml:"comment" required:"true" maxLength:"10000" minLength:"1"`
}

func (a *AddTicketCommentRequestModel) GetAgentqaTicketId() (value string) {
	if a == nil {
		return
	}
	return a.AgentqaTicketId
}

func (a *AddTicketCommentRequestModel) GetComment() (value string) {
	if a == nil {
		return
	}
	return a.Comment
}

func (a *AddTicketCommentRequestModel) SetAgentqaTicketId(agentqaTicketId string) {
	a.AgentqaTicketId = agentqaTicketId
}

func (a *AddTicketCommentRequestModel) SetComment(comment string) {
	a.Comment = comment
}

func (a *AddTicketCommentRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (a AddTicketCommentRequestModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddTicketCommentRequestModel to string"
	}
	return string(jsonData)
}

func (a *AddTicketCommentRequestModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AddTicketCommentRequestModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AddTicketCommentRequestModel(tmp)
	return nil
}

func (a *AddTicketCommentRequestModel) MarshalJSON() ([]byte, error) {
	type embed AddTicketCommentRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}

type MergeAgentConversationTicketsRequestModel struct {
	AgentqaTicketId string `json:"-" url:"-"`
	// Open issues of the same agent to fold into this one. They are kept with status 'merged', pointing at this issue.
	SourceTicketIds []string `json:"source_ticket_ids" xml:"source_ticket_ids" required:"true" minItems:"1" maxItems:"50"`
}

func (m *MergeAgentConversationTicketsRequestModel) GetAgentqaTicketId() (value string) {
	if m == nil {
		return
	}
	return m.AgentqaTicketId
}

func (m *MergeAgentConversationTicketsRequestModel) GetSourceTicketIds() (value []string) {
	if m == nil {
		return
	}
	return m.SourceTicketIds
}

func (m *MergeAgentConversationTicketsRequestModel) SetAgentqaTicketId(agentqaTicketId string) {
	m.AgentqaTicketId = agentqaTicketId
}

func (m *MergeAgentConversationTicketsRequestModel) SetSourceTicketIds(sourceTicketIds []string) {
	m.SourceTicketIds = sourceTicketIds
}

func (m *MergeAgentConversationTicketsRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (m MergeAgentConversationTicketsRequestModel) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MergeAgentConversationTicketsRequestModel to string"
	}
	return string(jsonData)
}

func (m *MergeAgentConversationTicketsRequestModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, m); err != nil {
		return err
	}
	type alias MergeAgentConversationTicketsRequestModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*m = MergeAgentConversationTicketsRequestModel(tmp)
	return nil
}

func (m *MergeAgentConversationTicketsRequestModel) MarshalJSON() ([]byte, error) {
	type embed MergeAgentConversationTicketsRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}

type AddTurnCommentRequestModel struct {
	AgentqaTicketId string `json:"-" url:"-"`
	// Zero-based index of the transcript turn this comment refers to.
	TurnIndex int `json:"turn_index" xml:"turn_index" required:"true" min:"0"`
	// What went wrong at this turn.
	Comment string `json:"comment" xml:"comment" required:"true" maxLength:"10000" minLength:"1"`
}

func (a *AddTurnCommentRequestModel) GetAgentqaTicketId() (value string) {
	if a == nil {
		return
	}
	return a.AgentqaTicketId
}

func (a *AddTurnCommentRequestModel) GetTurnIndex() (value int) {
	if a == nil {
		return
	}
	return a.TurnIndex
}

func (a *AddTurnCommentRequestModel) GetComment() (value string) {
	if a == nil {
		return
	}
	return a.Comment
}

func (a *AddTurnCommentRequestModel) SetAgentqaTicketId(agentqaTicketId string) {
	a.AgentqaTicketId = agentqaTicketId
}

func (a *AddTurnCommentRequestModel) SetTurnIndex(turnIndex int) {
	a.TurnIndex = turnIndex
}

func (a *AddTurnCommentRequestModel) SetComment(comment string) {
	a.Comment = comment
}

func (a *AddTurnCommentRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (a AddTurnCommentRequestModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddTurnCommentRequestModel to string"
	}
	return string(jsonData)
}

func (a *AddTurnCommentRequestModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AddTurnCommentRequestModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AddTurnCommentRequestModel(tmp)
	return nil
}

func (a *AddTurnCommentRequestModel) MarshalJSON() ([]byte, error) {
	type embed AddTurnCommentRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}

var (
	listTriageTicketsRequestFieldPageSize           = big.NewInt(1 << 0)
	listTriageTicketsRequestFieldConversationId     = big.NewInt(1 << 1)
	listTriageTicketsRequestFieldStatus             = big.NewInt(1 << 2)
	listTriageTicketsRequestFieldSources            = big.NewInt(1 << 3)
	listTriageTicketsRequestFieldPriorities         = big.NewInt(1 << 4)
	listTriageTicketsRequestFieldSortBy             = big.NewInt(1 << 5)
	listTriageTicketsRequestFieldOwnerUserId        = big.NewInt(1 << 6)
	listTriageTicketsRequestFieldAssigneeUserId     = big.NewInt(1 << 7)
	listTriageTicketsRequestFieldIssueType          = big.NewInt(1 << 8)
	listTriageTicketsRequestFieldLabel              = big.NewInt(1 << 9)
	listTriageTicketsRequestFieldMergedIntoTicketId = big.NewInt(1 << 10)
	listTriageTicketsRequestFieldSearch             = big.NewInt(1 << 11)
	listTriageTicketsRequestFieldCursor             = big.NewInt(1 << 12)
)

type ListTriageTicketsRequest struct {
	PageSize           *int                                     `json:"-" url:"page_size,omitempty"`
	ConversationId     *string                                  `json:"-" url:"conversation_id,omitempty"`
	Status             *types.AgentConversationTicketStatus     `json:"-" url:"status,omitempty"`
	Sources            []*types.AgentConversationTicketSource   `json:"-" url:"sources,omitempty"`
	Priorities         []*types.AgentConversationTicketPriority `json:"-" url:"priorities,omitempty"`
	SortBy             *types.AgentConversationTicketSortBy     `json:"-" url:"sort_by,omitempty"`
	OwnerUserId        *string                                  `json:"-" url:"owner_user_id,omitempty"`
	AssigneeUserId     *string                                  `json:"-" url:"assignee_user_id,omitempty"`
	IssueType          *types.AgentConversationTicketIssueType  `json:"-" url:"issue_type,omitempty"`
	Label              *string                                  `json:"-" url:"label,omitempty"`
	MergedIntoTicketId *string                                  `json:"-" url:"merged_into_ticket_id,omitempty"`
	Search             *string                                  `json:"-" url:"search,omitempty"`
	Cursor             *string                                  `json:"-" url:"cursor,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListTriageTicketsRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListTriageTicketsRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listTriageTicketsRequestFieldPageSize)
}

// SetConversationId sets the ConversationId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListTriageTicketsRequest) SetConversationId(conversationId *string) {
	l.ConversationId = conversationId
	l.require(listTriageTicketsRequestFieldConversationId)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListTriageTicketsRequest) SetStatus(status *types.AgentConversationTicketStatus) {
	l.Status = status
	l.require(listTriageTicketsRequestFieldStatus)
}

// SetSources sets the Sources field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListTriageTicketsRequest) SetSources(sources []*types.AgentConversationTicketSource) {
	l.Sources = sources
	l.require(listTriageTicketsRequestFieldSources)
}

// SetPriorities sets the Priorities field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListTriageTicketsRequest) SetPriorities(priorities []*types.AgentConversationTicketPriority) {
	l.Priorities = priorities
	l.require(listTriageTicketsRequestFieldPriorities)
}

// SetSortBy sets the SortBy field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListTriageTicketsRequest) SetSortBy(sortBy *types.AgentConversationTicketSortBy) {
	l.SortBy = sortBy
	l.require(listTriageTicketsRequestFieldSortBy)
}

// SetOwnerUserId sets the OwnerUserId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListTriageTicketsRequest) SetOwnerUserId(ownerUserId *string) {
	l.OwnerUserId = ownerUserId
	l.require(listTriageTicketsRequestFieldOwnerUserId)
}

// SetAssigneeUserId sets the AssigneeUserId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListTriageTicketsRequest) SetAssigneeUserId(assigneeUserId *string) {
	l.AssigneeUserId = assigneeUserId
	l.require(listTriageTicketsRequestFieldAssigneeUserId)
}

// SetIssueType sets the IssueType field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListTriageTicketsRequest) SetIssueType(issueType *types.AgentConversationTicketIssueType) {
	l.IssueType = issueType
	l.require(listTriageTicketsRequestFieldIssueType)
}

// SetLabel sets the Label field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListTriageTicketsRequest) SetLabel(label *string) {
	l.Label = label
	l.require(listTriageTicketsRequestFieldLabel)
}

// SetMergedIntoTicketId sets the MergedIntoTicketId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListTriageTicketsRequest) SetMergedIntoTicketId(mergedIntoTicketId *string) {
	l.MergedIntoTicketId = mergedIntoTicketId
	l.require(listTriageTicketsRequestFieldMergedIntoTicketId)
}

// SetSearch sets the Search field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListTriageTicketsRequest) SetSearch(search *string) {
	l.Search = search
	l.require(listTriageTicketsRequestFieldSearch)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListTriageTicketsRequest) SetCursor(cursor *string) {
	l.Cursor = cursor
	l.require(listTriageTicketsRequestFieldCursor)
}

var (
	listForWorkspaceTriageTicketsRequestFieldPageSize           = big.NewInt(1 << 0)
	listForWorkspaceTriageTicketsRequestFieldAgentId            = big.NewInt(1 << 1)
	listForWorkspaceTriageTicketsRequestFieldStatus             = big.NewInt(1 << 2)
	listForWorkspaceTriageTicketsRequestFieldSources            = big.NewInt(1 << 3)
	listForWorkspaceTriageTicketsRequestFieldPriorities         = big.NewInt(1 << 4)
	listForWorkspaceTriageTicketsRequestFieldSortBy             = big.NewInt(1 << 5)
	listForWorkspaceTriageTicketsRequestFieldOwnerUserId        = big.NewInt(1 << 6)
	listForWorkspaceTriageTicketsRequestFieldAssigneeUserId     = big.NewInt(1 << 7)
	listForWorkspaceTriageTicketsRequestFieldIssueType          = big.NewInt(1 << 8)
	listForWorkspaceTriageTicketsRequestFieldLabel              = big.NewInt(1 << 9)
	listForWorkspaceTriageTicketsRequestFieldMergedIntoTicketId = big.NewInt(1 << 10)
	listForWorkspaceTriageTicketsRequestFieldSearch             = big.NewInt(1 << 11)
	listForWorkspaceTriageTicketsRequestFieldCursor             = big.NewInt(1 << 12)
)

type ListForWorkspaceTriageTicketsRequest struct {
	PageSize           *int                                     `json:"-" url:"page_size,omitempty"`
	AgentId            *string                                  `json:"-" url:"agent_id,omitempty"`
	Status             *types.AgentConversationTicketStatus     `json:"-" url:"status,omitempty"`
	Sources            []*types.AgentConversationTicketSource   `json:"-" url:"sources,omitempty"`
	Priorities         []*types.AgentConversationTicketPriority `json:"-" url:"priorities,omitempty"`
	SortBy             *types.AgentConversationTicketSortBy     `json:"-" url:"sort_by,omitempty"`
	OwnerUserId        *string                                  `json:"-" url:"owner_user_id,omitempty"`
	AssigneeUserId     *string                                  `json:"-" url:"assignee_user_id,omitempty"`
	IssueType          *types.AgentConversationTicketIssueType  `json:"-" url:"issue_type,omitempty"`
	Label              *string                                  `json:"-" url:"label,omitempty"`
	MergedIntoTicketId *string                                  `json:"-" url:"merged_into_ticket_id,omitempty"`
	Search             *string                                  `json:"-" url:"search,omitempty"`
	Cursor             *string                                  `json:"-" url:"cursor,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListForWorkspaceTriageTicketsRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListForWorkspaceTriageTicketsRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listForWorkspaceTriageTicketsRequestFieldPageSize)
}

// SetAgentId sets the AgentId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListForWorkspaceTriageTicketsRequest) SetAgentId(agentId *string) {
	l.AgentId = agentId
	l.require(listForWorkspaceTriageTicketsRequestFieldAgentId)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListForWorkspaceTriageTicketsRequest) SetStatus(status *types.AgentConversationTicketStatus) {
	l.Status = status
	l.require(listForWorkspaceTriageTicketsRequestFieldStatus)
}

// SetSources sets the Sources field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListForWorkspaceTriageTicketsRequest) SetSources(sources []*types.AgentConversationTicketSource) {
	l.Sources = sources
	l.require(listForWorkspaceTriageTicketsRequestFieldSources)
}

// SetPriorities sets the Priorities field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListForWorkspaceTriageTicketsRequest) SetPriorities(priorities []*types.AgentConversationTicketPriority) {
	l.Priorities = priorities
	l.require(listForWorkspaceTriageTicketsRequestFieldPriorities)
}

// SetSortBy sets the SortBy field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListForWorkspaceTriageTicketsRequest) SetSortBy(sortBy *types.AgentConversationTicketSortBy) {
	l.SortBy = sortBy
	l.require(listForWorkspaceTriageTicketsRequestFieldSortBy)
}

// SetOwnerUserId sets the OwnerUserId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListForWorkspaceTriageTicketsRequest) SetOwnerUserId(ownerUserId *string) {
	l.OwnerUserId = ownerUserId
	l.require(listForWorkspaceTriageTicketsRequestFieldOwnerUserId)
}

// SetAssigneeUserId sets the AssigneeUserId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListForWorkspaceTriageTicketsRequest) SetAssigneeUserId(assigneeUserId *string) {
	l.AssigneeUserId = assigneeUserId
	l.require(listForWorkspaceTriageTicketsRequestFieldAssigneeUserId)
}

// SetIssueType sets the IssueType field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListForWorkspaceTriageTicketsRequest) SetIssueType(issueType *types.AgentConversationTicketIssueType) {
	l.IssueType = issueType
	l.require(listForWorkspaceTriageTicketsRequestFieldIssueType)
}

// SetLabel sets the Label field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListForWorkspaceTriageTicketsRequest) SetLabel(label *string) {
	l.Label = label
	l.require(listForWorkspaceTriageTicketsRequestFieldLabel)
}

// SetMergedIntoTicketId sets the MergedIntoTicketId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListForWorkspaceTriageTicketsRequest) SetMergedIntoTicketId(mergedIntoTicketId *string) {
	l.MergedIntoTicketId = mergedIntoTicketId
	l.require(listForWorkspaceTriageTicketsRequestFieldMergedIntoTicketId)
}

// SetSearch sets the Search field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListForWorkspaceTriageTicketsRequest) SetSearch(search *string) {
	l.Search = search
	l.require(listForWorkspaceTriageTicketsRequestFieldSearch)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListForWorkspaceTriageTicketsRequest) SetCursor(cursor *string) {
	l.Cursor = cursor
	l.require(listForWorkspaceTriageTicketsRequestFieldCursor)
}
