package types

import (
	"fmt"
)

type ConversationTokenPurpose string

const (
	ConversationTokenPurposeSignedUrl     ConversationTokenPurpose = "signed_url"
	ConversationTokenPurposeShareableLink ConversationTokenPurpose = "shareable_link"
)

func NewConversationTokenPurposeFromString(s string) (ConversationTokenPurpose, error) {
	switch s {
	case "signed_url":
		return ConversationTokenPurposeSignedUrl, nil
	case "shareable_link":
		return ConversationTokenPurposeShareableLink, nil
	}
	var t ConversationTokenPurpose
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ConversationTokenPurpose) Ptr() *ConversationTokenPurpose {
	return &e
}
