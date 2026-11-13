package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SystemEvaluationCriteria struct {
	UserSentimentScore   float64 `json:"user_sentiment_score" xml:"user_sentiment_score" required:"true" min:"-1" max:"1"`
	UserFrustrationScore float64 `json:"user_frustration_score" xml:"user_frustration_score" required:"true" min:"0" max:"1"`
	extraProperties      map[string]interface{}
}

func (s *SystemEvaluationCriteria) GetUserSentimentScore() (value float64) {
	if s == nil {
		return
	}
	return s.UserSentimentScore
}

func (s *SystemEvaluationCriteria) GetUserFrustrationScore() (value float64) {
	if s == nil {
		return
	}
	return s.UserFrustrationScore
}

func (s *SystemEvaluationCriteria) SetUserSentimentScore(userSentimentScore float64) {
	s.UserSentimentScore = userSentimentScore
}

func (s *SystemEvaluationCriteria) SetUserFrustrationScore(userFrustrationScore float64) {
	s.UserFrustrationScore = userFrustrationScore
}

func (s *SystemEvaluationCriteria) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SystemEvaluationCriteria) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SystemEvaluationCriteria to string"
	}
	return string(jsonData)
}

func (s *SystemEvaluationCriteria) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SystemEvaluationCriteria
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SystemEvaluationCriteria(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SystemEvaluationCriteria) MarshalJSON() ([]byte, error) {
	type embed SystemEvaluationCriteria
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
