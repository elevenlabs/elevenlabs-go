package types

import (
	"fmt"
)

type TicketSearchMatchField string

const (
	TicketSearchMatchFieldTitle       TicketSearchMatchField = "title"
	TicketSearchMatchFieldDescription TicketSearchMatchField = "description"
	TicketSearchMatchFieldComment     TicketSearchMatchField = "comment"
	TicketSearchMatchFieldTurnComment TicketSearchMatchField = "turn_comment"
)

func NewTicketSearchMatchFieldFromString(s string) (TicketSearchMatchField, error) {
	switch s {
	case "title":
		return TicketSearchMatchFieldTitle, nil
	case "description":
		return TicketSearchMatchFieldDescription, nil
	case "comment":
		return TicketSearchMatchFieldComment, nil
	case "turn_comment":
		return TicketSearchMatchFieldTurnComment, nil
	}
	var t TicketSearchMatchField
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TicketSearchMatchField) Ptr() *TicketSearchMatchField {
	return &e
}
