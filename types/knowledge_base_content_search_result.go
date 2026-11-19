package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type KnowledgeBaseContentSearchResult struct {
	Document        *KnowledgeBaseContentSearchResultDocument `json:"document" xml:"document" required:"true"`
	SearchSnippet   []*SearchHighlightSegment                 `json:"search_snippet,omitempty" xml:"search_snippet,omitempty"`
	Score           float64                                   `json:"score" xml:"score" required:"true"`
	extraProperties map[string]interface{}
}

func (k *KnowledgeBaseContentSearchResult) GetDocument() (value *KnowledgeBaseContentSearchResultDocument) {
	if k == nil {
		return
	}
	return k.Document
}

func (k *KnowledgeBaseContentSearchResult) GetSearchSnippet() (value []*SearchHighlightSegment) {
	if k == nil {
		return
	}
	return k.SearchSnippet
}

func (k *KnowledgeBaseContentSearchResult) GetScore() (value float64) {
	if k == nil {
		return
	}
	return k.Score
}

func (k *KnowledgeBaseContentSearchResult) SetDocument(document *KnowledgeBaseContentSearchResultDocument) {
	k.Document = document
}

func (k *KnowledgeBaseContentSearchResult) SetSearchSnippet(searchSnippet []*SearchHighlightSegment) {
	k.SearchSnippet = searchSnippet
}

func (k *KnowledgeBaseContentSearchResult) SetScore(score float64) {
	k.Score = score
}

func (k *KnowledgeBaseContentSearchResult) GetExtraProperties() map[string]interface{} {
	if k == nil {
		return nil
	}
	return k.extraProperties
}

func (k KnowledgeBaseContentSearchResult) String() string {
	jsonData, err := json.MarshalIndent(k, "", "  ")
	if err != nil {
		return "error converting struct: KnowledgeBaseContentSearchResult to string"
	}
	return string(jsonData)
}

func (k *KnowledgeBaseContentSearchResult) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, k); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *k)
	if err != nil {
		return err
	}
	k.extraProperties = extraProperties
	return nil
}

func (k *KnowledgeBaseContentSearchResult) MarshalJSON() ([]byte, error) {
	type embed KnowledgeBaseContentSearchResult
	var marshaler = struct {
		embed
	}{
		embed: embed(*k),
	}
	return json.Marshal(marshaler)
}
