package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PromptAgentApiModelOutput struct {
	// The prompt for the agent
	Prompt *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	Llm    *Llm    `json:"llm,omitempty" xml:"llm,omitempty"`
	// Reasoning effort of the model. Only available for some models.
	ReasoningEffort *LlmReasoningEffort `json:"reasoning_effort,omitempty" xml:"reasoning_effort,omitempty"`
	// Request-time processing tier for a model.
	//
	// STANDARD: default price and throughput.
	// FAST: pays a per-token price premium for higher throughput / lower latency.
	// See OpenAI's fast mode: https://openai.com/api-fast-mode/
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
	BuiltInTools *BuiltInToolsOutput `json:"built_in_tools,omitempty" xml:"built_in_tools,omitempty"`
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
	IgnoreDefaultPersonality *bool            `json:"ignore_default_personality,omitempty" xml:"ignore_default_personality,omitempty"`
	Rag                      *RagConfigOutput `json:"rag,omitempty" xml:"rag,omitempty"`
	// Timezone for displaying current time in system prompt. If set, the current time will be included in the system prompt using this timezone. Must be a valid timezone name (e.g., 'America/New_York', 'Europe/London', 'UTC'). Recommended for accurate time-aware responses; without this, the agent has no knowledge of the current date/time unless you provide it via dynamic variables or tools, which can lead to incorrect or hallucinated time references.
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// Configuration for backup LLM cascading. Can be disabled, use system defaults, or specify custom order.
	BackupLlmConfig *PromptAgentApiModelOutputBackupLlmConfig `json:"backup_llm_config,omitempty" xml:"backup_llm_config,omitempty"`
	// Time in seconds before cascading to backup LLM. Must be between 2 and 15 seconds.
	CascadeTimeoutSeconds *float64 `json:"cascade_timeout_seconds,omitempty" xml:"cascade_timeout_seconds,omitempty" min:"2" max:"15"`
	// A list of tools that the agent can use over the course of the conversation, use tool_ids instead
	Tools           []*PromptAgentApiModelOutputToolsItem `json:"tools,omitempty" xml:"tools,omitempty"`
	extraProperties map[string]interface{}
}

func (p *PromptAgentApiModelOutput) GetPrompt() (value *string) {
	if p == nil {
		return
	}
	return p.Prompt
}

func (p *PromptAgentApiModelOutput) GetLlm() (value *Llm) {
	if p == nil {
		return
	}
	return p.Llm
}

func (p *PromptAgentApiModelOutput) GetReasoningEffort() (value *LlmReasoningEffort) {
	if p == nil {
		return
	}
	return p.ReasoningEffort
}

func (p *PromptAgentApiModelOutput) GetServiceTier() (value *LlmServiceTier) {
	if p == nil {
		return
	}
	return p.ServiceTier
}

func (p *PromptAgentApiModelOutput) GetThinkingBudget() (value *int) {
	if p == nil {
		return
	}
	return p.ThinkingBudget
}

func (p *PromptAgentApiModelOutput) GetEnableReasoningSummary() (value *bool) {
	if p == nil {
		return
	}
	return p.EnableReasoningSummary
}

func (p *PromptAgentApiModelOutput) GetTemperature() (value *float64) {
	if p == nil {
		return
	}
	return p.Temperature
}

func (p *PromptAgentApiModelOutput) GetMaxTokens() (value *int) {
	if p == nil {
		return
	}
	return p.MaxTokens
}

func (p *PromptAgentApiModelOutput) GetToolIds() (value []string) {
	if p == nil {
		return
	}
	return p.ToolIds
}

func (p *PromptAgentApiModelOutput) GetToolIdsStrict() (value *bool) {
	if p == nil {
		return
	}
	return p.ToolIdsStrict
}

func (p *PromptAgentApiModelOutput) GetBuiltInTools() (value *BuiltInToolsOutput) {
	if p == nil {
		return
	}
	return p.BuiltInTools
}

func (p *PromptAgentApiModelOutput) GetEnableParallelToolCalls() (value *bool) {
	if p == nil {
		return
	}
	return p.EnableParallelToolCalls
}

func (p *PromptAgentApiModelOutput) GetMcpServerIds() (value []string) {
	if p == nil {
		return
	}
	return p.McpServerIds
}

func (p *PromptAgentApiModelOutput) GetNativeMcpServerIds() (value []string) {
	if p == nil {
		return
	}
	return p.NativeMcpServerIds
}

func (p *PromptAgentApiModelOutput) GetKnowledgeBase() (value []*KnowledgeBaseLocator) {
	if p == nil {
		return
	}
	return p.KnowledgeBase
}

func (p *PromptAgentApiModelOutput) GetCustomLlm() (value *CustomLlm) {
	if p == nil {
		return
	}
	return p.CustomLlm
}

func (p *PromptAgentApiModelOutput) GetIgnoreDefaultPersonality() (value *bool) {
	if p == nil {
		return
	}
	return p.IgnoreDefaultPersonality
}

