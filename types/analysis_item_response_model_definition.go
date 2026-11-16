package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type AnalysisItemResponseModelDefinition struct {
	Kind           string
	DataCollection *UserDataCollectionAnalysisItemOutput
	Evaluation     *UserEvaluationAnalysisItem
}

func (a *AnalysisItemResponseModelDefinition) GetDataCollection() *UserDataCollectionAnalysisItemOutput {
	if a == nil {
		return nil
	}
	if a.DataCollection == nil {
		return nil
	}
	return a.DataCollection
}

func (a *AnalysisItemResponseModelDefinition) GetEvaluation() *UserEvaluationAnalysisItem {
	if a == nil {
		return nil
	}
	if a.Evaluation == nil {
		return nil
	}
	return a.Evaluation
}

func (a *AnalysisItemResponseModelDefinition) GetKind() string {
	if a == nil {
		return ""
	}
	return a.Kind
}

func (a *AnalysisItemResponseModelDefinition) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Kind string `json:"kind"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	a.Kind = unmarshaler.Kind
	if unmarshaler.Kind == "" {
		return fmt.Errorf("%T did not include discriminant kind", a)
	}
	switch unmarshaler.Kind {
	case "data_collection":
		value := new(UserDataCollectionAnalysisItemOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.DataCollection = value
	case "evaluation":
		value := new(UserEvaluationAnalysisItem)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.Evaluation = value
	}
	return nil
}

func (a AnalysisItemResponseModelDefinition) MarshalJSON() ([]byte, error) {
	if err := a.validate(); err != nil {
		return nil, err
	}
	if a.DataCollection != nil {
		return internal.MarshalJSONWithDiscriminant(a.DataCollection, "kind", "data_collection")
	}
	if a.Evaluation != nil {
		return internal.MarshalJSONWithDiscriminant(a.Evaluation, "kind", "evaluation")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", a)
}

type AnalysisItemResponseModelDefinitionVisitor interface {
	VisitDataCollection(*UserDataCollectionAnalysisItemOutput) error
	VisitEvaluation(*UserEvaluationAnalysisItem) error
}

func (a *AnalysisItemResponseModelDefinition) Accept(visitor AnalysisItemResponseModelDefinitionVisitor) error {
	if a.DataCollection != nil {
		return visitor.VisitDataCollection(a.DataCollection)
	}
	if a.Evaluation != nil {
		return visitor.VisitEvaluation(a.Evaluation)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", a)
}

func (a *AnalysisItemResponseModelDefinition) validate() error {
	if a == nil {
		return fmt.Errorf("type %T is nil", a)
	}
	var fields []string
	if a.DataCollection != nil {
		fields = append(fields, "data_collection")
	}
	if a.Evaluation != nil {
		fields = append(fields, "evaluation")
	}
	if len(fields) == 0 {
		if a.Kind != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", a, a.Kind)
		}
		return fmt.Errorf("type %T is empty", a)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", a, fields)
	}
	if a.Kind != "" {
		field := fields[0]
		if a.Kind != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				a,
				a.Kind,
				a,
			)
		}
	}
	return nil
}
