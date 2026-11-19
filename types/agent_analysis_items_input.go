package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type AgentAnalysisItemsInput struct {
	EvaluationCriteria []*AgentAnalysisItemsInputEvaluationCriteriaItem `json:"evaluation_criteria,omitempty" xml:"evaluation_criteria,omitempty"`
	DataCollection     []*AgentAnalysisItemsInputDataCollectionItem     `json:"data_collection,omitempty" xml:"data_collection,omitempty"`
	extraProperties    map[string]interface{}
}

func (a *AgentAnalysisItemsInput) GetEvaluationCriteria() (value []*AgentAnalysisItemsInputEvaluationCriteriaItem) {
	if a == nil {
		return
	}
	return a.EvaluationCriteria
}

func (a *AgentAnalysisItemsInput) GetDataCollection() (value []*AgentAnalysisItemsInputDataCollectionItem) {
	if a == nil {
		return
	}
	return a.DataCollection
}

func (a *AgentAnalysisItemsInput) SetEvaluationCriteria(evaluationCriteria []*AgentAnalysisItemsInputEvaluationCriteriaItem) {
	a.EvaluationCriteria = evaluationCriteria
}

func (a *AgentAnalysisItemsInput) SetDataCollection(dataCollection []*AgentAnalysisItemsInputDataCollectionItem) {
	a.DataCollection = dataCollection
}

func (a *AgentAnalysisItemsInput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentAnalysisItemsInput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentAnalysisItemsInput to string"
	}
	return string(jsonData)
}

func (a *AgentAnalysisItemsInput) UnmarshalJSON(data []byte) error {
	type unmarshaler AgentAnalysisItemsInput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AgentAnalysisItemsInput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentAnalysisItemsInput) MarshalJSON() ([]byte, error) {
	type embed AgentAnalysisItemsInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
