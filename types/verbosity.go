package types

import (
	"fmt"
)

type Verbosity string

const (
	VerbosityAuto     Verbosity = "auto"
	VerbosityConcise  Verbosity = "concise"
	VerbosityThorough Verbosity = "thorough"
)

func NewVerbosityFromString(s string) (Verbosity, error) {
	switch s {
	case "auto":
		return VerbosityAuto, nil
	case "concise":
		return VerbosityConcise, nil
	case "thorough":
		return VerbosityThorough, nil
	}
	var t Verbosity
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e Verbosity) Ptr() *Verbosity {
	return &e
}
