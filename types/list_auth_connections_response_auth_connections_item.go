package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// The type of auth connection config
type ListAuthConnectionsResponseAuthConnectionsItem struct {
	AuthType                      string
	ApiIntegrationOauth2AuthCode  *ApiIntegrationOAuth2AuthCodeResponse
	ApiIntegrationOauth2CustomApp *ApiIntegrationOAuth2CustomAppResponse
	BasicAuth                     *BasicAuthResponse
	BearerAuth                    *BearerAuthResponse
	CustomHeaderAuth              *CustomHeaderAuthResponse
	Mtls                          *MtlsAuthResponse
	Oauth2ClientCredentials       *OAuth2ClientCredsResponse
	Oauth2Jwt                     *OAuth2JwtResponse
	PrivateKeyJwt                 *PrivateKeyJwtResponse
	RefreshTokenAuth              *RefreshTokenAuthResponse
	SlackBotAuth                  *SlackBotAuthResponse
	UrlSecret                     *UrlSecretAuthResponse
	WhatsappAuth                  *WhatsAppAuthResponse
}

func (l *ListAuthConnectionsResponseAuthConnectionsItem) GetApiIntegrationOauth2AuthCode() *ApiIntegrationOAuth2AuthCodeResponse {
	if l == nil {
		return nil
	}
	if l.ApiIntegrationOauth2AuthCode == nil {
		return nil
	}
	return l.ApiIntegrationOauth2AuthCode
}

func (l *ListAuthConnectionsResponseAuthConnectionsItem) GetApiIntegrationOauth2CustomApp() *ApiIntegrationOAuth2CustomAppResponse {
	if l == nil {
		return nil
	}
	if l.ApiIntegrationOauth2CustomApp == nil {
		return nil
	}
	return l.ApiIntegrationOauth2CustomApp
}

func (l *ListAuthConnectionsResponseAuthConnectionsItem) GetBasicAuth() *BasicAuthResponse {
	if l == nil {
		return nil
	}
	if l.BasicAuth == nil {
		return nil
	}
	return l.BasicAuth
}

func (l *ListAuthConnectionsResponseAuthConnectionsItem) GetBearerAuth() *BearerAuthResponse {
	if l == nil {
		return nil
	}
	if l.BearerAuth == nil {
		return nil
	}
	return l.BearerAuth
}

func (l *ListAuthConnectionsResponseAuthConnectionsItem) GetCustomHeaderAuth() *CustomHeaderAuthResponse {
	if l == nil {
		return nil
	}
	if l.CustomHeaderAuth == nil {
		return nil
	}
	return l.CustomHeaderAuth
}

func (l *ListAuthConnectionsResponseAuthConnectionsItem) GetMtls() *MtlsAuthResponse {
	if l == nil {
		return nil
	}
	if l.Mtls == nil {
		return nil
	}
	return l.Mtls
}

func (l *ListAuthConnectionsResponseAuthConnectionsItem) GetOauth2ClientCredentials() *OAuth2ClientCredsResponse {
	if l == nil {
		return nil
	}
	if l.Oauth2ClientCredentials == nil {
		return nil
	}
	return l.Oauth2ClientCredentials
}

func (l *ListAuthConnectionsResponseAuthConnectionsItem) GetOauth2Jwt() *OAuth2JwtResponse {
	if l == nil {
		return nil
	}
	if l.Oauth2Jwt == nil {
		return nil
	}
	return l.Oauth2Jwt
}

func (l *ListAuthConnectionsResponseAuthConnectionsItem) GetPrivateKeyJwt() *PrivateKeyJwtResponse {
	if l == nil {
		return nil
	}
	if l.PrivateKeyJwt == nil {
		return nil
	}
	return l.PrivateKeyJwt
}

func (l *ListAuthConnectionsResponseAuthConnectionsItem) GetRefreshTokenAuth() *RefreshTokenAuthResponse {
	if l == nil {
		return nil
	}
	if l.RefreshTokenAuth == nil {
		return nil
	}
	return l.RefreshTokenAuth
}

func (l *ListAuthConnectionsResponseAuthConnectionsItem) GetSlackBotAuth() *SlackBotAuthResponse {
	if l == nil {
		return nil
	}
	if l.SlackBotAuth == nil {
		return nil
	}
	return l.SlackBotAuth
}

func (l *ListAuthConnectionsResponseAuthConnectionsItem) GetUrlSecret() *UrlSecretAuthResponse {
	if l == nil {
		return nil
	}
	if l.UrlSecret == nil {
		return nil
	}
	return l.UrlSecret
}

func (l *ListAuthConnectionsResponseAuthConnectionsItem) GetWhatsappAuth() *WhatsAppAuthResponse {
	if l == nil {
		return nil
	}
	if l.WhatsappAuth == nil {
		return nil
	}
	return l.WhatsappAuth
}

