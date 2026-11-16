package types

import (
	"fmt"
)

type IconTheme string

const (
	IconThemeLight IconTheme = "light"
	IconThemeDark  IconTheme = "dark"
)

func NewIconThemeFromString(s string) (IconTheme, error) {
	switch s {
	case "light":
		return IconThemeLight, nil
	case "dark":
		return IconThemeDark, nil
	}
	var t IconTheme
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e IconTheme) Ptr() *IconTheme {
	return &e
}
