package types

import (
	"fmt"
)

// Single status field shared by every auth type's stored credential.
//
// OAuth values (“REFRESH_FAILED“, “REVOKED“) are written by the OAuth
// token-manager refresh path. “CREDENTIAL_INVALID“ is written by the
// tool execution path when an upstream response matches a credential's
// “failure_signatures“ entry (Bearer, Basic auth, etc.).
type AuthConnectionStatus string

const (
	AuthConnectionStatusActive            AuthConnectionStatus = "active"
	AuthConnectionStatusRefreshFailed     AuthConnectionStatus = "refresh_failed"
	AuthConnectionStatusRevoked           AuthConnectionStatus = "revoked"
	AuthConnectionStatusCredentialInvalid AuthConnectionStatus = "credential_invalid"
)

func NewAuthConnectionStatusFromString(s string) (AuthConnectionStatus, error) {
	switch s {
	case "active":
		return AuthConnectionStatusActive, nil
	case "refresh_failed":
		return AuthConnectionStatusRefreshFailed, nil
	case "revoked":
		return AuthConnectionStatusRevoked, nil
	case "credential_invalid":
		return AuthConnectionStatusCredentialInvalid, nil
	}
	var t AuthConnectionStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AuthConnectionStatus) Ptr() *AuthConnectionStatus {
	return &e
}