func (l *ListAuthConnectionsResponseAuthConnectionsItem) GetAuthType() string {
	if l == nil {
		return ""
	}
	return l.AuthType
}

func (l *ListAuthConnectionsResponseAuthConnectionsItem) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		AuthType string `json:"auth_type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	l.AuthType = unmarshaler.AuthType
	if unmarshaler.AuthType == "" {
		return fmt.Errorf("%T did not include discriminant auth_type", l)
	}
	switch unmarshaler.AuthType {
	case "api_integration_oauth2_auth_code":
		value := new(ApiIntegrationOAuth2AuthCodeResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		l.ApiIntegrationOauth2AuthCode = value
	case "api_integration_oauth2_custom_app":
		value := new(ApiIntegrationOAuth2CustomAppResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		l.ApiIntegrationOauth2CustomApp = value
	case "basic_auth":
		value := new(BasicAuthResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		l.BasicAuth = value
	case "bearer_auth":
		value := new(BearerAuthResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		l.BearerAuth = value
	case "custom_header_auth":
		value := new(CustomHeaderAuthResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		l.CustomHeaderAuth = value
	case "mtls":
		value := new(MtlsAuthResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		l.Mtls = value
	case "oauth2_client_credentials":
		value := new(OAuth2ClientCredsResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		l.Oauth2ClientCredentials = value
	case "oauth2_jwt":
		value := new(OAuth2JwtResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		l.Oauth2Jwt = value
	case "private_key_jwt":
		value := new(PrivateKeyJwtResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		l.PrivateKeyJwt = value
	case "refresh_token_auth":
		value := new(RefreshTokenAuthResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		l.RefreshTokenAuth = value
	case "slack_bot_auth":
		value := new(SlackBotAuthResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		l.SlackBotAuth = value
	case "url_secret":
		value := new(UrlSecretAuthResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		l.UrlSecret = value
	case "whatsapp_auth":
		value := new(WhatsAppAuthResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		l.WhatsappAuth = value
	}
	return nil
}

func (l ListAuthConnectionsResponseAuthConnectionsItem) MarshalJSON() ([]byte, error) {
	if err := l.validate(); err != nil {
		return nil, err
	}
	if l.ApiIntegrationOauth2AuthCode != nil {
		return internal.MarshalJSONWithDiscriminant(l.ApiIntegrationOauth2AuthCode, "auth_type", "api_integration_oauth2_auth_code")
	}
	if l.ApiIntegrationOauth2CustomApp != nil {
		return internal.MarshalJSONWithDiscriminant(l.ApiIntegrationOauth2CustomApp, "auth_type", "api_integration_oauth2_custom_app")
	}
	if l.BasicAuth != nil {
		return internal.MarshalJSONWithDiscriminant(l.BasicAuth, "auth_type", "basic_auth")
	}
	if l.BearerAuth != nil {
		return internal.MarshalJSONWithDiscriminant(l.BearerAuth, "auth_type", "bearer_auth")
	}
	if l.CustomHeaderAuth != nil {
		return internal.MarshalJSONWithDiscriminant(l.CustomHeaderAuth, "auth_type", "custom_header_auth")
	}
	if l.Mtls != nil {
		return internal.MarshalJSONWithDiscriminant(l.Mtls, "auth_type", "mtls")
	}
	if l.Oauth2ClientCredentials != nil {
		return internal.MarshalJSONWithDiscriminant(l.Oauth2ClientCredentials, "auth_type", "oauth2_client_credentials")
	}
	if l.Oauth2Jwt != nil {
		return internal.MarshalJSONWithDiscriminant(l.Oauth2Jwt, "auth_type", "oauth2_jwt")
	}
	if l.PrivateKeyJwt != nil {
		return internal.MarshalJSONWithDiscriminant(l.PrivateKeyJwt, "auth_type", "private_key_jwt")
	}
	if l.RefreshTokenAuth != nil {
		return internal.MarshalJSONWithDiscriminant(l.RefreshTokenAuth, "auth_type", "refresh_token_auth")
	}
	if l.SlackBotAuth != nil {
		return internal.MarshalJSONWithDiscriminant(l.SlackBotAuth, "auth_type", "slack_bot_auth")
	}
	if l.UrlSecret != nil {
		return internal.MarshalJSONWithDiscriminant(l.UrlSecret, "auth_type", "url_secret")
	}
	if l.WhatsappAuth != nil {
		return internal.MarshalJSONWithDiscriminant(l.WhatsappAuth, "auth_type", "whatsapp_auth")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", l)
}

type ListAuthConnectionsResponseAuthConnectionsItemVisitor interface {
	VisitApiIntegrationOauth2AuthCode(*ApiIntegrationOAuth2AuthCodeResponse) error
	VisitApiIntegrationOauth2CustomApp(*ApiIntegrationOAuth2CustomAppResponse) error
	VisitBasicAuth(*BasicAuthResponse) error
	VisitBearerAuth(*BearerAuthResponse) error
	VisitCustomHeaderAuth(*CustomHeaderAuthResponse) error
	VisitMtls(*MtlsAuthResponse) error
	VisitOauth2ClientCredentials(*OAuth2ClientCredsResponse) error
	VisitOauth2Jwt(*OAuth2JwtResponse) error
	VisitPrivateKeyJwt(*PrivateKeyJwtResponse) error
	VisitRefreshTokenAuth(*RefreshTokenAuthResponse) error
	VisitSlackBotAuth(*SlackBotAuthResponse) error
	VisitUrlSecret(*UrlSecretAuthResponse) error
	VisitWhatsappAuth(*WhatsAppAuthResponse) error
}

func (l *ListAuthConnectionsResponseAuthConnectionsItem) Accept(visitor ListAuthConnectionsResponseAuthConnectionsItemVisitor) error {
	if l.ApiIntegrationOauth2AuthCode != nil {
		return visitor.VisitApiIntegrationOauth2AuthCode(l.ApiIntegrationOauth2AuthCode)
	}
	if l.ApiIntegrationOauth2CustomApp != nil {
		return visitor.VisitApiIntegrationOauth2CustomApp(l.ApiIntegrationOauth2CustomApp)
	}
	if l.BasicAuth != nil {
		return visitor.VisitBasicAuth(l.BasicAuth)
	}
	if l.BearerAuth != nil {
		return visitor.VisitBearerAuth(l.BearerAuth)
	}
	if l.CustomHeaderAuth != nil {
		return visitor.VisitCustomHeaderAuth(l.CustomHeaderAuth)
	}
	if l.Mtls != nil {
		return visitor.VisitMtls(l.Mtls)
	}
	if l.Oauth2ClientCredentials != nil {
		return visitor.VisitOauth2ClientCredentials(l.Oauth2ClientCredentials)
	}
	if l.Oauth2Jwt != nil {
		return visitor.VisitOauth2Jwt(l.Oauth2Jwt)
	}
	if l.PrivateKeyJwt != nil {
		return visitor.VisitPrivateKeyJwt(l.PrivateKeyJwt)
	}
	if l.RefreshTokenAuth != nil {
		return visitor.VisitRefreshTokenAuth(l.RefreshTokenAuth)
	}
	if l.SlackBotAuth != nil {
		return visitor.VisitSlackBotAuth(l.SlackBotAuth)
	}
	if l.UrlSecret != nil {
		return visitor.VisitUrlSecret(l.UrlSecret)
	}
	if l.WhatsappAuth != nil {
		return visitor.VisitWhatsappAuth(l.WhatsappAuth)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", l)
}

func (l *ListAuthConnectionsResponseAuthConnectionsItem) validate() error {
	if l == nil {
		return fmt.Errorf("type %T is nil", l)
	}
	var fields []string
	if l.ApiIntegrationOauth2AuthCode != nil {
		fields = append(fields, "api_integration_oauth2_auth_code")
	}
	if l.ApiIntegrationOauth2CustomApp != nil {
		fields = append(fields, "api_integration_oauth2_custom_app")
	}
	if l.BasicAuth != nil {
		fields = append(fields, "basic_auth")
	}
	if l.BearerAuth != nil {
		fields = append(fields, "bearer_auth")
	}
	if l.CustomHeaderAuth != nil {
		fields = append(fields, "custom_header_auth")
	}
	if l.Mtls != nil {
		fields = append(fields, "mtls")
	}
	if l.Oauth2ClientCredentials != nil {
		fields = append(fields, "oauth2_client_credentials")
	}
	if l.Oauth2Jwt != nil {
		fields = append(fields, "oauth2_jwt")
	}
	if l.PrivateKeyJwt != nil {
		fields = append(fields, "private_key_jwt")
	}
	if l.RefreshTokenAuth != nil {
		fields = append(fields, "refresh_token_auth")
	}
	if l.SlackBotAuth != nil {
		fields = append(fields, "slack_bot_auth")
	}
	if l.UrlSecret != nil {
		fields = append(fields, "url_secret")
	}
	if l.WhatsappAuth != nil {
		fields = append(fields, "whatsapp_auth")
	}
	if len(fields) == 0 {
		if l.AuthType != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", l, l.AuthType)
		}
		return fmt.Errorf("type %T is empty", l)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", l, fields)
	}
	if l.AuthType != "" {
		field := fields[0]
		if l.AuthType != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				l,
				l.AuthType,
				l,
			)
		}
	}
	return nil
}
