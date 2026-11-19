package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationalConfigApiModelWorkflowOverrideInput struct {
	// Configuration for conversational transcription
	Asr *AsrConversationalConfigWorkflowOverride `json:"asr,omitempty" xml:"asr,omitempty"`
	// Configuration for turn detection
	Turn *TurnConfigWorkflowOverride `json:"turn,omitempty" xml:"turn,omitempty"`
	// Configuration for conversational text to speech
	Tts *TtsConversationalConfigWorkflowOverrideInput `json:"tts,omitempty" xml:"tts,omitempty"`
	// Configuration for conversational events
	Conversation *ConversationConfigWorkflowOverrideInput `json:"conversation,omitempty" xml:"conversation,omitempty"`
	// Language presets for conversations
	LanguagePresets map[string]*LanguagePresetInput `json:"language_presets,omitempty"`
	// Configuration for voice activity detection
	Vad *VadConfigWorkflowOverride `json:"vad,omitempty" xml:"vad,omitempty"`
	// Agent specific configuration
	Agent           *AgentConfigApiModelWorkflowOverrideInput `json:"agent,omitempty" xml:"agent,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ConversationalConfigApiModelWorkflowOverrideInput) GetAsr() (value *AsrConversationalConfigWorkflowOverride) {
	if c == nil {
		return
	}
	return c.Asr
}

func (c *ConversationalConfigApiModelWorkflowOverrideInput) GetTurn() (value *TurnConfigWorkflowOverride) {
	if c == nil {
		return
	}
	return c.Turn
}

func (c *ConversationalConfigApiModelWorkflowOverrideInput) GetTts() (value *TtsConversationalConfigWorkflowOverrideInput) {
	if c == nil {
		return
	}
	return c.Tts
}

func (c *ConversationalConfigApiModelWorkflowOverrideInput) GetConversation() (value *ConversationConfigWorkflowOverrideInput) {
	if c == nil {
		return
	}
	return c.Conversation
}

func (c *ConversationalConfigApiModelWorkflowOverrideInput) GetLanguagePresets() (value map[string]*LanguagePresetInput) {
	if c == nil {
		return
	}
	return c.LanguagePresets
}

func (c *ConversationalConfigApiModelWorkflowOverrideInput) GetVad() (value *VadConfigWorkflowOverride) {
	if c == nil {
		return
	}
	return c.Vad
}

func (c *ConversationalConfigApiModelWorkflowOverrideInput) GetAgent() (value *AgentConfigApiModelWorkflowOverrideInput) {
	if c == nil {
		return
	}
	return c.Agent
}

func (c *ConversationalConfigApiModelWorkflowOverrideInput) SetAsr(asr *AsrConversationalConfigWorkflowOverride) {
	c.Asr = asr
}

func (c *ConversationalConfigApiModelWorkflowOverrideInput) SetTurn(turn *TurnConfigWorkflowOverride) {
	c.Turn = turn
}

func (c *ConversationalConfigApiModelWorkflowOverrideInput) SetTts(tts *TtsConversationalConfigWorkflowOverrideInput) {
	c.Tts = tts
}

func (c *ConversationalConfigApiModelWorkflowOverrideInput) SetConversation(conversation *ConversationConfigWorkflowOverrideInput) {
	c.Conversation = conversation
}

func (c *ConversationalConfigApiModelWorkflowOverrideInput) SetLanguagePresets(languagePresets map[string]*LanguagePresetInput) {
	c.LanguagePresets = languagePresets
}

func (c *ConversationalConfigApiModelWorkflowOverrideInput) SetVad(vad *VadConfigWorkflowOverride) {
	c.Vad = vad
}

func (c *ConversationalConfigApiModelWorkflowOverrideInput) SetAgent(agent *AgentConfigApiModelWorkflowOverrideInput) {
	c.Agent = agent
}

func (c *ConversationalConfigApiModelWorkflowOverrideInput) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationalConfigApiModelWorkflowOverrideInput) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationalConfigApiModelWorkflowOverrideInput to string"
	}
	return string(jsonData)
}

func (c *ConversationalConfigApiModelWorkflowOverrideInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationalConfigApiModelWorkflowOverrideInput) MarshalJSON() ([]byte, error) {
	type embed ConversationalConfigApiModelWorkflowOverrideInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
