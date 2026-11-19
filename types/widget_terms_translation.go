package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WidgetTermsTranslation struct {
	SourceHash      string `json:"source_hash" xml:"source_hash" required:"true"`
	Text            string `json:"text" xml:"text" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WidgetTermsTranslation) GetSourceHash() (value string) {
	if w == nil {
		return
	}
	return w.SourceHash
}

func (w *WidgetTermsTranslation) GetText() (value string) {
	if w == nil {
		return
	}
	return w.Text
}

func (w *WidgetTermsTranslation) SetSourceHash(sourceHash string) {
	w.SourceHash = sourceHash
}

func (w *WidgetTermsTranslation) SetText(text string) {
	w.Text = text
}

func (w *WidgetTermsTranslation) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WidgetTermsTranslation) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WidgetTermsTranslation to string"
	}
	return string(jsonData)
}

func (w *WidgetTermsTranslation) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WidgetTermsTranslation
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WidgetTermsTranslation(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WidgetTermsTranslation) MarshalJSON() ([]byte, error) {
	type embed WidgetTermsTranslation
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
