package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type EnvironmentsListResponse struct {
	// All workspace environments, 'production' first. Variables without a value for an environment fall back to their production value.
	Environments    []*EnvironmentResponse `json:"environments" xml:"environments" required:"true"`
	extraProperties map[string]interface{}
}

func (e *EnvironmentsListResponse) GetEnvironments() (value []*EnvironmentResponse) {
	if e == nil {
		return
	}
	return e.Environments
}

func (e *EnvironmentsListResponse) SetEnvironments(environments []*EnvironmentResponse) {
	e.Environments = environments
}

func (e *EnvironmentsListResponse) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e EnvironmentsListResponse) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EnvironmentsListResponse to string"
	}
	return string(jsonData)
}

func (e *EnvironmentsListResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler EnvironmentsListResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EnvironmentsListResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *EnvironmentsListResponse) MarshalJSON() ([]byte, error) {
	type embed EnvironmentsListResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
