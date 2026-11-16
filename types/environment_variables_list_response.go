package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type EnvironmentVariablesListResponse struct {
	EnvironmentVariables []*EnvironmentVariableResponse `json:"environment_variables" xml:"environment_variables" required:"true"`
	NextCursor           *string                        `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	HasMore              bool                           `json:"has_more" xml:"has_more" required:"true"`
	extraProperties      map[string]interface{}
}

func (e *EnvironmentVariablesListResponse) GetEnvironmentVariables() (value []*EnvironmentVariableResponse) {
	if e == nil {
		return
	}
	return e.EnvironmentVariables
}

func (e *EnvironmentVariablesListResponse) GetNextCursor() (value *string) {
	if e == nil {
		return
	}
	return e.NextCursor
}

func (e *EnvironmentVariablesListResponse) GetHasMore() (value bool) {
	if e == nil {
		return
	}
	return e.HasMore
}

func (e *EnvironmentVariablesListResponse) SetEnvironmentVariables(environmentVariables []*EnvironmentVariableResponse) {
	e.EnvironmentVariables = environmentVariables
}

func (e *EnvironmentVariablesListResponse) SetNextCursor(nextCursor *string) {
	e.NextCursor = nextCursor
}

func (e *EnvironmentVariablesListResponse) SetHasMore(hasMore bool) {
	e.HasMore = hasMore
}

func (e *EnvironmentVariablesListResponse) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e EnvironmentVariablesListResponse) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EnvironmentVariablesListResponse to string"
	}
	return string(jsonData)
}

func (e *EnvironmentVariablesListResponse) UnmarshalJSON(data []byte) error {
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

func (e *EnvironmentVariablesListResponse) MarshalJSON() ([]byte, error) {
	type embed EnvironmentVariablesListResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
