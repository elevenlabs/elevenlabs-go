package types

import (
	"fmt"
)

type PreToolSpeechMode string

const (
	PreToolSpeechModeAuto  PreToolSpeechMode = "auto"
	PreToolSpeechModeForce PreToolSpeechMode = "force"
	PreToolSpeechModeOff   PreToolSpeechMode = "off"
)

func NewPreToolSpeechModeFromString(s string) (PreToolSpeechMode, error) {
	switch s {
	case "auto":
		return PreToolSpeechModeAuto, nil
	case "force":
		return PreToolSpeechModeForce, nil
	case "off":
		return PreToolSpeechModeOff, nil
	}
	var t PreToolSpeechMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e PreToolSpeechMode) Ptr() *PreToolSpeechMode {
	return &e
}
