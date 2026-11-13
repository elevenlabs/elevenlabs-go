package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type InternalSamplingConfigResponseModel struct {
	Focus           *float64 `json:"focus,omitempty" xml:"focus,omitempty" min:"0" max:"1"`
	Similarity      *float64 `json:"similarity,omitempty" xml:"similarity,omitempty" min:"0" max:"1"`
	Stability       *float64 `json:"stability,omitempty" xml:"stability,omitempty" min:"0" max:"1"`
	extraProperties map[string]interface{}
}

func (i *InternalSamplingConfigResponseModel) GetFocus() (value *float64) {
	if i == nil {
		return
	}
	return i.Focus
}

func (i *InternalSamplingConfigResponseModel) GetSimilarity() (value *float64) {
	if i == nil {
		return
	}
	return i.Similarity
}

func (i *InternalSamplingConfigResponseModel) GetStability() (value *float64) {
	if i == nil {
		return
	}
	return i.Stability
}

func (i *InternalSamplingConfigResponseModel) SetFocus(focus *float64) {
	i.Focus = focus
}

func (i *InternalSamplingConfigResponseModel) SetSimilarity(similarity *float64) {
	i.Similarity = similarity
}

func (i *InternalSamplingConfigResponseModel) SetStability(stability *float64) {
	i.Stability = stability
}

func (i *InternalSamplingConfigResponseModel) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i InternalSamplingConfigResponseModel) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InternalSamplingConfigResponseModel to string"
	}
	return string(jsonData)
}

func (i *InternalSamplingConfigResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, i); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	return nil
}

func (i *InternalSamplingConfigResponseModel) MarshalJSON() ([]byte, error) {
	type embed InternalSamplingConfigResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*i),
	}
	return json.Marshal(marshaler)
}
