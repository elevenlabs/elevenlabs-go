package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Response model for oauth2 client creds
type OAuth2ClientCredsResponse struct {
	Name        string            `json:"name" xml:"name" required:"true"`
	Provider    string            `json:"provider" xml:"provider" required:"true"`
	ClientId    string            `json:"client_id" xml:"client_id" required:"true"`
	TokenUrl    string            `json:"token_url" xml:"token_url" required:"true"`
	Scopes      []string          `json:"scopes,omitempty" xml:"scopes,omitempty"`
	ExtraParams map[string]string `json:"extra_params,omitempty"`
	// If True, send client credentials in Authorization header as Basic Auth instead of request body
	BasicAuthInHeader *bool                       `json:"basic_auth_in_header,omitempty" xml:"basic_auth_in_header,omitempty"`
	Id                string                      `json:"id" xml:"id" required:"true"`
	UsedBy            *AuthConnectionDependencies `json:"used_by,omitempty" xml:"used_by,omitempty"`
	// Single status field shared by every auth type's stored credential.
	//
	// OAuth values (``REFRESH_FAILED``, ``REVOKED``) are written by the OAuth
	// token-manager refresh path. ``CREDENTIAL_INVALID`` is written by the
	// tool execution path when an upstream response matches a credential's
	// ``failure_signatures`` entry (Bearer, Basic auth, etc.).
	Status          *AuthConnectionStatus `json:"status,omitempty" xml:"status,omitempty"`
	StatusDetail    *string               `json:"status_detail,omitempty" xml:"status_detail,omitempty"`
	StatusUpdatedAt *string               `json:"status_updated_at,omitempty" xml:"status_updated_at,omitempty"`
	// Custom headers configured for OAuth2 token requests
	CustomHeaders   map[string]string `json:"custom_headers,omitempty"`
	extraProperties map[string]interface{}
}

func (o *OAuth2ClientCredsResponse) GetName() (value string) {
	if o == nil {
		return
	}
	return o.Name
}

func (o *OAuth2ClientCredsResponse) GetProvider() (value string) {
	if o == nil {
		return
	}
	return o.Provider
}

func (o *OAuth2ClientCredsResponse) GetClientId() (value string) {
	if o == nil {
		return
	}
	return o.ClientId
}

func (o *OAuth2ClientCredsResponse) GetTokenUrl() (value string) {
	if o == nil {
		return
	}
	return o.TokenUrl
}

func (o *OAuth2ClientCredsResponse) GetScopes() (value []string) {
	if o == nil {
		return
	}
	return o.Scopes
}

func (o *OAuth2ClientCredsResponse) GetExtraParams() (value map[string]string) {
	if o == nil {
		return
	}
	return o.ExtraParams
}

func (o *OAuth2ClientCredsResponse) GetBasicAuthInHeader() (value *bool) {
	if o == nil {
		return
	}
	return o.BasicAuthInHeader
}

func (o *OAuth2ClientCredsResponse) GetId() (value string) {
	if o == nil {
		return
	}
	return o.Id
}

func (o *OAuth2ClientCredsResponse) GetUsedBy() (value *AuthConnectionDependencies) {
	if o == nil {
		return
	}
	return o.UsedBy
}

func (o *OAuth2ClientCredsResponse) GetStatus() (value *AuthConnectionStatus) {
	if o == nil {
		return
	}
	return o.Status
}

func (o *OAuth2ClientCredsResponse) GetStatusDetail() (value *string) {
	if o == nil {
		return
	}
	return o.StatusDetail
}

func (o *OAuth2ClientCredsResponse) GetStatusUpdatedAt() (value *string) {
	if o == nil {
		return
	}
	return o.StatusUpdatedAt
}

func (o *OAuth2ClientCredsResponse) GetCustomHeaders() (value map[string]string) {
	if o == nil {
		return
	}
	return o.CustomHeaders
}

func (o *OAuth2ClientCredsResponse) SetName(name string) {
	o.Name = name
}

func (o *OAuth2ClientCredsResponse) SetProvider(provider string) {
	o.Provider = provider
}

func (o *OAuth2ClientCredsResponse) SetClientId(clientId string) {
	o.ClientId = clientId
}

func (o *OAuth2ClientCredsResponse) SetTokenUrl(tokenUrl string) {
	o.TokenUrl = tokenUrl
}

func (o *OAuth2ClientCredsResponse) SetScopes(scopes []string) {
	o.Scopes = scopes
}

func (o *OAuth2ClientCredsResponse) SetExtraParams(extraParams map[string]string) {
	o.ExtraParams = extraParams
}

func (o *OAuth2ClientCredsResponse) SetBasicAuthInHeader(basicAuthInHeader *bool) {
	o.BasicAuthInHeader = basicAuthInHeader
}

func (o *OAuth2ClientCredsResponse) SetId(id string) {
	o.Id = id
}

func (o *OAuth2ClientCredsResponse) SetUsedBy(usedBy *AuthConnectionDependencies) {
	o.UsedBy = usedBy
}

func (o *OAuth2ClientCredsResponse) SetStatus(status *AuthConnectionStatus) {
	o.Status = status
}

func (o *OAuth2ClientCredsResponse) SetStatusDetail(statusDetail *string) {
	o.StatusDetail = statusDetail
}

func (o *OAuth2ClientCredsResponse) SetStatusUpdatedAt(statusUpdatedAt *string) {
	o.StatusUpdatedAt = statusUpdatedAt
}

func (o *OAuth2ClientCredsResponse) SetCustomHeaders(customHeaders map[string]string) {
	o.CustomHeaders = customHeaders
}

func (o *OAuth2ClientCredsResponse) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OAuth2ClientCredsResponse) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OAuth2ClientCredsResponse to string"
	}
	return string(jsonData)
}

func (o *OAuth2ClientCredsResponse) UnmarshalJSON(data []byte) error {
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

func (o *OAuth2ClientCredsResponse) MarshalJSON() ([]byte, error) {
	type embed OAuth2ClientCredsResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
