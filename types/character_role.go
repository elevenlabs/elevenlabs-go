package types

import (
	"fmt"
)

type CharacterRole string

const (
	CharacterRoleNarrator   CharacterRole = "narrator"
	CharacterRoleMain       CharacterRole = "main"
	CharacterRoleSupporting CharacterRole = "supporting"
	CharacterRoleMinor      CharacterRole = "minor"
)

func NewCharacterRoleFromString(s string) (CharacterRole, error) {
	switch s {
	case "narrator":
		return CharacterRoleNarrator, nil
	case "main":
		return CharacterRoleMain, nil
	case "supporting":
		return CharacterRoleSupporting, nil
	case "minor":
		return CharacterRoleMinor, nil
	}
	var t CharacterRole
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CharacterRole) Ptr() *CharacterRole {
	return &e
}
