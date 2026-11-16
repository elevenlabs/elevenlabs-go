package reception

import (
	big "math/big"
)

var (
	listInvoicesBillingRequestFieldCursor   = big.NewInt(1 << 0)
	listInvoicesBillingRequestFieldPageSize = big.NewInt(1 << 1)
)

type ListInvoicesBillingRequest struct {
	Cursor   *string `json:"-" url:"cursor,omitempty"`
	PageSize *int    `json:"-" url:"page_size,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListInvoicesBillingRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListInvoicesBillingRequest) SetCursor(cursor *string) {
	l.Cursor = cursor
	l.require(listInvoicesBillingRequestFieldCursor)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListInvoicesBillingRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listInvoicesBillingRequestFieldPageSize)
}
