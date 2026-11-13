package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ConversationalConfig struct {
	Asr          *AsrConversationalConfig      `json:"asr,omitempty" xml:"asr,omitempty"`
	Turn         *TurnConfig                   `json:"turn,omitempty" xml:"turn,omitempty"`
	Tts          *TtsConversationalConfigInput `json:"tts,omitempty" xml:"tts,omitempty"`
	Conversation *ConversationConfigInput      `json:"conversation,omitempty" xml:"conversation,omitempty"`
	// Language presets for conversations
	LanguagePresets map[string]*LanguagePresetInput `json:"language_presets,omitempty"`
	Vad             *VadConfig                      `json:"vad,omitempty" xml:"vad,omitempty"`
	Agent           *AgentConfig                    `json:"agent,omitempty" xml:"agent,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ConversationalConfig) GetAsr() (value *AsrConversationalConfig) {
	if c == nil {
		return
	}
	return c.Asr
}

func (c *ConversationalConfig) GetTurn() (value *TurnConfig) {
	if c == nil {
		return
	}
	return c.Turn
}

func (c *ConversationalConfig) GetTts() (value *TtsConversationalConfigInput) {
	if c == nil {
		return
	}
	return c.Tts
}

func (c *ConversationalConfig) GetConversation() (value *ConversationConfigInput) {
	if c == nil {
		return
	}
	return c.Conversation
}

func (c *ConversationalConfig) GetLanguagePresets() (value map[string]*LanguagePresetInput) {
	if c == nil {
		return
	}
	return c.LanguagePresets
}

func (c *ConversationalConfig) GetVad() (value *VadConfig) {
	if c == nil {
		return
	}
	return c.Vad
}

func (c *ConversationalConfig) GetAgent() (value *AgentConfig) {
	if c == nil {
		return
	}
	return c.Agent
}

func (c *ConversationalConfig) SetAsr(asr *AsrConversationalConfig) {
	c.Asr = asr
}

func (c *ConversationalConfig) SetTurn(turn *TurnConfig) {
	c.Turn = turn
}

func (c *ConversationalConfig) SetTts(tts *TtsConversationalConfigInput) {
	c.Tts = tts
}

func (c *ConversationalConfig) SetConversation(conversation *ConversationConfigInput) {
	c.Conversation = conversation
}

func (c *ConversationalConfig) SetLanguagePresets(languagePresets map[string]*LanguagePresetInput) {
	c.LanguagePresets = languagePresets
}

func (c *ConversationalConfig) SetVad(vad *VadConfig) {
	c.Vad = vad
}

func (c *ConversationalConfig) SetAgent(agent *AgentConfig) {
	c.Agent = agent
}

func (c *ConversationalConfig) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationalConfig) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationalConfig to string"
	}
	return string(jsonData)
}

func (c *ConversationalConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler ConversationalConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConversationalConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationalConfig) MarshalJSON() ([]byte, error) {
	type embed ConversationalConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
