package flows

import (
	"fmt"
	big "math/big"
)

type ListImageRequestStatus string

const (
	ListImageRequestStatusPending    ListImageRequestStatus = "pending"
	ListImageRequestStatusGenerating ListImageRequestStatus = "generating"
	ListImageRequestStatusCompleted  ListImageRequestStatus = "completed"
	ListImageRequestStatusFailed     ListImageRequestStatus = "failed"
)

func NewListImageRequestStatusFromString(s string) (ListImageRequestStatus, error) {
	switch s {
	case "pending":
		return ListImageRequestStatusPending, nil
	case "generating":
		return ListImageRequestStatusGenerating, nil
	case "completed":
		return ListImageRequestStatusCompleted, nil
	case "failed":
		return ListImageRequestStatusFailed, nil
	}
	var t ListImageRequestStatus
	return t, fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ListImageRequestStatus) Ptr() *ListImageRequestStatus {
	return &e
}

var (
	listImageRequestFieldCursor   = big.NewInt(1 << 0)
	listImageRequestFieldPageSize = big.NewInt(1 << 1)
	listImageRequestFieldStatus   = big.NewInt(1 << 2)
	listImageRequestFieldModelId  = big.NewInt(1 << 3)
)

type ListImageRequest struct {
	Cursor   *string                 `json:"-" url:"cursor,omitempty"`
	PageSize *int                    `json:"-" url:"page_size,omitempty"`
	Status   *ListImageRequestStatus `json:"-" url:"status,omitempty"`
	ModelId  *string                 `json:"-" url:"model_id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListImageRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListImageRequest) SetCursor(cursor *string) {
	l.Cursor = cursor
	l.require(listImageRequestFieldCursor)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListImageRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listImageRequestFieldPageSize)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListImageRequest) SetStatus(status *ListImageRequestStatus) {
	l.Status = status
	l.require(listImageRequestFieldStatus)
}

// SetModelId sets the ModelId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListImageRequest) SetModelId(modelId *string) {
	l.ModelId = modelId
	l.require(listImageRequestFieldModelId)
}
