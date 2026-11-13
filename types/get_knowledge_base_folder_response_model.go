package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetKnowledgeBaseFolderResponseModel struct {
	Id              string                                      `json:"id" xml:"id" required:"true"`
	Name            string                                      `json:"name" xml:"name" required:"true"`
	Metadata        *KnowledgeBaseDocumentMetadataResponseModel `json:"metadata" xml:"metadata" required:"true"`
	SupportedUsages []DocumentUsageModeEnum                     `json:"supported_usages" xml:"supported_usages" required:"true"`
	AccessInfo      *ResourceAccessInfo                         `json:"access_info" xml:"access_info" required:"true"`
	// The ID of the parent folder, or null if the document is at the root level.
	FolderParentId *string `json:"folder_parent_id,omitempty" xml:"folder_parent_id,omitempty"`
	// The folder path segments leading to this entity, from root to parent folder.
	FolderPath    []*KnowledgeBaseFolderPathSegmentResponseModel `json:"folder_path,omitempty" xml:"folder_path,omitempty"`
	type_         string
	ChildrenCount int `json:"children_count" xml:"children_count" required:"true"`
	// Number of non-folder documents anywhere in this folder's subtree (recursive). Counting stops past 1000;
	DocumentCount    int                     `json:"document_count" xml:"document_count" required:"true"`
	AutoSyncInfo     *AutoSyncInfo           `json:"auto_sync_info,omitempty" xml:"auto_sync_info,omitempty"`
	ExternalSyncInfo *ExternalFolderSyncInfo `json:"external_sync_info,omitempty" xml:"external_sync_info,omitempty"`
	IsFrozen         *bool                   `json:"is_frozen,omitempty" xml:"is_frozen,omitempty"`
	// Most recent (in-flight or terminal) external sync job for this folder, if any. Used by clients to render sync progress.
	ActiveSyncJob   *KbExternalSyncJob `json:"active_sync_job,omitempty" xml:"active_sync_job,omitempty"`
	extraProperties map[string]interface{}
}

func (g *GetKnowledgeBaseFolderResponseModel) GetId() (value string) {
	if g == nil {
		return
	}
	return g.Id
}

func (g *GetKnowledgeBaseFolderResponseModel) GetName() (value string) {
	if g == nil {
		return
	}
	return g.Name
}

func (g *GetKnowledgeBaseFolderResponseModel) GetMetadata() (value *KnowledgeBaseDocumentMetadataResponseModel) {
	if g == nil {
		return
	}
	return g.Metadata
}

func (g *GetKnowledgeBaseFolderResponseModel) GetSupportedUsages() (value []DocumentUsageModeEnum) {
	if g == nil {
		return
	}
	return g.SupportedUsages
}

func (g *GetKnowledgeBaseFolderResponseModel) GetAccessInfo() (value *ResourceAccessInfo) {
	if g == nil {
		return
	}
	return g.AccessInfo
}

func (g *GetKnowledgeBaseFolderResponseModel) GetFolderParentId() (value *string) {
	if g == nil {
		return
	}
	return g.FolderParentId
}

func (g *GetKnowledgeBaseFolderResponseModel) GetFolderPath() (value []*KnowledgeBaseFolderPathSegmentResponseModel) {
	if g == nil {
		return
	}
	return g.FolderPath
}

func (g *GetKnowledgeBaseFolderResponseModel) GetChildrenCount() (value int) {
	if g == nil {
		return
	}
	return g.ChildrenCount
}

func (g *GetKnowledgeBaseFolderResponseModel) GetDocumentCount() (value int) {
	if g == nil {
		return
	}
	return g.DocumentCount
}

func (g *GetKnowledgeBaseFolderResponseModel) GetAutoSyncInfo() (value *AutoSyncInfo) {
	if g == nil {
		return
	}
	return g.AutoSyncInfo
}

func (g *GetKnowledgeBaseFolderResponseModel) GetExternalSyncInfo() (value *ExternalFolderSyncInfo) {
	if g == nil {
		return
	}
	return g.ExternalSyncInfo
}

func (g *GetKnowledgeBaseFolderResponseModel) GetIsFrozen() (value *bool) {
	if g == nil {
		return
	}
	return g.IsFrozen
}

func (g *GetKnowledgeBaseFolderResponseModel) GetActiveSyncJob() (value *KbExternalSyncJob) {
	if g == nil {
		return
	}
	return g.ActiveSyncJob
}

func (g *GetKnowledgeBaseFolderResponseModel) SetId(id string) {
	g.Id = id
}

func (g *GetKnowledgeBaseFolderResponseModel) SetName(name string) {
	g.Name = name
}

func (g *GetKnowledgeBaseFolderResponseModel) SetMetadata(metadata *KnowledgeBaseDocumentMetadataResponseModel) {
	g.Metadata = metadata
}

func (g *GetKnowledgeBaseFolderResponseModel) SetSupportedUsages(supportedUsages []DocumentUsageModeEnum) {
	g.SupportedUsages = supportedUsages
}

func (g *GetKnowledgeBaseFolderResponseModel) SetAccessInfo(accessInfo *ResourceAccessInfo) {
	g.AccessInfo = accessInfo
}

func (g *GetKnowledgeBaseFolderResponseModel) SetFolderParentId(folderParentId *string) {
	g.FolderParentId = folderParentId
}

func (g *GetKnowledgeBaseFolderResponseModel) SetFolderPath(folderPath []*KnowledgeBaseFolderPathSegmentResponseModel) {
	g.FolderPath = folderPath
}

func (g *GetKnowledgeBaseFolderResponseModel) SetChildrenCount(childrenCount int) {
	g.ChildrenCount = childrenCount
}

func (g *GetKnowledgeBaseFolderResponseModel) SetDocumentCount(documentCount int) {
	g.DocumentCount = documentCount
}

func (g *GetKnowledgeBaseFolderResponseModel) SetAutoSyncInfo(autoSyncInfo *AutoSyncInfo) {
	g.AutoSyncInfo = autoSyncInfo
}

func (g *GetKnowledgeBaseFolderResponseModel) SetExternalSyncInfo(externalSyncInfo *ExternalFolderSyncInfo) {
	g.ExternalSyncInfo = externalSyncInfo
}

func (g *GetKnowledgeBaseFolderResponseModel) SetIsFrozen(isFrozen *bool) {
	g.IsFrozen = isFrozen
}

func (g *GetKnowledgeBaseFolderResponseModel) SetActiveSyncJob(activeSyncJob *KbExternalSyncJob) {
	g.ActiveSyncJob = activeSyncJob
}

func (g *GetKnowledgeBaseFolderResponseModel) Type() string {
	return g.type_
}

func (g *GetKnowledgeBaseFolderResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetKnowledgeBaseFolderResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetKnowledgeBaseFolderResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetKnowledgeBaseFolderResponseModel) UnmarshalJSON(data []byte) error {
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

func (g *GetKnowledgeBaseFolderResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetKnowledgeBaseFolderResponseModel
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*g),
		Type:  "folder",
	}
	return json.Marshal(marshaler)
}
