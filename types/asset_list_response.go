package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// One page of the workspace's assets, most recently created first.
type AssetListResponse struct {
	// List of Asset objects.
	Assets []*AssetResponse `json:"assets" xml:"assets" required:"true"`
	// Pass as `cursor` to fetch the next page. `null` if there are no more results.
	NextCursor *string `json:"next_cursor" xml:"next_cursor" required:"true"`
	// Whether there are more results to fetch.
	HasMore         *bool `json:"has_more,omitempty" xml:"has_more,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AssetListResponse) GetAssets() (value []*AssetResponse) {
	if a == nil {
		return
	}
	return a.Assets
}

func (a *AssetListResponse) GetNextCursor() (value *string) {
	if a == nil {
		return
	}
	return a.NextCursor
}

func (a *AssetListResponse) GetHasMore() (value *bool) {
	if a == nil {
		return
	}
	return a.HasMore
}

func (a *AssetListResponse) SetAssets(assets []*AssetResponse) {
	a.Assets = assets
}

func (a *AssetListResponse) SetNextCursor(nextCursor *string) {
	a.NextCursor = nextCursor
}

func (a *AssetListResponse) SetHasMore(hasMore *bool) {
	a.HasMore = hasMore
}

func (a *AssetListResponse) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AssetListResponse) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AssetListResponse to string"
	}
	return string(jsonData)
}

func (a *AssetListResponse) UnmarshalJSON(data []byte) error {
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

func (a *AssetListResponse) MarshalJSON() ([]byte, error) {
	type embed AssetListResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
