package types

import (
	"fmt"
)

type ClientToolResultItemErrorType string

const (
	ClientToolResultItemErrorTypeUserRejected   ClientToolResultItemErrorType = "user_rejected"
	ClientToolResultItemErrorTypeExternalServer ClientToolResultItemErrorType = "external_server"
	ClientToolResultItemErrorTypeExternalClient ClientToolResultItemErrorType = "external_client"
	ClientToolResultItemErrorTypeCustomerAuth   ClientToolResultItemErrorType = "customer_auth"
	ClientToolResultItemErrorTypeClientTimeout  ClientToolResultItemErrorType = "client_timeout"
	ClientToolResultItemErrorTypeUnknown        ClientToolResultItemErrorType = "unknown"
)

func NewClientToolResultItemErrorTypeFromString(s string) (ClientToolResultItemErrorType, error) {
	switch s {
	case "user_rejected":
		return ClientToolResultItemErrorTypeUserRejected, nil
	case "external_server":
		return ClientToolResultItemErrorTypeExternalServer, nil
	case "external_client":
		return ClientToolResultItemErrorTypeExternalClient, nil
	case "customer_auth":
		return ClientToolResultItemErrorTypeCustomerAuth, nil
	case "client_timeout":
		return ClientToolResultItemErrorTypeClientTimeout, nil
	case "unknown":
		return ClientToolResultItemErrorTypeUnknown, nil
	}
	var t ClientToolResultItemErrorType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ClientToolResultItemErrorType) Ptr() *ClientToolResultItemErrorType {
	return &e
}
