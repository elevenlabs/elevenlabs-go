package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetPronunciationDictionariesMetadataResponseModel struct {
	// A list of pronunciation dictionaries and their metadata.
	PronunciationDictionaries []*GetPronunciationDictionaryMetadataResponse `json:"pronunciation_dictionaries" xml:"pronunciation_dictionaries" required:"true"`
	// The next cursor to use for pagination.
	NextCursor *string `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	// Whether there are more pronunciation dictionaries to fetch.
	HasMore         bool `json:"has_more" xml:"has_more" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetPronunciationDictionariesMetadataResponseModel) GetPronunciationDictionaries() (value []*GetPronunciationDictionaryMetadataResponse) {
	if g == nil {
		return
	}
	return g.PronunciationDictionaries
}

func (g *GetPronunciationDictionariesMetadataResponseModel) GetNextCursor() (value *string) {
	if g == nil {
		return
	}
	return g.NextCursor
}

func (g *GetPronunciationDictionariesMetadataResponseModel) GetHasMore() (value bool) {
	if g == nil {
		return
	}
	return g.HasMore
}

func (g *GetPronunciationDictionariesMetadataResponseModel) SetPronunciationDictionaries(pronunciationDictionaries []*GetPronunciationDictionaryMetadataResponse) {
	g.PronunciationDictionaries = pronunciationDictionaries
}

func (g *GetPronunciationDictionariesMetadataResponseModel) SetNextCursor(nextCursor *string) {
	g.NextCursor = nextCursor
}

func (g *GetPronunciationDictionariesMetadataResponseModel) SetHasMore(hasMore bool) {
	g.HasMore = hasMore
}

func (g *GetPronunciationDictionariesMetadataResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetPronunciationDictionariesMetadataResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetPronunciationDictionariesMetadataResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetPronunciationDictionariesMetadataResponseModel) UnmarshalJSON(data []byte) error {
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

func (g *GetPronunciationDictionariesMetadataResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetPronunciationDictionariesMetadataResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
