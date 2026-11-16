package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WatermarkDetails struct {
	// Product that generated the audio, when known.
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
	// Generating model, when known.
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// Generation time, when known.
	GeneratedAt *string `json:"generated_at,omitempty" xml:"generated_at,omitempty"`
	// Identifier of the matched generation record, when known.
	UniqueId        *string `json:"unique_id,omitempty" xml:"unique_id,omitempty"`
	extraProperties map[string]interface{}
}

func (w *WatermarkDetails) GetProduct() (value *string) {
	if w == nil {
		return
	}
	return w.Product
}

func (w *WatermarkDetails) GetModel() (value *string) {
	if w == nil {
		return
	}
	return w.Model
}

func (w *WatermarkDetails) GetGeneratedAt() (value *string) {
	if w == nil {
		return
	}
	return w.GeneratedAt
}

func (w *WatermarkDetails) GetUniqueId() (value *string) {
	if w == nil {
		return
	}
	return w.UniqueId
}

func (w *WatermarkDetails) SetProduct(product *string) {
	w.Product = product
}

func (w *WatermarkDetails) SetModel(model *string) {
	w.Model = model
}

func (w *WatermarkDetails) SetGeneratedAt(generatedAt *string) {
	w.GeneratedAt = generatedAt
}

func (w *WatermarkDetails) SetUniqueId(uniqueId *string) {
	w.UniqueId = uniqueId
}

func (w *WatermarkDetails) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WatermarkDetails) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WatermarkDetails to string"
	}
	return string(jsonData)
}

func (w *WatermarkDetails) UnmarshalJSON(data []byte) error {
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

func (w *WatermarkDetails) MarshalJSON() ([]byte, error) {
	type embed WatermarkDetails
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
