package agents

import (
	big "math/big"
)

var (
	getDependentsSubagentsRequestFieldPageSize = big.NewInt(1 << 0)
	getDependentsSubagentsRequestFieldCursor   = big.NewInt(1 << 1)
)

type GetDependentsSubagentsRequest struct {
	PageSize *int    `json:"-" url:"page_size,omitempty"`
	Cursor   *string `json:"-" url:"cursor,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetDependentsSubagentsRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetDependentsSubagentsRequest) SetPageSize(pageSize *int) {
	g.PageSize = pageSize
	g.require(getDependentsSubagentsRequestFieldPageSize)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetDependentsSubagentsRequest) SetCursor(cursor *string) {
	g.Cursor = cursor
	g.require(getDependentsSubagentsRequestFieldCursor)
}
