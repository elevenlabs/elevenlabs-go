package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Response model for OAuth2 JWT auth connections
type OAuth2JwtResponse struct {
	Name     string `json:"name" xml:"name" required:"true"`
	Provider string `json:"provider" xml:"provider" required:"true"`
	// JWT signing algorithm
	Algorithm *OAuth2JwtResponseAlgorithm `json:"algorithm,omitempty" xml:"algorithm,omitempty"`
	// Key ID (kid) for JWT header - useful for key rotation
	KeyId *string `json:"key_id,omitempty" xml:"key_id,omitempty"`
	// JWT issuer (iss claim)
	Issuer string `json:"issuer" xml:"issuer" required:"true"`
	// JWT audience (aud claim)
	Audience string `json:"audience" xml:"audience" required:"true"`
	// JWT subject (sub claim)
	Subject string `json:"subject" xml:"subject" required:"true"`
	// Token expiration time in seconds
	ExpirationSeconds *int `json:"expiration_seconds,omitempty" xml:"expiration_seconds,omitempty" min:"60" max:"86400"`
	// Additional custom claims to include in the JWT
	ExtraParams map[string]string `json:"extra_params,omitempty"`
	// Additional JWT header parameters (e.g. x5t). 'alg' and 'kid' are set from algorithm and key_id.
	ExtraHeaders map[string]string `json:"extra_headers,omitempty"`
	// Token endpoint URL for exchanging JWT for access token
	TokenUrl string `json:"token_url" xml:"token_url" required:"true"`
	// OAuth2 scopes to request when exchanging JWT for access token
	Scopes []string `json:"scopes,omitempty" xml:"scopes,omitempty"`
	// Token field to extract from the token endpoint response.
	TokenResponseField *OAuth2JwtResponseTokenResponseField `json:"token_response_field,omitempty" xml:"token_response_field,omitempty"`
	// How the JWT is used at the token endpoint. 'authorization_grant' sends it as a jwt-bearer grant. 'client_authentication' sends it as a client_assertion on a client_credentials request, with the subject as client_id and scopes in the request body.
	JwtUsage *OAuth2JwtResponseJwtUsage  `json:"jwt_usage,omitempty" xml:"jwt_usage,omitempty"`
	Id       string                      `json:"id" xml:"id" required:"true"`
	UsedBy   *AuthConnectionDependencies `json:"used_by,omitempty" xml:"used_by,omitempty"`
	// Single status field shared by every auth type's stored credential.
	//
	// OAuth values (``REFRESH_FAILED``, ``REVOKED``) are written by the OAuth
	// token-manager refresh path. ``CREDENTIAL_INVALID`` is written by the
	// tool execution path when an upstream response matches a credential's
	// ``failure_signatures`` entry (Bearer, Basic auth, etc.).
	Status          *AuthConnectionStatus `json:"status,omitempty" xml:"status,omitempty"`
	StatusDetail    *string               `json:"status_detail,omitempty" xml:"status_detail,omitempty"`
	StatusUpdatedAt *string               `json:"status_updated_at,omitempty" xml:"status_updated_at,omitempty"`
	extraProperties map[string]interface{}
}

func (o *OAuth2JwtResponse) GetName() (value string) {
	if o == nil {
		return
	}
	return o.Name
}

func (o *OAuth2JwtResponse) GetProvider() (value string) {
	if o == nil {
		return
	}
	return o.Provider
}

func (o *OAuth2JwtResponse) GetAlgorithm() (value *OAuth2JwtResponseAlgorithm) {
	if o == nil {
		return
	}
	return o.Algorithm
}

func (o *OAuth2JwtResponse) GetKeyId() (value *string) {
	if o == nil {
		return
	}
	return o.KeyId
}

func (o *OAuth2JwtResponse) GetIssuer() (value string) {
	if o == nil {
		return
	}
	return o.Issuer
}

func (o *OAuth2JwtResponse) GetAudience() (value string) {
	if o == nil {
		return
	}
	return o.Audience
}

func (o *OAuth2JwtResponse) GetSubject() (value string) {
	if o == nil {
		return
	}
	return o.Subject
}

func (o *OAuth2JwtResponse) GetExpirationSeconds() (value *int) {
	if o == nil {
		return
	}
	return o.ExpirationSeconds
}

func (o *OAuth2JwtResponse) GetExtraParams() (value map[string]string) {
	if o == nil {
		return
	}
	return o.ExtraParams
}

