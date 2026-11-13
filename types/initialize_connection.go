package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type InitializeConnection struct {
	// The initial text that must be sent is a blank space.
	text                            string
	VoiceSettings                   *InitializeConnectionVoiceSettings                         `json:"voice_settings,omitempty" xml:"voice_settings,omitempty"`
	GenerationConfig                *InitializeConnectionGenerationConfig                      `json:"generation_config,omitempty" xml:"generation_config,omitempty"`
	PronunciationDictionaryLocators []*InitializeConnectionPronunciationDictionaryLocatorsItem `json:"pronunciation_dictionary_locators,omitempty" xml:"pronunciation_dictionary_locators,omitempty"`
	XiApiKey                        *string                                                    `json:"xi_api_key,omitempty" xml:"xi_api_key,omitempty"`
	Authorization                   *string                                                    `json:"authorization,omitempty" xml:"authorization,omitempty"`
	extraProperties                 map[string]interface{}
}

func (i *InitializeConnection) GetVoiceSettings() (value *InitializeConnectionVoiceSettings) {
	if i == nil {
		return
	}
	return i.VoiceSettings
}

func (i *InitializeConnection) GetGenerationConfig() (value *InitializeConnectionGenerationConfig) {
	if i == nil {
		return
	}
	return i.GenerationConfig
}

func (i *InitializeConnection) GetPronunciationDictionaryLocators() (value []*InitializeConnectionPronunciationDictionaryLocatorsItem) {
	if i == nil {
		return
	}
	return i.PronunciationDictionaryLocators
}

func (i *InitializeConnection) GetXiApiKey() (value *string) {
	if i == nil {
		return
	}
	return i.XiApiKey
}

func (i *InitializeConnection) GetAuthorization() (value *string) {
	if i == nil {
		return
	}
	return i.Authorization
}

func (i *InitializeConnection) SetVoiceSettings(voiceSettings *InitializeConnectionVoiceSettings) {
	i.VoiceSettings = voiceSettings
}

func (i *InitializeConnection) SetGenerationConfig(generationConfig *InitializeConnectionGenerationConfig) {
	i.GenerationConfig = generationConfig
}

func (i *InitializeConnection) SetPronunciationDictionaryLocators(pronunciationDictionaryLocators []*InitializeConnectionPronunciationDictionaryLocatorsItem) {
	i.PronunciationDictionaryLocators = pronunciationDictionaryLocators
}

func (i *InitializeConnection) SetXiApiKey(xiApiKey *string) {
	i.XiApiKey = xiApiKey
}

func (i *InitializeConnection) SetAuthorization(authorization *string) {
	i.Authorization = authorization
}

func (i *InitializeConnection) Text() string {
	return i.text
}

func (i *InitializeConnection) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i InitializeConnection) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InitializeConnection to string"
	}
	return string(jsonData)
}

func (i *InitializeConnection) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, i); err != nil {
		return err
	}
	type alias InitializeConnection
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*i = InitializeConnection(tmp)
	var literals struct {
		Text *string `json:"text"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Text != nil && *literals.Text != " " {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", i, " ", *literals.Text)
	}
	i.text = " "
	extraProperties, err := internal.ExtractExtraProperties(data, *i, "text")
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	return nil
}

func (i *InitializeConnection) MarshalJSON() ([]byte, error) {
	type embed InitializeConnection
	var marshaler = struct {
		embed
		Text string `json:"text"`
	}{
		embed: embed(*i),
		Text:  " ",
	}
	return json.Marshal(marshaler)
}
