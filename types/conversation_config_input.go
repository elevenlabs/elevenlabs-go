package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationConfigInput struct {
	// If enabled audio will not be processed and only text will be used, use to avoid audio pricing.
	TextOnly *bool `json:"text_only,omitempty" xml:"text_only,omitempty"`
	// The maximum duration of a conversation in seconds
	MaxDurationSeconds *int `json:"max_duration_seconds,omitempty" xml:"max_duration_seconds,omitempty"`
	// The events that will be sent to the client
	ClientEvents []ClientEvent    `json:"client_events,omitempty" xml:"client_events,omitempty"`
	FileInput    *FileInputConfig `json:"file_input,omitempty" xml:"file_input,omitempty"`
	// Enable real-time monitoring of conversations via WebSocket
	MonitoringEnabled *bool `json:"monitoring_enabled,omitempty" xml:"monitoring_enabled,omitempty"`
	// The events that will be sent to monitoring connections.
	MonitoringEvents []ClientEvent `json:"monitoring_events,omitempty" xml:"monitoring_events,omitempty"`
	// Configure DTMF (keypad) input collection during phone calls
	DtmfInputSettings *DtmfInputConfig       `json:"dtmf_input_settings,omitempty" xml:"dtmf_input_settings,omitempty"`
	BackgroundSound   *BackgroundSoundConfig `json:"background_sound,omitempty" xml:"background_sound,omitempty"`
	// When enabled and knowledge base content is present, the LLM is instructed to report which sources it used.
	SourceAttribution *bool `json:"source_attribution,omitempty" xml:"source_attribution,omitempty"`
	extraProperties   map[string]interface{}
}

func (c *ConversationConfigInput) GetTextOnly() (value *bool) {
	if c == nil {
		return
	}
	return c.TextOnly
}

func (c *ConversationConfigInput) GetMaxDurationSeconds() (value *int) {
	if c == nil {
		return
	}
	return c.MaxDurationSeconds
}

func (c *ConversationConfigInput) GetClientEvents() (value []ClientEvent) {
	if c == nil {
		return
	}
	return c.ClientEvents
}

func (c *ConversationConfigInput) GetFileInput() (value *FileInputConfig) {
	if c == nil {
		return
	}
	return c.FileInput
}

func (c *ConversationConfigInput) GetMonitoringEnabled() (value *bool) {
	if c == nil {
		return
	}
	return c.MonitoringEnabled
}

func (c *ConversationConfigInput) GetMonitoringEvents() (value []ClientEvent) {
	if c == nil {
		return
	}
	return c.MonitoringEvents
}

func (c *ConversationConfigInput) GetDtmfInputSettings() (value *DtmfInputConfig) {
	if c == nil {
		return
	}
	return c.DtmfInputSettings
}

func (c *ConversationConfigInput) GetBackgroundSound() (value *BackgroundSoundConfig) {
	if c == nil {
		return
	}
	return c.BackgroundSound
}

func (c *ConversationConfigInput) GetSourceAttribution() (value *bool) {
	if c == nil {
		return
	}
	return c.SourceAttribution
}

func (c *ConversationConfigInput) SetTextOnly(textOnly *bool) {
	c.TextOnly = textOnly
}

func (c *ConversationConfigInput) SetMaxDurationSeconds(maxDurationSeconds *int) {
	c.MaxDurationSeconds = maxDurationSeconds
}

func (c *ConversationConfigInput) SetClientEvents(clientEvents []ClientEvent) {
	c.ClientEvents = clientEvents
}

func (c *ConversationConfigInput) SetFileInput(fileInput *FileInputConfig) {
	c.FileInput = fileInput
}

func (c *ConversationConfigInput) SetMonitoringEnabled(monitoringEnabled *bool) {
	c.MonitoringEnabled = monitoringEnabled
}

func (c *ConversationConfigInput) SetMonitoringEvents(monitoringEvents []ClientEvent) {
	c.MonitoringEvents = monitoringEvents
}

func (c *ConversationConfigInput) SetDtmfInputSettings(dtmfInputSettings *DtmfInputConfig) {
	c.DtmfInputSettings = dtmfInputSettings
}

func (c *ConversationConfigInput) SetBackgroundSound(backgroundSound *BackgroundSoundConfig) {
	c.BackgroundSound = backgroundSound
}

func (c *ConversationConfigInput) SetSourceAttribution(sourceAttribution *bool) {
	c.SourceAttribution = sourceAttribution
}

func (c *ConversationConfigInput) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationConfigInput) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationConfigInput to string"
	}
	return string(jsonData)
}

func (c *ConversationConfigInput) UnmarshalJSON(data []byte) error {
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

func (c *ConversationConfigInput) MarshalJSON() ([]byte, error) {
	type embed ConversationConfigInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
