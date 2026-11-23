package types

import (
	"fmt"
)

type UserMessageSourceMedium string

const (
	UserMessageSourceMediumAudio UserMessageSourceMedium = "audio"
	UserMessageSourceMediumDtmf  UserMessageSourceMedium = "dtmf"
	UserMessageSourceMediumText  UserMessageSourceMedium = "text"
	UserMessageSourceMediumImage UserMessageSourceMedium = "image"
	UserMessageSourceMediumFile  UserMessageSourceMedium = "file"
)

func NewUserMessageSourceMediumFromString(s string) (UserMessageSourceMedium, error) {
	switch s {
	case "audio":
		return UserMessageSourceMediumAudio, nil
	case "dtmf":
		return UserMessageSourceMediumDtmf, nil
	case "text":
		return UserMessageSourceMediumText, nil
	case "image":
		return UserMessageSourceMediumImage, nil
	case "file":
		return UserMessageSourceMediumFile, nil
	}
	var t UserMessageSourceMedium
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e UserMessageSourceMedium) Ptr() *UserMessageSourceMedium {
	return &e
}
