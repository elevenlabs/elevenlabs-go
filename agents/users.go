package agents

import (
	types "github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

var (
	listUsersRequestFieldAgentId             = big.NewInt(1 << 0)
	listUsersRequestFieldBranchId            = big.NewInt(1 << 1)
	listUsersRequestFieldCallStartBeforeUnix = big.NewInt(1 << 2)
	listUsersRequestFieldCallStartAfterUnix  = big.NewInt(1 << 3)
	listUsersRequestFieldSearch              = big.NewInt(1 << 4)
	listUsersRequestFieldPageSize            = big.NewInt(1 << 5)
	listUsersRequestFieldSortBy              = big.NewInt(1 << 6)
	listUsersRequestFieldSortDirection       = big.NewInt(1 << 7)
	listUsersRequestFieldCursor              = big.NewInt(1 << 8)
)

type ListUsersRequest struct {
	AgentId             *string              `json:"-" url:"agent_id,omitempty"`
	BranchId            *string              `json:"-" url:"branch_id,omitempty"`
	CallStartBeforeUnix *int                 `json:"-" url:"call_start_before_unix,omitempty"`
	CallStartAfterUnix  *int                 `json:"-" url:"call_start_after_unix,omitempty"`
	Search              *string              `json:"-" url:"search,omitempty"`
	PageSize            *int                 `json:"-" url:"page_size,omitempty"`
	SortBy              *types.UsersSortBy   `json:"-" url:"sort_by,omitempty"`
	SortDirection       *types.SortDirection `json:"-" url:"sort_direction,omitempty"`
	Cursor              *string              `json:"-" url:"cursor,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListUsersRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetAgentId sets the AgentId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListUsersRequest) SetAgentId(agentId *string) {
	l.AgentId = agentId
	l.require(listUsersRequestFieldAgentId)
}

// SetBranchId sets the BranchId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListUsersRequest) SetBranchId(branchId *string) {
	l.BranchId = branchId
	l.require(listUsersRequestFieldBranchId)
}

// SetCallStartBeforeUnix sets the CallStartBeforeUnix field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListUsersRequest) SetCallStartBeforeUnix(callStartBeforeUnix *int) {
	l.CallStartBeforeUnix = callStartBeforeUnix
	l.require(listUsersRequestFieldCallStartBeforeUnix)
}

// SetCallStartAfterUnix sets the CallStartAfterUnix field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListUsersRequest) SetCallStartAfterUnix(callStartAfterUnix *int) {
	l.CallStartAfterUnix = callStartAfterUnix
	l.require(listUsersRequestFieldCallStartAfterUnix)
}

// SetSearch sets the Search field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListUsersRequest) SetSearch(search *string) {
	l.Search = search
	l.require(listUsersRequestFieldSearch)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListUsersRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listUsersRequestFieldPageSize)
}

// SetSortBy sets the SortBy field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListUsersRequest) SetSortBy(sortBy *types.UsersSortBy) {
	l.SortBy = sortBy
	l.require(listUsersRequestFieldSortBy)
}

// SetSortDirection sets the SortDirection field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListUsersRequest) SetSortDirection(sortDirection *types.SortDirection) {
	l.SortDirection = sortDirection
	l.require(listUsersRequestFieldSortDirection)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListUsersRequest) SetCursor(cursor *string) {
	l.Cursor = cursor
	l.require(listUsersRequestFieldCursor)
}
