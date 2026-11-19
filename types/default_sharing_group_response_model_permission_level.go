package types

import (
	"fmt"
)

// The permission level to grant to the group
type DefaultSharingGroupResponseModelPermissionLevel string

const (
	DefaultSharingGroupResponseModelPermissionLevelAdmin  DefaultSharingGroupResponseModelPermissionLevel = "admin"
	DefaultSharingGroupResponseModelPermissionLevelEditor DefaultSharingGroupResponseModelPermissionLevel = "editor"
	DefaultSharingGroupResponseModelPermissionLevelViewer DefaultSharingGroupResponseModelPermissionLevel = "viewer"
)

func NewDefaultSharingGroupResponseModelPermissionLevelFromString(s string) (DefaultSharingGroupResponseModelPermissionLevel, error) {
	switch s {
	case "admin":
		return DefaultSharingGroupResponseModelPermissionLevelAdmin, nil
	case "editor":
		return DefaultSharingGroupResponseModelPermissionLevelEditor, nil
	case "viewer":
		return DefaultSharingGroupResponseModelPermissionLevelViewer, nil
	}
	var t DefaultSharingGroupResponseModelPermissionLevel
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e DefaultSharingGroupResponseModelPermissionLevel) Ptr() *DefaultSharingGroupResponseModelPermissionLevel {
	return &e
}
