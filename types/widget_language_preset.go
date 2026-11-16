package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WidgetLanguagePreset struct {
	// The text contents for the selected language
	TextContents *WidgetTextContents `json:"text_contents,omitempty" xml:"text_contents,omitempty"`
	// The translation cache for the text contents
	TextContentsTranslation *WidgetTextContentsTranslation `json:"text_contents_translation,omitempty" xml:"text_contents_translation,omitempty"`
	// The text to display for terms and conditions in this language
	TermsText *string `json:"terms_text,omitempty" xml:"terms_text,omitempty"`
	// The HTML to display for terms and conditions in this language
	TermsHtml *string `json:"terms_html,omitempty" xml:"terms_html,omitempty"`
	// The key to display for terms and conditions in this language
	TermsKey *string `json:"terms_key,omitempty" xml:"terms_key,omitempty"`
	// The translation cache for the terms
	TermsTranslation *WidgetTermsTranslation `json:"terms_translation,omitempty" xml:"terms_translation,omitempty"`
	extraProperties  map[string]interface{}
}

func (w *WidgetLanguagePreset) GetTextContents() (value *WidgetTextContents) {
	if w == nil {
		return
	}
	return w.TextContents
}

func (w *WidgetLanguagePreset) GetTextContentsTranslation() (value *WidgetTextContentsTranslation) {
	if w == nil {
		return
	}
	return w.TextContentsTranslation
}

func (w *WidgetLanguagePreset) GetTermsText() (value *string) {
	if w == nil {
		return
	}
	return w.TermsText
}

func (w *WidgetLanguagePreset) GetTermsHtml() (value *string) {
	if w == nil {
		return
	}
	return w.TermsHtml
}

func (w *WidgetLanguagePreset) GetTermsKey() (value *string) {
	if w == nil {
		return
	}
	return w.TermsKey
}

func (w *WidgetLanguagePreset) GetTermsTranslation() (value *WidgetTermsTranslation) {
	if w == nil {
		return
	}
	return w.TermsTranslation
}

func (w *WidgetLanguagePreset) SetTextContents(textContents *WidgetTextContents) {
	w.TextContents = textContents
}

func (w *WidgetLanguagePreset) SetTextContentsTranslation(textContentsTranslation *WidgetTextContentsTranslation) {
	w.TextContentsTranslation = textContentsTranslation
}

func (w *WidgetLanguagePreset) SetTermsText(termsText *string) {
	w.TermsText = termsText
}

func (w *WidgetLanguagePreset) SetTermsHtml(termsHtml *string) {
	w.TermsHtml = termsHtml
}

func (w *WidgetLanguagePreset) SetTermsKey(termsKey *string) {
	w.TermsKey = termsKey
}

func (w *WidgetLanguagePreset) SetTermsTranslation(termsTranslation *WidgetTermsTranslation) {
	w.TermsTranslation = termsTranslation
}

func (w *WidgetLanguagePreset) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WidgetLanguagePreset) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WidgetLanguagePreset to string"
	}
	return string(jsonData)
}

func (w *WidgetLanguagePreset) UnmarshalJSON(data []byte) error {
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

func (w *WidgetLanguagePreset) MarshalJSON() ([]byte, error) {
	type embed WidgetLanguagePreset
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
