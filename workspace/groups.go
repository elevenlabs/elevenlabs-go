package workspace

import (
	big "math/big"
)

var (
	searchGroupsRequestFieldName = big.NewInt(1 << 0)
)

type SearchGroupsRequest struct {
	Name string `json:"-" url:"name,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (s *SearchGroupsRequest) require(field *big.Int) {
	if s.explicitFields == nil {
		s.explicitFields = big.NewInt(0)
	}
	s.explicitFields.Or(s.explicitFields, field)
}

// SetName sets the Name field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SearchGroupsRequest) SetName(name string) {
	s.Name = name
	s.require(searchGroupsRequestFieldName)
}
