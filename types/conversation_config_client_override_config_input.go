package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ConversationConfigClientOverrideConfigInput struct {
	Asr             *AsrConversationalConfigOverrideConfig `json:"asr,omitempty" xml:"asr,omitempty"`
	Turn            *TurnConfigOverrideConfig              `json:"turn,omitempty" xml:"turn,omitempty"`
	Tts             *TtsConversationalConfigOverrideConfig `json:"tts,omitempty" xml:"tts,omitempty"`
	Conversation    *ConversationConfigOverrideConfig      `json:"conversation,omitempty" xml:"conversation,omitempty"`
	Agent           *AgentConfigOverrideConfig             `json:"agent,omitempty" xml:"agent,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ConversationConfigClientOverrideConfigInput) GetAsr() (value *AsrConversationalConfigOverrideConfig) {
	if c == nil {
		return
	}
	return c.Asr
}

func (c *ConversationConfigClientOverrideConfigInput) GetTurn() (value *TurnConfigOverrideConfig) {
	if c == nil {
		return
	}
	return c.Turn
}

func (c *ConversationConfigClientOverrideConfigInput) GetTts() (value *TtsConversationalConfigOverrideConfig) {
	if c == nil {
		return
	}
	return c.Tts
}

func (c *ConversationConfigClientOverrideConfigInput) GetConversation() (value *ConversationConfigOverrideConfig) {
	if c == nil {
		return
	}
	return c.Conversation
}

func (c *ConversationConfigClientOverrideConfigInput) GetAgent() (value *AgentConfigOverrideConfig) {
	if c == nil {
		return
	}
	return c.Agent
}

func (c *ConversationConfigClientOverrideConfigInput) SetAsr(asr *AsrConversationalConfigOverrideConfig) {
	c.Asr = asr
}

func (c *ConversationConfigClientOverrideConfigInput) SetTurn(turn *TurnConfigOverrideConfig) {
	c.Turn = turn
}

func (c *ConversationConfigClientOverrideConfigInput) SetTts(tts *TtsConversationalConfigOverrideConfig) {
	c.Tts = tts
}

func (c *ConversationConfigClientOverrideConfigInput) SetConversation(conversation *ConversationConfigOverrideConfig) {
	c.Conversation = conversation
}

func (c *ConversationConfigClientOverrideConfigInput) SetAgent(agent *AgentConfigOverrideConfig) {
	c.Agent = agent
}

func (c *ConversationConfigClientOverrideConfigInput) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationConfigClientOverrideConfigInput) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationConfigClientOverrideConfigInput to string"
	}
	return string(jsonData)
}

func (c *ConversationConfigClientOverrideConfigInput) UnmarshalJSON(data []byte) error {
	type unmarshaler ConversationConfigClientOverrideConfigInput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConversationConfigClientOverrideConfigInput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationConfigClientOverrideConfigInput) MarshalJSON() ([]byte, error) {
	type embed ConversationConfigClientOverrideConfigInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
