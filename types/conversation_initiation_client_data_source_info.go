package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Information about the source of conversation initiation
type ConversationInitiationClientDataSourceInfo struct {
	// Source of the conversation initiation
	Source *ConversationInitiationClientDataSourceInfoSource `json:"source,omitempty" xml:"source,omitempty"`
	// The SDK version number
	Version         *string `json:"version,omitempty" xml:"version,omitempty" maxLength:"50"`
	extraProperties map[string]interface{}
}

func (c *ConversationInitiationClientDataSourceInfo) GetSource() (value *ConversationInitiationClientDataSourceInfoSource) {
	if c == nil {
		return
	}
	return c.Source
}

func (c *ConversationInitiationClientDataSourceInfo) GetVersion() (value *string) {
	if c == nil {
		return
	}
	return c.Version
}

func (c *ConversationInitiationClientDataSourceInfo) SetSource(source *ConversationInitiationClientDataSourceInfoSource) {
	c.Source = source
}

func (c *ConversationInitiationClientDataSourceInfo) SetVersion(version *string) {
	c.Version = version
}

func (c *ConversationInitiationClientDataSourceInfo) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationInitiationClientDataSourceInfo) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationInitiationClientDataSourceInfo to string"
	}
	return string(jsonData)
}

func (c *ConversationInitiationClientDataSourceInfo) UnmarshalJSON(data []byte) error {
	type unmarshaler ConversationInitiationClientDataSourceInfo
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConversationInitiationClientDataSourceInfo(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationInitiationClientDataSourceInfo) MarshalJSON() ([]byte, error) {
	type embed ConversationInitiationClientDataSourceInfo
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
