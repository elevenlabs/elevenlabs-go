package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type VideoAnalysis struct {
	Status          VideoAnalysisStatus  `json:"status" xml:"status" required:"true"`
	Data            *VideoAnalysisResult `json:"data" xml:"data" required:"true"`
	UpdatedAtMs     *int                 `json:"updated_at_ms,omitempty" xml:"updated_at_ms,omitempty"`
	extraProperties map[string]interface{}
}

func (v *VideoAnalysis) GetStatus() (value VideoAnalysisStatus) {
	if v == nil {
		return
	}
	return v.Status
}

func (v *VideoAnalysis) GetData() (value *VideoAnalysisResult) {
	if v == nil {
		return
	}
	return v.Data
}

func (v *VideoAnalysis) GetUpdatedAtMs() (value *int) {
	if v == nil {
		return
	}
	return v.UpdatedAtMs
}

func (v *VideoAnalysis) SetStatus(status VideoAnalysisStatus) {
	v.Status = status
}

func (v *VideoAnalysis) SetData(data *VideoAnalysisResult) {
	v.Data = data
}

func (v *VideoAnalysis) SetUpdatedAtMs(updatedAtMs *int) {
	v.UpdatedAtMs = updatedAtMs
}

func (v *VideoAnalysis) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v VideoAnalysis) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VideoAnalysis to string"
	}
	return string(jsonData)
}

func (v *VideoAnalysis) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, v); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	return nil
}

func (v *VideoAnalysis) MarshalJSON() ([]byte, error) {
	type embed VideoAnalysis
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
