package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AlertingPagerDutyNotifierResponse struct {
	Type            *string `json:"type,omitempty" xml:"type,omitempty"`
	IntegrationType *string `json:"integration_type,omitempty" xml:"integration_type,omitempty"`
	ConnectionId    string  `json:"connection_id" xml:"connection_id" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AlertingPagerDutyNotifierResponse) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AlertingPagerDutyNotifierResponse) GetIntegrationType() (value *string) {
	if a == nil {
		return
	}
	return a.IntegrationType
}

func (a *AlertingPagerDutyNotifierResponse) GetConnectionId() (value string) {
	if a == nil {
		return
	}
	return a.ConnectionId
}

func (a *AlertingPagerDutyNotifierResponse) SetType(type_ *string) {
	a.Type = type_
}

func (a *AlertingPagerDutyNotifierResponse) SetIntegrationType(integrationType *string) {
	a.IntegrationType = integrationType
}

func (a *AlertingPagerDutyNotifierResponse) SetConnectionId(connectionId string) {
	a.ConnectionId = connectionId
}

func (a *AlertingPagerDutyNotifierResponse) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AlertingPagerDutyNotifierResponse) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AlertingPagerDutyNotifierResponse to string"
	}
	return string(jsonData)
}

func (a *AlertingPagerDutyNotifierResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AlertingPagerDutyNotifierResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AlertingPagerDutyNotifierResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AlertingPagerDutyNotifierResponse) MarshalJSON() ([]byte, error) {
	type embed AlertingPagerDutyNotifierResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
