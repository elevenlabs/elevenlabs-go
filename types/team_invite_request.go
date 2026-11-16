package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TeamInviteRequest struct {
	// The email address to invite.
	Email string `json:"email" xml:"email" required:"true" maxLength:"254" minLength:"1"`
	// The role the invitee joins with: 'workspace_admin' or 'workspace_member'.
	Role            *TeamInviteRequestRole `json:"role,omitempty" xml:"role,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TeamInviteRequest) GetEmail() (value string) {
	if t == nil {
		return
	}
	return t.Email
}

func (t *TeamInviteRequest) GetRole() (value *TeamInviteRequestRole) {
	if t == nil {
		return
	}
	return t.Role
}

func (t *TeamInviteRequest) SetEmail(email string) {
	t.Email = email
}

func (t *TeamInviteRequest) SetRole(role *TeamInviteRequestRole) {
	t.Role = role
}

func (t *TeamInviteRequest) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TeamInviteRequest) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TeamInviteRequest to string"
	}
	return string(jsonData)
}

func (t *TeamInviteRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TeamInviteRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TeamInviteRequest(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TeamInviteRequest) MarshalJSON() ([]byte, error) {
	type embed TeamInviteRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
