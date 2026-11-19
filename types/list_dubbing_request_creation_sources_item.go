package types

import (
	"fmt"
)

type ListDubbingRequestCreationSourcesItem string

const (
	ListDubbingRequestCreationSourcesItemFlowNode   ListDubbingRequestCreationSourcesItem = "flow_node"
	ListDubbingRequestCreationSourcesItemDubbingUi  ListDubbingRequestCreationSourcesItem = "dubbing_ui"
	ListDubbingRequestCreationSourcesItemDubbingApi ListDubbingRequestCreationSourcesItem = "dubbing_api"
)

func NewListDubbingRequestCreationSourcesItemFromString(s string) (ListDubbingRequestCreationSourcesItem, error) {
	switch s {
	case "flow_node":
		return ListDubbingRequestCreationSourcesItemFlowNode, nil
	case "dubbing_ui":
		return ListDubbingRequestCreationSourcesItemDubbingUi, nil
	case "dubbing_api":
		return ListDubbingRequestCreationSourcesItemDubbingApi, nil
	}
	var t ListDubbingRequestCreationSourcesItem
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ListDubbingRequestCreationSourcesItem) Ptr() *ListDubbingRequestCreationSourcesItem {
	return &e
}
