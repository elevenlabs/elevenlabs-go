package types

import (
	"fmt"
)

// The category of the voice.
type Category string

const (
	CategoryGenerated    Category = "generated"
	CategoryCloned       Category = "cloned"
	CategoryPremade      Category = "premade"
	CategoryProfessional Category = "professional"
	CategoryFamous       Category = "famous"
	CategoryHighQuality  Category = "high_quality"
)

func NewCategoryFromString(s string) (Category, error) {
	switch s {
	case "generated":
		return CategoryGenerated, nil
	case "cloned":
		return CategoryCloned, nil
	case "premade":
		return CategoryPremade, nil
	case "professional":
		return CategoryProfessional, nil
	case "famous":
		return CategoryFamous, nil
	case "high_quality":
		return CategoryHighQuality, nil
	}
	var t Category
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e Category) Ptr() *Category {
	return &e
}
