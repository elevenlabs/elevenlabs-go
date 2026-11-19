package documents

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

type GetSummariesResponse struct {
	Status  string
	Failure *types.BatchFailureResponseModel
	Success *types.KnowledgeBaseSummaryBatchSuccessfulResponseModel
}

func (g *GetSummariesResponse) GetFailure() *types.BatchFailureResponseModel {
	if g == nil {
		return nil
	}
	if g.Failure == nil {
		return nil
	}
	return g.Failure
}

func (g *GetSummariesResponse) GetSuccess() *types.KnowledgeBaseSummaryBatchSuccessfulResponseModel {
	if g == nil {
		return nil
	}
	if g.Success == nil {
		return nil
	}
	return g.Success
}

func (g *GetSummariesResponse) GetStatus() string {
	if g == nil {
		return ""
	}
	return g.Status
}

func (g *GetSummariesResponse) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Status string `json:"status"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	g.Status = unmarshaler.Status
	if unmarshaler.Status == "" {
		return fmt.Errorf("%T did not include discriminant status", g)
	}
	switch unmarshaler.Status {
	case "failure":
		value := new(types.BatchFailureResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		g.Failure = value
	case "success":
		value := new(types.KnowledgeBaseSummaryBatchSuccessfulResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		g.Success = value
	}
	return nil
}

func (g GetSummariesResponse) MarshalJSON() ([]byte, error) {
	if err := g.validate(); err != nil {
		return nil, err
	}
	if g.Failure != nil {
		return internal.MarshalJSONWithDiscriminant(g.Failure, "status", "failure")
	}
	if g.Success != nil {
		return internal.MarshalJSONWithDiscriminant(g.Success, "status", "success")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", g)
}

type GetSummariesResponseVisitor interface {
	VisitFailure(*types.BatchFailureResponseModel) error
	VisitSuccess(*types.KnowledgeBaseSummaryBatchSuccessfulResponseModel) error
}

func (g *GetSummariesResponse) Accept(visitor GetSummariesResponseVisitor) error {
	if g.Failure != nil {
		return visitor.VisitFailure(g.Failure)
	}
	if g.Success != nil {
		return visitor.VisitSuccess(g.Success)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", g)
}

func (g *GetSummariesResponse) validate() error {
	if g == nil {
		return fmt.Errorf("type %T is nil", g)
	}
	var fields []string
	if g.Failure != nil {
		fields = append(fields, "failure")
	}
	if g.Success != nil {
		fields = append(fields, "success")
	}
	if len(fields) == 0 {
		if g.Status != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", g, g.Status)
		}
		return fmt.Errorf("type %T is empty", g)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", g, fields)
	}
	if g.Status != "" {
		field := fields[0]
		if g.Status != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				g,
				g.Status,
				g,
			)
		}
	}
	return nil
}

var (
	getSummariesRequestFieldDocumentIds = big.NewInt(1 << 0)
)

type GetSummariesRequest struct {
	DocumentIds []*string `json:"-" url:"document_ids,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetSummariesRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetDocumentIds sets the DocumentIds field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSummariesRequest) SetDocumentIds(documentIds []*string) {
	g.DocumentIds = documentIds
	g.require(getSummariesRequestFieldDocumentIds)
}
