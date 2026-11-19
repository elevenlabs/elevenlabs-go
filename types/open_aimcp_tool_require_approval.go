package types

import (
	"fmt"
)

type OpenAimcpToolRequireApproval string

const (
	OpenAimcpToolRequireApprovalNever  OpenAimcpToolRequireApproval = "never"
	OpenAimcpToolRequireApprovalAlways OpenAimcpToolRequireApproval = "always"
)

func NewOpenAimcpToolRequireApprovalFromString(s string) (OpenAimcpToolRequireApproval, error) {
	switch s {
	case "never":
		return OpenAimcpToolRequireApprovalNever, nil
	case "always":
		return OpenAimcpToolRequireApprovalAlways, nil
	}
	var t OpenAimcpToolRequireApproval
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e OpenAimcpToolRequireApproval) Ptr() *OpenAimcpToolRequireApproval {
	return &e
}
