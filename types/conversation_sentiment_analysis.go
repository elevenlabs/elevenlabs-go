package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationSentimentAnalysis struct {
	OverallLabel            ConversationSentimentAnalysisOverallLabel `json:"overall_label" xml:"overall_label" required:"true"`
	OverallSentimentScore   float64                                   `json:"overall_sentiment_score" xml:"overall_sentiment_score" required:"true" min:"-1" max:"1"`
	OverallFrustrationScore float64                                   `json:"overall_frustration_score" xml:"overall_frustration_score" required:"true" min:"0" max:"1"`
	MinUserSentimentScore   float64                                   `json:"min_user_sentiment_score" xml:"min_user_sentiment_score" required:"true" min:"-1" max:"1"`
	MaxUserFrustrationScore float64                                   `json:"max_user_frustration_score" xml:"max_user_frustration_score" required:"true" min:"0" max:"1"`
	NumScoredUserTurns      int                                       `json:"num_scored_user_turns" xml:"num_scored_user_turns" required:"true" min:"0"`
	extraProperties         map[string]interface{}
}

func (c *ConversationSentimentAnalysis) GetOverallLabel() (value ConversationSentimentAnalysisOverallLabel) {
	if c == nil {
		return
	}
	return c.OverallLabel
}

func (c *ConversationSentimentAnalysis) GetOverallSentimentScore() (value float64) {
	if c == nil {
		return
	}
	return c.OverallSentimentScore
}

func (c *ConversationSentimentAnalysis) GetOverallFrustrationScore() (value float64) {
	if c == nil {
		return
	}
	return c.OverallFrustrationScore
}

func (c *ConversationSentimentAnalysis) GetMinUserSentimentScore() (value float64) {
	if c == nil {
		return
	}
	return c.MinUserSentimentScore
}

func (c *ConversationSentimentAnalysis) GetMaxUserFrustrationScore() (value float64) {
	if c == nil {
		return
	}
	return c.MaxUserFrustrationScore
}

func (c *ConversationSentimentAnalysis) GetNumScoredUserTurns() (value int) {
	if c == nil {
		return
	}
	return c.NumScoredUserTurns
}

func (c *ConversationSentimentAnalysis) SetOverallLabel(overallLabel ConversationSentimentAnalysisOverallLabel) {
	c.OverallLabel = overallLabel
}

func (c *ConversationSentimentAnalysis) SetOverallSentimentScore(overallSentimentScore float64) {
	c.OverallSentimentScore = overallSentimentScore
}

func (c *ConversationSentimentAnalysis) SetOverallFrustrationScore(overallFrustrationScore float64) {
	c.OverallFrustrationScore = overallFrustrationScore
}

func (c *ConversationSentimentAnalysis) SetMinUserSentimentScore(minUserSentimentScore float64) {
	c.MinUserSentimentScore = minUserSentimentScore
}

func (c *ConversationSentimentAnalysis) SetMaxUserFrustrationScore(maxUserFrustrationScore float64) {
	c.MaxUserFrustrationScore = maxUserFrustrationScore
}

func (c *ConversationSentimentAnalysis) SetNumScoredUserTurns(numScoredUserTurns int) {
	c.NumScoredUserTurns = numScoredUserTurns
}

func (c *ConversationSentimentAnalysis) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationSentimentAnalysis) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationSentimentAnalysis to string"
	}
	return string(jsonData)
}

func (c *ConversationSentimentAnalysis) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ConversationSentimentAnalysis
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ConversationSentimentAnalysis(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationSentimentAnalysis) MarshalJSON() ([]byte, error) {
	type embed ConversationSentimentAnalysis
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
