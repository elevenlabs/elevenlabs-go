package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetKnowledgeBaseUrlResponseModel struct {
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
	FolderPath         []*KnowledgeBaseFolderPathSegmentResponseModel `json:"folder_path,omitempty" xml:"folder_path,omitempty"`
	type_              string
	Url                string `json:"url" xml:"url" required:"true"`
	ExtractedInnerHtml string `json:"extracted_inner_html" xml:"extracted_inner_html" required:"true"`
	// Canonical representation of a knowledge base document's stored content.
	//
	// HTML is the legacy default; documents created before this field existed are
	// interpreted as HTML.
	ContentFormat   *ContentFormat `json:"content_format,omitempty" xml:"content_format,omitempty"`
	AutoSyncInfo    *AutoSyncInfo  `json:"auto_sync_info,omitempty" xml:"auto_sync_info,omitempty"`
	extraProperties map[string]interface{}
}

func (g *GetKnowledgeBaseUrlResponseModel) GetId() (value string) {
	if g == nil {
		return
	}
	return g.Id
}

func (g *GetKnowledgeBaseUrlResponseModel) GetName() (value string) {
	if g == nil {
		return
	}
	return g.Name
}

func (g *GetKnowledgeBaseUrlResponseModel) GetMetadata() (value *KnowledgeBaseDocumentMetadataResponseModel) {
	if g == nil {
		return
	}
	return g.Metadata
}

func (g *GetKnowledgeBaseUrlResponseModel) GetSupportedUsages() (value []DocumentUsageModeEnum) {
	if g == nil {
		return
	}
	return g.SupportedUsages
}

func (g *GetKnowledgeBaseUrlResponseModel) GetAccessInfo() (value *ResourceAccessInfo) {
	if g == nil {
		return
	}
	return g.AccessInfo
}

func (g *GetKnowledgeBaseUrlResponseModel) GetFolderParentId() (value *string) {
	if g == nil {
		return
	}
	return g.FolderParentId
}

func (g *GetKnowledgeBaseUrlResponseModel) GetChunkingOverride() (value *ChunkingOverride) {
	if g == nil {
		return
	}
	return g.ChunkingOverride
}

func (g *GetKnowledgeBaseUrlResponseModel) GetFolderPath() (value []*KnowledgeBaseFolderPathSegmentResponseModel) {
	if g == nil {
		return
	}
	return g.FolderPath
}

func (g *GetKnowledgeBaseUrlResponseModel) GetUrl() (value string) {
	if g == nil {
		return
	}
	return g.Url
}

func (g *GetKnowledgeBaseUrlResponseModel) GetExtractedInnerHtml() (value string) {
	if g == nil {
		return
	}
	return g.ExtractedInnerHtml
}

func (g *GetKnowledgeBaseUrlResponseModel) GetContentFormat() (value *ContentFormat) {
	if g == nil {
		return
	}
	return g.ContentFormat
}

func (g *GetKnowledgeBaseUrlResponseModel) GetAutoSyncInfo() (value *AutoSyncInfo) {
	if g == nil {
		return
	}
	return g.AutoSyncInfo
}

func (g *GetKnowledgeBaseUrlResponseModel) SetId(id string) {
	g.Id = id
}

func (g *GetKnowledgeBaseUrlResponseModel) SetName(name string) {
	g.Name = name
}

func (g *GetKnowledgeBaseUrlResponseModel) SetMetadata(metadata *KnowledgeBaseDocumentMetadataResponseModel) {
	g.Metadata = metadata
}

func (g *GetKnowledgeBaseUrlResponseModel) SetSupportedUsages(supportedUsages []DocumentUsageModeEnum) {
	g.SupportedUsages = supportedUsages
}

func (g *GetKnowledgeBaseUrlResponseModel) SetAccessInfo(accessInfo *ResourceAccessInfo) {
	g.AccessInfo = accessInfo
}

func (g *GetKnowledgeBaseUrlResponseModel) SetFolderParentId(folderParentId *string) {
	g.FolderParentId = folderParentId
}

func (g *GetKnowledgeBaseUrlResponseModel) SetChunkingOverride(chunkingOverride *ChunkingOverride) {
	g.ChunkingOverride = chunkingOverride
}

func (g *GetKnowledgeBaseUrlResponseModel) SetFolderPath(folderPath []*KnowledgeBaseFolderPathSegmentResponseModel) {
	g.FolderPath = folderPath
}

func (g *GetKnowledgeBaseUrlResponseModel) SetUrl(url string) {
	g.Url = url
}

func (g *GetKnowledgeBaseUrlResponseModel) SetExtractedInnerHtml(extractedInnerHtml string) {
	g.ExtractedInnerHtml = extractedInnerHtml
}

func (g *GetKnowledgeBaseUrlResponseModel) SetContentFormat(contentFormat *ContentFormat) {
	g.ContentFormat = contentFormat
}

func (g *GetKnowledgeBaseUrlResponseModel) SetAutoSyncInfo(autoSyncInfo *AutoSyncInfo) {
	g.AutoSyncInfo = autoSyncInfo
}

func (g *GetKnowledgeBaseUrlResponseModel) Type() string {
	return g.type_
}

func (g *GetKnowledgeBaseUrlResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetKnowledgeBaseUrlResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetKnowledgeBaseUrlResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetKnowledgeBaseUrlResponseModel) UnmarshalJSON(data []byte) error {
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

func (g *GetKnowledgeBaseUrlResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetKnowledgeBaseUrlResponseModel
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*g),
		Type:  "url",
	}
	return json.Marshal(marshaler)
}
