package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConflictResolution struct {
	// `id` of the conflict this resolves, as the merge or rebase preview reported it.
	ConflictId string                   `json:"conflict_id" xml:"conflict_id" required:"true"`
	Choice     ConflictResolutionChoice `json:"choice" xml:"choice" required:"true"`
	// Text to keep when `choice` is `custom`, or for a procedure conflict the id of the procedure version to pin. Only conflicts with `accepts_custom_value` take one.
	CustomValue     *string `json:"custom_value,omitempty" xml:"custom_value,omitempty" maxLength:"1000000"`
	extraProperties map[string]interface{}
}

func (c *ConflictResolution) GetConflictId() (value string) {
	if c == nil {
		return
	}
	return c.ConflictId
}

func (c *ConflictResolution) GetChoice() (value ConflictResolutionChoice) {
	if c == nil {
		return
	}
	return c.Choice
}

func (c *ConflictResolution) GetCustomValue() (value *string) {
	if c == nil {
		return
	}
	return c.CustomValue
}

func (c *ConflictResolution) SetConflictId(conflictId string) {
	c.ConflictId = conflictId
}

func (c *ConflictResolution) SetChoice(choice ConflictResolutionChoice) {
	c.Choice = choice
}

func (c *ConflictResolution) SetCustomValue(customValue *string) {
	c.CustomValue = customValue
}

func (c *ConflictResolution) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConflictResolution) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConflictResolution to string"
	}
	return string(jsonData)
}

func (c *ConflictResolution) UnmarshalJSON(data []byte) error {
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

func (c *ConflictResolution) MarshalJSON() ([]byte, error) {
	type embed ConflictResolution
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
