package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkspaceGroupByNameResponseModel struct {
	// The name of the workspace group.
	Name string `json:"name" xml:"name" required:"true"`
	// The ID of the workspace group.
	Id string `json:"id" xml:"id" required:"true"`
	// The emails of the members of the workspace group.
	MembersEmails   []string `json:"members_emails" xml:"members_emails" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WorkspaceGroupByNameResponseModel) GetName() (value string) {
	if w == nil {
		return
	}
	return w.Name
}

func (w *WorkspaceGroupByNameResponseModel) GetId() (value string) {
	if w == nil {
		return
	}
	return w.Id
}

func (w *WorkspaceGroupByNameResponseModel) GetMembersEmails() (value []string) {
	if w == nil {
		return
	}
	return w.MembersEmails
}

func (w *WorkspaceGroupByNameResponseModel) SetName(name string) {
	w.Name = name
}

func (w *WorkspaceGroupByNameResponseModel) SetId(id string) {
	w.Id = id
}

func (w *WorkspaceGroupByNameResponseModel) SetMembersEmails(membersEmails []string) {
	w.MembersEmails = membersEmails
}

func (w *WorkspaceGroupByNameResponseModel) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkspaceGroupByNameResponseModel) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkspaceGroupByNameResponseModel to string"
	}
	return string(jsonData)
}

func (w *WorkspaceGroupByNameResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WorkspaceGroupByNameResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WorkspaceGroupByNameResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkspaceGroupByNameResponseModel) MarshalJSON() ([]byte, error) {
	type embed WorkspaceGroupByNameResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
