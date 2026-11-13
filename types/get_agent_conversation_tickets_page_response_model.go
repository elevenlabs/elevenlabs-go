package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetAgentConversationTicketsPageResponseModel struct {
	AgentConversationTickets []*AgentConversationTicketResponseModel `json:"agent_conversation_tickets" xml:"agent_conversation_tickets" required:"true"`
	NextCursor               *string                                 `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	HasMore                  bool                                    `json:"has_more" xml:"has_more" required:"true"`
	extraProperties          map[string]interface{}
}

func (g *GetAgentConversationTicketsPageResponseModel) GetAgentConversationTickets() (value []*AgentConversationTicketResponseModel) {
	if g == nil {
		return
	}
	return g.AgentConversationTickets
}

func (g *GetAgentConversationTicketsPageResponseModel) GetNextCursor() (value *string) {
	if g == nil {
		return
	}
	return g.NextCursor
}

func (g *GetAgentConversationTicketsPageResponseModel) GetHasMore() (value bool) {
	if g == nil {
		return
	}
	return g.HasMore
}

func (g *GetAgentConversationTicketsPageResponseModel) SetAgentConversationTickets(agentConversationTickets []*AgentConversationTicketResponseModel) {
	g.AgentConversationTickets = agentConversationTickets
}

func (g *GetAgentConversationTicketsPageResponseModel) SetNextCursor(nextCursor *string) {
	g.NextCursor = nextCursor
}

func (g *GetAgentConversationTicketsPageResponseModel) SetHasMore(hasMore bool) {
	g.HasMore = hasMore
}

func (g *GetAgentConversationTicketsPageResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetAgentConversationTicketsPageResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetAgentConversationTicketsPageResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetAgentConversationTicketsPageResponseModel) UnmarshalJSON(data []byte) error {
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

func (g *GetAgentConversationTicketsPageResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetAgentConversationTicketsPageResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
