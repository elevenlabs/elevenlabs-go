package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type WidgetEndFeedbackConfig struct {
	Type            *WidgetEndFeedbackType `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties map[string]interface{}
}

func (w *WidgetEndFeedbackConfig) GetType() (value *WidgetEndFeedbackType) {
	if w == nil {
		return
	}
	return w.Type
}

func (w *WidgetEndFeedbackConfig) SetType(type_ *WidgetEndFeedbackType) {
	w.Type = type_
}

func (w *WidgetEndFeedbackConfig) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WidgetEndFeedbackConfig) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WidgetEndFeedbackConfig to string"
	}
	return string(jsonData)
}

func (w *WidgetEndFeedbackConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler WidgetEndFeedbackConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*w = WidgetEndFeedbackConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WidgetEndFeedbackConfig) MarshalJSON() ([]byte, error) {
	type embed WidgetEndFeedbackConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
