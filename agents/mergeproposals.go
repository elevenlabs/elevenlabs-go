package agents

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	types "github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

type BodyCreateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsPost struct {
	// The id of an agent. This is returned on agent creation.
	AgentId string `json:"-" url:"-"`
	// Branch whose changes should be merged.
	SourceBranchId string `json:"source_branch_id" xml:"source_branch_id" required:"true"`
	// Branch that should receive the changes.
	TargetBranchId string `json:"target_branch_id" xml:"target_branch_id" required:"true"`
	// Short title for the merge_proposal.
	Title string `json:"title" xml:"title" required:"true" maxLength:"140" minLength:"1"`
	// Optional longer description for reviewers.
	Description *string `json:"description,omitempty" xml:"description,omitempty" maxLength:"4096"`
	// User IDs to request a review from.
	RequestedReviewerUserIds []string `json:"requested_reviewer_user_ids,omitempty" xml:"requested_reviewer_user_ids,omitempty" maxItems:"100"`
	// Triage ticket of this agent that the merge_proposal resolves. Merging it resolves the ticket if still open.
	TriageTicketId *string `json:"triage_ticket_id,omitempty" xml:"triage_ticket_id,omitempty"`
}

func (b *BodyCreateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsPost) GetAgentId() (value string) {
	if b == nil {
		return
	}
	return b.AgentId
}

func (b *BodyCreateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsPost) GetSourceBranchId() (value string) {
	if b == nil {
		return
	}
	return b.SourceBranchId
}

func (b *BodyCreateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsPost) GetTargetBranchId() (value string) {
	if b == nil {
		return
	}
	return b.TargetBranchId
}

func (b *BodyCreateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsPost) GetTitle() (value string) {
	if b == nil {
		return
	}
	return b.Title
}

func (b *BodyCreateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsPost) GetDescription() (value *string) {
	if b == nil {
		return
	}
	return b.Description
}

func (b *BodyCreateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsPost) GetRequestedReviewerUserIds() (value []string) {
	if b == nil {
		return
	}
	return b.RequestedReviewerUserIds
}

func (b *BodyCreateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsPost) GetTriageTicketId() (value *string) {
	if b == nil {
		return
	}
	return b.TriageTicketId
}

func (b *BodyCreateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsPost) SetAgentId(agentId string) {
	b.AgentId = agentId
}

func (b *BodyCreateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsPost) SetSourceBranchId(sourceBranchId string) {
	b.SourceBranchId = sourceBranchId
}

func (b *BodyCreateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsPost) SetTargetBranchId(targetBranchId string) {
	b.TargetBranchId = targetBranchId
}

func (b *BodyCreateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsPost) SetTitle(title string) {
	b.Title = title
}

func (b *BodyCreateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsPost) SetDescription(description *string) {
	b.Description = description
}

func (b *BodyCreateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsPost) SetRequestedReviewerUserIds(requestedReviewerUserIds []string) {
	b.RequestedReviewerUserIds = requestedReviewerUserIds
}

func (b *BodyCreateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsPost) SetTriageTicketId(triageTicketId *string) {
	b.TriageTicketId = triageTicketId
}

func (b *BodyCreateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyCreateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyCreateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsPost to string"
	}
	return string(jsonData)
}

