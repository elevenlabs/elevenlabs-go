package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkflowSayNodeModelInput struct {
	Position *PositionInput `json:"position,omitempty" xml:"position,omitempty"`
	// The ids of outgoing edges in the order they should be evaluated.
	EdgeOrder          []string                                           `json:"edge_order,omitempty" xml:"edge_order,omitempty"`
	ConversationConfig *ConversationalConfigApiModelWorkflowOverrideInput `json:"conversation_config,omitempty" xml:"conversation_config,omitempty"`
	// Accepted for consistency with subagent nodes but not used. Describe the message in `message.prompt`, or replace the system prompt with `conversation_config.agent.prompt.prompt`.
	AdditionalPrompt *string `json:"additional_prompt,omitempty" xml:"additional_prompt,omitempty"`
	// Knowledge base documents available to the LLM while this node generates its message, in addition to the agent's documents.
	AdditionalKnowledgeBase []*KnowledgeBaseLocator `json:"additional_knowledge_base,omitempty" xml:"additional_knowledge_base,omitempty"`
	// Accepted for consistency with subagent nodes but not used. A say node cannot call tools.
	AdditionalToolIds []string `json:"additional_tool_ids,omitempty" xml:"additional_tool_ids,omitempty"`
	// Human-readable label for the node used throughout the UI.
	Label string `json:"label" xml:"label" required:"true"`
	// Message spoken by this node. Use `literal` for exact text or `prompt` to have the LLM generate the message from a description.
	Message *WorkflowSayNodeModelInputMessage `json:"message" xml:"message" required:"true"`
	// When true, the workflow only moves to the next node after the message has been fully spoken.
	WaitForSpeechCompletion *bool `json:"wait_for_speech_completion,omitempty" xml:"wait_for_speech_completion,omitempty"`
	// When true, the node speaks its message and then waits for the user to reply before evaluating outgoing edges, so edges can branch on the reply. If the user stays silent, the workflow advances after the turn timeout.
	WaitForUserResponse *bool `json:"wait_for_user_response,omitempty" xml:"wait_for_user_response,omitempty"`
	extraProperties     map[string]interface{}
}

func (w *WorkflowSayNodeModelInput) GetPosition() (value *PositionInput) {
	if w == nil {
		return
	}
	return w.Position
}

func (w *WorkflowSayNodeModelInput) GetEdgeOrder() (value []string) {
	if w == nil {
		return
	}
	return w.EdgeOrder
}

func (w *WorkflowSayNodeModelInput) GetConversationConfig() (value *ConversationalConfigApiModelWorkflowOverrideInput) {
	if w == nil {
		return
	}
	return w.ConversationConfig
}

func (w *WorkflowSayNodeModelInput) GetAdditionalPrompt() (value *string) {
	if w == nil {
		return
	}
	return w.AdditionalPrompt
}

func (w *WorkflowSayNodeModelInput) GetAdditionalKnowledgeBase() (value []*KnowledgeBaseLocator) {
	if w == nil {
		return
	}
	return w.AdditionalKnowledgeBase
}

func (w *WorkflowSayNodeModelInput) GetAdditionalToolIds() (value []string) {
	if w == nil {
		return
	}
	return w.AdditionalToolIds
}

func (w *WorkflowSayNodeModelInput) GetLabel() (value string) {
	if w == nil {
		return
	}
	return w.Label
}

func (w *WorkflowSayNodeModelInput) GetMessage() (value *WorkflowSayNodeModelInputMessage) {
	if w == nil {
		return
	}
	return w.Message
}

func (w *WorkflowSayNodeModelInput) GetWaitForSpeechCompletion() (value *bool) {
	if w == nil {
		return
	}
	return w.WaitForSpeechCompletion
}

func (w *WorkflowSayNodeModelInput) GetWaitForUserResponse() (value *bool) {
	if w == nil {
		return
	}
	return w.WaitForUserResponse
}

func (w *WorkflowSayNodeModelInput) SetPosition(position *PositionInput) {
	w.Position = position
}

func (w *WorkflowSayNodeModelInput) SetEdgeOrder(edgeOrder []string) {
	w.EdgeOrder = edgeOrder
}

func (w *WorkflowSayNodeModelInput) SetConversationConfig(conversationConfig *ConversationalConfigApiModelWorkflowOverrideInput) {
	w.ConversationConfig = conversationConfig
}

func (w *WorkflowSayNodeModelInput) SetAdditionalPrompt(additionalPrompt *string) {
	w.AdditionalPrompt = additionalPrompt
}

func (w *WorkflowSayNodeModelInput) SetAdditionalKnowledgeBase(additionalKnowledgeBase []*KnowledgeBaseLocator) {
	w.AdditionalKnowledgeBase = additionalKnowledgeBase
}

func (w *WorkflowSayNodeModelInput) SetAdditionalToolIds(additionalToolIds []string) {
	w.AdditionalToolIds = additionalToolIds
}

func (w *WorkflowSayNodeModelInput) SetLabel(label string) {
	w.Label = label
}

func (w *WorkflowSayNodeModelInput) SetMessage(message *WorkflowSayNodeModelInputMessage) {
	w.Message = message
}

func (w *WorkflowSayNodeModelInput) SetWaitForSpeechCompletion(waitForSpeechCompletion *bool) {
	w.WaitForSpeechCompletion = waitForSpeechCompletion
}

func (w *WorkflowSayNodeModelInput) SetWaitForUserResponse(waitForUserResponse *bool) {
	w.WaitForUserResponse = waitForUserResponse
}

func (w *WorkflowSayNodeModelInput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowSayNodeModelInput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowSayNodeModelInput to string"
	}
	return string(jsonData)
}

func (w *WorkflowSayNodeModelInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WorkflowSayNodeModelInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WorkflowSayNodeModelInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowSayNodeModelInput) MarshalJSON() ([]byte, error) {
	type embed WorkflowSayNodeModelInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
