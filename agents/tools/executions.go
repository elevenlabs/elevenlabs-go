package tools

import (
	big "math/big"
)

var (
	getExecutionsRequestFieldCursor    = big.NewInt(1 << 0)
	getExecutionsRequestFieldPageSize  = big.NewInt(1 << 1)
	getExecutionsRequestFieldIsError   = big.NewInt(1 << 2)
	getExecutionsRequestFieldAgentId   = big.NewInt(1 << 3)
	getExecutionsRequestFieldBranchId  = big.NewInt(1 << 4)
	getExecutionsRequestFieldStartTime = big.NewInt(1 << 5)
	getExecutionsRequestFieldEndTime   = big.NewInt(1 << 6)
)

type GetExecutionsRequest struct {
	Cursor    *string  `json:"-" url:"cursor,omitempty"`
	PageSize  *int     `json:"-" url:"page_size,omitempty"`
	IsError   *bool    `json:"-" url:"is_error,omitempty"`
	AgentId   *string  `json:"-" url:"agent_id,omitempty"`
	BranchId  *string  `json:"-" url:"branch_id,omitempty"`
	StartTime *float64 `json:"-" url:"start_time,omitempty"`
	EndTime   *float64 `json:"-" url:"end_time,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetExecutionsRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetExecutionsRequest) SetCursor(cursor *string) {
	g.Cursor = cursor
	g.require(getExecutionsRequestFieldCursor)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetExecutionsRequest) SetPageSize(pageSize *int) {
	g.PageSize = pageSize
	g.require(getExecutionsRequestFieldPageSize)
}

// SetIsError sets the IsError field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetExecutionsRequest) SetIsError(isError *bool) {
	g.IsError = isError
	g.require(getExecutionsRequestFieldIsError)
}

// SetAgentId sets the AgentId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetExecutionsRequest) SetAgentId(agentId *string) {
	g.AgentId = agentId
	g.require(getExecutionsRequestFieldAgentId)
}

// SetBranchId sets the BranchId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetExecutionsRequest) SetBranchId(branchId *string) {
	g.BranchId = branchId
	g.require(getExecutionsRequestFieldBranchId)
}

// SetStartTime sets the StartTime field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetExecutionsRequest) SetStartTime(startTime *float64) {
	g.StartTime = startTime
	g.require(getExecutionsRequestFieldStartTime)
}

// SetEndTime sets the EndTime field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetExecutionsRequest) SetEndTime(endTime *float64) {
	g.EndTime = endTime
	g.require(getExecutionsRequestFieldEndTime)
}
