package types

import (
	"fmt"
)

type ColumnUnit string

const (
	ColumnUnitMs       ColumnUnit = "ms"
	ColumnUnitS        ColumnUnit = "s"
	ColumnUnitMin      ColumnUnit = "min"
	ColumnUnitDuration ColumnUnit = "duration"
	ColumnUnitCredits  ColumnUnit = "credits"
	ColumnUnitUsd      ColumnUnit = "usd"
	ColumnUnitEur      ColumnUnit = "eur"
	ColumnUnitInr      ColumnUnit = "inr"
	ColumnUnitPln      ColumnUnit = "pln"
	ColumnUnitGbp      ColumnUnit = "gbp"
	ColumnUnitRatio    ColumnUnit = "ratio"
	ColumnUnitRating   ColumnUnit = "rating"
)

func NewColumnUnitFromString(s string) (ColumnUnit, error) {
	switch s {
	case "ms":
		return ColumnUnitMs, nil
	case "s":
		return ColumnUnitS, nil
	case "min":
		return ColumnUnitMin, nil
	case "duration":
		return ColumnUnitDuration, nil
	case "credits":
		return ColumnUnitCredits, nil
	case "usd":
		return ColumnUnitUsd, nil
	case "eur":
		return ColumnUnitEur, nil
	case "inr":
		return ColumnUnitInr, nil
	case "pln":
		return ColumnUnitPln, nil
	case "gbp":
		return ColumnUnitGbp, nil
	case "ratio":
		return ColumnUnitRatio, nil
	case "rating":
		return ColumnUnitRating, nil
	}
	var t ColumnUnit
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ColumnUnit) Ptr() *ColumnUnit {
	return &e
}
