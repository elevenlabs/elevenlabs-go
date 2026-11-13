package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetKnowledgeBaseSummaryUrlResponseModel struct {
	Id              string                                      `json:"id" xml:"id" required:"true"`
	Name            string                                      `json:"name" xml:"name" required:"true"`
	Metadata        *KnowledgeBaseDocumentMetadataResponseModel `json:"metadata" xml:"metadata" required:"true"`
	SupportedUsages []DocumentUsageModeEnum                     `json:"supported_usages" xml:"supported_usages" required:"true"`
	AccessInfo      *ResourceAccessInfo                         `json:"access_info" xml:"access_info" required:"true"`
	// The ID of the parent folder, or null if the document is at the root level.
	FolderParentId *string `json:"folder_parent_id,omitempty" xml:"folder_parent_id,omitempty"`
	// The folder path segments leading to this entity, from root to parent folder.
	FolderPath []*KnowledgeBaseFolderPathSegmentSummaryResponseModel `json:"folder_path,omitempty" xml:"folder_path,omitempty"`
	// This field is deprecated and will be removed in the future, use the separate endpoint to get dependent agents instead.
	DependentAgents []*GetKnowledgeBaseSummaryUrlResponseModelDependentAgentsItem `json:"dependent_agents" xml:"dependent_agents" required:"true"`
	type_           string
	Url             string        `json:"url" xml:"url" required:"true"`
	AutoSyncInfo    *AutoSyncInfo `json:"auto_sync_info,omitempty" xml:"auto_sync_info,omitempty"`
	extraProperties map[string]interface{}
}

func (g *GetKnowledgeBaseSummaryUrlResponseModel) GetId() (value string) {
	if g == nil {
		return
	}
	return g.Id
}

func (g *GetKnowledgeBaseSummaryUrlResponseModel) GetName() (value string) {
	if g == nil {
		return
	}
	return g.Name
}

func (g *GetKnowledgeBaseSummaryUrlResponseModel) GetMetadata() (value *KnowledgeBaseDocumentMetadataResponseModel) {
	if g == nil {
		return
	}
	return g.Metadata
}

func (g *GetKnowledgeBaseSummaryUrlResponseModel) GetSupportedUsages() (value []DocumentUsageModeEnum) {
	if g == nil {
		return
	}
	return g.SupportedUsages
}

func (g *GetKnowledgeBaseSummaryUrlResponseModel) GetAccessInfo() (value *ResourceAccessInfo) {
	if g == nil {
		return
	}
	return g.AccessInfo
}

func (g *GetKnowledgeBaseSummaryUrlResponseModel) GetFolderParentId() (value *string) {
	if g == nil {
		return
	}
	return g.FolderParentId
}

func (g *GetKnowledgeBaseSummaryUrlResponseModel) GetFolderPath() (value []*KnowledgeBaseFolderPathSegmentSummaryResponseModel) {
	if g == nil {
		return
	}
	return g.FolderPath
}

func (g *GetKnowledgeBaseSummaryUrlResponseModel) GetDependentAgents() (value []*GetKnowledgeBaseSummaryUrlResponseModelDependentAgentsItem) {
	if g == nil {
		return
	}
	return g.DependentAgents
}

func (g *GetKnowledgeBaseSummaryUrlResponseModel) GetUrl() (value string) {
	if g == nil {
		return
	}
	return g.Url
}

func (g *GetKnowledgeBaseSummaryUrlResponseModel) GetAutoSyncInfo() (value *AutoSyncInfo) {
	if g == nil {
		return
	}
	return g.AutoSyncInfo
}

func (g *GetKnowledgeBaseSummaryUrlResponseModel) SetId(id string) {
	g.Id = id
}

func (g *GetKnowledgeBaseSummaryUrlResponseModel) SetName(name string) {
	g.Name = name
}

func (g *GetKnowledgeBaseSummaryUrlResponseModel) SetMetadata(metadata *KnowledgeBaseDocumentMetadataResponseModel) {
	g.Metadata = metadata
}

func (g *GetKnowledgeBaseSummaryUrlResponseModel) SetSupportedUsages(supportedUsages []DocumentUsageModeEnum) {
	g.SupportedUsages = supportedUsages
}

func (g *GetKnowledgeBaseSummaryUrlResponseModel) SetAccessInfo(accessInfo *ResourceAccessInfo) {
	g.AccessInfo = accessInfo
}

func (g *GetKnowledgeBaseSummaryUrlResponseModel) SetFolderParentId(folderParentId *string) {
	g.FolderParentId = folderParentId
}

func (g *GetKnowledgeBaseSummaryUrlResponseModel) SetFolderPath(folderPath []*KnowledgeBaseFolderPathSegmentSummaryResponseModel) {
	g.FolderPath = folderPath
}

func (g *GetKnowledgeBaseSummaryUrlResponseModel) SetDependentAgents(dependentAgents []*GetKnowledgeBaseSummaryUrlResponseModelDependentAgentsItem) {
	g.DependentAgents = dependentAgents
}

func (g *GetKnowledgeBaseSummaryUrlResponseModel) SetUrl(url string) {
	g.Url = url
}

func (g *GetKnowledgeBaseSummaryUrlResponseModel) SetAutoSyncInfo(autoSyncInfo *AutoSyncInfo) {
	g.AutoSyncInfo = autoSyncInfo
}

func (g *GetKnowledgeBaseSummaryUrlResponseModel) Type() string {
	return g.type_
}

func (g *GetKnowledgeBaseSummaryUrlResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetKnowledgeBaseSummaryUrlResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetKnowledgeBaseSummaryUrlResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetKnowledgeBaseSummaryUrlResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, g); err != nil {
		return err
	}
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "url" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", g, "url", *literals.Type)
	}
	g.type_ = "url"
	extraProperties, err := internal.ExtractExtraProperties(data, *g, "type")
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetKnowledgeBaseSummaryUrlResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetKnowledgeBaseSummaryUrlResponseModel
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*g),
		Type:  "url",
	}
	return json.Marshal(marshaler)
}
