package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentMessagesOverrides struct {
	// The prompt for the agent
	Prompt *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
	// The LLM to query with the prompt and the chat history. If using data residency, the LLM must be supported in the data residency environment
	Llm *Llm `json:"llm,omitempty" xml:"llm,omitempty"`
	// Language of the agent
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// If set, only these tools are available in the conversation, including tools on workflow nodes. Each ID must be attached to the agent.
	ToolIds []string `json:"tool_ids,omitempty" xml:"tool_ids,omitempty"`
	// If set, the agent uses only these native MCP servers. Each ID must be attached to the agent.
	NativeMcpServerIds []string `json:"native_mcp_server_ids,omitempty" xml:"native_mcp_server_ids,omitempty"`
	// If set, the agent uses only these knowledge base documents. Each ID must be attached to the agent.
	KnowledgeBaseIds []string `json:"knowledge_base_ids,omitempty" xml:"knowledge_base_ids,omitempty"`
	// If set, only these procedures are available to the starting agent. Each ID must be attached to that agent. An empty list disables all of that agent's procedures. Not applied after an agent transfer. Requires enable_procedure_ids_from_client.
	ProcedureIds []string `json:"procedure_ids,omitempty" xml:"procedure_ids,omitempty"`
	// Extra fields for the request body that goes to the custom LLM.
	CustomLlmExtraBody map[string]any `json:"custom_llm_extra_body,omitempty"`
	extraProperties    map[string]interface{}
}

func (a *AgentMessagesOverrides) GetPrompt() (value *string) {
	if a == nil {
		return
	}
	return a.Prompt
}

func (a *AgentMessagesOverrides) GetLlm() (value *Llm) {
	if a == nil {
		return
	}
	return a.Llm
}

func (a *AgentMessagesOverrides) GetLanguage() (value *string) {
	if a == nil {
		return
	}
	return a.Language
}

func (a *AgentMessagesOverrides) GetToolIds() (value []string) {
	if a == nil {
		return
	}
	return a.ToolIds
}

func (a *AgentMessagesOverrides) GetNativeMcpServerIds() (value []string) {
	if a == nil {
		return
	}
	return a.NativeMcpServerIds
}

func (a *AgentMessagesOverrides) GetKnowledgeBaseIds() (value []string) {
	if a == nil {
		return
	}
	return a.KnowledgeBaseIds
}

func (a *AgentMessagesOverrides) GetProcedureIds() (value []string) {
	if a == nil {
		return
	}
	return a.ProcedureIds
}

func (a *AgentMessagesOverrides) GetCustomLlmExtraBody() (value map[string]any) {
	if a == nil {
		return
	}
	return a.CustomLlmExtraBody
}

func (a *AgentMessagesOverrides) SetPrompt(prompt *string) {
	a.Prompt = prompt
}

func (a *AgentMessagesOverrides) SetLlm(llm *Llm) {
	a.Llm = llm
}

func (a *AgentMessagesOverrides) SetLanguage(language *string) {
	a.Language = language
}

func (a *AgentMessagesOverrides) SetToolIds(toolIds []string) {
	a.ToolIds = toolIds
}

func (a *AgentMessagesOverrides) SetNativeMcpServerIds(nativeMcpServerIds []string) {
	a.NativeMcpServerIds = nativeMcpServerIds
}

func (a *AgentMessagesOverrides) SetKnowledgeBaseIds(knowledgeBaseIds []string) {
	a.KnowledgeBaseIds = knowledgeBaseIds
}

func (a *AgentMessagesOverrides) SetProcedureIds(procedureIds []string) {
	a.ProcedureIds = procedureIds
}

func (a *AgentMessagesOverrides) SetCustomLlmExtraBody(customLlmExtraBody map[string]any) {
	a.CustomLlmExtraBody = customLlmExtraBody
}

func (a *AgentMessagesOverrides) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentMessagesOverrides) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentMessagesOverrides to string"
	}
	return string(jsonData)
}

func (a *AgentMessagesOverrides) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentMessagesOverrides) MarshalJSON() ([]byte, error) {
	type embed AgentMessagesOverrides
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
