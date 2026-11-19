package types

import (
	"fmt"
)

type GetPronunciationDictionaryMetadataResponsePermissionOnResource string

const (
	GetPronunciationDictionaryMetadataResponsePermissionOnResourceAdmin     GetPronunciationDictionaryMetadataResponsePermissionOnResource = "admin"
	GetPronunciationDictionaryMetadataResponsePermissionOnResourceEditor    GetPronunciationDictionaryMetadataResponsePermissionOnResource = "editor"
	GetPronunciationDictionaryMetadataResponsePermissionOnResourceCommenter GetPronunciationDictionaryMetadataResponsePermissionOnResource = "commenter"
	GetPronunciationDictionaryMetadataResponsePermissionOnResourceViewer    GetPronunciationDictionaryMetadataResponsePermissionOnResource = "viewer"
)

func NewGetPronunciationDictionaryMetadataResponsePermissionOnResourceFromString(s string) (GetPronunciationDictionaryMetadataResponsePermissionOnResource, error) {
	switch s {
	case "admin":
		return GetPronunciationDictionaryMetadataResponsePermissionOnResourceAdmin, nil
	case "editor":
		return GetPronunciationDictionaryMetadataResponsePermissionOnResourceEditor, nil
	case "commenter":
		return GetPronunciationDictionaryMetadataResponsePermissionOnResourceCommenter, nil
	case "viewer":
		return GetPronunciationDictionaryMetadataResponsePermissionOnResourceViewer, nil
	}
	var t GetPronunciationDictionaryMetadataResponsePermissionOnResource
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e GetPronunciationDictionaryMetadataResponsePermissionOnResource) Ptr() *GetPronunciationDictionaryMetadataResponsePermissionOnResource {
	return &e
}
