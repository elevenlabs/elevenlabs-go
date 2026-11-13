package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type TopicSentimentAggregate struct {
	Sentiment       *NumericDistributionAggregate `json:"sentiment,omitempty" xml:"sentiment,omitempty"`
	Frustration     *NumericDistributionAggregate `json:"frustration,omitempty" xml:"frustration,omitempty"`
	PositiveCount   *int                          `json:"positive_count,omitempty" xml:"positive_count,omitempty"`
	NeutralCount    *int                          `json:"neutral_count,omitempty" xml:"neutral_count,omitempty"`
	NegativeCount   *int                          `json:"negative_count,omitempty" xml:"negative_count,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TopicSentimentAggregate) GetSentiment() (value *NumericDistributionAggregate) {
	if t == nil {
		return
	}
	return t.Sentiment
}

func (t *TopicSentimentAggregate) GetFrustration() (value *NumericDistributionAggregate) {
	if t == nil {
		return
	}
	return t.Frustration
}

func (t *TopicSentimentAggregate) GetPositiveCount() (value *int) {
	if t == nil {
		return
	}
	return t.PositiveCount
}

func (t *TopicSentimentAggregate) GetNeutralCount() (value *int) {
	if t == nil {
		return
	}
	return t.NeutralCount
}

func (t *TopicSentimentAggregate) GetNegativeCount() (value *int) {
	if t == nil {
		return
	}
	return t.NegativeCount
}

func (t *TopicSentimentAggregate) SetSentiment(sentiment *NumericDistributionAggregate) {
	t.Sentiment = sentiment
}

func (t *TopicSentimentAggregate) SetFrustration(frustration *NumericDistributionAggregate) {
	t.Frustration = frustration
}

func (t *TopicSentimentAggregate) SetPositiveCount(positiveCount *int) {
	t.PositiveCount = positiveCount
}

func (t *TopicSentimentAggregate) SetNeutralCount(neutralCount *int) {
	t.NeutralCount = neutralCount
}

func (t *TopicSentimentAggregate) SetNegativeCount(negativeCount *int) {
	t.NegativeCount = negativeCount
}

func (t *TopicSentimentAggregate) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TopicSentimentAggregate) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TopicSentimentAggregate to string"
	}
	return string(jsonData)
}

func (t *TopicSentimentAggregate) UnmarshalJSON(data []byte) error {
	type unmarshaler TopicSentimentAggregate
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TopicSentimentAggregate(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TopicSentimentAggregate) MarshalJSON() ([]byte, error) {
	type embed TopicSentimentAggregate
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
