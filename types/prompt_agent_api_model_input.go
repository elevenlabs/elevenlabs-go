package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PromptAgentApiModelInput struct {
	// The prompt for the agent
	Prompt *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	Llm    *Llm    `json:"llm,omitempty" xml:"llm,omitempty"`
	// Reasoning effort of the model. Only available for some models.
	ReasoningEffort *LlmReasoningEffort `json:"reasoning_effort,omitempty" xml:"reasoning_effort,omitempty"`
	// Request-time processing tier for a model.
	//
	// STANDARD: default price and throughput.
	// FAST: pays a per-token price premium for higher throughput / lower latency.
	// See OpenAI's fast mode: https://openai.com/api-fast-mode/ and, for Gemini, Vertex
	// Priority PayGo: https://docs.cloud.google.com/gemini-enterprise-agent-platform/models/priority-paygo
	ServiceTier *LlmServiceTier `json:"service_tier,omitempty" xml:"service_tier,omitempty"`
	// Max number of tokens used for thinking. Use 0 to turn off if supported by the model.
	ThinkingBudget *int `json:"thinking_budget,omitempty" xml:"thinking_budget,omitempty"`
	// Enable model reasoning summaries. When disabled, we do not request summaries from provider if possible for faster TTFB. Not ZRM compatible.
	EnableReasoningSummary *bool `json:"enable_reasoning_summary,omitempty" xml:"enable_reasoning_summary,omitempty"`
	// The temperature for the LLM. Defaults to 0. Set to null to omit the parameter from the LLM request entirely (useful for custom LLMs that reject the temperature field).
	Temperature *float64 `json:"temperature,omitempty" xml:"temperature,omitempty"`
	// If greater than 0, maximum number of tokens the LLM can predict
	MaxTokens *int `json:"max_tokens,omitempty" xml:"max_tokens,omitempty"`
	// A list of IDs of tools used by the agent
	ToolIds []string `json:"tool_ids,omitempty" xml:"tool_ids,omitempty"`
	// Only applies as a conversation override, allowed whenever tool_ids may be overridden. When true, the overridden tool_ids is the full list of tools allowed in the conversation, including tools on workflow nodes. Ignored when set on the agent.
	ToolIdsStrict *bool `json:"tool_ids_strict,omitempty" xml:"tool_ids_strict,omitempty"`
	// System tools a conversational agent can be given.
	//
	// Deliberately not named ConversationalBuiltInTools: BuiltInToolsInput and
	// BuiltInToolsOutput are part of the public API spec and the generated SDKs.
	BuiltInTools *BuiltInToolsInput `json:"built_in_tools,omitempty" xml:"built_in_tools,omitempty"`
	// Enable parallel tool calling. When enabled, the agent can execute multiple tools in parallel within a single turn. Not supported by all models.
	EnableParallelToolCalls *bool `json:"enable_parallel_tool_calls,omitempty" xml:"enable_parallel_tool_calls,omitempty"`
	// A list of MCP server ids to be used by the agent
	McpServerIds []string `json:"mcp_server_ids,omitempty" xml:"mcp_server_ids,omitempty" maxItems:"10"`
	// A list of Native MCP server ids to be used by the agent
	NativeMcpServerIds []string `json:"native_mcp_server_ids,omitempty" xml:"native_mcp_server_ids,omitempty" maxItems:"10"`
	// A list of knowledge bases to be used by the agent
	KnowledgeBase []*KnowledgeBaseLocator `json:"knowledge_base,omitempty" xml:"knowledge_base,omitempty"`
	// Definition for a custom LLM if LLM field is set to 'CUSTOM_LLM'
	CustomLlm *CustomLlm `json:"custom_llm,omitempty" xml:"custom_llm,omitempty"`
	// Whether to remove the default personality lines from the system prompt
	IgnoreDefaultPersonality *bool           `json:"ignore_default_personality,omitempty" xml:"ignore_default_personality,omitempty"`
	Rag                      *RagConfigInput `json:"rag,omitempty" xml:"rag,omitempty"`
	// Timezone for displaying current time in system prompt. If set, the current time will be included in the system prompt using this timezone. Must be a valid timezone name (e.g., 'America/New_York', 'Europe/London', 'UTC'). Recommended for accurate time-aware responses; without this, the agent has no knowledge of the current date/time unless you provide it via dynamic variables or tools, which can lead to incorrect or hallucinated time references.
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// Configuration for backup LLM cascading. Can be disabled, use system defaults, or specify custom order.
	BackupLlmConfig *PromptAgentApiModelInputBackupLlmConfig `json:"backup_llm_config,omitempty" xml:"backup_llm_config,omitempty"`
	// Time in seconds before cascading to backup LLM. Must be between 2 and 15 seconds.
	CascadeTimeoutSeconds *float64 `json:"cascade_timeout_seconds,omitempty" xml:"cascade_timeout_seconds,omitempty" min:"2" max:"15"`
	// A list of tools that the agent can use over the course of the conversation, use tool_ids instead
	Tools           []*PromptAgentApiModelInputToolsItem `json:"tools,omitempty" xml:"tools,omitempty"`
	extraProperties map[string]interface{}
}

