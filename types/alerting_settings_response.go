package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Customer-facing view of alerting settings. Unlike AdminAlertingSettingsResponse,
// it has no internal_notifiers field: those are ElevenLabs-internal delivery
// channels whose URLs must never be returned outside the admin API.
type AlertingSettingsResponse struct {
	MonitorConfigs                  map[string]*AlertingMonitorConfig        `json:"monitor_configs,omitempty"`
	AutoResolveAfterInactiveMinutes *int                                     `json:"auto_resolve_after_inactive_minutes,omitempty" xml:"auto_resolve_after_inactive_minutes,omitempty"`
	Environments                    []string                                 `json:"environments,omitempty" xml:"environments,omitempty"`
	Notifiers                       []*AlertingSettingsResponseNotifiersItem `json:"notifiers,omitempty" xml:"notifiers,omitempty"`
	extraProperties                 map[string]interface{}
}

func (a *AlertingSettingsResponse) GetMonitorConfigs() (value map[string]*AlertingMonitorConfig) {
	if a == nil {
		return
	}
	return a.MonitorConfigs
}

func (a *AlertingSettingsResponse) GetAutoResolveAfterInactiveMinutes() (value *int) {
	if a == nil {
		return
	}
	return a.AutoResolveAfterInactiveMinutes
}

func (a *AlertingSettingsResponse) GetEnvironments() (value []string) {
	if a == nil {
		return
	}
	return a.Environments
}

func (a *AlertingSettingsResponse) GetNotifiers() (value []*AlertingSettingsResponseNotifiersItem) {
	if a == nil {
		return
	}
	return a.Notifiers
}

func (a *AlertingSettingsResponse) SetMonitorConfigs(monitorConfigs map[string]*AlertingMonitorConfig) {
	a.MonitorConfigs = monitorConfigs
}

func (a *AlertingSettingsResponse) SetAutoResolveAfterInactiveMinutes(autoResolveAfterInactiveMinutes *int) {
	a.AutoResolveAfterInactiveMinutes = autoResolveAfterInactiveMinutes
}

func (a *AlertingSettingsResponse) SetEnvironments(environments []string) {
	a.Environments = environments
}

func (a *AlertingSettingsResponse) SetNotifiers(notifiers []*AlertingSettingsResponseNotifiersItem) {
	a.Notifiers = notifiers
}

func (a *AlertingSettingsResponse) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AlertingSettingsResponse) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AlertingSettingsResponse to string"
	}
	return string(jsonData)
}

func (a *AlertingSettingsResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AlertingSettingsResponse) MarshalJSON() ([]byte, error) {
	type embed AlertingSettingsResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
