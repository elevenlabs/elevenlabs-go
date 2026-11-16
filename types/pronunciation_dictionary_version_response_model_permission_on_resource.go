package types

import (
	"fmt"
)

type PronunciationDictionaryVersionResponseModelPermissionOnResource string

const (
	PronunciationDictionaryVersionResponseModelPermissionOnResourceAdmin     PronunciationDictionaryVersionResponseModelPermissionOnResource = "admin"
	PronunciationDictionaryVersionResponseModelPermissionOnResourceEditor    PronunciationDictionaryVersionResponseModelPermissionOnResource = "editor"
	PronunciationDictionaryVersionResponseModelPermissionOnResourceCommenter PronunciationDictionaryVersionResponseModelPermissionOnResource = "commenter"
	PronunciationDictionaryVersionResponseModelPermissionOnResourceViewer    PronunciationDictionaryVersionResponseModelPermissionOnResource = "viewer"
)

func NewPronunciationDictionaryVersionResponseModelPermissionOnResourceFromString(s string) (PronunciationDictionaryVersionResponseModelPermissionOnResource, error) {
	switch s {
	case "admin":
		return PronunciationDictionaryVersionResponseModelPermissionOnResourceAdmin, nil
	case "editor":
		return PronunciationDictionaryVersionResponseModelPermissionOnResourceEditor, nil
	case "commenter":
		return PronunciationDictionaryVersionResponseModelPermissionOnResourceCommenter, nil
	case "viewer":
		return PronunciationDictionaryVersionResponseModelPermissionOnResourceViewer, nil
	}
	var t PronunciationDictionaryVersionResponseModelPermissionOnResource
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e PronunciationDictionaryVersionResponseModelPermissionOnResource) Ptr() *PronunciationDictionaryVersionResponseModelPermissionOnResource {
	return &e
}
