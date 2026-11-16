package types

import (
	"fmt"
)

type WebhookAuthMethodType string

const (
	WebhookAuthMethodTypeHmac   WebhookAuthMethodType = "hmac"
	WebhookAuthMethodTypeOauth2 WebhookAuthMethodType = "oauth2"
	WebhookAuthMethodTypeMtls   WebhookAuthMethodType = "mtls"
)

func NewWebhookAuthMethodTypeFromString(s string) (WebhookAuthMethodType, error) {
	switch s {
	case "hmac":
		return WebhookAuthMethodTypeHmac, nil
	case "oauth2":
		return WebhookAuthMethodTypeOauth2, nil
	case "mtls":
		return WebhookAuthMethodTypeMtls, nil
	}
	var t WebhookAuthMethodType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e WebhookAuthMethodType) Ptr() *WebhookAuthMethodType {
	return &e
}
