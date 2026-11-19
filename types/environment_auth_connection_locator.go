package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// References an environment variable of type 'auth_connection' by label.
// At runtime, resolves to the auth connection for the current environment,
// falling back to the default environment.
type EnvironmentAuthConnectionLocator struct {
	EnvVarLabel     string `json:"env_var_label" xml:"env_var_label" required:"true"`
	extraProperties map[string]interface{}
}

func (e *EnvironmentAuthConnectionLocator) GetEnvVarLabel() (value string) {
	if e == nil {
		return
	}
	return e.EnvVarLabel
}

func (e *EnvironmentAuthConnectionLocator) SetEnvVarLabel(envVarLabel string) {
	e.EnvVarLabel = envVarLabel
}

func (e *EnvironmentAuthConnectionLocator) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e EnvironmentAuthConnectionLocator) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EnvironmentAuthConnectionLocator to string"
	}
	return string(jsonData)
}

func (e *EnvironmentAuthConnectionLocator) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, e); err != nil {
		return err
	}
	type alias EnvironmentAuthConnectionLocator
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*e = EnvironmentAuthConnectionLocator(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *EnvironmentAuthConnectionLocator) MarshalJSON() ([]byte, error) {
	type embed EnvironmentAuthConnectionLocator
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
