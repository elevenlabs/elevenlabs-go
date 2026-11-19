package types

import (
	"fmt"
)

type WidgetConfigSyntaxHighlightTheme string

const (
	WidgetConfigSyntaxHighlightThemeLight WidgetConfigSyntaxHighlightTheme = "light"
	WidgetConfigSyntaxHighlightThemeDark  WidgetConfigSyntaxHighlightTheme = "dark"
)

func NewWidgetConfigSyntaxHighlightThemeFromString(s string) (WidgetConfigSyntaxHighlightTheme, error) {
	switch s {
	case "light":
		return WidgetConfigSyntaxHighlightThemeLight, nil
	case "dark":
		return WidgetConfigSyntaxHighlightThemeDark, nil
	}
	var t WidgetConfigSyntaxHighlightTheme
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e WidgetConfigSyntaxHighlightTheme) Ptr() *WidgetConfigSyntaxHighlightTheme {
	return &e
}
