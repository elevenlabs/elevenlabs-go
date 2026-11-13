package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Response model for Custom Header Auth auth connections
type CustomHeaderAuthResponse struct {
	Name     string `json:"name" xml:"name" required:"true"`
	Provider string `json:"provider" xml:"provider" required:"true"`
	// The name of the header to use for authentication (e.g., 'x-api-key')
	HeaderName string                      `json:"header_name" xml:"header_name" required:"true"`
	Id         string                      `json:"id" xml:"id" required:"true"`
	UsedBy     *AuthConnectionDependencies `json:"used_by,omitempty" xml:"used_by,omitempty"`
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

func (c *CustomHeaderAuthResponse) GetName() (value string) {
	if c == nil {
		return
	}
	return c.Name
}

func (c *CustomHeaderAuthResponse) GetProvider() (value string) {
	if c == nil {
		return
	}
	return c.Provider
}

func (c *CustomHeaderAuthResponse) GetHeaderName() (value string) {
	if c == nil {
		return
	}
	return c.HeaderName
}

func (c *CustomHeaderAuthResponse) GetId() (value string) {
	if c == nil {
		return
	}
	return c.Id
}

func (c *CustomHeaderAuthResponse) GetUsedBy() (value *AuthConnectionDependencies) {
	if c == nil {
		return
	}
	return c.UsedBy
}

func (c *CustomHeaderAuthResponse) GetStatus() (value *AuthConnectionStatus) {
	if c == nil {
		return
	}
	return c.Status
}

func (c *CustomHeaderAuthResponse) GetStatusDetail() (value *string) {
	if c == nil {
		return
	}
	return c.StatusDetail
}

func (c *CustomHeaderAuthResponse) GetStatusUpdatedAt() (value *string) {
	if c == nil {
		return
	}
	return c.StatusUpdatedAt
}

func (c *CustomHeaderAuthResponse) SetName(name string) {
	c.Name = name
}

func (c *CustomHeaderAuthResponse) SetProvider(provider string) {
	c.Provider = provider
}

func (c *CustomHeaderAuthResponse) SetHeaderName(headerName string) {
	c.HeaderName = headerName
}

func (c *CustomHeaderAuthResponse) SetId(id string) {
	c.Id = id
}

func (c *CustomHeaderAuthResponse) SetUsedBy(usedBy *AuthConnectionDependencies) {
	c.UsedBy = usedBy
}

func (c *CustomHeaderAuthResponse) SetStatus(status *AuthConnectionStatus) {
	c.Status = status
}

func (c *CustomHeaderAuthResponse) SetStatusDetail(statusDetail *string) {
	c.StatusDetail = statusDetail
}

func (c *CustomHeaderAuthResponse) SetStatusUpdatedAt(statusUpdatedAt *string) {
	c.StatusUpdatedAt = statusUpdatedAt
}

func (c *CustomHeaderAuthResponse) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CustomHeaderAuthResponse) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CustomHeaderAuthResponse to string"
	}
	return string(jsonData)
}

func (c *CustomHeaderAuthResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "auth_type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CustomHeaderAuthResponse) MarshalJSON() ([]byte, error) {
	type embed CustomHeaderAuthResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
