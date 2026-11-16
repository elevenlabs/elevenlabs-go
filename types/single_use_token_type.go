package types

import (
	"fmt"
)

type SingleUseTokenType string

const (
	SingleUseTokenTypeRealtimeScribe SingleUseTokenType = "realtime_scribe"
	SingleUseTokenTypeBatchScribe    SingleUseTokenType = "batch_scribe"
	SingleUseTokenTypeTtsWebsocket   SingleUseTokenType = "tts_websocket"
)

func NewSingleUseTokenTypeFromString(s string) (SingleUseTokenType, error) {
	switch s {
	case "realtime_scribe":
		return SingleUseTokenTypeRealtimeScribe, nil
	case "batch_scribe":
		return SingleUseTokenTypeBatchScribe, nil
	case "tts_websocket":
		return SingleUseTokenTypeTtsWebsocket, nil
	}
	var t SingleUseTokenType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SingleUseTokenType) Ptr() *SingleUseTokenType {
	return &e
}
