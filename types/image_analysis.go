package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ImageAnalysis struct {
	Status          ImageAnalysisStatus  `json:"status" xml:"status" required:"true"`
	Data            *ImageAnalysisResult `json:"data" xml:"data" required:"true"`
	UpdatedAtMs     *int                 `json:"updated_at_ms,omitempty" xml:"updated_at_ms,omitempty"`
	extraProperties map[string]interface{}
}

func (i *ImageAnalysis) GetStatus() (value ImageAnalysisStatus) {
	if i == nil {
		return
	}
	return i.Status
}

func (i *ImageAnalysis) GetData() (value *ImageAnalysisResult) {
	if i == nil {
		return
	}
	return i.Data
}

func (i *ImageAnalysis) GetUpdatedAtMs() (value *int) {
	if i == nil {
		return
	}
	return i.UpdatedAtMs
}

func (i *ImageAnalysis) SetStatus(status ImageAnalysisStatus) {
	i.Status = status
}

func (i *ImageAnalysis) SetData(data *ImageAnalysisResult) {
	i.Data = data
}

func (i *ImageAnalysis) SetUpdatedAtMs(updatedAtMs *int) {
	i.UpdatedAtMs = updatedAtMs
}

func (i *ImageAnalysis) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i ImageAnalysis) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: ImageAnalysis to string"
	}
	return string(jsonData)
}

func (i *ImageAnalysis) UnmarshalJSON(data []byte) error {
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

func (i *ImageAnalysis) MarshalJSON() ([]byte, error) {
	type embed ImageAnalysis
	var marshaler = struct {
		embed
	}{
		embed: embed(*i),
	}
	return json.Marshal(marshaler)
}
