package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationalConfigApiModelWorkflowOverrideOutput struct {
	// Configuration for conversational transcription
	Asr *AsrConversationalConfigWorkflowOverride `json:"asr,omitempty" xml:"asr,omitempty"`
	// Configuration for turn detection
	Turn *TurnConfigWorkflowOverride `json:"turn,omitempty" xml:"turn,omitempty"`
	// Configuration for conversational text to speech
	Tts *TtsConversationalConfigWorkflowOverrideOutput `json:"tts,omitempty" xml:"tts,omitempty"`
	// Configuration for conversational events
	Conversation *ConversationConfigWorkflowOverrideOutput `json:"conversation,omitempty" xml:"conversation,omitempty"`
	// Language presets for conversations
	LanguagePresets map[string]*LanguagePresetOutput `json:"language_presets,omitempty"`
	// Configuration for voice activity detection
	Vad *VadConfigWorkflowOverride `json:"vad,omitempty" xml:"vad,omitempty"`
	// Agent specific configuration
	Agent           *AgentConfigApiModelWorkflowOverrideOutput `json:"agent,omitempty" xml:"agent,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ConversationalConfigApiModelWorkflowOverrideOutput) GetAsr() (value *AsrConversationalConfigWorkflowOverride) {
	if c == nil {
		return
	}
	return c.Asr
}

func (c *ConversationalConfigApiModelWorkflowOverrideOutput) GetTurn() (value *TurnConfigWorkflowOverride) {
	if c == nil {
		return
	}
	return c.Turn
}

func (c *ConversationalConfigApiModelWorkflowOverrideOutput) GetTts() (value *TtsConversationalConfigWorkflowOverrideOutput) {
	if c == nil {
		return
	}
	return c.Tts
}

func (c *ConversationalConfigApiModelWorkflowOverrideOutput) GetConversation() (value *ConversationConfigWorkflowOverrideOutput) {
	if c == nil {
		return
	}
	return c.Conversation
}

func (c *ConversationalConfigApiModelWorkflowOverrideOutput) GetLanguagePresets() (value map[string]*LanguagePresetOutput) {
	if c == nil {
		return
	}
	return c.LanguagePresets
}

func (c *ConversationalConfigApiModelWorkflowOverrideOutput) GetVad() (value *VadConfigWorkflowOverride) {
	if c == nil {
		return
	}
	return c.Vad
}

func (c *ConversationalConfigApiModelWorkflowOverrideOutput) GetAgent() (value *AgentConfigApiModelWorkflowOverrideOutput) {
	if c == nil {
		return
	}
	return c.Agent
}

func (c *ConversationalConfigApiModelWorkflowOverrideOutput) SetAsr(asr *AsrConversationalConfigWorkflowOverride) {
	c.Asr = asr
}

func (c *ConversationalConfigApiModelWorkflowOverrideOutput) SetTurn(turn *TurnConfigWorkflowOverride) {
	c.Turn = turn
}

func (c *ConversationalConfigApiModelWorkflowOverrideOutput) SetTts(tts *TtsConversationalConfigWorkflowOverrideOutput) {
	c.Tts = tts
}

func (c *ConversationalConfigApiModelWorkflowOverrideOutput) SetConversation(conversation *ConversationConfigWorkflowOverrideOutput) {
	c.Conversation = conversation
}

func (c *ConversationalConfigApiModelWorkflowOverrideOutput) SetLanguagePresets(languagePresets map[string]*LanguagePresetOutput) {
	c.LanguagePresets = languagePresets
}

func (c *ConversationalConfigApiModelWorkflowOverrideOutput) SetVad(vad *VadConfigWorkflowOverride) {
	c.Vad = vad
}

func (c *ConversationalConfigApiModelWorkflowOverrideOutput) SetAgent(agent *AgentConfigApiModelWorkflowOverrideOutput) {
	c.Agent = agent
}

func (c *ConversationalConfigApiModelWorkflowOverrideOutput) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationalConfigApiModelWorkflowOverrideOutput) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationalConfigApiModelWorkflowOverrideOutput to string"
	}
	return string(jsonData)
}

func (c *ConversationalConfigApiModelWorkflowOverrideOutput) UnmarshalJSON(data []byte) error {
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

func (c *ConversationalConfigApiModelWorkflowOverrideOutput) MarshalJSON() ([]byte, error) {
	type embed ConversationalConfigApiModelWorkflowOverrideOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
