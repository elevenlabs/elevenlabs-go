package types

import (
	"fmt"
)

type AssetTranscriptionStatus string

const (
	AssetTranscriptionStatusProcessing AssetTranscriptionStatus = "processing"
	AssetTranscriptionStatusCompleted  AssetTranscriptionStatus = "completed"
	AssetTranscriptionStatusFailed     AssetTranscriptionStatus = "failed"
)

func NewAssetTranscriptionStatusFromString(s string) (AssetTranscriptionStatus, error) {
	switch s {
	case "processing":
		return AssetTranscriptionStatusProcessing, nil
	case "completed":
		return AssetTranscriptionStatusCompleted, nil
	case "failed":
		return AssetTranscriptionStatusFailed, nil
	}
	var t AssetTranscriptionStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AssetTranscriptionStatus) Ptr() *AssetTranscriptionStatus {
	return &e
}
