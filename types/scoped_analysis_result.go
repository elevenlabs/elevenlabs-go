package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ScopedAnalysisResult struct {
	Scope         AnalysisScope `json:"scope" xml:"scope" required:"true"`
	SourceAgentId string        `json:"source_agent_id" xml:"source_agent_id" required:"true"`
	// Branch of the agent for this scoped block; disambiguates repeated agent_id.
	SourceBranchId            *string                                                            `json:"source_branch_id,omitempty" xml:"source_branch_id,omitempty"`
	EvaluationCriteriaResults map[string]*ConversationHistoryEvaluationCriteriaResultCommonModel `json:"evaluation_criteria_results,omitempty"`
	DataCollectionResults     map[string]*DataCollectionResultCommonModel                        `json:"data_collection_results,omitempty"`
	Successful                EvaluationSuccessResult                                            `json:"successful" xml:"successful" required:"true"`
	SuccessScore              *float64                                                           `json:"success_score,omitempty" xml:"success_score,omitempty"`
	extraProperties           map[string]interface{}
}

func (s *ScopedAnalysisResult) GetScope() (value AnalysisScope) {
	if s == nil {
		return
	}
	return s.Scope
}

func (s *ScopedAnalysisResult) GetSourceAgentId() (value string) {
	if s == nil {
		return
	}
	return s.SourceAgentId
}

func (s *ScopedAnalysisResult) GetSourceBranchId() (value *string) {
	if s == nil {
		return
	}
	return s.SourceBranchId
}

func (s *ScopedAnalysisResult) GetEvaluationCriteriaResults() (value map[string]*ConversationHistoryEvaluationCriteriaResultCommonModel) {
	if s == nil {
		return
	}
	return s.EvaluationCriteriaResults
}

func (s *ScopedAnalysisResult) GetDataCollectionResults() (value map[string]*DataCollectionResultCommonModel) {
	if s == nil {
		return
	}
	return s.DataCollectionResults
}

func (s *ScopedAnalysisResult) GetSuccessful() (value EvaluationSuccessResult) {
	if s == nil {
		return
	}
	return s.Successful
}

func (s *ScopedAnalysisResult) GetSuccessScore() (value *float64) {
	if s == nil {
		return
	}
	return s.SuccessScore
}

func (s *ScopedAnalysisResult) SetScope(scope AnalysisScope) {
	s.Scope = scope
}

func (s *ScopedAnalysisResult) SetSourceAgentId(sourceAgentId string) {
	s.SourceAgentId = sourceAgentId
}

func (s *ScopedAnalysisResult) SetSourceBranchId(sourceBranchId *string) {
	s.SourceBranchId = sourceBranchId
}

func (s *ScopedAnalysisResult) SetEvaluationCriteriaResults(evaluationCriteriaResults map[string]*ConversationHistoryEvaluationCriteriaResultCommonModel) {
	s.EvaluationCriteriaResults = evaluationCriteriaResults
}

func (s *ScopedAnalysisResult) SetDataCollectionResults(dataCollectionResults map[string]*DataCollectionResultCommonModel) {
	s.DataCollectionResults = dataCollectionResults
}

func (s *ScopedAnalysisResult) SetSuccessful(successful EvaluationSuccessResult) {
	s.Successful = successful
}

func (s *ScopedAnalysisResult) SetSuccessScore(successScore *float64) {
	s.SuccessScore = successScore
}

func (s *ScopedAnalysisResult) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s ScopedAnalysisResult) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: ScopedAnalysisResult to string"
	}
	return string(jsonData)
}

func (s *ScopedAnalysisResult) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, s); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *ScopedAnalysisResult) MarshalJSON() ([]byte, error) {
	type embed ScopedAnalysisResult
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
