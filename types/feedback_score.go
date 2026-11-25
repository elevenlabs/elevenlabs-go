package types

import (
	"fmt"
)

// Feedback score for the referenced agent response. Null clears any previously submitted feedback for the event.
type FeedbackScore string

const (
	FeedbackScoreLike    FeedbackScore = "like"
	FeedbackScoreDislike FeedbackScore = "dislike"
)

func NewFeedbackScoreFromString(s string) (FeedbackScore, error) {
	switch s {
	case "like":
		return FeedbackScoreLike, nil
	case "dislike":
		return FeedbackScoreDislike, nil
	}
	var t FeedbackScore
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e FeedbackScore) Ptr() *FeedbackScore {
	return &e
}
