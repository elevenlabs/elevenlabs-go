package types

import (
	"fmt"
)

// Where a classic service is delivered.
type ServiceLocationType string

const (
	ServiceLocationTypeBusiness ServiceLocationType = "business"
	ServiceLocationTypeCustomer ServiceLocationType = "customer"
)

func NewServiceLocationTypeFromString(s string) (ServiceLocationType, error) {
	switch s {
	case "business":
		return ServiceLocationTypeBusiness, nil
	case "customer":
		return ServiceLocationTypeCustomer, nil
	}
	var t ServiceLocationType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ServiceLocationType) Ptr() *ServiceLocationType {
	return &e
}
