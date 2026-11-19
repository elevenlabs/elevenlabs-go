package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type WidgetTextContentsTranslation struct {
	// The source text each translated field was derived from
	Source map[string]string `json:"source,omitempty"`
	// The last auto-translated output for each translated field
	Text            map[string]string `json:"text,omitempty"`
	extraProperties map[string]interface{}
}

func (w *WidgetTextContentsTranslation) GetSource() (value map[string]string) {
	if w == nil {
		return
	}
	return w.Source
}

func (w *WidgetTextContentsTranslation) GetText() (value map[string]string) {
	if w == nil {
		return
	}
	return w.Text
}

func (w *WidgetTextContentsTranslation) SetSource(source map[string]string) {
	w.Source = source
}

func (w *WidgetTextContentsTranslation) SetText(text map[string]string) {
	w.Text = text
}

func (w *WidgetTextContentsTranslation) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WidgetTextContentsTranslation) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WidgetTextContentsTranslation to string"
	}
	return string(jsonData)
}

func (w *WidgetTextContentsTranslation) UnmarshalJSON(data []byte) error {
	type unmarshaler WidgetTextContentsTranslation
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*w = WidgetTextContentsTranslation(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WidgetTextContentsTranslation) MarshalJSON() ([]byte, error) {
	type embed WidgetTextContentsTranslation
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
