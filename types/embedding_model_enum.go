package types

import (
	"fmt"
)

type EmbeddingModelEnum string

const (
	EmbeddingModelEnumE5Mistral7BInstruct         EmbeddingModelEnum = "e5_mistral_7b_instruct"
	EmbeddingModelEnumMultilingualE5LargeInstruct EmbeddingModelEnum = "multilingual_e5_large_instruct"
)

func NewEmbeddingModelEnumFromString(s string) (EmbeddingModelEnum, error) {
	switch s {
	case "e5_mistral_7b_instruct":
		return EmbeddingModelEnumE5Mistral7BInstruct, nil
	case "multilingual_e5_large_instruct":
		return EmbeddingModelEnumMultilingualE5LargeInstruct, nil
	}
	var t EmbeddingModelEnum
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e EmbeddingModelEnum) Ptr() *EmbeddingModelEnum {
	return &e
}
