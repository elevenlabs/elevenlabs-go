package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Published versions, newest first; current_version_id is the main that unpinned refs resolve to.
type AnalysisItemVersionsResponseModel struct {
	Versions         []*AnalysisItemVersionResponseModel `json:"versions" xml:"versions" required:"true"`
	CurrentVersionId string                              `json:"current_version_id" xml:"current_version_id" required:"true"`
	NextCursor       *string                             `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	HasMore          bool                                `json:"has_more" xml:"has_more" required:"true"`
	extraProperties  map[string]interface{}
}

func (a *AnalysisItemVersionsResponseModel) GetVersions() (value []*AnalysisItemVersionResponseModel) {
	if a == nil {
		return
	}
	return a.Versions
}

func (a *AnalysisItemVersionsResponseModel) GetCurrentVersionId() (value string) {
	if a == nil {
		return
	}
	return a.CurrentVersionId
}

func (a *AnalysisItemVersionsResponseModel) GetNextCursor() (value *string) {
	if a == nil {
		return
	}
	return a.NextCursor
}

func (a *AnalysisItemVersionsResponseModel) GetHasMore() (value bool) {
	if a == nil {
		return
	}
	return a.HasMore
}

func (a *AnalysisItemVersionsResponseModel) SetVersions(versions []*AnalysisItemVersionResponseModel) {
	a.Versions = versions
}

func (a *AnalysisItemVersionsResponseModel) SetCurrentVersionId(currentVersionId string) {
	a.CurrentVersionId = currentVersionId
}

func (a *AnalysisItemVersionsResponseModel) SetNextCursor(nextCursor *string) {
	a.NextCursor = nextCursor
}

func (a *AnalysisItemVersionsResponseModel) SetHasMore(hasMore bool) {
	a.HasMore = hasMore
}

func (a *AnalysisItemVersionsResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AnalysisItemVersionsResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AnalysisItemVersionsResponseModel to string"
	}
	return string(jsonData)
}

func (a *AnalysisItemVersionsResponseModel) UnmarshalJSON(data []byte) error {
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

func (a *AnalysisItemVersionsResponseModel) MarshalJSON() ([]byte, error) {
	type embed AnalysisItemVersionsResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
