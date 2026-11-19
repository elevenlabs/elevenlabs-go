package types

import (
	"fmt"
)

type CharacterGender string

const (
	CharacterGenderMale    CharacterGender = "male"
	CharacterGenderFemale  CharacterGender = "female"
	CharacterGenderNeutral CharacterGender = "neutral"
)

func NewCharacterGenderFromString(s string) (CharacterGender, error) {
	switch s {
	case "male":
		return CharacterGenderMale, nil
	case "female":
		return CharacterGenderFemale, nil
	case "neutral":
		return CharacterGenderNeutral, nil
	}
	var t CharacterGender
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CharacterGender) Ptr() *CharacterGender {
	return &e
}
