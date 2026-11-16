package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CreateEnvironmentRequest struct {
	// Unique name of the environment, e.g. 'staging'.
	Name string `json:"name" xml:"name" required:"true"`
}

func (c *CreateEnvironmentRequest) GetName() (value string) {
	if c == nil {
		return
	}
	return c.Name
}

func (c *CreateEnvironmentRequest) SetName(name string) {
	c.Name = name
}

func (c *CreateEnvironmentRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (c CreateEnvironmentRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateEnvironmentRequest to string"
	}
	return string(jsonData)
}

func (c *CreateEnvironmentRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreateEnvironmentRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreateEnvironmentRequest(tmp)
	return nil
}

func (c *CreateEnvironmentRequest) MarshalJSON() ([]byte, error) {
	type embed CreateEnvironmentRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
