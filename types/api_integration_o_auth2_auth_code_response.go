package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Response model for integration-managed OAuth2 Auth Code auth connections
type ApiIntegrationOAuth2AuthCodeResponse struct {
	Name     string   `json:"name" xml:"name" required:"true"`
	Provider string   `json:"provider" xml:"provider" required:"true"`
	TokenUrl string   `json:"token_url" xml:"token_url" required:"true"`
	Scopes   []string `json:"scopes,omitempty" xml:"scopes,omitempty"`
	// Separator for scopes
	ScopeSeparator *ApiIntegrationOAuth2AuthCodeResponseScopeSeparator `json:"scope_separator,omitempty" xml:"scope_separator,omitempty"`
	// ISO 8601 timestamp of when the access token expires
	ExpiresAt     string                      `json:"expires_at" xml:"expires_at" required:"true"`
	IntegrationId string                      `json:"integration_id" xml:"integration_id" required:"true"`
	CredentialId  string                      `json:"credential_id" xml:"credential_id" required:"true"`
	Id            string                      `json:"id" xml:"id" required:"true"`
	UsedBy        *AuthConnectionDependencies `json:"used_by,omitempty" xml:"used_by,omitempty"`
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

func (a *ApiIntegrationOAuth2AuthCodeResponse) GetName() (value string) {
	if a == nil {
		return
	}
	return a.Name
}

func (a *ApiIntegrationOAuth2AuthCodeResponse) GetProvider() (value string) {
	if a == nil {
		return
	}
	return a.Provider
}

func (a *ApiIntegrationOAuth2AuthCodeResponse) GetTokenUrl() (value string) {
	if a == nil {
		return
	}
	return a.TokenUrl
}

func (a *ApiIntegrationOAuth2AuthCodeResponse) GetScopes() (value []string) {
	if a == nil {
		return
	}
	return a.Scopes
}

func (a *ApiIntegrationOAuth2AuthCodeResponse) GetScopeSeparator() (value *ApiIntegrationOAuth2AuthCodeResponseScopeSeparator) {
	if a == nil {
		return
	}
	return a.ScopeSeparator
}

func (a *ApiIntegrationOAuth2AuthCodeResponse) GetExpiresAt() (value string) {
	if a == nil {
		return
	}
	return a.ExpiresAt
}

func (a *ApiIntegrationOAuth2AuthCodeResponse) GetIntegrationId() (value string) {
	if a == nil {
		return
	}
	return a.IntegrationId
}

func (a *ApiIntegrationOAuth2AuthCodeResponse) GetCredentialId() (value string) {
	if a == nil {
		return
	}
	return a.CredentialId
}

func (a *ApiIntegrationOAuth2AuthCodeResponse) GetId() (value string) {
	if a == nil {
		return
	}
	return a.Id
}

func (a *ApiIntegrationOAuth2AuthCodeResponse) GetUsedBy() (value *AuthConnectionDependencies) {
	if a == nil {
		return
	}
	return a.UsedBy
}

func (a *ApiIntegrationOAuth2AuthCodeResponse) GetStatus() (value *AuthConnectionStatus) {
	if a == nil {
		return
	}
	return a.Status
}

func (a *ApiIntegrationOAuth2AuthCodeResponse) GetStatusDetail() (value *string) {
	if a == nil {
		return
	}
	return a.StatusDetail
}

func (a *ApiIntegrationOAuth2AuthCodeResponse) GetStatusUpdatedAt() (value *string) {
	if a == nil {
		return
	}
	return a.StatusUpdatedAt
}

func (a *ApiIntegrationOAuth2AuthCodeResponse) SetName(name string) {
	a.Name = name
}

func (a *ApiIntegrationOAuth2AuthCodeResponse) SetProvider(provider string) {
	a.Provider = provider
}

func (a *ApiIntegrationOAuth2AuthCodeResponse) SetTokenUrl(tokenUrl string) {
	a.TokenUrl = tokenUrl
}

func (a *ApiIntegrationOAuth2AuthCodeResponse) SetScopes(scopes []string) {
	a.Scopes = scopes
}

func (a *ApiIntegrationOAuth2AuthCodeResponse) SetScopeSeparator(scopeSeparator *ApiIntegrationOAuth2AuthCodeResponseScopeSeparator) {
	a.ScopeSeparator = scopeSeparator
}

func (a *ApiIntegrationOAuth2AuthCodeResponse) SetExpiresAt(expiresAt string) {
	a.ExpiresAt = expiresAt
}

func (a *ApiIntegrationOAuth2AuthCodeResponse) SetIntegrationId(integrationId string) {
	a.IntegrationId = integrationId
}

func (a *ApiIntegrationOAuth2AuthCodeResponse) SetCredentialId(credentialId string) {
	a.CredentialId = credentialId
}

func (a *ApiIntegrationOAuth2AuthCodeResponse) SetId(id string) {
	a.Id = id
}

func (a *ApiIntegrationOAuth2AuthCodeResponse) SetUsedBy(usedBy *AuthConnectionDependencies) {
	a.UsedBy = usedBy
}

func (a *ApiIntegrationOAuth2AuthCodeResponse) SetStatus(status *AuthConnectionStatus) {
	a.Status = status
}

func (a *ApiIntegrationOAuth2AuthCodeResponse) SetStatusDetail(statusDetail *string) {
	a.StatusDetail = statusDetail
}

func (a *ApiIntegrationOAuth2AuthCodeResponse) SetStatusUpdatedAt(statusUpdatedAt *string) {
	a.StatusUpdatedAt = statusUpdatedAt
}

func (a *ApiIntegrationOAuth2AuthCodeResponse) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a ApiIntegrationOAuth2AuthCodeResponse) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: ApiIntegrationOAuth2AuthCodeResponse to string"
	}
	return string(jsonData)
}

func (a *ApiIntegrationOAuth2AuthCodeResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "auth_type")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *ApiIntegrationOAuth2AuthCodeResponse) MarshalJSON() ([]byte, error) {
	type embed ApiIntegrationOAuth2AuthCodeResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
