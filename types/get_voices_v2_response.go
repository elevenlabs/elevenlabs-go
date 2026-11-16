package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetVoicesV2Response struct {
	// The list of voices matching the query.
	Voices []*Voice `json:"voices" xml:"voices" required:"true"`
	// Indicates whether there are more voices available in subsequent pages. Use this flag (and next_page_token) for reliable pagination instead of relying on total_count.
	HasMore bool `json:"has_more" xml:"has_more" required:"true"`
	// The total count of voices matching the query. This value is a live snapshot that reflects the current state of the database and may change between requests as users create, modify, or delete voices. For reliable pagination, use the has_more flag instead of relying on this value. Only request this field when you actually need the total count (e.g., for display purposes), as calculating it incurs a performance cost.
	TotalCount int `json:"total_count" xml:"total_count" required:"true"`
	// Token to retrieve the next page of results. Pass this value to the next request to continue pagination. Null if there are no more results.
	NextPageToken   *string `json:"next_page_token,omitempty" xml:"next_page_token,omitempty"`
	extraProperties map[string]interface{}
}

func (g *GetVoicesV2Response) GetVoices() (value []*Voice) {
	if g == nil {
		return
	}
	return g.Voices
}

func (g *GetVoicesV2Response) GetHasMore() (value bool) {
	if g == nil {
		return
	}
	return g.HasMore
}

func (g *GetVoicesV2Response) GetTotalCount() (value int) {
	if g == nil {
		return
	}
	return g.TotalCount
}

func (g *GetVoicesV2Response) GetNextPageToken() (value *string) {
	if g == nil {
		return
	}
	return g.NextPageToken
}

func (g *GetVoicesV2Response) SetVoices(voices []*Voice) {
	g.Voices = voices
}

func (g *GetVoicesV2Response) SetHasMore(hasMore bool) {
	g.HasMore = hasMore
}

func (g *GetVoicesV2Response) SetTotalCount(totalCount int) {
	g.TotalCount = totalCount
}

func (g *GetVoicesV2Response) SetNextPageToken(nextPageToken *string) {
	g.NextPageToken = nextPageToken
}

func (g *GetVoicesV2Response) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetVoicesV2Response) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetVoicesV2Response to string"
	}
	return string(jsonData)
}

func (g *GetVoicesV2Response) UnmarshalJSON(data []byte) error {
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

func (g *GetVoicesV2Response) MarshalJSON() ([]byte, error) {
	type embed GetVoicesV2Response
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
