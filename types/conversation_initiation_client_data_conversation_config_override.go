package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ConversationInitiationClientDataConversationConfigOverride struct {
	// Configuration for conversational transcription
	Asr *ConversationInitiationClientDataConversationConfigOverrideAsr `json:"asr,omitempty" xml:"asr,omitempty"`
	// Configuration for turn detection
	Turn *ConversationInitiationClientDataConversationConfigOverrideTurn `json:"turn,omitempty" xml:"turn,omitempty"`
	// Configuration for conversational text to speech
	Tts *ConversationInitiationClientDataConversationConfigOverrideTts `json:"tts,omitempty" xml:"tts,omitempty"`
	// Configuration for conversational events
	Conversation *ConversationInitiationClientDataConversationConfigOverrideConversation `json:"conversation,omitempty" xml:"conversation,omitempty"`
	// Agent specific configuration
	Agent           *ConversationInitiationClientDataConversationConfigOverrideAgent `json:"agent,omitempty" xml:"agent,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ConversationInitiationClientDataConversationConfigOverride) GetAsr() (value *ConversationInitiationClientDataConversationConfigOverrideAsr) {
	if c == nil {
		return
	}
	return c.Asr
}

func (c *ConversationInitiationClientDataConversationConfigOverride) GetTurn() (value *ConversationInitiationClientDataConversationConfigOverrideTurn) {
	if c == nil {
		return
	}
	return c.Turn
}

func (c *ConversationInitiationClientDataConversationConfigOverride) GetTts() (value *ConversationInitiationClientDataConversationConfigOverrideTts) {
	if c == nil {
		return
	}
	return c.Tts
}

func (c *ConversationInitiationClientDataConversationConfigOverride) GetConversation() (value *ConversationInitiationClientDataConversationConfigOverrideConversation) {
	if c == nil {
		return
	}
	return c.Conversation
}

func (c *ConversationInitiationClientDataConversationConfigOverride) GetAgent() (value *ConversationInitiationClientDataConversationConfigOverrideAgent) {
	if c == nil {
		return
	}
	return c.Agent
}

func (c *ConversationInitiationClientDataConversationConfigOverride) SetAsr(asr *ConversationInitiationClientDataConversationConfigOverrideAsr) {
	c.Asr = asr
}

func (c *ConversationInitiationClientDataConversationConfigOverride) SetTurn(turn *ConversationInitiationClientDataConversationConfigOverrideTurn) {
	c.Turn = turn
}

func (c *ConversationInitiationClientDataConversationConfigOverride) SetTts(tts *ConversationInitiationClientDataConversationConfigOverrideTts) {
	c.Tts = tts
}

func (c *ConversationInitiationClientDataConversationConfigOverride) SetConversation(conversation *ConversationInitiationClientDataConversationConfigOverrideConversation) {
	c.Conversation = conversation
}

func (c *ConversationInitiationClientDataConversationConfigOverride) SetAgent(agent *ConversationInitiationClientDataConversationConfigOverrideAgent) {
	c.Agent = agent
}

func (c *ConversationInitiationClientDataConversationConfigOverride) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationInitiationClientDataConversationConfigOverride) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationInitiationClientDataConversationConfigOverride to string"
	}
	return string(jsonData)
}

func (c *ConversationInitiationClientDataConversationConfigOverride) UnmarshalJSON(data []byte) error {
	type unmarshaler ConversationInitiationClientDataConversationConfigOverride
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConversationInitiationClientDataConversationConfigOverride(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationInitiationClientDataConversationConfigOverride) MarshalJSON() ([]byte, error) {
	type embed ConversationInitiationClientDataConversationConfigOverride
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
