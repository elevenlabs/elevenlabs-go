package types

import (
	"fmt"
)

type ConversationSentimentAnalysisOverallLabel string

const (
	ConversationSentimentAnalysisOverallLabelPositive ConversationSentimentAnalysisOverallLabel = "positive"
	ConversationSentimentAnalysisOverallLabelNeutral  ConversationSentimentAnalysisOverallLabel = "neutral"
	ConversationSentimentAnalysisOverallLabelNegative ConversationSentimentAnalysisOverallLabel = "negative"
)

func NewConversationSentimentAnalysisOverallLabelFromString(s string) (ConversationSentimentAnalysisOverallLabel, error) {
	switch s {
	case "positive":
		return ConversationSentimentAnalysisOverallLabelPositive, nil
	case "neutral":
		return ConversationSentimentAnalysisOverallLabelNeutral, nil
	case "negative":
		return ConversationSentimentAnalysisOverallLabelNegative, nil
	}
	var t ConversationSentimentAnalysisOverallLabel
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ConversationSentimentAnalysisOverallLabel) Ptr() *ConversationSentimentAnalysisOverallLabel {
	return &e
}
