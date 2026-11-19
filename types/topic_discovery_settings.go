package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Per-agent topic-discovery configuration. Cadence and analysis window are
// managed internally; this only exposes the customer-facing on/off toggle.
type TopicDiscoverySettings struct {
	extraProperties map[string]interface{}
}

func (t *TopicDiscoverySettings) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TopicDiscoverySettings) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TopicDiscoverySettings to string"
	}
	return string(jsonData)
}

func (t *TopicDiscoverySettings) UnmarshalJSON(data []byte) error {
	type unmarshaler TopicDiscoverySettings
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TopicDiscoverySettings(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TopicDiscoverySettings) MarshalJSON() ([]byte, error) {
	type embed TopicDiscoverySettings
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
