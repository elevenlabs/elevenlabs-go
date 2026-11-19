package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Information about the source of conversation initiation
type ConversationInitiationSourceInfo struct {
	// Source of the conversation initiation
	Source *ConversationInitiationSource `json:"source,omitempty" xml:"source,omitempty"`
	// The SDK version number
	Version         *string `json:"version,omitempty" xml:"version,omitempty" maxLength:"50"`
	extraProperties map[string]interface{}
}

func (c *ConversationInitiationSourceInfo) GetSource() (value *ConversationInitiationSource) {
	if c == nil {
		return
	}
	return c.Source
}

func (c *ConversationInitiationSourceInfo) GetVersion() (value *string) {
	if c == nil {
		return
	}
	return c.Version
}

func (c *ConversationInitiationSourceInfo) SetSource(source *ConversationInitiationSource) {
	c.Source = source
}

func (c *ConversationInitiationSourceInfo) SetVersion(version *string) {
	c.Version = version
}

func (c *ConversationInitiationSourceInfo) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationInitiationSourceInfo) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationInitiationSourceInfo to string"
	}
	return string(jsonData)
}

func (c *ConversationInitiationSourceInfo) UnmarshalJSON(data []byte) error {
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

func (c *ConversationInitiationSourceInfo) MarshalJSON() ([]byte, error) {
	type embed ConversationInitiationSourceInfo
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
