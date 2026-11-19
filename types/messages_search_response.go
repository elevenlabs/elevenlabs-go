package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type MessagesSearchResponse struct {
	Meta    *ListResponseMeta       `json:"meta,omitempty" xml:"meta,omitempty"`
	Results []*MessagesSearchResult `json:"results" xml:"results" required:"true"`
	// Cursor for the next page of results
	NextCursor *string `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	// Whether there are more results available
	HasMore         bool `json:"has_more" xml:"has_more" required:"true"`
	extraProperties map[string]interface{}
}

func (m *MessagesSearchResponse) GetMeta() (value *ListResponseMeta) {
	if m == nil {
		return
	}
	return m.Meta
}

func (m *MessagesSearchResponse) GetResults() (value []*MessagesSearchResult) {
	if m == nil {
		return
	}
	return m.Results
}

func (m *MessagesSearchResponse) GetNextCursor() (value *string) {
	if m == nil {
		return
	}
	return m.NextCursor
}

func (m *MessagesSearchResponse) GetHasMore() (value bool) {
	if m == nil {
		return
	}
	return m.HasMore
}

func (m *MessagesSearchResponse) SetMeta(meta *ListResponseMeta) {
	m.Meta = meta
}

func (m *MessagesSearchResponse) SetResults(results []*MessagesSearchResult) {
	m.Results = results
}

func (m *MessagesSearchResponse) SetNextCursor(nextCursor *string) {
	m.NextCursor = nextCursor
}

func (m *MessagesSearchResponse) SetHasMore(hasMore bool) {
	m.HasMore = hasMore
}

func (m *MessagesSearchResponse) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m MessagesSearchResponse) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MessagesSearchResponse to string"
	}
	return string(jsonData)
}

func (m *MessagesSearchResponse) UnmarshalJSON(data []byte) error {
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

func (m *MessagesSearchResponse) MarshalJSON() ([]byte, error) {
	type embed MessagesSearchResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
