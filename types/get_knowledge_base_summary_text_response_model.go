package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetKnowledgeBaseSummaryTextResponseModel struct {
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
	DependentAgents []*GetKnowledgeBaseSummaryTextResponseModelDependentAgentsItem `json:"dependent_agents" xml:"dependent_agents" required:"true"`
	type_           string
	extraProperties map[string]interface{}
}

func (g *GetKnowledgeBaseSummaryTextResponseModel) GetId() (value string) {
	if g == nil {
		return
	}
	return g.Id
}

func (g *GetKnowledgeBaseSummaryTextResponseModel) GetName() (value string) {
	if g == nil {
		return
	}
	return g.Name
}

func (g *GetKnowledgeBaseSummaryTextResponseModel) GetMetadata() (value *KnowledgeBaseDocumentMetadataResponseModel) {
	if g == nil {
		return
	}
	return g.Metadata
}

func (g *GetKnowledgeBaseSummaryTextResponseModel) GetSupportedUsages() (value []DocumentUsageModeEnum) {
	if g == nil {
		return
	}
	return g.SupportedUsages
}

func (g *GetKnowledgeBaseSummaryTextResponseModel) GetAccessInfo() (value *ResourceAccessInfo) {
	if g == nil {
		return
	}
	return g.AccessInfo
}

func (g *GetKnowledgeBaseSummaryTextResponseModel) GetFolderParentId() (value *string) {
	if g == nil {
		return
	}
	return g.FolderParentId
}

func (g *GetKnowledgeBaseSummaryTextResponseModel) GetChunkingOverride() (value *ChunkingOverride) {
	if g == nil {
		return
	}
	return g.ChunkingOverride
}

func (g *GetKnowledgeBaseSummaryTextResponseModel) GetFolderPath() (value []*KnowledgeBaseFolderPathSegmentSummaryResponseModel) {
	if g == nil {
		return
	}
	return g.FolderPath
}

func (g *GetKnowledgeBaseSummaryTextResponseModel) GetDependentAgents() (value []*GetKnowledgeBaseSummaryTextResponseModelDependentAgentsItem) {
	if g == nil {
		return
	}
	return g.DependentAgents
}

func (g *GetKnowledgeBaseSummaryTextResponseModel) SetId(id string) {
	g.Id = id
}

func (g *GetKnowledgeBaseSummaryTextResponseModel) SetName(name string) {
	g.Name = name
}

func (g *GetKnowledgeBaseSummaryTextResponseModel) SetMetadata(metadata *KnowledgeBaseDocumentMetadataResponseModel) {
	g.Metadata = metadata
}

func (g *GetKnowledgeBaseSummaryTextResponseModel) SetSupportedUsages(supportedUsages []DocumentUsageModeEnum) {
	g.SupportedUsages = supportedUsages
}

func (g *GetKnowledgeBaseSummaryTextResponseModel) SetAccessInfo(accessInfo *ResourceAccessInfo) {
	g.AccessInfo = accessInfo
}

func (g *GetKnowledgeBaseSummaryTextResponseModel) SetFolderParentId(folderParentId *string) {
	g.FolderParentId = folderParentId
}

func (g *GetKnowledgeBaseSummaryTextResponseModel) SetChunkingOverride(chunkingOverride *ChunkingOverride) {
	g.ChunkingOverride = chunkingOverride
}

func (g *GetKnowledgeBaseSummaryTextResponseModel) SetFolderPath(folderPath []*KnowledgeBaseFolderPathSegmentSummaryResponseModel) {
	g.FolderPath = folderPath
}

func (g *GetKnowledgeBaseSummaryTextResponseModel) SetDependentAgents(dependentAgents []*GetKnowledgeBaseSummaryTextResponseModelDependentAgentsItem) {
	g.DependentAgents = dependentAgents
}

func (g *GetKnowledgeBaseSummaryTextResponseModel) Type() string {
	return g.type_
}

func (g *GetKnowledgeBaseSummaryTextResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetKnowledgeBaseSummaryTextResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetKnowledgeBaseSummaryTextResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetKnowledgeBaseSummaryTextResponseModel) UnmarshalJSON(data []byte) error {
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

func (g *GetKnowledgeBaseSummaryTextResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetKnowledgeBaseSummaryTextResponseModel
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*g),
		Type:  "text",
	}
	return json.Marshal(marshaler)
}
