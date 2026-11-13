package types

import (
	"fmt"
)

type SessionStartedConfigTimestampsGranularity string

const (
	SessionStartedConfigTimestampsGranularityNone      SessionStartedConfigTimestampsGranularity = "none"
	SessionStartedConfigTimestampsGranularityWord      SessionStartedConfigTimestampsGranularity = "word"
	SessionStartedConfigTimestampsGranularityCharacter SessionStartedConfigTimestampsGranularity = "character"
)

func NewSessionStartedConfigTimestampsGranularityFromString(s string) (SessionStartedConfigTimestampsGranularity, error) {
	switch s {
	case "none":
		return SessionStartedConfigTimestampsGranularityNone, nil
	case "word":
		return SessionStartedConfigTimestampsGranularityWord, nil
	case "character":
		return SessionStartedConfigTimestampsGranularityCharacter, nil
	}
	var t SessionStartedConfigTimestampsGranularity
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SessionStartedConfigTimestampsGranularity) Ptr() *SessionStartedConfigTimestampsGranularity {
	return &e
}
