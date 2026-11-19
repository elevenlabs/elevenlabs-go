package types

import (
	"fmt"
)

type OrderItemKind string

const (
	OrderItemKindDub           OrderItemKind = "dub"
	OrderItemKindSubtitles     OrderItemKind = "subtitles"
	OrderItemKindTranscription OrderItemKind = "transcription"
)

func NewOrderItemKindFromString(s string) (OrderItemKind, error) {
	switch s {
	case "dub":
		return OrderItemKindDub, nil
	case "subtitles":
		return OrderItemKindSubtitles, nil
	case "transcription":
		return OrderItemKindTranscription, nil
	}
	var t OrderItemKind
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e OrderItemKind) Ptr() *OrderItemKind {
	return &e
}
