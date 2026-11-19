package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetKnowledgeBaseTextResponseModel struct {
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
	ExtractedInnerHtml string `json:"extracted_inner_html" xml:"extracted_inner_html" required:"true"`
	// Canonical representation of a knowledge base document's stored content.
	//
	// HTML is the legacy default; documents created before this field existed are
	// interpreted as HTML.
	ContentFormat   *ContentFormat `json:"content_format,omitempty" xml:"content_format,omitempty"`
	extraProperties map[string]interface{}
}

func (g *GetKnowledgeBaseTextResponseModel) GetId() (value string) {
	if g == nil {
		return
	}
	return g.Id
}

func (g *GetKnowledgeBaseTextResponseModel) GetName() (value string) {
	if g == nil {
		return
	}
	return g.Name
}

func (g *GetKnowledgeBaseTextResponseModel) GetMetadata() (value *KnowledgeBaseDocumentMetadataResponseModel) {
	if g == nil {
		return
	}
	return g.Metadata
}

func (g *GetKnowledgeBaseTextResponseModel) GetSupportedUsages() (value []DocumentUsageModeEnum) {
	if g == nil {
		return
	}
	return g.SupportedUsages
}

func (g *GetKnowledgeBaseTextResponseModel) GetAccessInfo() (value *ResourceAccessInfo) {
	if g == nil {
		return
	}
	return g.AccessInfo
}

func (g *GetKnowledgeBaseTextResponseModel) GetFolderParentId() (value *string) {
	if g == nil {
		return
	}
	return g.FolderParentId
}

func (g *GetKnowledgeBaseTextResponseModel) GetChunkingOverride() (value *ChunkingOverride) {
	if g == nil {
		return
	}
	return g.ChunkingOverride
}

func (g *GetKnowledgeBaseTextResponseModel) GetFolderPath() (value []*KnowledgeBaseFolderPathSegmentResponseModel) {
	if g == nil {
		return
	}
	return g.FolderPath
}

func (g *GetKnowledgeBaseTextResponseModel) GetExtractedInnerHtml() (value string) {
	if g == nil {
		return
	}
	return g.ExtractedInnerHtml
}

func (g *GetKnowledgeBaseTextResponseModel) GetContentFormat() (value *ContentFormat) {
	if g == nil {
		return
	}
	return g.ContentFormat
}

func (g *GetKnowledgeBaseTextResponseModel) SetId(id string) {
	g.Id = id
}

func (g *GetKnowledgeBaseTextResponseModel) SetName(name string) {
	g.Name = name
}

func (g *GetKnowledgeBaseTextResponseModel) SetMetadata(metadata *KnowledgeBaseDocumentMetadataResponseModel) {
	g.Metadata = metadata
}

func (g *GetKnowledgeBaseTextResponseModel) SetSupportedUsages(supportedUsages []DocumentUsageModeEnum) {
	g.SupportedUsages = supportedUsages
}

func (g *GetKnowledgeBaseTextResponseModel) SetAccessInfo(accessInfo *ResourceAccessInfo) {
	g.AccessInfo = accessInfo
}

func (g *GetKnowledgeBaseTextResponseModel) SetFolderParentId(folderParentId *string) {
	g.FolderParentId = folderParentId
}

func (g *GetKnowledgeBaseTextResponseModel) SetChunkingOverride(chunkingOverride *ChunkingOverride) {
	g.ChunkingOverride = chunkingOverride
}

func (g *GetKnowledgeBaseTextResponseModel) SetFolderPath(folderPath []*KnowledgeBaseFolderPathSegmentResponseModel) {
	g.FolderPath = folderPath
}

func (g *GetKnowledgeBaseTextResponseModel) SetExtractedInnerHtml(extractedInnerHtml string) {
	g.ExtractedInnerHtml = extractedInnerHtml
}

func (g *GetKnowledgeBaseTextResponseModel) SetContentFormat(contentFormat *ContentFormat) {
	g.ContentFormat = contentFormat
}

func (g *GetKnowledgeBaseTextResponseModel) Type() string {
	return g.type_
}

func (g *GetKnowledgeBaseTextResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetKnowledgeBaseTextResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetKnowledgeBaseTextResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetKnowledgeBaseTextResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, g); err != nil {
		return err
	}
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "text" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", g, "text", *literals.Type)
	}
	g.type_ = "text"
	extraProperties, err := internal.ExtractExtraProperties(data, *g, "type")
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetKnowledgeBaseTextResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetKnowledgeBaseTextResponseModel
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*g),
		Type:  "text",
	}
	return json.Marshal(marshaler)
}
