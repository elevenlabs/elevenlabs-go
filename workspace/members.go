package workspace

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
)

type UpdateMemberRequest struct {
	// Email of the target user.
	Email string `json:"email" xml:"email" required:"true"`
	// Whether to lock or unlock the user account.
	IsLocked *bool `json:"is_locked,omitempty" xml:"is_locked,omitempty"`
	// The workspace role of the user. This is deprecated, use `workspace_seat_type` instead.
	WorkspaceRole *types.SeatType `json:"workspace_role,omitempty" xml:"workspace_role,omitempty"`
	// The workspace seat type
	WorkspaceSeatType *types.SeatType `json:"workspace_seat_type,omitempty" xml:"workspace_seat_type,omitempty"`
	// When unlocking a user whose original seat type has no seats available but a basic seat is available, downgrade the user to a basic seat instead of failing.
	AllowDowngradeToBasicSeat *bool `json:"allow_downgrade_to_basic_seat,omitempty" xml:"allow_downgrade_to_basic_seat,omitempty"`
}

func (u *UpdateMemberRequest) GetEmail() (value string) {
	if u == nil {
		return
	}
	return u.Email
}

func (u *UpdateMemberRequest) GetIsLocked() (value *bool) {
	if u == nil {
		return
	}
	return u.IsLocked
}

func (u *UpdateMemberRequest) GetWorkspaceRole() (value *types.SeatType) {
	if u == nil {
		return
	}
	return u.WorkspaceRole
}

func (u *UpdateMemberRequest) GetWorkspaceSeatType() (value *types.SeatType) {
	if u == nil {
		return
	}
	return u.WorkspaceSeatType
}

func (u *UpdateMemberRequest) GetAllowDowngradeToBasicSeat() (value *bool) {
	if u == nil {
		return
	}
	return u.AllowDowngradeToBasicSeat
}

func (u *UpdateMemberRequest) SetEmail(email string) {
	u.Email = email
}

func (u *UpdateMemberRequest) SetIsLocked(isLocked *bool) {
	u.IsLocked = isLocked
}

func (u *UpdateMemberRequest) SetWorkspaceRole(workspaceRole *types.SeatType) {
	u.WorkspaceRole = workspaceRole
}

func (u *UpdateMemberRequest) SetWorkspaceSeatType(workspaceSeatType *types.SeatType) {
	u.WorkspaceSeatType = workspaceSeatType
}

func (u *UpdateMemberRequest) SetAllowDowngradeToBasicSeat(allowDowngradeToBasicSeat *bool) {
	u.AllowDowngradeToBasicSeat = allowDowngradeToBasicSeat
}

func (u *UpdateMemberRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (u UpdateMemberRequest) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateMemberRequest to string"
	}
	return string(jsonData)
}

func (u *UpdateMemberRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, u)
}

func (u *UpdateMemberRequest) MarshalJSON() ([]byte, error) {
	type embed UpdateMemberRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}

type UpdateMemberUsageLimitRequestModel struct {
	// User ID of the target user.
	UserId string `json:"-" url:"-"`
	// Monthly credit usage limit for the member. Set to null to remove the limit.
	UsageLimit *int `json:"usage_limit" xml:"usage_limit" required:"true" min:"0"`
}

func (u *UpdateMemberUsageLimitRequestModel) GetUserId() (value string) {
	if u == nil {
		return
	}
	return u.UserId
}

func (u *UpdateMemberUsageLimitRequestModel) GetUsageLimit() (value *int) {
	if u == nil {
		return
	}
	return u.UsageLimit
}

func (u *UpdateMemberUsageLimitRequestModel) SetUserId(userId string) {
	u.UserId = userId
}

func (u *UpdateMemberUsageLimitRequestModel) SetUsageLimit(usageLimit *int) {
	u.UsageLimit = usageLimit
}

func (u *UpdateMemberUsageLimitRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (u UpdateMemberUsageLimitRequestModel) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateMemberUsageLimitRequestModel to string"
	}
	return string(jsonData)
}

func (u *UpdateMemberUsageLimitRequestModel) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, u)
}

func (u *UpdateMemberUsageLimitRequestModel) MarshalJSON() ([]byte, error) {
	type embed UpdateMemberUsageLimitRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
