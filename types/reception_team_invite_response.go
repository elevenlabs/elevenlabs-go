package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ReceptionTeamInviteResponse struct {
	// The invited email address.
	Email string `json:"email" xml:"email" required:"true"`
	// Seat types for workspace members.
	SeatType SeatType `json:"seat_type" xml:"seat_type" required:"true"`
	// Email address of the member who sent the invite, if known.
	InvitingUserEmail *string `json:"inviting_user_email,omitempty" xml:"inviting_user_email,omitempty"`
	// When the invite was sent, as a Unix timestamp in seconds.
	CreatedAtUnix   *int `json:"created_at_unix,omitempty" xml:"created_at_unix,omitempty"`
	extraProperties map[string]interface{}
}

func (r *ReceptionTeamInviteResponse) GetEmail() (value string) {
	if r == nil {
		return
	}
	return r.Email
}

func (r *ReceptionTeamInviteResponse) GetSeatType() (value SeatType) {
	if r == nil {
		return
	}
	return r.SeatType
}

func (r *ReceptionTeamInviteResponse) GetInvitingUserEmail() (value *string) {
	if r == nil {
		return
	}
	return r.InvitingUserEmail
}

func (r *ReceptionTeamInviteResponse) GetCreatedAtUnix() (value *int) {
	if r == nil {
		return
	}
	return r.CreatedAtUnix
}

func (r *ReceptionTeamInviteResponse) SetEmail(email string) {
	r.Email = email
}

func (r *ReceptionTeamInviteResponse) SetSeatType(seatType SeatType) {
	r.SeatType = seatType
}

func (r *ReceptionTeamInviteResponse) SetInvitingUserEmail(invitingUserEmail *string) {
	r.InvitingUserEmail = invitingUserEmail
}

func (r *ReceptionTeamInviteResponse) SetCreatedAtUnix(createdAtUnix *int) {
	r.CreatedAtUnix = createdAtUnix
}

func (r *ReceptionTeamInviteResponse) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReceptionTeamInviteResponse) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionTeamInviteResponse to string"
	}
	return string(jsonData)
}

func (r *ReceptionTeamInviteResponse) UnmarshalJSON(data []byte) error {
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

func (r *ReceptionTeamInviteResponse) MarshalJSON() ([]byte, error) {
	type embed ReceptionTeamInviteResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
