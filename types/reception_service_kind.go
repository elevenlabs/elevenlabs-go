package types

import (
	"fmt"
)

// What a service books.
//
// - `classic`: a staff-delivered service (haircut, consultation, massage, ...).
// - `rental`: an asset reservation (court, room, equipment).
// - `group`: seats in a multi-participant session (class, workshop, tour).
type ReceptionServiceKind string

const (
	ReceptionServiceKindClassic ReceptionServiceKind = "classic"
	ReceptionServiceKindRental  ReceptionServiceKind = "rental"
	ReceptionServiceKindGroup   ReceptionServiceKind = "group"
)

func NewReceptionServiceKindFromString(s string) (ReceptionServiceKind, error) {
	switch s {
	case "classic":
		return ReceptionServiceKindClassic, nil
	case "rental":
		return ReceptionServiceKindRental, nil
	case "group":
		return ReceptionServiceKindGroup, nil
	}
	var t ReceptionServiceKind
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ReceptionServiceKind) Ptr() *ReceptionServiceKind {
	return &e
}
