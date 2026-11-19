package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TopicEvaluationCriteriaAggregate struct {
	CriteriaId      string `json:"criteria_id" xml:"criteria_id" required:"true"`
	SuccessCount    *int   `json:"success_count,omitempty" xml:"success_count,omitempty"`
	FailureCount    *int   `json:"failure_count,omitempty" xml:"failure_count,omitempty"`
	UnknownCount    *int   `json:"unknown_count,omitempty" xml:"unknown_count,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TopicEvaluationCriteriaAggregate) GetCriteriaId() (value string) {
	if t == nil {
		return
	}
	return t.CriteriaId
}

func (t *TopicEvaluationCriteriaAggregate) GetSuccessCount() (value *int) {
	if t == nil {
		return
	}
	return t.SuccessCount
}

func (t *TopicEvaluationCriteriaAggregate) GetFailureCount() (value *int) {
	if t == nil {
		return
	}
	return t.FailureCount
}

func (t *TopicEvaluationCriteriaAggregate) GetUnknownCount() (value *int) {
	if t == nil {
		return
	}
	return t.UnknownCount
}

func (t *TopicEvaluationCriteriaAggregate) SetCriteriaId(criteriaId string) {
	t.CriteriaId = criteriaId
}

func (t *TopicEvaluationCriteriaAggregate) SetSuccessCount(successCount *int) {
	t.SuccessCount = successCount
}

func (t *TopicEvaluationCriteriaAggregate) SetFailureCount(failureCount *int) {
	t.FailureCount = failureCount
}

func (t *TopicEvaluationCriteriaAggregate) SetUnknownCount(unknownCount *int) {
	t.UnknownCount = unknownCount
}

func (t *TopicEvaluationCriteriaAggregate) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TopicEvaluationCriteriaAggregate) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TopicEvaluationCriteriaAggregate to string"
	}
	return string(jsonData)
}

func (t *TopicEvaluationCriteriaAggregate) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TopicEvaluationCriteriaAggregate
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TopicEvaluationCriteriaAggregate(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TopicEvaluationCriteriaAggregate) MarshalJSON() ([]byte, error) {
	type embed TopicEvaluationCriteriaAggregate
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
