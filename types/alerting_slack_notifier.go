package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AlertingSlackNotifier struct {
	Type            *string `json:"type,omitempty" xml:"type,omitempty"`
	IntegrationType *string `json:"integration_type,omitempty" xml:"integration_type,omitempty"`
	// ID of the workspace integration connection to deliver alert lifecycle notifications to. The connection's integration must have the monitoring capability and match ``integration_type``.
	ConnectionId string `json:"connection_id" xml:"connection_id" required:"true" minLength:"1"`
	// ID of the Slack channel to post alert notifications to, e.g. ``C0123456789``. The Slack app must be a member of the channel and have the ``chat:write`` scope, or ``chat:write.public`` for public channels it has not joined.
	ChannelId       string `json:"channel_id" xml:"channel_id" required:"true" maxLength:"64" minLength:"1"`
	extraProperties map[string]interface{}
}

func (a *AlertingSlackNotifier) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AlertingSlackNotifier) GetIntegrationType() (value *string) {
	if a == nil {
		return
	}
	return a.IntegrationType
}

func (a *AlertingSlackNotifier) GetConnectionId() (value string) {
	if a == nil {
		return
	}
	return a.ConnectionId
}

func (a *AlertingSlackNotifier) GetChannelId() (value string) {
	if a == nil {
		return
	}
	return a.ChannelId
}

func (a *AlertingSlackNotifier) SetType(type_ *string) {
	a.Type = type_
}

func (a *AlertingSlackNotifier) SetIntegrationType(integrationType *string) {
	a.IntegrationType = integrationType
}

func (a *AlertingSlackNotifier) SetConnectionId(connectionId string) {
	a.ConnectionId = connectionId
}

func (a *AlertingSlackNotifier) SetChannelId(channelId string) {
	a.ChannelId = channelId
}

func (a *AlertingSlackNotifier) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AlertingSlackNotifier) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AlertingSlackNotifier to string"
	}
	return string(jsonData)
}

func (a *AlertingSlackNotifier) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AlertingSlackNotifier
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AlertingSlackNotifier(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AlertingSlackNotifier) MarshalJSON() ([]byte, error) {
	type embed AlertingSlackNotifier
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
