package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WidgetLanguagePresetResponse struct {
	FirstMessage *string `json:"first_message,omitempty" xml:"first_message,omitempty"`
	// The text contents for the selected language
	TextContents *WidgetTextContents `json:"text_contents,omitempty" xml:"text_contents,omitempty"`
	// The text to display for terms and conditions in this language
	TermsText *string `json:"terms_text,omitempty" xml:"terms_text,omitempty"`
	// The HTML to display for terms and conditions in this language
	TermsHtml *string `json:"terms_html,omitempty" xml:"terms_html,omitempty"`
	// The key to display for terms and conditions in this language
	TermsKey        *string `json:"terms_key,omitempty" xml:"terms_key,omitempty"`
	extraProperties map[string]interface{}
}

func (w *WidgetLanguagePresetResponse) GetFirstMessage() (value *string) {
	if w == nil {
		return
	}
	return w.FirstMessage
}

func (w *WidgetLanguagePresetResponse) GetTextContents() (value *WidgetTextContents) {
	if w == nil {
		return
	}
	return w.TextContents
}

func (w *WidgetLanguagePresetResponse) GetTermsText() (value *string) {
	if w == nil {
		return
	}
	return w.TermsText
}

func (w *WidgetLanguagePresetResponse) GetTermsHtml() (value *string) {
	if w == nil {
		return
	}
	return w.TermsHtml
}

func (w *WidgetLanguagePresetResponse) GetTermsKey() (value *string) {
	if w == nil {
		return
	}
	return w.TermsKey
}

func (w *WidgetLanguagePresetResponse) SetFirstMessage(firstMessage *string) {
	w.FirstMessage = firstMessage
}

func (w *WidgetLanguagePresetResponse) SetTextContents(textContents *WidgetTextContents) {
	w.TextContents = textContents
}

func (w *WidgetLanguagePresetResponse) SetTermsText(termsText *string) {
	w.TermsText = termsText
}

func (w *WidgetLanguagePresetResponse) SetTermsHtml(termsHtml *string) {
	w.TermsHtml = termsHtml
}

func (w *WidgetLanguagePresetResponse) SetTermsKey(termsKey *string) {
	w.TermsKey = termsKey
}

func (w *WidgetLanguagePresetResponse) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WidgetLanguagePresetResponse) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WidgetLanguagePresetResponse to string"
	}
	return string(jsonData)
}

func (w *WidgetLanguagePresetResponse) UnmarshalJSON(data []byte) error {
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

func (w *WidgetLanguagePresetResponse) MarshalJSON() ([]byte, error) {
	type embed WidgetLanguagePresetResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
