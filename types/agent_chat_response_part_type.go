package types

import (
	"fmt"
)

type AgentChatResponsePartType string

const (
	AgentChatResponsePartTypeStart AgentChatResponsePartType = "start"
	AgentChatResponsePartTypeDelta AgentChatResponsePartType = "delta"
	AgentChatResponsePartTypeStop  AgentChatResponsePartType = "stop"
)

func NewAgentChatResponsePartTypeFromString(s string) (AgentChatResponsePartType, error) {
	switch s {
	case "start":
		return AgentChatResponsePartTypeStart, nil
	case "delta":
		return AgentChatResponsePartTypeDelta, nil
	case "stop":
		return AgentChatResponsePartTypeStop, nil
	}
	var t AgentChatResponsePartType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AgentChatResponsePartType) Ptr() *AgentChatResponsePartType {
	return &e
}
