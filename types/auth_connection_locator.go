package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Used to reference an auth connection from the workspace's auth connection store.
type AuthConnectionLocator struct {
	AuthConnectionId string `json:"auth_connection_id" xml:"auth_connection_id" required:"true"`
	extraProperties  map[string]interface{}
}

func (a *AuthConnectionLocator) GetAuthConnectionId() (value string) {
	if a == nil {
		return
	}
	return a.AuthConnectionId
}

func (a *AuthConnectionLocator) SetAuthConnectionId(authConnectionId string) {
	a.AuthConnectionId = authConnectionId
}

func (a *AuthConnectionLocator) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AuthConnectionLocator) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AuthConnectionLocator to string"
	}
	return string(jsonData)
}

func (a *AuthConnectionLocator) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AuthConnectionLocator
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AuthConnectionLocator(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AuthConnectionLocator) MarshalJSON() ([]byte, error) {
	type embed AuthConnectionLocator
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
