package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type CloseSocket struct {
	CloseSocket     *bool `json:"close_socket,omitempty" xml:"close_socket,omitempty"`
	extraProperties map[string]interface{}
}

func (c *CloseSocket) GetCloseSocket() (value *bool) {
	if c == nil {
		return
	}
	return c.CloseSocket
}

func (c *CloseSocket) SetCloseSocket(closeSocket *bool) {
	c.CloseSocket = closeSocket
}

func (c *CloseSocket) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CloseSocket) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CloseSocket to string"
	}
	return string(jsonData)
}

func (c *CloseSocket) UnmarshalJSON(data []byte) error {
	type unmarshaler CloseSocket
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CloseSocket(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CloseSocket) MarshalJSON() ([]byte, error) {
	type embed CloseSocket
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
