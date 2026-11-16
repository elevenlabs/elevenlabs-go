package flows

import (
	"fmt"
	big "math/big"
)

type ListVideoRequestStatus string

const (
	ListVideoRequestStatusPending    ListVideoRequestStatus = "pending"
	ListVideoRequestStatusGenerating ListVideoRequestStatus = "generating"
	ListVideoRequestStatusCompleted  ListVideoRequestStatus = "completed"
	ListVideoRequestStatusFailed     ListVideoRequestStatus = "failed"
)

func NewListVideoRequestStatusFromString(s string) (ListVideoRequestStatus, error) {
	switch s {
	case "pending":
		return ListVideoRequestStatusPending, nil
	case "generating":
		return ListVideoRequestStatusGenerating, nil
	case "completed":
		return ListVideoRequestStatusCompleted, nil
	case "failed":
		return ListVideoRequestStatusFailed, nil
	}
	var t ListVideoRequestStatus
	return t, fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ListVideoRequestStatus) Ptr() *ListVideoRequestStatus {
	return &e
}

var (
	listVideoRequestFieldCursor   = big.NewInt(1 << 0)
	listVideoRequestFieldPageSize = big.NewInt(1 << 1)
	listVideoRequestFieldStatus   = big.NewInt(1 << 2)
	listVideoRequestFieldModelId  = big.NewInt(1 << 3)
)

type ListVideoRequest struct {
	Cursor   *string                 `json:"-" url:"cursor,omitempty"`
	PageSize *int                    `json:"-" url:"page_size,omitempty"`
	Status   *ListVideoRequestStatus `json:"-" url:"status,omitempty"`
	ModelId  *string                 `json:"-" url:"model_id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListVideoRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListVideoRequest) SetCursor(cursor *string) {
	l.Cursor = cursor
	l.require(listVideoRequestFieldCursor)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListVideoRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listVideoRequestFieldPageSize)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListVideoRequest) SetStatus(status *ListVideoRequestStatus) {
	l.Status = status
	l.require(listVideoRequestFieldStatus)
}

// SetModelId sets the ModelId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListVideoRequest) SetModelId(modelId *string) {
	l.ModelId = modelId
	l.require(listVideoRequestFieldModelId)
}
