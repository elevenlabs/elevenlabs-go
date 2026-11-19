package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type AnalysisItemVersionResponseModelDefinition struct {
	Kind           string
	DataCollection *UserDataCollectionAnalysisItemOutput
	Evaluation     *UserEvaluationAnalysisItem
}

func (analysisItemVersionResponseModelDefinition *AnalysisItemVersionResponseModelDefinition) GetDataCollection() *UserDataCollectionAnalysisItemOutput {
	if analysisItemVersionResponseModelDefinition == nil {
		return nil
	}
	if analysisItemVersionResponseModelDefinition.DataCollection == nil {
		return nil
	}
	return analysisItemVersionResponseModelDefinition.DataCollection
}

func (analysisItemVersionResponseModelDefinition *AnalysisItemVersionResponseModelDefinition) GetEvaluation() *UserEvaluationAnalysisItem {
	if analysisItemVersionResponseModelDefinition == nil {
		return nil
	}
	if analysisItemVersionResponseModelDefinition.Evaluation == nil {
		return nil
	}
	return analysisItemVersionResponseModelDefinition.Evaluation
}

func (analysisItemVersionResponseModelDefinition *AnalysisItemVersionResponseModelDefinition) GetKind() string {
	if analysisItemVersionResponseModelDefinition == nil {
		return ""
	}
	return analysisItemVersionResponseModelDefinition.Kind
}

func (analysisItemVersionResponseModelDefinition *AnalysisItemVersionResponseModelDefinition) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Kind string `json:"kind"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	analysisItemVersionResponseModelDefinition.Kind = unmarshaler.Kind
	if unmarshaler.Kind == "" {
		return fmt.Errorf("%T did not include discriminant kind", analysisItemVersionResponseModelDefinition)
	}
	switch unmarshaler.Kind {
	case "data_collection":
		value := new(UserDataCollectionAnalysisItemOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		analysisItemVersionResponseModelDefinition.DataCollection = value
	case "evaluation":
		value := new(UserEvaluationAnalysisItem)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		analysisItemVersionResponseModelDefinition.Evaluation = value
	}
	return nil
}

func (analysisItemVersionResponseModelDefinition AnalysisItemVersionResponseModelDefinition) MarshalJSON() ([]byte, error) {
	if err := analysisItemVersionResponseModelDefinition.validate(); err != nil {
		return nil, err
	}
	if analysisItemVersionResponseModelDefinition.DataCollection != nil {
		return internal.MarshalJSONWithDiscriminant(analysisItemVersionResponseModelDefinition.DataCollection, "kind", "data_collection")
	}
	if analysisItemVersionResponseModelDefinition.Evaluation != nil {
		return internal.MarshalJSONWithDiscriminant(analysisItemVersionResponseModelDefinition.Evaluation, "kind", "evaluation")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", analysisItemVersionResponseModelDefinition)
}

type AnalysisItemVersionResponseModelDefinitionVisitor interface {
	VisitDataCollection(*UserDataCollectionAnalysisItemOutput) error
	VisitEvaluation(*UserEvaluationAnalysisItem) error
}

func (analysisItemVersionResponseModelDefinition *AnalysisItemVersionResponseModelDefinition) Accept(visitor AnalysisItemVersionResponseModelDefinitionVisitor) error {
	if analysisItemVersionResponseModelDefinition.DataCollection != nil {
		return visitor.VisitDataCollection(analysisItemVersionResponseModelDefinition.DataCollection)
	}
	if analysisItemVersionResponseModelDefinition.Evaluation != nil {
		return visitor.VisitEvaluation(analysisItemVersionResponseModelDefinition.Evaluation)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", analysisItemVersionResponseModelDefinition)
}

func (analysisItemVersionResponseModelDefinition *AnalysisItemVersionResponseModelDefinition) validate() error {
	if analysisItemVersionResponseModelDefinition == nil {
		return fmt.Errorf("type %T is nil", analysisItemVersionResponseModelDefinition)
	}
	var fields []string
	if analysisItemVersionResponseModelDefinition.DataCollection != nil {
		fields = append(fields, "data_collection")
	}
	if analysisItemVersionResponseModelDefinition.Evaluation != nil {
		fields = append(fields, "evaluation")
	}
	if len(fields) == 0 {
		if analysisItemVersionResponseModelDefinition.Kind != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", analysisItemVersionResponseModelDefinition, analysisItemVersionResponseModelDefinition.Kind)
		}
		return fmt.Errorf("type %T is empty", analysisItemVersionResponseModelDefinition)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", analysisItemVersionResponseModelDefinition, fields)
	}
	if analysisItemVersionResponseModelDefinition.Kind != "" {
		field := fields[0]
		if analysisItemVersionResponseModelDefinition.Kind != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				analysisItemVersionResponseModelDefinition,
				analysisItemVersionResponseModelDefinition.Kind,
				analysisItemVersionResponseModelDefinition,
			)
		}
	}
	return nil
}
