package types

import (
	"fmt"
)

type OrderState string

const (
	OrderStateOpen       OrderState = "open"
	OrderStateSubmitted  OrderState = "submitted"
	OrderStatePaid       OrderState = "paid"
	OrderStateAccepted   OrderState = "accepted"
	OrderStateRejected   OrderState = "rejected"
	OrderStateDone       OrderState = "done"
	OrderStateCancelling OrderState = "cancelling"
	OrderStateCancelled  OrderState = "cancelled"
	OrderStateExpired    OrderState = "expired"
	OrderStateMerged     OrderState = "merged"
)

func NewOrderStateFromString(s string) (OrderState, error) {
	switch s {
	case "open":
		return OrderStateOpen, nil
	case "submitted":
		return OrderStateSubmitted, nil
	case "paid":
		return OrderStatePaid, nil
	case "accepted":
		return OrderStateAccepted, nil
	case "rejected":
		return OrderStateRejected, nil
	case "done":
		return OrderStateDone, nil
	case "cancelling":
		return OrderStateCancelling, nil
	case "cancelled":
		return OrderStateCancelled, nil
	case "expired":
		return OrderStateExpired, nil
	case "merged":
		return OrderStateMerged, nil
	}
	var t OrderState
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e OrderState) Ptr() *OrderState {
	return &e
}
