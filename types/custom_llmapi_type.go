package types

import (
	"fmt"
)

type CustomLlmapiType string

const (
	CustomLlmapiTypeChatCompletions CustomLlmapiType = "chat_completions"
	CustomLlmapiTypeResponses       CustomLlmapiType = "responses"
	CustomLlmapiTypeWebsocket       CustomLlmapiType = "websocket"
)

func NewCustomLlmapiTypeFromString(s string) (CustomLlmapiType, error) {
	switch s {
	case "chat_completions":
		return CustomLlmapiTypeChatCompletions, nil
	case "responses":
		return CustomLlmapiTypeResponses, nil
	case "websocket":
		return CustomLlmapiTypeWebsocket, nil
	}
	var t CustomLlmapiType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CustomLlmapiType) Ptr() *CustomLlmapiType {
	return &e
}
