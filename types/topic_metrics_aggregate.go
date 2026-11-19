package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TopicMetricsAggregate struct {
	ConversationCount  *int                                `json:"conversation_count,omitempty" xml:"conversation_count,omitempty"`
	Sentiment          *TopicSentimentAggregate            `json:"sentiment,omitempty" xml:"sentiment,omitempty"`
	EvaluationCriteria []*TopicEvaluationCriteriaAggregate `json:"evaluation_criteria,omitempty" xml:"evaluation_criteria,omitempty"`
	extraProperties    map[string]interface{}
}

func (t *TopicMetricsAggregate) GetConversationCount() (value *int) {
	if t == nil {
		return
	}
	return t.ConversationCount
}

func (t *TopicMetricsAggregate) GetSentiment() (value *TopicSentimentAggregate) {
	if t == nil {
		return
	}
	return t.Sentiment
}

func (t *TopicMetricsAggregate) GetEvaluationCriteria() (value []*TopicEvaluationCriteriaAggregate) {
	if t == nil {
		return
	}
	return t.EvaluationCriteria
}

func (t *TopicMetricsAggregate) SetConversationCount(conversationCount *int) {
	t.ConversationCount = conversationCount
}

func (t *TopicMetricsAggregate) SetSentiment(sentiment *TopicSentimentAggregate) {
	t.Sentiment = sentiment
}

func (t *TopicMetricsAggregate) SetEvaluationCriteria(evaluationCriteria []*TopicEvaluationCriteriaAggregate) {
	t.EvaluationCriteria = evaluationCriteria
}

func (t *TopicMetricsAggregate) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TopicMetricsAggregate) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TopicMetricsAggregate to string"
	}
	return string(jsonData)
}

func (t *TopicMetricsAggregate) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TopicMetricsAggregate) MarshalJSON() ([]byte, error) {
	type embed TopicMetricsAggregate
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
