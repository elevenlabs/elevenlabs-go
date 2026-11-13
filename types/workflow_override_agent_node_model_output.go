package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkflowOverrideAgentNodeModelOutput struct {
	type_    string
	Position *PositionOutput `json:"position" xml:"position" required:"true"`
	// The ids of outgoing edges in the order they should be evaluated.
	EdgeOrder          []string                                            `json:"edge_order" xml:"edge_order" required:"true"`
	ConversationConfig *ConversationalConfigApiModelWorkflowOverrideOutput `json:"conversation_config" xml:"conversation_config" required:"true"`
	// Specific goal for this subagent. It will be added to the system prompt and can be used to further refine the agent's behavior in this specific context.
	AdditionalPrompt string `json:"additional_prompt" xml:"additional_prompt" required:"true"`
	// Additional knowledge base documents that the subagent has access to. These will be used in addition to the main agent's documents.
	AdditionalKnowledgeBase []*KnowledgeBaseLocator `json:"additional_knowledge_base" xml:"additional_knowledge_base" required:"true"`
	// IDs of additional tools that the subagent has access to. These will be used in addition to the main agent's tools.
	AdditionalToolIds []string `json:"additional_tool_ids" xml:"additional_tool_ids" required:"true"`
	// Human-readable label for the node used throughout the UI.
	Label           string        `json:"label" xml:"label" required:"true"`
	EntryBehavior   EntryBehavior `json:"entry_behavior" xml:"entry_behavior" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WorkflowOverrideAgentNodeModelOutput) GetPosition() (value *PositionOutput) {
	if w == nil {
		return
	}
	return w.Position
}

func (w *WorkflowOverrideAgentNodeModelOutput) GetEdgeOrder() (value []string) {
	if w == nil {
		return
	}
	return w.EdgeOrder
}

func (w *WorkflowOverrideAgentNodeModelOutput) GetConversationConfig() (value *ConversationalConfigApiModelWorkflowOverrideOutput) {
	if w == nil {
		return
	}
	return w.ConversationConfig
}

func (w *WorkflowOverrideAgentNodeModelOutput) GetAdditionalPrompt() (value string) {
	if w == nil {
		return
	}
	return w.AdditionalPrompt
}

func (w *WorkflowOverrideAgentNodeModelOutput) GetAdditionalKnowledgeBase() (value []*KnowledgeBaseLocator) {
	if w == nil {
		return
	}
	return w.AdditionalKnowledgeBase
}

func (w *WorkflowOverrideAgentNodeModelOutput) GetAdditionalToolIds() (value []string) {
	if w == nil {
		return
	}
	return w.AdditionalToolIds
}

func (w *WorkflowOverrideAgentNodeModelOutput) GetLabel() (value string) {
	if w == nil {
		return
	}
	return w.Label
}

func (w *WorkflowOverrideAgentNodeModelOutput) GetEntryBehavior() (value EntryBehavior) {
	if w == nil {
		return
	}
	return w.EntryBehavior
}

func (w *WorkflowOverrideAgentNodeModelOutput) SetPosition(position *PositionOutput) {
	w.Position = position
}

func (w *WorkflowOverrideAgentNodeModelOutput) SetEdgeOrder(edgeOrder []string) {
	w.EdgeOrder = edgeOrder
}

func (w *WorkflowOverrideAgentNodeModelOutput) SetConversationConfig(conversationConfig *ConversationalConfigApiModelWorkflowOverrideOutput) {
	w.ConversationConfig = conversationConfig
}

func (w *WorkflowOverrideAgentNodeModelOutput) SetAdditionalPrompt(additionalPrompt string) {
	w.AdditionalPrompt = additionalPrompt
}

func (w *WorkflowOverrideAgentNodeModelOutput) SetAdditionalKnowledgeBase(additionalKnowledgeBase []*KnowledgeBaseLocator) {
	w.AdditionalKnowledgeBase = additionalKnowledgeBase
}

func (w *WorkflowOverrideAgentNodeModelOutput) SetAdditionalToolIds(additionalToolIds []string) {
	w.AdditionalToolIds = additionalToolIds
}

func (w *WorkflowOverrideAgentNodeModelOutput) SetLabel(label string) {
	w.Label = label
}

func (w *WorkflowOverrideAgentNodeModelOutput) SetEntryBehavior(entryBehavior EntryBehavior) {
	w.EntryBehavior = entryBehavior
}

func (w *WorkflowOverrideAgentNodeModelOutput) Type() string {
	return w.type_
}

func (w *WorkflowOverrideAgentNodeModelOutput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowOverrideAgentNodeModelOutput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowOverrideAgentNodeModelOutput to string"
	}
	return string(jsonData)
}

func (w *WorkflowOverrideAgentNodeModelOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WorkflowOverrideAgentNodeModelOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WorkflowOverrideAgentNodeModelOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "override_agent" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", w, "override_agent", *literals.Type)
	}
	w.type_ = "override_agent"
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowOverrideAgentNodeModelOutput) MarshalJSON() ([]byte, error) {
	type embed WorkflowOverrideAgentNodeModelOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*w),
		Type:  "override_agent",
	}
	return json.Marshal(marshaler)
}
