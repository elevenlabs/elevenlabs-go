package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetConversationTagsPageResponseModel struct {
	ConversationTags []*ConversationTagResponseModel `json:"conversation_tags" xml:"conversation_tags" required:"true"`
	NextCursor       *string                         `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	HasMore          bool                            `json:"has_more" xml:"has_more" required:"true"`
	extraProperties  map[string]interface{}
}

func (g *GetConversationTagsPageResponseModel) GetConversationTags() (value []*ConversationTagResponseModel) {
	if g == nil {
		return
	}
	return g.ConversationTags
}

func (g *GetConversationTagsPageResponseModel) GetNextCursor() (value *string) {
	if g == nil {
		return
	}
	return g.NextCursor
}

func (g *GetConversationTagsPageResponseModel) GetHasMore() (value bool) {
	if g == nil {
		return
	}
	return g.HasMore
}

func (g *GetConversationTagsPageResponseModel) SetConversationTags(conversationTags []*ConversationTagResponseModel) {
	g.ConversationTags = conversationTags
}

func (g *GetConversationTagsPageResponseModel) SetNextCursor(nextCursor *string) {
	g.NextCursor = nextCursor
}

func (g *GetConversationTagsPageResponseModel) SetHasMore(hasMore bool) {
	g.HasMore = hasMore
}

func (g *GetConversationTagsPageResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetConversationTagsPageResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetConversationTagsPageResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetConversationTagsPageResponseModel) UnmarshalJSON(data []byte) error {
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

func (g *GetConversationTagsPageResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetConversationTagsPageResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
