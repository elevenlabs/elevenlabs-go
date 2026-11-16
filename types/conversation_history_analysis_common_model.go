package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationHistoryAnalysisCommonModel struct {
	EvaluationCriteriaResults     map[string]*ConversationHistoryEvaluationCriteriaResultCommonModel `json:"evaluation_criteria_results,omitempty"`
	DataCollectionResults         map[string]*DataCollectionResultCommonModel                        `json:"data_collection_results,omitempty"`
	EvaluationCriteriaResultsList []*ConversationHistoryEvaluationCriteriaResultCommonModel          `json:"evaluation_criteria_results_list,omitempty" xml:"evaluation_criteria_results_list,omitempty"`
	DataCollectionResultsList     []*DataCollectionResultCommonModel                                 `json:"data_collection_results_list,omitempty" xml:"data_collection_results_list,omitempty"`
	CallSuccessful                EvaluationSuccessResult                                            `json:"call_successful" xml:"call_successful" required:"true"`
	CallSuccessScore              *float64                                                           `json:"call_success_score,omitempty" xml:"call_success_score,omitempty"`
	TranscriptSummary             string                                                             `json:"transcript_summary" xml:"transcript_summary" required:"true"`
	CallSummaryTitle              *string                                                            `json:"call_summary_title,omitempty" xml:"call_summary_title,omitempty"`
	Scoped                        []*ScopedAnalysisResult                                            `json:"scoped,omitempty" xml:"scoped,omitempty"`
	extraProperties               map[string]interface{}
}

func (c *ConversationHistoryAnalysisCommonModel) GetEvaluationCriteriaResults() (value map[string]*ConversationHistoryEvaluationCriteriaResultCommonModel) {
	if c == nil {
		return
	}
	return c.EvaluationCriteriaResults
}

func (c *ConversationHistoryAnalysisCommonModel) GetDataCollectionResults() (value map[string]*DataCollectionResultCommonModel) {
	if c == nil {
		return
	}
	return c.DataCollectionResults
}

func (c *ConversationHistoryAnalysisCommonModel) GetEvaluationCriteriaResultsList() (value []*ConversationHistoryEvaluationCriteriaResultCommonModel) {
	if c == nil {
		return
	}
	return c.EvaluationCriteriaResultsList
}

func (c *ConversationHistoryAnalysisCommonModel) GetDataCollectionResultsList() (value []*DataCollectionResultCommonModel) {
	if c == nil {
		return
	}
	return c.DataCollectionResultsList
}

func (c *ConversationHistoryAnalysisCommonModel) GetCallSuccessful() (value EvaluationSuccessResult) {
	if c == nil {
		return
	}
	return c.CallSuccessful
}

func (c *ConversationHistoryAnalysisCommonModel) GetCallSuccessScore() (value *float64) {
	if c == nil {
		return
	}
	return c.CallSuccessScore
}

func (c *ConversationHistoryAnalysisCommonModel) GetTranscriptSummary() (value string) {
	if c == nil {
		return
	}
	return c.TranscriptSummary
}

func (c *ConversationHistoryAnalysisCommonModel) GetCallSummaryTitle() (value *string) {
	if c == nil {
		return
	}
	return c.CallSummaryTitle
}

func (c *ConversationHistoryAnalysisCommonModel) GetScoped() (value []*ScopedAnalysisResult) {
	if c == nil {
		return
	}
	return c.Scoped
}

func (c *ConversationHistoryAnalysisCommonModel) SetEvaluationCriteriaResults(evaluationCriteriaResults map[string]*ConversationHistoryEvaluationCriteriaResultCommonModel) {
	c.EvaluationCriteriaResults = evaluationCriteriaResults
}

func (c *ConversationHistoryAnalysisCommonModel) SetDataCollectionResults(dataCollectionResults map[string]*DataCollectionResultCommonModel) {
	c.DataCollectionResults = dataCollectionResults
}

func (c *ConversationHistoryAnalysisCommonModel) SetEvaluationCriteriaResultsList(evaluationCriteriaResultsList []*ConversationHistoryEvaluationCriteriaResultCommonModel) {
	c.EvaluationCriteriaResultsList = evaluationCriteriaResultsList
}

func (c *ConversationHistoryAnalysisCommonModel) SetDataCollectionResultsList(dataCollectionResultsList []*DataCollectionResultCommonModel) {
	c.DataCollectionResultsList = dataCollectionResultsList
}

func (c *ConversationHistoryAnalysisCommonModel) SetCallSuccessful(callSuccessful EvaluationSuccessResult) {
	c.CallSuccessful = callSuccessful
}

func (c *ConversationHistoryAnalysisCommonModel) SetCallSuccessScore(callSuccessScore *float64) {
	c.CallSuccessScore = callSuccessScore
}

func (c *ConversationHistoryAnalysisCommonModel) SetTranscriptSummary(transcriptSummary string) {
	c.TranscriptSummary = transcriptSummary
}

func (c *ConversationHistoryAnalysisCommonModel) SetCallSummaryTitle(callSummaryTitle *string) {
	c.CallSummaryTitle = callSummaryTitle
}

func (c *ConversationHistoryAnalysisCommonModel) SetScoped(scoped []*ScopedAnalysisResult) {
	c.Scoped = scoped
}

func (c *ConversationHistoryAnalysisCommonModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationHistoryAnalysisCommonModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationHistoryAnalysisCommonModel to string"
	}
	return string(jsonData)
}

func (c *ConversationHistoryAnalysisCommonModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationHistoryAnalysisCommonModel) MarshalJSON() ([]byte, error) {
	type embed ConversationHistoryAnalysisCommonModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
