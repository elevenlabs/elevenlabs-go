package types

import (
	"fmt"
)

// Filters who created the resources being listed, whether it was the user running the request or someone else that shared the resource with them.
type ListDubbingRequestFilterByCreator string

const (
	ListDubbingRequestFilterByCreatorPersonal ListDubbingRequestFilterByCreator = "personal"
	ListDubbingRequestFilterByCreatorOthers   ListDubbingRequestFilterByCreator = "others"
	ListDubbingRequestFilterByCreatorAll      ListDubbingRequestFilterByCreator = "all"
)

func NewListDubbingRequestFilterByCreatorFromString(s string) (ListDubbingRequestFilterByCreator, error) {
	switch s {
	case "personal":
		return ListDubbingRequestFilterByCreatorPersonal, nil
	case "others":
		return ListDubbingRequestFilterByCreatorOthers, nil
	case "all":
		return ListDubbingRequestFilterByCreatorAll, nil
	}
	var t ListDubbingRequestFilterByCreator
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ListDubbingRequestFilterByCreator) Ptr() *ListDubbingRequestFilterByCreator {
	return &e
}
