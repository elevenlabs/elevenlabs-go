package types

import (
	"fmt"
)

// Where the LLM view renders a contextual update.
type ContextualUpdatePlacement string

const (
	ContextualUpdatePlacementInPlace ContextualUpdatePlacement = "in_place"
	ContextualUpdatePlacementEnd     ContextualUpdatePlacement = "end"
)

func NewContextualUpdatePlacementFromString(s string) (ContextualUpdatePlacement, error) {
	switch s {
	case "in_place":
		return ContextualUpdatePlacementInPlace, nil
	case "end":
		return ContextualUpdatePlacementEnd, nil
	}
	var t ContextualUpdatePlacement
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ContextualUpdatePlacement) Ptr() *ContextualUpdatePlacement {
	return &e
}
