package types

import (
	"fmt"
)

type ListHistoryRequestSourcesItem string

const (
	ListHistoryRequestSourcesItemTts   ListHistoryRequestSourcesItem = "TTS"
	ListHistoryRequestSourcesItemSts   ListHistoryRequestSourcesItem = "STS"
	ListHistoryRequestSourcesItemFlows ListHistoryRequestSourcesItem = "Flows"
)

func NewListHistoryRequestSourcesItemFromString(s string) (ListHistoryRequestSourcesItem, error) {
	switch s {
	case "TTS":
		return ListHistoryRequestSourcesItemTts, nil
	case "STS":
		return ListHistoryRequestSourcesItemSts, nil
	case "Flows":
		return ListHistoryRequestSourcesItemFlows, nil
	}
	var t ListHistoryRequestSourcesItem
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ListHistoryRequestSourcesItem) Ptr() *ListHistoryRequestSourcesItem {
	return &e
}
