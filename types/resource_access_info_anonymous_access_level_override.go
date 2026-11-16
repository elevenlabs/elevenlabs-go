package types

import (
	"fmt"
)

type ResourceAccessInfoAnonymousAccessLevelOverride string

const (
	ResourceAccessInfoAnonymousAccessLevelOverrideAdmin     ResourceAccessInfoAnonymousAccessLevelOverride = "admin"
	ResourceAccessInfoAnonymousAccessLevelOverrideEditor    ResourceAccessInfoAnonymousAccessLevelOverride = "editor"
	ResourceAccessInfoAnonymousAccessLevelOverrideCommenter ResourceAccessInfoAnonymousAccessLevelOverride = "commenter"
	ResourceAccessInfoAnonymousAccessLevelOverrideViewer    ResourceAccessInfoAnonymousAccessLevelOverride = "viewer"
)

func NewResourceAccessInfoAnonymousAccessLevelOverrideFromString(s string) (ResourceAccessInfoAnonymousAccessLevelOverride, error) {
	switch s {
	case "admin":
		return ResourceAccessInfoAnonymousAccessLevelOverrideAdmin, nil
	case "editor":
		return ResourceAccessInfoAnonymousAccessLevelOverrideEditor, nil
	case "commenter":
		return ResourceAccessInfoAnonymousAccessLevelOverrideCommenter, nil
	case "viewer":
		return ResourceAccessInfoAnonymousAccessLevelOverrideViewer, nil
	}
	var t ResourceAccessInfoAnonymousAccessLevelOverride
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ResourceAccessInfoAnonymousAccessLevelOverride) Ptr() *ResourceAccessInfoAnonymousAccessLevelOverride {
	return &e
}
