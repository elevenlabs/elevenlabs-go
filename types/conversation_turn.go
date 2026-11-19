package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationTurn struct {
	TurnId string            `json:"turn_id" xml:"turn_id" required:"true"`
	Status AgentOutputStatus `json:"status" xml:"status" required:"true"`
	Output []*TextOutputItem `json:"output,omitempty" xml:"output,omitempty"`
	Error  *AgentOutputError `json:"error,omitempty" xml:"error,omitempty"`
	// The actions the agent waits on. Answer each `client_tool_call` with a `client_tool_result` item in `input` to continue.
	RequiredActions []*ClientToolCallRequiredAction `json:"required_actions,omitempty" xml:"required_actions,omitempty"`
	// Resume token for this turn's event stream. Present while the turn is in progress; pass it as Last-Event-ID to stream the turn's events.
	StreamCursor    *string `json:"stream_cursor,omitempty" xml:"stream_cursor,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ConversationTurn) GetTurnId() (value string) {
	if c == nil {
		return
	}
	return c.TurnId
}

func (c *ConversationTurn) GetStatus() (value AgentOutputStatus) {
	if c == nil {
		return
	}
	return c.Status
}

func (c *ConversationTurn) GetOutput() (value []*TextOutputItem) {
	if c == nil {
		return
	}
	return c.Output
}

func (c *ConversationTurn) GetError() (value *AgentOutputError) {
	if c == nil {
		return
	}
	return c.Error
}

func (c *ConversationTurn) GetRequiredActions() (value []*ClientToolCallRequiredAction) {
	if c == nil {
		return
	}
	return c.RequiredActions
}

func (c *ConversationTurn) GetStreamCursor() (value *string) {
	if c == nil {
		return
	}
	return c.StreamCursor
}

func (c *ConversationTurn) SetTurnId(turnId string) {
	c.TurnId = turnId
}

func (c *ConversationTurn) SetStatus(status AgentOutputStatus) {
	c.Status = status
}

func (c *ConversationTurn) SetOutput(output []*TextOutputItem) {
	c.Output = output
}

func (c *ConversationTurn) SetError(error *AgentOutputError) {
	c.Error = error
}

func (c *ConversationTurn) SetRequiredActions(requiredActions []*ClientToolCallRequiredAction) {
	c.RequiredActions = requiredActions
}

func (c *ConversationTurn) SetStreamCursor(streamCursor *string) {
	c.StreamCursor = streamCursor
}

func (c *ConversationTurn) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationTurn) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationTurn to string"
	}
	return string(jsonData)
}

func (c *ConversationTurn) UnmarshalJSON(data []byte) error {
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

func (c *ConversationTurn) MarshalJSON() ([]byte, error) {
	type embed ConversationTurn
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
