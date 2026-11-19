package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Alerting configuration used at both per-agent and per-workspace level.
// Cascade order for per-monitor threshold and auto-resolve: agent → workspace →
// system default.
type AlertingSettings struct {
	// Alerting configuration keyed by monitor name.
	MonitorConfigs map[string]*AlertingMonitorConfig `json:"monitor_configs,omitempty"`
	// How many minutes an alert can stay inactive before it is auto-resolved. Unset values fall through to the next layer.
	AutoResolveAfterInactiveMinutes *int `json:"auto_resolve_after_inactive_minutes,omitempty" xml:"auto_resolve_after_inactive_minutes,omitempty" min:"1"`
	// Delivery channels for alert lifecycle notifications. Stacked with other layers and deduped by ``webhook_id``, PagerDuty ``connection_id``, or Slack ``(connection_id, channel_id)``.
	Notifiers       []*AlertingSettingsNotifiersItem `json:"notifiers,omitempty" xml:"notifiers,omitempty" maxItems:"10"`
	extraProperties map[string]interface{}
}

func (a *AlertingSettings) GetMonitorConfigs() (value map[string]*AlertingMonitorConfig) {
	if a == nil {
		return
	}
	return a.MonitorConfigs
}

func (a *AlertingSettings) GetAutoResolveAfterInactiveMinutes() (value *int) {
	if a == nil {
		return
	}
	return a.AutoResolveAfterInactiveMinutes
}

func (a *AlertingSettings) GetNotifiers() (value []*AlertingSettingsNotifiersItem) {
	if a == nil {
		return
	}
	return a.Notifiers
}

func (a *AlertingSettings) SetMonitorConfigs(monitorConfigs map[string]*AlertingMonitorConfig) {
	a.MonitorConfigs = monitorConfigs
}

func (a *AlertingSettings) SetAutoResolveAfterInactiveMinutes(autoResolveAfterInactiveMinutes *int) {
	a.AutoResolveAfterInactiveMinutes = autoResolveAfterInactiveMinutes
}

func (a *AlertingSettings) SetNotifiers(notifiers []*AlertingSettingsNotifiersItem) {
	a.Notifiers = notifiers
}

func (a *AlertingSettings) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AlertingSettings) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AlertingSettings to string"
	}
	return string(jsonData)
}

func (a *AlertingSettings) UnmarshalJSON(data []byte) error {
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

func (a *AlertingSettings) MarshalJSON() ([]byte, error) {
	type embed AlertingSettings
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
