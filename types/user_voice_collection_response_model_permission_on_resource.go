package types

import (
	"fmt"
)

// The caller's access role on the voice collection.
type UserVoiceCollectionResponseModelPermissionOnResource string

const (
	UserVoiceCollectionResponseModelPermissionOnResourceAdmin     UserVoiceCollectionResponseModelPermissionOnResource = "admin"
	UserVoiceCollectionResponseModelPermissionOnResourceEditor    UserVoiceCollectionResponseModelPermissionOnResource = "editor"
	UserVoiceCollectionResponseModelPermissionOnResourceCommenter UserVoiceCollectionResponseModelPermissionOnResource = "commenter"
	UserVoiceCollectionResponseModelPermissionOnResourceViewer    UserVoiceCollectionResponseModelPermissionOnResource = "viewer"
)

func NewUserVoiceCollectionResponseModelPermissionOnResourceFromString(s string) (UserVoiceCollectionResponseModelPermissionOnResource, error) {
	switch s {
	case "admin":
		return UserVoiceCollectionResponseModelPermissionOnResourceAdmin, nil
	case "editor":
		return UserVoiceCollectionResponseModelPermissionOnResourceEditor, nil
	case "commenter":
		return UserVoiceCollectionResponseModelPermissionOnResourceCommenter, nil
	case "viewer":
		return UserVoiceCollectionResponseModelPermissionOnResourceViewer, nil
	}
	var t UserVoiceCollectionResponseModelPermissionOnResource
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e UserVoiceCollectionResponseModelPermissionOnResource) Ptr() *UserVoiceCollectionResponseModelPermissionOnResource {
	return &e
}
