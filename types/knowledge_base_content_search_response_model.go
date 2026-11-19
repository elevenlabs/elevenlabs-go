package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type KnowledgeBaseContentSearchResponseModel struct {
	Results         []*KnowledgeBaseContentSearchResult `json:"results" xml:"results" required:"true"`
	NextCursor      *string                             `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	extraProperties map[string]interface{}
}

func (k *KnowledgeBaseContentSearchResponseModel) GetResults() (value []*KnowledgeBaseContentSearchResult) {
	if k == nil {
		return
	}
	return k.Results
}

func (k *KnowledgeBaseContentSearchResponseModel) GetNextCursor() (value *string) {
	if k == nil {
		return
	}
	return k.NextCursor
}

func (k *KnowledgeBaseContentSearchResponseModel) SetResults(results []*KnowledgeBaseContentSearchResult) {
	k.Results = results
}

func (k *KnowledgeBaseContentSearchResponseModel) SetNextCursor(nextCursor *string) {
	k.NextCursor = nextCursor
}

func (k *KnowledgeBaseContentSearchResponseModel) GetExtraProperties() map[string]interface{} {
	if k == nil {
		return nil
	}
	return k.extraProperties
}

func (k KnowledgeBaseContentSearchResponseModel) String() string {
	jsonData, err := json.MarshalIndent(k, "", "  ")
	if err != nil {
		return "error converting struct: KnowledgeBaseContentSearchResponseModel to string"
	}
	return string(jsonData)
}

func (k *KnowledgeBaseContentSearchResponseModel) UnmarshalJSON(data []byte) error {
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

func (k *KnowledgeBaseContentSearchResponseModel) MarshalJSON() ([]byte, error) {
	type embed KnowledgeBaseContentSearchResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*k),
	}
	return json.Marshal(marshaler)
}
