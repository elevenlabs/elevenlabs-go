package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationInitiationMetadataConversationInitiationMetadataEvent struct {
	ConversationId string `json:"conversation_id" xml:"conversation_id" required:"true"`
	// Audio format specification for agent's speech output.
	AgentOutputAudioFormat ConversationInitiationMetadataConversationInitiationMetadataEventAgentOutputAudioFormat `json:"agent_output_audio_format" xml:"agent_output_audio_format" required:"true"`
	// Audio format specification for user's speech input.
	UserInputAudioFormat ConversationInitiationMetadataConversationInitiationMetadataEventUserInputAudioFormat `json:"user_input_audio_format" xml:"user_input_audio_format" required:"true"`
	extraProperties      map[string]interface{}
}

func (c *ConversationInitiationMetadataConversationInitiationMetadataEvent) GetConversationId() (value string) {
	if c == nil {
		return
	}
	return c.ConversationId
}

func (c *ConversationInitiationMetadataConversationInitiationMetadataEvent) GetAgentOutputAudioFormat() (value ConversationInitiationMetadataConversationInitiationMetadataEventAgentOutputAudioFormat) {
	if c == nil {
		return
	}
	return c.AgentOutputAudioFormat
}

func (c *ConversationInitiationMetadataConversationInitiationMetadataEvent) GetUserInputAudioFormat() (value ConversationInitiationMetadataConversationInitiationMetadataEventUserInputAudioFormat) {
	if c == nil {
		return
	}
	return c.UserInputAudioFormat
}

func (c *ConversationInitiationMetadataConversationInitiationMetadataEvent) SetConversationId(conversationId string) {
	c.ConversationId = conversationId
}

func (c *ConversationInitiationMetadataConversationInitiationMetadataEvent) SetAgentOutputAudioFormat(agentOutputAudioFormat ConversationInitiationMetadataConversationInitiationMetadataEventAgentOutputAudioFormat) {
	c.AgentOutputAudioFormat = agentOutputAudioFormat
}

func (c *ConversationInitiationMetadataConversationInitiationMetadataEvent) SetUserInputAudioFormat(userInputAudioFormat ConversationInitiationMetadataConversationInitiationMetadataEventUserInputAudioFormat) {
	c.UserInputAudioFormat = userInputAudioFormat
}

func (c *ConversationInitiationMetadataConversationInitiationMetadataEvent) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationInitiationMetadataConversationInitiationMetadataEvent) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationInitiationMetadataConversationInitiationMetadataEvent to string"
	}
	return string(jsonData)
}

func (c *ConversationInitiationMetadataConversationInitiationMetadataEvent) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ConversationInitiationMetadataConversationInitiationMetadataEvent
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ConversationInitiationMetadataConversationInitiationMetadataEvent(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationInitiationMetadataConversationInitiationMetadataEvent) MarshalJSON() ([]byte, error) {
	type embed ConversationInitiationMetadataConversationInitiationMetadataEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
