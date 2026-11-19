package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AsrConversationalConfigWorkflowOverride struct {
	// The quality of the transcription
	Quality *AsrQuality `json:"quality,omitempty" xml:"quality,omitempty"`
	// The provider of the transcription service
	Provider *AsrProvider `json:"provider,omitempty" xml:"provider,omitempty"`
	// The format of the audio to be transcribed
	UserInputAudioFormat *AsrInputFormat `json:"user_input_audio_format,omitempty" xml:"user_input_audio_format,omitempty"`
	// Keywords to boost prediction probability for
	Keywords        []string `json:"keywords,omitempty" xml:"keywords,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AsrConversationalConfigWorkflowOverride) GetQuality() (value *AsrQuality) {
	if a == nil {
		return
	}
	return a.Quality
}

func (a *AsrConversationalConfigWorkflowOverride) GetProvider() (value *AsrProvider) {
	if a == nil {
		return
	}
	return a.Provider
}

func (a *AsrConversationalConfigWorkflowOverride) GetUserInputAudioFormat() (value *AsrInputFormat) {
	if a == nil {
		return
	}
	return a.UserInputAudioFormat
}

func (a *AsrConversationalConfigWorkflowOverride) GetKeywords() (value []string) {
	if a == nil {
		return
	}
	return a.Keywords
}

func (a *AsrConversationalConfigWorkflowOverride) SetQuality(quality *AsrQuality) {
	a.Quality = quality
}

func (a *AsrConversationalConfigWorkflowOverride) SetProvider(provider *AsrProvider) {
	a.Provider = provider
}

func (a *AsrConversationalConfigWorkflowOverride) SetUserInputAudioFormat(userInputAudioFormat *AsrInputFormat) {
	a.UserInputAudioFormat = userInputAudioFormat
}

func (a *AsrConversationalConfigWorkflowOverride) SetKeywords(keywords []string) {
	a.Keywords = keywords
}

func (a *AsrConversationalConfigWorkflowOverride) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AsrConversationalConfigWorkflowOverride) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AsrConversationalConfigWorkflowOverride to string"
	}
	return string(jsonData)
}

func (a *AsrConversationalConfigWorkflowOverride) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AsrConversationalConfigWorkflowOverride) MarshalJSON() ([]byte, error) {
	type embed AsrConversationalConfigWorkflowOverride
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
