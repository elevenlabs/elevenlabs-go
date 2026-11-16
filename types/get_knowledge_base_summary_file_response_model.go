package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetKnowledgeBaseSummaryFileResponseModel struct {
	Id              string                                      `json:"id" xml:"id" required:"true"`
	Name            string                                      `json:"name" xml:"name" required:"true"`
	Metadata        *KnowledgeBaseDocumentMetadataResponseModel `json:"metadata" xml:"metadata" required:"true"`
	SupportedUsages []DocumentUsageModeEnum                     `json:"supported_usages" xml:"supported_usages" required:"true"`
	AccessInfo      *ResourceAccessInfo                         `json:"access_info" xml:"access_info" required:"true"`
	// The ID of the parent folder, or null if the document is at the root level.
	FolderParentId *string `json:"folder_parent_id,omitempty" xml:"folder_parent_id,omitempty"`
	// Per-entity RAG chunking override. Null means inherit from ancestors or the embedding model default.
	ChunkingOverride *ChunkingOverride `json:"chunking_override,omitempty" xml:"chunking_override,omitempty"`
	// The folder path segments leading to this entity, from root to parent folder.
	FolderPath []*KnowledgeBaseFolderPathSegmentSummaryResponseModel `json:"folder_path,omitempty" xml:"folder_path,omitempty"`
	// This field is deprecated, always empty, and will be removed in the future. Use the separate endpoint to get dependent agents instead.
	DependentAgents  []*GetKnowledgeBaseSummaryFileResponseModelDependentAgentsItem `json:"dependent_agents" xml:"dependent_agents" required:"true"`
	type_            string
	ExternalSyncInfo *ExternalFileSyncInfo `json:"external_sync_info,omitempty" xml:"external_sync_info,omitempty"`
	AutoSyncInfo     *AutoSyncInfo         `json:"auto_sync_info,omitempty" xml:"auto_sync_info,omitempty"`
	RefreshStatus    *FileRefreshStatus    `json:"refresh_status,omitempty" xml:"refresh_status,omitempty"`
	IsFrozen         *bool                 `json:"is_frozen,omitempty" xml:"is_frozen,omitempty"`
	extraProperties  map[string]interface{}
}

func (g *GetKnowledgeBaseSummaryFileResponseModel) GetId() (value string) {
	if g == nil {
		return
	}
	return g.Id
}

func (g *GetKnowledgeBaseSummaryFileResponseModel) GetName() (value string) {
	if g == nil {
		return
	}
	return g.Name
}

func (g *GetKnowledgeBaseSummaryFileResponseModel) GetMetadata() (value *KnowledgeBaseDocumentMetadataResponseModel) {
	if g == nil {
		return
	}
	return g.Metadata
}

func (g *GetKnowledgeBaseSummaryFileResponseModel) GetSupportedUsages() (value []DocumentUsageModeEnum) {
	if g == nil {
		return
	}
	return g.SupportedUsages
}

func (g *GetKnowledgeBaseSummaryFileResponseModel) GetAccessInfo() (value *ResourceAccessInfo) {
	if g == nil {
		return
	}
	return g.AccessInfo
}

func (g *GetKnowledgeBaseSummaryFileResponseModel) GetFolderParentId() (value *string) {
	if g == nil {
		return
	}
	return g.FolderParentId
}

func (g *GetKnowledgeBaseSummaryFileResponseModel) GetChunkingOverride() (value *ChunkingOverride) {
	if g == nil {
		return
	}
	return g.ChunkingOverride
}

func (g *GetKnowledgeBaseSummaryFileResponseModel) GetFolderPath() (value []*KnowledgeBaseFolderPathSegmentSummaryResponseModel) {
	if g == nil {
		return
	}
	return g.FolderPath
}

func (g *GetKnowledgeBaseSummaryFileResponseModel) GetDependentAgents() (value []*GetKnowledgeBaseSummaryFileResponseModelDependentAgentsItem) {
	if g == nil {
		return
	}
	return g.DependentAgents
}

