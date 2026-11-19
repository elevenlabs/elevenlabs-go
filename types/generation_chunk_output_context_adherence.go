package types

import (
	"fmt"
)

// How much the model adheres to the context of its surrounding chunks. Low adherence means the model can deviate from the context and be more creative. High adherence means the model will be more consistent with the context.
type GenerationChunkOutputContextAdherence string

const (
	GenerationChunkOutputContextAdherenceLow    GenerationChunkOutputContextAdherence = "low"
	GenerationChunkOutputContextAdherenceMedium GenerationChunkOutputContextAdherence = "medium"
	GenerationChunkOutputContextAdherenceHigh   GenerationChunkOutputContextAdherence = "high"
)

func NewGenerationChunkOutputContextAdherenceFromString(s string) (GenerationChunkOutputContextAdherence, error) {
	switch s {
	case "low":
		return GenerationChunkOutputContextAdherenceLow, nil
	case "medium":
		return GenerationChunkOutputContextAdherenceMedium, nil
	case "high":
		return GenerationChunkOutputContextAdherenceHigh, nil
	}
	var t GenerationChunkOutputContextAdherence
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e GenerationChunkOutputContextAdherence) Ptr() *GenerationChunkOutputContextAdherence {
	return &e
}
