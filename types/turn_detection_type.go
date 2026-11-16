package types

import (
	"fmt"
)

type TurnDetectionType string

const (
	TurnDetectionTypeSemanticVad TurnDetectionType = "semantic_vad"
	TurnDetectionTypeServerVad   TurnDetectionType = "server_vad"
)

func NewTurnDetectionTypeFromString(s string) (TurnDetectionType, error) {
	switch s {
	case "semantic_vad":
		return TurnDetectionTypeSemanticVad, nil
	case "server_vad":
		return TurnDetectionTypeServerVad, nil
	}
	var t TurnDetectionType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TurnDetectionType) Ptr() *TurnDetectionType {
	return &e
}
