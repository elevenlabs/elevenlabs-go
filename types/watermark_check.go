package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WatermarkCheck struct {
	// True when an ElevenLabs watermark was found.
	IsDetected bool `json:"is_detected" xml:"is_detected" required:"true"`
	// Matched generation record. Null when none matched.
	Details         *WatermarkDetails `json:"details" xml:"details" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WatermarkCheck) GetIsDetected() (value bool) {
	if w == nil {
		return
	}
	return w.IsDetected
}

func (w *WatermarkCheck) GetDetails() (value *WatermarkDetails) {
	if w == nil {
		return
	}
	return w.Details
}

func (w *WatermarkCheck) SetIsDetected(isDetected bool) {
	w.IsDetected = isDetected
}

func (w *WatermarkCheck) SetDetails(details *WatermarkDetails) {
	w.Details = details
}

func (w *WatermarkCheck) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WatermarkCheck) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WatermarkCheck to string"
	}
	return string(jsonData)
}

func (w *WatermarkCheck) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, w); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WatermarkCheck) MarshalJSON() ([]byte, error) {
	type embed WatermarkCheck
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
