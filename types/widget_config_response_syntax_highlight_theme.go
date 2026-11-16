package types

import (
	"fmt"
)

type WidgetConfigResponseSyntaxHighlightTheme string

const (
	WidgetConfigResponseSyntaxHighlightThemeLight WidgetConfigResponseSyntaxHighlightTheme = "light"
	WidgetConfigResponseSyntaxHighlightThemeDark  WidgetConfigResponseSyntaxHighlightTheme = "dark"
)

func NewWidgetConfigResponseSyntaxHighlightThemeFromString(s string) (WidgetConfigResponseSyntaxHighlightTheme, error) {
	switch s {
	case "light":
		return WidgetConfigResponseSyntaxHighlightThemeLight, nil
	case "dark":
		return WidgetConfigResponseSyntaxHighlightThemeDark, nil
	}
	var t WidgetConfigResponseSyntaxHighlightTheme
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e WidgetConfigResponseSyntaxHighlightTheme) Ptr() *WidgetConfigResponseSyntaxHighlightTheme {
	return &e
}
