package types

import (
	"fmt"
)

// The role of the user making the request
type ResourceAccessInfoRole string

const (
	ResourceAccessInfoRoleAdmin     ResourceAccessInfoRole = "admin"
	ResourceAccessInfoRoleEditor    ResourceAccessInfoRole = "editor"
	ResourceAccessInfoRoleCommenter ResourceAccessInfoRole = "commenter"
	ResourceAccessInfoRoleViewer    ResourceAccessInfoRole = "viewer"
)

func NewResourceAccessInfoRoleFromString(s string) (ResourceAccessInfoRole, error) {
	switch s {
	case "admin":
		return ResourceAccessInfoRoleAdmin, nil
	case "editor":
		return ResourceAccessInfoRoleEditor, nil
	case "commenter":
		return ResourceAccessInfoRoleCommenter, nil
	case "viewer":
		return ResourceAccessInfoRoleViewer, nil
	}
	var t ResourceAccessInfoRole
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ResourceAccessInfoRole) Ptr() *ResourceAccessInfoRole {
	return &e
}
