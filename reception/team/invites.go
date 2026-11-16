package team

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	big "math/big"
)

type CreateTeamInviteRequest struct {
	// The customer workspace a reseller's API key acts on.
	XiWorkspaceId string `json:"-" url:"-" explode:"false" serializationStyle:"simple" headerParam:"xi-workspace-id" required:"true"`
	// The email address to invite.
	Email string `json:"email" xml:"email" required:"true" maxLength:"254" minLength:"1"`
	// The role the invitee joins with: 'workspace_admin' or 'workspace_member'.
	Role *CreateTeamInviteRequestRole `json:"role,omitempty" xml:"role,omitempty"`
}

func (c *CreateTeamInviteRequest) GetXiWorkspaceId() (value string) {
	if c == nil {
		return
	}
	return c.XiWorkspaceId
}

func (c *CreateTeamInviteRequest) GetEmail() (value string) {
	if c == nil {
		return
	}
	return c.Email
}

func (c *CreateTeamInviteRequest) GetRole() (value *CreateTeamInviteRequestRole) {
	if c == nil {
		return
	}
	return c.Role
}

func (c *CreateTeamInviteRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	c.XiWorkspaceId = xiWorkspaceId
}

func (c *CreateTeamInviteRequest) SetEmail(email string) {
	c.Email = email
}

func (c *CreateTeamInviteRequest) SetRole(role *CreateTeamInviteRequestRole) {
	c.Role = role
}

func (c *CreateTeamInviteRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (c CreateTeamInviteRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateTeamInviteRequest to string"
	}
	return string(jsonData)
}

func (c *CreateTeamInviteRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreateTeamInviteRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreateTeamInviteRequest(tmp)
	return nil
}

func (c *CreateTeamInviteRequest) MarshalJSON() ([]byte, error) {
	type embed CreateTeamInviteRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}

// The role the invitee joins with: 'workspace_admin' or 'workspace_member'.
type CreateTeamInviteRequestRole string

const (
	CreateTeamInviteRequestRoleWorkspaceAdmin  CreateTeamInviteRequestRole = "workspace_admin"
	CreateTeamInviteRequestRoleWorkspaceMember CreateTeamInviteRequestRole = "workspace_member"
)

func NewCreateTeamInviteRequestRoleFromString(s string) (CreateTeamInviteRequestRole, error) {
	switch s {
	case "workspace_admin":
		return CreateTeamInviteRequestRoleWorkspaceAdmin, nil
	case "workspace_member":
		return CreateTeamInviteRequestRoleWorkspaceMember, nil
	}
	var t CreateTeamInviteRequestRole
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CreateTeamInviteRequestRole) Ptr() *CreateTeamInviteRequestRole {
	return &e
}

type BodySmbCancelTeamInvite struct {
	// The customer workspace a reseller's API key acts on.
	XiWorkspaceId string `json:"-" url:"-" explode:"false" serializationStyle:"simple" headerParam:"xi-workspace-id" required:"true"`
	// Email address the pending invite was sent to.
	Email string `json:"email" xml:"email" required:"true"`
}

func (b *BodySmbCancelTeamInvite) GetXiWorkspaceId() (value string) {
	if b == nil {
		return
	}
	return b.XiWorkspaceId
}

func (b *BodySmbCancelTeamInvite) GetEmail() (value string) {
	if b == nil {
		return
	}
	return b.Email
}

func (b *BodySmbCancelTeamInvite) SetXiWorkspaceId(xiWorkspaceId string) {
	b.XiWorkspaceId = xiWorkspaceId
}

func (b *BodySmbCancelTeamInvite) SetEmail(email string) {
	b.Email = email
}

func (b *BodySmbCancelTeamInvite) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodySmbCancelTeamInvite) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodySmbCancelTeamInvite to string"
	}
	return string(jsonData)
}

func (b *BodySmbCancelTeamInvite) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, b); err != nil {
		return err
	}
	type alias BodySmbCancelTeamInvite
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = BodySmbCancelTeamInvite(tmp)
	return nil
}

func (b *BodySmbCancelTeamInvite) MarshalJSON() ([]byte, error) {
	type embed BodySmbCancelTeamInvite
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

var (
	listInvitesRequestFieldXiWorkspaceId = big.NewInt(1 << 0)
)

type ListInvitesRequest struct {
	XiWorkspaceId string `json:"-" url:"xi-workspace-id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListInvitesRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetXiWorkspaceId sets the XiWorkspaceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListInvitesRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	l.XiWorkspaceId = xiWorkspaceId
	l.require(listInvitesRequestFieldXiWorkspaceId)
}
