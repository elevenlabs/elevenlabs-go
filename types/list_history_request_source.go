package types

import (
	"fmt"
)

type ListHistoryRequestSource string

const (
	ListHistoryRequestSourceTts   ListHistoryRequestSource = "TTS"
	ListHistoryRequestSourceSts   ListHistoryRequestSource = "STS"
	ListHistoryRequestSourceFlows ListHistoryRequestSource = "Flows"
)

func NewListHistoryRequestSourceFromString(s string) (ListHistoryRequestSource, error) {
	switch s {
	case "TTS":
		return ListHistoryRequestSourceTts, nil
	case "STS":
		return ListHistoryRequestSourceSts, nil
	case "Flows":
		return ListHistoryRequestSourceFlows, nil
	}
	var t ListHistoryRequestSource
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ListHistoryRequestSource) Ptr() *ListHistoryRequestSource {
	return &e
}
