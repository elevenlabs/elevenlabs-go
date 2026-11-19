package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type LanguageDetectionToolResultModel struct {
	Status          *string `json:"status,omitempty" xml:"status,omitempty"`
	Reason          *string `json:"reason,omitempty" xml:"reason,omitempty"`
	Language        *string `json:"language,omitempty" xml:"language,omitempty"`
	extraProperties map[string]interface{}
}

func (l *LanguageDetectionToolResultModel) GetStatus() (value *string) {
	if l == nil {
		return
	}
	return l.Status
}

func (l *LanguageDetectionToolResultModel) GetReason() (value *string) {
	if l == nil {
		return
	}
	return l.Reason
}

func (l *LanguageDetectionToolResultModel) GetLanguage() (value *string) {
	if l == nil {
		return
	}
	return l.Language
}

func (l *LanguageDetectionToolResultModel) SetStatus(status *string) {
	l.Status = status
}

func (l *LanguageDetectionToolResultModel) SetReason(reason *string) {
	l.Reason = reason
}

func (l *LanguageDetectionToolResultModel) SetLanguage(language *string) {
	l.Language = language
}

func (l *LanguageDetectionToolResultModel) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LanguageDetectionToolResultModel) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LanguageDetectionToolResultModel to string"
	}
	return string(jsonData)
}

func (l *LanguageDetectionToolResultModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, l); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *l, "result_type")
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *LanguageDetectionToolResultModel) MarshalJSON() ([]byte, error) {
	type embed LanguageDetectionToolResultModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
