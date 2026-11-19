package types

import (
	"fmt"
)

// The state of the chapter.
type ChapterState string

const (
	ChapterStateDefault    ChapterState = "default"
	ChapterStateConverting ChapterState = "converting"
)

func NewChapterStateFromString(s string) (ChapterState, error) {
	switch s {
	case "default":
		return ChapterStateDefault, nil
	case "converting":
		return ChapterStateConverting, nil
	}
	var t ChapterState
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ChapterState) Ptr() *ChapterState {
	return &e
}
