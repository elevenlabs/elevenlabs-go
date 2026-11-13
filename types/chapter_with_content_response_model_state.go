package types

import (
	"fmt"
)

// The state of the chapter.
type ChapterWithContentResponseModelState string

const (
	ChapterWithContentResponseModelStateDefault    ChapterWithContentResponseModelState = "default"
	ChapterWithContentResponseModelStateConverting ChapterWithContentResponseModelState = "converting"
)

func NewChapterWithContentResponseModelStateFromString(s string) (ChapterWithContentResponseModelState, error) {
	switch s {
	case "default":
		return ChapterWithContentResponseModelStateDefault, nil
	case "converting":
		return ChapterWithContentResponseModelStateConverting, nil
	}
	var t ChapterWithContentResponseModelState
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ChapterWithContentResponseModelState) Ptr() *ChapterWithContentResponseModelState {
	return &e
}
