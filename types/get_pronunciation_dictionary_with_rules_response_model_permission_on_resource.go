package types

import (
	"fmt"
)

type GetPronunciationDictionaryWithRulesResponseModelPermissionOnResource string

const (
	GetPronunciationDictionaryWithRulesResponseModelPermissionOnResourceAdmin     GetPronunciationDictionaryWithRulesResponseModelPermissionOnResource = "admin"
	GetPronunciationDictionaryWithRulesResponseModelPermissionOnResourceEditor    GetPronunciationDictionaryWithRulesResponseModelPermissionOnResource = "editor"
	GetPronunciationDictionaryWithRulesResponseModelPermissionOnResourceCommenter GetPronunciationDictionaryWithRulesResponseModelPermissionOnResource = "commenter"
	GetPronunciationDictionaryWithRulesResponseModelPermissionOnResourceViewer    GetPronunciationDictionaryWithRulesResponseModelPermissionOnResource = "viewer"
)

func NewGetPronunciationDictionaryWithRulesResponseModelPermissionOnResourceFromString(s string) (GetPronunciationDictionaryWithRulesResponseModelPermissionOnResource, error) {
	switch s {
	case "admin":
		return GetPronunciationDictionaryWithRulesResponseModelPermissionOnResourceAdmin, nil
	case "editor":
		return GetPronunciationDictionaryWithRulesResponseModelPermissionOnResourceEditor, nil
	case "commenter":
		return GetPronunciationDictionaryWithRulesResponseModelPermissionOnResourceCommenter, nil
	case "viewer":
		return GetPronunciationDictionaryWithRulesResponseModelPermissionOnResourceViewer, nil
	}
	var t GetPronunciationDictionaryWithRulesResponseModelPermissionOnResource
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e GetPronunciationDictionaryWithRulesResponseModelPermissionOnResource) Ptr() *GetPronunciationDictionaryWithRulesResponseModelPermissionOnResource {
	return &e
}
