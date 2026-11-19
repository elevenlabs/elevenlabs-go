package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type EnvironmentVariableAuthConnectionValueRequest struct {
	AuthConnectionId string `json:"auth_connection_id" xml:"auth_connection_id" required:"true"`
	extraProperties  map[string]interface{}
}

func (e *EnvironmentVariableAuthConnectionValueRequest) GetAuthConnectionId() (value string) {
	if e == nil {
		return
	}
	return e.AuthConnectionId
}

func (e *EnvironmentVariableAuthConnectionValueRequest) SetAuthConnectionId(authConnectionId string) {
	e.AuthConnectionId = authConnectionId
}

func (e *EnvironmentVariableAuthConnectionValueRequest) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e EnvironmentVariableAuthConnectionValueRequest) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EnvironmentVariableAuthConnectionValueRequest to string"
	}
	return string(jsonData)
}

func (e *EnvironmentVariableAuthConnectionValueRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, e); err != nil {
		return err
	}
	type alias EnvironmentVariableAuthConnectionValueRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*e = EnvironmentVariableAuthConnectionValueRequest(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *EnvironmentVariableAuthConnectionValueRequest) MarshalJSON() ([]byte, error) {
	type embed EnvironmentVariableAuthConnectionValueRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
