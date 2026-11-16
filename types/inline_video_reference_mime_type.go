package types

import (
	"fmt"
)

// The MIME type of the encoded video.
type InlineVideoReferenceMimeType string

const (
	InlineVideoReferenceMimeTypeVideoMp4       InlineVideoReferenceMimeType = "video/mp4"
	InlineVideoReferenceMimeTypeVideoQuicktime InlineVideoReferenceMimeType = "video/quicktime"
	InlineVideoReferenceMimeTypeVideoWebm      InlineVideoReferenceMimeType = "video/webm"
)

func NewInlineVideoReferenceMimeTypeFromString(s string) (InlineVideoReferenceMimeType, error) {
	switch s {
	case "video/mp4":
		return InlineVideoReferenceMimeTypeVideoMp4, nil
	case "video/quicktime":
		return InlineVideoReferenceMimeTypeVideoQuicktime, nil
	case "video/webm":
		return InlineVideoReferenceMimeTypeVideoWebm, nil
	}
	var t InlineVideoReferenceMimeType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e InlineVideoReferenceMimeType) Ptr() *InlineVideoReferenceMimeType {
	return &e
}
