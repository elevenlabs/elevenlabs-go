package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type EnvironmentVariableSecretValue struct {
	SecretId        string `json:"secret_id" xml:"secret_id" required:"true"`
	extraProperties map[string]interface{}
}

func (e *EnvironmentVariableSecretValue) GetSecretId() (value string) {
	if e == nil {
		return
	}
	return e.SecretId
}

func (e *EnvironmentVariableSecretValue) SetSecretId(secretId string) {
	e.SecretId = secretId
}

func (e *EnvironmentVariableSecretValue) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e EnvironmentVariableSecretValue) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EnvironmentVariableSecretValue to string"
	}
	return string(jsonData)
}

func (e *EnvironmentVariableSecretValue) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, e); err != nil {
		return err
	}
	type alias EnvironmentVariableSecretValue
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*e = EnvironmentVariableSecretValue(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *EnvironmentVariableSecretValue) MarshalJSON() ([]byte, error) {
	type embed EnvironmentVariableSecretValue
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
