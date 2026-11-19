package mcpservers

import (
	big "math/big"
)

var (
	listToolsRequestFieldEnvironment = big.NewInt(1 << 0)
)

type ListToolsRequest struct {
	Environment *string `json:"-" url:"environment,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListToolsRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetEnvironment sets the Environment field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListToolsRequest) SetEnvironment(environment *string) {
	l.Environment = environment
	l.require(listToolsRequestFieldEnvironment)
}
