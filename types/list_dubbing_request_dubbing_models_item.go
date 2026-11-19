package types

import (
	"fmt"
)

type ListDubbingRequestDubbingModelsItem string

const (
	ListDubbingRequestDubbingModelsItemDubbingV1 ListDubbingRequestDubbingModelsItem = "dubbing_v1"
	ListDubbingRequestDubbingModelsItemDubbingV2 ListDubbingRequestDubbingModelsItem = "dubbing_v2"
)

func NewListDubbingRequestDubbingModelsItemFromString(s string) (ListDubbingRequestDubbingModelsItem, error) {
	switch s {
	case "dubbing_v1":
		return ListDubbingRequestDubbingModelsItemDubbingV1, nil
	case "dubbing_v2":
		return ListDubbingRequestDubbingModelsItemDubbingV2, nil
	}
	var t ListDubbingRequestDubbingModelsItem
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ListDubbingRequestDubbingModelsItem) Ptr() *ListDubbingRequestDubbingModelsItem {
	return &e
}
