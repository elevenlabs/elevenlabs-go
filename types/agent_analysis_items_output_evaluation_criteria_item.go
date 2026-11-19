package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type AgentAnalysisItemsOutputEvaluationCriteriaItem struct {
	Source string
	System *AttachedSystemEvaluationRef
	User   *AttachedUserEvaluationRef
}

func (a *AgentAnalysisItemsOutputEvaluationCriteriaItem) GetSystem() *AttachedSystemEvaluationRef {
	if a == nil {
		return nil
	}
	if a.System == nil {
		return nil
	}
	return a.System
}

func (a *AgentAnalysisItemsOutputEvaluationCriteriaItem) GetUser() *AttachedUserEvaluationRef {
	if a == nil {
		return nil
	}
	if a.User == nil {
		return nil
	}
	return a.User
}

func (a *AgentAnalysisItemsOutputEvaluationCriteriaItem) GetSource() string {
	if a == nil {
		return ""
	}
	return a.Source
}

func (a *AgentAnalysisItemsOutputEvaluationCriteriaItem) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Source string `json:"source"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	a.Source = unmarshaler.Source
	if unmarshaler.Source == "" {
		return fmt.Errorf("%T did not include discriminant source", a)
	}
	switch unmarshaler.Source {
	case "system":
		value := new(AttachedSystemEvaluationRef)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.System = value
	case "user":
		value := new(AttachedUserEvaluationRef)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.User = value
	}
	return nil
}

func (a AgentAnalysisItemsOutputEvaluationCriteriaItem) MarshalJSON() ([]byte, error) {
	if err := a.validate(); err != nil {
		return nil, err
	}
	if a.System != nil {
		return internal.MarshalJSONWithDiscriminant(a.System, "source", "system")
	}
	if a.User != nil {
		return internal.MarshalJSONWithDiscriminant(a.User, "source", "user")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", a)
}

type AgentAnalysisItemsOutputEvaluationCriteriaItemVisitor interface {
	VisitSystem(*AttachedSystemEvaluationRef) error
	VisitUser(*AttachedUserEvaluationRef) error
}

func (a *AgentAnalysisItemsOutputEvaluationCriteriaItem) Accept(visitor AgentAnalysisItemsOutputEvaluationCriteriaItemVisitor) error {
	if a.System != nil {
		return visitor.VisitSystem(a.System)
	}
	if a.User != nil {
		return visitor.VisitUser(a.User)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", a)
}

func (a *AgentAnalysisItemsOutputEvaluationCriteriaItem) validate() error {
	if a == nil {
		return fmt.Errorf("type %T is nil", a)
	}
	var fields []string
	if a.System != nil {
		fields = append(fields, "system")
	}
	if a.User != nil {
		fields = append(fields, "user")
	}
	if len(fields) == 0 {
		if a.Source != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", a, a.Source)
		}
		return fmt.Errorf("type %T is empty", a)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", a, fields)
	}
	if a.Source != "" {
		field := fields[0]
		if a.Source != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				a,
				a.Source,
				a,
			)
		}
	}
	return nil
}
