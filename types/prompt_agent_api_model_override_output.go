package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PromptAgentApiModelOverrideOutput struct {
	// The prompt for the agent
	Prompt *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	// The LLM to query with the prompt and the chat history. If using data residency, the LLM must be supported in the data residency environment
	Llm *Llm `json:"llm,omitempty" xml:"llm,omitempty"`
	// Reasoning effort of the model. Only available for some models.
	ReasoningEffort *LlmReasoningEffort `json:"reasoning_effort,omitempty" xml:"reasoning_effort,omitempty"`
	// A list of IDs of tools used by the agent
	ToolIds []string `json:"tool_ids,omitempty" xml:"tool_ids,omitempty"`
	// Only applies as a conversation override, allowed whenever tool_ids may be overridden. When true, the overridden tool_ids is the full list of tools allowed in the conversation, including tools on workflow nodes. Ignored when set on the agent.
	ToolIdsStrict *bool `json:"tool_ids_strict,omitempty" xml:"tool_ids_strict,omitempty"`
	// A list of Native MCP server ids to be used by the agent
	NativeMcpServerIds []string `json:"native_mcp_server_ids,omitempty" xml:"native_mcp_server_ids,omitempty"`
	// A list of knowledge bases to be used by the agent
	KnowledgeBase   []*KnowledgeBaseLocator `json:"knowledge_base,omitempty" xml:"knowledge_base,omitempty"`
	extraProperties map[string]interface{}
}

func (p *PromptAgentApiModelOverrideOutput) GetPrompt() (value *string) {
	if p == nil {
		return
	}
	return p.Prompt
}

func (p *PromptAgentApiModelOverrideOutput) GetLlm() (value *Llm) {
	if p == nil {
		return
	}
	return p.Llm
}

func (p *PromptAgentApiModelOverrideOutput) GetReasoningEffort() (value *LlmReasoningEffort) {
	if p == nil {
		return
	}
	return p.ReasoningEffort
}

func (p *PromptAgentApiModelOverrideOutput) GetToolIds() (value []string) {
	if p == nil {
		return
	}
	return p.ToolIds
}

func (p *PromptAgentApiModelOverrideOutput) GetToolIdsStrict() (value *bool) {
	if p == nil {
		return
	}
	return p.ToolIdsStrict
}

func (p *PromptAgentApiModelOverrideOutput) GetNativeMcpServerIds() (value []string) {
	if p == nil {
		return
	}
	return p.NativeMcpServerIds
}

func (p *PromptAgentApiModelOverrideOutput) GetKnowledgeBase() (value []*KnowledgeBaseLocator) {
	if p == nil {
		return
	}
	return p.KnowledgeBase
}

func (p *PromptAgentApiModelOverrideOutput) SetPrompt(prompt *string) {
	p.Prompt = prompt
}

func (p *PromptAgentApiModelOverrideOutput) SetLlm(llm *Llm) {
	p.Llm = llm
}

func (p *PromptAgentApiModelOverrideOutput) SetReasoningEffort(reasoningEffort *LlmReasoningEffort) {
	p.ReasoningEffort = reasoningEffort
}

func (p *PromptAgentApiModelOverrideOutput) SetToolIds(toolIds []string) {
	p.ToolIds = toolIds
}

func (p *PromptAgentApiModelOverrideOutput) SetToolIdsStrict(toolIdsStrict *bool) {
	p.ToolIdsStrict = toolIdsStrict
}

func (p *PromptAgentApiModelOverrideOutput) SetNativeMcpServerIds(nativeMcpServerIds []string) {
	p.NativeMcpServerIds = nativeMcpServerIds
}

func (p *PromptAgentApiModelOverrideOutput) SetKnowledgeBase(knowledgeBase []*KnowledgeBaseLocator) {
	p.KnowledgeBase = knowledgeBase
}

func (p *PromptAgentApiModelOverrideOutput) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PromptAgentApiModelOverrideOutput) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PromptAgentApiModelOverrideOutput to string"
	}
	return string(jsonData)
}

func (p *PromptAgentApiModelOverrideOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, p); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PromptAgentApiModelOverrideOutput) MarshalJSON() ([]byte, error) {
	type embed PromptAgentApiModelOverrideOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
