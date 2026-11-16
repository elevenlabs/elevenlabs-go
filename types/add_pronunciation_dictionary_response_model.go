package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AddPronunciationDictionaryResponseModel struct {
	// The ID of the created pronunciation dictionary.
	Id string `json:"id" xml:"id" required:"true"`
	// The name of the created pronunciation dictionary.
	Name string `json:"name" xml:"name" required:"true"`
	// The user ID of the creator of the pronunciation dictionary.
	CreatedBy string `json:"created_by" xml:"created_by" required:"true"`
	// The creation time of the pronunciation dictionary in Unix timestamp.
	CreationTimeUnix int `json:"creation_time_unix" xml:"creation_time_unix" required:"true"`
	// The ID of the created pronunciation dictionary version.
	VersionId string `json:"version_id" xml:"version_id" required:"true"`
	// The number of rules in the version of the pronunciation dictionary.
	VersionRulesNum int `json:"version_rules_num" xml:"version_rules_num" required:"true"`
	// The description of the pronunciation dictionary.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The permission on the resource of the pronunciation dictionary.
	PermissionOnResource *AddPronunciationDictionaryResponseModelPermissionOnResource `json:"permission_on_resource" xml:"permission_on_resource" required:"true"`
	extraProperties      map[string]interface{}
}

func (a *AddPronunciationDictionaryResponseModel) GetId() (value string) {
	if a == nil {
		return
	}
	return a.Id
}

func (a *AddPronunciationDictionaryResponseModel) GetName() (value string) {
	if a == nil {
		return
	}
	return a.Name
}

func (a *AddPronunciationDictionaryResponseModel) GetCreatedBy() (value string) {
	if a == nil {
		return
	}
	return a.CreatedBy
}

func (a *AddPronunciationDictionaryResponseModel) GetCreationTimeUnix() (value int) {
	if a == nil {
		return
	}
	return a.CreationTimeUnix
}

func (a *AddPronunciationDictionaryResponseModel) GetVersionId() (value string) {
	if a == nil {
		return
	}
	return a.VersionId
}

func (a *AddPronunciationDictionaryResponseModel) GetVersionRulesNum() (value int) {
	if a == nil {
		return
	}
	return a.VersionRulesNum
}

func (a *AddPronunciationDictionaryResponseModel) GetDescription() (value *string) {
	if a == nil {
		return
	}
	return a.Description
}

func (a *AddPronunciationDictionaryResponseModel) GetPermissionOnResource() (value *AddPronunciationDictionaryResponseModelPermissionOnResource) {
	if a == nil {
		return
	}
	return a.PermissionOnResource
}

func (a *AddPronunciationDictionaryResponseModel) SetId(id string) {
	a.Id = id
}

func (a *AddPronunciationDictionaryResponseModel) SetName(name string) {
	a.Name = name
}

func (a *AddPronunciationDictionaryResponseModel) SetCreatedBy(createdBy string) {
	a.CreatedBy = createdBy
}

func (a *AddPronunciationDictionaryResponseModel) SetCreationTimeUnix(creationTimeUnix int) {
	a.CreationTimeUnix = creationTimeUnix
}

func (a *AddPronunciationDictionaryResponseModel) SetVersionId(versionId string) {
	a.VersionId = versionId
}

func (a *AddPronunciationDictionaryResponseModel) SetVersionRulesNum(versionRulesNum int) {
	a.VersionRulesNum = versionRulesNum
}

func (a *AddPronunciationDictionaryResponseModel) SetDescription(description *string) {
	a.Description = description
}

func (a *AddPronunciationDictionaryResponseModel) SetPermissionOnResource(permissionOnResource *AddPronunciationDictionaryResponseModelPermissionOnResource) {
	a.PermissionOnResource = permissionOnResource
}

func (a *AddPronunciationDictionaryResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AddPronunciationDictionaryResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddPronunciationDictionaryResponseModel to string"
	}
	return string(jsonData)
}

func (a *AddPronunciationDictionaryResponseModel) UnmarshalJSON(data []byte) error {
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

func (a *AddPronunciationDictionaryResponseModel) MarshalJSON() ([]byte, error) {
	type embed AddPronunciationDictionaryResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
