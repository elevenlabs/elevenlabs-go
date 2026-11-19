package outboundcampaigns

import (
	"fmt"
	types "github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

type ListCallsRequestStage string

const (
	ListCallsRequestStageQueued    ListCallsRequestStage = "queued"
	ListCallsRequestStageSent      ListCallsRequestStage = "sent"
	ListCallsRequestStageDelivered ListCallsRequestStage = "delivered"
	ListCallsRequestStageRead      ListCallsRequestStage = "read"
	ListCallsRequestStageReplied   ListCallsRequestStage = "replied"
	ListCallsRequestStageFailed    ListCallsRequestStage = "failed"
)

func NewListCallsRequestStageFromString(s string) (ListCallsRequestStage, error) {
	switch s {
	case "queued":
		return ListCallsRequestStageQueued, nil
	case "sent":
		return ListCallsRequestStageSent, nil
	case "delivered":
		return ListCallsRequestStageDelivered, nil
	case "read":
		return ListCallsRequestStageRead, nil
	case "replied":
		return ListCallsRequestStageReplied, nil
	case "failed":
		return ListCallsRequestStageFailed, nil
	}
	var t ListCallsRequestStage
	return t, fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ListCallsRequestStage) Ptr() *ListCallsRequestStage {
	return &e
}

var (
	listCallsRequestFieldPageSize   = big.NewInt(1 << 0)
	listCallsRequestFieldLeadId     = big.NewInt(1 << 1)
	listCallsRequestFieldStatus     = big.NewInt(1 << 2)
	listCallsRequestFieldCallResult = big.NewInt(1 << 3)
	listCallsRequestFieldStage      = big.NewInt(1 << 4)
	listCallsRequestFieldCursor     = big.NewInt(1 << 5)
)

type ListCallsRequest struct {
	PageSize   *int                   `json:"-" url:"page_size,omitempty"`
	LeadId     *string                `json:"-" url:"lead_id,omitempty"`
	Status     *types.CallStatus      `json:"-" url:"status,omitempty"`
	CallResult []*types.CallResult    `json:"-" url:"call_result,omitempty"`
	Stage      *ListCallsRequestStage `json:"-" url:"stage,omitempty"`
	Cursor     *string                `json:"-" url:"cursor,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListCallsRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListCallsRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listCallsRequestFieldPageSize)
}

// SetLeadId sets the LeadId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListCallsRequest) SetLeadId(leadId *string) {
	l.LeadId = leadId
	l.require(listCallsRequestFieldLeadId)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListCallsRequest) SetStatus(status *types.CallStatus) {
	l.Status = status
	l.require(listCallsRequestFieldStatus)
}

// SetCallResult sets the CallResult field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListCallsRequest) SetCallResult(callResult []*types.CallResult) {
	l.CallResult = callResult
	l.require(listCallsRequestFieldCallResult)
}

// SetStage sets the Stage field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListCallsRequest) SetStage(stage *ListCallsRequestStage) {
	l.Stage = stage
	l.require(listCallsRequestFieldStage)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListCallsRequest) SetCursor(cursor *string) {
	l.Cursor = cursor
	l.require(listCallsRequestFieldCursor)
}

var (
	exportCallsRequestFieldLimit = big.NewInt(1 << 0)
)

type ExportCallsRequest struct {
	Limit *int `json:"-" url:"limit,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (e *ExportCallsRequest) require(field *big.Int) {
	if e.explicitFields == nil {
		e.explicitFields = big.NewInt(0)
	}
	e.explicitFields.Or(e.explicitFields, field)
}

// SetLimit sets the Limit field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (e *ExportCallsRequest) SetLimit(limit *int) {
	e.Limit = limit
	e.require(exportCallsRequestFieldLimit)
}
