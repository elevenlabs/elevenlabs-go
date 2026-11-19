package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AnalysisItemsResponseModel struct {
	AnalysisItems   []*AnalysisItemResponseModel `json:"analysis_items" xml:"analysis_items" required:"true"`
	NextCursor      *string                      `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	HasMore         bool                         `json:"has_more" xml:"has_more" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AnalysisItemsResponseModel) GetAnalysisItems() (value []*AnalysisItemResponseModel) {
	if a == nil {
		return
	}
	return a.AnalysisItems
}

func (a *AnalysisItemsResponseModel) GetNextCursor() (value *string) {
	if a == nil {
		return
	}
	return a.NextCursor
}

func (a *AnalysisItemsResponseModel) GetHasMore() (value bool) {
	if a == nil {
		return
	}
	return a.HasMore
}

func (a *AnalysisItemsResponseModel) SetAnalysisItems(analysisItems []*AnalysisItemResponseModel) {
	a.AnalysisItems = analysisItems
}

func (a *AnalysisItemsResponseModel) SetNextCursor(nextCursor *string) {
	a.NextCursor = nextCursor
}

func (a *AnalysisItemsResponseModel) SetHasMore(hasMore bool) {
	a.HasMore = hasMore
}

func (a *AnalysisItemsResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AnalysisItemsResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AnalysisItemsResponseModel to string"
	}
	return string(jsonData)
}

func (a *AnalysisItemsResponseModel) UnmarshalJSON(data []byte) error {
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

func (a *AnalysisItemsResponseModel) MarshalJSON() ([]byte, error) {
	type embed AnalysisItemsResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
