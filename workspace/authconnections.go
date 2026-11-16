package workspace

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/types"
)

// The type of auth connection config
type CreateAuthConnectionsResponse struct {
	AuthType                      string
	ApiIntegrationOauth2AuthCode  *types.ApiIntegrationOAuth2AuthCodeResponse
	ApiIntegrationOauth2CustomApp *types.ApiIntegrationOAuth2CustomAppResponse
	BasicAuth                     *types.BasicAuthResponse
	BearerAuth                    *types.BearerAuthResponse
	CustomHeaderAuth              *types.CustomHeaderAuthResponse
	Mtls                          *types.MtlsAuthResponse
	Oauth2ClientCredentials       *types.OAuth2ClientCredsResponse
	Oauth2Jwt                     *types.OAuth2JwtResponse
	PrivateKeyJwt                 *types.PrivateKeyJwtResponse
	RefreshTokenAuth              *types.RefreshTokenAuthResponse
	SlackBotAuth                  *types.SlackBotAuthResponse
	UrlSecret                     *types.UrlSecretAuthResponse
	WhatsappAuth                  *types.WhatsAppAuthResponse
}

func (c *CreateAuthConnectionsResponse) GetApiIntegrationOauth2AuthCode() *types.ApiIntegrationOAuth2AuthCodeResponse {
	if c == nil {
		return nil
	}
	if c.ApiIntegrationOauth2AuthCode == nil {
		return nil
	}
	return c.ApiIntegrationOauth2AuthCode
}

func (c *CreateAuthConnectionsResponse) GetApiIntegrationOauth2CustomApp() *types.ApiIntegrationOAuth2CustomAppResponse {
	if c == nil {
		return nil
	}
	if c.ApiIntegrationOauth2CustomApp == nil {
		return nil
	}
	return c.ApiIntegrationOauth2CustomApp
}

func (c *CreateAuthConnectionsResponse) GetBasicAuth() *types.BasicAuthResponse {
	if c == nil {
		return nil
	}
	if c.BasicAuth == nil {
		return nil
	}
	return c.BasicAuth
}

func (c *CreateAuthConnectionsResponse) GetBearerAuth() *types.BearerAuthResponse {
	if c == nil {
		return nil
	}
	if c.BearerAuth == nil {
		return nil
	}
	return c.BearerAuth
}

func (c *CreateAuthConnectionsResponse) GetCustomHeaderAuth() *types.CustomHeaderAuthResponse {
	if c == nil {
		return nil
	}
	if c.CustomHeaderAuth == nil {
		return nil
	}
	return c.CustomHeaderAuth
}

func (c *CreateAuthConnectionsResponse) GetMtls() *types.MtlsAuthResponse {
	if c == nil {
		return nil
	}
	if c.Mtls == nil {
		return nil
	}
	return c.Mtls
}

func (c *CreateAuthConnectionsResponse) GetOauth2ClientCredentials() *types.OAuth2ClientCredsResponse {
	if c == nil {
		return nil
	}
	if c.Oauth2ClientCredentials == nil {
		return nil
	}
	return c.Oauth2ClientCredentials
}

func (c *CreateAuthConnectionsResponse) GetOauth2Jwt() *types.OAuth2JwtResponse {
	if c == nil {
		return nil
	}
	if c.Oauth2Jwt == nil {
		return nil
	}
	return c.Oauth2Jwt
}

func (c *CreateAuthConnectionsResponse) GetPrivateKeyJwt() *types.PrivateKeyJwtResponse {
	if c == nil {
		return nil
	}
	if c.PrivateKeyJwt == nil {
		return nil
	}
	return c.PrivateKeyJwt
}

func (c *CreateAuthConnectionsResponse) GetRefreshTokenAuth() *types.RefreshTokenAuthResponse {
	if c == nil {
		return nil
	}
	if c.RefreshTokenAuth == nil {
		return nil
	}
	return c.RefreshTokenAuth
}

func (c *CreateAuthConnectionsResponse) GetSlackBotAuth() *types.SlackBotAuthResponse {
	if c == nil {
		return nil
	}
	if c.SlackBotAuth == nil {
		return nil
	}
	return c.SlackBotAuth
}

func (c *CreateAuthConnectionsResponse) GetUrlSecret() *types.UrlSecretAuthResponse {
	if c == nil {
		return nil
	}
	if c.UrlSecret == nil {
		return nil
	}
	return c.UrlSecret
}

func (c *CreateAuthConnectionsResponse) GetWhatsappAuth() *types.WhatsAppAuthResponse {
	if c == nil {
		return nil
	}
	if c.WhatsappAuth == nil {
		return nil
	}
	return c.WhatsappAuth
}

func (c *CreateAuthConnectionsResponse) GetAuthType() string {
	if c == nil {
		return ""
	}
	return c.AuthType
}

