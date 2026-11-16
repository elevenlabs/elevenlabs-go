package agents

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

type SummariesGetOkResponse1 struct {
	Status  string
	Failure *types.BatchFailureResponseModel
	Success *types.AgentSummaryBatchSuccessfulResponseModel
}

func (s *SummariesGetOkResponse1) GetFailure() *types.BatchFailureResponseModel {
	if s == nil {
		return nil
	}
	if s.Failure == nil {
		return nil
	}
	return s.Failure
}

func (s *SummariesGetOkResponse1) GetSuccess() *types.AgentSummaryBatchSuccessfulResponseModel {
	if s == nil {
		return nil
	}
	if s.Success == nil {
		return nil
	}
	return s.Success
}

func (s *SummariesGetOkResponse1) GetStatus() string {
	if s == nil {
		return ""
	}
	return s.Status
}

func (s *SummariesGetOkResponse1) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Status string `json:"status"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	s.Status = unmarshaler.Status
	if unmarshaler.Status == "" {
		return fmt.Errorf("%T did not include discriminant status", s)
	}
	switch unmarshaler.Status {
	case "failure":
		value := new(types.BatchFailureResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.Failure = value
	case "success":
		value := new(types.AgentSummaryBatchSuccessfulResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		s.Success = value
	}
	return nil
}

func (s SummariesGetOkResponse1) MarshalJSON() ([]byte, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	if s.Failure != nil {
		return internal.MarshalJSONWithDiscriminant(s.Failure, "status", "failure")
	}
	if s.Success != nil {
		return internal.MarshalJSONWithDiscriminant(s.Success, "status", "success")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", s)
}

type SummariesGetOkResponse1Visitor interface {
	VisitFailure(*types.BatchFailureResponseModel) error
	VisitSuccess(*types.AgentSummaryBatchSuccessfulResponseModel) error
}

func (s *SummariesGetOkResponse1) Accept(visitor SummariesGetOkResponse1Visitor) error {
	if s.Failure != nil {
		return visitor.VisitFailure(s.Failure)
	}
	if s.Success != nil {
		return visitor.VisitSuccess(s.Success)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", s)
}

func (s *SummariesGetOkResponse1) validate() error {
	if s == nil {
		return fmt.Errorf("type %T is nil", s)
	}
	var fields []string
	if s.Failure != nil {
		fields = append(fields, "failure")
	}
	if s.Success != nil {
		fields = append(fields, "success")
	}
	if len(fields) == 0 {
		if s.Status != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", s, s.Status)
		}
		return fmt.Errorf("type %T is empty", s)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", s, fields)
	}
	if s.Status != "" {
		field := fields[0]
		if s.Status != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				s,
				s.Status,
				s,
			)
		}
	}
	return nil
}

var (
	getSummariesRequestFieldAgentIds = big.NewInt(1 << 0)
)

type GetSummariesRequest struct {
	AgentIds []*string `json:"-" url:"agent_ids,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetSummariesRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetAgentIds sets the AgentIds field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSummariesRequest) SetAgentIds(agentIds []*string) {
	g.AgentIds = agentIds
	g.require(getSummariesRequestFieldAgentIds)
}
