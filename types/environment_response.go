package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type EnvironmentResponse struct {
	// Name of the environment, e.g. 'staging'.
	Name string `json:"name" xml:"name" required:"true"`
	// When the environment was created or one of its values last changed. Unset when neither is known.
	UpdatedAtUnixSecs *int `json:"updated_at_unix_secs,omitempty" xml:"updated_at_unix_secs,omitempty"`
	// Unset for 'production' and when the creator is not known.
	CreatedByUserId *string `json:"created_by_user_id,omitempty" xml:"created_by_user_id,omitempty"`
	extraProperties map[string]interface{}
}

func (e *EnvironmentResponse) GetName() (value string) {
	if e == nil {
		return
	}
	return e.Name
}

func (e *EnvironmentResponse) GetUpdatedAtUnixSecs() (value *int) {
	if e == nil {
		return
	}
	return e.UpdatedAtUnixSecs
}

func (e *EnvironmentResponse) GetCreatedByUserId() (value *string) {
	if e == nil {
		return
	}
	return e.CreatedByUserId
}

func (e *EnvironmentResponse) SetName(name string) {
	e.Name = name
}

func (e *EnvironmentResponse) SetUpdatedAtUnixSecs(updatedAtUnixSecs *int) {
	e.UpdatedAtUnixSecs = updatedAtUnixSecs
}

func (e *EnvironmentResponse) SetCreatedByUserId(createdByUserId *string) {
	e.CreatedByUserId = createdByUserId
}

func (e *EnvironmentResponse) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e EnvironmentResponse) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EnvironmentResponse to string"
	}
	return string(jsonData)
}

func (e *EnvironmentResponse) UnmarshalJSON(data []byte) error {
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

func (e *EnvironmentResponse) MarshalJSON() ([]byte, error) {
	type embed EnvironmentResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
