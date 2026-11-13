package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetLibraryVoicesResponse struct {
	// The list of shared voices
	Voices []*LibraryVoiceResponse `json:"voices" xml:"voices" required:"true"`
	// Whether there are more shared voices in subsequent pages.
	HasMore bool `json:"has_more" xml:"has_more" required:"true"`
	// The total number of shared voices matching the query.
	TotalCount      *int    `json:"total_count,omitempty" xml:"total_count,omitempty"`
	LastSortId      *string `json:"last_sort_id,omitempty" xml:"last_sort_id,omitempty"`
	extraProperties map[string]interface{}
}

func (g *GetLibraryVoicesResponse) GetVoices() (value []*LibraryVoiceResponse) {
	if g == nil {
		return
	}
	return g.Voices
}

func (g *GetLibraryVoicesResponse) GetHasMore() (value bool) {
	if g == nil {
		return
	}
	return g.HasMore
}

func (g *GetLibraryVoicesResponse) GetTotalCount() (value *int) {
	if g == nil {
		return
	}
	return g.TotalCount
}

func (g *GetLibraryVoicesResponse) GetLastSortId() (value *string) {
	if g == nil {
		return
	}
	return g.LastSortId
}

func (g *GetLibraryVoicesResponse) SetVoices(voices []*LibraryVoiceResponse) {
	g.Voices = voices
}

func (g *GetLibraryVoicesResponse) SetHasMore(hasMore bool) {
	g.HasMore = hasMore
}

func (g *GetLibraryVoicesResponse) SetTotalCount(totalCount *int) {
	g.TotalCount = totalCount
}

func (g *GetLibraryVoicesResponse) SetLastSortId(lastSortId *string) {
	g.LastSortId = lastSortId
}

func (g *GetLibraryVoicesResponse) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetLibraryVoicesResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetLibraryVoicesResponse to string"
	}
	return string(jsonData)
}

func (g *GetLibraryVoicesResponse) UnmarshalJSON(data []byte) error {
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

func (g *GetLibraryVoicesResponse) MarshalJSON() ([]byte, error) {
	type embed GetLibraryVoicesResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
