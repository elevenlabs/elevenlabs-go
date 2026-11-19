package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ConversationConfigClientOverrideConfigOutput struct {
	Asr             *AsrConversationalConfigOverrideConfig `json:"asr,omitempty" xml:"asr,omitempty"`
	Turn            *TurnConfigOverrideConfig              `json:"turn,omitempty" xml:"turn,omitempty"`
	Tts             *TtsConversationalConfigOverrideConfig `json:"tts,omitempty" xml:"tts,omitempty"`
	Conversation    *ConversationConfigOverrideConfig      `json:"conversation,omitempty" xml:"conversation,omitempty"`
	Agent           *AgentConfigOverrideConfig             `json:"agent,omitempty" xml:"agent,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ConversationConfigClientOverrideConfigOutput) GetAsr() (value *AsrConversationalConfigOverrideConfig) {
	if c == nil {
		return
	}
	return c.Asr
}

func (c *ConversationConfigClientOverrideConfigOutput) GetTurn() (value *TurnConfigOverrideConfig) {
	if c == nil {
		return
	}
	return c.Turn
}

func (c *ConversationConfigClientOverrideConfigOutput) GetTts() (value *TtsConversationalConfigOverrideConfig) {
	if c == nil {
		return
	}
	return c.Tts
}

func (c *ConversationConfigClientOverrideConfigOutput) GetConversation() (value *ConversationConfigOverrideConfig) {
	if c == nil {
		return
	}
	return c.Conversation
}

func (c *ConversationConfigClientOverrideConfigOutput) GetAgent() (value *AgentConfigOverrideConfig) {
	if c == nil {
		return
	}
	return c.Agent
}

func (c *ConversationConfigClientOverrideConfigOutput) SetAsr(asr *AsrConversationalConfigOverrideConfig) {
	c.Asr = asr
}

func (c *ConversationConfigClientOverrideConfigOutput) SetTurn(turn *TurnConfigOverrideConfig) {
	c.Turn = turn
}

func (c *ConversationConfigClientOverrideConfigOutput) SetTts(tts *TtsConversationalConfigOverrideConfig) {
	c.Tts = tts
}

func (c *ConversationConfigClientOverrideConfigOutput) SetConversation(conversation *ConversationConfigOverrideConfig) {
	c.Conversation = conversation
}

func (c *ConversationConfigClientOverrideConfigOutput) SetAgent(agent *AgentConfigOverrideConfig) {
	c.Agent = agent
}

func (c *ConversationConfigClientOverrideConfigOutput) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationConfigClientOverrideConfigOutput) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationConfigClientOverrideConfigOutput to string"
	}
	return string(jsonData)
}

func (c *ConversationConfigClientOverrideConfigOutput) UnmarshalJSON(data []byte) error {
	type unmarshaler ConversationConfigClientOverrideConfigOutput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConversationConfigClientOverrideConfigOutput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationConfigClientOverrideConfigOutput) MarshalJSON() ([]byte, error) {
	type embed ConversationConfigClientOverrideConfigOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
