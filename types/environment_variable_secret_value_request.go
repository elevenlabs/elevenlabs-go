package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type EnvironmentVariableSecretValueRequest struct {
	SecretId        string `json:"secret_id" xml:"secret_id" required:"true"`
	extraProperties map[string]interface{}
}

func (e *EnvironmentVariableSecretValueRequest) GetSecretId() (value string) {
	if e == nil {
		return
	}
	return e.SecretId
}

func (e *EnvironmentVariableSecretValueRequest) SetSecretId(secretId string) {
	e.SecretId = secretId
}

func (e *EnvironmentVariableSecretValueRequest) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e EnvironmentVariableSecretValueRequest) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EnvironmentVariableSecretValueRequest to string"
	}
	return string(jsonData)
}

func (e *EnvironmentVariableSecretValueRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, e); err != nil {
		return err
	}
	type alias EnvironmentVariableSecretValueRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*e = EnvironmentVariableSecretValueRequest(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *EnvironmentVariableSecretValueRequest) MarshalJSON() ([]byte, error) {
	type embed EnvironmentVariableSecretValueRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