func (p *PromptAgentApiModelInput) GetPrompt() (value *string) {
	if p == nil {
		return
	}
	return p.Prompt
}

func (p *PromptAgentApiModelInput) GetLlm() (value *Llm) {
	if p == nil {
		return
	}
	return p.Llm
}

func (p *PromptAgentApiModelInput) GetReasoningEffort() (value *LlmReasoningEffort) {
	if p == nil {
		return
	}
	return p.ReasoningEffort
}

func (p *PromptAgentApiModelInput) GetServiceTier() (value *LlmServiceTier) {
	if p == nil {
		return
	}
	return p.ServiceTier
}

func (p *PromptAgentApiModelInput) GetThinkingBudget() (value *int) {
	if p == nil {
		return
	}
	return p.ThinkingBudget
}

func (p *PromptAgentApiModelInput) GetEnableReasoningSummary() (value *bool) {
	if p == nil {
		return
	}
	return p.EnableReasoningSummary
}

func (p *PromptAgentApiModelInput) GetTemperature() (value *float64) {
	if p == nil {
		return
	}
	return p.Temperature
}

func (p *PromptAgentApiModelInput) GetMaxTokens() (value *int) {
	if p == nil {
		return
	}
	return p.MaxTokens
}

func (p *PromptAgentApiModelInput) GetToolIds() (value []string) {
	if p == nil {
		return
	}
	return p.ToolIds
}

func (p *PromptAgentApiModelInput) GetToolIdsStrict() (value *bool) {
	if p == nil {
		return
	}
	return p.ToolIdsStrict
}

func (p *PromptAgentApiModelInput) GetBuiltInTools() (value *BuiltInToolsInput) {
	if p == nil {
		return
	}
	return p.BuiltInTools
}

func (p *PromptAgentApiModelInput) GetEnableParallelToolCalls() (value *bool) {
	if p == nil {
		return
	}
	return p.EnableParallelToolCalls
}

func (p *PromptAgentApiModelInput) GetMcpServerIds() (value []string) {
	if p == nil {
		return
	}
	return p.McpServerIds
}

func (p *PromptAgentApiModelInput) GetNativeMcpServerIds() (value []string) {
	if p == nil {
		return
	}
	return p.NativeMcpServerIds
}

func (p *PromptAgentApiModelInput) GetKnowledgeBase() (value []*KnowledgeBaseLocator) {
	if p == nil {
		return
	}
	return p.KnowledgeBase
}

func (p *PromptAgentApiModelInput) GetCustomLlm() (value *CustomLlm) {
	if p == nil {
		return
	}
	return p.CustomLlm
}

func (p *PromptAgentApiModelInput) GetIgnoreDefaultPersonality() (value *bool) {
	if p == nil {
		return
	}
	return p.IgnoreDefaultPersonality
}

func (p *PromptAgentApiModelInput) GetRag() (value *RagConfigInput) {
	if p == nil {
		return
	}
	return p.Rag
}

func (p *PromptAgentApiModelInput) GetTimezone() (value *string) {
	if p == nil {
		return
	}
	return p.Timezone
}

func (p *PromptAgentApiModelInput) GetBackupLlmConfig() (value *PromptAgentApiModelInputBackupLlmConfig) {
	if p == nil {
		return
	}
	return p.BackupLlmConfig
}

