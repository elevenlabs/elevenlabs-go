package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Response model for user-owned OAuth2 Custom App auth connections
type ApiIntegrationOAuth2CustomAppResponse struct {
	Name     string   `json:"name" xml:"name" required:"true"`
	Provider string   `json:"provider" xml:"provider" required:"true"`
	TokenUrl string   `json:"token_url" xml:"token_url" required:"true"`
	Scopes   []string `json:"scopes,omitempty" xml:"scopes,omitempty"`
	// Separator for scopes
	ScopeSeparator *ApiIntegrationOAuth2AuthCodeResponseScopeSeparator `json:"scope_separator,omitempty" xml:"scope_separator,omitempty"`
	// ISO 8601 timestamp of when the access token expires
	ExpiresAt     string `json:"expires_at" xml:"expires_at" required:"true"`
	IntegrationId string `json:"integration_id" xml:"integration_id" required:"true"`
	CredentialId  string `json:"credential_id" xml:"credential_id" required:"true"`
	// OAuth client ID (rendered from template if credential uses templated credentials, None for legacy connections)
	ClientId string                      `json:"client_id" xml:"client_id" required:"true"`
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

func (a *ApiIntegrationOAuth2CustomAppResponse) GetName() (value string) {
	if a == nil {
		return
	}
	return a.Name
}

func (a *ApiIntegrationOAuth2CustomAppResponse) GetProvider() (value string) {
	if a == nil {
		return
	}
	return a.Provider
}

func (a *ApiIntegrationOAuth2CustomAppResponse) GetTokenUrl() (value string) {
	if a == nil {
		return
	}
	return a.TokenUrl
}

func (a *ApiIntegrationOAuth2CustomAppResponse) GetScopes() (value []string) {
	if a == nil {
		return
	}
	return a.Scopes
}

func (a *ApiIntegrationOAuth2CustomAppResponse) GetScopeSeparator() (value *ApiIntegrationOAuth2AuthCodeResponseScopeSeparator) {
	if a == nil {
		return
	}
	return a.ScopeSeparator
}

func (a *ApiIntegrationOAuth2CustomAppResponse) GetExpiresAt() (value string) {
	if a == nil {
		return
	}
	return a.ExpiresAt
}

func (a *ApiIntegrationOAuth2CustomAppResponse) GetIntegrationId() (value string) {
	if a == nil {
		return
	}
	return a.IntegrationId
}

func (a *ApiIntegrationOAuth2CustomAppResponse) GetCredentialId() (value string) {
	if a == nil {
		return
	}
	return a.CredentialId
}

func (a *ApiIntegrationOAuth2CustomAppResponse) GetClientId() (value string) {
	if a == nil {
		return
	}
	return a.ClientId
}

func (a *ApiIntegrationOAuth2CustomAppResponse) GetId() (value string) {
	if a == nil {
		return
	}
	return a.Id
}

func (a *ApiIntegrationOAuth2CustomAppResponse) GetUsedBy() (value *AuthConnectionDependencies) {
	if a == nil {
		return
	}
	return a.UsedBy
}

func (a *ApiIntegrationOAuth2CustomAppResponse) GetStatus() (value *AuthConnectionStatus) {
	if a == nil {
		return
	}
	return a.Status
}

func (a *ApiIntegrationOAuth2CustomAppResponse) GetStatusDetail() (value *string) {
	if a == nil {
		return
	}
	return a.StatusDetail
}

func (a *ApiIntegrationOAuth2CustomAppResponse) GetStatusUpdatedAt() (value *string) {
	if a == nil {
		return
	}
	return a.StatusUpdatedAt
}

func (a *ApiIntegrationOAuth2CustomAppResponse) SetName(name string) {
	a.Name = name
}

func (a *ApiIntegrationOAuth2CustomAppResponse) SetProvider(provider string) {
	a.Provider = provider
}

func (a *ApiIntegrationOAuth2CustomAppResponse) SetTokenUrl(tokenUrl string) {
	a.TokenUrl = tokenUrl
}

func (a *ApiIntegrationOAuth2CustomAppResponse) SetScopes(scopes []string) {
	a.Scopes = scopes
}

func (a *ApiIntegrationOAuth2CustomAppResponse) SetScopeSeparator(scopeSeparator *ApiIntegrationOAuth2AuthCodeResponseScopeSeparator) {
	a.ScopeSeparator = scopeSeparator
}

func (a *ApiIntegrationOAuth2CustomAppResponse) SetExpiresAt(expiresAt string) {
	a.ExpiresAt = expiresAt
}

func (a *ApiIntegrationOAuth2CustomAppResponse) SetIntegrationId(integrationId string) {
	a.IntegrationId = integrationId
}

func (a *ApiIntegrationOAuth2CustomAppResponse) SetCredentialId(credentialId string) {
	a.CredentialId = credentialId
}

func (a *ApiIntegrationOAuth2CustomAppResponse) SetClientId(clientId string) {
	a.ClientId = clientId
}

func (a *ApiIntegrationOAuth2CustomAppResponse) SetId(id string) {
	a.Id = id
}

func (a *ApiIntegrationOAuth2CustomAppResponse) SetUsedBy(usedBy *AuthConnectionDependencies) {
	a.UsedBy = usedBy
}

func (a *ApiIntegrationOAuth2CustomAppResponse) SetStatus(status *AuthConnectionStatus) {
	a.Status = status
}

func (a *ApiIntegrationOAuth2CustomAppResponse) SetStatusDetail(statusDetail *string) {
	a.StatusDetail = statusDetail
}

func (a *ApiIntegrationOAuth2CustomAppResponse) SetStatusUpdatedAt(statusUpdatedAt *string) {
	a.StatusUpdatedAt = statusUpdatedAt
}

func (a *ApiIntegrationOAuth2CustomAppResponse) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a ApiIntegrationOAuth2CustomAppResponse) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: ApiIntegrationOAuth2CustomAppResponse to string"
	}
	return string(jsonData)
}

func (a *ApiIntegrationOAuth2CustomAppResponse) UnmarshalJSON(data []byte) error {
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

func (a *ApiIntegrationOAuth2CustomAppResponse) MarshalJSON() ([]byte, error) {
	type embed ApiIntegrationOAuth2CustomAppResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
