package workspace

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
)

type InviteUserRequest struct {
	// The email of the customer
	Email string `json:"email" xml:"email" required:"true"`
	// The workspace permission of the user. This is deprecated, use `seat_type` instead.
	WorkspacePermission *string `json:"workspace_permission,omitempty" xml:"workspace_permission,omitempty"`
	// The seat type of the user
	SeatType *types.SeatType `json:"seat_type,omitempty" xml:"seat_type,omitempty"`
	// The group ids of the user
	GroupIds []string `json:"group_ids,omitempty" xml:"group_ids,omitempty"`
	// Monthly credit usage limit for the invitee. Omit or set to null for no custom cap.
	UsageLimit *int `json:"usage_limit,omitempty" xml:"usage_limit,omitempty" min:"0"`
}

func (i *InviteUserRequest) GetEmail() (value string) {
	if i == nil {
		return
	}
	return i.Email
}

func (i *InviteUserRequest) GetWorkspacePermission() (value *string) {
	if i == nil {
		return
	}
	return i.WorkspacePermission
}

func (i *InviteUserRequest) GetSeatType() (value *types.SeatType) {
	if i == nil {
		return
	}
	return i.SeatType
}

func (i *InviteUserRequest) GetGroupIds() (value []string) {
	if i == nil {
		return
	}
	return i.GroupIds
}

func (i *InviteUserRequest) GetUsageLimit() (value *int) {
	if i == nil {
		return
	}
	return i.UsageLimit
}

func (i *InviteUserRequest) SetEmail(email string) {
	i.Email = email
}

func (i *InviteUserRequest) SetWorkspacePermission(workspacePermission *string) {
	i.WorkspacePermission = workspacePermission
}

func (i *InviteUserRequest) SetSeatType(seatType *types.SeatType) {
	i.SeatType = seatType
}

func (i *InviteUserRequest) SetGroupIds(groupIds []string) {
	i.GroupIds = groupIds
}

func (i *InviteUserRequest) SetUsageLimit(usageLimit *int) {
	i.UsageLimit = usageLimit
}

func (i *InviteUserRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (i InviteUserRequest) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InviteUserRequest to string"
	}
	return string(jsonData)
}

func (i *InviteUserRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, i)
}

func (i *InviteUserRequest) MarshalJSON() ([]byte, error) {
	type embed InviteUserRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*i),
	}
	return json.Marshal(marshaler)
}

type BodyInviteMultipleUsersV1WorkspaceInvitesAddBulkPost struct {
	// The email of the customer
	Emails []string `json:"emails" xml:"emails" required:"true"`
	// The seat type of the user
	SeatType *types.SeatType `json:"seat_type,omitempty" xml:"seat_type,omitempty"`
	// The group ids of the user
	GroupIds []string `json:"group_ids,omitempty" xml:"group_ids,omitempty"`
	// Monthly credit usage limit for the invitee. Omit or set to null for no custom cap.
	UsageLimit *int `json:"usage_limit,omitempty" xml:"usage_limit,omitempty" min:"0"`
}

func (b *BodyInviteMultipleUsersV1WorkspaceInvitesAddBulkPost) GetEmails() (value []string) {
	if b == nil {
		return
	}
	return b.Emails
}

func (b *BodyInviteMultipleUsersV1WorkspaceInvitesAddBulkPost) GetSeatType() (value *types.SeatType) {
	if b == nil {
		return
	}
	return b.SeatType
}

func (b *BodyInviteMultipleUsersV1WorkspaceInvitesAddBulkPost) GetGroupIds() (value []string) {
	if b == nil {
		return
	}
	return b.GroupIds
}

func (b *BodyInviteMultipleUsersV1WorkspaceInvitesAddBulkPost) GetUsageLimit() (value *int) {
	if b == nil {
		return
	}
	return b.UsageLimit
}

func (b *BodyInviteMultipleUsersV1WorkspaceInvitesAddBulkPost) SetEmails(emails []string) {
	b.Emails = emails
}

func (b *BodyInviteMultipleUsersV1WorkspaceInvitesAddBulkPost) SetSeatType(seatType *types.SeatType) {
	b.SeatType = seatType
}

func (b *BodyInviteMultipleUsersV1WorkspaceInvitesAddBulkPost) SetGroupIds(groupIds []string) {
	b.GroupIds = groupIds
}

func (b *BodyInviteMultipleUsersV1WorkspaceInvitesAddBulkPost) SetUsageLimit(usageLimit *int) {
	b.UsageLimit = usageLimit
}

func (b *BodyInviteMultipleUsersV1WorkspaceInvitesAddBulkPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyInviteMultipleUsersV1WorkspaceInvitesAddBulkPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyInviteMultipleUsersV1WorkspaceInvitesAddBulkPost to string"
	}
	return string(jsonData)
}

func (b *BodyInviteMultipleUsersV1WorkspaceInvitesAddBulkPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyInviteMultipleUsersV1WorkspaceInvitesAddBulkPost) MarshalJSON() ([]byte, error) {
	type embed BodyInviteMultipleUsersV1WorkspaceInvitesAddBulkPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

type BodyDeleteExistingInvitationV1WorkspaceInvitesDelete struct {
	// The email of the customer
	Email string `json:"email" xml:"email" required:"true"`
}

func (b *BodyDeleteExistingInvitationV1WorkspaceInvitesDelete) GetEmail() (value string) {
	if b == nil {
		return
	}
	return b.Email
}

func (b *BodyDeleteExistingInvitationV1WorkspaceInvitesDelete) SetEmail(email string) {
	b.Email = email
}

func (b *BodyDeleteExistingInvitationV1WorkspaceInvitesDelete) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyDeleteExistingInvitationV1WorkspaceInvitesDelete) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyDeleteExistingInvitationV1WorkspaceInvitesDelete to string"
	}
	return string(jsonData)
}

func (b *BodyDeleteExistingInvitationV1WorkspaceInvitesDelete) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, b); err != nil {
		return err
	}
	type alias BodyDeleteExistingInvitationV1WorkspaceInvitesDelete
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = BodyDeleteExistingInvitationV1WorkspaceInvitesDelete(tmp)
	return nil
}

func (b *BodyDeleteExistingInvitationV1WorkspaceInvitesDelete) MarshalJSON() ([]byte, error) {
	type embed BodyDeleteExistingInvitationV1WorkspaceInvitesDelete
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
