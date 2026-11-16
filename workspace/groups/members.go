package groups

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BodyDeleteMemberFromUserGroupV1WorkspaceGroupsGroupIdMembersRemovePost struct {
	// The ID of the target group.
	GroupId string `json:"-" url:"-"`
	// The email of the target workspace member.
	Email string `json:"email" xml:"email" required:"true"`
}

func (b *BodyDeleteMemberFromUserGroupV1WorkspaceGroupsGroupIdMembersRemovePost) GetGroupId() (value string) {
	if b == nil {
		return
	}
	return b.GroupId
}

func (b *BodyDeleteMemberFromUserGroupV1WorkspaceGroupsGroupIdMembersRemovePost) GetEmail() (value string) {
	if b == nil {
		return
	}
	return b.Email
}

func (b *BodyDeleteMemberFromUserGroupV1WorkspaceGroupsGroupIdMembersRemovePost) SetGroupId(groupId string) {
	b.GroupId = groupId
}

func (b *BodyDeleteMemberFromUserGroupV1WorkspaceGroupsGroupIdMembersRemovePost) SetEmail(email string) {
	b.Email = email
}

func (b *BodyDeleteMemberFromUserGroupV1WorkspaceGroupsGroupIdMembersRemovePost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyDeleteMemberFromUserGroupV1WorkspaceGroupsGroupIdMembersRemovePost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyDeleteMemberFromUserGroupV1WorkspaceGroupsGroupIdMembersRemovePost to string"
	}
	return string(jsonData)
}

func (b *BodyDeleteMemberFromUserGroupV1WorkspaceGroupsGroupIdMembersRemovePost) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, b); err != nil {
		return err
	}
	type alias BodyDeleteMemberFromUserGroupV1WorkspaceGroupsGroupIdMembersRemovePost
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = BodyDeleteMemberFromUserGroupV1WorkspaceGroupsGroupIdMembersRemovePost(tmp)
	return nil
}

func (b *BodyDeleteMemberFromUserGroupV1WorkspaceGroupsGroupIdMembersRemovePost) MarshalJSON() ([]byte, error) {
	type embed BodyDeleteMemberFromUserGroupV1WorkspaceGroupsGroupIdMembersRemovePost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

type AddMemberToGroupRequest struct {
	// The ID of the target group.
	GroupId string `json:"-" url:"-"`
	// The email of the target workspace member.
	Email string `json:"email" xml:"email" required:"true"`
}

func (a *AddMemberToGroupRequest) GetGroupId() (value string) {
	if a == nil {
		return
	}
	return a.GroupId
}

func (a *AddMemberToGroupRequest) GetEmail() (value string) {
	if a == nil {
		return
	}
	return a.Email
}

func (a *AddMemberToGroupRequest) SetGroupId(groupId string) {
	a.GroupId = groupId
}

func (a *AddMemberToGroupRequest) SetEmail(email string) {
	a.Email = email
}

func (a *AddMemberToGroupRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (a AddMemberToGroupRequest) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddMemberToGroupRequest to string"
	}
	return string(jsonData)
}

func (a *AddMemberToGroupRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AddMemberToGroupRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AddMemberToGroupRequest(tmp)
	return nil
}

func (a *AddMemberToGroupRequest) MarshalJSON() ([]byte, error) {
	type embed AddMemberToGroupRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
