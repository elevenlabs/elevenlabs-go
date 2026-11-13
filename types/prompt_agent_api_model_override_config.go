package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type PromptAgentApiModelOverrideConfig struct {
	// Whether to allow overriding the prompt field.
	Prompt *bool `json:"prompt,omitempty" xml:"prompt,omitempty"`
	// Whether to allow overriding the llm field.
	Llm *bool `json:"llm,omitempty" xml:"llm,omitempty"`
	// Whether to allow overriding the reasoning_effort field.
	ReasoningEffort *bool `json:"reasoning_effort,omitempty" xml:"reasoning_effort,omitempty"`
	// Whether to allow overriding the tool_ids field.
	ToolIds *bool `json:"tool_ids,omitempty" xml:"tool_ids,omitempty"`
	// Whether to allow overriding the tool_ids_strict field.
	ToolIdsStrict *bool `json:"tool_ids_strict,omitempty" xml:"tool_ids_strict,omitempty"`
	// Whether to allow overriding the native_mcp_server_ids field.
	NativeMcpServerIds *bool `json:"native_mcp_server_ids,omitempty" xml:"native_mcp_server_ids,omitempty"`
	// Whether to allow overriding the knowledge_base field.
	KnowledgeBase   *bool `json:"knowledge_base,omitempty" xml:"knowledge_base,omitempty"`
	extraProperties map[string]interface{}
}

func (p *PromptAgentApiModelOverrideConfig) GetPrompt() (value *bool) {
	if p == nil {
		return
	}
	return p.Prompt
}

func (p *PromptAgentApiModelOverrideConfig) GetLlm() (value *bool) {
	if p == nil {
		return
	}
	return p.Llm
}

func (p *PromptAgentApiModelOverrideConfig) GetReasoningEffort() (value *bool) {
	if p == nil {
		return
	}
	return p.ReasoningEffort
}

func (p *PromptAgentApiModelOverrideConfig) GetToolIds() (value *bool) {
	if p == nil {
		return
	}
	return p.ToolIds
}

func (p *PromptAgentApiModelOverrideConfig) GetToolIdsStrict() (value *bool) {
	if p == nil {
		return
	}
	return p.ToolIdsStrict
}

func (p *PromptAgentApiModelOverrideConfig) GetNativeMcpServerIds() (value *bool) {
	if p == nil {
		return
	}
	return p.NativeMcpServerIds
}

func (p *PromptAgentApiModelOverrideConfig) GetKnowledgeBase() (value *bool) {
	if p == nil {
		return
	}
	return p.KnowledgeBase
}

func (p *PromptAgentApiModelOverrideConfig) SetPrompt(prompt *bool) {
	p.Prompt = prompt
}

func (p *PromptAgentApiModelOverrideConfig) SetLlm(llm *bool) {
	p.Llm = llm
}

func (p *PromptAgentApiModelOverrideConfig) SetReasoningEffort(reasoningEffort *bool) {
	p.ReasoningEffort = reasoningEffort
}

func (p *PromptAgentApiModelOverrideConfig) SetToolIds(toolIds *bool) {
	p.ToolIds = toolIds
}

func (p *PromptAgentApiModelOverrideConfig) SetToolIdsStrict(toolIdsStrict *bool) {
	p.ToolIdsStrict = toolIdsStrict
}

func (p *PromptAgentApiModelOverrideConfig) SetNativeMcpServerIds(nativeMcpServerIds *bool) {
	p.NativeMcpServerIds = nativeMcpServerIds
}

func (p *PromptAgentApiModelOverrideConfig) SetKnowledgeBase(knowledgeBase *bool) {
	p.KnowledgeBase = knowledgeBase
}

func (p *PromptAgentApiModelOverrideConfig) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PromptAgentApiModelOverrideConfig) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PromptAgentApiModelOverrideConfig to string"
	}
	return string(jsonData)
}

func (p *PromptAgentApiModelOverrideConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler PromptAgentApiModelOverrideConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PromptAgentApiModelOverrideConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PromptAgentApiModelOverrideConfig) MarshalJSON() ([]byte, error) {
	type embed PromptAgentApiModelOverrideConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
