package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetKnowledgeBaseListResponseModel struct {
	Documents       []*GetKnowledgeBaseListResponseModelDocumentsItem `json:"documents" xml:"documents" required:"true"`
	NextCursor      *string                                           `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	HasMore         bool                                              `json:"has_more" xml:"has_more" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetKnowledgeBaseListResponseModel) GetDocuments() (value []*GetKnowledgeBaseListResponseModelDocumentsItem) {
	if g == nil {
		return
	}
	return g.Documents
}

func (g *GetKnowledgeBaseListResponseModel) GetNextCursor() (value *string) {
	if g == nil {
		return
	}
	return g.NextCursor
}

func (g *GetKnowledgeBaseListResponseModel) GetHasMore() (value bool) {
	if g == nil {
		return
	}
	return g.HasMore
}

func (g *GetKnowledgeBaseListResponseModel) SetDocuments(documents []*GetKnowledgeBaseListResponseModelDocumentsItem) {
	g.Documents = documents
}

func (g *GetKnowledgeBaseListResponseModel) SetNextCursor(nextCursor *string) {
	g.NextCursor = nextCursor
}

func (g *GetKnowledgeBaseListResponseModel) SetHasMore(hasMore bool) {
	g.HasMore = hasMore
}

func (g *GetKnowledgeBaseListResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetKnowledgeBaseListResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetKnowledgeBaseListResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetKnowledgeBaseListResponseModel) UnmarshalJSON(data []byte) error {
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

func (g *GetKnowledgeBaseListResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetKnowledgeBaseListResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
