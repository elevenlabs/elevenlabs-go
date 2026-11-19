package types

import (
	"fmt"
)

type AuthorizationMethod string

const (
	AuthorizationMethodInvalid             AuthorizationMethod = "invalid"
	AuthorizationMethodPublic              AuthorizationMethod = "public"
	AuthorizationMethodAuthorizationHeader AuthorizationMethod = "authorization_header"
	AuthorizationMethodSignedUrl           AuthorizationMethod = "signed_url"
	AuthorizationMethodShareableLink       AuthorizationMethod = "shareable_link"
	AuthorizationMethodLivekitToken        AuthorizationMethod = "livekit_token"
	AuthorizationMethodLivekitTokenWebsite AuthorizationMethod = "livekit_token_website"
	AuthorizationMethodGenesysApiKey       AuthorizationMethod = "genesys_api_key"
	AuthorizationMethodAvayaApiKey         AuthorizationMethod = "avaya_api_key"
	AuthorizationMethodAudiocodesApiKey    AuthorizationMethod = "audiocodes_api_key"
	AuthorizationMethodWhatsapp            AuthorizationMethod = "whatsapp"
	AuthorizationMethodSms                 AuthorizationMethod = "sms"
)

func NewAuthorizationMethodFromString(s string) (AuthorizationMethod, error) {
	switch s {
	case "invalid":
		return AuthorizationMethodInvalid, nil
	case "public":
		return AuthorizationMethodPublic, nil
	case "authorization_header":
		return AuthorizationMethodAuthorizationHeader, nil
	case "signed_url":
		return AuthorizationMethodSignedUrl, nil
	case "shareable_link":
		return AuthorizationMethodShareableLink, nil
	case "livekit_token":
		return AuthorizationMethodLivekitToken, nil
	case "livekit_token_website":
		return AuthorizationMethodLivekitTokenWebsite, nil
	case "genesys_api_key":
		return AuthorizationMethodGenesysApiKey, nil
	case "avaya_api_key":
		return AuthorizationMethodAvayaApiKey, nil
	case "audiocodes_api_key":
		return AuthorizationMethodAudiocodesApiKey, nil
	case "whatsapp":
		return AuthorizationMethodWhatsapp, nil
	case "sms":
		return AuthorizationMethodSms, nil
	}
	var t AuthorizationMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AuthorizationMethod) Ptr() *AuthorizationMethod {
	return &e
}
