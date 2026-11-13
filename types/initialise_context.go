package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type InitialiseContext struct {
	Text                            string                                                  `json:"text" xml:"text" required:"true"`
	VoiceSettings                   *InitialiseContextVoiceSettings                         `json:"voice_settings,omitempty" xml:"voice_settings,omitempty"`
	GenerationConfig                *InitialiseContextGenerationConfig                      `json:"generation_config,omitempty" xml:"generation_config,omitempty"`
	PronunciationDictionaryLocators []*InitialiseContextPronunciationDictionaryLocatorsItem `json:"pronunciation_dictionary_locators,omitempty" xml:"pronunciation_dictionary_locators,omitempty"`
	XiApiKey                        *string                                                 `json:"xi_api_key,omitempty" xml:"xi_api_key,omitempty"`
	Authorization                   *string                                                 `json:"authorization,omitempty" xml:"authorization,omitempty"`
	ContextId                       *string                                                 `json:"context_id,omitempty" xml:"context_id,omitempty"`
	extraProperties                 map[string]interface{}
}

func (i *InitialiseContext) GetText() (value string) {
	if i == nil {
		return
	}
	return i.Text
}

func (i *InitialiseContext) GetVoiceSettings() (value *InitialiseContextVoiceSettings) {
	if i == nil {
		return
	}
	return i.VoiceSettings
}

func (i *InitialiseContext) GetGenerationConfig() (value *InitialiseContextGenerationConfig) {
	if i == nil {
		return
	}
	return i.GenerationConfig
}

func (i *InitialiseContext) GetPronunciationDictionaryLocators() (value []*InitialiseContextPronunciationDictionaryLocatorsItem) {
	if i == nil {
		return
	}
	return i.PronunciationDictionaryLocators
}

func (i *InitialiseContext) GetXiApiKey() (value *string) {
	if i == nil {
		return
	}
	return i.XiApiKey
}

func (i *InitialiseContext) GetAuthorization() (value *string) {
	if i == nil {
		return
	}
	return i.Authorization
}

func (i *InitialiseContext) GetContextId() (value *string) {
	if i == nil {
		return
	}
	return i.ContextId
}

func (i *InitialiseContext) SetText(text string) {
	i.Text = text
}

func (i *InitialiseContext) SetVoiceSettings(voiceSettings *InitialiseContextVoiceSettings) {
	i.VoiceSettings = voiceSettings
}

func (i *InitialiseContext) SetGenerationConfig(generationConfig *InitialiseContextGenerationConfig) {
	i.GenerationConfig = generationConfig
}

func (i *InitialiseContext) SetPronunciationDictionaryLocators(pronunciationDictionaryLocators []*InitialiseContextPronunciationDictionaryLocatorsItem) {
	i.PronunciationDictionaryLocators = pronunciationDictionaryLocators
}

func (i *InitialiseContext) SetXiApiKey(xiApiKey *string) {
	i.XiApiKey = xiApiKey
}

func (i *InitialiseContext) SetAuthorization(authorization *string) {
	i.Authorization = authorization
}

func (i *InitialiseContext) SetContextId(contextId *string) {
	i.ContextId = contextId
}

func (i *InitialiseContext) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i InitialiseContext) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InitialiseContext to string"
	}
	return string(jsonData)
}

func (i *InitialiseContext) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, i); err != nil {
		return err
	}
	type alias InitialiseContext
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*i = InitialiseContext(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	return nil
}

func (i *InitialiseContext) MarshalJSON() ([]byte, error) {
	type embed InitialiseContext
	var marshaler = struct {
		embed
	}{
		embed: embed(*i),
	}
	return json.Marshal(marshaler)
}