func (g *GetKnowledgeBaseSummaryFileResponseModel) GetExternalSyncInfo() (value *ExternalFileSyncInfo) {
	if g == nil {
		return
	}
	return g.ExternalSyncInfo
}

func (g *GetKnowledgeBaseSummaryFileResponseModel) GetAutoSyncInfo() (value *AutoSyncInfo) {
	if g == nil {
		return
	}
	return g.AutoSyncInfo
}

func (g *GetKnowledgeBaseSummaryFileResponseModel) GetRefreshStatus() (value *FileRefreshStatus) {
	if g == nil {
		return
	}
	return g.RefreshStatus
}

func (g *GetKnowledgeBaseSummaryFileResponseModel) GetIsFrozen() (value *bool) {
	if g == nil {
		return
	}
	return g.IsFrozen
}

func (g *GetKnowledgeBaseSummaryFileResponseModel) SetId(id string) {
	g.Id = id
}

func (g *GetKnowledgeBaseSummaryFileResponseModel) SetName(name string) {
	g.Name = name
}

func (g *GetKnowledgeBaseSummaryFileResponseModel) SetMetadata(metadata *KnowledgeBaseDocumentMetadataResponseModel) {
	g.Metadata = metadata
}

func (g *GetKnowledgeBaseSummaryFileResponseModel) SetSupportedUsages(supportedUsages []DocumentUsageModeEnum) {
	g.SupportedUsages = supportedUsages
}

func (g *GetKnowledgeBaseSummaryFileResponseModel) SetAccessInfo(accessInfo *ResourceAccessInfo) {
	g.AccessInfo = accessInfo
}

func (g *GetKnowledgeBaseSummaryFileResponseModel) SetFolderParentId(folderParentId *string) {
	g.FolderParentId = folderParentId
}

func (g *GetKnowledgeBaseSummaryFileResponseModel) SetChunkingOverride(chunkingOverride *ChunkingOverride) {
	g.ChunkingOverride = chunkingOverride
}

func (g *GetKnowledgeBaseSummaryFileResponseModel) SetFolderPath(folderPath []*KnowledgeBaseFolderPathSegmentSummaryResponseModel) {
	g.FolderPath = folderPath
}

func (g *GetKnowledgeBaseSummaryFileResponseModel) SetDependentAgents(dependentAgents []*GetKnowledgeBaseSummaryFileResponseModelDependentAgentsItem) {
	g.DependentAgents = dependentAgents
}

func (g *GetKnowledgeBaseSummaryFileResponseModel) SetExternalSyncInfo(externalSyncInfo *ExternalFileSyncInfo) {
	g.ExternalSyncInfo = externalSyncInfo
}

func (g *GetKnowledgeBaseSummaryFileResponseModel) SetAutoSyncInfo(autoSyncInfo *AutoSyncInfo) {
	g.AutoSyncInfo = autoSyncInfo
}

func (g *GetKnowledgeBaseSummaryFileResponseModel) SetRefreshStatus(refreshStatus *FileRefreshStatus) {
	g.RefreshStatus = refreshStatus
}

func (g *GetKnowledgeBaseSummaryFileResponseModel) SetIsFrozen(isFrozen *bool) {
	g.IsFrozen = isFrozen
}

func (g *GetKnowledgeBaseSummaryFileResponseModel) Type() string {
	return g.type_
}

func (g *GetKnowledgeBaseSummaryFileResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetKnowledgeBaseSummaryFileResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetKnowledgeBaseSummaryFileResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetKnowledgeBaseSummaryFileResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, g); err != nil {
		return err
	}
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "file" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", g, "file", *literals.Type)
	}
	g.type_ = "file"
	extraProperties, err := internal.ExtractExtraProperties(data, *g, "type")
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetKnowledgeBaseSummaryFileResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetKnowledgeBaseSummaryFileResponseModel
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*g),
		Type:  "file",
	}
	return json.Marshal(marshaler)
}
