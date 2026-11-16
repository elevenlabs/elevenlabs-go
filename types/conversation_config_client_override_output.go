package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationConfigClientOverrideOutput struct {
	// Configuration for conversational transcription
	Asr *AsrConversationalConfigOverride `json:"asr,omitempty" xml:"asr,omitempty"`
	// Configuration for turn detection
	Turn *TurnConfigOverride `json:"turn,omitempty" xml:"turn,omitempty"`
	// Configuration for conversational text to speech
	Tts *TtsConversationalConfigOverride `json:"tts,omitempty" xml:"tts,omitempty"`
	// Configuration for conversational events
	Conversation *ConversationConfigOverride `json:"conversation,omitempty" xml:"conversation,omitempty"`
	// Agent specific configuration
	Agent           *AgentConfigOverrideOutput `json:"agent,omitempty" xml:"agent,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ConversationConfigClientOverrideOutput) GetAsr() (value *AsrConversationalConfigOverride) {
	if c == nil {
		return
	}
	return c.Asr
}

func (c *ConversationConfigClientOverrideOutput) GetTurn() (value *TurnConfigOverride) {
	if c == nil {
		return
	}
	return c.Turn
}

func (c *ConversationConfigClientOverrideOutput) GetTts() (value *TtsConversationalConfigOverride) {
	if c == nil {
		return
	}
	return c.Tts
}

func (c *ConversationConfigClientOverrideOutput) GetConversation() (value *ConversationConfigOverride) {
	if c == nil {
		return
	}
	return c.Conversation
}

func (c *ConversationConfigClientOverrideOutput) GetAgent() (value *AgentConfigOverrideOutput) {
	if c == nil {
		return
	}
	return c.Agent
}

func (c *ConversationConfigClientOverrideOutput) SetAsr(asr *AsrConversationalConfigOverride) {
	c.Asr = asr
}

func (c *ConversationConfigClientOverrideOutput) SetTurn(turn *TurnConfigOverride) {
	c.Turn = turn
}

func (c *ConversationConfigClientOverrideOutput) SetTts(tts *TtsConversationalConfigOverride) {
	c.Tts = tts
}

func (c *ConversationConfigClientOverrideOutput) SetConversation(conversation *ConversationConfigOverride) {
	c.Conversation = conversation
}

func (c *ConversationConfigClientOverrideOutput) SetAgent(agent *AgentConfigOverrideOutput) {
	c.Agent = agent
}

func (c *ConversationConfigClientOverrideOutput) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationConfigClientOverrideOutput) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationConfigClientOverrideOutput to string"
	}
	return string(jsonData)
}

func (c *ConversationConfigClientOverrideOutput) UnmarshalJSON(data []byte) error {
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

func (c *ConversationConfigClientOverrideOutput) MarshalJSON() ([]byte, error) {
	type embed ConversationConfigClientOverrideOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
