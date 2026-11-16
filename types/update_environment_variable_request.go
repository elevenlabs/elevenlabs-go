package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type UpdateEnvironmentVariableRequest struct {
	EnvVarId string `json:"-" url:"-"`
	// Values to replace. Set to null to remove the value for an environment (except 'production').
	Values map[string]*UpdateEnvironmentVariableRequestValuesValue `json:"values" required:"true"`
}

func (u *UpdateEnvironmentVariableRequest) GetEnvVarId() (value string) {
	if u == nil {
		return
	}
	return u.EnvVarId
}

func (u *UpdateEnvironmentVariableRequest) GetValues() (value map[string]*UpdateEnvironmentVariableRequestValuesValue) {
	if u == nil {
		return
	}
	return u.Values
}

func (u *UpdateEnvironmentVariableRequest) SetEnvVarId(envVarId string) {
	u.EnvVarId = envVarId
}

func (u *UpdateEnvironmentVariableRequest) SetValues(values map[string]*UpdateEnvironmentVariableRequestValuesValue) {
	u.Values = values
}

func (u *UpdateEnvironmentVariableRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (u UpdateEnvironmentVariableRequest) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateEnvironmentVariableRequest to string"
	}
	return string(jsonData)
}

func (u *UpdateEnvironmentVariableRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, u); err != nil {
		return err
	}
	type alias UpdateEnvironmentVariableRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*u = UpdateEnvironmentVariableRequest(tmp)
	return nil
}

func (u *UpdateEnvironmentVariableRequest) MarshalJSON() ([]byte, error) {
	type embed UpdateEnvironmentVariableRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