func (p *PromptAgentApiModelOutput) GetRag() (value *RagConfigOutput) {
	if p == nil {
		return
	}
	return p.Rag
}

func (p *PromptAgentApiModelOutput) GetTimezone() (value *string) {
	if p == nil {
		return
	}
	return p.Timezone
}

func (p *PromptAgentApiModelOutput) GetBackupLlmConfig() (value *PromptAgentApiModelOutputBackupLlmConfig) {
	if p == nil {
		return
	}
	return p.BackupLlmConfig
}

func (p *PromptAgentApiModelOutput) GetCascadeTimeoutSeconds() (value *float64) {
	if p == nil {
		return
	}
	return p.CascadeTimeoutSeconds
}

func (p *PromptAgentApiModelOutput) GetTools() (value []*PromptAgentApiModelOutputToolsItem) {
	if p == nil {
		return
	}
	return p.Tools
}

func (p *PromptAgentApiModelOutput) SetPrompt(prompt *string) {
	p.Prompt = prompt
}

func (p *PromptAgentApiModelOutput) SetLlm(llm *Llm) {
	p.Llm = llm
}

func (p *PromptAgentApiModelOutput) SetReasoningEffort(reasoningEffort *LlmReasoningEffort) {
	p.ReasoningEffort = reasoningEffort
}

func (p *PromptAgentApiModelOutput) SetServiceTier(serviceTier *LlmServiceTier) {
	p.ServiceTier = serviceTier
}

func (p *PromptAgentApiModelOutput) SetThinkingBudget(thinkingBudget *int) {
	p.ThinkingBudget = thinkingBudget
}

func (p *PromptAgentApiModelOutput) SetEnableReasoningSummary(enableReasoningSummary *bool) {
	p.EnableReasoningSummary = enableReasoningSummary
}

func (p *PromptAgentApiModelOutput) SetTemperature(temperature *float64) {
	p.Temperature = temperature
}

func (p *PromptAgentApiModelOutput) SetMaxTokens(maxTokens *int) {
	p.MaxTokens = maxTokens
}

func (p *PromptAgentApiModelOutput) SetToolIds(toolIds []string) {
	p.ToolIds = toolIds
}

func (p *PromptAgentApiModelOutput) SetToolIdsStrict(toolIdsStrict *bool) {
	p.ToolIdsStrict = toolIdsStrict
}

func (p *PromptAgentApiModelOutput) SetBuiltInTools(builtInTools *BuiltInToolsOutput) {
	p.BuiltInTools = builtInTools
}

func (p *PromptAgentApiModelOutput) SetEnableParallelToolCalls(enableParallelToolCalls *bool) {
	p.EnableParallelToolCalls = enableParallelToolCalls
}

func (p *PromptAgentApiModelOutput) SetMcpServerIds(mcpServerIds []string) {
	p.McpServerIds = mcpServerIds
}

func (p *PromptAgentApiModelOutput) SetNativeMcpServerIds(nativeMcpServerIds []string) {
	p.NativeMcpServerIds = nativeMcpServerIds
}

func (p *PromptAgentApiModelOutput) SetKnowledgeBase(knowledgeBase []*KnowledgeBaseLocator) {
	p.KnowledgeBase = knowledgeBase
}

func (p *PromptAgentApiModelOutput) SetCustomLlm(customLlm *CustomLlm) {
	p.CustomLlm = customLlm
}

func (p *PromptAgentApiModelOutput) SetIgnoreDefaultPersonality(ignoreDefaultPersonality *bool) {
	p.IgnoreDefaultPersonality = ignoreDefaultPersonality
}

func (p *PromptAgentApiModelOutput) SetRag(rag *RagConfigOutput) {
	p.Rag = rag
}

func (p *PromptAgentApiModelOutput) SetTimezone(timezone *string) {
	p.Timezone = timezone
}

func (p *PromptAgentApiModelOutput) SetBackupLlmConfig(backupLlmConfig *PromptAgentApiModelOutputBackupLlmConfig) {
	p.BackupLlmConfig = backupLlmConfig
}

func (p *PromptAgentApiModelOutput) SetCascadeTimeoutSeconds(cascadeTimeoutSeconds *float64) {
	p.CascadeTimeoutSeconds = cascadeTimeoutSeconds
}

func (p *PromptAgentApiModelOutput) SetTools(tools []*PromptAgentApiModelOutputToolsItem) {
	p.Tools = tools
}

func (p *PromptAgentApiModelOutput) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PromptAgentApiModelOutput) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PromptAgentApiModelOutput to string"
	}
	return string(jsonData)
}

func (p *PromptAgentApiModelOutput) UnmarshalJSON(data []byte) error {
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

func (p *PromptAgentApiModelOutput) MarshalJSON() ([]byte, error) {
	type embed PromptAgentApiModelOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
