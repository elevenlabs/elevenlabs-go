package types

import (
	"fmt"
)

// The mode in which to run this Dubbing job. Defaults to automatic, use manual if specifically providing a CSV transcript to use. Note that manual mode is experimental and production use is strongly discouraged.
type CreateDubbingRequestMode string

const (
	CreateDubbingRequestModeAutomatic CreateDubbingRequestMode = "automatic"
	CreateDubbingRequestModeManual    CreateDubbingRequestMode = "manual"
)

func NewCreateDubbingRequestModeFromString(s string) (CreateDubbingRequestMode, error) {
	switch s {
	case "automatic":
		return CreateDubbingRequestModeAutomatic, nil
	case "manual":
		return CreateDubbingRequestModeManual, nil
	}
	var t CreateDubbingRequestMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CreateDubbingRequestMode) Ptr() *CreateDubbingRequestMode {
	return &e
}
