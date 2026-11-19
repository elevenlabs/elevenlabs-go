package types

import (
	"fmt"
)

type EmbedVariant string

const (
	EmbedVariantTiny       EmbedVariant = "tiny"
	EmbedVariantCompact    EmbedVariant = "compact"
	EmbedVariantFull       EmbedVariant = "full"
	EmbedVariantExpandable EmbedVariant = "expandable"
)

func NewEmbedVariantFromString(s string) (EmbedVariant, error) {
	switch s {
	case "tiny":
		return EmbedVariantTiny, nil
	case "compact":
		return EmbedVariantCompact, nil
	case "full":
		return EmbedVariantFull, nil
	case "expandable":
		return EmbedVariantExpandable, nil
	}
	var t EmbedVariant
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e EmbedVariant) Ptr() *EmbedVariant {
	return &e
}
