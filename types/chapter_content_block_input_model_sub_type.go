package types

import (
	"fmt"
)

type ChapterContentBlockInputModelSubType string

const (
	ChapterContentBlockInputModelSubTypeP  ChapterContentBlockInputModelSubType = "p"
	ChapterContentBlockInputModelSubTypeH1 ChapterContentBlockInputModelSubType = "h1"
	ChapterContentBlockInputModelSubTypeH2 ChapterContentBlockInputModelSubType = "h2"
	ChapterContentBlockInputModelSubTypeH3 ChapterContentBlockInputModelSubType = "h3"
)

func NewChapterContentBlockInputModelSubTypeFromString(s string) (ChapterContentBlockInputModelSubType, error) {
	switch s {
	case "p":
		return ChapterContentBlockInputModelSubTypeP, nil
	case "h1":
		return ChapterContentBlockInputModelSubTypeH1, nil
	case "h2":
		return ChapterContentBlockInputModelSubTypeH2, nil
	case "h3":
		return ChapterContentBlockInputModelSubTypeH3, nil
	}
	var t ChapterContentBlockInputModelSubType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ChapterContentBlockInputModelSubType) Ptr() *ChapterContentBlockInputModelSubType {
	return &e
}
