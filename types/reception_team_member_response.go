package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ReceptionTeamMemberResponse struct {
	// The member's user ID.
	UserId string `json:"user_id" xml:"user_id" required:"true"`
	// The member's email address.
	Email string `json:"email" xml:"email" required:"true"`
	// The member's first name, if set.
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// Whether the member owns the workspace.
	IsOwner bool `json:"is_owner" xml:"is_owner" required:"true"`
	// Whether the member is a workspace admin.
	IsAdmin bool `json:"is_admin" xml:"is_admin" required:"true"`
	// URL of the member's profile picture, if set.
	PictureUrl *string `json:"picture_url,omitempty" xml:"picture_url,omitempty"`
	// Whether the member's access to the workspace is suspended.
	IsSuspended     *bool `json:"is_suspended,omitempty" xml:"is_suspended,omitempty"`
	extraProperties map[string]interface{}
}

func (r *ReceptionTeamMemberResponse) GetUserId() (value string) {
	if r == nil {
		return
	}
	return r.UserId
}

func (r *ReceptionTeamMemberResponse) GetEmail() (value string) {
	if r == nil {
		return
	}
	return r.Email
}

func (r *ReceptionTeamMemberResponse) GetFirstName() (value *string) {
	if r == nil {
		return
	}
	return r.FirstName
}

func (r *ReceptionTeamMemberResponse) GetIsOwner() (value bool) {
	if r == nil {
		return
	}
	return r.IsOwner
}

func (r *ReceptionTeamMemberResponse) GetIsAdmin() (value bool) {
	if r == nil {
		return
	}
	return r.IsAdmin
}

func (r *ReceptionTeamMemberResponse) GetPictureUrl() (value *string) {
	if r == nil {
		return
	}
	return r.PictureUrl
}

func (r *ReceptionTeamMemberResponse) GetIsSuspended() (value *bool) {
	if r == nil {
		return
	}
	return r.IsSuspended
}

func (r *ReceptionTeamMemberResponse) SetUserId(userId string) {
	r.UserId = userId
}

func (r *ReceptionTeamMemberResponse) SetEmail(email string) {
	r.Email = email
}

func (r *ReceptionTeamMemberResponse) SetFirstName(firstName *string) {
	r.FirstName = firstName
}

func (r *ReceptionTeamMemberResponse) SetIsOwner(isOwner bool) {
	r.IsOwner = isOwner
}

func (r *ReceptionTeamMemberResponse) SetIsAdmin(isAdmin bool) {
	r.IsAdmin = isAdmin
}

func (r *ReceptionTeamMemberResponse) SetPictureUrl(pictureUrl *string) {
	r.PictureUrl = pictureUrl
}

func (r *ReceptionTeamMemberResponse) SetIsSuspended(isSuspended *bool) {
	r.IsSuspended = isSuspended
}

func (r *ReceptionTeamMemberResponse) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReceptionTeamMemberResponse) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionTeamMemberResponse to string"
	}
	return string(jsonData)
}

func (r *ReceptionTeamMemberResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, r); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *ReceptionTeamMemberResponse) MarshalJSON() ([]byte, error) {
	type embed ReceptionTeamMemberResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
