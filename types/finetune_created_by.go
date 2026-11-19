package types

import (
	"fmt"
)

type FinetuneCreatedBy string

const (
	FinetuneCreatedBySelf       FinetuneCreatedBy = "self"
	FinetuneCreatedByWorkspace  FinetuneCreatedBy = "workspace"
	FinetuneCreatedByElevenlabs FinetuneCreatedBy = "elevenlabs"
)

func NewFinetuneCreatedByFromString(s string) (FinetuneCreatedBy, error) {
	switch s {
	case "self":
		return FinetuneCreatedBySelf, nil
	case "workspace":
		return FinetuneCreatedByWorkspace, nil
	case "elevenlabs":
		return FinetuneCreatedByElevenlabs, nil
	}
	var t FinetuneCreatedBy
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e FinetuneCreatedBy) Ptr() *FinetuneCreatedBy {
	return &e
}
