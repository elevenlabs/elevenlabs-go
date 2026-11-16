package types

import (
	"fmt"
)

// The lifecycle status of the generation. It ends at `completed` or `failed`.
type MediaGenerationInProgressResponseStatus string

const (
	MediaGenerationInProgressResponseStatusPending    MediaGenerationInProgressResponseStatus = "pending"
	MediaGenerationInProgressResponseStatusGenerating MediaGenerationInProgressResponseStatus = "generating"
)

func NewMediaGenerationInProgressResponseStatusFromString(s string) (MediaGenerationInProgressResponseStatus, error) {
	switch s {
	case "pending":
		return MediaGenerationInProgressResponseStatusPending, nil
	case "generating":
		return MediaGenerationInProgressResponseStatusGenerating, nil
	}
	var t MediaGenerationInProgressResponseStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e MediaGenerationInProgressResponseStatus) Ptr() *MediaGenerationInProgressResponseStatus {
	return &e
}
