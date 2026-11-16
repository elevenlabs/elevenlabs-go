package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentKnowledgeBaseRagQueryRequestModel struct {
	Strategy *SearchStrategy `json:"strategy,omitempty" xml:"strategy,omitempty"`
	// Query to search the knowledge base for. Required for the semantic and keyword strategies.
	Query *string `json:"query,omitempty" xml:"query,omitempty" minLength:"1"`
	// IDs or names of knowledge base documents or folders. Semantic and keyword search only search these entries; cat reads the single document given; ls lists the single folder given, or the top level of the knowledge base when omitted.
	Documents []string `json:"documents,omitempty" xml:"documents,omitempty" maxItems:"100"`
	// When true (the default), retrieval uses the agent's own RAG settings, reproducing exactly what the agent would retrieve. Set to false to retrieve with neutral default RAG settings instead (the agent's embedding model is always kept, since it determines which vector index exists). Useful for auditing the knowledge base independently of how a particular agent is tuned.
	UseAgentDefaults *bool `json:"use_agent_defaults,omitempty" xml:"use_agent_defaults,omitempty"`
	// Optional maximum total character length of document chunks returned. Overrides the selected RAG settings for this query only.
	MaxDocumentsLength *int `json:"max_documents_length,omitempty" xml:"max_documents_length,omitempty" max:"50000"`
	// Optional maximum number of document chunks retrieved. Overrides the selected RAG settings for this query only.
	MaxRetrievedRagChunksCount *int `json:"max_retrieved_rag_chunks_count,omitempty" xml:"max_retrieved_rag_chunks_count,omitempty" max:"20"`
	extraProperties            map[string]interface{}
}

func (a *AgentKnowledgeBaseRagQueryRequestModel) GetStrategy() (value *SearchStrategy) {
	if a == nil {
		return
	}
	return a.Strategy
}

func (a *AgentKnowledgeBaseRagQueryRequestModel) GetQuery() (value *string) {
	if a == nil {
		return
	}
	return a.Query
}

func (a *AgentKnowledgeBaseRagQueryRequestModel) GetDocuments() (value []string) {
	if a == nil {
		return
	}
	return a.Documents
}

func (a *AgentKnowledgeBaseRagQueryRequestModel) GetUseAgentDefaults() (value *bool) {
	if a == nil {
		return
	}
	return a.UseAgentDefaults
}

func (a *AgentKnowledgeBaseRagQueryRequestModel) GetMaxDocumentsLength() (value *int) {
	if a == nil {
		return
	}
	return a.MaxDocumentsLength
}

func (a *AgentKnowledgeBaseRagQueryRequestModel) GetMaxRetrievedRagChunksCount() (value *int) {
	if a == nil {
		return
	}
	return a.MaxRetrievedRagChunksCount
}

func (a *AgentKnowledgeBaseRagQueryRequestModel) SetStrategy(strategy *SearchStrategy) {
	a.Strategy = strategy
}

func (a *AgentKnowledgeBaseRagQueryRequestModel) SetQuery(query *string) {
	a.Query = query
}

func (a *AgentKnowledgeBaseRagQueryRequestModel) SetDocuments(documents []string) {
	a.Documents = documents
}

func (a *AgentKnowledgeBaseRagQueryRequestModel) SetUseAgentDefaults(useAgentDefaults *bool) {
	a.UseAgentDefaults = useAgentDefaults
}

func (a *AgentKnowledgeBaseRagQueryRequestModel) SetMaxDocumentsLength(maxDocumentsLength *int) {
	a.MaxDocumentsLength = maxDocumentsLength
}

func (a *AgentKnowledgeBaseRagQueryRequestModel) SetMaxRetrievedRagChunksCount(maxRetrievedRagChunksCount *int) {
	a.MaxRetrievedRagChunksCount = maxRetrievedRagChunksCount
}

func (a *AgentKnowledgeBaseRagQueryRequestModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentKnowledgeBaseRagQueryRequestModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentKnowledgeBaseRagQueryRequestModel to string"
	}
	return string(jsonData)
}

func (a *AgentKnowledgeBaseRagQueryRequestModel) UnmarshalJSON(data []byte) error {
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

func (a *AgentKnowledgeBaseRagQueryRequestModel) MarshalJSON() ([]byte, error) {
	type embed AgentKnowledgeBaseRagQueryRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
