package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type AgentMergeProposalResponseOutcome struct {
	Status string
	Closed *ClosedOutcome
	Merged *MergedOutcome
	Open   *OpenOutcome
}

func (a *AgentMergeProposalResponseOutcome) GetClosed() *ClosedOutcome {
	if a == nil {
		return nil
	}
	if a.Closed == nil {
		return nil
	}
	return a.Closed
}

func (a *AgentMergeProposalResponseOutcome) GetMerged() *MergedOutcome {
	if a == nil {
		return nil
	}
	if a.Merged == nil {
		return nil
	}
	return a.Merged
}

func (a *AgentMergeProposalResponseOutcome) GetOpen() *OpenOutcome {
	if a == nil {
		return nil
	}
	if a.Open == nil {
		return nil
	}
	return a.Open
}

func (a *AgentMergeProposalResponseOutcome) GetStatus() string {
	if a == nil {
		return ""
	}
	return a.Status
}

func (a *AgentMergeProposalResponseOutcome) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Status string `json:"status"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	a.Status = unmarshaler.Status
	if unmarshaler.Status == "" {
		return fmt.Errorf("%T did not include discriminant status", a)
	}
	switch unmarshaler.Status {
	case "closed":
		value := new(ClosedOutcome)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.Closed = value
	case "merged":
		value := new(MergedOutcome)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.Merged = value
	case "open":
		value := new(OpenOutcome)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.Open = value
	}
	return nil
}

func (a AgentMergeProposalResponseOutcome) MarshalJSON() ([]byte, error) {
	if err := a.validate(); err != nil {
		return nil, err
	}
	if a.Closed != nil {
		return internal.MarshalJSONWithDiscriminant(a.Closed, "status", "closed")
	}
	if a.Merged != nil {
		return internal.MarshalJSONWithDiscriminant(a.Merged, "status", "merged")
	}
	if a.Open != nil {
		return internal.MarshalJSONWithDiscriminant(a.Open, "status", "open")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", a)
}

type AgentMergeProposalResponseOutcomeVisitor interface {
	VisitClosed(*ClosedOutcome) error
	VisitMerged(*MergedOutcome) error
	VisitOpen(*OpenOutcome) error
}

func (a *AgentMergeProposalResponseOutcome) Accept(visitor AgentMergeProposalResponseOutcomeVisitor) error {
	if a.Closed != nil {
		return visitor.VisitClosed(a.Closed)
	}
	if a.Merged != nil {
		return visitor.VisitMerged(a.Merged)
	}
	if a.Open != nil {
		return visitor.VisitOpen(a.Open)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", a)
}

func (a *AgentMergeProposalResponseOutcome) validate() error {
	if a == nil {
		return fmt.Errorf("type %T is nil", a)
	}
	var fields []string
	if a.Closed != nil {
		fields = append(fields, "closed")
	}
	if a.Merged != nil {
		fields = append(fields, "merged")
	}
	if a.Open != nil {
		fields = append(fields, "open")
	}
	if len(fields) == 0 {
		if a.Status != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", a, a.Status)
		}
		return fmt.Errorf("type %T is empty", a)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", a, fields)
	}
	if a.Status != "" {
		field := fields[0]
		if a.Status != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				a,
				a.Status,
				a,
			)
		}
	}
	return nil
}
