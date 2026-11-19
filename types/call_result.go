package types

import (
	"fmt"
)

type CallResult string

const (
	CallResultAnswered  CallResult = "answered"
	CallResultVoicemail CallResult = "voicemail"
	CallResultBusy      CallResult = "busy"
	CallResultNoAnswer  CallResult = "no_answer"
	CallResultFailed    CallResult = "failed"
)

func NewCallResultFromString(s string) (CallResult, error) {
	switch s {
	case "answered":
		return CallResultAnswered, nil
	case "voicemail":
		return CallResultVoicemail, nil
	case "busy":
		return CallResultBusy, nil
	case "no_answer":
		return CallResultNoAnswer, nil
	case "failed":
		return CallResultFailed, nil
	}
	var t CallResult
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CallResult) Ptr() *CallResult {
	return &e
}
