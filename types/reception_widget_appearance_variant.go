package types

import (
	"fmt"
)

// Widget size, from the smallest ('tiny') to the largest ('full').
type ReceptionWidgetAppearanceVariant string

const (
	ReceptionWidgetAppearanceVariantTiny    ReceptionWidgetAppearanceVariant = "tiny"
	ReceptionWidgetAppearanceVariantCompact ReceptionWidgetAppearanceVariant = "compact"
	ReceptionWidgetAppearanceVariantFull    ReceptionWidgetAppearanceVariant = "full"
)

func NewReceptionWidgetAppearanceVariantFromString(s string) (ReceptionWidgetAppearanceVariant, error) {
	switch s {
	case "tiny":
		return ReceptionWidgetAppearanceVariantTiny, nil
	case "compact":
		return ReceptionWidgetAppearanceVariantCompact, nil
	case "full":
		return ReceptionWidgetAppearanceVariantFull, nil
	}
	var t ReceptionWidgetAppearanceVariant
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ReceptionWidgetAppearanceVariant) Ptr() *ReceptionWidgetAppearanceVariant {
	return &e
}
