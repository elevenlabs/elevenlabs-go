package types

import (
	"fmt"
)

type WebhookHttpMethod string

const (
	WebhookHttpMethodGet    WebhookHttpMethod = "GET"
	WebhookHttpMethodPost   WebhookHttpMethod = "POST"
	WebhookHttpMethodPut    WebhookHttpMethod = "PUT"
	WebhookHttpMethodPatch  WebhookHttpMethod = "PATCH"
	WebhookHttpMethodDelete WebhookHttpMethod = "DELETE"
)

func NewWebhookHttpMethodFromString(s string) (WebhookHttpMethod, error) {
	switch s {
	case "GET":
		return WebhookHttpMethodGet, nil
	case "POST":
		return WebhookHttpMethodPost, nil
	case "PUT":
		return WebhookHttpMethodPut, nil
	case "PATCH":
		return WebhookHttpMethodPatch, nil
	case "DELETE":
		return WebhookHttpMethodDelete, nil
	}
	var t WebhookHttpMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e WebhookHttpMethod) Ptr() *WebhookHttpMethod {
	return &e
}
