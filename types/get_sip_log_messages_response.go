package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetSipLogMessagesResponse struct {
	SipMessages     []*SipLogMessage `json:"sip_messages" xml:"sip_messages" required:"true"`
	NextCursor      *string          `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	HasMore         *bool            `json:"has_more,omitempty" xml:"has_more,omitempty"`
	extraProperties map[string]interface{}
}

func (g *GetSipLogMessagesResponse) GetSipMessages() (value []*SipLogMessage) {
	if g == nil {
		return
	}
	return g.SipMessages
}

func (g *GetSipLogMessagesResponse) GetNextCursor() (value *string) {
	if g == nil {
		return
	}
	return g.NextCursor
}

func (g *GetSipLogMessagesResponse) GetHasMore() (value *bool) {
	if g == nil {
		return
	}
	return g.HasMore
}

func (g *GetSipLogMessagesResponse) SetSipMessages(sipMessages []*SipLogMessage) {
	g.SipMessages = sipMessages
}

func (g *GetSipLogMessagesResponse) SetNextCursor(nextCursor *string) {
	g.NextCursor = nextCursor
}

func (g *GetSipLogMessagesResponse) SetHasMore(hasMore *bool) {
	g.HasMore = hasMore
}

func (g *GetSipLogMessagesResponse) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetSipLogMessagesResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetSipLogMessagesResponse to string"
	}
	return string(jsonData)
}

func (g *GetSipLogMessagesResponse) UnmarshalJSON(data []byte) error {
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

func (g *GetSipLogMessagesResponse) MarshalJSON() ([]byte, error) {
	type embed GetSipLogMessagesResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