func (p *PromptAgentApiModelInput) GetCascadeTimeoutSeconds() (value *float64) {
	if p == nil {
		return
	}
	return p.CascadeTimeoutSeconds
}

func (p *PromptAgentApiModelInput) GetTools() (value []*PromptAgentApiModelInputToolsItem) {
	if p == nil {
		return
	}
	return p.Tools
}

func (p *PromptAgentApiModelInput) SetPrompt(prompt *string) {
	p.Prompt = prompt
}

func (p *PromptAgentApiModelInput) SetLlm(llm *Llm) {
	p.Llm = llm
}

func (p *PromptAgentApiModelInput) SetReasoningEffort(reasoningEffort *LlmReasoningEffort) {
	p.ReasoningEffort = reasoningEffort
}

func (p *PromptAgentApiModelInput) SetServiceTier(serviceTier *LlmServiceTier) {
	p.ServiceTier = serviceTier
}

func (p *PromptAgentApiModelInput) SetThinkingBudget(thinkingBudget *int) {
	p.ThinkingBudget = thinkingBudget
}

func (p *PromptAgentApiModelInput) SetEnableReasoningSummary(enableReasoningSummary *bool) {
	p.EnableReasoningSummary = enableReasoningSummary
}

func (p *PromptAgentApiModelInput) SetTemperature(temperature *float64) {
	p.Temperature = temperature
}

func (p *PromptAgentApiModelInput) SetMaxTokens(maxTokens *int) {
	p.MaxTokens = maxTokens
}

func (p *PromptAgentApiModelInput) SetToolIds(toolIds []string) {
	p.ToolIds = toolIds
}

func (p *PromptAgentApiModelInput) SetToolIdsStrict(toolIdsStrict *bool) {
	p.ToolIdsStrict = toolIdsStrict
}

func (p *PromptAgentApiModelInput) SetBuiltInTools(builtInTools *BuiltInToolsInput) {
	p.BuiltInTools = builtInTools
}

func (p *PromptAgentApiModelInput) SetEnableParallelToolCalls(enableParallelToolCalls *bool) {
	p.EnableParallelToolCalls = enableParallelToolCalls
}

func (p *PromptAgentApiModelInput) SetMcpServerIds(mcpServerIds []string) {
	p.McpServerIds = mcpServerIds
}

func (p *PromptAgentApiModelInput) SetNativeMcpServerIds(nativeMcpServerIds []string) {
	p.NativeMcpServerIds = nativeMcpServerIds
}

func (p *PromptAgentApiModelInput) SetKnowledgeBase(knowledgeBase []*KnowledgeBaseLocator) {
	p.KnowledgeBase = knowledgeBase
}

func (p *PromptAgentApiModelInput) SetCustomLlm(customLlm *CustomLlm) {
	p.CustomLlm = customLlm
}

func (p *PromptAgentApiModelInput) SetIgnoreDefaultPersonality(ignoreDefaultPersonality *bool) {
	p.IgnoreDefaultPersonality = ignoreDefaultPersonality
}

func (p *PromptAgentApiModelInput) SetRag(rag *RagConfigInput) {
	p.Rag = rag
}

func (p *PromptAgentApiModelInput) SetTimezone(timezone *string) {
	p.Timezone = timezone
}

func (p *PromptAgentApiModelInput) SetBackupLlmConfig(backupLlmConfig *PromptAgentApiModelInputBackupLlmConfig) {
	p.BackupLlmConfig = backupLlmConfig
}

func (p *PromptAgentApiModelInput) SetCascadeTimeoutSeconds(cascadeTimeoutSeconds *float64) {
	p.CascadeTimeoutSeconds = cascadeTimeoutSeconds
}

func (p *PromptAgentApiModelInput) SetTools(tools []*PromptAgentApiModelInputToolsItem) {
	p.Tools = tools
}

func (p *PromptAgentApiModelInput) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PromptAgentApiModelInput) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PromptAgentApiModelInput to string"
	}
	return string(jsonData)
}

func (p *PromptAgentApiModelInput) UnmarshalJSON(data []byte) error {
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

func (p *PromptAgentApiModelInput) MarshalJSON() ([]byte, error) {
	type embed PromptAgentApiModelInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
