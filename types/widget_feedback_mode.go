package types

import (
	"fmt"
)

type WidgetFeedbackMode string

const (
	WidgetFeedbackModeNone   WidgetFeedbackMode = "none"
	WidgetFeedbackModeDuring WidgetFeedbackMode = "during"
	WidgetFeedbackModeEnd    WidgetFeedbackMode = "end"
)

func NewWidgetFeedbackModeFromString(s string) (WidgetFeedbackMode, error) {
	switch s {
	case "none":
		return WidgetFeedbackModeNone, nil
	case "during":
		return WidgetFeedbackModeDuring, nil
	case "end":
		return WidgetFeedbackModeEnd, nil
	}
	var t WidgetFeedbackMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e WidgetFeedbackMode) Ptr() *WidgetFeedbackMode {
	return &e
}
