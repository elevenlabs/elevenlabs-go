package types

import (
	"fmt"
)

type RenderStatus string

const (
	RenderStatusComplete   RenderStatus = "complete"
	RenderStatusProcessing RenderStatus = "processing"
	RenderStatusFailed     RenderStatus = "failed"
)

func NewRenderStatusFromString(s string) (RenderStatus, error) {
	switch s {
	case "complete":
		return RenderStatusComplete, nil
	case "processing":
		return RenderStatusProcessing, nil
	case "failed":
		return RenderStatusFailed, nil
	}
	var t RenderStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e RenderStatus) Ptr() *RenderStatus {
	return &e
}
