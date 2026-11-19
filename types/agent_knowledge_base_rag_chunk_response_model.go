package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentKnowledgeBaseRagChunkResponseModel struct {
	// ID of the source knowledge base document.
	DocumentId string `json:"document_id" xml:"document_id" required:"true"`
	// Name of the source knowledge base document.
	DocumentName string `json:"document_name" xml:"document_name" required:"true"`
	// Tracked source URL for URL documents, or null for other document types.
	SourceUrl *string `json:"source_url" xml:"source_url" required:"true"`
	// ID of the retrieved chunk.
	ChunkId string `json:"chunk_id" xml:"chunk_id" required:"true"`
	// Text content of the retrieved chunk.
	Text string `json:"text" xml:"text" required:"true"`
	// Similarity distance when exposed by the retrieval strategy.
	VectorDistance *float64 `json:"vector_distance" xml:"vector_distance" required:"true"`
	// Canonical representation of a knowledge base document's stored content.
	//
	// HTML is the legacy default; documents created before this field existed are
	// interpreted as HTML.
	ContentFormat   ContentFormat             `json:"content_format" xml:"content_format" required:"true"`
	DocumentType    KnowledgeBaseDocumentType `json:"document_type" xml:"document_type" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AgentKnowledgeBaseRagChunkResponseModel) GetDocumentId() (value string) {
	if a == nil {
		return
	}
	return a.DocumentId
}

func (a *AgentKnowledgeBaseRagChunkResponseModel) GetDocumentName() (value string) {
	if a == nil {
		return
	}
	return a.DocumentName
}

func (a *AgentKnowledgeBaseRagChunkResponseModel) GetSourceUrl() (value *string) {
	if a == nil {
		return
	}
	return a.SourceUrl
}

func (a *AgentKnowledgeBaseRagChunkResponseModel) GetChunkId() (value string) {
	if a == nil {
		return
	}
	return a.ChunkId
}

func (a *AgentKnowledgeBaseRagChunkResponseModel) GetText() (value string) {
	if a == nil {
		return
	}
	return a.Text
}

func (a *AgentKnowledgeBaseRagChunkResponseModel) GetVectorDistance() (value *float64) {
	if a == nil {
		return
	}
	return a.VectorDistance
}

func (a *AgentKnowledgeBaseRagChunkResponseModel) GetContentFormat() (value ContentFormat) {
	if a == nil {
		return
	}
	return a.ContentFormat
}

func (a *AgentKnowledgeBaseRagChunkResponseModel) GetDocumentType() (value KnowledgeBaseDocumentType) {
	if a == nil {
		return
	}
	return a.DocumentType
}

func (a *AgentKnowledgeBaseRagChunkResponseModel) SetDocumentId(documentId string) {
	a.DocumentId = documentId
}

func (a *AgentKnowledgeBaseRagChunkResponseModel) SetDocumentName(documentName string) {
	a.DocumentName = documentName
}

func (a *AgentKnowledgeBaseRagChunkResponseModel) SetSourceUrl(sourceUrl *string) {
	a.SourceUrl = sourceUrl
}

func (a *AgentKnowledgeBaseRagChunkResponseModel) SetChunkId(chunkId string) {
	a.ChunkId = chunkId
}

func (a *AgentKnowledgeBaseRagChunkResponseModel) SetText(text string) {
	a.Text = text
}

func (a *AgentKnowledgeBaseRagChunkResponseModel) SetVectorDistance(vectorDistance *float64) {
	a.VectorDistance = vectorDistance
}

func (a *AgentKnowledgeBaseRagChunkResponseModel) SetContentFormat(contentFormat ContentFormat) {
	a.ContentFormat = contentFormat
}

func (a *AgentKnowledgeBaseRagChunkResponseModel) SetDocumentType(documentType KnowledgeBaseDocumentType) {
	a.DocumentType = documentType
}

func (a *AgentKnowledgeBaseRagChunkResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentKnowledgeBaseRagChunkResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentKnowledgeBaseRagChunkResponseModel to string"
	}
	return string(jsonData)
}

func (a *AgentKnowledgeBaseRagChunkResponseModel) UnmarshalJSON(data []byte) error {
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

func (a *AgentKnowledgeBaseRagChunkResponseModel) MarshalJSON() ([]byte, error) {
	type embed AgentKnowledgeBaseRagChunkResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
