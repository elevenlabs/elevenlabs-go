package types

import (
	"fmt"
)

type ReceptionAgentTarget string

const (
	ReceptionAgentTargetCustomerFacing ReceptionAgentTarget = "customer_facing"
	ReceptionAgentTargetAssistant      ReceptionAgentTarget = "assistant"
)

func NewReceptionAgentTargetFromString(s string) (ReceptionAgentTarget, error) {
	switch s {
	case "customer_facing":
		return ReceptionAgentTargetCustomerFacing, nil
	case "assistant":
		return ReceptionAgentTargetAssistant, nil
	}
	var t ReceptionAgentTarget
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ReceptionAgentTarget) Ptr() *ReceptionAgentTarget {
	return &e
}
