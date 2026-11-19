package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type UpdateOAuth2JwtRequest struct {
	AuthType           *string                                   `json:"auth_type,omitempty" xml:"auth_type,omitempty"`
	Provider           *string                                   `json:"provider,omitempty" xml:"provider,omitempty"`
	Algorithm          *UpdateOAuth2JwtRequestAlgorithm          `json:"algorithm,omitempty" xml:"algorithm,omitempty"`
	KeyId              *string                                   `json:"key_id,omitempty" xml:"key_id,omitempty"`
	Issuer             *string                                   `json:"issuer,omitempty" xml:"issuer,omitempty"`
	Audience           *string                                   `json:"audience,omitempty" xml:"audience,omitempty"`
	Subject            *string                                   `json:"subject,omitempty" xml:"subject,omitempty"`
	ExpirationSeconds  *int                                      `json:"expiration_seconds,omitempty" xml:"expiration_seconds,omitempty" min:"60" max:"86400"`
	ExtraParams        map[string]*string                        `json:"extra_params,omitempty"`
	Scopes             []string                                  `json:"scopes,omitempty" xml:"scopes,omitempty"`
	TokenResponseField *UpdateOAuth2JwtRequestTokenResponseField `json:"token_response_field,omitempty" xml:"token_response_field,omitempty"`
	JwtUsage           *UpdateOAuth2JwtRequestJwtUsage           `json:"jwt_usage,omitempty" xml:"jwt_usage,omitempty"`
	ExtraHeaders       map[string]*string                        `json:"extra_headers,omitempty"`
	SecretKey          *string                                   `json:"secret_key,omitempty" xml:"secret_key,omitempty"`
	extraProperties    map[string]interface{}
}

func (u *UpdateOAuth2JwtRequest) GetAuthType() (value *string) {
	if u == nil {
		return
	}
	return u.AuthType
}

func (u *UpdateOAuth2JwtRequest) GetProvider() (value *string) {
	if u == nil {
		return
	}
	return u.Provider
}

func (u *UpdateOAuth2JwtRequest) GetAlgorithm() (value *UpdateOAuth2JwtRequestAlgorithm) {
	if u == nil {
		return
	}
	return u.Algorithm
}

func (u *UpdateOAuth2JwtRequest) GetKeyId() (value *string) {
	if u == nil {
		return
	}
	return u.KeyId
}

func (u *UpdateOAuth2JwtRequest) GetIssuer() (value *string) {
	if u == nil {
		return
	}
	return u.Issuer
}

func (u *UpdateOAuth2JwtRequest) GetAudience() (value *string) {
	if u == nil {
		return
	}
	return u.Audience
}

func (u *UpdateOAuth2JwtRequest) GetSubject() (value *string) {
	if u == nil {
		return
	}
	return u.Subject
}

func (u *UpdateOAuth2JwtRequest) GetExpirationSeconds() (value *int) {
	if u == nil {
		return
	}
	return u.ExpirationSeconds
}

func (u *UpdateOAuth2JwtRequest) GetExtraParams() (value map[string]*string) {
	if u == nil {
		return
	}
	return u.ExtraParams
}

func (u *UpdateOAuth2JwtRequest) GetScopes() (value []string) {
	if u == nil {
		return
	}
	return u.Scopes
}

func (u *UpdateOAuth2JwtRequest) GetTokenResponseField() (value *UpdateOAuth2JwtRequestTokenResponseField) {
	if u == nil {
		return
	}
	return u.TokenResponseField
}

func (u *UpdateOAuth2JwtRequest) GetJwtUsage() (value *UpdateOAuth2JwtRequestJwtUsage) {
	if u == nil {
		return
	}
	return u.JwtUsage
}

func (u *UpdateOAuth2JwtRequest) GetExtraHeaders() (value map[string]*string) {
	if u == nil {
		return
	}
	return u.ExtraHeaders
}

func (u *UpdateOAuth2JwtRequest) GetSecretKey() (value *string) {
	if u == nil {
		return
	}
	return u.SecretKey
}

func (u *UpdateOAuth2JwtRequest) SetAuthType(authType *string) {
	u.AuthType = authType
}

func (u *UpdateOAuth2JwtRequest) SetProvider(provider *string) {
	u.Provider = provider
}

func (u *UpdateOAuth2JwtRequest) SetAlgorithm(algorithm *UpdateOAuth2JwtRequestAlgorithm) {
	u.Algorithm = algorithm
}

func (u *UpdateOAuth2JwtRequest) SetKeyId(keyId *string) {
	u.KeyId = keyId
}

func (u *UpdateOAuth2JwtRequest) SetIssuer(issuer *string) {
	u.Issuer = issuer
}

func (u *UpdateOAuth2JwtRequest) SetAudience(audience *string) {
	u.Audience = audience
}

func (u *UpdateOAuth2JwtRequest) SetSubject(subject *string) {
	u.Subject = subject
}

func (u *UpdateOAuth2JwtRequest) SetExpirationSeconds(expirationSeconds *int) {
	u.ExpirationSeconds = expirationSeconds
}

func (u *UpdateOAuth2JwtRequest) SetExtraParams(extraParams map[string]*string) {
	u.ExtraParams = extraParams
}

func (u *UpdateOAuth2JwtRequest) SetScopes(scopes []string) {
	u.Scopes = scopes
}

func (u *UpdateOAuth2JwtRequest) SetTokenResponseField(tokenResponseField *UpdateOAuth2JwtRequestTokenResponseField) {
	u.TokenResponseField = tokenResponseField
}

func (u *UpdateOAuth2JwtRequest) SetJwtUsage(jwtUsage *UpdateOAuth2JwtRequestJwtUsage) {
	u.JwtUsage = jwtUsage
}

func (u *UpdateOAuth2JwtRequest) SetExtraHeaders(extraHeaders map[string]*string) {
	u.ExtraHeaders = extraHeaders
}

func (u *UpdateOAuth2JwtRequest) SetSecretKey(secretKey *string) {
	u.SecretKey = secretKey
}

func (u *UpdateOAuth2JwtRequest) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UpdateOAuth2JwtRequest) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateOAuth2JwtRequest to string"
	}
	return string(jsonData)
}

func (u *UpdateOAuth2JwtRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, u); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UpdateOAuth2JwtRequest) MarshalJSON() ([]byte, error) {
	type embed UpdateOAuth2JwtRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
