package types

import (
	"fmt"
)

type BranchProtectionStatus string

const (
	BranchProtectionStatusWriterPermsRequired BranchProtectionStatus = "writer_perms_required"
	BranchProtectionStatusAdminPermsRequired  BranchProtectionStatus = "admin_perms_required"
)

func NewBranchProtectionStatusFromString(s string) (BranchProtectionStatus, error) {
	switch s {
	case "writer_perms_required":
		return BranchProtectionStatusWriterPermsRequired, nil
	case "admin_perms_required":
		return BranchProtectionStatusAdminPermsRequired, nil
	}
	var t BranchProtectionStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BranchProtectionStatus) Ptr() *BranchProtectionStatus {
	return &e
}
