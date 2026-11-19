package outboundcampaigns

import (
	"fmt"
	types "github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

type ListLeadsRequestStatus string

const (
	ListLeadsRequestStatusInProgress ListLeadsRequestStatus = "in_progress"
	ListLeadsRequestStatusFinished   ListLeadsRequestStatus = "finished"
)

func NewListLeadsRequestStatusFromString(s string) (ListLeadsRequestStatus, error) {
	switch s {
	case "in_progress":
		return ListLeadsRequestStatusInProgress, nil
	case "finished":
		return ListLeadsRequestStatusFinished, nil
	}
	var t ListLeadsRequestStatus
	return t, fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ListLeadsRequestStatus) Ptr() *ListLeadsRequestStatus {
	return &e
}

var (
	listLeadsRequestFieldPageSize    = big.NewInt(1 << 0)
	listLeadsRequestFieldSearch      = big.NewInt(1 << 1)
	listLeadsRequestFieldStatus      = big.NewInt(1 << 2)
	listLeadsRequestFieldDisposition = big.NewInt(1 << 3)
	listLeadsRequestFieldCursor      = big.NewInt(1 << 4)
)

type ListLeadsRequest struct {
	PageSize    *int                         `json:"-" url:"page_size,omitempty"`
	Search      *string                      `json:"-" url:"search,omitempty"`
	Status      *ListLeadsRequestStatus      `json:"-" url:"status,omitempty"`
	Disposition []*types.TerminalDisposition `json:"-" url:"disposition,omitempty"`
	Cursor      *string                      `json:"-" url:"cursor,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListLeadsRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListLeadsRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listLeadsRequestFieldPageSize)
}

// SetSearch sets the Search field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListLeadsRequest) SetSearch(search *string) {
	l.Search = search
	l.require(listLeadsRequestFieldSearch)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListLeadsRequest) SetStatus(status *ListLeadsRequestStatus) {
	l.Status = status
	l.require(listLeadsRequestFieldStatus)
}

// SetDisposition sets the Disposition field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListLeadsRequest) SetDisposition(disposition []*types.TerminalDisposition) {
	l.Disposition = disposition
	l.require(listLeadsRequestFieldDisposition)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListLeadsRequest) SetCursor(cursor *string) {
	l.Cursor = cursor
	l.require(listLeadsRequestFieldCursor)
}
