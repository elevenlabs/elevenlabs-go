package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationConfigClientOverrideInput struct {
	// Configuration for conversational transcription
	Asr *AsrConversationalConfigOverride `json:"asr,omitempty" xml:"asr,omitempty"`
	// Configuration for turn detection
	Turn *TurnConfigOverride `json:"turn,omitempty" xml:"turn,omitempty"`
	// Configuration for conversational text to speech
	Tts *TtsConversationalConfigOverride `json:"tts,omitempty" xml:"tts,omitempty"`
	// Configuration for conversational events
	Conversation *ConversationConfigOverride `json:"conversation,omitempty" xml:"conversation,omitempty"`
	// Agent specific configuration
	Agent           *AgentConfigOverrideInput `json:"agent,omitempty" xml:"agent,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ConversationConfigClientOverrideInput) GetAsr() (value *AsrConversationalConfigOverride) {
	if c == nil {
		return
	}
	return c.Asr
}

func (c *ConversationConfigClientOverrideInput) GetTurn() (value *TurnConfigOverride) {
	if c == nil {
		return
	}
	return c.Turn
}

func (c *ConversationConfigClientOverrideInput) GetTts() (value *TtsConversationalConfigOverride) {
	if c == nil {
		return
	}
	return c.Tts
}

func (c *ConversationConfigClientOverrideInput) GetConversation() (value *ConversationConfigOverride) {
	if c == nil {
		return
	}
	return c.Conversation
}

func (c *ConversationConfigClientOverrideInput) GetAgent() (value *AgentConfigOverrideInput) {
	if c == nil {
		return
	}
	return c.Agent
}

func (c *ConversationConfigClientOverrideInput) SetAsr(asr *AsrConversationalConfigOverride) {
	c.Asr = asr
}

func (c *ConversationConfigClientOverrideInput) SetTurn(turn *TurnConfigOverride) {
	c.Turn = turn
}

func (c *ConversationConfigClientOverrideInput) SetTts(tts *TtsConversationalConfigOverride) {
	c.Tts = tts
}

func (c *ConversationConfigClientOverrideInput) SetConversation(conversation *ConversationConfigOverride) {
	c.Conversation = conversation
}

func (c *ConversationConfigClientOverrideInput) SetAgent(agent *AgentConfigOverrideInput) {
	c.Agent = agent
}

func (c *ConversationConfigClientOverrideInput) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationConfigClientOverrideInput) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationConfigClientOverrideInput to string"
	}
	return string(jsonData)
}

func (c *ConversationConfigClientOverrideInput) UnmarshalJSON(data []byte) error {
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

func (c *ConversationConfigClientOverrideInput) MarshalJSON() ([]byte, error) {
	type embed ConversationConfigClientOverrideInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
