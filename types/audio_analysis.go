package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AudioAnalysis struct {
	Status          AudioAnalysisStatus  `json:"status" xml:"status" required:"true"`
	Data            *AudioAnalysisResult `json:"data" xml:"data" required:"true"`
	UpdatedAtMs     *int                 `json:"updated_at_ms,omitempty" xml:"updated_at_ms,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AudioAnalysis) GetStatus() (value AudioAnalysisStatus) {
	if a == nil {
		return
	}
	return a.Status
}

func (a *AudioAnalysis) GetData() (value *AudioAnalysisResult) {
	if a == nil {
		return
	}
	return a.Data
}

func (a *AudioAnalysis) GetUpdatedAtMs() (value *int) {
	if a == nil {
		return
	}
	return a.UpdatedAtMs
}

func (a *AudioAnalysis) SetStatus(status AudioAnalysisStatus) {
	a.Status = status
}

func (a *AudioAnalysis) SetData(data *AudioAnalysisResult) {
	a.Data = data
}

func (a *AudioAnalysis) SetUpdatedAtMs(updatedAtMs *int) {
	a.UpdatedAtMs = updatedAtMs
}

func (a *AudioAnalysis) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AudioAnalysis) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AudioAnalysis to string"
	}
	return string(jsonData)
}

func (a *AudioAnalysis) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AudioAnalysis) MarshalJSON() ([]byte, error) {
	type embed AudioAnalysis
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
