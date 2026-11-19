package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Composition plan for the `music_v2` and `music_v2_5` models. Using this field with any other model will result in an error.
type CompositionPlan struct {
	// The chunks that make up the generation.
	Chunks          []*CompositionPlanChunksItem `json:"chunks" xml:"chunks" required:"true" maxItems:"30"`
	extraProperties map[string]interface{}
}

func (c *CompositionPlan) GetChunks() (value []*CompositionPlanChunksItem) {
	if c == nil {
		return
	}
	return c.Chunks
}

func (c *CompositionPlan) SetChunks(chunks []*CompositionPlanChunksItem) {
	c.Chunks = chunks
}

func (c *CompositionPlan) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CompositionPlan) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CompositionPlan to string"
	}
	return string(jsonData)
}

func (c *CompositionPlan) UnmarshalJSON(data []byte) error {
	type unmarshaler CompositionPlan
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CompositionPlan(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CompositionPlan) MarshalJSON() ([]byte, error) {
	type embed CompositionPlan
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
