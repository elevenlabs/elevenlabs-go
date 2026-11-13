package types

import (
	"fmt"
)

type QualityPresetType string

const (
	QualityPresetTypeStandard      QualityPresetType = "standard"
	QualityPresetTypeHigh          QualityPresetType = "high"
	QualityPresetTypeUltra         QualityPresetType = "ultra"
	QualityPresetTypeUltraLossless QualityPresetType = "ultra_lossless"
)

func NewQualityPresetTypeFromString(s string) (QualityPresetType, error) {
	switch s {
	case "standard":
		return QualityPresetTypeStandard, nil
	case "high":
		return QualityPresetTypeHigh, nil
	case "ultra":
		return QualityPresetTypeUltra, nil
	case "ultra_lossless":
		return QualityPresetTypeUltraLossless, nil
	}
	var t QualityPresetType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e QualityPresetType) Ptr() *QualityPresetType {
	return &e
}
