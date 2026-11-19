package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ListAuthConnectionsResponse struct {
	AuthConnections []*ListAuthConnectionsResponseAuthConnectionsItem `json:"auth_connections" xml:"auth_connections" required:"true"`
	extraProperties map[string]interface{}
}

func (l *ListAuthConnectionsResponse) GetAuthConnections() (value []*ListAuthConnectionsResponseAuthConnectionsItem) {
	if l == nil {
		return
	}
	return l.AuthConnections
}

func (l *ListAuthConnectionsResponse) SetAuthConnections(authConnections []*ListAuthConnectionsResponseAuthConnectionsItem) {
	l.AuthConnections = authConnections
}

func (l *ListAuthConnectionsResponse) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l ListAuthConnectionsResponse) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListAuthConnectionsResponse to string"
	}
	return string(jsonData)
}

func (l *ListAuthConnectionsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListAuthConnectionsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListAuthConnectionsResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *ListAuthConnectionsResponse) MarshalJSON() ([]byte, error) {
	type embed ListAuthConnectionsResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
