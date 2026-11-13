package types

import (
	"fmt"
)

type ExternalSyncJobTrigger string

const (
	ExternalSyncJobTriggerOnDemand  ExternalSyncJobTrigger = "on_demand"
	ExternalSyncJobTriggerOnConnect ExternalSyncJobTrigger = "on_connect"
	ExternalSyncJobTriggerAuto      ExternalSyncJobTrigger = "auto"
)

func NewExternalSyncJobTriggerFromString(s string) (ExternalSyncJobTrigger, error) {
	switch s {
	case "on_demand":
		return ExternalSyncJobTriggerOnDemand, nil
	case "on_connect":
		return ExternalSyncJobTriggerOnConnect, nil
	case "auto":
		return ExternalSyncJobTriggerAuto, nil
	}
	var t ExternalSyncJobTrigger
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ExternalSyncJobTrigger) Ptr() *ExternalSyncJobTrigger {
	return &e
}
