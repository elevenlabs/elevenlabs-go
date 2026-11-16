package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetPronunciationDictionaryMetadataResponse struct {
	// The ID of the pronunciation dictionary.
	Id string `json:"id" xml:"id" required:"true"`
	// The ID of the latest version of the pronunciation dictionary.
	LatestVersionId string `json:"latest_version_id" xml:"latest_version_id" required:"true"`
	// The number of rules in the latest version of the pronunciation dictionary.
	LatestVersionRulesNum int `json:"latest_version_rules_num" xml:"latest_version_rules_num" required:"true"`
	// The name of the pronunciation dictionary.
	Name string `json:"name" xml:"name" required:"true"`
	// The permission on the resource of the pronunciation dictionary.
	PermissionOnResource *GetPronunciationDictionaryMetadataResponsePermissionOnResource `json:"permission_on_resource" xml:"permission_on_resource" required:"true"`
	// The user ID of the creator of the pronunciation dictionary.
	CreatedBy string `json:"created_by" xml:"created_by" required:"true"`
	// The creation time of the pronunciation dictionary in Unix timestamp.
	CreationTimeUnix int `json:"creation_time_unix" xml:"creation_time_unix" required:"true"`
	// The archive time of the pronunciation dictionary in Unix timestamp.
	ArchivedTimeUnix *int `json:"archived_time_unix,omitempty" xml:"archived_time_unix,omitempty"`
	// The description of the pronunciation dictionary.
	Description     *string `json:"description,omitempty" xml:"description,omitempty"`
	extraProperties map[string]interface{}
}

func (g *GetPronunciationDictionaryMetadataResponse) GetId() (value string) {
	if g == nil {
		return
	}
	return g.Id
}

func (g *GetPronunciationDictionaryMetadataResponse) GetLatestVersionId() (value string) {
	if g == nil {
		return
	}
	return g.LatestVersionId
}

func (g *GetPronunciationDictionaryMetadataResponse) GetLatestVersionRulesNum() (value int) {
	if g == nil {
		return
	}
	return g.LatestVersionRulesNum
}

func (g *GetPronunciationDictionaryMetadataResponse) GetName() (value string) {
	if g == nil {
		return
	}
	return g.Name
}

func (g *GetPronunciationDictionaryMetadataResponse) GetPermissionOnResource() (value *GetPronunciationDictionaryMetadataResponsePermissionOnResource) {
	if g == nil {
		return
	}
	return g.PermissionOnResource
}

func (g *GetPronunciationDictionaryMetadataResponse) GetCreatedBy() (value string) {
	if g == nil {
		return
	}
	return g.CreatedBy
}

func (g *GetPronunciationDictionaryMetadataResponse) GetCreationTimeUnix() (value int) {
	if g == nil {
		return
	}
	return g.CreationTimeUnix
}

func (g *GetPronunciationDictionaryMetadataResponse) GetArchivedTimeUnix() (value *int) {
	if g == nil {
		return
	}
	return g.ArchivedTimeUnix
}

func (g *GetPronunciationDictionaryMetadataResponse) GetDescription() (value *string) {
	if g == nil {
		return
	}
	return g.Description
}

func (g *GetPronunciationDictionaryMetadataResponse) SetId(id string) {
	g.Id = id
}

func (g *GetPronunciationDictionaryMetadataResponse) SetLatestVersionId(latestVersionId string) {
	g.LatestVersionId = latestVersionId
}

func (g *GetPronunciationDictionaryMetadataResponse) SetLatestVersionRulesNum(latestVersionRulesNum int) {
	g.LatestVersionRulesNum = latestVersionRulesNum
}

func (g *GetPronunciationDictionaryMetadataResponse) SetName(name string) {
	g.Name = name
}

func (g *GetPronunciationDictionaryMetadataResponse) SetPermissionOnResource(permissionOnResource *GetPronunciationDictionaryMetadataResponsePermissionOnResource) {
	g.PermissionOnResource = permissionOnResource
}

func (g *GetPronunciationDictionaryMetadataResponse) SetCreatedBy(createdBy string) {
	g.CreatedBy = createdBy
}

func (g *GetPronunciationDictionaryMetadataResponse) SetCreationTimeUnix(creationTimeUnix int) {
	g.CreationTimeUnix = creationTimeUnix
}

func (g *GetPronunciationDictionaryMetadataResponse) SetArchivedTimeUnix(archivedTimeUnix *int) {
	g.ArchivedTimeUnix = archivedTimeUnix
}

func (g *GetPronunciationDictionaryMetadataResponse) SetDescription(description *string) {
	g.Description = description
}

func (g *GetPronunciationDictionaryMetadataResponse) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetPronunciationDictionaryMetadataResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetPronunciationDictionaryMetadataResponse to string"
	}
	return string(jsonData)
}

func (g *GetPronunciationDictionaryMetadataResponse) UnmarshalJSON(data []byte) error {
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

func (g *GetPronunciationDictionaryMetadataResponse) MarshalJSON() ([]byte, error) {
	type embed GetPronunciationDictionaryMetadataResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
