package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type VoiceSampleVisualWaveformResponseModel struct {
	// The ID of the sample.
	SampleId string `json:"sample_id" xml:"sample_id" required:"true"`
	// The visual waveform of the sample, represented as a list of floats.
	VisualWaveform  []float64 `json:"visual_waveform" xml:"visual_waveform" required:"true"`
	extraProperties map[string]interface{}
}

func (v *VoiceSampleVisualWaveformResponseModel) GetSampleId() (value string) {
	if v == nil {
		return
	}
	return v.SampleId
}

func (v *VoiceSampleVisualWaveformResponseModel) GetVisualWaveform() (value []float64) {
	if v == nil {
		return
	}
	return v.VisualWaveform
}

func (v *VoiceSampleVisualWaveformResponseModel) SetSampleId(sampleId string) {
	v.SampleId = sampleId
}

func (v *VoiceSampleVisualWaveformResponseModel) SetVisualWaveform(visualWaveform []float64) {
	v.VisualWaveform = visualWaveform
}

func (v *VoiceSampleVisualWaveformResponseModel) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v VoiceSampleVisualWaveformResponseModel) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VoiceSampleVisualWaveformResponseModel to string"
	}
	return string(jsonData)
}

func (v *VoiceSampleVisualWaveformResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, v); err != nil {
		return err
	}
	type alias VoiceSampleVisualWaveformResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*v = VoiceSampleVisualWaveformResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	return nil
}

func (v *VoiceSampleVisualWaveformResponseModel) MarshalJSON() ([]byte, error) {
	type embed VoiceSampleVisualWaveformResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
