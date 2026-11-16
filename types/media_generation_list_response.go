package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// One page of the caller's generations, newest first.
type MediaGenerationListResponse struct {
	// The generations on this page, newest first. Each item has the same shape as the corresponding GET endpoint's response.
	Generations []*MediaGenerationResponse `json:"generations" xml:"generations" required:"true"`
	// Pass as `cursor` to fetch the next page. `null` when there is no further page.
	NextCursor *string `json:"next_cursor" xml:"next_cursor" required:"true"`
	// Whether more generations exist beyond this page.
	HasMore         *bool `json:"has_more,omitempty" xml:"has_more,omitempty"`
	extraProperties map[string]interface{}
}

func (m *MediaGenerationListResponse) GetGenerations() (value []*MediaGenerationResponse) {
	if m == nil {
		return
	}
	return m.Generations
}

func (m *MediaGenerationListResponse) GetNextCursor() (value *string) {
	if m == nil {
		return
	}
	return m.NextCursor
}

func (m *MediaGenerationListResponse) GetHasMore() (value *bool) {
	if m == nil {
		return
	}
	return m.HasMore
}

func (m *MediaGenerationListResponse) SetGenerations(generations []*MediaGenerationResponse) {
	m.Generations = generations
}

func (m *MediaGenerationListResponse) SetNextCursor(nextCursor *string) {
	m.NextCursor = nextCursor
}

func (m *MediaGenerationListResponse) SetHasMore(hasMore *bool) {
	m.HasMore = hasMore
}

func (m *MediaGenerationListResponse) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m MediaGenerationListResponse) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MediaGenerationListResponse to string"
	}
	return string(jsonData)
}

func (m *MediaGenerationListResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, m); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *MediaGenerationListResponse) MarshalJSON() ([]byte, error) {
	type embed MediaGenerationListResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
