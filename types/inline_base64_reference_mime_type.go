package types

import (
	"fmt"
)

// The MIME type of the encoded media.
type InlineBase64ReferenceMimeType string

const (
	InlineBase64ReferenceMimeTypeImageJpeg      InlineBase64ReferenceMimeType = "image/jpeg"
	InlineBase64ReferenceMimeTypeImagePng       InlineBase64ReferenceMimeType = "image/png"
	InlineBase64ReferenceMimeTypeImageWebp      InlineBase64ReferenceMimeType = "image/webp"
	InlineBase64ReferenceMimeTypeImageHeic      InlineBase64ReferenceMimeType = "image/heic"
	InlineBase64ReferenceMimeTypeImageHeif      InlineBase64ReferenceMimeType = "image/heif"
	InlineBase64ReferenceMimeTypeAudioMpeg      InlineBase64ReferenceMimeType = "audio/mpeg"
	InlineBase64ReferenceMimeTypeAudioWav       InlineBase64ReferenceMimeType = "audio/wav"
	InlineBase64ReferenceMimeTypeVideoMp4       InlineBase64ReferenceMimeType = "video/mp4"
	InlineBase64ReferenceMimeTypeVideoQuicktime InlineBase64ReferenceMimeType = "video/quicktime"
	InlineBase64ReferenceMimeTypeVideoWebm      InlineBase64ReferenceMimeType = "video/webm"
)

func NewInlineBase64ReferenceMimeTypeFromString(s string) (InlineBase64ReferenceMimeType, error) {
	switch s {
	case "image/jpeg":
		return InlineBase64ReferenceMimeTypeImageJpeg, nil
	case "image/png":
		return InlineBase64ReferenceMimeTypeImagePng, nil
	case "image/webp":
		return InlineBase64ReferenceMimeTypeImageWebp, nil
	case "image/heic":
		return InlineBase64ReferenceMimeTypeImageHeic, nil
	case "image/heif":
		return InlineBase64ReferenceMimeTypeImageHeif, nil
	case "audio/mpeg":
		return InlineBase64ReferenceMimeTypeAudioMpeg, nil
	case "audio/wav":
		return InlineBase64ReferenceMimeTypeAudioWav, nil
	case "video/mp4":
		return InlineBase64ReferenceMimeTypeVideoMp4, nil
	case "video/quicktime":
		return InlineBase64ReferenceMimeTypeVideoQuicktime, nil
	case "video/webm":
		return InlineBase64ReferenceMimeTypeVideoWebm, nil
	}
	var t InlineBase64ReferenceMimeType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e InlineBase64ReferenceMimeType) Ptr() *InlineBase64ReferenceMimeType {
	return &e
}
