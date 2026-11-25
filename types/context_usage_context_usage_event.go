package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ContextUsageContextUsageEvent struct {
	EventId int `json:"event_id" xml:"event_id" required:"true"`
	// Identifier of the LLM whose context window is reported.
	Model string `json:"model" xml:"model" required:"true"`
	// Prompt size of the turn's last LLM generation in provider tokens (uncached input plus cache reads and writes).
	ContextTokens int `json:"context_tokens" xml:"context_tokens" required:"true"`
	// The model's maximum context window size in tokens.
	ContextLimitTokens int `json:"context_limit_tokens" xml:"context_limit_tokens" required:"true"`
	extraProperties    map[string]interface{}
}

func (c *ContextUsageContextUsageEvent) GetEventId() (value int) {
	if c == nil {
		return
	}
	return c.EventId
}

func (c *ContextUsageContextUsageEvent) GetModel() (value string) {
	if c == nil {
		return
	}
	return c.Model
}

func (c *ContextUsageContextUsageEvent) GetContextTokens() (value int) {
	if c == nil {
		return
	}
	return c.ContextTokens
}

func (c *ContextUsageContextUsageEvent) GetContextLimitTokens() (value int) {
	if c == nil {
		return
	}
	return c.ContextLimitTokens
}

func (c *ContextUsageContextUsageEvent) SetEventId(eventId int) {
	c.EventId = eventId
}

func (c *ContextUsageContextUsageEvent) SetModel(model string) {
	c.Model = model
}

func (c *ContextUsageContextUsageEvent) SetContextTokens(contextTokens int) {
	c.ContextTokens = contextTokens
}

func (c *ContextUsageContextUsageEvent) SetContextLimitTokens(contextLimitTokens int) {
	c.ContextLimitTokens = contextLimitTokens
}

func (c *ContextUsageContextUsageEvent) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ContextUsageContextUsageEvent) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContextUsageContextUsageEvent to string"
	}
	return string(jsonData)
}

func (c *ContextUsageContextUsageEvent) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ContextUsageContextUsageEvent
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ContextUsageContextUsageEvent(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ContextUsageContextUsageEvent) MarshalJSON() ([]byte, error) {
	type embed ContextUsageContextUsageEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
