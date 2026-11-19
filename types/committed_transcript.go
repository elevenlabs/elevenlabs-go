package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CommittedTranscript struct {
	MessageType *string `json:"message_type,omitempty" xml:"message_type,omitempty"`
	Text        string  `json:"text" xml:"text" required:"true"`
	// Zero-based commit index when diarize is enabled. Matches the segment_index in the corresponding diarized_transcript.
	SegmentIndex    *int `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
	extraProperties map[string]interface{}
}

func (c *CommittedTranscript) GetMessageType() (value *string) {
	if c == nil {
		return
	}
	return c.MessageType
}

func (c *CommittedTranscript) GetText() (value string) {
	if c == nil {
		return
	}
	return c.Text
}

func (c *CommittedTranscript) GetSegmentIndex() (value *int) {
	if c == nil {
		return
	}
	return c.SegmentIndex
}

func (c *CommittedTranscript) SetMessageType(messageType *string) {
	c.MessageType = messageType
}

func (c *CommittedTranscript) SetText(text string) {
	c.Text = text
}

func (c *CommittedTranscript) SetSegmentIndex(segmentIndex *int) {
	c.SegmentIndex = segmentIndex
}

func (c *CommittedTranscript) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CommittedTranscript) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CommittedTranscript to string"
	}
	return string(jsonData)
}

func (c *CommittedTranscript) UnmarshalJSON(data []byte) error {
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

func (c *CommittedTranscript) MarshalJSON() ([]byte, error) {
	type embed CommittedTranscript
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
