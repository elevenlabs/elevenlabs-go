package types

import (
	"fmt"
)

type FinetuneVisibility string

const (
	FinetuneVisibilityPrivate   FinetuneVisibility = "private"
	FinetuneVisibilityWorkspace FinetuneVisibility = "workspace"
	FinetuneVisibilityPublic    FinetuneVisibility = "public"
)

func NewFinetuneVisibilityFromString(s string) (FinetuneVisibility, error) {
	switch s {
	case "private":
		return FinetuneVisibilityPrivate, nil
	case "workspace":
		return FinetuneVisibilityWorkspace, nil
	case "public":
		return FinetuneVisibilityPublic, nil
	}
	var t FinetuneVisibility
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e FinetuneVisibility) Ptr() *FinetuneVisibility {
	return &e
}
