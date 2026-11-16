package music

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
)

type CreateCompositionPlanResponse struct {
	MusicPrompt     *types.MusicPrompt
	CompositionPlan *types.CompositionPlan
	typ             string
}

func (c *CreateCompositionPlanResponse) GetMusicPrompt() *types.MusicPrompt {
	if c == nil {
		return nil
	}
	if c.MusicPrompt == nil {
		return nil
	}
	return c.MusicPrompt
}

func (c *CreateCompositionPlanResponse) GetCompositionPlan() *types.CompositionPlan {
	if c == nil {
		return nil
	}
	if c.CompositionPlan == nil {
		return nil
	}
	return c.CompositionPlan
}

func (c *CreateCompositionPlanResponse) UnmarshalJSON(data []byte) error {
	var valueMusicPrompt *types.MusicPrompt
	if err := json.Unmarshal(data, &valueMusicPrompt); err == nil {
		c.typ = "MusicPrompt"
		c.MusicPrompt = valueMusicPrompt
		return nil
	}
	var valueCompositionPlan *types.CompositionPlan
	if err := json.Unmarshal(data, &valueCompositionPlan); err == nil {
		c.typ = "CompositionPlan"
		c.CompositionPlan = valueCompositionPlan
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, c)
}

func (c CreateCompositionPlanResponse) MarshalJSON() ([]byte, error) {
	if c.typ == "MusicPrompt" || c.MusicPrompt != nil {
		return json.Marshal(c.MusicPrompt)
	}
	if c.typ == "CompositionPlan" || c.CompositionPlan != nil {
		return json.Marshal(c.CompositionPlan)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", c)
}

type CreateCompositionPlanResponseVisitor interface {
	VisitMusicPrompt(*types.MusicPrompt) error
	VisitCompositionPlan(*types.CompositionPlan) error
}

func (c *CreateCompositionPlanResponse) Accept(visitor CreateCompositionPlanResponseVisitor) error {
	if c.typ == "MusicPrompt" || c.MusicPrompt != nil {
		return visitor.VisitMusicPrompt(c.MusicPrompt)
	}
	if c.typ == "CompositionPlan" || c.CompositionPlan != nil {
		return visitor.VisitCompositionPlan(c.CompositionPlan)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", c)
}

type BodyGenerateCompositionPlanV1MusicPlanPost struct {
	// When enable_logging is set to false zero retention mode will be used for the request. Zero retention mode may only be used by enterprise customers.
	EnableLogging *bool `json:"-" url:"enable_logging,omitempty" explode:"true" serializationStyle:"form" queryParam:"enable_logging"`
	// A simple text prompt to compose a plan from.
	Prompt string `json:"prompt" xml:"prompt" required:"true" maxLength:"4100"`
	// The length of the composition plan to generate in milliseconds. Must be between 3000ms and 600000ms. Optional - if not provided, the model will choose a length based on the prompt.
	MusicLengthMs *int `json:"music_length_ms,omitempty" xml:"music_length_ms,omitempty" min:"3000" max:"600000"`
	// An optional composition plan to use as a source for the new composition plan.
	SourceCompositionPlan *BodyGenerateCompositionPlanV1MusicPlanPostSourceCompositionPlan `json:"source_composition_plan,omitempty" xml:"source_composition_plan,omitempty"`
	ModelId               *types.MusicModelId                                              `json:"model_id,omitempty" xml:"model_id,omitempty"`
}

func (b *BodyGenerateCompositionPlanV1MusicPlanPost) GetEnableLogging() (value *bool) {
	if b == nil {
		return
	}
	return b.EnableLogging
}

func (b *BodyGenerateCompositionPlanV1MusicPlanPost) GetPrompt() (value string) {
	if b == nil {
		return
	}
	return b.Prompt
}

func (b *BodyGenerateCompositionPlanV1MusicPlanPost) GetMusicLengthMs() (value *int) {
	if b == nil {
		return
	}
	return b.MusicLengthMs
}

func (b *BodyGenerateCompositionPlanV1MusicPlanPost) GetSourceCompositionPlan() (value *BodyGenerateCompositionPlanV1MusicPlanPostSourceCompositionPlan) {
	if b == nil {
		return
	}
	return b.SourceCompositionPlan
}

func (b *BodyGenerateCompositionPlanV1MusicPlanPost) GetModelId() (value *types.MusicModelId) {
	if b == nil {
		return
	}
	return b.ModelId
}

func (b *BodyGenerateCompositionPlanV1MusicPlanPost) SetEnableLogging(enableLogging *bool) {
	b.EnableLogging = enableLogging
}

func (b *BodyGenerateCompositionPlanV1MusicPlanPost) SetPrompt(prompt string) {
	b.Prompt = prompt
}

func (b *BodyGenerateCompositionPlanV1MusicPlanPost) SetMusicLengthMs(musicLengthMs *int) {
	b.MusicLengthMs = musicLengthMs
}

func (b *BodyGenerateCompositionPlanV1MusicPlanPost) SetSourceCompositionPlan(sourceCompositionPlan *BodyGenerateCompositionPlanV1MusicPlanPostSourceCompositionPlan) {
	b.SourceCompositionPlan = sourceCompositionPlan
}

func (b *BodyGenerateCompositionPlanV1MusicPlanPost) SetModelId(modelId *types.MusicModelId) {
	b.ModelId = modelId
}

func (b *BodyGenerateCompositionPlanV1MusicPlanPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyGenerateCompositionPlanV1MusicPlanPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyGenerateCompositionPlanV1MusicPlanPost to string"
	}
	return string(jsonData)
}

func (b *BodyGenerateCompositionPlanV1MusicPlanPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyGenerateCompositionPlanV1MusicPlanPost) MarshalJSON() ([]byte, error) {
	type embed BodyGenerateCompositionPlanV1MusicPlanPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

// An optional composition plan to use as a source for the new composition plan.
type BodyGenerateCompositionPlanV1MusicPlanPostSourceCompositionPlan struct {
	MusicPrompt     *types.MusicPrompt
	CompositionPlan *types.CompositionPlan
	typ             string
}

func (b *BodyGenerateCompositionPlanV1MusicPlanPostSourceCompositionPlan) GetMusicPrompt() *types.MusicPrompt {
	if b == nil {
		return nil
	}
	if b.MusicPrompt == nil {
		return nil
	}
	return b.MusicPrompt
}

func (b *BodyGenerateCompositionPlanV1MusicPlanPostSourceCompositionPlan) GetCompositionPlan() *types.CompositionPlan {
	if b == nil {
		return nil
	}
	if b.CompositionPlan == nil {
		return nil
	}
	return b.CompositionPlan
}

func (b *BodyGenerateCompositionPlanV1MusicPlanPostSourceCompositionPlan) UnmarshalJSON(data []byte) error {
	var valueMusicPrompt *types.MusicPrompt
	if err := json.Unmarshal(data, &valueMusicPrompt); err == nil {
		b.typ = "MusicPrompt"
		b.MusicPrompt = valueMusicPrompt
		return nil
	}
	var valueCompositionPlan *types.CompositionPlan
	if err := json.Unmarshal(data, &valueCompositionPlan); err == nil {
		b.typ = "CompositionPlan"
		b.CompositionPlan = valueCompositionPlan
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, b)
}

func (b BodyGenerateCompositionPlanV1MusicPlanPostSourceCompositionPlan) MarshalJSON() ([]byte, error) {
	if b.typ == "MusicPrompt" || b.MusicPrompt != nil {
		return json.Marshal(b.MusicPrompt)
	}
	if b.typ == "CompositionPlan" || b.CompositionPlan != nil {
		return json.Marshal(b.CompositionPlan)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BodyGenerateCompositionPlanV1MusicPlanPostSourceCompositionPlanVisitor interface {
	VisitMusicPrompt(*types.MusicPrompt) error
	VisitCompositionPlan(*types.CompositionPlan) error
}

func (b *BodyGenerateCompositionPlanV1MusicPlanPostSourceCompositionPlan) Accept(visitor BodyGenerateCompositionPlanV1MusicPlanPostSourceCompositionPlanVisitor) error {
	if b.typ == "MusicPrompt" || b.MusicPrompt != nil {
		return visitor.VisitMusicPrompt(b.MusicPrompt)
	}
	if b.typ == "CompositionPlan" || b.CompositionPlan != nil {
		return visitor.VisitCompositionPlan(b.CompositionPlan)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", b)
}
