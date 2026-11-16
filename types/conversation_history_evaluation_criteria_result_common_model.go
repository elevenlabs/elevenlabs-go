package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationHistoryEvaluationCriteriaResultCommonModel struct {
	CriteriaId      string                  `json:"criteria_id" xml:"criteria_id" required:"true"`
	Result          EvaluationSuccessResult `json:"result" xml:"result" required:"true"`
	Rationale       string                  `json:"rationale" xml:"rationale" required:"true"`
	ScoringMode     *CriteriaScoringMode    `json:"scoring_mode,omitempty" xml:"scoring_mode,omitempty"`
	Score           *int                    `json:"score,omitempty" xml:"score,omitempty"`
	MaxScore        *int                    `json:"max_score,omitempty" xml:"max_score,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ConversationHistoryEvaluationCriteriaResultCommonModel) GetCriteriaId() (value string) {
	if c == nil {
		return
	}
	return c.CriteriaId
}

func (c *ConversationHistoryEvaluationCriteriaResultCommonModel) GetResult() (value EvaluationSuccessResult) {
	if c == nil {
		return
	}
	return c.Result
}

func (c *ConversationHistoryEvaluationCriteriaResultCommonModel) GetRationale() (value string) {
	if c == nil {
		return
	}
	return c.Rationale
}

func (c *ConversationHistoryEvaluationCriteriaResultCommonModel) GetScoringMode() (value *CriteriaScoringMode) {
	if c == nil {
		return
	}
	return c.ScoringMode
}

func (c *ConversationHistoryEvaluationCriteriaResultCommonModel) GetScore() (value *int) {
	if c == nil {
		return
	}
	return c.Score
}

func (c *ConversationHistoryEvaluationCriteriaResultCommonModel) GetMaxScore() (value *int) {
	if c == nil {
		return
	}
	return c.MaxScore
}

func (c *ConversationHistoryEvaluationCriteriaResultCommonModel) SetCriteriaId(criteriaId string) {
	c.CriteriaId = criteriaId
}

func (c *ConversationHistoryEvaluationCriteriaResultCommonModel) SetResult(result EvaluationSuccessResult) {
	c.Result = result
}

func (c *ConversationHistoryEvaluationCriteriaResultCommonModel) SetRationale(rationale string) {
	c.Rationale = rationale
}

func (c *ConversationHistoryEvaluationCriteriaResultCommonModel) SetScoringMode(scoringMode *CriteriaScoringMode) {
	c.ScoringMode = scoringMode
}

func (c *ConversationHistoryEvaluationCriteriaResultCommonModel) SetScore(score *int) {
	c.Score = score
}

func (c *ConversationHistoryEvaluationCriteriaResultCommonModel) SetMaxScore(maxScore *int) {
	c.MaxScore = maxScore
}

func (c *ConversationHistoryEvaluationCriteriaResultCommonModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationHistoryEvaluationCriteriaResultCommonModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationHistoryEvaluationCriteriaResultCommonModel to string"
	}
	return string(jsonData)
}

func (c *ConversationHistoryEvaluationCriteriaResultCommonModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationHistoryEvaluationCriteriaResultCommonModel) MarshalJSON() ([]byte, error) {
	type embed ConversationHistoryEvaluationCriteriaResultCommonModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
