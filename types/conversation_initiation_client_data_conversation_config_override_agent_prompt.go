package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// The prompt for the agent
type ConversationInitiationClientDataConversationConfigOverrideAgentPrompt struct {
	// The prompt for the agent
	Prompt *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	// The LLM to query with the prompt and the chat history. If using data residency, the LLM must be supported in the data residency environment
	Llm *ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm `json:"llm,omitempty" xml:"llm,omitempty"`
	// Reasoning effort of the model. Only available for some models.
	ReasoningEffort *ConversationInitiationClientDataConversationConfigOverrideAgentPromptReasoningEffort `json:"reasoning_effort,omitempty" xml:"reasoning_effort,omitempty"`
	// A list of IDs of tools used by the agent
	ToolIds []string `json:"tool_ids,omitempty" xml:"tool_ids,omitempty"`
	// Only applies as a conversation override, allowed whenever tool_ids may be overridden. When true, the overridden tool_ids is the full list of tools allowed in the conversation, including tools on workflow nodes. Ignored when set on the agent.
	ToolIdsStrict *bool `json:"tool_ids_strict,omitempty" xml:"tool_ids_strict,omitempty"`
	// A list of Native MCP server ids to be used by the agent
	NativeMcpServerIds []string `json:"native_mcp_server_ids,omitempty" xml:"native_mcp_server_ids,omitempty"`
	// A list of knowledge bases to be used by the agent
	KnowledgeBase   []*ConversationInitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItem `json:"knowledge_base,omitempty" xml:"knowledge_base,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgentPrompt) GetPrompt() (value *string) {
	if c == nil {
		return
	}
	return c.Prompt
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgentPrompt) GetLlm() (value *ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm) {
	if c == nil {
		return
	}
	return c.Llm
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgentPrompt) GetReasoningEffort() (value *ConversationInitiationClientDataConversationConfigOverrideAgentPromptReasoningEffort) {
	if c == nil {
		return
	}
	return c.ReasoningEffort
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgentPrompt) GetToolIds() (value []string) {
	if c == nil {
		return
	}
	return c.ToolIds
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgentPrompt) GetToolIdsStrict() (value *bool) {
	if c == nil {
		return
	}
	return c.ToolIdsStrict
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgentPrompt) GetNativeMcpServerIds() (value []string) {
	if c == nil {
		return
	}
	return c.NativeMcpServerIds
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgentPrompt) GetKnowledgeBase() (value []*ConversationInitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItem) {
	if c == nil {
		return
	}
	return c.KnowledgeBase
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgentPrompt) SetPrompt(prompt *string) {
	c.Prompt = prompt
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgentPrompt) SetLlm(llm *ConversationInitiationClientDataConversationConfigOverrideAgentPromptLlm) {
	c.Llm = llm
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgentPrompt) SetReasoningEffort(reasoningEffort *ConversationInitiationClientDataConversationConfigOverrideAgentPromptReasoningEffort) {
	c.ReasoningEffort = reasoningEffort
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgentPrompt) SetToolIds(toolIds []string) {
	c.ToolIds = toolIds
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgentPrompt) SetToolIdsStrict(toolIdsStrict *bool) {
	c.ToolIdsStrict = toolIdsStrict
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgentPrompt) SetNativeMcpServerIds(nativeMcpServerIds []string) {
	c.NativeMcpServerIds = nativeMcpServerIds
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgentPrompt) SetKnowledgeBase(knowledgeBase []*ConversationInitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItem) {
	c.KnowledgeBase = knowledgeBase
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgentPrompt) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationInitiationClientDataConversationConfigOverrideAgentPrompt) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationInitiationClientDataConversationConfigOverrideAgentPrompt to string"
	}
	return string(jsonData)
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgentPrompt) UnmarshalJSON(data []byte) error {
	type unmarshaler ConversationInitiationClientDataConversationConfigOverrideAgentPrompt
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConversationInitiationClientDataConversationConfigOverrideAgentPrompt(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgentPrompt) MarshalJSON() ([]byte, error) {
	type embed ConversationInitiationClientDataConversationConfigOverrideAgentPrompt
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
