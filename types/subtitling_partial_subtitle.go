package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// The cue still being filled. Each one replaces the previous message with the same
// `cue_index` and only adds text at the end, though line breaks may move as the cue
// fills. A cue gets zero or more of these before its `committed_subtitle`, so it can
// arrive already committed.
type SubtitlingPartialSubtitle struct {
	messageType string
	// The cue's position in the session, counting from 0.
	CueIndex int `json:"cue_index" xml:"cue_index" required:"true"`
	// The cue's text, one entry per line.
	Lines []string `json:"lines" xml:"lines" required:"true"`
	// When to show the cue, in milliseconds since the session_started message. It doesn't change after the cue's first message.
	StartMs         int `json:"start_ms" xml:"start_ms" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SubtitlingPartialSubtitle) GetCueIndex() (value int) {
	if s == nil {
		return
	}
	return s.CueIndex
}

func (s *SubtitlingPartialSubtitle) GetLines() (value []string) {
	if s == nil {
		return
	}
	return s.Lines
}

func (s *SubtitlingPartialSubtitle) GetStartMs() (value int) {
	if s == nil {
		return
	}
	return s.StartMs
}

func (s *SubtitlingPartialSubtitle) SetCueIndex(cueIndex int) {
	s.CueIndex = cueIndex
}

func (s *SubtitlingPartialSubtitle) SetLines(lines []string) {
	s.Lines = lines
}

func (s *SubtitlingPartialSubtitle) SetStartMs(startMs int) {
	s.StartMs = startMs
}

func (s *SubtitlingPartialSubtitle) MessageType() string {
	return s.messageType
}

func (s *SubtitlingPartialSubtitle) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SubtitlingPartialSubtitle) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SubtitlingPartialSubtitle to string"
	}
	return string(jsonData)
}

func (s *SubtitlingPartialSubtitle) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SubtitlingPartialSubtitle
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SubtitlingPartialSubtitle(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "partial_subtitle" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "partial_subtitle", *literals.MessageType)
	}
	s.messageType = "partial_subtitle"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "message_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SubtitlingPartialSubtitle) MarshalJSON() ([]byte, error) {
	type embed SubtitlingPartialSubtitle
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*s),
		MessageType: "partial_subtitle",
	}
	return json.Marshal(marshaler)
}
