package types

import (
	"fmt"
)

// How an order reaches the customer.
type OrderFulfillmentType string

const (
	OrderFulfillmentTypePickup   OrderFulfillmentType = "pickup"
	OrderFulfillmentTypeDelivery OrderFulfillmentType = "delivery"
	OrderFulfillmentTypeDigital  OrderFulfillmentType = "digital"
)

func NewOrderFulfillmentTypeFromString(s string) (OrderFulfillmentType, error) {
	switch s {
	case "pickup":
		return OrderFulfillmentTypePickup, nil
	case "delivery":
		return OrderFulfillmentTypeDelivery, nil
	case "digital":
		return OrderFulfillmentTypeDigital, nil
	}
	var t OrderFulfillmentType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e OrderFulfillmentType) Ptr() *OrderFulfillmentType {
	return &e
}
