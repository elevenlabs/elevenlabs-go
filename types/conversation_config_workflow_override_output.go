package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationConfigWorkflowOverrideOutput struct {
	// If enabled audio will not be processed and only text will be used, use to avoid audio pricing.
	TextOnly *bool `json:"text_only,omitempty" xml:"text_only,omitempty"`
	// The maximum duration of a conversation in seconds
	MaxDurationSeconds *int `json:"max_duration_seconds,omitempty" xml:"max_duration_seconds,omitempty"`
	// The events that will be sent to the client
	ClientEvents []ClientEvent `json:"client_events,omitempty" xml:"client_events,omitempty"`
	// Configuration for file input (image/PDF uploads) during conversations.
	FileInput *FileInputConfigWorkflowOverride `json:"file_input,omitempty" xml:"file_input,omitempty"`
	// Enable real-time monitoring of conversations via WebSocket
	MonitoringEnabled *bool `json:"monitoring_enabled,omitempty" xml:"monitoring_enabled,omitempty"`
	// The events that will be sent to monitoring connections.
	MonitoringEvents []ClientEvent `json:"monitoring_events,omitempty" xml:"monitoring_events,omitempty"`
	// Configure DTMF (keypad) input collection during phone calls
	DtmfInputSettings *DtmfInputConfig `json:"dtmf_input_settings,omitempty" xml:"dtmf_input_settings,omitempty"`
	// Configuration for background sound during conversations.
	BackgroundSound *BackgroundSoundConfigWorkflowOverride `json:"background_sound,omitempty" xml:"background_sound,omitempty"`
	// When enabled and knowledge base content is present, the LLM is instructed to report which sources it used.
	SourceAttribution *bool `json:"source_attribution,omitempty" xml:"source_attribution,omitempty"`
	extraProperties   map[string]interface{}
}

func (c *ConversationConfigWorkflowOverrideOutput) GetTextOnly() (value *bool) {
	if c == nil {
		return
	}
	return c.TextOnly
}

func (c *ConversationConfigWorkflowOverrideOutput) GetMaxDurationSeconds() (value *int) {
	if c == nil {
		return
	}
	return c.MaxDurationSeconds
}

func (c *ConversationConfigWorkflowOverrideOutput) GetClientEvents() (value []ClientEvent) {
	if c == nil {
		return
	}
	return c.ClientEvents
}

func (c *ConversationConfigWorkflowOverrideOutput) GetFileInput() (value *FileInputConfigWorkflowOverride) {
	if c == nil {
		return
	}
	return c.FileInput
}

func (c *ConversationConfigWorkflowOverrideOutput) GetMonitoringEnabled() (value *bool) {
	if c == nil {
		return
	}
	return c.MonitoringEnabled
}

func (c *ConversationConfigWorkflowOverrideOutput) GetMonitoringEvents() (value []ClientEvent) {
	if c == nil {
		return
	}
	return c.MonitoringEvents
}

func (c *ConversationConfigWorkflowOverrideOutput) GetDtmfInputSettings() (value *DtmfInputConfig) {
	if c == nil {
		return
	}
	return c.DtmfInputSettings
}

func (c *ConversationConfigWorkflowOverrideOutput) GetBackgroundSound() (value *BackgroundSoundConfigWorkflowOverride) {
	if c == nil {
		return
	}
	return c.BackgroundSound
}

func (c *ConversationConfigWorkflowOverrideOutput) GetSourceAttribution() (value *bool) {
	if c == nil {
		return
	}
	return c.SourceAttribution
}

func (c *ConversationConfigWorkflowOverrideOutput) SetTextOnly(textOnly *bool) {
	c.TextOnly = textOnly
}

func (c *ConversationConfigWorkflowOverrideOutput) SetMaxDurationSeconds(maxDurationSeconds *int) {
	c.MaxDurationSeconds = maxDurationSeconds
}

func (c *ConversationConfigWorkflowOverrideOutput) SetClientEvents(clientEvents []ClientEvent) {
	c.ClientEvents = clientEvents
}

func (c *ConversationConfigWorkflowOverrideOutput) SetFileInput(fileInput *FileInputConfigWorkflowOverride) {
	c.FileInput = fileInput
}

func (c *ConversationConfigWorkflowOverrideOutput) SetMonitoringEnabled(monitoringEnabled *bool) {
	c.MonitoringEnabled = monitoringEnabled
}

func (c *ConversationConfigWorkflowOverrideOutput) SetMonitoringEvents(monitoringEvents []ClientEvent) {
	c.MonitoringEvents = monitoringEvents
}

func (c *ConversationConfigWorkflowOverrideOutput) SetDtmfInputSettings(dtmfInputSettings *DtmfInputConfig) {
	c.DtmfInputSettings = dtmfInputSettings
}

func (c *ConversationConfigWorkflowOverrideOutput) SetBackgroundSound(backgroundSound *BackgroundSoundConfigWorkflowOverride) {
	c.BackgroundSound = backgroundSound
}

func (c *ConversationConfigWorkflowOverrideOutput) SetSourceAttribution(sourceAttribution *bool) {
	c.SourceAttribution = sourceAttribution
}

func (c *ConversationConfigWorkflowOverrideOutput) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationConfigWorkflowOverrideOutput) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationConfigWorkflowOverrideOutput to string"
	}
	return string(jsonData)
}

func (c *ConversationConfigWorkflowOverrideOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationConfigWorkflowOverrideOutput) MarshalJSON() ([]byte, error) {
	type embed ConversationConfigWorkflowOverrideOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
