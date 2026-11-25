package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A finished cue, and the last message with its `cue_index`. Each cue is committed
// before any message for the next, so at most one cue is open at a time, and a session
// that ends normally commits every cue.
//
// The cue keeps the text its partials showed, except that its last word can move to
// the next cue: when the text after it starts with punctuation, such as ",", that has
// no room in this cue, the word moves with it so the next cue doesn't start with it.
type SubtitlingCommittedSubtitle struct {
	messageType string
	// The cue's position in the session, counting from 0.
	CueIndex int `json:"cue_index" xml:"cue_index" required:"true"`
	// The cue's text, one entry per line.
	Lines []string `json:"lines" xml:"lines" required:"true"`
	// When to show the cue, in milliseconds since the session_started message. It doesn't change after the cue's first message.
	StartMs int `json:"start_ms" xml:"start_ms" required:"true"`
	// When to hide the cue, in milliseconds since the session_started message.
	EndMs           int `json:"end_ms" xml:"end_ms" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SubtitlingCommittedSubtitle) GetCueIndex() (value int) {
	if s == nil {
		return
	}
	return s.CueIndex
}

func (s *SubtitlingCommittedSubtitle) GetLines() (value []string) {
	if s == nil {
		return
	}
	return s.Lines
}

func (s *SubtitlingCommittedSubtitle) GetStartMs() (value int) {
	if s == nil {
		return
	}
	return s.StartMs
}

func (s *SubtitlingCommittedSubtitle) GetEndMs() (value int) {
	if s == nil {
		return
	}
	return s.EndMs
}

func (s *SubtitlingCommittedSubtitle) SetCueIndex(cueIndex int) {
	s.CueIndex = cueIndex
}

func (s *SubtitlingCommittedSubtitle) SetLines(lines []string) {
	s.Lines = lines
}

func (s *SubtitlingCommittedSubtitle) SetStartMs(startMs int) {
	s.StartMs = startMs
}

func (s *SubtitlingCommittedSubtitle) SetEndMs(endMs int) {
	s.EndMs = endMs
}

func (s *SubtitlingCommittedSubtitle) MessageType() string {
	return s.messageType
}

func (s *SubtitlingCommittedSubtitle) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SubtitlingCommittedSubtitle) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SubtitlingCommittedSubtitle to string"
	}
	return string(jsonData)
}

func (s *SubtitlingCommittedSubtitle) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SubtitlingCommittedSubtitle
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SubtitlingCommittedSubtitle(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "committed_subtitle" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "committed_subtitle", *literals.MessageType)
	}
	s.messageType = "committed_subtitle"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "message_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SubtitlingCommittedSubtitle) MarshalJSON() ([]byte, error) {
	type embed SubtitlingCommittedSubtitle
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*s),
		MessageType: "committed_subtitle",
	}
	return json.Marshal(marshaler)
}
