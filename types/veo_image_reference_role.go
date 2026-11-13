package types

import (
	"fmt"
)

// How the model uses the image: `subject` places its subject or scene elements into the video; `style` transfers its visual style.
type VeoImageReferenceRole string

const (
	VeoImageReferenceRoleSubject VeoImageReferenceRole = "subject"
	VeoImageReferenceRoleStyle   VeoImageReferenceRole = "style"
)

func NewVeoImageReferenceRoleFromString(s string) (VeoImageReferenceRole, error) {
	switch s {
	case "subject":
		return VeoImageReferenceRoleSubject, nil
	case "style":
		return VeoImageReferenceRoleStyle, nil
	}
	var t VeoImageReferenceRole
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e VeoImageReferenceRole) Ptr() *VeoImageReferenceRole {
	return &e
}
