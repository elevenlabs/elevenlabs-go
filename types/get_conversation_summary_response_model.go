package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetConversationSummaryResponseModel struct {
	ConversationId string                                    `json:"conversation_id" xml:"conversation_id" required:"true"`
	AgentId        string                                    `json:"agent_id" xml:"agent_id" required:"true"`
	Status         GetConversationSummaryResponseModelStatus `json:"status" xml:"status" required:"true"`
	// Short generated title for the conversation.
	CallSummaryTitle *string `json:"call_summary_title,omitempty" xml:"call_summary_title,omitempty"`
	// Generated natural-language summary of the call.
	TranscriptSummary *string                  `json:"transcript_summary,omitempty" xml:"transcript_summary,omitempty"`
	CallSuccessful    *EvaluationSuccessResult `json:"call_successful,omitempty" xml:"call_successful,omitempty"`
	// Number of plain chat message turns in the conversation.
	MessageCount int `json:"message_count" xml:"message_count" required:"true"`
	// The plain chat messages (role and text only). Included only when message_count does not exceed the requested max_messages; otherwise null and messages_omitted is true.
	Messages []*ConversationSummaryMessageModel `json:"messages,omitempty" xml:"messages,omitempty"`
	// True when the chat messages were omitted because the conversation was too long. Fetch the full transcript for the messages.
	MessagesOmitted *bool `json:"messages_omitted,omitempty" xml:"messages_omitted,omitempty"`
	// Guidance telling the agent how to get the full transcript.
	Note            string `json:"note" xml:"note" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetConversationSummaryResponseModel) GetConversationId() (value string) {
	if g == nil {
		return
	}
	return g.ConversationId
}

func (g *GetConversationSummaryResponseModel) GetAgentId() (value string) {
	if g == nil {
		return
	}
	return g.AgentId
}

func (g *GetConversationSummaryResponseModel) GetStatus() (value GetConversationSummaryResponseModelStatus) {
	if g == nil {
		return
	}
	return g.Status
}

func (g *GetConversationSummaryResponseModel) GetCallSummaryTitle() (value *string) {
	if g == nil {
		return
	}
	return g.CallSummaryTitle
}

func (g *GetConversationSummaryResponseModel) GetTranscriptSummary() (value *string) {
	if g == nil {
		return
	}
	return g.TranscriptSummary
}

func (g *GetConversationSummaryResponseModel) GetCallSuccessful() (value *EvaluationSuccessResult) {
	if g == nil {
		return
	}
	return g.CallSuccessful
}

func (g *GetConversationSummaryResponseModel) GetMessageCount() (value int) {
	if g == nil {
		return
	}
	return g.MessageCount
}

func (g *GetConversationSummaryResponseModel) GetMessages() (value []*ConversationSummaryMessageModel) {
	if g == nil {
		return
	}
	return g.Messages
}

func (g *GetConversationSummaryResponseModel) GetMessagesOmitted() (value *bool) {
	if g == nil {
		return
	}
	return g.MessagesOmitted
}

func (g *GetConversationSummaryResponseModel) GetNote() (value string) {
	if g == nil {
		return
	}
	return g.Note
}

func (g *GetConversationSummaryResponseModel) SetConversationId(conversationId string) {
	g.ConversationId = conversationId
}

func (g *GetConversationSummaryResponseModel) SetAgentId(agentId string) {
	g.AgentId = agentId
}

func (g *GetConversationSummaryResponseModel) SetStatus(status GetConversationSummaryResponseModelStatus) {
	g.Status = status
}

func (g *GetConversationSummaryResponseModel) SetCallSummaryTitle(callSummaryTitle *string) {
	g.CallSummaryTitle = callSummaryTitle
}

func (g *GetConversationSummaryResponseModel) SetTranscriptSummary(transcriptSummary *string) {
	g.TranscriptSummary = transcriptSummary
}

func (g *GetConversationSummaryResponseModel) SetCallSuccessful(callSuccessful *EvaluationSuccessResult) {
	g.CallSuccessful = callSuccessful
}

func (g *GetConversationSummaryResponseModel) SetMessageCount(messageCount int) {
	g.MessageCount = messageCount
}

func (g *GetConversationSummaryResponseModel) SetMessages(messages []*ConversationSummaryMessageModel) {
	g.Messages = messages
}

func (g *GetConversationSummaryResponseModel) SetMessagesOmitted(messagesOmitted *bool) {
	g.MessagesOmitted = messagesOmitted
}

func (g *GetConversationSummaryResponseModel) SetNote(note string) {
	g.Note = note
}

func (g *GetConversationSummaryResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetConversationSummaryResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetConversationSummaryResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetConversationSummaryResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, g); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetConversationSummaryResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetConversationSummaryResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
