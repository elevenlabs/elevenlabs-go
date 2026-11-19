package types

import (
	"fmt"
)

type SmbAgentType string

const (
	SmbAgentTypeCustomerFacing SmbAgentType = "customer_facing"
	SmbAgentTypeAssistant      SmbAgentType = "assistant"
)

func NewSmbAgentTypeFromString(s string) (SmbAgentType, error) {
	switch s {
	case "customer_facing":
		return SmbAgentTypeCustomerFacing, nil
	case "assistant":
		return SmbAgentTypeAssistant, nil
	}
	var t SmbAgentType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SmbAgentType) Ptr() *SmbAgentType {
	return &e
}
