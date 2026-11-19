package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Public workspace member fields exposed via GET /v1/workspace/members.
type WorkspaceMemberResponseModel struct {
	// The user ID of the workspace member.
	UserId string `json:"user_id" xml:"user_id" required:"true"`
	// The email address of the workspace member.
	Email string `json:"email" xml:"email" required:"true"`
	// The first name of the workspace member, if available.
	FirstName *string `json:"first_name" xml:"first_name" required:"true"`
	// The seat type (role) of the workspace member.
	SeatType *SeatType `json:"seat_type" xml:"seat_type" required:"true"`
	// Whether the member is the workspace owner.
	IsOwner bool `json:"is_owner" xml:"is_owner" required:"true"`
	// Whether the member's account is locked in this workspace.
	IsLocked        bool `json:"is_locked" xml:"is_locked" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WorkspaceMemberResponseModel) GetUserId() (value string) {
	if w == nil {
		return
	}
	return w.UserId
}

func (w *WorkspaceMemberResponseModel) GetEmail() (value string) {
	if w == nil {
		return
	}
	return w.Email
}

func (w *WorkspaceMemberResponseModel) GetFirstName() (value *string) {
	if w == nil {
		return
	}
	return w.FirstName
}

func (w *WorkspaceMemberResponseModel) GetSeatType() (value *SeatType) {
	if w == nil {
		return
	}
	return w.SeatType
}

func (w *WorkspaceMemberResponseModel) GetIsOwner() (value bool) {
	if w == nil {
		return
	}
	return w.IsOwner
}

func (w *WorkspaceMemberResponseModel) GetIsLocked() (value bool) {
	if w == nil {
		return
	}
	return w.IsLocked
}

func (w *WorkspaceMemberResponseModel) SetUserId(userId string) {
	w.UserId = userId
}

func (w *WorkspaceMemberResponseModel) SetEmail(email string) {
	w.Email = email
}

func (w *WorkspaceMemberResponseModel) SetFirstName(firstName *string) {
	w.FirstName = firstName
}

func (w *WorkspaceMemberResponseModel) SetSeatType(seatType *SeatType) {
	w.SeatType = seatType
}

func (w *WorkspaceMemberResponseModel) SetIsOwner(isOwner bool) {
	w.IsOwner = isOwner
}

func (w *WorkspaceMemberResponseModel) SetIsLocked(isLocked bool) {
	w.IsLocked = isLocked
}

func (w *WorkspaceMemberResponseModel) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkspaceMemberResponseModel) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkspaceMemberResponseModel to string"
	}
	return string(jsonData)
}

func (w *WorkspaceMemberResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, w); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkspaceMemberResponseModel) MarshalJSON() ([]byte, error) {
	type embed WorkspaceMemberResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
