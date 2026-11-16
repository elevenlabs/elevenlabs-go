package types

import (
	"fmt"
)

type ListPronunciationDictionariesRequestSort string

const (
	ListPronunciationDictionariesRequestSortCreationTimeUnix ListPronunciationDictionariesRequestSort = "creation_time_unix"
	ListPronunciationDictionariesRequestSortName             ListPronunciationDictionariesRequestSort = "name"
)

func NewListPronunciationDictionariesRequestSortFromString(s string) (ListPronunciationDictionariesRequestSort, error) {
	switch s {
	case "creation_time_unix":
		return ListPronunciationDictionariesRequestSortCreationTimeUnix, nil
	case "name":
		return ListPronunciationDictionariesRequestSortName, nil
	}
	var t ListPronunciationDictionariesRequestSort
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ListPronunciationDictionariesRequestSort) Ptr() *ListPronunciationDictionariesRequestSort {
	return &e
}
