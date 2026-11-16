package types

import (
	"fmt"
)

type ResourceAccessInfoAccessSource string

const (
	ResourceAccessInfoAccessSourceCreator          ResourceAccessInfoAccessSource = "creator"
	ResourceAccessInfoAccessSourceExplicit         ResourceAccessInfoAccessSource = "explicit"
	ResourceAccessInfoAccessSourceWorkspaceAdmin   ResourceAccessInfoAccessSource = "workspace_admin"
	ResourceAccessInfoAccessSourceWorkspaceDefault ResourceAccessInfoAccessSource = "workspace_default"
)

func NewResourceAccessInfoAccessSourceFromString(s string) (ResourceAccessInfoAccessSource, error) {
	switch s {
	case "creator":
		return ResourceAccessInfoAccessSourceCreator, nil
	case "explicit":
		return ResourceAccessInfoAccessSourceExplicit, nil
	case "workspace_admin":
		return ResourceAccessInfoAccessSourceWorkspaceAdmin, nil
	case "workspace_default":
		return ResourceAccessInfoAccessSourceWorkspaceDefault, nil
	}
	var t ResourceAccessInfoAccessSource
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ResourceAccessInfoAccessSource) Ptr() *ResourceAccessInfoAccessSource {
	return &e
}
