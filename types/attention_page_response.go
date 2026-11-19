package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AttentionPageResponse struct {
	Attention       []*AttentionView `json:"attention" xml:"attention" required:"true"`
	NextCursor      *string          `json:"next_cursor" xml:"next_cursor" required:"true"`
	HasMore         bool             `json:"has_more" xml:"has_more" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AttentionPageResponse) GetAttention() (value []*AttentionView) {
	if a == nil {
		return
	}
	return a.Attention
}

func (a *AttentionPageResponse) GetNextCursor() (value *string) {
	if a == nil {
		return
	}
	return a.NextCursor
}

func (a *AttentionPageResponse) GetHasMore() (value bool) {
	if a == nil {
		return
	}
	return a.HasMore
}

func (a *AttentionPageResponse) SetAttention(attention []*AttentionView) {
	a.Attention = attention
}

func (a *AttentionPageResponse) SetNextCursor(nextCursor *string) {
	a.NextCursor = nextCursor
}

func (a *AttentionPageResponse) SetHasMore(hasMore bool) {
	a.HasMore = hasMore
}

func (a *AttentionPageResponse) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AttentionPageResponse) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AttentionPageResponse to string"
	}
	return string(jsonData)
}

func (a *AttentionPageResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AttentionPageResponse) MarshalJSON() ([]byte, error) {
	type embed AttentionPageResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
