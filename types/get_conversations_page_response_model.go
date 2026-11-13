package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetConversationsPageResponseModel struct {
	Conversations   []*ConversationSummaryResponseModel `json:"conversations" xml:"conversations" required:"true"`
	NextCursor      *string                             `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	HasMore         bool                                `json:"has_more" xml:"has_more" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetConversationsPageResponseModel) GetConversations() (value []*ConversationSummaryResponseModel) {
	if g == nil {
		return
	}
	return g.Conversations
}

func (g *GetConversationsPageResponseModel) GetNextCursor() (value *string) {
	if g == nil {
		return
	}
	return g.NextCursor
}

func (g *GetConversationsPageResponseModel) GetHasMore() (value bool) {
	if g == nil {
		return
	}
	return g.HasMore
}

func (g *GetConversationsPageResponseModel) SetConversations(conversations []*ConversationSummaryResponseModel) {
	g.Conversations = conversations
}

func (g *GetConversationsPageResponseModel) SetNextCursor(nextCursor *string) {
	g.NextCursor = nextCursor
}

func (g *GetConversationsPageResponseModel) SetHasMore(hasMore bool) {
	g.HasMore = hasMore
}

func (g *GetConversationsPageResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetConversationsPageResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetConversationsPageResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetConversationsPageResponseModel) UnmarshalJSON(data []byte) error {
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

func (g *GetConversationsPageResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetConversationsPageResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
