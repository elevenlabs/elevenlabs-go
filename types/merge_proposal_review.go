package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type MergeProposalReview struct {
	UserId          string                           `json:"user_id" xml:"user_id" required:"true"`
	State           MergeProposalReviewState         `json:"state" xml:"state" required:"true"`
	SubmittedAt     int                              `json:"submitted_at" xml:"submitted_at" required:"true"`
	Comment         *string                          `json:"comment,omitempty" xml:"comment,omitempty"`
	ReviewerRole    *MergeProposalReviewReviewerRole `json:"reviewer_role,omitempty" xml:"reviewer_role,omitempty"`
	extraProperties map[string]interface{}
}

func (m *MergeProposalReview) GetUserId() (value string) {
	if m == nil {
		return
	}
	return m.UserId
}

func (m *MergeProposalReview) GetState() (value MergeProposalReviewState) {
	if m == nil {
		return
	}
	return m.State
}

func (m *MergeProposalReview) GetSubmittedAt() (value int) {
	if m == nil {
		return
	}
	return m.SubmittedAt
}

func (m *MergeProposalReview) GetComment() (value *string) {
	if m == nil {
		return
	}
	return m.Comment
}

func (m *MergeProposalReview) GetReviewerRole() (value *MergeProposalReviewReviewerRole) {
	if m == nil {
		return
	}
	return m.ReviewerRole
}

func (m *MergeProposalReview) SetUserId(userId string) {
	m.UserId = userId
}

func (m *MergeProposalReview) SetState(state MergeProposalReviewState) {
	m.State = state
}

func (m *MergeProposalReview) SetSubmittedAt(submittedAt int) {
	m.SubmittedAt = submittedAt
}

func (m *MergeProposalReview) SetComment(comment *string) {
	m.Comment = comment
}

func (m *MergeProposalReview) SetReviewerRole(reviewerRole *MergeProposalReviewReviewerRole) {
	m.ReviewerRole = reviewerRole
}

func (m *MergeProposalReview) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m MergeProposalReview) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MergeProposalReview to string"
	}
	return string(jsonData)
}

func (m *MergeProposalReview) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, m); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *MergeProposalReview) MarshalJSON() ([]byte, error) {
	type embed MergeProposalReview
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
