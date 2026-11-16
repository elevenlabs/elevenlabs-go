package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type TextToDialogueClientMessage struct {
	Inputs         []*TextToDialogueClientMessageInputsItem `json:"inputs,omitempty" xml:"inputs,omitempty"`
	Flush          *bool                                    `json:"flush,omitempty" xml:"flush,omitempty"`
	CloseSocket    *bool                                    `json:"close_socket,omitempty" xml:"close_socket,omitempty"`
	KeepAlive      *bool                                    `json:"keep_alive,omitempty" xml:"keep_alive,omitempty"`
	XiApiKey       *string                                  `json:"xi_api_key,omitempty" xml:"xi_api_key,omitempty"`
	Authorization  *string                                  `json:"authorization,omitempty" xml:"authorization,omitempty"`
	SingleUseToken *string                                  `json:"single_use_token,omitempty" xml:"single_use_token,omitempty"`
	// List of voice IDs to register for the connection.
	Voices []string `json:"voices,omitempty" xml:"voices,omitempty"`
	// Voice settings for the connection.
	VoiceSettings                   *TextToDialogueClientMessageVoiceSettings                         `json:"voice_settings,omitempty" xml:"voice_settings,omitempty"`
	PronunciationDictionaryLocators []*TextToDialogueClientMessagePronunciationDictionaryLocatorsItem `json:"pronunciation_dictionary_locators,omitempty" xml:"pronunciation_dictionary_locators,omitempty"`
	extraProperties                 map[string]interface{}
}

func (t *TextToDialogueClientMessage) GetInputs() (value []*TextToDialogueClientMessageInputsItem) {
	if t == nil {
		return
	}
	return t.Inputs
}

func (t *TextToDialogueClientMessage) GetFlush() (value *bool) {
	if t == nil {
		return
	}
	return t.Flush
}

func (t *TextToDialogueClientMessage) GetCloseSocket() (value *bool) {
	if t == nil {
		return
	}
	return t.CloseSocket
}

func (t *TextToDialogueClientMessage) GetKeepAlive() (value *bool) {
	if t == nil {
		return
	}
	return t.KeepAlive
}

func (t *TextToDialogueClientMessage) GetXiApiKey() (value *string) {
	if t == nil {
		return
	}
	return t.XiApiKey
}

func (t *TextToDialogueClientMessage) GetAuthorization() (value *string) {
	if t == nil {
		return
	}
	return t.Authorization
}

func (t *TextToDialogueClientMessage) GetSingleUseToken() (value *string) {
	if t == nil {
		return
	}
	return t.SingleUseToken
}

func (t *TextToDialogueClientMessage) GetVoices() (value []string) {
	if t == nil {
		return
	}
	return t.Voices
}

func (t *TextToDialogueClientMessage) GetVoiceSettings() (value *TextToDialogueClientMessageVoiceSettings) {
	if t == nil {
		return
	}
	return t.VoiceSettings
}

func (t *TextToDialogueClientMessage) GetPronunciationDictionaryLocators() (value []*TextToDialogueClientMessagePronunciationDictionaryLocatorsItem) {
	if t == nil {
		return
	}
	return t.PronunciationDictionaryLocators
}

func (t *TextToDialogueClientMessage) SetInputs(inputs []*TextToDialogueClientMessageInputsItem) {
	t.Inputs = inputs
}

func (t *TextToDialogueClientMessage) SetFlush(flush *bool) {
	t.Flush = flush
}

func (t *TextToDialogueClientMessage) SetCloseSocket(closeSocket *bool) {
	t.CloseSocket = closeSocket
}

func (t *TextToDialogueClientMessage) SetKeepAlive(keepAlive *bool) {
	t.KeepAlive = keepAlive
}

func (t *TextToDialogueClientMessage) SetXiApiKey(xiApiKey *string) {
	t.XiApiKey = xiApiKey
}

func (t *TextToDialogueClientMessage) SetAuthorization(authorization *string) {
	t.Authorization = authorization
}

func (t *TextToDialogueClientMessage) SetSingleUseToken(singleUseToken *string) {
	t.SingleUseToken = singleUseToken
}

func (t *TextToDialogueClientMessage) SetVoices(voices []string) {
	t.Voices = voices
}

func (t *TextToDialogueClientMessage) SetVoiceSettings(voiceSettings *TextToDialogueClientMessageVoiceSettings) {
	t.VoiceSettings = voiceSettings
}

func (t *TextToDialogueClientMessage) SetPronunciationDictionaryLocators(pronunciationDictionaryLocators []*TextToDialogueClientMessagePronunciationDictionaryLocatorsItem) {
	t.PronunciationDictionaryLocators = pronunciationDictionaryLocators
}

func (t *TextToDialogueClientMessage) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TextToDialogueClientMessage) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TextToDialogueClientMessage to string"
	}
	return string(jsonData)
}

func (t *TextToDialogueClientMessage) UnmarshalJSON(data []byte) error {
	type unmarshaler TextToDialogueClientMessage
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TextToDialogueClientMessage(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TextToDialogueClientMessage) MarshalJSON() ([]byte, error) {
	type embed TextToDialogueClientMessage
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
