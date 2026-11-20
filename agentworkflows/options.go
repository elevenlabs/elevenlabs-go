package agentworkflows

import (
	"fmt"

	"github.com/elevenlabs/elevenlabs-go/internal/wirejson"
	"github.com/elevenlabs/elevenlabs-go/types"
)

// WorkflowOptions are the optional settings of New.
type WorkflowOptions struct {
	// PreventSubagentLoops stops the conversation from re-entering a subagent it already visited.
	// Nil leaves the field out.
	PreventSubagentLoops *bool
}

// Options are the options of the node verbs (Converse, Tool, Transfer, End) and of Node.Update.
// Every verb takes the same struct; a field that doesn't apply to the node's kind is reported by
// Build as an invalid_option error. Nil (or empty, for ID) means "not set".
//
// The json tags are the options' canonical names, so options stored as JSON decode into this
// struct directly. Nested values use the SDK's generated types.
type Options struct {
	// Position is the node's place in the workflow editor. Applies to every node kind; nodes
	// without one are laid out automatically.
	Position *types.PositionInput `json:"position,omitempty"`

	// Label is the subagent's name in the editor (converse). Defaults to the node id.
	Label *string `json:"label,omitempty"`
	// Prompt replaces a subagent's prompt (converse, through Update).
	Prompt *string `json:"prompt,omitempty"`
	// ToolIDs are a subagent's additional tools (converse), or the tools a tool node runs (tool,
	// through Update), by id. A non-nil empty slice clears them.
	ToolIDs []string `json:"tool_ids,omitempty"`
	// Tools are tools written as code (such as agenttools tools), used like ToolIDs: a subagent's
	// additional tools (converse), or more tools a tool node runs (tool, and through Update). They
	// come after ToolIDs, and each one's id is filled in when the workflow is built with
	// BuildOptions.ToolIDs, as agentbuilder's push does. Setting ToolIDs or Tools in Update replaces
	// both; a non-nil empty slice clears them.
	Tools []Tool `json:"-"`
	// KnowledgeBase is a subagent's additional knowledge base (converse).
	KnowledgeBase []*types.KnowledgeBaseLocator `json:"knowledge_base,omitempty"`
	// ConversationConfig overrides the agent's conversation config in a subagent (converse).
	ConversationConfig *types.ConversationalConfigApiModelWorkflowOverrideInput `json:"conversation_config,omitempty"`
	// EntryBehavior is what a subagent does when the conversation enters it (converse).
	EntryBehavior *types.EntryBehavior `json:"entry_behavior,omitempty"`

	// SchemaOverrides are per-tool parameter overrides, keyed by tool id (or tool name, for a Tool
	// in Tools) and then by dotted parameter path (tool).
	SchemaOverrides map[string]map[string]*types.WorkflowToolLocatorSchemaOverridesValue `json:"schema_overrides,omitempty"`

	// TransferType, CustomSipHeaders, SipReferPlayDialtone, Uui and PostDialDigits apply to phone
	// and SIP transfers.
	TransferType         *types.TransferTypeEnum                                        `json:"transfer_type,omitempty"`
	CustomSipHeaders     []*types.WorkflowPhoneNumberNodeModelInputCustomSipHeadersItem `json:"custom_sip_headers,omitempty"`
	SipReferPlayDialtone *bool                                                          `json:"sip_refer_play_dialtone,omitempty"`
	Uui                  *types.UuiTransferConfig                                       `json:"uui,omitempty"`
	PostDialDigits       *types.WorkflowPhoneNumberNodeModelInputPostDialDigits         `json:"post_dial_digits,omitempty"`

	// DelayMs, TransferMessage, EnableTransferredAgentFirstMessage, PreserveClientTtsOverrides and
	// PreserveVoiceSettings apply to agent transfers.
	DelayMs                            *int    `json:"delay_ms,omitempty"`
	TransferMessage                    *string `json:"transfer_message,omitempty"`
	EnableTransferredAgentFirstMessage *bool   `json:"enable_transferred_agent_first_message,omitempty"`
	PreserveClientTtsOverrides         *bool   `json:"preserve_client_tts_overrides,omitempty"`
	PreserveVoiceSettings              *bool   `json:"preserve_voice_settings,omitempty"`

	// Destination changes where a transfer goes (transfer, through Update). A destination of the
	// other transfer kind changes the node's type and keeps only its position.
	Destination *Destination `json:"-"`

	// ID is an end node's id (End only). Empty means "end". Reusing an end node's id reuses that
	// node.
	ID string `json:"id,omitempty"`
}

