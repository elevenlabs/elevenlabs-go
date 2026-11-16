package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CommittedTranscriptEntities struct {
	MessageType     *string                                    `json:"message_type,omitempty" xml:"message_type,omitempty"`
	Text            string                                     `json:"text" xml:"text" required:"true"`
	Entities        []*CommittedTranscriptEntitiesEntitiesItem `json:"entities,omitempty" xml:"entities,omitempty"`
	extraProperties map[string]interface{}
}

func (c *CommittedTranscriptEntities) GetMessageType() (value *string) {
	if c == nil {
		return
	}
	return c.MessageType
}

func (c *CommittedTranscriptEntities) GetText() (value string) {
	if c == nil {
		return
	}
	return c.Text
}

func (c *CommittedTranscriptEntities) GetEntities() (value []*CommittedTranscriptEntitiesEntitiesItem) {
	if c == nil {
		return
	}
	return c.Entities
}

func (c *CommittedTranscriptEntities) SetMessageType(messageType *string) {
	c.MessageType = messageType
}

func (c *CommittedTranscriptEntities) SetText(text string) {
	c.Text = text
}

func (c *CommittedTranscriptEntities) SetEntities(entities []*CommittedTranscriptEntitiesEntitiesItem) {
	c.Entities = entities
}

func (c *CommittedTranscriptEntities) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CommittedTranscriptEntities) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CommittedTranscriptEntities to string"
	}
	return string(jsonData)
}

func (c *CommittedTranscriptEntities) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CommittedTranscriptEntities
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CommittedTranscriptEntities(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CommittedTranscriptEntities) MarshalJSON() ([]byte, error) {
	type embed CommittedTranscriptEntities
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
