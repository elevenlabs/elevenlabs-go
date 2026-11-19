package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ToDialogueSettingsResponseModel struct {
	// Determines how stable the voice is and the randomness between each generation. Lower values introduce broader emotional range for the voice. Higher values can result in a monotonous voice with limited emotion.
	Stability *float64 `json:"stability,omitempty" xml:"stability,omitempty" min:"0" max:"1"`
	// Determines how strongly the model is guided while generating. Higher values make the model adhere more closely to the voice, at the cost of variation. Not supported by every model.
	Similarity *float64 `json:"similarity,omitempty" xml:"similarity,omitempty" min:"0" max:"1"`
	// Controls how closely the output matches your professional voice clone samples (0–1). Lower values sound more versatile, while higher values preserve more of the original voice characteristics. We recommend 0 for most use cases, but if you do run into issues 0.5 could be a good fallback. Applies only to professional voice clones fine-tuned for the selected v4 model.
	CloneStrength   *float64 `json:"clone_strength,omitempty" xml:"clone_strength,omitempty" min:"0" max:"1"`
	extraProperties map[string]interface{}
}

func (t *ToDialogueSettingsResponseModel) GetStability() (value *float64) {
	if t == nil {
		return
	}
	return t.Stability
}

func (t *ToDialogueSettingsResponseModel) GetSimilarity() (value *float64) {
	if t == nil {
		return
	}
	return t.Similarity
}

func (t *ToDialogueSettingsResponseModel) GetCloneStrength() (value *float64) {
	if t == nil {
		return
	}
	return t.CloneStrength
}

func (t *ToDialogueSettingsResponseModel) SetStability(stability *float64) {
	t.Stability = stability
}

func (t *ToDialogueSettingsResponseModel) SetSimilarity(similarity *float64) {
	t.Similarity = similarity
}

func (t *ToDialogueSettingsResponseModel) SetCloneStrength(cloneStrength *float64) {
	t.CloneStrength = cloneStrength
}

func (t *ToDialogueSettingsResponseModel) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t ToDialogueSettingsResponseModel) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: ToDialogueSettingsResponseModel to string"
	}
	return string(jsonData)
}

func (t *ToDialogueSettingsResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *ToDialogueSettingsResponseModel) MarshalJSON() ([]byte, error) {
	type embed ToDialogueSettingsResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
