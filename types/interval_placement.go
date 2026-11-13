package types

import (
	"fmt"
)

type IntervalPlacement string

const (
	IntervalPlacementStart  IntervalPlacement = "start"
	IntervalPlacementRandom IntervalPlacement = "random"
	IntervalPlacementRotate IntervalPlacement = "rotate"
)

func NewIntervalPlacementFromString(s string) (IntervalPlacement, error) {
	switch s {
	case "start":
		return IntervalPlacementStart, nil
	case "random":
		return IntervalPlacementRandom, nil
	case "rotate":
		return IntervalPlacementRotate, nil
	}
	var t IntervalPlacement
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e IntervalPlacement) Ptr() *IntervalPlacement {
	return &e
}
