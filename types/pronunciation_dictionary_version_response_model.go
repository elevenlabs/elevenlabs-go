package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PronunciationDictionaryVersionResponseModel struct {
	VersionId                 string                                                           `json:"version_id" xml:"version_id" required:"true"`
	VersionRulesNum           int                                                              `json:"version_rules_num" xml:"version_rules_num" required:"true"`
	PronunciationDictionaryId string                                                           `json:"pronunciation_dictionary_id" xml:"pronunciation_dictionary_id" required:"true"`
	DictionaryName            string                                                           `json:"dictionary_name" xml:"dictionary_name" required:"true"`
	VersionName               string                                                           `json:"version_name" xml:"version_name" required:"true"`
	PermissionOnResource      *PronunciationDictionaryVersionResponseModelPermissionOnResource `json:"permission_on_resource" xml:"permission_on_resource" required:"true"`
	CreatedBy                 string                                                           `json:"created_by" xml:"created_by" required:"true"`
	CreationTimeUnix          int                                                              `json:"creation_time_unix" xml:"creation_time_unix" required:"true"`
	ArchivedTimeUnix          *int                                                             `json:"archived_time_unix,omitempty" xml:"archived_time_unix,omitempty"`
	extraProperties           map[string]interface{}
}

func (p *PronunciationDictionaryVersionResponseModel) GetVersionId() (value string) {
	if p == nil {
		return
	}
	return p.VersionId
}

func (p *PronunciationDictionaryVersionResponseModel) GetVersionRulesNum() (value int) {
	if p == nil {
		return
	}
	return p.VersionRulesNum
}

func (p *PronunciationDictionaryVersionResponseModel) GetPronunciationDictionaryId() (value string) {
	if p == nil {
		return
	}
	return p.PronunciationDictionaryId
}

func (p *PronunciationDictionaryVersionResponseModel) GetDictionaryName() (value string) {
	if p == nil {
		return
	}
	return p.DictionaryName
}

func (p *PronunciationDictionaryVersionResponseModel) GetVersionName() (value string) {
	if p == nil {
		return
	}
	return p.VersionName
}

func (p *PronunciationDictionaryVersionResponseModel) GetPermissionOnResource() (value *PronunciationDictionaryVersionResponseModelPermissionOnResource) {
	if p == nil {
		return
	}
	return p.PermissionOnResource
}

func (p *PronunciationDictionaryVersionResponseModel) GetCreatedBy() (value string) {
	if p == nil {
		return
	}
	return p.CreatedBy
}

func (p *PronunciationDictionaryVersionResponseModel) GetCreationTimeUnix() (value int) {
	if p == nil {
		return
	}
	return p.CreationTimeUnix
}

func (p *PronunciationDictionaryVersionResponseModel) GetArchivedTimeUnix() (value *int) {
	if p == nil {
		return
	}
	return p.ArchivedTimeUnix
}

func (p *PronunciationDictionaryVersionResponseModel) SetVersionId(versionId string) {
	p.VersionId = versionId
}

func (p *PronunciationDictionaryVersionResponseModel) SetVersionRulesNum(versionRulesNum int) {
	p.VersionRulesNum = versionRulesNum
}

func (p *PronunciationDictionaryVersionResponseModel) SetPronunciationDictionaryId(pronunciationDictionaryId string) {
	p.PronunciationDictionaryId = pronunciationDictionaryId
}

func (p *PronunciationDictionaryVersionResponseModel) SetDictionaryName(dictionaryName string) {
	p.DictionaryName = dictionaryName
}

func (p *PronunciationDictionaryVersionResponseModel) SetVersionName(versionName string) {
	p.VersionName = versionName
}

func (p *PronunciationDictionaryVersionResponseModel) SetPermissionOnResource(permissionOnResource *PronunciationDictionaryVersionResponseModelPermissionOnResource) {
	p.PermissionOnResource = permissionOnResource
}

func (p *PronunciationDictionaryVersionResponseModel) SetCreatedBy(createdBy string) {
	p.CreatedBy = createdBy
}

func (p *PronunciationDictionaryVersionResponseModel) SetCreationTimeUnix(creationTimeUnix int) {
	p.CreationTimeUnix = creationTimeUnix
}

func (p *PronunciationDictionaryVersionResponseModel) SetArchivedTimeUnix(archivedTimeUnix *int) {
	p.ArchivedTimeUnix = archivedTimeUnix
}

func (p *PronunciationDictionaryVersionResponseModel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PronunciationDictionaryVersionResponseModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PronunciationDictionaryVersionResponseModel to string"
	}
	return string(jsonData)
}

func (p *PronunciationDictionaryVersionResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, p); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PronunciationDictionaryVersionResponseModel) MarshalJSON() ([]byte, error) {
	type embed PronunciationDictionaryVersionResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
