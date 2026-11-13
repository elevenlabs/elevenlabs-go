package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentKnowledgeBaseRagDocumentResponseModel struct {
	// ID of the knowledge base document.
	DocumentId string `json:"document_id" xml:"document_id" required:"true"`
	// Name of the knowledge base document.
	DocumentName string `json:"document_name" xml:"document_name" required:"true"`
	// Tracked source URL for URL documents, or null for other document types.
	SourceUrl *string `json:"source_url" xml:"source_url" required:"true"`
	// Content the agent would read, cut to the maximum documents length. Empty when the document has no extractable content.
	Content string `json:"content" xml:"content" required:"true"`
	// Whether the content was cut to the maximum documents length.
	Truncated bool `json:"truncated" xml:"truncated" required:"true"`
	// Canonical representation of a knowledge base document's stored content.
	//
	// HTML is the legacy default; documents created before this field existed are
	// interpreted as HTML.
	ContentFormat   ContentFormat             `json:"content_format" xml:"content_format" required:"true"`
	DocumentType    KnowledgeBaseDocumentType `json:"document_type" xml:"document_type" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AgentKnowledgeBaseRagDocumentResponseModel) GetDocumentId() (value string) {
	if a == nil {
		return
	}
	return a.DocumentId
}

func (a *AgentKnowledgeBaseRagDocumentResponseModel) GetDocumentName() (value string) {
	if a == nil {
		return
	}
	return a.DocumentName
}

func (a *AgentKnowledgeBaseRagDocumentResponseModel) GetSourceUrl() (value *string) {
	if a == nil {
		return
	}
	return a.SourceUrl
}

func (a *AgentKnowledgeBaseRagDocumentResponseModel) GetContent() (value string) {
	if a == nil {
		return
	}
	return a.Content
}

func (a *AgentKnowledgeBaseRagDocumentResponseModel) GetTruncated() (value bool) {
	if a == nil {
		return
	}
	return a.Truncated
}

func (a *AgentKnowledgeBaseRagDocumentResponseModel) GetContentFormat() (value ContentFormat) {
	if a == nil {
		return
	}
	return a.ContentFormat
}

func (a *AgentKnowledgeBaseRagDocumentResponseModel) GetDocumentType() (value KnowledgeBaseDocumentType) {
	if a == nil {
		return
	}
	return a.DocumentType
}

func (a *AgentKnowledgeBaseRagDocumentResponseModel) SetDocumentId(documentId string) {
	a.DocumentId = documentId
}

func (a *AgentKnowledgeBaseRagDocumentResponseModel) SetDocumentName(documentName string) {
	a.DocumentName = documentName
}

func (a *AgentKnowledgeBaseRagDocumentResponseModel) SetSourceUrl(sourceUrl *string) {
	a.SourceUrl = sourceUrl
}

func (a *AgentKnowledgeBaseRagDocumentResponseModel) SetContent(content string) {
	a.Content = content
}

func (a *AgentKnowledgeBaseRagDocumentResponseModel) SetTruncated(truncated bool) {
	a.Truncated = truncated
}

func (a *AgentKnowledgeBaseRagDocumentResponseModel) SetContentFormat(contentFormat ContentFormat) {
	a.ContentFormat = contentFormat
}

func (a *AgentKnowledgeBaseRagDocumentResponseModel) SetDocumentType(documentType KnowledgeBaseDocumentType) {
	a.DocumentType = documentType
}

func (a *AgentKnowledgeBaseRagDocumentResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentKnowledgeBaseRagDocumentResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentKnowledgeBaseRagDocumentResponseModel to string"
	}
	return string(jsonData)
}

func (a *AgentKnowledgeBaseRagDocumentResponseModel) UnmarshalJSON(data []byte) error {
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

func (a *AgentKnowledgeBaseRagDocumentResponseModel) MarshalJSON() ([]byte, error) {
	type embed AgentKnowledgeBaseRagDocumentResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
