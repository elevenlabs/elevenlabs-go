package types

import (
	"fmt"
)

type ExternalSyncJobType string

const (
	ExternalSyncJobTypeFull        ExternalSyncJobType = "full"
	ExternalSyncJobTypeIncremental ExternalSyncJobType = "incremental"
)

func NewExternalSyncJobTypeFromString(s string) (ExternalSyncJobType, error) {
	switch s {
	case "full":
		return ExternalSyncJobTypeFull, nil
	case "incremental":
		return ExternalSyncJobTypeIncremental, nil
	}
	var t ExternalSyncJobType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ExternalSyncJobType) Ptr() *ExternalSyncJobType {
	return &e
}
