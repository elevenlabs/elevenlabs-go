package types

import (
	"fmt"
)

type ChatSourceMedium string

const (
	ChatSourceMediumAudio ChatSourceMedium = "audio"
	ChatSourceMediumDtmf  ChatSourceMedium = "dtmf"
	ChatSourceMediumText  ChatSourceMedium = "text"
	ChatSourceMediumImage ChatSourceMedium = "image"
	ChatSourceMediumFile  ChatSourceMedium = "file"
)

func NewChatSourceMediumFromString(s string) (ChatSourceMedium, error) {
	switch s {
	case "audio":
		return ChatSourceMediumAudio, nil
	case "dtmf":
		return ChatSourceMediumDtmf, nil
	case "text":
		return ChatSourceMediumText, nil
	case "image":
		return ChatSourceMediumImage, nil
	case "file":
		return ChatSourceMediumFile, nil
	}
	var t ChatSourceMedium
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ChatSourceMedium) Ptr() *ChatSourceMedium {
	return &e
}
