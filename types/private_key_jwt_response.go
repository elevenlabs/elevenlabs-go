package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Response model for Private Key JWT auth connections
type PrivateKeyJwtResponse struct {
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
	ExtraHeaders map[string]string           `json:"extra_headers,omitempty"`
	Id           string                      `json:"id" xml:"id" required:"true"`
	UsedBy       *AuthConnectionDependencies `json:"used_by,omitempty" xml:"used_by,omitempty"`
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

func (p *PrivateKeyJwtResponse) GetName() (value string) {
	if p == nil {
		return
	}
	return p.Name
}

func (p *PrivateKeyJwtResponse) GetProvider() (value string) {
	if p == nil {
		return
	}
	return p.Provider
}

func (p *PrivateKeyJwtResponse) GetAlgorithm() (value *OAuth2JwtResponseAlgorithm) {
	if p == nil {
		return
	}
	return p.Algorithm
}

func (p *PrivateKeyJwtResponse) GetKeyId() (value *string) {
	if p == nil {
		return
	}
	return p.KeyId
}

func (p *PrivateKeyJwtResponse) GetIssuer() (value string) {
	if p == nil {
		return
	}
	return p.Issuer
}

func (p *PrivateKeyJwtResponse) GetAudience() (value string) {
	if p == nil {
		return
	}
	return p.Audience
}

func (p *PrivateKeyJwtResponse) GetSubject() (value string) {
	if p == nil {
		return
	}
	return p.Subject
}

func (p *PrivateKeyJwtResponse) GetExpirationSeconds() (value *int) {
	if p == nil {
		return
	}
	return p.ExpirationSeconds
}

func (p *PrivateKeyJwtResponse) GetExtraParams() (value map[string]string) {
	if p == nil {
		return
	}
	return p.ExtraParams
}

func (p *PrivateKeyJwtResponse) GetExtraHeaders() (value map[string]string) {
	if p == nil {
		return
	}
	return p.ExtraHeaders
}

func (p *PrivateKeyJwtResponse) GetId() (value string) {
	if p == nil {
		return
	}
	return p.Id
}

func (p *PrivateKeyJwtResponse) GetUsedBy() (value *AuthConnectionDependencies) {
	if p == nil {
		return
	}
	return p.UsedBy
}

func (p *PrivateKeyJwtResponse) GetStatus() (value *AuthConnectionStatus) {
	if p == nil {
		return
	}
	return p.Status
}

func (p *PrivateKeyJwtResponse) GetStatusDetail() (value *string) {
	if p == nil {
		return
	}
	return p.StatusDetail
}

func (p *PrivateKeyJwtResponse) GetStatusUpdatedAt() (value *string) {
	if p == nil {
		return
	}
	return p.StatusUpdatedAt
}

func (p *PrivateKeyJwtResponse) SetName(name string) {
	p.Name = name
}

func (p *PrivateKeyJwtResponse) SetProvider(provider string) {
	p.Provider = provider
}

func (p *PrivateKeyJwtResponse) SetAlgorithm(algorithm *OAuth2JwtResponseAlgorithm) {
	p.Algorithm = algorithm
}

func (p *PrivateKeyJwtResponse) SetKeyId(keyId *string) {
	p.KeyId = keyId
}

func (p *PrivateKeyJwtResponse) SetIssuer(issuer string) {
	p.Issuer = issuer
}

func (p *PrivateKeyJwtResponse) SetAudience(audience string) {
	p.Audience = audience
}

func (p *PrivateKeyJwtResponse) SetSubject(subject string) {
	p.Subject = subject
}

func (p *PrivateKeyJwtResponse) SetExpirationSeconds(expirationSeconds *int) {
	p.ExpirationSeconds = expirationSeconds
}

func (p *PrivateKeyJwtResponse) SetExtraParams(extraParams map[string]string) {
	p.ExtraParams = extraParams
}

func (p *PrivateKeyJwtResponse) SetExtraHeaders(extraHeaders map[string]string) {
	p.ExtraHeaders = extraHeaders
}

func (p *PrivateKeyJwtResponse) SetId(id string) {
	p.Id = id
}

func (p *PrivateKeyJwtResponse) SetUsedBy(usedBy *AuthConnectionDependencies) {
	p.UsedBy = usedBy
}

func (p *PrivateKeyJwtResponse) SetStatus(status *AuthConnectionStatus) {
	p.Status = status
}

func (p *PrivateKeyJwtResponse) SetStatusDetail(statusDetail *string) {
	p.StatusDetail = statusDetail
}

func (p *PrivateKeyJwtResponse) SetStatusUpdatedAt(statusUpdatedAt *string) {
	p.StatusUpdatedAt = statusUpdatedAt
}

func (p *PrivateKeyJwtResponse) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PrivateKeyJwtResponse) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PrivateKeyJwtResponse to string"
	}
	return string(jsonData)
}

func (p *PrivateKeyJwtResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, p); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *p, "auth_type")
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PrivateKeyJwtResponse) MarshalJSON() ([]byte, error) {
	type embed PrivateKeyJwtResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
