package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GenerationChunkInput struct {
	// The text config to be generated for this chunk. Can contain an optional section name in square brackets at the beginning, e.g. [Verse 1], lyrics lines, and inline directions in curly braces, e.g. {scratching}. Section names must be between 1 and 100 characters. At most 30 lines are allowed, each at most 200 characters.
	Text string `json:"text" xml:"text" required:"true" maxLength:"6132"`
	// The duration of the chunk in milliseconds. Must be between 3000ms and 120000ms.
	DurationMs int `json:"duration_ms" xml:"duration_ms" required:"true" min:"3000" max:"120000"`
	// The styles and musical directions that should be present in this chunk. Use English language for best results. The styles for the first chunk are the most important as they set the overall tone and genre. Styles for subsequent chunks can be used to add nuance, progression, emphasis, or change the direction of the song. Aim to have at least 6-7 styles in early chunks until the direction is established. Generic styles like 'great production quality' are good default styles to append to the list.
	PositiveStyles []string `json:"positive_styles" xml:"positive_styles" required:"true" maxItems:"50"`
	// The styles and musical directions that should not be present in this chunk. Use English language for best results. Leaving empty is a good default, only use this field if you want to explicitly avoid a particular style or direction.
	NegativeStyles []string `json:"negative_styles,omitempty" xml:"negative_styles,omitempty" maxItems:"50"`
	// How much the model adheres to the context of its surrounding chunks. Low adherence means the model can deviate from the context and be more creative. High adherence means the model will be more consistent with the context.
	ContextAdherence *GenerationChunkInputContextAdherence `json:"context_adherence,omitempty" xml:"context_adherence,omitempty"`
	// The audio reference to condition the generation on. The first chunk is the most important as it will influence the generation of all subsequent chunks. Thus, if you want to apply conditioning to the entire song, start conditioning from the first chunk.
	ConditioningRef *AudioRefChunk `json:"conditioning_ref,omitempty" xml:"conditioning_ref,omitempty"`
	// How strongly the model adheres to the conditioning reference. Low strength means the model will be more creative and deviate from the reference. High strength means the model will be more consistent with the reference.
	ConditionStrength *ConditionStrength `json:"condition_strength,omitempty" xml:"condition_strength,omitempty"`
	extraProperties   map[string]interface{}
}

func (g *GenerationChunkInput) GetText() (value string) {
	if g == nil {
		return
	}
	return g.Text
}

func (g *GenerationChunkInput) GetDurationMs() (value int) {
	if g == nil {
		return
	}
	return g.DurationMs
}

func (g *GenerationChunkInput) GetPositiveStyles() (value []string) {
	if g == nil {
		return
	}
	return g.PositiveStyles
}

func (g *GenerationChunkInput) GetNegativeStyles() (value []string) {
	if g == nil {
		return
	}
	return g.NegativeStyles
}

func (g *GenerationChunkInput) GetContextAdherence() (value *GenerationChunkInputContextAdherence) {
	if g == nil {
		return
	}
	return g.ContextAdherence
}

func (g *GenerationChunkInput) GetConditioningRef() (value *AudioRefChunk) {
	if g == nil {
		return
	}
	return g.ConditioningRef
}

func (g *GenerationChunkInput) GetConditionStrength() (value *ConditionStrength) {
	if g == nil {
		return
	}
	return g.ConditionStrength
}

func (g *GenerationChunkInput) SetText(text string) {
	g.Text = text
}

func (g *GenerationChunkInput) SetDurationMs(durationMs int) {
	g.DurationMs = durationMs
}

func (g *GenerationChunkInput) SetPositiveStyles(positiveStyles []string) {
	g.PositiveStyles = positiveStyles
}

func (g *GenerationChunkInput) SetNegativeStyles(negativeStyles []string) {
	g.NegativeStyles = negativeStyles
}

func (g *GenerationChunkInput) SetContextAdherence(contextAdherence *GenerationChunkInputContextAdherence) {
	g.ContextAdherence = contextAdherence
}

func (g *GenerationChunkInput) SetConditioningRef(conditioningRef *AudioRefChunk) {
	g.ConditioningRef = conditioningRef
}

func (g *GenerationChunkInput) SetConditionStrength(conditionStrength *ConditionStrength) {
	g.ConditionStrength = conditionStrength
}

func (g *GenerationChunkInput) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GenerationChunkInput) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GenerationChunkInput to string"
	}
	return string(jsonData)
}

func (g *GenerationChunkInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, g); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GenerationChunkInput) MarshalJSON() ([]byte, error) {
	type embed GenerationChunkInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
