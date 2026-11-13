package flows

import (
	"fmt"
	big "math/big"
)

type ListTextToSpeechRequestStatus string

const (
	ListTextToSpeechRequestStatusPending    ListTextToSpeechRequestStatus = "pending"
	ListTextToSpeechRequestStatusGenerating ListTextToSpeechRequestStatus = "generating"
	ListTextToSpeechRequestStatusCompleted  ListTextToSpeechRequestStatus = "completed"
	ListTextToSpeechRequestStatusFailed     ListTextToSpeechRequestStatus = "failed"
)

func NewListTextToSpeechRequestStatusFromString(s string) (ListTextToSpeechRequestStatus, error) {
	switch s {
	case "pending":
		return ListTextToSpeechRequestStatusPending, nil
	case "generating":
		return ListTextToSpeechRequestStatusGenerating, nil
	case "completed":
		return ListTextToSpeechRequestStatusCompleted, nil
	case "failed":
		return ListTextToSpeechRequestStatusFailed, nil
	}
	var t ListTextToSpeechRequestStatus
	return t, fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ListTextToSpeechRequestStatus) Ptr() *ListTextToSpeechRequestStatus {
	return &e
}

var (
	listTextToSpeechRequestFieldCursor   = big.NewInt(1 << 0)
	listTextToSpeechRequestFieldPageSize = big.NewInt(1 << 1)
	listTextToSpeechRequestFieldStatus   = big.NewInt(1 << 2)
	listTextToSpeechRequestFieldModelId  = big.NewInt(1 << 3)
)

type ListTextToSpeechRequest struct {
	Cursor   *string                        `json:"-" url:"cursor,omitempty"`
	PageSize *int                           `json:"-" url:"page_size,omitempty"`
	Status   *ListTextToSpeechRequestStatus `json:"-" url:"status,omitempty"`
	ModelId  *string                        `json:"-" url:"model_id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListTextToSpeechRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListTextToSpeechRequest) SetCursor(cursor *string) {
	l.Cursor = cursor
	l.require(listTextToSpeechRequestFieldCursor)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListTextToSpeechRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listTextToSpeechRequestFieldPageSize)
}

// SetStatus sets the Status field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListTextToSpeechRequest) SetStatus(status *ListTextToSpeechRequestStatus) {
	l.Status = status
	l.require(listTextToSpeechRequestFieldStatus)
}

// SetModelId sets the ModelId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListTextToSpeechRequest) SetModelId(modelId *string) {
	l.ModelId = modelId
	l.require(listTextToSpeechRequestFieldModelId)
}
