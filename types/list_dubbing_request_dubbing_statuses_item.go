package types

import (
	"fmt"
)

type ListDubbingRequestDubbingStatusesItem string

const (
	ListDubbingRequestDubbingStatusesItemQueued    ListDubbingRequestDubbingStatusesItem = "queued"
	ListDubbingRequestDubbingStatusesItemPreparing ListDubbingRequestDubbingStatusesItem = "preparing"
	ListDubbingRequestDubbingStatusesItemDubbing   ListDubbingRequestDubbingStatusesItem = "dubbing"
	ListDubbingRequestDubbingStatusesItemDubbed    ListDubbingRequestDubbingStatusesItem = "dubbed"
	ListDubbingRequestDubbingStatusesItemFailed    ListDubbingRequestDubbingStatusesItem = "failed"
)

func NewListDubbingRequestDubbingStatusesItemFromString(s string) (ListDubbingRequestDubbingStatusesItem, error) {
	switch s {
	case "queued":
		return ListDubbingRequestDubbingStatusesItemQueued, nil
	case "preparing":
		return ListDubbingRequestDubbingStatusesItemPreparing, nil
	case "dubbing":
		return ListDubbingRequestDubbingStatusesItemDubbing, nil
	case "dubbed":
		return ListDubbingRequestDubbingStatusesItemDubbed, nil
	case "failed":
		return ListDubbingRequestDubbingStatusesItemFailed, nil
	}
	var t ListDubbingRequestDubbingStatusesItem
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ListDubbingRequestDubbingStatusesItem) Ptr() *ListDubbingRequestDubbingStatusesItem {
	return &e
}
