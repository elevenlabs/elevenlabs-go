package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type InitializeConnectionMulti struct {
	// The initial text that must be sent is a blank space.
	text                            string
	VoiceSettings                   *InitializeConnectionMultiVoiceSettings                         `json:"voice_settings,omitempty" xml:"voice_settings,omitempty"`
	GenerationConfig                *InitializeConnectionMultiGenerationConfig                      `json:"generation_config,omitempty" xml:"generation_config,omitempty"`
	PronunciationDictionaryLocators []*InitializeConnectionMultiPronunciationDictionaryLocatorsItem `json:"pronunciation_dictionary_locators,omitempty" xml:"pronunciation_dictionary_locators,omitempty"`
	XiApiKey                        *string                                                         `json:"xi_api_key,omitempty" xml:"xi_api_key,omitempty"`
	Authorization                   *string                                                         `json:"authorization,omitempty" xml:"authorization,omitempty"`
	ContextId                       *string                                                         `json:"context_id,omitempty" xml:"context_id,omitempty"`
	extraProperties                 map[string]interface{}
}

func (i *InitializeConnectionMulti) GetVoiceSettings() (value *InitializeConnectionMultiVoiceSettings) {
	if i == nil {
		return
	}
	return i.VoiceSettings
}

func (i *InitializeConnectionMulti) GetGenerationConfig() (value *InitializeConnectionMultiGenerationConfig) {
	if i == nil {
		return
	}
	return i.GenerationConfig
}

func (i *InitializeConnectionMulti) GetPronunciationDictionaryLocators() (value []*InitializeConnectionMultiPronunciationDictionaryLocatorsItem) {
	if i == nil {
		return
	}
	return i.PronunciationDictionaryLocators
}

func (i *InitializeConnectionMulti) GetXiApiKey() (value *string) {
	if i == nil {
		return
	}
	return i.XiApiKey
}

func (i *InitializeConnectionMulti) GetAuthorization() (value *string) {
	if i == nil {
		return
	}
	return i.Authorization
}

func (i *InitializeConnectionMulti) GetContextId() (value *string) {
	if i == nil {
		return
	}
	return i.ContextId
}

func (i *InitializeConnectionMulti) SetVoiceSettings(voiceSettings *InitializeConnectionMultiVoiceSettings) {
	i.VoiceSettings = voiceSettings
}

func (i *InitializeConnectionMulti) SetGenerationConfig(generationConfig *InitializeConnectionMultiGenerationConfig) {
	i.GenerationConfig = generationConfig
}

func (i *InitializeConnectionMulti) SetPronunciationDictionaryLocators(pronunciationDictionaryLocators []*InitializeConnectionMultiPronunciationDictionaryLocatorsItem) {
	i.PronunciationDictionaryLocators = pronunciationDictionaryLocators
}

func (i *InitializeConnectionMulti) SetXiApiKey(xiApiKey *string) {
	i.XiApiKey = xiApiKey
}

func (i *InitializeConnectionMulti) SetAuthorization(authorization *string) {
	i.Authorization = authorization
}

func (i *InitializeConnectionMulti) SetContextId(contextId *string) {
	i.ContextId = contextId
}

func (i *InitializeConnectionMulti) Text() string {
	return i.text
}

func (i *InitializeConnectionMulti) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i InitializeConnectionMulti) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InitializeConnectionMulti to string"
	}
	return string(jsonData)
}

func (i *InitializeConnectionMulti) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, i); err != nil {
		return err
	}
	type alias InitializeConnectionMulti
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*i = InitializeConnectionMulti(tmp)
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

func (i *InitializeConnectionMulti) MarshalJSON() ([]byte, error) {
	type embed InitializeConnectionMulti
	var marshaler = struct {
		embed
		Text string `json:"text"`
	}{
		embed: embed(*i),
		Text:  " ",
	}
	return json.Marshal(marshaler)
}
