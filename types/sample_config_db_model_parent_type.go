package types

import (
	"fmt"
)

type SampleConfigDbModelParentType string

const (
	SampleConfigDbModelParentTypeRead       SampleConfigDbModelParentType = "read"
	SampleConfigDbModelParentTypeCollection SampleConfigDbModelParentType = "collection"
)

func NewSampleConfigDbModelParentTypeFromString(s string) (SampleConfigDbModelParentType, error) {
	switch s {
	case "read":
		return SampleConfigDbModelParentTypeRead, nil
	case "collection":
		return SampleConfigDbModelParentTypeCollection, nil
	}
	var t SampleConfigDbModelParentType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SampleConfigDbModelParentType) Ptr() *SampleConfigDbModelParentType {
	return &e
}
