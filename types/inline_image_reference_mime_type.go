package types

import (
	"fmt"
)

// The MIME type of the encoded image.
type InlineImageReferenceMimeType string

const (
	InlineImageReferenceMimeTypeImageJpeg InlineImageReferenceMimeType = "image/jpeg"
	InlineImageReferenceMimeTypeImagePng  InlineImageReferenceMimeType = "image/png"
	InlineImageReferenceMimeTypeImageWebp InlineImageReferenceMimeType = "image/webp"
	InlineImageReferenceMimeTypeImageHeic InlineImageReferenceMimeType = "image/heic"
	InlineImageReferenceMimeTypeImageHeif InlineImageReferenceMimeType = "image/heif"
)

func NewInlineImageReferenceMimeTypeFromString(s string) (InlineImageReferenceMimeType, error) {
	switch s {
	case "image/jpeg":
		return InlineImageReferenceMimeTypeImageJpeg, nil
	case "image/png":
		return InlineImageReferenceMimeTypeImagePng, nil
	case "image/webp":
		return InlineImageReferenceMimeTypeImageWebp, nil
	case "image/heic":
		return InlineImageReferenceMimeTypeImageHeic, nil
	case "image/heif":
		return InlineImageReferenceMimeTypeImageHeif, nil
	}
	var t InlineImageReferenceMimeType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e InlineImageReferenceMimeType) Ptr() *InlineImageReferenceMimeType {
	return &e
}
