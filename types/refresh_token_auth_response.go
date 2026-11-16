package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Response model for OAuth2 refresh-token-grant auth connections
type RefreshTokenAuthResponse struct {
	Name        string                      `json:"name" xml:"name" required:"true"`
	Provider    string                      `json:"provider" xml:"provider" required:"true"`
	ClientId    string                      `json:"client_id" xml:"client_id" required:"true"`
	TokenUrl    string                      `json:"token_url" xml:"token_url" required:"true"`
	Scopes      []string                    `json:"scopes,omitempty" xml:"scopes,omitempty"`
	ExtraParams map[string]string           `json:"extra_params,omitempty"`
	Id          string                      `json:"id" xml:"id" required:"true"`
	UsedBy      *AuthConnectionDependencies `json:"used_by,omitempty" xml:"used_by,omitempty"`
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

func (r *RefreshTokenAuthResponse) GetName() (value string) {
	if r == nil {
		return
	}
	return r.Name
}

func (r *RefreshTokenAuthResponse) GetProvider() (value string) {
	if r == nil {
		return
	}
	return r.Provider
}

func (r *RefreshTokenAuthResponse) GetClientId() (value string) {
	if r == nil {
		return
	}
	return r.ClientId
}

func (r *RefreshTokenAuthResponse) GetTokenUrl() (value string) {
	if r == nil {
		return
	}
	return r.TokenUrl
}

func (r *RefreshTokenAuthResponse) GetScopes() (value []string) {
	if r == nil {
		return
	}
	return r.Scopes
}

func (r *RefreshTokenAuthResponse) GetExtraParams() (value map[string]string) {
	if r == nil {
		return
	}
	return r.ExtraParams
}

func (r *RefreshTokenAuthResponse) GetId() (value string) {
	if r == nil {
		return
	}
	return r.Id
}

func (r *RefreshTokenAuthResponse) GetUsedBy() (value *AuthConnectionDependencies) {
	if r == nil {
		return
	}
	return r.UsedBy
}

func (r *RefreshTokenAuthResponse) GetStatus() (value *AuthConnectionStatus) {
	if r == nil {
		return
	}
	return r.Status
}

func (r *RefreshTokenAuthResponse) GetStatusDetail() (value *string) {
	if r == nil {
		return
	}
	return r.StatusDetail
}

func (r *RefreshTokenAuthResponse) GetStatusUpdatedAt() (value *string) {
	if r == nil {
		return
	}
	return r.StatusUpdatedAt
}

func (r *RefreshTokenAuthResponse) SetName(name string) {
	r.Name = name
}

func (r *RefreshTokenAuthResponse) SetProvider(provider string) {
	r.Provider = provider
}

func (r *RefreshTokenAuthResponse) SetClientId(clientId string) {
	r.ClientId = clientId
}

func (r *RefreshTokenAuthResponse) SetTokenUrl(tokenUrl string) {
	r.TokenUrl = tokenUrl
}

func (r *RefreshTokenAuthResponse) SetScopes(scopes []string) {
	r.Scopes = scopes
}

func (r *RefreshTokenAuthResponse) SetExtraParams(extraParams map[string]string) {
	r.ExtraParams = extraParams
}

func (r *RefreshTokenAuthResponse) SetId(id string) {
	r.Id = id
}

func (r *RefreshTokenAuthResponse) SetUsedBy(usedBy *AuthConnectionDependencies) {
	r.UsedBy = usedBy
}

func (r *RefreshTokenAuthResponse) SetStatus(status *AuthConnectionStatus) {
	r.Status = status
}

func (r *RefreshTokenAuthResponse) SetStatusDetail(statusDetail *string) {
	r.StatusDetail = statusDetail
}

func (r *RefreshTokenAuthResponse) SetStatusUpdatedAt(statusUpdatedAt *string) {
	r.StatusUpdatedAt = statusUpdatedAt
}

func (r *RefreshTokenAuthResponse) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r RefreshTokenAuthResponse) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RefreshTokenAuthResponse to string"
	}
	return string(jsonData)
}

func (r *RefreshTokenAuthResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, r); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *r, "auth_type")
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *RefreshTokenAuthResponse) MarshalJSON() ([]byte, error) {
	type embed RefreshTokenAuthResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
