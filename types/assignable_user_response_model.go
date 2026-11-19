package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AssignableUserResponseModel struct {
	UserId           string  `json:"user_id" xml:"user_id" required:"true"`
	Email            string  `json:"email" xml:"email" required:"true"`
	FirstName        *string `json:"first_name" xml:"first_name" required:"true"`
	IsServiceAccount bool    `json:"is_service_account" xml:"is_service_account" required:"true"`
	// Whether this workspace member currently has at least viewer access to the agent. Members without access are still returned so they can be surfaced (e.g. grayed out) and granted access before being assigned.
	HasAccess       bool `json:"has_access" xml:"has_access" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AssignableUserResponseModel) GetUserId() (value string) {
	if a == nil {
		return
	}
	return a.UserId
}

func (a *AssignableUserResponseModel) GetEmail() (value string) {
	if a == nil {
		return
	}
	return a.Email
}

func (a *AssignableUserResponseModel) GetFirstName() (value *string) {
	if a == nil {
		return
	}
	return a.FirstName
}

func (a *AssignableUserResponseModel) GetIsServiceAccount() (value bool) {
	if a == nil {
		return
	}
	return a.IsServiceAccount
}

func (a *AssignableUserResponseModel) GetHasAccess() (value bool) {
	if a == nil {
		return
	}
	return a.HasAccess
}

func (a *AssignableUserResponseModel) SetUserId(userId string) {
	a.UserId = userId
}

func (a *AssignableUserResponseModel) SetEmail(email string) {
	a.Email = email
}

func (a *AssignableUserResponseModel) SetFirstName(firstName *string) {
	a.FirstName = firstName
}

func (a *AssignableUserResponseModel) SetIsServiceAccount(isServiceAccount bool) {
	a.IsServiceAccount = isServiceAccount
}

func (a *AssignableUserResponseModel) SetHasAccess(hasAccess bool) {
	a.HasAccess = hasAccess
}

func (a *AssignableUserResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AssignableUserResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AssignableUserResponseModel to string"
	}
	return string(jsonData)
}

func (a *AssignableUserResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AssignableUserResponseModel) MarshalJSON() ([]byte, error) {
	type embed AssignableUserResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
