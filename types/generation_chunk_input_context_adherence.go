package types

import (
	"fmt"
)

// How much the model adheres to the context of its surrounding chunks. Low adherence means the model can deviate from the context and be more creative. High adherence means the model will be more consistent with the context.
type GenerationChunkInputContextAdherence string

const (
	GenerationChunkInputContextAdherenceLow    GenerationChunkInputContextAdherence = "low"
	GenerationChunkInputContextAdherenceMedium GenerationChunkInputContextAdherence = "medium"
	GenerationChunkInputContextAdherenceHigh   GenerationChunkInputContextAdherence = "high"
)

func NewGenerationChunkInputContextAdherenceFromString(s string) (GenerationChunkInputContextAdherence, error) {
	switch s {
	case "low":
		return GenerationChunkInputContextAdherenceLow, nil
	case "medium":
		return GenerationChunkInputContextAdherenceMedium, nil
	case "high":
		return GenerationChunkInputContextAdherenceHigh, nil
	}
	var t GenerationChunkInputContextAdherence
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e GenerationChunkInputContextAdherence) Ptr() *GenerationChunkInputContextAdherence {
	return &e
}
