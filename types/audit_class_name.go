package types

import (
	"fmt"
)

// OCSF event classes emitted in audit logs.
type AuditClassName string

const (
	AuditClassNameAccountChange        AuditClassName = "Account Change"
	AuditClassNameAuthentication       AuditClassName = "Authentication"
	AuditClassNameEntityManagement     AuditClassName = "Entity Management"
	AuditClassNameUserAccessManagement AuditClassName = "User Access Management"
	AuditClassNameGroupManagement      AuditClassName = "Group Management"
)

func NewAuditClassNameFromString(s string) (AuditClassName, error) {
	switch s {
	case "Account Change":
		return AuditClassNameAccountChange, nil
	case "Authentication":
		return AuditClassNameAuthentication, nil
	case "Entity Management":
		return AuditClassNameEntityManagement, nil
	case "User Access Management":
		return AuditClassNameUserAccessManagement, nil
	case "Group Management":
		return AuditClassNameGroupManagement, nil
	}
	var t AuditClassName
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AuditClassName) Ptr() *AuditClassName {
	return &e
}
