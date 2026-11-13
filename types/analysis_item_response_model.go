package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AnalysisItemResponseModel struct {
	Id                string                               `json:"id" xml:"id" required:"true"`
	Definition        *AnalysisItemResponseModelDefinition `json:"definition" xml:"definition" required:"true"`
	AccessInfo        *ResourceAccessInfo                  `json:"access_info" xml:"access_info" required:"true"`
	CreatedAtUnixSecs int                                  `json:"created_at_unix_secs" xml:"created_at_unix_secs" required:"true"`
	UpdatedAtUnixSecs int                                  `json:"updated_at_unix_secs" xml:"updated_at_unix_secs" required:"true"`
	CurrentVersionId  string                               `json:"current_version_id" xml:"current_version_id" required:"true"`
	LegacySourceId    *string                              `json:"legacy_source_id,omitempty" xml:"legacy_source_id,omitempty"`
	extraProperties   map[string]interface{}
}

func (a *AnalysisItemResponseModel) GetId() (value string) {
	if a == nil {
		return
	}
	return a.Id
}

func (a *AnalysisItemResponseModel) GetDefinition() (value *AnalysisItemResponseModelDefinition) {
	if a == nil {
		return
	}
	return a.Definition
}

func (a *AnalysisItemResponseModel) GetAccessInfo() (value *ResourceAccessInfo) {
	if a == nil {
		return
	}
	return a.AccessInfo
}

func (a *AnalysisItemResponseModel) GetCreatedAtUnixSecs() (value int) {
	if a == nil {
		return
	}
	return a.CreatedAtUnixSecs
}

func (a *AnalysisItemResponseModel) GetUpdatedAtUnixSecs() (value int) {
	if a == nil {
		return
	}
	return a.UpdatedAtUnixSecs
}

func (a *AnalysisItemResponseModel) GetCurrentVersionId() (value string) {
	if a == nil {
		return
	}
	return a.CurrentVersionId
}

func (a *AnalysisItemResponseModel) GetLegacySourceId() (value *string) {
	if a == nil {
		return
	}
	return a.LegacySourceId
}

func (a *AnalysisItemResponseModel) SetId(id string) {
	a.Id = id
}

func (a *AnalysisItemResponseModel) SetDefinition(definition *AnalysisItemResponseModelDefinition) {
	a.Definition = definition
}

func (a *AnalysisItemResponseModel) SetAccessInfo(accessInfo *ResourceAccessInfo) {
	a.AccessInfo = accessInfo
}

func (a *AnalysisItemResponseModel) SetCreatedAtUnixSecs(createdAtUnixSecs int) {
	a.CreatedAtUnixSecs = createdAtUnixSecs
}

func (a *AnalysisItemResponseModel) SetUpdatedAtUnixSecs(updatedAtUnixSecs int) {
	a.UpdatedAtUnixSecs = updatedAtUnixSecs
}

func (a *AnalysisItemResponseModel) SetCurrentVersionId(currentVersionId string) {
	a.CurrentVersionId = currentVersionId
}

func (a *AnalysisItemResponseModel) SetLegacySourceId(legacySourceId *string) {
	a.LegacySourceId = legacySourceId
}

func (a *AnalysisItemResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AnalysisItemResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AnalysisItemResponseModel to string"
	}
	return string(jsonData)
}

func (a *AnalysisItemResponseModel) UnmarshalJSON(data []byte) error {
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

func (a *AnalysisItemResponseModel) MarshalJSON() ([]byte, error) {
	type embed AnalysisItemResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
