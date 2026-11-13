package types

import (
	"fmt"
)

// The id of the stem variation to use.
type SeparateStemsMusicRequestStemVariationId string

const (
	SeparateStemsMusicRequestStemVariationIdTwoStemsV1 SeparateStemsMusicRequestStemVariationId = "two_stems_v1"
	SeparateStemsMusicRequestStemVariationIdSixStemsV1 SeparateStemsMusicRequestStemVariationId = "six_stems_v1"
)

func NewSeparateStemsMusicRequestStemVariationIdFromString(s string) (SeparateStemsMusicRequestStemVariationId, error) {
	switch s {
	case "two_stems_v1":
		return SeparateStemsMusicRequestStemVariationIdTwoStemsV1, nil
	case "six_stems_v1":
		return SeparateStemsMusicRequestStemVariationIdSixStemsV1, nil
	}
	var t SeparateStemsMusicRequestStemVariationId
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SeparateStemsMusicRequestStemVariationId) Ptr() *SeparateStemsMusicRequestStemVariationId {
	return &e
}
