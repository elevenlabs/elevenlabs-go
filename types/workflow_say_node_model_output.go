package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkflowSayNodeModelOutput struct {
	type_    string
	Position *PositionOutput `json:"position" xml:"position" required:"true"`
	// The ids of outgoing edges in the order they should be evaluated.
	EdgeOrder          []string                                            `json:"edge_order" xml:"edge_order" required:"true"`
	ConversationConfig *ConversationalConfigApiModelWorkflowOverrideOutput `json:"conversation_config" xml:"conversation_config" required:"true"`
	// Accepted for consistency with subagent nodes but not used. Describe the message in `message.prompt`, or replace the system prompt with `conversation_config.agent.prompt.prompt`.
	AdditionalPrompt string `json:"additional_prompt" xml:"additional_prompt" required:"true"`
	// Knowledge base documents available to the LLM while this node generates its message, in addition to the agent's documents.
	AdditionalKnowledgeBase []*KnowledgeBaseLocator `json:"additional_knowledge_base" xml:"additional_knowledge_base" required:"true"`
	// Accepted for consistency with subagent nodes but not used. A say node cannot call tools.
	AdditionalToolIds []string `json:"additional_tool_ids" xml:"additional_tool_ids" required:"true"`
	// Human-readable label for the node used throughout the UI.
	Label string `json:"label" xml:"label" required:"true"`
	// Message spoken by this node. Use `literal` for exact text or `prompt` to have the LLM generate the message from a description.
	Message *WorkflowSayNodeModelOutputMessage `json:"message" xml:"message" required:"true"`
	// When true, the workflow only moves to the next node after the message has been fully spoken.
	WaitForSpeechCompletion bool `json:"wait_for_speech_completion" xml:"wait_for_speech_completion" required:"true"`
	// When true, the node speaks its message and then waits for the user to reply before evaluating outgoing edges, so edges can branch on the reply. If the user stays silent, the workflow advances after the turn timeout.
	WaitForUserResponse bool `json:"wait_for_user_response" xml:"wait_for_user_response" required:"true"`
	extraProperties     map[string]interface{}
}

func (w *WorkflowSayNodeModelOutput) GetPosition() (value *PositionOutput) {
	if w == nil {
		return
	}
	return w.Position
}

func (w *WorkflowSayNodeModelOutput) GetEdgeOrder() (value []string) {
	if w == nil {
		return
	}
	return w.EdgeOrder
}

func (w *WorkflowSayNodeModelOutput) GetConversationConfig() (value *ConversationalConfigApiModelWorkflowOverrideOutput) {
	if w == nil {
		return
	}
	return w.ConversationConfig
}

func (w *WorkflowSayNodeModelOutput) GetAdditionalPrompt() (value string) {
	if w == nil {
		return
	}
	return w.AdditionalPrompt
}

func (w *WorkflowSayNodeModelOutput) GetAdditionalKnowledgeBase() (value []*KnowledgeBaseLocator) {
	if w == nil {
		return
	}
	return w.AdditionalKnowledgeBase
}

func (w *WorkflowSayNodeModelOutput) GetAdditionalToolIds() (value []string) {
	if w == nil {
		return
	}
	return w.AdditionalToolIds
}

func (w *WorkflowSayNodeModelOutput) GetLabel() (value string) {
	if w == nil {
		return
	}
	return w.Label
}

func (w *WorkflowSayNodeModelOutput) GetMessage() (value *WorkflowSayNodeModelOutputMessage) {
	if w == nil {
		return
	}
	return w.Message
}

func (w *WorkflowSayNodeModelOutput) GetWaitForSpeechCompletion() (value bool) {
	if w == nil {
		return
	}
	return w.WaitForSpeechCompletion
}

func (w *WorkflowSayNodeModelOutput) GetWaitForUserResponse() (value bool) {
	if w == nil {
		return
	}
	return w.WaitForUserResponse
}

func (w *WorkflowSayNodeModelOutput) SetPosition(position *PositionOutput) {
	w.Position = position
}

func (w *WorkflowSayNodeModelOutput) SetEdgeOrder(edgeOrder []string) {
	w.EdgeOrder = edgeOrder
}

func (w *WorkflowSayNodeModelOutput) SetConversationConfig(conversationConfig *ConversationalConfigApiModelWorkflowOverrideOutput) {
	w.ConversationConfig = conversationConfig
}

func (w *WorkflowSayNodeModelOutput) SetAdditionalPrompt(additionalPrompt string) {
	w.AdditionalPrompt = additionalPrompt
}

func (w *WorkflowSayNodeModelOutput) SetAdditionalKnowledgeBase(additionalKnowledgeBase []*KnowledgeBaseLocator) {
	w.AdditionalKnowledgeBase = additionalKnowledgeBase
}

func (w *WorkflowSayNodeModelOutput) SetAdditionalToolIds(additionalToolIds []string) {
	w.AdditionalToolIds = additionalToolIds
}

func (w *WorkflowSayNodeModelOutput) SetLabel(label string) {
	w.Label = label
}

func (w *WorkflowSayNodeModelOutput) SetMessage(message *WorkflowSayNodeModelOutputMessage) {
	w.Message = message
}

func (w *WorkflowSayNodeModelOutput) SetWaitForSpeechCompletion(waitForSpeechCompletion bool) {
	w.WaitForSpeechCompletion = waitForSpeechCompletion
}

func (w *WorkflowSayNodeModelOutput) SetWaitForUserResponse(waitForUserResponse bool) {
	w.WaitForUserResponse = waitForUserResponse
}

func (w *WorkflowSayNodeModelOutput) Type() string {
	return w.type_
}

func (w *WorkflowSayNodeModelOutput) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkflowSayNodeModelOutput) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkflowSayNodeModelOutput to string"
	}
	return string(jsonData)
}

func (w *WorkflowSayNodeModelOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WorkflowSayNodeModelOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WorkflowSayNodeModelOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "say" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", w, "say", *literals.Type)
	}
	w.type_ = "say"
	extraProperties, err := internal.ExtractExtraProperties(data, *w, "type")
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkflowSayNodeModelOutput) MarshalJSON() ([]byte, error) {
	type embed WorkflowSayNodeModelOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*w),
		Type:  "say",
	}
	return json.Marshal(marshaler)
}
