package types

import (
	"fmt"
)

type UserFeedbackScore string

const (
	UserFeedbackScoreLike    UserFeedbackScore = "like"
	UserFeedbackScoreDislike UserFeedbackScore = "dislike"
)

func NewUserFeedbackScoreFromString(s string) (UserFeedbackScore, error) {
	switch s {
	case "like":
		return UserFeedbackScoreLike, nil
	case "dislike":
		return UserFeedbackScoreDislike, nil
	}
	var t UserFeedbackScore
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e UserFeedbackScore) Ptr() *UserFeedbackScore {
	return &e
}
