package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AlertingPagerDutyNotifier struct {
	Type            *string `json:"type,omitempty" xml:"type,omitempty"`
	IntegrationType *string `json:"integration_type,omitempty" xml:"integration_type,omitempty"`
	// ID of the workspace integration connection to deliver alert lifecycle notifications to. The connection's integration must have the monitoring capability and match ``integration_type``.
	ConnectionId    string `json:"connection_id" xml:"connection_id" required:"true" minLength:"1"`
	extraProperties map[string]interface{}
}

func (a *AlertingPagerDutyNotifier) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AlertingPagerDutyNotifier) GetIntegrationType() (value *string) {
	if a == nil {
		return
	}
	return a.IntegrationType
}

func (a *AlertingPagerDutyNotifier) GetConnectionId() (value string) {
	if a == nil {
		return
	}
	return a.ConnectionId
}

func (a *AlertingPagerDutyNotifier) SetType(type_ *string) {
	a.Type = type_
}

func (a *AlertingPagerDutyNotifier) SetIntegrationType(integrationType *string) {
	a.IntegrationType = integrationType
}

func (a *AlertingPagerDutyNotifier) SetConnectionId(connectionId string) {
	a.ConnectionId = connectionId
}

func (a *AlertingPagerDutyNotifier) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AlertingPagerDutyNotifier) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AlertingPagerDutyNotifier to string"
	}
	return string(jsonData)
}

func (a *AlertingPagerDutyNotifier) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AlertingPagerDutyNotifier
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AlertingPagerDutyNotifier(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AlertingPagerDutyNotifier) MarshalJSON() ([]byte, error) {
	type embed AlertingPagerDutyNotifier
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
