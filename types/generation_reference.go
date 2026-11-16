package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// The output of a prior generation on this API.
type GenerationReference struct {
	type_ string
	// The ID of the generation whose output to use, as returned when the generation was created.
	GenerationId    string `json:"generation_id" xml:"generation_id" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GenerationReference) GetGenerationId() (value string) {
	if g == nil {
		return
	}
	return g.GenerationId
}

func (g *GenerationReference) SetGenerationId(generationId string) {
	g.GenerationId = generationId
}

func (g *GenerationReference) Type() string {
	return g.type_
}

func (g *GenerationReference) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GenerationReference) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GenerationReference to string"
	}
	return string(jsonData)
}

func (g *GenerationReference) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, g); err != nil {
		return err
	}
	type alias GenerationReference
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*g = GenerationReference(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "generation" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", g, "generation", *literals.Type)
	}
	g.type_ = "generation"
	extraProperties, err := internal.ExtractExtraProperties(data, *g, "type")
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GenerationReference) MarshalJSON() ([]byte, error) {
	type embed GenerationReference
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*g),
		Type:  "generation",
	}
	return json.Marshal(marshaler)
}
