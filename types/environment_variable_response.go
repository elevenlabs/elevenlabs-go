package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type EnvironmentVariableResponse struct {
	Label             string                             `json:"label" xml:"label" required:"true"`
	CreatedAtUnixSecs int                                `json:"created_at_unix_secs" xml:"created_at_unix_secs" required:"true"`
	UpdatedAtUnixSecs int                                `json:"updated_at_unix_secs" xml:"updated_at_unix_secs" required:"true"`
	CreatedByUserId   *string                            `json:"created_by_user_id,omitempty" xml:"created_by_user_id,omitempty"`
	Type              EnvironmentVariableResponseType    `json:"type" xml:"type" required:"true"`
	Id                string                             `json:"id" xml:"id" required:"true"`
	WorkspaceId       string                             `json:"workspace_id" xml:"workspace_id" required:"true"`
	Values            *EnvironmentVariableResponseValues `json:"values" xml:"values" required:"true"`
	extraProperties   map[string]interface{}
}

func (e *EnvironmentVariableResponse) GetLabel() (value string) {
	if e == nil {
		return
	}
	return e.Label
}

func (e *EnvironmentVariableResponse) GetCreatedAtUnixSecs() (value int) {
	if e == nil {
		return
	}
	return e.CreatedAtUnixSecs
}

func (e *EnvironmentVariableResponse) GetUpdatedAtUnixSecs() (value int) {
	if e == nil {
		return
	}
	return e.UpdatedAtUnixSecs
}

func (e *EnvironmentVariableResponse) GetCreatedByUserId() (value *string) {
	if e == nil {
		return
	}
	return e.CreatedByUserId
}

func (e *EnvironmentVariableResponse) GetType() (value EnvironmentVariableResponseType) {
	if e == nil {
		return
	}
	return e.Type
}

func (e *EnvironmentVariableResponse) GetId() (value string) {
	if e == nil {
		return
	}
	return e.Id
}

func (e *EnvironmentVariableResponse) GetWorkspaceId() (value string) {
	if e == nil {
		return
	}
	return e.WorkspaceId
}

func (e *EnvironmentVariableResponse) GetValues() (value *EnvironmentVariableResponseValues) {
	if e == nil {
		return
	}
	return e.Values
}

func (e *EnvironmentVariableResponse) SetLabel(label string) {
	e.Label = label
}

func (e *EnvironmentVariableResponse) SetCreatedAtUnixSecs(createdAtUnixSecs int) {
	e.CreatedAtUnixSecs = createdAtUnixSecs
}

func (e *EnvironmentVariableResponse) SetUpdatedAtUnixSecs(updatedAtUnixSecs int) {
	e.UpdatedAtUnixSecs = updatedAtUnixSecs
}

func (e *EnvironmentVariableResponse) SetCreatedByUserId(createdByUserId *string) {
	e.CreatedByUserId = createdByUserId
}

func (e *EnvironmentVariableResponse) SetType(type_ EnvironmentVariableResponseType) {
	e.Type = type_
}

func (e *EnvironmentVariableResponse) SetId(id string) {
	e.Id = id
}

func (e *EnvironmentVariableResponse) SetWorkspaceId(workspaceId string) {
	e.WorkspaceId = workspaceId
}

func (e *EnvironmentVariableResponse) SetValues(values *EnvironmentVariableResponseValues) {
	e.Values = values
}

func (e *EnvironmentVariableResponse) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e EnvironmentVariableResponse) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EnvironmentVariableResponse to string"
	}
	return string(jsonData)
}

func (e *EnvironmentVariableResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, e); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *EnvironmentVariableResponse) MarshalJSON() ([]byte, error) {
	type embed EnvironmentVariableResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
