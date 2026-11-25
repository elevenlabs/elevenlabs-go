package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationInitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItem struct {
	// The type of the knowledge base
	Type ConversationInitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemType `json:"type" xml:"type" required:"true"`
	// The name of the knowledge base
	Name string `json:"name" xml:"name" required:"true"`
	// The ID of the knowledge base
	Id string `json:"id" xml:"id" required:"true"`
	// The usage mode of the knowledge base
	UsageMode       *ConversationInitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemUsageMode `json:"usage_mode,omitempty" xml:"usage_mode,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItem) GetType() (value ConversationInitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemType) {
	if c == nil {
		return
	}
	return c.Type
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItem) GetName() (value string) {
	if c == nil {
		return
	}
	return c.Name
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItem) GetId() (value string) {
	if c == nil {
		return
	}
	return c.Id
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItem) GetUsageMode() (value *ConversationInitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemUsageMode) {
	if c == nil {
		return
	}
	return c.UsageMode
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItem) SetType(type_ ConversationInitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemType) {
	c.Type = type_
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItem) SetName(name string) {
	c.Name = name
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItem) SetId(id string) {
	c.Id = id
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItem) SetUsageMode(usageMode *ConversationInitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemUsageMode) {
	c.UsageMode = usageMode
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItem) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationInitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItem) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationInitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItem to string"
	}
	return string(jsonData)
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItem) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ConversationInitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItem
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ConversationInitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItem(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItem) MarshalJSON() ([]byte, error) {
	type embed ConversationInitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
