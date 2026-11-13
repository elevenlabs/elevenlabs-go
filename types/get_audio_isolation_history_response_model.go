package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetAudioIsolationHistoryResponseModel struct {
	Items           []*AudioIsolationHistoryItemResponseModel `json:"items" xml:"items" required:"true"`
	HasMore         bool                                      `json:"has_more" xml:"has_more" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetAudioIsolationHistoryResponseModel) GetItems() (value []*AudioIsolationHistoryItemResponseModel) {
	if g == nil {
		return
	}
	return g.Items
}

func (g *GetAudioIsolationHistoryResponseModel) GetHasMore() (value bool) {
	if g == nil {
		return
	}
	return g.HasMore
}

func (g *GetAudioIsolationHistoryResponseModel) SetItems(items []*AudioIsolationHistoryItemResponseModel) {
	g.Items = items
}

func (g *GetAudioIsolationHistoryResponseModel) SetHasMore(hasMore bool) {
	g.HasMore = hasMore
}

func (g *GetAudioIsolationHistoryResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetAudioIsolationHistoryResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetAudioIsolationHistoryResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetAudioIsolationHistoryResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, g); err != nil {
		return err
	}
	type alias GetAudioIsolationHistoryResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*g = GetAudioIsolationHistoryResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetAudioIsolationHistoryResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetAudioIsolationHistoryResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
