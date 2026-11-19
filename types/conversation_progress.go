package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationProgress struct {
	// Calls due to be placed that haven't started.
	Scheduled int `json:"scheduled" xml:"scheduled" required:"true"`
	// Scheduled calls by call type. Missing keys mean none.
	ScheduledByType map[string]int `json:"scheduled_by_type" required:"true"`
	// WhatsApp calls waiting on the contact to grant call permission.
	AwaitingPermission int `json:"awaiting_permission" xml:"awaiting_permission" required:"true"`
	// Calls under way.
	InProgress int `json:"in_progress" xml:"in_progress" required:"true"`
	// Ended calls whose analysis the campaign is waiting on.
	AwaitingAnalysis int `json:"awaiting_analysis" xml:"awaiting_analysis" required:"true"`
	extraProperties  map[string]interface{}
}

func (c *ConversationProgress) GetScheduled() (value int) {
	if c == nil {
		return
	}
	return c.Scheduled
}

func (c *ConversationProgress) GetScheduledByType() (value map[string]int) {
	if c == nil {
		return
	}
	return c.ScheduledByType
}

func (c *ConversationProgress) GetAwaitingPermission() (value int) {
	if c == nil {
		return
	}
	return c.AwaitingPermission
}

func (c *ConversationProgress) GetInProgress() (value int) {
	if c == nil {
		return
	}
	return c.InProgress
}

func (c *ConversationProgress) GetAwaitingAnalysis() (value int) {
	if c == nil {
		return
	}
	return c.AwaitingAnalysis
}

func (c *ConversationProgress) SetScheduled(scheduled int) {
	c.Scheduled = scheduled
}

func (c *ConversationProgress) SetScheduledByType(scheduledByType map[string]int) {
	c.ScheduledByType = scheduledByType
}

func (c *ConversationProgress) SetAwaitingPermission(awaitingPermission int) {
	c.AwaitingPermission = awaitingPermission
}

func (c *ConversationProgress) SetInProgress(inProgress int) {
	c.InProgress = inProgress
}

func (c *ConversationProgress) SetAwaitingAnalysis(awaitingAnalysis int) {
	c.AwaitingAnalysis = awaitingAnalysis
}

func (c *ConversationProgress) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationProgress) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationProgress to string"
	}
	return string(jsonData)
}

func (c *ConversationProgress) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ConversationProgress
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ConversationProgress(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationProgress) MarshalJSON() ([]byte, error) {
	type embed ConversationProgress
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
