package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentMergeProposalResponse struct {
	Id                           string                             `json:"id" xml:"id" required:"true"`
	AgentId                      string                             `json:"agent_id" xml:"agent_id" required:"true"`
	SourceBranchId               string                             `json:"source_branch_id" xml:"source_branch_id" required:"true"`
	TargetBranchId               string                             `json:"target_branch_id" xml:"target_branch_id" required:"true"`
	SourceTipVersionIdAtCreation string                             `json:"source_tip_version_id_at_creation" xml:"source_tip_version_id_at_creation" required:"true"`
	Title                        string                             `json:"title" xml:"title" required:"true"`
	Description                  *string                            `json:"description,omitempty" xml:"description,omitempty"`
	RequestedReviewerUserIds     []string                           `json:"requested_reviewer_user_ids,omitempty" xml:"requested_reviewer_user_ids,omitempty"`
	Reviews                      []*MergeProposalReview             `json:"reviews,omitempty" xml:"reviews,omitempty"`
	Comments                     []*MergeProposalComment            `json:"comments,omitempty" xml:"comments,omitempty"`
	Outcome                      *AgentMergeProposalResponseOutcome `json:"outcome,omitempty" xml:"outcome,omitempty"`
	CreatedAt                    int                                `json:"created_at" xml:"created_at" required:"true"`
	UpdatedAt                    int                                `json:"updated_at" xml:"updated_at" required:"true"`
	// User ID of the merge_proposal's author.
	AuthorUserId *string `json:"author_user_id,omitempty" xml:"author_user_id,omitempty"`
	// Access information for the merge_proposal, including creator.
	AccessInfo      *ResourceAccessInfo `json:"access_info,omitempty" xml:"access_info,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AgentMergeProposalResponse) GetId() (value string) {
	if a == nil {
		return
	}
	return a.Id
}

func (a *AgentMergeProposalResponse) GetAgentId() (value string) {
	if a == nil {
		return
	}
	return a.AgentId
}

func (a *AgentMergeProposalResponse) GetSourceBranchId() (value string) {
	if a == nil {
		return
	}
	return a.SourceBranchId
}

func (a *AgentMergeProposalResponse) GetTargetBranchId() (value string) {
	if a == nil {
		return
	}
	return a.TargetBranchId
}

func (a *AgentMergeProposalResponse) GetSourceTipVersionIdAtCreation() (value string) {
	if a == nil {
		return
	}
	return a.SourceTipVersionIdAtCreation
}

func (a *AgentMergeProposalResponse) GetTitle() (value string) {
	if a == nil {
		return
	}
	return a.Title
}

func (a *AgentMergeProposalResponse) GetDescription() (value *string) {
	if a == nil {
		return
	}
	return a.Description
}

func (a *AgentMergeProposalResponse) GetRequestedReviewerUserIds() (value []string) {
	if a == nil {
		return
	}
	return a.RequestedReviewerUserIds
}

func (a *AgentMergeProposalResponse) GetReviews() (value []*MergeProposalReview) {
	if a == nil {
		return
	}
	return a.Reviews
}

func (a *AgentMergeProposalResponse) GetComments() (value []*MergeProposalComment) {
	if a == nil {
		return
	}
	return a.Comments
}

func (a *AgentMergeProposalResponse) GetOutcome() (value *AgentMergeProposalResponseOutcome) {
	if a == nil {
		return
	}
	return a.Outcome
}

func (a *AgentMergeProposalResponse) GetCreatedAt() (value int) {
	if a == nil {
		return
	}
	return a.CreatedAt
}

func (a *AgentMergeProposalResponse) GetUpdatedAt() (value int) {
	if a == nil {
		return
	}
	return a.UpdatedAt
}

func (a *AgentMergeProposalResponse) GetAuthorUserId() (value *string) {
	if a == nil {
		return
	}
	return a.AuthorUserId
}

func (a *AgentMergeProposalResponse) GetAccessInfo() (value *ResourceAccessInfo) {
	if a == nil {
		return
	}
	return a.AccessInfo
}

func (a *AgentMergeProposalResponse) SetId(id string) {
	a.Id = id
}

func (a *AgentMergeProposalResponse) SetAgentId(agentId string) {
	a.AgentId = agentId
}

func (a *AgentMergeProposalResponse) SetSourceBranchId(sourceBranchId string) {
	a.SourceBranchId = sourceBranchId
}

func (a *AgentMergeProposalResponse) SetTargetBranchId(targetBranchId string) {
	a.TargetBranchId = targetBranchId
}

func (a *AgentMergeProposalResponse) SetSourceTipVersionIdAtCreation(sourceTipVersionIdAtCreation string) {
	a.SourceTipVersionIdAtCreation = sourceTipVersionIdAtCreation
}

func (a *AgentMergeProposalResponse) SetTitle(title string) {
	a.Title = title
}

func (a *AgentMergeProposalResponse) SetDescription(description *string) {
	a.Description = description
}

func (a *AgentMergeProposalResponse) SetRequestedReviewerUserIds(requestedReviewerUserIds []string) {
	a.RequestedReviewerUserIds = requestedReviewerUserIds
}

func (a *AgentMergeProposalResponse) SetReviews(reviews []*MergeProposalReview) {
	a.Reviews = reviews
}

func (a *AgentMergeProposalResponse) SetComments(comments []*MergeProposalComment) {
	a.Comments = comments
}

func (a *AgentMergeProposalResponse) SetOutcome(outcome *AgentMergeProposalResponseOutcome) {
	a.Outcome = outcome
}

func (a *AgentMergeProposalResponse) SetCreatedAt(createdAt int) {
	a.CreatedAt = createdAt
}

func (a *AgentMergeProposalResponse) SetUpdatedAt(updatedAt int) {
	a.UpdatedAt = updatedAt
}

func (a *AgentMergeProposalResponse) SetAuthorUserId(authorUserId *string) {
	a.AuthorUserId = authorUserId
}

func (a *AgentMergeProposalResponse) SetAccessInfo(accessInfo *ResourceAccessInfo) {
	a.AccessInfo = accessInfo
}

func (a *AgentMergeProposalResponse) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentMergeProposalResponse) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentMergeProposalResponse to string"
	}
	return string(jsonData)
}

func (a *AgentMergeProposalResponse) UnmarshalJSON(data []byte) error {
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

func (a *AgentMergeProposalResponse) MarshalJSON() ([]byte, error) {
	type embed AgentMergeProposalResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
