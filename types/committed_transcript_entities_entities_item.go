package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CommittedTranscriptEntitiesEntitiesItem struct {
	// The text that was identified as an entity.
	Text string `json:"text" xml:"text" required:"true"`
	// The type of entity detected (e.g., 'credit_card', 'email_address', 'person_name').
	EntityType string `json:"entity_type" xml:"entity_type" required:"true"`
	// Start character position in the transcript text.
	StartChar int `json:"start_char" xml:"start_char" required:"true"`
	// End character position in the transcript text.
	EndChar         int `json:"end_char" xml:"end_char" required:"true"`
	extraProperties map[string]interface{}
}

func (c *CommittedTranscriptEntitiesEntitiesItem) GetText() (value string) {
	if c == nil {
		return
	}
	return c.Text
}

func (c *CommittedTranscriptEntitiesEntitiesItem) GetEntityType() (value string) {
	if c == nil {
		return
	}
	return c.EntityType
}

func (c *CommittedTranscriptEntitiesEntitiesItem) GetStartChar() (value int) {
	if c == nil {
		return
	}
	return c.StartChar
}

func (c *CommittedTranscriptEntitiesEntitiesItem) GetEndChar() (value int) {
	if c == nil {
		return
	}
	return c.EndChar
}

func (c *CommittedTranscriptEntitiesEntitiesItem) SetText(text string) {
	c.Text = text
}

func (c *CommittedTranscriptEntitiesEntitiesItem) SetEntityType(entityType string) {
	c.EntityType = entityType
}

func (c *CommittedTranscriptEntitiesEntitiesItem) SetStartChar(startChar int) {
	c.StartChar = startChar
}

func (c *CommittedTranscriptEntitiesEntitiesItem) SetEndChar(endChar int) {
	c.EndChar = endChar
}

func (c *CommittedTranscriptEntitiesEntitiesItem) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CommittedTranscriptEntitiesEntitiesItem) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CommittedTranscriptEntitiesEntitiesItem to string"
	}
	return string(jsonData)
}

func (c *CommittedTranscriptEntitiesEntitiesItem) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CommittedTranscriptEntitiesEntitiesItem
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CommittedTranscriptEntitiesEntitiesItem(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CommittedTranscriptEntitiesEntitiesItem) MarshalJSON() ([]byte, error) {
	type embed CommittedTranscriptEntitiesEntitiesItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
