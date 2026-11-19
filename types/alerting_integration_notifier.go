package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AlertingIntegrationNotifier struct {
	// Integration to deliver to. The server treats an omitted value as ``pagerduty``.
	IntegrationType *AlertingIntegrationNotifierIntegrationType `json:"integration_type,omitempty" xml:"integration_type,omitempty"`
	// ID of the workspace integration connection to deliver alert lifecycle notifications to. The connection's integration must have the monitoring capability and match ``integration_type``.
	ConnectionId string `json:"connection_id" xml:"connection_id" required:"true" minLength:"1"`
	// ID of the Slack channel to post alert notifications to, e.g. ``C0123456789``. Required when ``integration_type`` is ``slack``. The Slack app must be a member of the channel and have the ``chat:write`` scope, or ``chat:write.public`` for public channels it has not joined.
	ChannelId       *string `json:"channel_id,omitempty" xml:"channel_id,omitempty" maxLength:"64" minLength:"1"`
	extraProperties map[string]interface{}
}

func (a *AlertingIntegrationNotifier) GetIntegrationType() (value *AlertingIntegrationNotifierIntegrationType) {
	if a == nil {
		return
	}
	return a.IntegrationType
}

func (a *AlertingIntegrationNotifier) GetConnectionId() (value string) {
	if a == nil {
		return
	}
	return a.ConnectionId
}

func (a *AlertingIntegrationNotifier) GetChannelId() (value *string) {
	if a == nil {
		return
	}
	return a.ChannelId
}

func (a *AlertingIntegrationNotifier) SetIntegrationType(integrationType *AlertingIntegrationNotifierIntegrationType) {
	a.IntegrationType = integrationType
}

func (a *AlertingIntegrationNotifier) SetConnectionId(connectionId string) {
	a.ConnectionId = connectionId
}

func (a *AlertingIntegrationNotifier) SetChannelId(channelId *string) {
	a.ChannelId = channelId
}

func (a *AlertingIntegrationNotifier) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AlertingIntegrationNotifier) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AlertingIntegrationNotifier to string"
	}
	return string(jsonData)
}

func (a *AlertingIntegrationNotifier) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AlertingIntegrationNotifier
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AlertingIntegrationNotifier(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "type")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AlertingIntegrationNotifier) MarshalJSON() ([]byte, error) {
	type embed AlertingIntegrationNotifier
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
