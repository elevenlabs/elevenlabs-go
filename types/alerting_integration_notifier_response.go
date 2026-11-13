package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AlertingIntegrationNotifierResponse struct {
	IntegrationType *AlertingIntegrationNotifierResponseIntegrationType `json:"integration_type,omitempty" xml:"integration_type,omitempty"`
	ConnectionId    string                                              `json:"connection_id" xml:"connection_id" required:"true"`
	ChannelId       *string                                             `json:"channel_id,omitempty" xml:"channel_id,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AlertingIntegrationNotifierResponse) GetIntegrationType() (value *AlertingIntegrationNotifierResponseIntegrationType) {
	if a == nil {
		return
	}
	return a.IntegrationType
}

func (a *AlertingIntegrationNotifierResponse) GetConnectionId() (value string) {
	if a == nil {
		return
	}
	return a.ConnectionId
}

func (a *AlertingIntegrationNotifierResponse) GetChannelId() (value *string) {
	if a == nil {
		return
	}
	return a.ChannelId
}

func (a *AlertingIntegrationNotifierResponse) SetIntegrationType(integrationType *AlertingIntegrationNotifierResponseIntegrationType) {
	a.IntegrationType = integrationType
}

func (a *AlertingIntegrationNotifierResponse) SetConnectionId(connectionId string) {
	a.ConnectionId = connectionId
}

func (a *AlertingIntegrationNotifierResponse) SetChannelId(channelId *string) {
	a.ChannelId = channelId
}

func (a *AlertingIntegrationNotifierResponse) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AlertingIntegrationNotifierResponse) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AlertingIntegrationNotifierResponse to string"
	}
	return string(jsonData)
}

func (a *AlertingIntegrationNotifierResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AlertingIntegrationNotifierResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AlertingIntegrationNotifierResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "type")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AlertingIntegrationNotifierResponse) MarshalJSON() ([]byte, error) {
	type embed AlertingIntegrationNotifierResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
