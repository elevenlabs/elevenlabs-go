package types

import (
	"fmt"
)

type FileInputMimeType string

const (
	FileInputMimeTypeImagePng       FileInputMimeType = "image/png"
	FileInputMimeTypeImageJpeg      FileInputMimeType = "image/jpeg"
	FileInputMimeTypeImageGif       FileInputMimeType = "image/gif"
	FileInputMimeTypeImageWebp      FileInputMimeType = "image/webp"
	FileInputMimeTypeApplicationPdf FileInputMimeType = "application/pdf"
	FileInputMimeTypeTextPlain      FileInputMimeType = "text/plain"
)

func NewFileInputMimeTypeFromString(s string) (FileInputMimeType, error) {
	switch s {
	case "image/png":
		return FileInputMimeTypeImagePng, nil
	case "image/jpeg":
		return FileInputMimeTypeImageJpeg, nil
	case "image/gif":
		return FileInputMimeTypeImageGif, nil
	case "image/webp":
		return FileInputMimeTypeImageWebp, nil
	case "application/pdf":
		return FileInputMimeTypeApplicationPdf, nil
	case "text/plain":
		return FileInputMimeTypeTextPlain, nil
	}
	var t FileInputMimeType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e FileInputMimeType) Ptr() *FileInputMimeType {
	return &e
}
