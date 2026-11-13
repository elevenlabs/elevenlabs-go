package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetKnowledgeBaseSummaryFolderResponseModel struct {
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
	DependentAgents []*GetKnowledgeBaseSummaryFolderResponseModelDependentAgentsItem `json:"dependent_agents" xml:"dependent_agents" required:"true"`
	type_           string
	ChildrenCount   int `json:"children_count" xml:"children_count" required:"true"`
	// Number of non-folder documents anywhere in this folder's subtree (recursive). Counting stops past 1000;
	DocumentCount    int                     `json:"document_count" xml:"document_count" required:"true"`
	AutoSyncInfo     *AutoSyncInfo           `json:"auto_sync_info,omitempty" xml:"auto_sync_info,omitempty"`
	ExternalSyncInfo *ExternalFolderSyncInfo `json:"external_sync_info,omitempty" xml:"external_sync_info,omitempty"`
	IsFrozen         *bool                   `json:"is_frozen,omitempty" xml:"is_frozen,omitempty"`
	extraProperties  map[string]interface{}
}

func (g *GetKnowledgeBaseSummaryFolderResponseModel) GetId() (value string) {
	if g == nil {
		return
	}
	return g.Id
}

func (g *GetKnowledgeBaseSummaryFolderResponseModel) GetName() (value string) {
	if g == nil {
		return
	}
	return g.Name
}

func (g *GetKnowledgeBaseSummaryFolderResponseModel) GetMetadata() (value *KnowledgeBaseDocumentMetadataResponseModel) {
	if g == nil {
		return
	}
	return g.Metadata
}

func (g *GetKnowledgeBaseSummaryFolderResponseModel) GetSupportedUsages() (value []DocumentUsageModeEnum) {
	if g == nil {
		return
	}
	return g.SupportedUsages
}

func (g *GetKnowledgeBaseSummaryFolderResponseModel) GetAccessInfo() (value *ResourceAccessInfo) {
	if g == nil {
		return
	}
	return g.AccessInfo
}

func (g *GetKnowledgeBaseSummaryFolderResponseModel) GetFolderParentId() (value *string) {
	if g == nil {
		return
	}
	return g.FolderParentId
}

func (g *GetKnowledgeBaseSummaryFolderResponseModel) GetFolderPath() (value []*KnowledgeBaseFolderPathSegmentSummaryResponseModel) {
	if g == nil {
		return
	}
	return g.FolderPath
}

func (g *GetKnowledgeBaseSummaryFolderResponseModel) GetDependentAgents() (value []*GetKnowledgeBaseSummaryFolderResponseModelDependentAgentsItem) {
	if g == nil {
		return
	}
	return g.DependentAgents
}

func (g *GetKnowledgeBaseSummaryFolderResponseModel) GetChildrenCount() (value int) {
	if g == nil {
		return
	}
	return g.ChildrenCount
}

func (g *GetKnowledgeBaseSummaryFolderResponseModel) GetDocumentCount() (value int) {
	if g == nil {
		return
	}
	return g.DocumentCount
}

func (g *GetKnowledgeBaseSummaryFolderResponseModel) GetAutoSyncInfo() (value *AutoSyncInfo) {
	if g == nil {
		return
	}
	return g.AutoSyncInfo
}

func (g *GetKnowledgeBaseSummaryFolderResponseModel) GetExternalSyncInfo() (value *ExternalFolderSyncInfo) {
	if g == nil {
		return
	}
	return g.ExternalSyncInfo
}

func (g *GetKnowledgeBaseSummaryFolderResponseModel) GetIsFrozen() (value *bool) {
	if g == nil {
		return
	}
	return g.IsFrozen
}

func (g *GetKnowledgeBaseSummaryFolderResponseModel) SetId(id string) {
	g.Id = id
}

func (g *GetKnowledgeBaseSummaryFolderResponseModel) SetName(name string) {
	g.Name = name
}

func (g *GetKnowledgeBaseSummaryFolderResponseModel) SetMetadata(metadata *KnowledgeBaseDocumentMetadataResponseModel) {
	g.Metadata = metadata
}

func (g *GetKnowledgeBaseSummaryFolderResponseModel) SetSupportedUsages(supportedUsages []DocumentUsageModeEnum) {
	g.SupportedUsages = supportedUsages
}

func (g *GetKnowledgeBaseSummaryFolderResponseModel) SetAccessInfo(accessInfo *ResourceAccessInfo) {
	g.AccessInfo = accessInfo
}

func (g *GetKnowledgeBaseSummaryFolderResponseModel) SetFolderParentId(folderParentId *string) {
	g.FolderParentId = folderParentId
}

func (g *GetKnowledgeBaseSummaryFolderResponseModel) SetFolderPath(folderPath []*KnowledgeBaseFolderPathSegmentSummaryResponseModel) {
	g.FolderPath = folderPath
}

func (g *GetKnowledgeBaseSummaryFolderResponseModel) SetDependentAgents(dependentAgents []*GetKnowledgeBaseSummaryFolderResponseModelDependentAgentsItem) {
	g.DependentAgents = dependentAgents
}

func (g *GetKnowledgeBaseSummaryFolderResponseModel) SetChildrenCount(childrenCount int) {
	g.ChildrenCount = childrenCount
}

func (g *GetKnowledgeBaseSummaryFolderResponseModel) SetDocumentCount(documentCount int) {
	g.DocumentCount = documentCount
}

func (g *GetKnowledgeBaseSummaryFolderResponseModel) SetAutoSyncInfo(autoSyncInfo *AutoSyncInfo) {
	g.AutoSyncInfo = autoSyncInfo
}

func (g *GetKnowledgeBaseSummaryFolderResponseModel) SetExternalSyncInfo(externalSyncInfo *ExternalFolderSyncInfo) {
	g.ExternalSyncInfo = externalSyncInfo
}

func (g *GetKnowledgeBaseSummaryFolderResponseModel) SetIsFrozen(isFrozen *bool) {
	g.IsFrozen = isFrozen
}

func (g *GetKnowledgeBaseSummaryFolderResponseModel) Type() string {
	return g.type_
}

func (g *GetKnowledgeBaseSummaryFolderResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetKnowledgeBaseSummaryFolderResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetKnowledgeBaseSummaryFolderResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetKnowledgeBaseSummaryFolderResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, g); err != nil {
		return err
	}
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "folder" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", g, "folder", *literals.Type)
	}
	g.type_ = "folder"
	extraProperties, err := internal.ExtractExtraProperties(data, *g, "type")
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetKnowledgeBaseSummaryFolderResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetKnowledgeBaseSummaryFolderResponseModel
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*g),
		Type:  "folder",
	}
	return json.Marshal(marshaler)
}
