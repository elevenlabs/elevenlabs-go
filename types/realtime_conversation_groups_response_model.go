package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type RealtimeConversationGroupsResponseModel struct {
	ConversationGroups []*RealtimeConversationGroupResponseModel `json:"conversation_groups" xml:"conversation_groups" required:"true"`
	// Cursor to pass to the next request to fetch the following page.
	NextCursor *string `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	// Whether more conversation groups are available beyond this page.
	HasMore         *bool `json:"has_more,omitempty" xml:"has_more,omitempty"`
	extraProperties map[string]interface{}
}

func (r *RealtimeConversationGroupsResponseModel) GetConversationGroups() (value []*RealtimeConversationGroupResponseModel) {
	if r == nil {
		return
	}
	return r.ConversationGroups
}

func (r *RealtimeConversationGroupsResponseModel) GetNextCursor() (value *string) {
	if r == nil {
		return
	}
	return r.NextCursor
}

func (r *RealtimeConversationGroupsResponseModel) GetHasMore() (value *bool) {
	if r == nil {
		return
	}
	return r.HasMore
}

func (r *RealtimeConversationGroupsResponseModel) SetConversationGroups(conversationGroups []*RealtimeConversationGroupResponseModel) {
	r.ConversationGroups = conversationGroups
}

func (r *RealtimeConversationGroupsResponseModel) SetNextCursor(nextCursor *string) {
	r.NextCursor = nextCursor
}

func (r *RealtimeConversationGroupsResponseModel) SetHasMore(hasMore *bool) {
	r.HasMore = hasMore
}

func (r *RealtimeConversationGroupsResponseModel) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r RealtimeConversationGroupsResponseModel) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RealtimeConversationGroupsResponseModel to string"
	}
	return string(jsonData)
}

func (r *RealtimeConversationGroupsResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, r); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *RealtimeConversationGroupsResponseModel) MarshalJSON() ([]byte, error) {
	type embed RealtimeConversationGroupsResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
