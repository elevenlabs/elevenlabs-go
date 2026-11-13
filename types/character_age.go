package types

import (
	"fmt"
)

type CharacterAge string

const (
	CharacterAgeYoung      CharacterAge = "young"
	CharacterAgeMiddleAged CharacterAge = "middle_aged"
	CharacterAgeOld        CharacterAge = "old"
)

func NewCharacterAgeFromString(s string) (CharacterAge, error) {
	switch s {
	case "young":
		return CharacterAgeYoung, nil
	case "middle_aged":
		return CharacterAgeMiddleAged, nil
	case "old":
		return CharacterAgeOld, nil
	}
	var t CharacterAge
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CharacterAge) Ptr() *CharacterAge {
	return &e
}
