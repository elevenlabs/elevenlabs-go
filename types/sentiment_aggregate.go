package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SentimentAggregate struct {
	ScoredConversationCount       int      `json:"scored_conversation_count" xml:"scored_conversation_count" required:"true"`
	PositiveCount                 int      `json:"positive_count" xml:"positive_count" required:"true"`
	NeutralCount                  int      `json:"neutral_count" xml:"neutral_count" required:"true"`
	NegativeCount                 int      `json:"negative_count" xml:"negative_count" required:"true"`
	AverageSentimentScore         *float64 `json:"average_sentiment_score" xml:"average_sentiment_score" required:"true"`
	AverageFrustrationScore       *float64 `json:"average_frustration_score" xml:"average_frustration_score" required:"true"`
	RecentScoredConversationCount int      `json:"recent_scored_conversation_count" xml:"recent_scored_conversation_count" required:"true"`
	RecentPositiveCount           int      `json:"recent_positive_count" xml:"recent_positive_count" required:"true"`
	RecentNeutralCount            int      `json:"recent_neutral_count" xml:"recent_neutral_count" required:"true"`
	RecentNegativeCount           int      `json:"recent_negative_count" xml:"recent_negative_count" required:"true"`
	RecentAverageSentimentScore   *float64 `json:"recent_average_sentiment_score" xml:"recent_average_sentiment_score" required:"true"`
	RecentAverageFrustrationScore *float64 `json:"recent_average_frustration_score" xml:"recent_average_frustration_score" required:"true"`
	extraProperties               map[string]interface{}
}

func (s *SentimentAggregate) GetScoredConversationCount() (value int) {
	if s == nil {
		return
	}
	return s.ScoredConversationCount
}

func (s *SentimentAggregate) GetPositiveCount() (value int) {
	if s == nil {
		return
	}
	return s.PositiveCount
}

func (s *SentimentAggregate) GetNeutralCount() (value int) {
	if s == nil {
		return
	}
	return s.NeutralCount
}

func (s *SentimentAggregate) GetNegativeCount() (value int) {
	if s == nil {
		return
	}
	return s.NegativeCount
}

func (s *SentimentAggregate) GetAverageSentimentScore() (value *float64) {
	if s == nil {
		return
	}
	return s.AverageSentimentScore
}

func (s *SentimentAggregate) GetAverageFrustrationScore() (value *float64) {
	if s == nil {
		return
	}
	return s.AverageFrustrationScore
}

func (s *SentimentAggregate) GetRecentScoredConversationCount() (value int) {
	if s == nil {
		return
	}
	return s.RecentScoredConversationCount
}

func (s *SentimentAggregate) GetRecentPositiveCount() (value int) {
	if s == nil {
		return
	}
	return s.RecentPositiveCount
}

func (s *SentimentAggregate) GetRecentNeutralCount() (value int) {
	if s == nil {
		return
	}
	return s.RecentNeutralCount
}

func (s *SentimentAggregate) GetRecentNegativeCount() (value int) {
	if s == nil {
		return
	}
	return s.RecentNegativeCount
}

func (s *SentimentAggregate) GetRecentAverageSentimentScore() (value *float64) {
	if s == nil {
		return
	}
	return s.RecentAverageSentimentScore
}

func (s *SentimentAggregate) GetRecentAverageFrustrationScore() (value *float64) {
	if s == nil {
		return
	}
	return s.RecentAverageFrustrationScore
}

func (s *SentimentAggregate) SetScoredConversationCount(scoredConversationCount int) {
	s.ScoredConversationCount = scoredConversationCount
}

func (s *SentimentAggregate) SetPositiveCount(positiveCount int) {
	s.PositiveCount = positiveCount
}

func (s *SentimentAggregate) SetNeutralCount(neutralCount int) {
	s.NeutralCount = neutralCount
}

func (s *SentimentAggregate) SetNegativeCount(negativeCount int) {
	s.NegativeCount = negativeCount
}

func (s *SentimentAggregate) SetAverageSentimentScore(averageSentimentScore *float64) {
	s.AverageSentimentScore = averageSentimentScore
}

func (s *SentimentAggregate) SetAverageFrustrationScore(averageFrustrationScore *float64) {
	s.AverageFrustrationScore = averageFrustrationScore
}

func (s *SentimentAggregate) SetRecentScoredConversationCount(recentScoredConversationCount int) {
	s.RecentScoredConversationCount = recentScoredConversationCount
}

func (s *SentimentAggregate) SetRecentPositiveCount(recentPositiveCount int) {
	s.RecentPositiveCount = recentPositiveCount
}

func (s *SentimentAggregate) SetRecentNeutralCount(recentNeutralCount int) {
	s.RecentNeutralCount = recentNeutralCount
}

func (s *SentimentAggregate) SetRecentNegativeCount(recentNegativeCount int) {
	s.RecentNegativeCount = recentNegativeCount
}

func (s *SentimentAggregate) SetRecentAverageSentimentScore(recentAverageSentimentScore *float64) {
	s.RecentAverageSentimentScore = recentAverageSentimentScore
}

func (s *SentimentAggregate) SetRecentAverageFrustrationScore(recentAverageFrustrationScore *float64) {
	s.RecentAverageFrustrationScore = recentAverageFrustrationScore
}

func (s *SentimentAggregate) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SentimentAggregate) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SentimentAggregate to string"
	}
	return string(jsonData)
}

func (s *SentimentAggregate) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, s); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SentimentAggregate) MarshalJSON() ([]byte, error) {
	type embed SentimentAggregate
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
