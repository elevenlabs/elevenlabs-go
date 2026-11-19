package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetPronunciationDictionaryWithRulesResponseModel struct {
	// The ID of the pronunciation dictionary.
	Id string `json:"id" xml:"id" required:"true"`
	// The ID of the latest version of the pronunciation dictionary.
	LatestVersionId string `json:"latest_version_id" xml:"latest_version_id" required:"true"`
	// The number of rules in the latest version of the pronunciation dictionary.
	LatestVersionRulesNum int `json:"latest_version_rules_num" xml:"latest_version_rules_num" required:"true"`
	// The name of the pronunciation dictionary.
	Name string `json:"name" xml:"name" required:"true"`
	// The permission on the resource of the pronunciation dictionary.
	PermissionOnResource *GetPronunciationDictionaryWithRulesResponseModelPermissionOnResource `json:"permission_on_resource" xml:"permission_on_resource" required:"true"`
	// The user ID of the creator of the pronunciation dictionary.
	CreatedBy string `json:"created_by" xml:"created_by" required:"true"`
	// The creation time of the pronunciation dictionary in Unix timestamp.
	CreationTimeUnix int `json:"creation_time_unix" xml:"creation_time_unix" required:"true"`
	// The archive time of the pronunciation dictionary in Unix timestamp.
	ArchivedTimeUnix *int `json:"archived_time_unix,omitempty" xml:"archived_time_unix,omitempty"`
	// The description of the pronunciation dictionary.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The rules in the latest version of the pronunciation dictionary.
	Rules           []*GetPronunciationDictionaryWithRulesResponseModelRulesItem `json:"rules" xml:"rules" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetPronunciationDictionaryWithRulesResponseModel) GetId() (value string) {
	if g == nil {
		return
	}
	return g.Id
}

func (g *GetPronunciationDictionaryWithRulesResponseModel) GetLatestVersionId() (value string) {
	if g == nil {
		return
	}
	return g.LatestVersionId
}

func (g *GetPronunciationDictionaryWithRulesResponseModel) GetLatestVersionRulesNum() (value int) {
	if g == nil {
		return
	}
	return g.LatestVersionRulesNum
}

func (g *GetPronunciationDictionaryWithRulesResponseModel) GetName() (value string) {
	if g == nil {
		return
	}
	return g.Name
}

func (g *GetPronunciationDictionaryWithRulesResponseModel) GetPermissionOnResource() (value *GetPronunciationDictionaryWithRulesResponseModelPermissionOnResource) {
	if g == nil {
		return
	}
	return g.PermissionOnResource
}

func (g *GetPronunciationDictionaryWithRulesResponseModel) GetCreatedBy() (value string) {
	if g == nil {
		return
	}
	return g.CreatedBy
}

func (g *GetPronunciationDictionaryWithRulesResponseModel) GetCreationTimeUnix() (value int) {
	if g == nil {
		return
	}
	return g.CreationTimeUnix
}

func (g *GetPronunciationDictionaryWithRulesResponseModel) GetArchivedTimeUnix() (value *int) {
	if g == nil {
		return
	}
	return g.ArchivedTimeUnix
}

func (g *GetPronunciationDictionaryWithRulesResponseModel) GetDescription() (value *string) {
	if g == nil {
		return
	}
	return g.Description
}

func (g *GetPronunciationDictionaryWithRulesResponseModel) GetRules() (value []*GetPronunciationDictionaryWithRulesResponseModelRulesItem) {
	if g == nil {
		return
	}
	return g.Rules
}

func (g *GetPronunciationDictionaryWithRulesResponseModel) SetId(id string) {
	g.Id = id
}

func (g *GetPronunciationDictionaryWithRulesResponseModel) SetLatestVersionId(latestVersionId string) {
	g.LatestVersionId = latestVersionId
}

func (g *GetPronunciationDictionaryWithRulesResponseModel) SetLatestVersionRulesNum(latestVersionRulesNum int) {
	g.LatestVersionRulesNum = latestVersionRulesNum
}

func (g *GetPronunciationDictionaryWithRulesResponseModel) SetName(name string) {
	g.Name = name
}

func (g *GetPronunciationDictionaryWithRulesResponseModel) SetPermissionOnResource(permissionOnResource *GetPronunciationDictionaryWithRulesResponseModelPermissionOnResource) {
	g.PermissionOnResource = permissionOnResource
}

func (g *GetPronunciationDictionaryWithRulesResponseModel) SetCreatedBy(createdBy string) {
	g.CreatedBy = createdBy
}

func (g *GetPronunciationDictionaryWithRulesResponseModel) SetCreationTimeUnix(creationTimeUnix int) {
	g.CreationTimeUnix = creationTimeUnix
}

func (g *GetPronunciationDictionaryWithRulesResponseModel) SetArchivedTimeUnix(archivedTimeUnix *int) {
	g.ArchivedTimeUnix = archivedTimeUnix
}

func (g *GetPronunciationDictionaryWithRulesResponseModel) SetDescription(description *string) {
	g.Description = description
}

func (g *GetPronunciationDictionaryWithRulesResponseModel) SetRules(rules []*GetPronunciationDictionaryWithRulesResponseModelRulesItem) {
	g.Rules = rules
}

func (g *GetPronunciationDictionaryWithRulesResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetPronunciationDictionaryWithRulesResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetPronunciationDictionaryWithRulesResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetPronunciationDictionaryWithRulesResponseModel) UnmarshalJSON(data []byte) error {
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

func (g *GetPronunciationDictionaryWithRulesResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetPronunciationDictionaryWithRulesResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