func (b *BodyCreateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyCreateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsPost) MarshalJSON() ([]byte, error) {
	type embed BodyCreateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

type BodyUpdateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdPatch struct {
	// Unique identifier for the merge_proposal.
	MergeProposalId string `json:"-" url:"-"`
	// The id of an agent. This is returned on agent creation.
	AgentId string `json:"-" url:"-"`
	// New title for the merge_proposal.
	Title *string `json:"title,omitempty" xml:"title,omitempty" maxLength:"140" minLength:"1"`
	// New description for the merge_proposal.
	Description *string `json:"description,omitempty" xml:"description,omitempty" maxLength:"4096"`
	// Replacement list of user IDs to request a review from.
	RequestedReviewerUserIds []string `json:"requested_reviewer_user_ids,omitempty" xml:"requested_reviewer_user_ids,omitempty" maxItems:"100"`
	// When true, close the merge_proposal without merging.
	Close *bool `json:"close,omitempty" xml:"close,omitempty"`
}

func (b *BodyUpdateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdPatch) GetMergeProposalId() (value string) {
	if b == nil {
		return
	}
	return b.MergeProposalId
}

func (b *BodyUpdateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdPatch) GetAgentId() (value string) {
	if b == nil {
		return
	}
	return b.AgentId
}

func (b *BodyUpdateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdPatch) GetTitle() (value *string) {
	if b == nil {
		return
	}
	return b.Title
}

func (b *BodyUpdateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdPatch) GetDescription() (value *string) {
	if b == nil {
		return
	}
	return b.Description
}

func (b *BodyUpdateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdPatch) GetRequestedReviewerUserIds() (value []string) {
	if b == nil {
		return
	}
	return b.RequestedReviewerUserIds
}

func (b *BodyUpdateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdPatch) GetClose() (value *bool) {
	if b == nil {
		return
	}
	return b.Close
}

func (b *BodyUpdateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdPatch) SetMergeProposalId(mergeProposalId string) {
	b.MergeProposalId = mergeProposalId
}

func (b *BodyUpdateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdPatch) SetAgentId(agentId string) {
	b.AgentId = agentId
}

func (b *BodyUpdateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdPatch) SetTitle(title *string) {
	b.Title = title
}

func (b *BodyUpdateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdPatch) SetDescription(description *string) {
	b.Description = description
}

func (b *BodyUpdateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdPatch) SetRequestedReviewerUserIds(requestedReviewerUserIds []string) {
	b.RequestedReviewerUserIds = requestedReviewerUserIds
}

func (b *BodyUpdateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdPatch) SetClose(close *bool) {
	b.Close = close
}

func (b *BodyUpdateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdPatch) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyUpdateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdPatch) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyUpdateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdPatch to string"
	}
	return string(jsonData)
}

func (b *BodyUpdateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdPatch) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyUpdateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdPatch) MarshalJSON() ([]byte, error) {
	type embed BodyUpdateAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdPatch
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

type ReviewAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdReviewsPost struct {
	State types.MergeProposalReviewState `json:"state" xml:"state" required:"true"`
	// Optional comment to leave with the review.
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty" maxLength:"4096"`
}

func (r *ReviewAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdReviewsPost) GetState() (value types.MergeProposalReviewState) {
	if r == nil {
		return
	}
	return r.State
}

func (r *ReviewAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdReviewsPost) GetComment() (value *string) {
	if r == nil {
		return
	}
	return r.Comment
}

func (r *ReviewAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdReviewsPost) SetState(state types.MergeProposalReviewState) {
	r.State = state
}

func (r *ReviewAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdReviewsPost) SetComment(comment *string) {
	r.Comment = comment
}

func (r *ReviewAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdReviewsPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (r ReviewAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdReviewsPost) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReviewAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdReviewsPost to string"
	}
	return string(jsonData)
}

func (r *ReviewAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdReviewsPost) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias ReviewAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdReviewsPost
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = ReviewAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdReviewsPost(tmp)
	return nil
}

func (r *ReviewAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdReviewsPost) MarshalJSON() ([]byte, error) {
	type embed ReviewAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdReviewsPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}

type OnAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdCommentsPost struct {
	// The comment text. Markdown is supported.
	Body string `json:"body" xml:"body" required:"true" maxLength:"4096" minLength:"1"`
}

func (o *OnAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdCommentsPost) GetBody() (value string) {
	if o == nil {
		return
	}
	return o.Body
}

func (o *OnAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdCommentsPost) SetBody(body string) {
	o.Body = body
}

func (o *OnAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdCommentsPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (o OnAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdCommentsPost) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OnAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdCommentsPost to string"
	}
	return string(jsonData)
}

func (o *OnAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdCommentsPost) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, o); err != nil {
		return err
	}
	type alias OnAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdCommentsPost
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*o = OnAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdCommentsPost(tmp)
	return nil
}

func (o *OnAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdCommentsPost) MarshalJSON() ([]byte, error) {
	type embed OnAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdCommentsPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}

type BodyMergeAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdMergePost struct {
	// Unique identifier for the merge_proposal.
	MergeProposalId string `json:"-" url:"-"`
	// The id of an agent. This is returned on agent creation.
	AgentId string `json:"-" url:"-"`
	// Whether to archive the source branch after merging.
	ArchiveSourceBranch *bool `json:"archive_source_branch,omitempty" xml:"archive_source_branch,omitempty"`
	// Force source branch changes onto the target, overriding timestamp-based conflict resolution.
	Force *bool `json:"force,omitempty" xml:"force,omitempty"`
}

func (b *BodyMergeAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdMergePost) GetMergeProposalId() (value string) {
	if b == nil {
		return
	}
	return b.MergeProposalId
}

func (b *BodyMergeAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdMergePost) GetAgentId() (value string) {
	if b == nil {
		return
	}
	return b.AgentId
}

func (b *BodyMergeAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdMergePost) GetArchiveSourceBranch() (value *bool) {
	if b == nil {
		return
	}
	return b.ArchiveSourceBranch
}

func (b *BodyMergeAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdMergePost) GetForce() (value *bool) {
	if b == nil {
		return
	}
	return b.Force
}

func (b *BodyMergeAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdMergePost) SetMergeProposalId(mergeProposalId string) {
	b.MergeProposalId = mergeProposalId
}

func (b *BodyMergeAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdMergePost) SetAgentId(agentId string) {
	b.AgentId = agentId
}

func (b *BodyMergeAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdMergePost) SetArchiveSourceBranch(archiveSourceBranch *bool) {
	b.ArchiveSourceBranch = archiveSourceBranch
}

func (b *BodyMergeAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdMergePost) SetForce(force *bool) {
	b.Force = force
}

func (b *BodyMergeAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdMergePost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyMergeAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdMergePost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyMergeAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdMergePost to string"
	}
	return string(jsonData)
}

func (b *BodyMergeAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdMergePost) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, b); err != nil {
		return err
	}
	type alias BodyMergeAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdMergePost
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = BodyMergeAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdMergePost(tmp)
	return nil
}

func (b *BodyMergeAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdMergePost) MarshalJSON() ([]byte, error) {
	type embed BodyMergeAMergeProposalV1ConvaiAgentsAgentIdMergeProposalsMergeProposalIdMergePost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

var (
	listMergeProposalsRequestFieldStatus         = big.NewInt(1 << 0)
	listMergeProposalsRequestFieldSourceBranchId = big.NewInt(1 << 1)
	listMergeProposalsRequestFieldTargetBranchId = big.NewInt(1 << 2)
	listMergeProposalsRequestFieldSearch         = big.NewInt(1 << 3)
	listMergeProposalsRequestFieldPageSize       = big.NewInt(1 << 4)
	listMergeProposalsRequestFieldCursor         = big.NewInt(1 << 5)
)

type ListMergeProposalsRequest struct {
	Status         *types.MergeProposalStatus `json:"-" url:"status,omitempty"`
	SourceBranchId *string                    `json:"-" url:"source_branch_id,omitempty"`
	TargetBranchId *string                    `json:"-" url:"target_branch_id,omitempty"`
	Search         *string                    `json:"-" url:"search,omitempty"`
	PageSize       *int                       `json:"-" url:"page_size,omitempty"`
	Cursor         *string                    `json:"-" url:"cursor,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListMergeProposalsRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListMergeProposalsRequest) SetStatus(status *types.MergeProposalStatus) {
	l.Status = status
	l.require(listMergeProposalsRequestFieldStatus)
}

// SetSourceBranchId sets the SourceBranchId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListMergeProposalsRequest) SetSourceBranchId(sourceBranchId *string) {
	l.SourceBranchId = sourceBranchId
	l.require(listMergeProposalsRequestFieldSourceBranchId)
}

// SetTargetBranchId sets the TargetBranchId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListMergeProposalsRequest) SetTargetBranchId(targetBranchId *string) {
	l.TargetBranchId = targetBranchId
	l.require(listMergeProposalsRequestFieldTargetBranchId)
}

// SetSearch sets the Search field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListMergeProposalsRequest) SetSearch(search *string) {
	l.Search = search
	l.require(listMergeProposalsRequestFieldSearch)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListMergeProposalsRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listMergeProposalsRequestFieldPageSize)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListMergeProposalsRequest) SetCursor(cursor *string) {
	l.Cursor = cursor
	l.require(listMergeProposalsRequestFieldCursor)
}
