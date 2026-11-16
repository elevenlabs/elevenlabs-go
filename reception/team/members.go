package team

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	big "math/big"
)

type BodySmbUpdateTeamMemberRole struct {
	// The customer workspace a reseller's API key acts on.
	XiWorkspaceId string `json:"-" url:"-" explode:"false" serializationStyle:"simple" headerParam:"xi-workspace-id" required:"true"`
	// User ID of the member to update
	UserId string `json:"-" url:"-"`
	// The member's new role.
	Role BodySmbUpdateTeamMemberRoleRole `json:"role" xml:"role" required:"true"`
}

func (b *BodySmbUpdateTeamMemberRole) GetXiWorkspaceId() (value string) {
	if b == nil {
		return
	}
	return b.XiWorkspaceId
}

func (b *BodySmbUpdateTeamMemberRole) GetUserId() (value string) {
	if b == nil {
		return
	}
	return b.UserId
}

func (b *BodySmbUpdateTeamMemberRole) GetRole() (value BodySmbUpdateTeamMemberRoleRole) {
	if b == nil {
		return
	}
	return b.Role
}

func (b *BodySmbUpdateTeamMemberRole) SetXiWorkspaceId(xiWorkspaceId string) {
	b.XiWorkspaceId = xiWorkspaceId
}

func (b *BodySmbUpdateTeamMemberRole) SetUserId(userId string) {
	b.UserId = userId
}

func (b *BodySmbUpdateTeamMemberRole) SetRole(role BodySmbUpdateTeamMemberRoleRole) {
	b.Role = role
}

func (b *BodySmbUpdateTeamMemberRole) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodySmbUpdateTeamMemberRole) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodySmbUpdateTeamMemberRole to string"
	}
	return string(jsonData)
}

func (b *BodySmbUpdateTeamMemberRole) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, b); err != nil {
		return err
	}
	type alias BodySmbUpdateTeamMemberRole
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = BodySmbUpdateTeamMemberRole(tmp)
	return nil
}

func (b *BodySmbUpdateTeamMemberRole) MarshalJSON() ([]byte, error) {
	type embed BodySmbUpdateTeamMemberRole
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

// The member's new role.
type BodySmbUpdateTeamMemberRoleRole string

const (
	BodySmbUpdateTeamMemberRoleRoleWorkspaceAdmin  BodySmbUpdateTeamMemberRoleRole = "workspace_admin"
	BodySmbUpdateTeamMemberRoleRoleWorkspaceMember BodySmbUpdateTeamMemberRoleRole = "workspace_member"
)

func NewBodySmbUpdateTeamMemberRoleRoleFromString(s string) (BodySmbUpdateTeamMemberRoleRole, error) {
	switch s {
	case "workspace_admin":
		return BodySmbUpdateTeamMemberRoleRoleWorkspaceAdmin, nil
	case "workspace_member":
		return BodySmbUpdateTeamMemberRoleRoleWorkspaceMember, nil
	}
	var t BodySmbUpdateTeamMemberRoleRole
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BodySmbUpdateTeamMemberRoleRole) Ptr() *BodySmbUpdateTeamMemberRoleRole {
	return &e
}

var (
	listMembersRequestFieldXiWorkspaceId = big.NewInt(1 << 0)
)

type ListMembersRequest struct {
	XiWorkspaceId string `json:"-" url:"xi-workspace-id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListMembersRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetXiWorkspaceId sets the XiWorkspaceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListMembersRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	l.XiWorkspaceId = xiWorkspaceId
	l.require(listMembersRequestFieldXiWorkspaceId)
}

var (
	removeMembersRequestFieldXiWorkspaceId = big.NewInt(1 << 0)
)

type RemoveMembersRequest struct {
	XiWorkspaceId string `json:"-" url:"xi-workspace-id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (r *RemoveMembersRequest) require(field *big.Int) {
	if r.explicitFields == nil {
		r.explicitFields = big.NewInt(0)
	}
	r.explicitFields.Or(r.explicitFields, field)
}

// SetXiWorkspaceId sets the XiWorkspaceId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *RemoveMembersRequest) SetXiWorkspaceId(xiWorkspaceId string) {
	r.XiWorkspaceId = xiWorkspaceId
	r.require(removeMembersRequestFieldXiWorkspaceId)
}
