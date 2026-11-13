package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Client -> server message for the multi-context endpoint.
//
// Adds a “context_id“ so several independent dialogue streams can be
// multiplexed over a single socket. “voices“/“voice_settings“/
// “pronunciation_dictionary_locators“ apply to a context's first message
// (not the connection).
type TextToDialogueMultiClientMessage struct {
	Inputs         []*TextToDialogueMultiClientMessageInputsItem `json:"inputs,omitempty" xml:"inputs,omitempty"`
	Flush          *bool                                         `json:"flush,omitempty" xml:"flush,omitempty"`
	CloseSocket    *bool                                         `json:"close_socket,omitempty" xml:"close_socket,omitempty"`
	KeepAlive      *bool                                         `json:"keep_alive,omitempty" xml:"keep_alive,omitempty"`
	XiApiKey       *string                                       `json:"xi_api_key,omitempty" xml:"xi_api_key,omitempty"`
	Authorization  *string                                       `json:"authorization,omitempty" xml:"authorization,omitempty"`
	SingleUseToken *string                                       `json:"single_use_token,omitempty" xml:"single_use_token,omitempty"`
	// List of voice IDs to register for the connection.
	Voices []string `json:"voices,omitempty" xml:"voices,omitempty"`
	// Voice settings for the connection.
	VoiceSettings                   *TextToDialogueMultiClientMessageVoiceSettings                         `json:"voice_settings,omitempty" xml:"voice_settings,omitempty"`
	PronunciationDictionaryLocators []*TextToDialogueMultiClientMessagePronunciationDictionaryLocatorsItem `json:"pronunciation_dictionary_locators,omitempty" xml:"pronunciation_dictionary_locators,omitempty"`
	// Context ID for an independent dialogue stream within the socket.
	ContextId *string `json:"context_id,omitempty" xml:"context_id,omitempty"`
	// Flush and close this context, leaving other contexts open.
	CloseContext    *bool `json:"close_context,omitempty" xml:"close_context,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TextToDialogueMultiClientMessage) GetInputs() (value []*TextToDialogueMultiClientMessageInputsItem) {
	if t == nil {
		return
	}
	return t.Inputs
}

func (t *TextToDialogueMultiClientMessage) GetFlush() (value *bool) {
	if t == nil {
		return
	}
	return t.Flush
}

func (t *TextToDialogueMultiClientMessage) GetCloseSocket() (value *bool) {
	if t == nil {
		return
	}
	return t.CloseSocket
}

func (t *TextToDialogueMultiClientMessage) GetKeepAlive() (value *bool) {
	if t == nil {
		return
	}
	return t.KeepAlive
}

func (t *TextToDialogueMultiClientMessage) GetXiApiKey() (value *string) {
	if t == nil {
		return
	}
	return t.XiApiKey
}

func (t *TextToDialogueMultiClientMessage) GetAuthorization() (value *string) {
	if t == nil {
		return
	}
	return t.Authorization
}

func (t *TextToDialogueMultiClientMessage) GetSingleUseToken() (value *string) {
	if t == nil {
		return
	}
	return t.SingleUseToken
}

func (t *TextToDialogueMultiClientMessage) GetVoices() (value []string) {
	if t == nil {
		return
	}
	return t.Voices
}

func (t *TextToDialogueMultiClientMessage) GetVoiceSettings() (value *TextToDialogueMultiClientMessageVoiceSettings) {
	if t == nil {
		return
	}
	return t.VoiceSettings
}

func (t *TextToDialogueMultiClientMessage) GetPronunciationDictionaryLocators() (value []*TextToDialogueMultiClientMessagePronunciationDictionaryLocatorsItem) {
	if t == nil {
		return
	}
	return t.PronunciationDictionaryLocators
}

func (t *TextToDialogueMultiClientMessage) GetContextId() (value *string) {
	if t == nil {
		return
	}
	return t.ContextId
}

func (t *TextToDialogueMultiClientMessage) GetCloseContext() (value *bool) {
	if t == nil {
		return
	}
	return t.CloseContext
}

func (t *TextToDialogueMultiClientMessage) SetInputs(inputs []*TextToDialogueMultiClientMessageInputsItem) {
	t.Inputs = inputs
}

func (t *TextToDialogueMultiClientMessage) SetFlush(flush *bool) {
	t.Flush = flush
}

func (t *TextToDialogueMultiClientMessage) SetCloseSocket(closeSocket *bool) {
	t.CloseSocket = closeSocket
}

func (t *TextToDialogueMultiClientMessage) SetKeepAlive(keepAlive *bool) {
	t.KeepAlive = keepAlive
}

func (t *TextToDialogueMultiClientMessage) SetXiApiKey(xiApiKey *string) {
	t.XiApiKey = xiApiKey
}

func (t *TextToDialogueMultiClientMessage) SetAuthorization(authorization *string) {
	t.Authorization = authorization
}

func (t *TextToDialogueMultiClientMessage) SetSingleUseToken(singleUseToken *string) {
	t.SingleUseToken = singleUseToken
}

func (t *TextToDialogueMultiClientMessage) SetVoices(voices []string) {
	t.Voices = voices
}

func (t *TextToDialogueMultiClientMessage) SetVoiceSettings(voiceSettings *TextToDialogueMultiClientMessageVoiceSettings) {
	t.VoiceSettings = voiceSettings
}

func (t *TextToDialogueMultiClientMessage) SetPronunciationDictionaryLocators(pronunciationDictionaryLocators []*TextToDialogueMultiClientMessagePronunciationDictionaryLocatorsItem) {
	t.PronunciationDictionaryLocators = pronunciationDictionaryLocators
}

func (t *TextToDialogueMultiClientMessage) SetContextId(contextId *string) {
	t.ContextId = contextId
}

func (t *TextToDialogueMultiClientMessage) SetCloseContext(closeContext *bool) {
	t.CloseContext = closeContext
}

func (t *TextToDialogueMultiClientMessage) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TextToDialogueMultiClientMessage) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TextToDialogueMultiClientMessage to string"
	}
	return string(jsonData)
}

func (t *TextToDialogueMultiClientMessage) UnmarshalJSON(data []byte) error {
	type unmarshaler TextToDialogueMultiClientMessage
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TextToDialogueMultiClientMessage(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TextToDialogueMultiClientMessage) MarshalJSON() ([]byte, error) {
	type embed TextToDialogueMultiClientMessage
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
