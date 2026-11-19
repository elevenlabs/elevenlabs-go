package types

import (
	"fmt"
)

type SessionStartedConfigModelId string

const (
	SessionStartedConfigModelIdScribeV2Realtime SessionStartedConfigModelId = "scribe_v2_realtime"
	SessionStartedConfigModelIdScribeV3Live     SessionStartedConfigModelId = "scribe_v3_live"
)

func NewSessionStartedConfigModelIdFromString(s string) (SessionStartedConfigModelId, error) {
	switch s {
	case "scribe_v2_realtime":
		return SessionStartedConfigModelIdScribeV2Realtime, nil
	case "scribe_v3_live":
		return SessionStartedConfigModelIdScribeV3Live, nil
	}
	var t SessionStartedConfigModelId
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SessionStartedConfigModelId) Ptr() *SessionStartedConfigModelId {
	return &e
}
