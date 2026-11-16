package types

import (
	"fmt"
)

type AgentResponseType string

const (
	AgentResponseTypeText  AgentResponseType = "text"
	AgentResponseTypeVoice AgentResponseType = "voice"
)

func NewAgentResponseTypeFromString(s string) (AgentResponseType, error) {
	switch s {
	case "text":
		return AgentResponseTypeText, nil
	case "voice":
		return AgentResponseTypeVoice, nil
	}
	var t AgentResponseType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AgentResponseType) Ptr() *AgentResponseType {
	return &e
}
