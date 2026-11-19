package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type RagConfigWorkflowOverrideOutput struct {
	Enabled        *bool               `json:"enabled,omitempty" xml:"enabled,omitempty"`
	EmbeddingModel *EmbeddingModelEnum `json:"embedding_model,omitempty" xml:"embedding_model,omitempty"`
	// Maximum vector distance of retrieved chunks.
	MaxVectorDistance *float64 `json:"max_vector_distance,omitempty" xml:"max_vector_distance,omitempty"`
	// Maximum total length of document chunks retrieved from RAG.
	MaxDocumentsLength *int `json:"max_documents_length,omitempty" xml:"max_documents_length,omitempty"`
	// Maximum number of RAG document chunks to initially retrieve from the vector store. These are then further filtered by vector distance and total length.
	MaxRetrievedRagChunksCount *int `json:"max_retrieved_rag_chunks_count,omitempty" xml:"max_retrieved_rag_chunks_count,omitempty"`
	// Number of candidates evaluated in ANN vector search. Higher number means better results, but higher latency. Minimum recommended value is 100. If disabled, the default value is used.
	NumCandidates *int `json:"num_candidates,omitempty" xml:"num_candidates,omitempty"`
	// Custom prompt for rewriting user queries before RAG retrieval. The conversation history will be automatically appended at the end. If not set, the default prompt will be used.
	QueryRewritePromptOverride *string `json:"query_rewrite_prompt_override,omitempty" xml:"query_rewrite_prompt_override,omitempty"`
	// When set, the agent uses the knowledge_base tool instead of the legacy knowledge_base_rag tool. None means the agent is not opted in.
	KnowledgeBaseToolInfo *KnowledgeBaseToolInfo `json:"knowledge_base_tool_info,omitempty" xml:"knowledge_base_tool_info,omitempty"`
	extraProperties       map[string]interface{}
}

func (r *RagConfigWorkflowOverrideOutput) GetEnabled() (value *bool) {
	if r == nil {
		return
	}
	return r.Enabled
}

func (r *RagConfigWorkflowOverrideOutput) GetEmbeddingModel() (value *EmbeddingModelEnum) {
	if r == nil {
		return
	}
	return r.EmbeddingModel
}

func (r *RagConfigWorkflowOverrideOutput) GetMaxVectorDistance() (value *float64) {
	if r == nil {
		return
	}
	return r.MaxVectorDistance
}

func (r *RagConfigWorkflowOverrideOutput) GetMaxDocumentsLength() (value *int) {
	if r == nil {
		return
	}
	return r.MaxDocumentsLength
}

func (r *RagConfigWorkflowOverrideOutput) GetMaxRetrievedRagChunksCount() (value *int) {
	if r == nil {
		return
	}
	return r.MaxRetrievedRagChunksCount
}

func (r *RagConfigWorkflowOverrideOutput) GetNumCandidates() (value *int) {
	if r == nil {
		return
	}
	return r.NumCandidates
}

func (r *RagConfigWorkflowOverrideOutput) GetQueryRewritePromptOverride() (value *string) {
	if r == nil {
		return
	}
	return r.QueryRewritePromptOverride
}

func (r *RagConfigWorkflowOverrideOutput) GetKnowledgeBaseToolInfo() (value *KnowledgeBaseToolInfo) {
	if r == nil {
		return
	}
	return r.KnowledgeBaseToolInfo
}

func (r *RagConfigWorkflowOverrideOutput) SetEnabled(enabled *bool) {
	r.Enabled = enabled
}

func (r *RagConfigWorkflowOverrideOutput) SetEmbeddingModel(embeddingModel *EmbeddingModelEnum) {
	r.EmbeddingModel = embeddingModel
}

func (r *RagConfigWorkflowOverrideOutput) SetMaxVectorDistance(maxVectorDistance *float64) {
	r.MaxVectorDistance = maxVectorDistance
}

func (r *RagConfigWorkflowOverrideOutput) SetMaxDocumentsLength(maxDocumentsLength *int) {
	r.MaxDocumentsLength = maxDocumentsLength
}

func (r *RagConfigWorkflowOverrideOutput) SetMaxRetrievedRagChunksCount(maxRetrievedRagChunksCount *int) {
	r.MaxRetrievedRagChunksCount = maxRetrievedRagChunksCount
}

func (r *RagConfigWorkflowOverrideOutput) SetNumCandidates(numCandidates *int) {
	r.NumCandidates = numCandidates
}

func (r *RagConfigWorkflowOverrideOutput) SetQueryRewritePromptOverride(queryRewritePromptOverride *string) {
	r.QueryRewritePromptOverride = queryRewritePromptOverride
}

func (r *RagConfigWorkflowOverrideOutput) SetKnowledgeBaseToolInfo(knowledgeBaseToolInfo *KnowledgeBaseToolInfo) {
	r.KnowledgeBaseToolInfo = knowledgeBaseToolInfo
}

func (r *RagConfigWorkflowOverrideOutput) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r RagConfigWorkflowOverrideOutput) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RagConfigWorkflowOverrideOutput to string"
	}
	return string(jsonData)
}

func (r *RagConfigWorkflowOverrideOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, r); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *RagConfigWorkflowOverrideOutput) MarshalJSON() ([]byte, error) {
	type embed RagConfigWorkflowOverrideOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
