package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GenerationSourceContext struct {
	SourceType          *string         `json:"source_type,omitempty" xml:"source_type,omitempty"`
	GenerationId        string          `json:"generation_id" xml:"generation_id" required:"true"`
	Prompt              *string         `json:"prompt,omitempty" xml:"prompt,omitempty"`
	ModelId             string          `json:"model_id" xml:"model_id" required:"true"`
	ModelProvider       *string         `json:"model_provider,omitempty" xml:"model_provider,omitempty"`
	GenerationSessionId *string         `json:"generation_session_id,omitempty" xml:"generation_session_id,omitempty"`
	SessionIterationId  *string         `json:"session_iteration_id,omitempty" xml:"session_iteration_id,omitempty"`
	ModelParameters     map[string]any  `json:"model_parameters,omitempty"`
	ExtendVideo         *ReferenceVideo `json:"extend_video,omitempty" xml:"extend_video,omitempty"`
	extraProperties     map[string]interface{}
}

func (g *GenerationSourceContext) GetSourceType() (value *string) {
	if g == nil {
		return
	}
	return g.SourceType
}

func (g *GenerationSourceContext) GetGenerationId() (value string) {
	if g == nil {
		return
	}
	return g.GenerationId
}

func (g *GenerationSourceContext) GetPrompt() (value *string) {
	if g == nil {
		return
	}
	return g.Prompt
}

func (g *GenerationSourceContext) GetModelId() (value string) {
	if g == nil {
		return
	}
	return g.ModelId
}

func (g *GenerationSourceContext) GetModelProvider() (value *string) {
	if g == nil {
		return
	}
	return g.ModelProvider
}

func (g *GenerationSourceContext) GetGenerationSessionId() (value *string) {
	if g == nil {
		return
	}
	return g.GenerationSessionId
}

func (g *GenerationSourceContext) GetSessionIterationId() (value *string) {
	if g == nil {
		return
	}
	return g.SessionIterationId
}

func (g *GenerationSourceContext) GetModelParameters() (value map[string]any) {
	if g == nil {
		return
	}
	return g.ModelParameters
}

func (g *GenerationSourceContext) GetExtendVideo() (value *ReferenceVideo) {
	if g == nil {
		return
	}
	return g.ExtendVideo
}

func (g *GenerationSourceContext) SetSourceType(sourceType *string) {
	g.SourceType = sourceType
}

func (g *GenerationSourceContext) SetGenerationId(generationId string) {
	g.GenerationId = generationId
}

func (g *GenerationSourceContext) SetPrompt(prompt *string) {
	g.Prompt = prompt
}

func (g *GenerationSourceContext) SetModelId(modelId string) {
	g.ModelId = modelId
}

func (g *GenerationSourceContext) SetModelProvider(modelProvider *string) {
	g.ModelProvider = modelProvider
}

func (g *GenerationSourceContext) SetGenerationSessionId(generationSessionId *string) {
	g.GenerationSessionId = generationSessionId
}

func (g *GenerationSourceContext) SetSessionIterationId(sessionIterationId *string) {
	g.SessionIterationId = sessionIterationId
}

func (g *GenerationSourceContext) SetModelParameters(modelParameters map[string]any) {
	g.ModelParameters = modelParameters
}

func (g *GenerationSourceContext) SetExtendVideo(extendVideo *ReferenceVideo) {
	g.ExtendVideo = extendVideo
}

func (g *GenerationSourceContext) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GenerationSourceContext) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GenerationSourceContext to string"
	}
	return string(jsonData)
}

func (g *GenerationSourceContext) UnmarshalJSON(data []byte) error {
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

func (g *GenerationSourceContext) MarshalJSON() ([]byte, error) {
	type embed GenerationSourceContext
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
