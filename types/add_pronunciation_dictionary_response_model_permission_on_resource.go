package types

import (
	"fmt"
)

type AddPronunciationDictionaryResponseModelPermissionOnResource string

const (
	AddPronunciationDictionaryResponseModelPermissionOnResourceAdmin     AddPronunciationDictionaryResponseModelPermissionOnResource = "admin"
	AddPronunciationDictionaryResponseModelPermissionOnResourceEditor    AddPronunciationDictionaryResponseModelPermissionOnResource = "editor"
	AddPronunciationDictionaryResponseModelPermissionOnResourceCommenter AddPronunciationDictionaryResponseModelPermissionOnResource = "commenter"
	AddPronunciationDictionaryResponseModelPermissionOnResourceViewer    AddPronunciationDictionaryResponseModelPermissionOnResource = "viewer"
)

func NewAddPronunciationDictionaryResponseModelPermissionOnResourceFromString(s string) (AddPronunciationDictionaryResponseModelPermissionOnResource, error) {
	switch s {
	case "admin":
		return AddPronunciationDictionaryResponseModelPermissionOnResourceAdmin, nil
	case "editor":
		return AddPronunciationDictionaryResponseModelPermissionOnResourceEditor, nil
	case "commenter":
		return AddPronunciationDictionaryResponseModelPermissionOnResourceCommenter, nil
	case "viewer":
		return AddPronunciationDictionaryResponseModelPermissionOnResourceViewer, nil
	}
	var t AddPronunciationDictionaryResponseModelPermissionOnResource
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AddPronunciationDictionaryResponseModelPermissionOnResource) Ptr() *AddPronunciationDictionaryResponseModelPermissionOnResource {
	return &e
}
