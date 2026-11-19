package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkflowOverrideAgentNodeModelInput struct {
	Position *PositionInput `json:"position,omitempty" xml:"position,omitempty"`
	// The ids of outgoing edges in the order they should be evaluated.
	EdgeOrder          []string                                           `json:"edge_order,omitempty" xml:"edge_order,omitempty"`
	ConversationConfig *ConversationalConfigApiModelWorkflowOverrideInput `json:"conversation_config,omitempty" xml:"conversation_config,omitempty"`
	// Specific goal for this subagent. It will be added to the system prompt and can be used to further refine the agent's behavior in this specific context.
	AdditionalPrompt *string `json:"additional_prompt,omitempty" xml:"additional_prompt,omitempty"`
	// Additional knowledge base documents that the subagent has access to. These will be used in addition to the main agent's documents.
	AdditionalKnowledgeBase []*KnowledgeBaseLocator `json:"additional_knowledge_base,omitempty" xml:"additional_knowledge_base,omitempty"`
	// IDs of additional tools that the subagent has access to. These will be used in addition to the main agent's tools.
	AdditionalToolIds []string `json:"additional_tool_ids,omitempty" xml:"additional_tool_ids,omitempty"`
	// Human-readable label for the node used throughout the UI.
	Label         string         `json:"label" xml:"label" required:"true"`
	EntryBehavior *EntryBehavior `json:"entry_behavior,omitempty" xml:"entry_behavior,omitempty"`
	// Lets the caller interrupt the response this node gives on entry. Ignored unless entry_behavior is "generate_immediately".
	AllowInterruptionsOnEntry *bool `json:"allow_interruptions_on_entry,omitempty" xml:"allow_interruptions_on_entry,omitempty"`
	extraProperties           map[string]interface{}
}

func (w *WorkflowOverrideAgentNodeModelInput) GetPosition() (value *PositionInput) {
	if w == nil {
		return
	}
	return w.Position
}

func (w *WorkflowOverrideAgentNodeModelInput) GetEdgeOrder() (value []string) {
	if w == nil {
		return
	}
	return w.EdgeOrder
}

func (w *WorkflowOverrideAgentNodeModelInput) GetConversationConfig() (value *ConversationalConfigApiModelWorkflowOverrideInput) {
	if w == nil {
		return
	}
	return w.ConversationConfig
}

func (w *WorkflowOverrideAgentNodeModelInput) GetAdditionalPrompt() (value *string) {
	if w == nil {
		return
	}
	return w.AdditionalPrompt
}

func (w *WorkflowOverrideAgentNodeModelInput) GetAdditionalKnowledgeBase() (value []*KnowledgeBaseLocator) {
	if w == nil {
		return
	}
	return w.AdditionalKnowledgeBase
}

func (w *WorkflowOverrideAgentNodeModelInput) GetAdditionalToolIds() (value []string) {
	if w == nil {
		return
	}
	return w.AdditionalToolIds
}

func (w *WorkflowOverrideAgentNodeModelInput) GetLabel() (value string) {
	if w == nil {
		return
	}
	return w.Label
}

func (w *WorkflowOverrideAgentNodeModelInput) GetEntryBehavior() (value *EntryBehavior) {
	if w == nil {
		return
	}
	return w.EntryBehavior
}

func (w *WorkflowOverrideAgentNodeModelInput) GetAllowInterruptionsOnEntry() (value *bool) {
	if w == nil {
		return
	}
	return w.AllowInterruptionsOnEntry
}

func (w *WorkflowOverrideAgentNodeModelInput) SetPosition(position *PositionInput) {
	w.Position = position
}

func (w *WorkflowOverrideAgentNodeModelInput) SetEdgeOrder(edgeOrder []string) {
	w.EdgeOrder = edgeOrder
}

func (w *WorkflowOverrideAgentNodeModelInput) SetConversationConfig(conversationConfig *ConversationalConfigApiModelWorkflowOverrideInput) {
	w.ConversationConfig = conversationConfig
}

func (w *WorkflowOverrideAgentNodeModelInput) SetAdditionalPrompt(additionalPrompt *string) {
	w.AdditionalPrompt = additionalPrompt
}

func (w *WorkflowOverrideAgentNodeModelInput) SetAdditionalKnowledgeBase(additionalKnowledgeBase []*KnowledgeBaseLocator) {
	w.AdditionalKnowledgeBase = additionalKnowledgeBase
}

func (w *WorkflowOverrideAgentNodeModelInput) SetAdditionalToolIds(additionalToolIds []string) {
	w.AdditionalToolIds = additionalToolIds
}

func (w *WorkflowOverrideAgentNodeModelInput) SetLabel(label string) {
	w.Label = label
}

func (w *WorkflowOverrideAgentNodeModelInput) SetEntryBehavior(entryBehavior *EntryBehavior) {
	w.EntryBehavior = entryBehavior
}

func (w *WorkflowOverrideAgentNodeModelInput) SetAllowInterruptionsOnEntry(allowInterruptionsOnEntry *bool) {
	w.AllowInterruptionsOnEntry = allowInterruptionsOnEntry
}

func (w *WorkflowOverrideAgentNodeModelInput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowOverrideAgentNodeModelInput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowOverrideAgentNodeModelInput to string"
	}
	return string(jsonData)
}

func (w *WorkflowOverrideAgentNodeModelInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WorkflowOverrideAgentNodeModelInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WorkflowOverrideAgentNodeModelInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowOverrideAgentNodeModelInput) MarshalJSON() ([]byte, error) {
	type embed WorkflowOverrideAgentNodeModelInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
