package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type AsrConversationalConfig struct {
	Quality              *AsrQuality     `json:"quality,omitempty" xml:"quality,omitempty"`
	Provider             *AsrProvider    `json:"provider,omitempty" xml:"provider,omitempty"`
	UserInputAudioFormat *AsrInputFormat `json:"user_input_audio_format,omitempty" xml:"user_input_audio_format,omitempty"`
	// Keywords to boost prediction probability for
	Keywords        []string `json:"keywords,omitempty" xml:"keywords,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AsrConversationalConfig) GetQuality() (value *AsrQuality) {
	if a == nil {
		return
	}
	return a.Quality
}

func (a *AsrConversationalConfig) GetProvider() (value *AsrProvider) {
	if a == nil {
		return
	}
	return a.Provider
}

func (a *AsrConversationalConfig) GetUserInputAudioFormat() (value *AsrInputFormat) {
	if a == nil {
		return
	}
	return a.UserInputAudioFormat
}

func (a *AsrConversationalConfig) GetKeywords() (value []string) {
	if a == nil {
		return
	}
	return a.Keywords
}

func (a *AsrConversationalConfig) SetQuality(quality *AsrQuality) {
	a.Quality = quality
}

func (a *AsrConversationalConfig) SetProvider(provider *AsrProvider) {
	a.Provider = provider
}

func (a *AsrConversationalConfig) SetUserInputAudioFormat(userInputAudioFormat *AsrInputFormat) {
	a.UserInputAudioFormat = userInputAudioFormat
}

func (a *AsrConversationalConfig) SetKeywords(keywords []string) {
	a.Keywords = keywords
}

func (a *AsrConversationalConfig) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AsrConversationalConfig) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AsrConversationalConfig to string"
	}
	return string(jsonData)
}

func (a *AsrConversationalConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler AsrConversationalConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AsrConversationalConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AsrConversationalConfig) MarshalJSON() ([]byte, error) {
	type embed AsrConversationalConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
