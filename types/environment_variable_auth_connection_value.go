package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type EnvironmentVariableAuthConnectionValue struct {
	AuthConnectionId string `json:"auth_connection_id" xml:"auth_connection_id" required:"true"`
	extraProperties  map[string]interface{}
}

func (e *EnvironmentVariableAuthConnectionValue) GetAuthConnectionId() (value string) {
	if e == nil {
		return
	}
	return e.AuthConnectionId
}

func (e *EnvironmentVariableAuthConnectionValue) SetAuthConnectionId(authConnectionId string) {
	e.AuthConnectionId = authConnectionId
}

func (e *EnvironmentVariableAuthConnectionValue) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e EnvironmentVariableAuthConnectionValue) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EnvironmentVariableAuthConnectionValue to string"
	}
	return string(jsonData)
}

func (e *EnvironmentVariableAuthConnectionValue) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, e); err != nil {
		return err
	}
	type alias EnvironmentVariableAuthConnectionValue
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*e = EnvironmentVariableAuthConnectionValue(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *EnvironmentVariableAuthConnectionValue) MarshalJSON() ([]byte, error) {
	type embed EnvironmentVariableAuthConnectionValue
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
