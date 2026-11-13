package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type EvaluationCriteriaSummaryResult struct {
	Result          EvaluationSuccessResult `json:"result" xml:"result" required:"true"`
	Score           *int                    `json:"score,omitempty" xml:"score,omitempty"`
	MaxScore        *int                    `json:"max_score,omitempty" xml:"max_score,omitempty"`
	extraProperties map[string]interface{}
}

func (e *EvaluationCriteriaSummaryResult) GetResult() (value EvaluationSuccessResult) {
	if e == nil {
		return
	}
	return e.Result
}

func (e *EvaluationCriteriaSummaryResult) GetScore() (value *int) {
	if e == nil {
		return
	}
	return e.Score
}

func (e *EvaluationCriteriaSummaryResult) GetMaxScore() (value *int) {
	if e == nil {
		return
	}
	return e.MaxScore
}

func (e *EvaluationCriteriaSummaryResult) SetResult(result EvaluationSuccessResult) {
	e.Result = result
}

func (e *EvaluationCriteriaSummaryResult) SetScore(score *int) {
	e.Score = score
}

func (e *EvaluationCriteriaSummaryResult) SetMaxScore(maxScore *int) {
	e.MaxScore = maxScore
}

func (e *EvaluationCriteriaSummaryResult) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e EvaluationCriteriaSummaryResult) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EvaluationCriteriaSummaryResult to string"
	}
	return string(jsonData)
}

func (e *EvaluationCriteriaSummaryResult) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, e); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *EvaluationCriteriaSummaryResult) MarshalJSON() ([]byte, error) {
	type embed EvaluationCriteriaSummaryResult
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
