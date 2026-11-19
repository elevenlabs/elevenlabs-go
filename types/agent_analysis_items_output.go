package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type AgentAnalysisItemsOutput struct {
	EvaluationCriteria []*AgentAnalysisItemsOutputEvaluationCriteriaItem `json:"evaluation_criteria,omitempty" xml:"evaluation_criteria,omitempty"`
	DataCollection     []*AgentAnalysisItemsOutputDataCollectionItem     `json:"data_collection,omitempty" xml:"data_collection,omitempty"`
	extraProperties    map[string]interface{}
}

func (a *AgentAnalysisItemsOutput) GetEvaluationCriteria() (value []*AgentAnalysisItemsOutputEvaluationCriteriaItem) {
	if a == nil {
		return
	}
	return a.EvaluationCriteria
}

func (a *AgentAnalysisItemsOutput) GetDataCollection() (value []*AgentAnalysisItemsOutputDataCollectionItem) {
	if a == nil {
		return
	}
	return a.DataCollection
}

func (a *AgentAnalysisItemsOutput) SetEvaluationCriteria(evaluationCriteria []*AgentAnalysisItemsOutputEvaluationCriteriaItem) {
	a.EvaluationCriteria = evaluationCriteria
}

func (a *AgentAnalysisItemsOutput) SetDataCollection(dataCollection []*AgentAnalysisItemsOutputDataCollectionItem) {
	a.DataCollection = dataCollection
}

func (a *AgentAnalysisItemsOutput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentAnalysisItemsOutput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentAnalysisItemsOutput to string"
	}
	return string(jsonData)
}

func (a *AgentAnalysisItemsOutput) UnmarshalJSON(data []byte) error {
	type unmarshaler AgentAnalysisItemsOutput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AgentAnalysisItemsOutput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentAnalysisItemsOutput) MarshalJSON() ([]byte, error) {
	type embed AgentAnalysisItemsOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
