package types

import (
	"fmt"
)

// How to format redacted entities. 'redacted' replaces with {REDACTED}, 'entity_type' replaces with {ENTITY_TYPE}, 'enumerated_entity_type' replaces with {ENTITY_TYPE_N} where N enumerates each occurrence. Only used when entity_redaction is set.
type SpeechToTextRealtimeEntityRedactionMode string

const (
	SpeechToTextRealtimeEntityRedactionModeEntityType           SpeechToTextRealtimeEntityRedactionMode = "entity_type"
	SpeechToTextRealtimeEntityRedactionModeEnumeratedEntityType SpeechToTextRealtimeEntityRedactionMode = "enumerated_entity_type"
	SpeechToTextRealtimeEntityRedactionModeRedacted             SpeechToTextRealtimeEntityRedactionMode = "redacted"
)

func NewSpeechToTextRealtimeEntityRedactionModeFromString(s string) (SpeechToTextRealtimeEntityRedactionMode, error) {
	switch s {
	case "entity_type":
		return SpeechToTextRealtimeEntityRedactionModeEntityType, nil
	case "enumerated_entity_type":
		return SpeechToTextRealtimeEntityRedactionModeEnumeratedEntityType, nil
	case "redacted":
		return SpeechToTextRealtimeEntityRedactionModeRedacted, nil
	}
	var t SpeechToTextRealtimeEntityRedactionMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SpeechToTextRealtimeEntityRedactionMode) Ptr() *SpeechToTextRealtimeEntityRedactionMode {
	return &e
}
