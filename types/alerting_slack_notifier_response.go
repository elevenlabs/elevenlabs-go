package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AlertingSlackNotifierResponse struct {
	Type            *string `json:"type,omitempty" xml:"type,omitempty"`
	IntegrationType *string `json:"integration_type,omitempty" xml:"integration_type,omitempty"`
	ConnectionId    string  `json:"connection_id" xml:"connection_id" required:"true"`
	ChannelId       string  `json:"channel_id" xml:"channel_id" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AlertingSlackNotifierResponse) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AlertingSlackNotifierResponse) GetIntegrationType() (value *string) {
	if a == nil {
		return
	}
	return a.IntegrationType
}

func (a *AlertingSlackNotifierResponse) GetConnectionId() (value string) {
	if a == nil {
		return
	}
	return a.ConnectionId
}

func (a *AlertingSlackNotifierResponse) GetChannelId() (value string) {
	if a == nil {
		return
	}
	return a.ChannelId
}

func (a *AlertingSlackNotifierResponse) SetType(type_ *string) {
	a.Type = type_
}

func (a *AlertingSlackNotifierResponse) SetIntegrationType(integrationType *string) {
	a.IntegrationType = integrationType
}

func (a *AlertingSlackNotifierResponse) SetConnectionId(connectionId string) {
	a.ConnectionId = connectionId
}

func (a *AlertingSlackNotifierResponse) SetChannelId(channelId string) {
	a.ChannelId = channelId
}

func (a *AlertingSlackNotifierResponse) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AlertingSlackNotifierResponse) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AlertingSlackNotifierResponse to string"
	}
	return string(jsonData)
}

func (a *AlertingSlackNotifierResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AlertingSlackNotifierResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AlertingSlackNotifierResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AlertingSlackNotifierResponse) MarshalJSON() ([]byte, error) {
	type embed AlertingSlackNotifierResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