// WhenOptions are the optional settings of When.
type WhenOptions struct {
	// Back is the condition for coming back along the edge, from its target to its source.
	Back *Condition
}

// GoToOptions are the optional settings of GoTo.
type GoToOptions struct {
	// Back is the condition for coming back along the edge, from its target to its source.
	Back *Condition
}

// ResultOptions are the optional settings of OnSuccess and OnFailure.
type ResultOptions struct {
	// Label is shown on the edge in the workflow editor.
	Label string
}

// ResultBranches are the callbacks of OnResult. Either may be nil. Success always runs first.
type ResultBranches struct {
	Success func(ok *Branch)
	Failure func(failed *Branch)
}

// BuildOptions are the optional settings of Build, ToJSON and ToJSONString.
type BuildOptions struct {
	// OnWarning is called once per warning (unreachable_node, unknown_node_type). Build only.
	OnWarning func(Issue)
	// ToolIDs maps the name of each Tool the workflow's nodes use (see Options.Tools) to its id.
	// agentbuilder's push passes these after creating the tools. Every Tool needs one; a missing
	// one is reported as unresolved_tool.
	ToolIDs map[string]string
}

// resolvedOptions are Options in wire format: the special options apart, and the rest as
// canonical option names mapped to wire values, in a fixed order.
type resolvedOptions struct {
	destination     *Destination
	toolIDs         []string
	tools           []Tool
	hasToolIDs      bool
	schemaOverrides map[string]any
	hasOverrides    bool
	rest            []option
}

type option struct {
	key   string
	value any
}

// resolve converts the options to wire format. A value that can't be converted to JSON is
// recorded as an invalid_option error on node id.
func (o *Options) resolve(g *graph, id string) resolvedOptions {
	var r resolvedOptions
	if o == nil {
		return r
	}
	r.destination = o.Destination
	if o.ToolIDs != nil || o.Tools != nil {
		r.toolIDs = append([]string{}, o.ToolIDs...)
		r.tools = append([]Tool{}, o.Tools...)
		r.hasToolIDs = true
	}
	if o.SchemaOverrides != nil {
		if value, ok := g.wireValue(id, "schema_overrides", o.SchemaOverrides); ok {
			if m, isMap := value.(map[string]any); isMap {
				r.schemaOverrides = m
				r.hasOverrides = true
			}
		}
	}
	add := func(key string, set bool, value any) {
		if !set {
			return
		}
		if wire, ok := g.wireValue(id, key, value); ok {
			r.rest = append(r.rest, option{key: key, value: wire})
		}
	}
	add("position", o.Position != nil, o.Position)
	add("label", o.Label != nil, o.Label)
	add("prompt", o.Prompt != nil, o.Prompt)
	add("knowledge_base", o.KnowledgeBase != nil, o.KnowledgeBase)
	add("conversation_config", o.ConversationConfig != nil, o.ConversationConfig)
	add("entry_behavior", o.EntryBehavior != nil, o.EntryBehavior)
	add("transfer_type", o.TransferType != nil, o.TransferType)
	add("custom_sip_headers", o.CustomSipHeaders != nil, o.CustomSipHeaders)
	add("sip_refer_play_dialtone", o.SipReferPlayDialtone != nil, o.SipReferPlayDialtone)
	add("uui", o.Uui != nil, o.Uui)
	add("post_dial_digits", o.PostDialDigits != nil, o.PostDialDigits)
	add("delay_ms", o.DelayMs != nil, o.DelayMs)
	add("transfer_message", o.TransferMessage != nil, o.TransferMessage)
	add("enable_transferred_agent_first_message", o.EnableTransferredAgentFirstMessage != nil, o.EnableTransferredAgentFirstMessage)
	add("preserve_client_tts_overrides", o.PreserveClientTtsOverrides != nil, o.PreserveClientTtsOverrides)
	add("preserve_voice_settings", o.PreserveVoiceSettings != nil, o.PreserveVoiceSettings)
	add("id", o.ID != "", o.ID)
	return r
}

// wireValue converts an option value to a wire tree, recording invalid_option when it can't be.
func (g *graph) wireValue(id, key string, value any) (any, bool) {
	wire, err := wirejson.Normalize(value)
	if err != nil {
		g.issue(InvalidOption, fmt.Sprintf("Option %q of node %q can't be converted to JSON: %v", key, id, err), id)
		return nil, false
	}
	return wire, true
}