func (o *OAuth2JwtResponse) GetExtraHeaders() (value map[string]string) {
	if o == nil {
		return
	}
	return o.ExtraHeaders
}

func (o *OAuth2JwtResponse) GetTokenUrl() (value string) {
	if o == nil {
		return
	}
	return o.TokenUrl
}

func (o *OAuth2JwtResponse) GetScopes() (value []string) {
	if o == nil {
		return
	}
	return o.Scopes
}

func (o *OAuth2JwtResponse) GetTokenResponseField() (value *OAuth2JwtResponseTokenResponseField) {
	if o == nil {
		return
	}
	return o.TokenResponseField
}

func (o *OAuth2JwtResponse) GetJwtUsage() (value *OAuth2JwtResponseJwtUsage) {
	if o == nil {
		return
	}
	return o.JwtUsage
}

func (o *OAuth2JwtResponse) GetId() (value string) {
	if o == nil {
		return
	}
	return o.Id
}

func (o *OAuth2JwtResponse) GetUsedBy() (value *AuthConnectionDependencies) {
	if o == nil {
		return
	}
	return o.UsedBy
}

func (o *OAuth2JwtResponse) GetStatus() (value *AuthConnectionStatus) {
	if o == nil {
		return
	}
	return o.Status
}

func (o *OAuth2JwtResponse) GetStatusDetail() (value *string) {
	if o == nil {
		return
	}
	return o.StatusDetail
}

func (o *OAuth2JwtResponse) GetStatusUpdatedAt() (value *string) {
	if o == nil {
		return
	}
	return o.StatusUpdatedAt
}

func (o *OAuth2JwtResponse) SetName(name string) {
	o.Name = name
}

func (o *OAuth2JwtResponse) SetProvider(provider string) {
	o.Provider = provider
}

func (o *OAuth2JwtResponse) SetAlgorithm(algorithm *OAuth2JwtResponseAlgorithm) {
	o.Algorithm = algorithm
}

func (o *OAuth2JwtResponse) SetKeyId(keyId *string) {
	o.KeyId = keyId
}

func (o *OAuth2JwtResponse) SetIssuer(issuer string) {
	o.Issuer = issuer
}

func (o *OAuth2JwtResponse) SetAudience(audience string) {
	o.Audience = audience
}

func (o *OAuth2JwtResponse) SetSubject(subject string) {
	o.Subject = subject
}

func (o *OAuth2JwtResponse) SetExpirationSeconds(expirationSeconds *int) {
	o.ExpirationSeconds = expirationSeconds
}

func (o *OAuth2JwtResponse) SetExtraParams(extraParams map[string]string) {
	o.ExtraParams = extraParams
}

func (o *OAuth2JwtResponse) SetExtraHeaders(extraHeaders map[string]string) {
	o.ExtraHeaders = extraHeaders
}

func (o *OAuth2JwtResponse) SetTokenUrl(tokenUrl string) {
	o.TokenUrl = tokenUrl
}

func (o *OAuth2JwtResponse) SetScopes(scopes []string) {
	o.Scopes = scopes
}

func (o *OAuth2JwtResponse) SetTokenResponseField(tokenResponseField *OAuth2JwtResponseTokenResponseField) {
	o.TokenResponseField = tokenResponseField
}

func (o *OAuth2JwtResponse) SetJwtUsage(jwtUsage *OAuth2JwtResponseJwtUsage) {
	o.JwtUsage = jwtUsage
}

func (o *OAuth2JwtResponse) SetId(id string) {
	o.Id = id
}

func (o *OAuth2JwtResponse) SetUsedBy(usedBy *AuthConnectionDependencies) {
	o.UsedBy = usedBy
}

func (o *OAuth2JwtResponse) SetStatus(status *AuthConnectionStatus) {
	o.Status = status
}

func (o *OAuth2JwtResponse) SetStatusDetail(statusDetail *string) {
	o.StatusDetail = statusDetail
}

func (o *OAuth2JwtResponse) SetStatusUpdatedAt(statusUpdatedAt *string) {
	o.StatusUpdatedAt = statusUpdatedAt
}

func (o *OAuth2JwtResponse) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OAuth2JwtResponse) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OAuth2JwtResponse to string"
	}
	return string(jsonData)
}

func (o *OAuth2JwtResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, o); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *o, "auth_type")
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OAuth2JwtResponse) MarshalJSON() ([]byte, error) {
	type embed OAuth2JwtResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
