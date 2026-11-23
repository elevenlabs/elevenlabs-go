package types

import (
	"fmt"
)

type AgentChatResponsePartTextResponsePartType string

const (
	AgentChatResponsePartTextResponsePartTypeStart AgentChatResponsePartTextResponsePartType = "start"
	AgentChatResponsePartTextResponsePartTypeDelta AgentChatResponsePartTextResponsePartType = "delta"
	AgentChatResponsePartTextResponsePartTypeStop  AgentChatResponsePartTextResponsePartType = "stop"
)

func NewAgentChatResponsePartTextResponsePartTypeFromString(s string) (AgentChatResponsePartTextResponsePartType, error) {
	switch s {
	case "start":
		return AgentChatResponsePartTextResponsePartTypeStart, nil
	case "delta":
		return AgentChatResponsePartTextResponsePartTypeDelta, nil
	case "stop":
		return AgentChatResponsePartTextResponsePartTypeStop, nil
	}
	var t AgentChatResponsePartTextResponsePartType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AgentChatResponsePartTextResponsePartType) Ptr() *AgentChatResponsePartTextResponsePartType {
	return &e
}
