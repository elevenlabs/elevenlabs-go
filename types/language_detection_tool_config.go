package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type LanguageDetectionToolConfig struct {
	// If no language switch happens in the first 2 user turns, later attempts fail and the conversation stays in the current language. If the language switches during those turns, later switching stays available. Enable to reduce the possibility of false switching.
	OnlyAtConversationStart *bool `json:"only_at_conversation_start,omitempty" xml:"only_at_conversation_start,omitempty"`
	extraProperties         map[string]interface{}
}

func (l *LanguageDetectionToolConfig) GetOnlyAtConversationStart() (value *bool) {
	if l == nil {
		return
	}
	return l.OnlyAtConversationStart
}

func (l *LanguageDetectionToolConfig) SetOnlyAtConversationStart(onlyAtConversationStart *bool) {
	l.OnlyAtConversationStart = onlyAtConversationStart
}

func (l *LanguageDetectionToolConfig) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LanguageDetectionToolConfig) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LanguageDetectionToolConfig to string"
	}
	return string(jsonData)
}

func (l *LanguageDetectionToolConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler LanguageDetectionToolConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LanguageDetectionToolConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l, "system_tool_type")
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *LanguageDetectionToolConfig) MarshalJSON() ([]byte, error) {
	type embed LanguageDetectionToolConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
