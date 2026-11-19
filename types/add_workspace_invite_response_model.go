package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AddWorkspaceInviteResponseModel struct {
	// The status of the workspace invite request. If the request was successful, the status will be 'ok'. Otherwise an error message with status 500 will be returned.
	Status          string `json:"status" xml:"status" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AddWorkspaceInviteResponseModel) GetStatus() (value string) {
	if a == nil {
		return
	}
	return a.Status
}

func (a *AddWorkspaceInviteResponseModel) SetStatus(status string) {
	a.Status = status
}

func (a *AddWorkspaceInviteResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AddWorkspaceInviteResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddWorkspaceInviteResponseModel to string"
	}
	return string(jsonData)
}

func (a *AddWorkspaceInviteResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AddWorkspaceInviteResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AddWorkspaceInviteResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AddWorkspaceInviteResponseModel) MarshalJSON() ([]byte, error) {
	type embed AddWorkspaceInviteResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
