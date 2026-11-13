package types

import (
	"fmt"
)

// Which product surface owns this agent document.
type ConversationProduct string

const (
	ConversationProductAgents       ConversationProduct = "agents"
	ConversationProductSpeechEngine ConversationProduct = "speech_engine"
)

func NewConversationProductFromString(s string) (ConversationProduct, error) {
	switch s {
	case "agents":
		return ConversationProductAgents, nil
	case "speech_engine":
		return ConversationProductSpeechEngine, nil
	}
	var t ConversationProduct
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ConversationProduct) Ptr() *ConversationProduct {
	return &e
}
