package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type UpdateWorkspaceMemberResponseModel struct {
	// The status of the workspace member update request. If the request was successful, the status will be 'ok'. Otherwise an error message with status 500 will be returned.
	Status          string `json:"status" xml:"status" required:"true"`
	extraProperties map[string]interface{}
}

func (u *UpdateWorkspaceMemberResponseModel) GetStatus() (value string) {
	if u == nil {
		return
	}
	return u.Status
}

func (u *UpdateWorkspaceMemberResponseModel) SetStatus(status string) {
	u.Status = status
}

func (u *UpdateWorkspaceMemberResponseModel) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UpdateWorkspaceMemberResponseModel) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateWorkspaceMemberResponseModel to string"
	}
	return string(jsonData)
}

func (u *UpdateWorkspaceMemberResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, u); err != nil {
		return err
	}
	type alias UpdateWorkspaceMemberResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*u = UpdateWorkspaceMemberResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UpdateWorkspaceMemberResponseModel) MarshalJSON() ([]byte, error) {
	type embed UpdateWorkspaceMemberResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
