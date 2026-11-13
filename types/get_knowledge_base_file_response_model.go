package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetKnowledgeBaseFileResponseModel struct {
	Id              string                                      `json:"id" xml:"id" required:"true"`
	Name            string                                      `json:"name" xml:"name" required:"true"`
	Metadata        *KnowledgeBaseDocumentMetadataResponseModel `json:"metadata" xml:"metadata" required:"true"`
	SupportedUsages []DocumentUsageModeEnum                     `json:"supported_usages" xml:"supported_usages" required:"true"`
	AccessInfo      *ResourceAccessInfo                         `json:"access_info" xml:"access_info" required:"true"`
	// The ID of the parent folder, or null if the document is at the root level.
	FolderParentId *string `json:"folder_parent_id,omitempty" xml:"folder_parent_id,omitempty"`
	// The folder path segments leading to this entity, from root to parent folder.
	FolderPath         []*KnowledgeBaseFolderPathSegmentResponseModel `json:"folder_path,omitempty" xml:"folder_path,omitempty"`
	type_              string
	ExtractedInnerHtml string `json:"extracted_inner_html" xml:"extracted_inner_html" required:"true"`
	// Canonical representation of a knowledge base document's stored content.
	//
	// HTML is the legacy default; documents created before this field existed are
	// interpreted as HTML.
	ContentFormat    *ContentFormat        `json:"content_format,omitempty" xml:"content_format,omitempty"`
	Filename         string                `json:"filename" xml:"filename" required:"true"`
	ExternalSyncInfo *ExternalFileSyncInfo `json:"external_sync_info,omitempty" xml:"external_sync_info,omitempty"`
	AutoSyncInfo     *AutoSyncInfo         `json:"auto_sync_info,omitempty" xml:"auto_sync_info,omitempty"`
	// In-flight or last refresh state for an externally-synced file. Used by clients to render sync progress and disable re-sync while a refresh is queued or processing.
	RefreshStatus   *FileRefreshStatus `json:"refresh_status,omitempty" xml:"refresh_status,omitempty"`
	IsFrozen        *bool              `json:"is_frozen,omitempty" xml:"is_frozen,omitempty"`
	extraProperties map[string]interface{}
}

func (g *GetKnowledgeBaseFileResponseModel) GetId() (value string) {
	if g == nil {
		return
	}
	return g.Id
}

func (g *GetKnowledgeBaseFileResponseModel) GetName() (value string) {
	if g == nil {
		return
	}
	return g.Name
}

func (g *GetKnowledgeBaseFileResponseModel) GetMetadata() (value *KnowledgeBaseDocumentMetadataResponseModel) {
	if g == nil {
		return
	}
	return g.Metadata
}

func (g *GetKnowledgeBaseFileResponseModel) GetSupportedUsages() (value []DocumentUsageModeEnum) {
	if g == nil {
		return
	}
	return g.SupportedUsages
}

func (g *GetKnowledgeBaseFileResponseModel) GetAccessInfo() (value *ResourceAccessInfo) {
	if g == nil {
		return
	}
	return g.AccessInfo
}

func (g *GetKnowledgeBaseFileResponseModel) GetFolderParentId() (value *string) {
	if g == nil {
		return
	}
	return g.FolderParentId
}

func (g *GetKnowledgeBaseFileResponseModel) GetFolderPath() (value []*KnowledgeBaseFolderPathSegmentResponseModel) {
	if g == nil {
		return
	}
	return g.FolderPath
}

func (g *GetKnowledgeBaseFileResponseModel) GetExtractedInnerHtml() (value string) {
	if g == nil {
		return
	}
	return g.ExtractedInnerHtml
}

func (g *GetKnowledgeBaseFileResponseModel) GetContentFormat() (value *ContentFormat) {
	if g == nil {
		return
	}
	return g.ContentFormat
}

func (g *GetKnowledgeBaseFileResponseModel) GetFilename() (value string) {
	if g == nil {
		return
	}
	return g.Filename
}

func (g *GetKnowledgeBaseFileResponseModel) GetExternalSyncInfo() (value *ExternalFileSyncInfo) {
	if g == nil {
		return
	}
	return g.ExternalSyncInfo
}

func (g *GetKnowledgeBaseFileResponseModel) GetAutoSyncInfo() (value *AutoSyncInfo) {
	if g == nil {
		return
	}
	return g.AutoSyncInfo
}

func (g *GetKnowledgeBaseFileResponseModel) GetRefreshStatus() (value *FileRefreshStatus) {
	if g == nil {
		return
	}
	return g.RefreshStatus
}

func (g *GetKnowledgeBaseFileResponseModel) GetIsFrozen() (value *bool) {
	if g == nil {
		return
	}
	return g.IsFrozen
}

func (g *GetKnowledgeBaseFileResponseModel) SetId(id string) {
	g.Id = id
}

func (g *GetKnowledgeBaseFileResponseModel) SetName(name string) {
	g.Name = name
}

func (g *GetKnowledgeBaseFileResponseModel) SetMetadata(metadata *KnowledgeBaseDocumentMetadataResponseModel) {
	g.Metadata = metadata
}

func (g *GetKnowledgeBaseFileResponseModel) SetSupportedUsages(supportedUsages []DocumentUsageModeEnum) {
	g.SupportedUsages = supportedUsages
}

func (g *GetKnowledgeBaseFileResponseModel) SetAccessInfo(accessInfo *ResourceAccessInfo) {
	g.AccessInfo = accessInfo
}

func (g *GetKnowledgeBaseFileResponseModel) SetFolderParentId(folderParentId *string) {
	g.FolderParentId = folderParentId
}

func (g *GetKnowledgeBaseFileResponseModel) SetFolderPath(folderPath []*KnowledgeBaseFolderPathSegmentResponseModel) {
	g.FolderPath = folderPath
}

func (g *GetKnowledgeBaseFileResponseModel) SetExtractedInnerHtml(extractedInnerHtml string) {
	g.ExtractedInnerHtml = extractedInnerHtml
}

func (g *GetKnowledgeBaseFileResponseModel) SetContentFormat(contentFormat *ContentFormat) {
	g.ContentFormat = contentFormat
}

func (g *GetKnowledgeBaseFileResponseModel) SetFilename(filename string) {
	g.Filename = filename
}

func (g *GetKnowledgeBaseFileResponseModel) SetExternalSyncInfo(externalSyncInfo *ExternalFileSyncInfo) {
	g.ExternalSyncInfo = externalSyncInfo
}

func (g *GetKnowledgeBaseFileResponseModel) SetAutoSyncInfo(autoSyncInfo *AutoSyncInfo) {
	g.AutoSyncInfo = autoSyncInfo
}

func (g *GetKnowledgeBaseFileResponseModel) SetRefreshStatus(refreshStatus *FileRefreshStatus) {
	g.RefreshStatus = refreshStatus
}

func (g *GetKnowledgeBaseFileResponseModel) SetIsFrozen(isFrozen *bool) {
	g.IsFrozen = isFrozen
}

func (g *GetKnowledgeBaseFileResponseModel) Type() string {
	return g.type_
}

func (g *GetKnowledgeBaseFileResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetKnowledgeBaseFileResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetKnowledgeBaseFileResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetKnowledgeBaseFileResponseModel) UnmarshalJSON(data []byte) error {
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

func (g *GetKnowledgeBaseFileResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetKnowledgeBaseFileResponseModel
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*g),
		Type:  "file",
	}
	return json.Marshal(marshaler)
}
