package types

import (
	"fmt"
)

// Optional category for the failure reason. When set, `is_error` is treated as true. `user_rejected` is special-cased to mark the tool as not having been called.
type ClientToolResultErrorType string

const (
	ClientToolResultErrorTypeUserRejected   ClientToolResultErrorType = "user_rejected"
	ClientToolResultErrorTypeExternalServer ClientToolResultErrorType = "external_server"
	ClientToolResultErrorTypeExternalClient ClientToolResultErrorType = "external_client"
	ClientToolResultErrorTypeCustomerAuth   ClientToolResultErrorType = "customer_auth"
	ClientToolResultErrorTypeClientTimeout  ClientToolResultErrorType = "client_timeout"
	ClientToolResultErrorTypeUnknown        ClientToolResultErrorType = "unknown"
)

func NewClientToolResultErrorTypeFromString(s string) (ClientToolResultErrorType, error) {
	switch s {
	case "user_rejected":
		return ClientToolResultErrorTypeUserRejected, nil
	case "external_server":
		return ClientToolResultErrorTypeExternalServer, nil
	case "external_client":
		return ClientToolResultErrorTypeExternalClient, nil
	case "customer_auth":
		return ClientToolResultErrorTypeCustomerAuth, nil
	case "client_timeout":
		return ClientToolResultErrorTypeClientTimeout, nil
	case "unknown":
		return ClientToolResultErrorTypeUnknown, nil
	}
	var t ClientToolResultErrorType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ClientToolResultErrorType) Ptr() *ClientToolResultErrorType {
	return &e
}
