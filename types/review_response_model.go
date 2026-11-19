package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ReviewResponseModel struct {
	ReviewStatus    ReviewResponseModelReviewStatus        `json:"review_status" xml:"review_status" required:"true"`
	ReviewedAtUnix  int                                    `json:"reviewed_at_unix" xml:"reviewed_at_unix" required:"true"`
	ReviewedBy      *string                                `json:"reviewed_by,omitempty" xml:"reviewed_by,omitempty"`
	RejectReasons   []ReviewResponseModelRejectReasonsItem `json:"reject_reasons,omitempty" xml:"reject_reasons,omitempty"`
	ScoresBreakdown map[string]*int                        `json:"scores_breakdown,omitempty"`
	RejectedDetails *string                                `json:"rejected_details,omitempty" xml:"rejected_details,omitempty"`
	Explanation     *string                                `json:"explanation,omitempty" xml:"explanation,omitempty"`
	extraProperties map[string]interface{}
}

func (r *ReviewResponseModel) GetReviewStatus() (value ReviewResponseModelReviewStatus) {
	if r == nil {
		return
	}
	return r.ReviewStatus
}

func (r *ReviewResponseModel) GetReviewedAtUnix() (value int) {
	if r == nil {
		return
	}
	return r.ReviewedAtUnix
}

func (r *ReviewResponseModel) GetReviewedBy() (value *string) {
	if r == nil {
		return
	}
	return r.ReviewedBy
}

func (r *ReviewResponseModel) GetRejectReasons() (value []ReviewResponseModelRejectReasonsItem) {
	if r == nil {
		return
	}
	return r.RejectReasons
}

func (r *ReviewResponseModel) GetScoresBreakdown() (value map[string]*int) {
	if r == nil {
		return
	}
	return r.ScoresBreakdown
}

func (r *ReviewResponseModel) GetRejectedDetails() (value *string) {
	if r == nil {
		return
	}
	return r.RejectedDetails
}

func (r *ReviewResponseModel) GetExplanation() (value *string) {
	if r == nil {
		return
	}
	return r.Explanation
}

func (r *ReviewResponseModel) SetReviewStatus(reviewStatus ReviewResponseModelReviewStatus) {
	r.ReviewStatus = reviewStatus
}

func (r *ReviewResponseModel) SetReviewedAtUnix(reviewedAtUnix int) {
	r.ReviewedAtUnix = reviewedAtUnix
}

func (r *ReviewResponseModel) SetReviewedBy(reviewedBy *string) {
	r.ReviewedBy = reviewedBy
}

func (r *ReviewResponseModel) SetRejectReasons(rejectReasons []ReviewResponseModelRejectReasonsItem) {
	r.RejectReasons = rejectReasons
}

func (r *ReviewResponseModel) SetScoresBreakdown(scoresBreakdown map[string]*int) {
	r.ScoresBreakdown = scoresBreakdown
}

func (r *ReviewResponseModel) SetRejectedDetails(rejectedDetails *string) {
	r.RejectedDetails = rejectedDetails
}

func (r *ReviewResponseModel) SetExplanation(explanation *string) {
	r.Explanation = explanation
}

func (r *ReviewResponseModel) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReviewResponseModel) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReviewResponseModel to string"
	}
	return string(jsonData)
}

func (r *ReviewResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, r); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *ReviewResponseModel) MarshalJSON() ([]byte, error) {
	type embed ReviewResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
