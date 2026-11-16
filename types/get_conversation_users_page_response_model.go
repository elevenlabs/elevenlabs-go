package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetConversationUsersPageResponseModel struct {
	Users           []*ConversationUserResponseModel `json:"users" xml:"users" required:"true"`
	NextCursor      *string                          `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	HasMore         bool                             `json:"has_more" xml:"has_more" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetConversationUsersPageResponseModel) GetUsers() (value []*ConversationUserResponseModel) {
	if g == nil {
		return
	}
	return g.Users
}

func (g *GetConversationUsersPageResponseModel) GetNextCursor() (value *string) {
	if g == nil {
		return
	}
	return g.NextCursor
}

func (g *GetConversationUsersPageResponseModel) GetHasMore() (value bool) {
	if g == nil {
		return
	}
	return g.HasMore
}

func (g *GetConversationUsersPageResponseModel) SetUsers(users []*ConversationUserResponseModel) {
	g.Users = users
}

func (g *GetConversationUsersPageResponseModel) SetNextCursor(nextCursor *string) {
	g.NextCursor = nextCursor
}

func (g *GetConversationUsersPageResponseModel) SetHasMore(hasMore bool) {
	g.HasMore = hasMore
}

func (g *GetConversationUsersPageResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetConversationUsersPageResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetConversationUsersPageResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetConversationUsersPageResponseModel) UnmarshalJSON(data []byte) error {
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

func (g *GetConversationUsersPageResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetConversationUsersPageResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
