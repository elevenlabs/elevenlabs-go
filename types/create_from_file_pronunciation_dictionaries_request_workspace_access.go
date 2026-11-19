package types

import (
	"fmt"
)

type CreateFromFilePronunciationDictionariesRequestWorkspaceAccess string

const (
	CreateFromFilePronunciationDictionariesRequestWorkspaceAccessAdmin     CreateFromFilePronunciationDictionariesRequestWorkspaceAccess = "admin"
	CreateFromFilePronunciationDictionariesRequestWorkspaceAccessEditor    CreateFromFilePronunciationDictionariesRequestWorkspaceAccess = "editor"
	CreateFromFilePronunciationDictionariesRequestWorkspaceAccessCommenter CreateFromFilePronunciationDictionariesRequestWorkspaceAccess = "commenter"
	CreateFromFilePronunciationDictionariesRequestWorkspaceAccessViewer    CreateFromFilePronunciationDictionariesRequestWorkspaceAccess = "viewer"
)

func NewCreateFromFilePronunciationDictionariesRequestWorkspaceAccessFromString(s string) (CreateFromFilePronunciationDictionariesRequestWorkspaceAccess, error) {
	switch s {
	case "admin":
		return CreateFromFilePronunciationDictionariesRequestWorkspaceAccessAdmin, nil
	case "editor":
		return CreateFromFilePronunciationDictionariesRequestWorkspaceAccessEditor, nil
	case "commenter":
		return CreateFromFilePronunciationDictionariesRequestWorkspaceAccessCommenter, nil
	case "viewer":
		return CreateFromFilePronunciationDictionariesRequestWorkspaceAccessViewer, nil
	}
	var t CreateFromFilePronunciationDictionariesRequestWorkspaceAccess
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CreateFromFilePronunciationDictionariesRequestWorkspaceAccess) Ptr() *CreateFromFilePronunciationDictionariesRequestWorkspaceAccess {
	return &e
}