func (c *CreateAuthConnectionsResponse) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		AuthType string `json:"auth_type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	c.AuthType = unmarshaler.AuthType
	if unmarshaler.AuthType == "" {
		return fmt.Errorf("%T did not include discriminant auth_type", c)
	}
	switch unmarshaler.AuthType {
	case "api_integration_oauth2_auth_code":
		value := new(types.ApiIntegrationOAuth2AuthCodeResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.ApiIntegrationOauth2AuthCode = value
	case "api_integration_oauth2_custom_app":
		value := new(types.ApiIntegrationOAuth2CustomAppResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.ApiIntegrationOauth2CustomApp = value
	case "basic_auth":
		value := new(types.BasicAuthResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.BasicAuth = value
	case "bearer_auth":
		value := new(types.BearerAuthResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.BearerAuth = value
	case "custom_header_auth":
		value := new(types.CustomHeaderAuthResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.CustomHeaderAuth = value
	case "mtls":
		value := new(types.MtlsAuthResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Mtls = value
	case "oauth2_client_credentials":
		value := new(types.OAuth2ClientCredsResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Oauth2ClientCredentials = value
	case "oauth2_jwt":
		value := new(types.OAuth2JwtResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Oauth2Jwt = value
	case "private_key_jwt":
		value := new(types.PrivateKeyJwtResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.PrivateKeyJwt = value
	case "refresh_token_auth":
		value := new(types.RefreshTokenAuthResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.RefreshTokenAuth = value
	case "slack_bot_auth":
		value := new(types.SlackBotAuthResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.SlackBotAuth = value
	case "url_secret":
		value := new(types.UrlSecretAuthResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.UrlSecret = value
	case "whatsapp_auth":
		value := new(types.WhatsAppAuthResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.WhatsappAuth = value
	}
	return nil
}

func (c CreateAuthConnectionsResponse) MarshalJSON() ([]byte, error) {
	if err := c.validate(); err != nil {
		return nil, err
	}
	if c.ApiIntegrationOauth2AuthCode != nil {
		return internal.MarshalJSONWithDiscriminant(c.ApiIntegrationOauth2AuthCode, "auth_type", "api_integration_oauth2_auth_code")
	}
	if c.ApiIntegrationOauth2CustomApp != nil {
		return internal.MarshalJSONWithDiscriminant(c.ApiIntegrationOauth2CustomApp, "auth_type", "api_integration_oauth2_custom_app")
	}
	if c.BasicAuth != nil {
		return internal.MarshalJSONWithDiscriminant(c.BasicAuth, "auth_type", "basic_auth")
	}
	if c.BearerAuth != nil {
		return internal.MarshalJSONWithDiscriminant(c.BearerAuth, "auth_type", "bearer_auth")
	}
	if c.CustomHeaderAuth != nil {
		return internal.MarshalJSONWithDiscriminant(c.CustomHeaderAuth, "auth_type", "custom_header_auth")
	}
	if c.Mtls != nil {
		return internal.MarshalJSONWithDiscriminant(c.Mtls, "auth_type", "mtls")
	}
	if c.Oauth2ClientCredentials != nil {
		return internal.MarshalJSONWithDiscriminant(c.Oauth2ClientCredentials, "auth_type", "oauth2_client_credentials")
	}
	if c.Oauth2Jwt != nil {
		return internal.MarshalJSONWithDiscriminant(c.Oauth2Jwt, "auth_type", "oauth2_jwt")
	}
	if c.PrivateKeyJwt != nil {
		return internal.MarshalJSONWithDiscriminant(c.PrivateKeyJwt, "auth_type", "private_key_jwt")
	}
	if c.RefreshTokenAuth != nil {
		return internal.MarshalJSONWithDiscriminant(c.RefreshTokenAuth, "auth_type", "refresh_token_auth")
	}
	if c.SlackBotAuth != nil {
		return internal.MarshalJSONWithDiscriminant(c.SlackBotAuth, "auth_type", "slack_bot_auth")
	}
	if c.UrlSecret != nil {
		return internal.MarshalJSONWithDiscriminant(c.UrlSecret, "auth_type", "url_secret")
	}
	if c.WhatsappAuth != nil {
		return internal.MarshalJSONWithDiscriminant(c.WhatsappAuth, "auth_type", "whatsapp_auth")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", c)
}

type CreateAuthConnectionsResponseVisitor interface {
	VisitApiIntegrationOauth2AuthCode(*types.ApiIntegrationOAuth2AuthCodeResponse) error
	VisitApiIntegrationOauth2CustomApp(*types.ApiIntegrationOAuth2CustomAppResponse) error
	VisitBasicAuth(*types.BasicAuthResponse) error
	VisitBearerAuth(*types.BearerAuthResponse) error
	VisitCustomHeaderAuth(*types.CustomHeaderAuthResponse) error
	VisitMtls(*types.MtlsAuthResponse) error
	VisitOauth2ClientCredentials(*types.OAuth2ClientCredsResponse) error
	VisitOauth2Jwt(*types.OAuth2JwtResponse) error
	VisitPrivateKeyJwt(*types.PrivateKeyJwtResponse) error
	VisitRefreshTokenAuth(*types.RefreshTokenAuthResponse) error
	VisitSlackBotAuth(*types.SlackBotAuthResponse) error
	VisitUrlSecret(*types.UrlSecretAuthResponse) error
	VisitWhatsappAuth(*types.WhatsAppAuthResponse) error
}

func (c *CreateAuthConnectionsResponse) Accept(visitor CreateAuthConnectionsResponseVisitor) error {
	if c.ApiIntegrationOauth2AuthCode != nil {
		return visitor.VisitApiIntegrationOauth2AuthCode(c.ApiIntegrationOauth2AuthCode)
	}
	if c.ApiIntegrationOauth2CustomApp != nil {
		return visitor.VisitApiIntegrationOauth2CustomApp(c.ApiIntegrationOauth2CustomApp)
	}
	if c.BasicAuth != nil {
		return visitor.VisitBasicAuth(c.BasicAuth)
	}
	if c.BearerAuth != nil {
		return visitor.VisitBearerAuth(c.BearerAuth)
	}
	if c.CustomHeaderAuth != nil {
		return visitor.VisitCustomHeaderAuth(c.CustomHeaderAuth)
	}
	if c.Mtls != nil {
		return visitor.VisitMtls(c.Mtls)
	}
	if c.Oauth2ClientCredentials != nil {
		return visitor.VisitOauth2ClientCredentials(c.Oauth2ClientCredentials)
	}
	if c.Oauth2Jwt != nil {
		return visitor.VisitOauth2Jwt(c.Oauth2Jwt)
	}
	if c.PrivateKeyJwt != nil {
		return visitor.VisitPrivateKeyJwt(c.PrivateKeyJwt)
	}
	if c.RefreshTokenAuth != nil {
		return visitor.VisitRefreshTokenAuth(c.RefreshTokenAuth)
	}
	if c.SlackBotAuth != nil {
		return visitor.VisitSlackBotAuth(c.SlackBotAuth)
	}
	if c.UrlSecret != nil {
		return visitor.VisitUrlSecret(c.UrlSecret)
	}
	if c.WhatsappAuth != nil {
		return visitor.VisitWhatsappAuth(c.WhatsappAuth)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", c)
}

func (c *CreateAuthConnectionsResponse) validate() error {
	if c == nil {
		return fmt.Errorf("type %T is nil", c)
	}
	var fields []string
	if c.ApiIntegrationOauth2AuthCode != nil {
		fields = append(fields, "api_integration_oauth2_auth_code")
	}
	if c.ApiIntegrationOauth2CustomApp != nil {
		fields = append(fields, "api_integration_oauth2_custom_app")
	}
	if c.BasicAuth != nil {
		fields = append(fields, "basic_auth")
	}
	if c.BearerAuth != nil {
		fields = append(fields, "bearer_auth")
	}
	if c.CustomHeaderAuth != nil {
		fields = append(fields, "custom_header_auth")
	}
	if c.Mtls != nil {
		fields = append(fields, "mtls")
	}
	if c.Oauth2ClientCredentials != nil {
		fields = append(fields, "oauth2_client_credentials")
	}
	if c.Oauth2Jwt != nil {
		fields = append(fields, "oauth2_jwt")
	}
	if c.PrivateKeyJwt != nil {
		fields = append(fields, "private_key_jwt")
	}
	if c.RefreshTokenAuth != nil {
		fields = append(fields, "refresh_token_auth")
	}
	if c.SlackBotAuth != nil {
		fields = append(fields, "slack_bot_auth")
	}
	if c.UrlSecret != nil {
		fields = append(fields, "url_secret")
	}
	if c.WhatsappAuth != nil {
		fields = append(fields, "whatsapp_auth")
	}
	if len(fields) == 0 {
		if c.AuthType != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", c, c.AuthType)
		}
		return fmt.Errorf("type %T is empty", c)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", c, fields)
	}
	if c.AuthType != "" {
		field := fields[0]
		if c.AuthType != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				c,
				c.AuthType,
				c,
			)
		}
	}
	return nil
}

// Auth connection to create
type CreateAuthConnectionsRequestBody struct {
	AuthType                string
	BasicAuth               *types.CreateBasicAuthRequest
	BearerAuth              *types.CreateBearerAuthRequest
	CustomHeaderAuth        *types.CreateCustomHeaderAuthRequest
	Mtls                    *types.CreateMtlsAuthRequest
	Oauth2ClientCredentials *types.CreateOAuth2ClientCredsRequest
	Oauth2Jwt               *types.CreateOAuth2JwtRequest
	PrivateKeyJwt           *types.CreatePrivateKeyJwtRequest
}

func (c *CreateAuthConnectionsRequestBody) GetBasicAuth() *types.CreateBasicAuthRequest {
	if c == nil {
		return nil
	}
	if c.BasicAuth == nil {
		return nil
	}
	return c.BasicAuth
}

func (c *CreateAuthConnectionsRequestBody) GetBearerAuth() *types.CreateBearerAuthRequest {
	if c == nil {
		return nil
	}
	if c.BearerAuth == nil {
		return nil
	}
	return c.BearerAuth
}

func (c *CreateAuthConnectionsRequestBody) GetCustomHeaderAuth() *types.CreateCustomHeaderAuthRequest {
	if c == nil {
		return nil
	}
	if c.CustomHeaderAuth == nil {
		return nil
	}
	return c.CustomHeaderAuth
}

func (c *CreateAuthConnectionsRequestBody) GetMtls() *types.CreateMtlsAuthRequest {
	if c == nil {
		return nil
	}
	if c.Mtls == nil {
		return nil
	}
	return c.Mtls
}

func (c *CreateAuthConnectionsRequestBody) GetOauth2ClientCredentials() *types.CreateOAuth2ClientCredsRequest {
	if c == nil {
		return nil
	}
	if c.Oauth2ClientCredentials == nil {
		return nil
	}
	return c.Oauth2ClientCredentials
}

func (c *CreateAuthConnectionsRequestBody) GetOauth2Jwt() *types.CreateOAuth2JwtRequest {
	if c == nil {
		return nil
	}
	if c.Oauth2Jwt == nil {
		return nil
	}
	return c.Oauth2Jwt
}

func (c *CreateAuthConnectionsRequestBody) GetPrivateKeyJwt() *types.CreatePrivateKeyJwtRequest {
	if c == nil {
		return nil
	}
	if c.PrivateKeyJwt == nil {
		return nil
	}
	return c.PrivateKeyJwt
}

func (c *CreateAuthConnectionsRequestBody) GetAuthType() string {
	if c == nil {
		return ""
	}
	return c.AuthType
}

func (c *CreateAuthConnectionsRequestBody) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		AuthType string `json:"auth_type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	c.AuthType = unmarshaler.AuthType
	if unmarshaler.AuthType == "" {
		return fmt.Errorf("%T did not include discriminant auth_type", c)
	}
	switch unmarshaler.AuthType {
	case "basic_auth":
		value := new(types.CreateBasicAuthRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.BasicAuth = value
	case "bearer_auth":
		value := new(types.CreateBearerAuthRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.BearerAuth = value
	case "custom_header_auth":
		value := new(types.CreateCustomHeaderAuthRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.CustomHeaderAuth = value
	case "mtls":
		value := new(types.CreateMtlsAuthRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Mtls = value
	case "oauth2_client_credentials":
		value := new(types.CreateOAuth2ClientCredsRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Oauth2ClientCredentials = value
	case "oauth2_jwt":
		value := new(types.CreateOAuth2JwtRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Oauth2Jwt = value
	case "private_key_jwt":
		value := new(types.CreatePrivateKeyJwtRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.PrivateKeyJwt = value
	}
	return nil
}

func (c CreateAuthConnectionsRequestBody) MarshalJSON() ([]byte, error) {
	if err := c.validate(); err != nil {
		return nil, err
	}
	if c.BasicAuth != nil {
		return internal.MarshalJSONWithDiscriminant(c.BasicAuth, "auth_type", "basic_auth")
	}
	if c.BearerAuth != nil {
		return internal.MarshalJSONWithDiscriminant(c.BearerAuth, "auth_type", "bearer_auth")
	}
	if c.CustomHeaderAuth != nil {
		return internal.MarshalJSONWithDiscriminant(c.CustomHeaderAuth, "auth_type", "custom_header_auth")
	}
	if c.Mtls != nil {
		return internal.MarshalJSONWithDiscriminant(c.Mtls, "auth_type", "mtls")
	}
	if c.Oauth2ClientCredentials != nil {
		return internal.MarshalJSONWithDiscriminant(c.Oauth2ClientCredentials, "auth_type", "oauth2_client_credentials")
	}
	if c.Oauth2Jwt != nil {
		return internal.MarshalJSONWithDiscriminant(c.Oauth2Jwt, "auth_type", "oauth2_jwt")
	}
	if c.PrivateKeyJwt != nil {
		return internal.MarshalJSONWithDiscriminant(c.PrivateKeyJwt, "auth_type", "private_key_jwt")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", c)
}

type CreateAuthConnectionsRequestBodyVisitor interface {
	VisitBasicAuth(*types.CreateBasicAuthRequest) error
	VisitBearerAuth(*types.CreateBearerAuthRequest) error
	VisitCustomHeaderAuth(*types.CreateCustomHeaderAuthRequest) error
	VisitMtls(*types.CreateMtlsAuthRequest) error
	VisitOauth2ClientCredentials(*types.CreateOAuth2ClientCredsRequest) error
	VisitOauth2Jwt(*types.CreateOAuth2JwtRequest) error
	VisitPrivateKeyJwt(*types.CreatePrivateKeyJwtRequest) error
}

func (c *CreateAuthConnectionsRequestBody) Accept(visitor CreateAuthConnectionsRequestBodyVisitor) error {
	if c.BasicAuth != nil {
		return visitor.VisitBasicAuth(c.BasicAuth)
	}
	if c.BearerAuth != nil {
		return visitor.VisitBearerAuth(c.BearerAuth)
	}
	if c.CustomHeaderAuth != nil {
		return visitor.VisitCustomHeaderAuth(c.CustomHeaderAuth)
	}
	if c.Mtls != nil {
		return visitor.VisitMtls(c.Mtls)
	}
	if c.Oauth2ClientCredentials != nil {
		return visitor.VisitOauth2ClientCredentials(c.Oauth2ClientCredentials)
	}
	if c.Oauth2Jwt != nil {
		return visitor.VisitOauth2Jwt(c.Oauth2Jwt)
	}
	if c.PrivateKeyJwt != nil {
		return visitor.VisitPrivateKeyJwt(c.PrivateKeyJwt)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", c)
}

func (c *CreateAuthConnectionsRequestBody) validate() error {
	if c == nil {
		return fmt.Errorf("type %T is nil", c)
	}
	var fields []string
	if c.BasicAuth != nil {
		fields = append(fields, "basic_auth")
	}
	if c.BearerAuth != nil {
		fields = append(fields, "bearer_auth")
	}
	if c.CustomHeaderAuth != nil {
		fields = append(fields, "custom_header_auth")
	}
	if c.Mtls != nil {
		fields = append(fields, "mtls")
	}
	if c.Oauth2ClientCredentials != nil {
		fields = append(fields, "oauth2_client_credentials")
	}
	if c.Oauth2Jwt != nil {
		fields = append(fields, "oauth2_jwt")
	}
	if c.PrivateKeyJwt != nil {
		fields = append(fields, "private_key_jwt")
	}
	if len(fields) == 0 {
		if c.AuthType != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", c, c.AuthType)
		}
		return fmt.Errorf("type %T is empty", c)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", c, fields)
	}
	if c.AuthType != "" {
		field := fields[0]
		if c.AuthType != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				c,
				c.AuthType,
				c,
			)
		}
	}
	return nil
}

// The type of auth connection config
type UpdateAuthConnectionsResponse struct {
	AuthType                      string
	ApiIntegrationOauth2AuthCode  *types.ApiIntegrationOAuth2AuthCodeResponse
	ApiIntegrationOauth2CustomApp *types.ApiIntegrationOAuth2CustomAppResponse
	BasicAuth                     *types.BasicAuthResponse
	BearerAuth                    *types.BearerAuthResponse
	CustomHeaderAuth              *types.CustomHeaderAuthResponse
	Mtls                          *types.MtlsAuthResponse
	Oauth2ClientCredentials       *types.OAuth2ClientCredsResponse
	Oauth2Jwt                     *types.OAuth2JwtResponse
	PrivateKeyJwt                 *types.PrivateKeyJwtResponse
	RefreshTokenAuth              *types.RefreshTokenAuthResponse
	SlackBotAuth                  *types.SlackBotAuthResponse
	UrlSecret                     *types.UrlSecretAuthResponse
	WhatsappAuth                  *types.WhatsAppAuthResponse
}

func (u *UpdateAuthConnectionsResponse) GetApiIntegrationOauth2AuthCode() *types.ApiIntegrationOAuth2AuthCodeResponse {
	if u == nil {
		return nil
	}
	if u.ApiIntegrationOauth2AuthCode == nil {
		return nil
	}
	return u.ApiIntegrationOauth2AuthCode
}

func (u *UpdateAuthConnectionsResponse) GetApiIntegrationOauth2CustomApp() *types.ApiIntegrationOAuth2CustomAppResponse {
	if u == nil {
		return nil
	}
	if u.ApiIntegrationOauth2CustomApp == nil {
		return nil
	}
	return u.ApiIntegrationOauth2CustomApp
}

func (u *UpdateAuthConnectionsResponse) GetBasicAuth() *types.BasicAuthResponse {
	if u == nil {
		return nil
	}
	if u.BasicAuth == nil {
		return nil
	}
	return u.BasicAuth
}

func (u *UpdateAuthConnectionsResponse) GetBearerAuth() *types.BearerAuthResponse {
	if u == nil {
		return nil
	}
	if u.BearerAuth == nil {
		return nil
	}
	return u.BearerAuth
}

func (u *UpdateAuthConnectionsResponse) GetCustomHeaderAuth() *types.CustomHeaderAuthResponse {
	if u == nil {
		return nil
	}
	if u.CustomHeaderAuth == nil {
		return nil
	}
	return u.CustomHeaderAuth
}

func (u *UpdateAuthConnectionsResponse) GetMtls() *types.MtlsAuthResponse {
	if u == nil {
		return nil
	}
	if u.Mtls == nil {
		return nil
	}
	return u.Mtls
}

func (u *UpdateAuthConnectionsResponse) GetOauth2ClientCredentials() *types.OAuth2ClientCredsResponse {
	if u == nil {
		return nil
	}
	if u.Oauth2ClientCredentials == nil {
		return nil
	}
	return u.Oauth2ClientCredentials
}

func (u *UpdateAuthConnectionsResponse) GetOauth2Jwt() *types.OAuth2JwtResponse {
	if u == nil {
		return nil
	}
	if u.Oauth2Jwt == nil {
		return nil
	}
	return u.Oauth2Jwt
}

func (u *UpdateAuthConnectionsResponse) GetPrivateKeyJwt() *types.PrivateKeyJwtResponse {
	if u == nil {
		return nil
	}
	if u.PrivateKeyJwt == nil {
		return nil
	}
	return u.PrivateKeyJwt
}

func (u *UpdateAuthConnectionsResponse) GetRefreshTokenAuth() *types.RefreshTokenAuthResponse {
	if u == nil {
		return nil
	}
	if u.RefreshTokenAuth == nil {
		return nil
	}
	return u.RefreshTokenAuth
}

func (u *UpdateAuthConnectionsResponse) GetSlackBotAuth() *types.SlackBotAuthResponse {
	if u == nil {
		return nil
	}
	if u.SlackBotAuth == nil {
		return nil
	}
	return u.SlackBotAuth
}

func (u *UpdateAuthConnectionsResponse) GetUrlSecret() *types.UrlSecretAuthResponse {
	if u == nil {
		return nil
	}
	if u.UrlSecret == nil {
		return nil
	}
	return u.UrlSecret
}

func (u *UpdateAuthConnectionsResponse) GetWhatsappAuth() *types.WhatsAppAuthResponse {
	if u == nil {
		return nil
	}
	if u.WhatsappAuth == nil {
		return nil
	}
	return u.WhatsappAuth
}

func (u *UpdateAuthConnectionsResponse) GetAuthType() string {
	if u == nil {
		return ""
	}
	return u.AuthType
}

func (u *UpdateAuthConnectionsResponse) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		AuthType string `json:"auth_type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	u.AuthType = unmarshaler.AuthType
	if unmarshaler.AuthType == "" {
		return fmt.Errorf("%T did not include discriminant auth_type", u)
	}
	switch unmarshaler.AuthType {
	case "api_integration_oauth2_auth_code":
		value := new(types.ApiIntegrationOAuth2AuthCodeResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.ApiIntegrationOauth2AuthCode = value
	case "api_integration_oauth2_custom_app":
		value := new(types.ApiIntegrationOAuth2CustomAppResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.ApiIntegrationOauth2CustomApp = value
	case "basic_auth":
		value := new(types.BasicAuthResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.BasicAuth = value
	case "bearer_auth":
		value := new(types.BearerAuthResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.BearerAuth = value
	case "custom_header_auth":
		value := new(types.CustomHeaderAuthResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.CustomHeaderAuth = value
	case "mtls":
		value := new(types.MtlsAuthResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Mtls = value
	case "oauth2_client_credentials":
		value := new(types.OAuth2ClientCredsResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Oauth2ClientCredentials = value
	case "oauth2_jwt":
		value := new(types.OAuth2JwtResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Oauth2Jwt = value
	case "private_key_jwt":
		value := new(types.PrivateKeyJwtResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.PrivateKeyJwt = value
	case "refresh_token_auth":
		value := new(types.RefreshTokenAuthResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.RefreshTokenAuth = value
	case "slack_bot_auth":
		value := new(types.SlackBotAuthResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.SlackBotAuth = value
	case "url_secret":
		value := new(types.UrlSecretAuthResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.UrlSecret = value
	case "whatsapp_auth":
		value := new(types.WhatsAppAuthResponse)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.WhatsappAuth = value
	}
	return nil
}

func (u UpdateAuthConnectionsResponse) MarshalJSON() ([]byte, error) {
	if err := u.validate(); err != nil {
		return nil, err
	}
	if u.ApiIntegrationOauth2AuthCode != nil {
		return internal.MarshalJSONWithDiscriminant(u.ApiIntegrationOauth2AuthCode, "auth_type", "api_integration_oauth2_auth_code")
	}
	if u.ApiIntegrationOauth2CustomApp != nil {
		return internal.MarshalJSONWithDiscriminant(u.ApiIntegrationOauth2CustomApp, "auth_type", "api_integration_oauth2_custom_app")
	}
	if u.BasicAuth != nil {
		return internal.MarshalJSONWithDiscriminant(u.BasicAuth, "auth_type", "basic_auth")
	}
	if u.BearerAuth != nil {
		return internal.MarshalJSONWithDiscriminant(u.BearerAuth, "auth_type", "bearer_auth")
	}
	if u.CustomHeaderAuth != nil {
		return internal.MarshalJSONWithDiscriminant(u.CustomHeaderAuth, "auth_type", "custom_header_auth")
	}
	if u.Mtls != nil {
		return internal.MarshalJSONWithDiscriminant(u.Mtls, "auth_type", "mtls")
	}
	if u.Oauth2ClientCredentials != nil {
		return internal.MarshalJSONWithDiscriminant(u.Oauth2ClientCredentials, "auth_type", "oauth2_client_credentials")
	}
	if u.Oauth2Jwt != nil {
		return internal.MarshalJSONWithDiscriminant(u.Oauth2Jwt, "auth_type", "oauth2_jwt")
	}
	if u.PrivateKeyJwt != nil {
		return internal.MarshalJSONWithDiscriminant(u.PrivateKeyJwt, "auth_type", "private_key_jwt")
	}
	if u.RefreshTokenAuth != nil {
		return internal.MarshalJSONWithDiscriminant(u.RefreshTokenAuth, "auth_type", "refresh_token_auth")
	}
	if u.SlackBotAuth != nil {
		return internal.MarshalJSONWithDiscriminant(u.SlackBotAuth, "auth_type", "slack_bot_auth")
	}
	if u.UrlSecret != nil {
		return internal.MarshalJSONWithDiscriminant(u.UrlSecret, "auth_type", "url_secret")
	}
	if u.WhatsappAuth != nil {
		return internal.MarshalJSONWithDiscriminant(u.WhatsappAuth, "auth_type", "whatsapp_auth")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", u)
}

type UpdateAuthConnectionsResponseVisitor interface {
	VisitApiIntegrationOauth2AuthCode(*types.ApiIntegrationOAuth2AuthCodeResponse) error
	VisitApiIntegrationOauth2CustomApp(*types.ApiIntegrationOAuth2CustomAppResponse) error
	VisitBasicAuth(*types.BasicAuthResponse) error
	VisitBearerAuth(*types.BearerAuthResponse) error
	VisitCustomHeaderAuth(*types.CustomHeaderAuthResponse) error
	VisitMtls(*types.MtlsAuthResponse) error
	VisitOauth2ClientCredentials(*types.OAuth2ClientCredsResponse) error
	VisitOauth2Jwt(*types.OAuth2JwtResponse) error
	VisitPrivateKeyJwt(*types.PrivateKeyJwtResponse) error
	VisitRefreshTokenAuth(*types.RefreshTokenAuthResponse) error
	VisitSlackBotAuth(*types.SlackBotAuthResponse) error
	VisitUrlSecret(*types.UrlSecretAuthResponse) error
	VisitWhatsappAuth(*types.WhatsAppAuthResponse) error
}

func (u *UpdateAuthConnectionsResponse) Accept(visitor UpdateAuthConnectionsResponseVisitor) error {
	if u.ApiIntegrationOauth2AuthCode != nil {
		return visitor.VisitApiIntegrationOauth2AuthCode(u.ApiIntegrationOauth2AuthCode)
	}
	if u.ApiIntegrationOauth2CustomApp != nil {
		return visitor.VisitApiIntegrationOauth2CustomApp(u.ApiIntegrationOauth2CustomApp)
	}
	if u.BasicAuth != nil {
		return visitor.VisitBasicAuth(u.BasicAuth)
	}
	if u.BearerAuth != nil {
		return visitor.VisitBearerAuth(u.BearerAuth)
	}
	if u.CustomHeaderAuth != nil {
		return visitor.VisitCustomHeaderAuth(u.CustomHeaderAuth)
	}
	if u.Mtls != nil {
		return visitor.VisitMtls(u.Mtls)
	}
	if u.Oauth2ClientCredentials != nil {
		return visitor.VisitOauth2ClientCredentials(u.Oauth2ClientCredentials)
	}
	if u.Oauth2Jwt != nil {
		return visitor.VisitOauth2Jwt(u.Oauth2Jwt)
	}
	if u.PrivateKeyJwt != nil {
		return visitor.VisitPrivateKeyJwt(u.PrivateKeyJwt)
	}
	if u.RefreshTokenAuth != nil {
		return visitor.VisitRefreshTokenAuth(u.RefreshTokenAuth)
	}
	if u.SlackBotAuth != nil {
		return visitor.VisitSlackBotAuth(u.SlackBotAuth)
	}
	if u.UrlSecret != nil {
		return visitor.VisitUrlSecret(u.UrlSecret)
	}
	if u.WhatsappAuth != nil {
		return visitor.VisitWhatsappAuth(u.WhatsappAuth)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", u)
}

func (u *UpdateAuthConnectionsResponse) validate() error {
	if u == nil {
		return fmt.Errorf("type %T is nil", u)
	}
	var fields []string
	if u.ApiIntegrationOauth2AuthCode != nil {
		fields = append(fields, "api_integration_oauth2_auth_code")
	}
	if u.ApiIntegrationOauth2CustomApp != nil {
		fields = append(fields, "api_integration_oauth2_custom_app")
	}
	if u.BasicAuth != nil {
		fields = append(fields, "basic_auth")
	}
	if u.BearerAuth != nil {
		fields = append(fields, "bearer_auth")
	}
	if u.CustomHeaderAuth != nil {
		fields = append(fields, "custom_header_auth")
	}
	if u.Mtls != nil {
		fields = append(fields, "mtls")
	}
	if u.Oauth2ClientCredentials != nil {
		fields = append(fields, "oauth2_client_credentials")
	}
	if u.Oauth2Jwt != nil {
		fields = append(fields, "oauth2_jwt")
	}
	if u.PrivateKeyJwt != nil {
		fields = append(fields, "private_key_jwt")
	}
	if u.RefreshTokenAuth != nil {
		fields = append(fields, "refresh_token_auth")
	}
	if u.SlackBotAuth != nil {
		fields = append(fields, "slack_bot_auth")
	}
	if u.UrlSecret != nil {
		fields = append(fields, "url_secret")
	}
	if u.WhatsappAuth != nil {
		fields = append(fields, "whatsapp_auth")
	}
	if len(fields) == 0 {
		if u.AuthType != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", u, u.AuthType)
		}
		return fmt.Errorf("type %T is empty", u)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", u, fields)
	}
	if u.AuthType != "" {
		field := fields[0]
		if u.AuthType != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				u,
				u.AuthType,
				u,
			)
		}
	}
	return nil
}

// Updated auth connection fields
type UpdateAuthConnectionsRequestBody struct {
	AuthType                string
	BasicAuth               *types.UpdateBasicAuthRequest
	BearerAuth              *types.UpdateBearerAuthRequest
	Oauth2ClientCredentials *types.UpdateOAuth2ClientCredsRequest
	Oauth2Jwt               *types.UpdateOAuth2JwtRequest
}

func (u *UpdateAuthConnectionsRequestBody) GetBasicAuth() *types.UpdateBasicAuthRequest {
	if u == nil {
		return nil
	}
	if u.BasicAuth == nil {
		return nil
	}
	return u.BasicAuth
}

func (u *UpdateAuthConnectionsRequestBody) GetBearerAuth() *types.UpdateBearerAuthRequest {
	if u == nil {
		return nil
	}
	if u.BearerAuth == nil {
		return nil
	}
	return u.BearerAuth
}

func (u *UpdateAuthConnectionsRequestBody) GetOauth2ClientCredentials() *types.UpdateOAuth2ClientCredsRequest {
	if u == nil {
		return nil
	}
	if u.Oauth2ClientCredentials == nil {
		return nil
	}
	return u.Oauth2ClientCredentials
}

func (u *UpdateAuthConnectionsRequestBody) GetOauth2Jwt() *types.UpdateOAuth2JwtRequest {
	if u == nil {
		return nil
	}
	if u.Oauth2Jwt == nil {
		return nil
	}
	return u.Oauth2Jwt
}

func (u *UpdateAuthConnectionsRequestBody) GetAuthType() string {
	if u == nil {
		return ""
	}
	return u.AuthType
}

func (u *UpdateAuthConnectionsRequestBody) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		AuthType string `json:"auth_type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	u.AuthType = unmarshaler.AuthType
	if unmarshaler.AuthType == "" {
		return fmt.Errorf("%T did not include discriminant auth_type", u)
	}
	switch unmarshaler.AuthType {
	case "basic_auth":
		value := new(types.UpdateBasicAuthRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.BasicAuth = value
	case "bearer_auth":
		value := new(types.UpdateBearerAuthRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.BearerAuth = value
	case "oauth2_client_credentials":
		value := new(types.UpdateOAuth2ClientCredsRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Oauth2ClientCredentials = value
	case "oauth2_jwt":
		value := new(types.UpdateOAuth2JwtRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Oauth2Jwt = value
	}
	return nil
}

func (u UpdateAuthConnectionsRequestBody) MarshalJSON() ([]byte, error) {
	if err := u.validate(); err != nil {
		return nil, err
	}
	if u.BasicAuth != nil {
		return internal.MarshalJSONWithDiscriminant(u.BasicAuth, "auth_type", "basic_auth")
	}
	if u.BearerAuth != nil {
		return internal.MarshalJSONWithDiscriminant(u.BearerAuth, "auth_type", "bearer_auth")
	}
	if u.Oauth2ClientCredentials != nil {
		return internal.MarshalJSONWithDiscriminant(u.Oauth2ClientCredentials, "auth_type", "oauth2_client_credentials")
	}
	if u.Oauth2Jwt != nil {
		return internal.MarshalJSONWithDiscriminant(u.Oauth2Jwt, "auth_type", "oauth2_jwt")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", u)
}

type UpdateAuthConnectionsRequestBodyVisitor interface {
	VisitBasicAuth(*types.UpdateBasicAuthRequest) error
	VisitBearerAuth(*types.UpdateBearerAuthRequest) error
	VisitOauth2ClientCredentials(*types.UpdateOAuth2ClientCredsRequest) error
	VisitOauth2Jwt(*types.UpdateOAuth2JwtRequest) error
}

func (u *UpdateAuthConnectionsRequestBody) Accept(visitor UpdateAuthConnectionsRequestBodyVisitor) error {
	if u.BasicAuth != nil {
		return visitor.VisitBasicAuth(u.BasicAuth)
	}
	if u.BearerAuth != nil {
		return visitor.VisitBearerAuth(u.BearerAuth)
	}
	if u.Oauth2ClientCredentials != nil {
		return visitor.VisitOauth2ClientCredentials(u.Oauth2ClientCredentials)
	}
	if u.Oauth2Jwt != nil {
		return visitor.VisitOauth2Jwt(u.Oauth2Jwt)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", u)
}

func (u *UpdateAuthConnectionsRequestBody) validate() error {
	if u == nil {
		return fmt.Errorf("type %T is nil", u)
	}
	var fields []string
	if u.BasicAuth != nil {
		fields = append(fields, "basic_auth")
	}
	if u.BearerAuth != nil {
		fields = append(fields, "bearer_auth")
	}
	if u.Oauth2ClientCredentials != nil {
		fields = append(fields, "oauth2_client_credentials")
	}
	if u.Oauth2Jwt != nil {
		fields = append(fields, "oauth2_jwt")
	}
	if len(fields) == 0 {
		if u.AuthType != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", u, u.AuthType)
		}
		return fmt.Errorf("type %T is empty", u)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", u, fields)
	}
	if u.AuthType != "" {
		field := fields[0]
		if u.AuthType != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				u,
				u.AuthType,
				u,
			)
		}
	}
	return nil
}
