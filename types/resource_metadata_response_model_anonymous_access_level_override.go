package types

import (
	"fmt"
)

type ResourceMetadataResponseModelAnonymousAccessLevelOverride string

const (
	ResourceMetadataResponseModelAnonymousAccessLevelOverrideAdmin     ResourceMetadataResponseModelAnonymousAccessLevelOverride = "admin"
	ResourceMetadataResponseModelAnonymousAccessLevelOverrideEditor    ResourceMetadataResponseModelAnonymousAccessLevelOverride = "editor"
	ResourceMetadataResponseModelAnonymousAccessLevelOverrideCommenter ResourceMetadataResponseModelAnonymousAccessLevelOverride = "commenter"
	ResourceMetadataResponseModelAnonymousAccessLevelOverrideViewer    ResourceMetadataResponseModelAnonymousAccessLevelOverride = "viewer"
)

func NewResourceMetadataResponseModelAnonymousAccessLevelOverrideFromString(s string) (ResourceMetadataResponseModelAnonymousAccessLevelOverride, error) {
	switch s {
	case "admin":
		return ResourceMetadataResponseModelAnonymousAccessLevelOverrideAdmin, nil
	case "editor":
		return ResourceMetadataResponseModelAnonymousAccessLevelOverrideEditor, nil
	case "commenter":
		return ResourceMetadataResponseModelAnonymousAccessLevelOverrideCommenter, nil
	case "viewer":
		return ResourceMetadataResponseModelAnonymousAccessLevelOverrideViewer, nil
	}
	var t ResourceMetadataResponseModelAnonymousAccessLevelOverride
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ResourceMetadataResponseModelAnonymousAccessLevelOverride) Ptr() *ResourceMetadataResponseModelAnonymousAccessLevelOverride {
	return &e
}
