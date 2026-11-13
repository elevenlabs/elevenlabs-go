package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CompactionSettingsWorkflowOverride struct {
	// Whether context compaction is enabled for this agent.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// Context window fraction at which compaction is triggered.
	SoftTriggerFraction *float64 `json:"soft_trigger_fraction,omitempty" xml:"soft_trigger_fraction,omitempty"`
	// Number of recent messages kept verbatim after compaction.
	TailSize *int `json:"tail_size,omitempty" xml:"tail_size,omitempty"`
	// Minimum number of tokens that compaction must reclaim.
	MinReclaimableTokens *int `json:"min_reclaimable_tokens,omitempty" xml:"min_reclaimable_tokens,omitempty"`
	extraProperties      map[string]interface{}
}

func (c *CompactionSettingsWorkflowOverride) GetEnabled() (value *bool) {
	if c == nil {
		return
	}
	return c.Enabled
}

func (c *CompactionSettingsWorkflowOverride) GetSoftTriggerFraction() (value *float64) {
	if c == nil {
		return
	}
	return c.SoftTriggerFraction
}

func (c *CompactionSettingsWorkflowOverride) GetTailSize() (value *int) {
	if c == nil {
		return
	}
	return c.TailSize
}

func (c *CompactionSettingsWorkflowOverride) GetMinReclaimableTokens() (value *int) {
	if c == nil {
		return
	}
	return c.MinReclaimableTokens
}

func (c *CompactionSettingsWorkflowOverride) SetEnabled(enabled *bool) {
	c.Enabled = enabled
}

func (c *CompactionSettingsWorkflowOverride) SetSoftTriggerFraction(softTriggerFraction *float64) {
	c.SoftTriggerFraction = softTriggerFraction
}

func (c *CompactionSettingsWorkflowOverride) SetTailSize(tailSize *int) {
	c.TailSize = tailSize
}

func (c *CompactionSettingsWorkflowOverride) SetMinReclaimableTokens(minReclaimableTokens *int) {
	c.MinReclaimableTokens = minReclaimableTokens
}

func (c *CompactionSettingsWorkflowOverride) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CompactionSettingsWorkflowOverride) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CompactionSettingsWorkflowOverride to string"
	}
	return string(jsonData)
}

func (c *CompactionSettingsWorkflowOverride) UnmarshalJSON(data []byte) error {
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

func (c *CompactionSettingsWorkflowOverride) MarshalJSON() ([]byte, error) {
	type embed CompactionSettingsWorkflowOverride
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
