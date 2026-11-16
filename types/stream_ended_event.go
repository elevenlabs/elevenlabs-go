package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type StreamEndedEvent struct {
	StreamEndedEvent *StreamEndedData `json:"stream_ended_event" xml:"stream_ended_event" required:"true"`
	extraProperties  map[string]interface{}
}

func (s *StreamEndedEvent) GetStreamEndedEvent() (value *StreamEndedData) {
	if s == nil {
		return
	}
	return s.StreamEndedEvent
}

func (s *StreamEndedEvent) SetStreamEndedEvent(streamEndedEvent *StreamEndedData) {
	s.StreamEndedEvent = streamEndedEvent
}

func (s *StreamEndedEvent) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s StreamEndedEvent) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: StreamEndedEvent to string"
	}
	return string(jsonData)
}

func (s *StreamEndedEvent) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias StreamEndedEvent
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = StreamEndedEvent(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *StreamEndedEvent) MarshalJSON() ([]byte, error) {
	type embed StreamEndedEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
