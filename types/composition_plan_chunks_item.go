package types

import (
	"encoding/json"
	"fmt"
)

type CompositionPlanChunksItem struct {
	GenerationChunkInput *GenerationChunkInput
	AudioRefChunk        *AudioRefChunk
	typ                  string
}

func (c *CompositionPlanChunksItem) GetGenerationChunkInput() *GenerationChunkInput {
	if c == nil {
		return nil
	}
	if c.GenerationChunkInput == nil {
		return nil
	}
	return c.GenerationChunkInput
}

func (c *CompositionPlanChunksItem) GetAudioRefChunk() *AudioRefChunk {
	if c == nil {
		return nil
	}
	if c.AudioRefChunk == nil {
		return nil
	}
	return c.AudioRefChunk
}

func (c *CompositionPlanChunksItem) UnmarshalJSON(data []byte) error {
	var valueGenerationChunkInput *GenerationChunkInput
	if err := json.Unmarshal(data, &valueGenerationChunkInput); err == nil {
		c.typ = "GenerationChunkInput"
		c.GenerationChunkInput = valueGenerationChunkInput
		return nil
	}
	var valueAudioRefChunk *AudioRefChunk
	if err := json.Unmarshal(data, &valueAudioRefChunk); err == nil {
		c.typ = "AudioRefChunk"
		c.AudioRefChunk = valueAudioRefChunk
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, c)
}

func (c CompositionPlanChunksItem) MarshalJSON() ([]byte, error) {
	if c.typ == "GenerationChunkInput" || c.GenerationChunkInput != nil {
		return json.Marshal(c.GenerationChunkInput)
	}
	if c.typ == "AudioRefChunk" || c.AudioRefChunk != nil {
		return json.Marshal(c.AudioRefChunk)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", c)
}

type CompositionPlanChunksItemVisitor interface {
	VisitGenerationChunkInput(*GenerationChunkInput) error
	VisitAudioRefChunk(*AudioRefChunk) error
}

func (c *CompositionPlanChunksItem) Accept(visitor CompositionPlanChunksItemVisitor) error {
	if c.typ == "GenerationChunkInput" || c.GenerationChunkInput != nil {
		return visitor.VisitGenerationChunkInput(c.GenerationChunkInput)
	}
	if c.typ == "AudioRefChunk" || c.AudioRefChunk != nil {
		return visitor.VisitAudioRefChunk(c.AudioRefChunk)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", c)
}
