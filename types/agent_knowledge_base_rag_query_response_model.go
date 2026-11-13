package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentKnowledgeBaseRagQueryResponseModel struct {
	Strategy SearchStrategy `json:"strategy" xml:"strategy" required:"true"`
	// Raw query used for retrieval (echoes the request query), or null for the cat and ls strategies.
	RetrievalQuery *string `json:"retrieval_query" xml:"retrieval_query" required:"true"`
	// Ranked chunks the agent would retrieve with semantic or keyword search, after distance and length filtering. Empty for the cat and ls strategies.
	Chunks []*AgentKnowledgeBaseRagChunkResponseModel `json:"chunks" xml:"chunks" required:"true"`
	// Document read by the cat strategy, or null when the reference matched no document or another strategy ran.
	Document *AgentKnowledgeBaseRagDocumentResponseModel `json:"document,omitempty" xml:"document,omitempty"`
	// Entries listed by the ls strategy, or null when the reference matched nothing or another strategy ran.
	Listing         *AgentKnowledgeBaseRagListingResponseModel `json:"listing,omitempty" xml:"listing,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AgentKnowledgeBaseRagQueryResponseModel) GetStrategy() (value SearchStrategy) {
	if a == nil {
		return
	}
	return a.Strategy
}

func (a *AgentKnowledgeBaseRagQueryResponseModel) GetRetrievalQuery() (value *string) {
	if a == nil {
		return
	}
	return a.RetrievalQuery
}

func (a *AgentKnowledgeBaseRagQueryResponseModel) GetChunks() (value []*AgentKnowledgeBaseRagChunkResponseModel) {
	if a == nil {
		return
	}
	return a.Chunks
}

func (a *AgentKnowledgeBaseRagQueryResponseModel) GetDocument() (value *AgentKnowledgeBaseRagDocumentResponseModel) {
	if a == nil {
		return
	}
	return a.Document
}

func (a *AgentKnowledgeBaseRagQueryResponseModel) GetListing() (value *AgentKnowledgeBaseRagListingResponseModel) {
	if a == nil {
		return
	}
	return a.Listing
}

func (a *AgentKnowledgeBaseRagQueryResponseModel) SetStrategy(strategy SearchStrategy) {
	a.Strategy = strategy
}

func (a *AgentKnowledgeBaseRagQueryResponseModel) SetRetrievalQuery(retrievalQuery *string) {
	a.RetrievalQuery = retrievalQuery
}

func (a *AgentKnowledgeBaseRagQueryResponseModel) SetChunks(chunks []*AgentKnowledgeBaseRagChunkResponseModel) {
	a.Chunks = chunks
}

func (a *AgentKnowledgeBaseRagQueryResponseModel) SetDocument(document *AgentKnowledgeBaseRagDocumentResponseModel) {
	a.Document = document
}

func (a *AgentKnowledgeBaseRagQueryResponseModel) SetListing(listing *AgentKnowledgeBaseRagListingResponseModel) {
	a.Listing = listing
}

func (a *AgentKnowledgeBaseRagQueryResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentKnowledgeBaseRagQueryResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentKnowledgeBaseRagQueryResponseModel to string"
	}
	return string(jsonData)
}

func (a *AgentKnowledgeBaseRagQueryResponseModel) UnmarshalJSON(data []byte) error {
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

func (a *AgentKnowledgeBaseRagQueryResponseModel) MarshalJSON() ([]byte, error) {
	type embed AgentKnowledgeBaseRagQueryResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
