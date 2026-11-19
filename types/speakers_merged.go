package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// With diarize, two speaker_ids turned out to be one person: every earlier use of
// from_speaker_id, in partial or diarized transcripts, is to_speaker_id. Ids are never
// reused; merges chain, so apply them in the order received.
type SpeakersMerged struct {
	MessageType     *string `json:"message_type,omitempty" xml:"message_type,omitempty"`
	FromSpeakerId   string  `json:"from_speaker_id" xml:"from_speaker_id" required:"true"`
	ToSpeakerId     string  `json:"to_speaker_id" xml:"to_speaker_id" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SpeakersMerged) GetMessageType() (value *string) {
	if s == nil {
		return
	}
	return s.MessageType
}

func (s *SpeakersMerged) GetFromSpeakerId() (value string) {
	if s == nil {
		return
	}
	return s.FromSpeakerId
}

func (s *SpeakersMerged) GetToSpeakerId() (value string) {
	if s == nil {
		return
	}
	return s.ToSpeakerId
}

func (s *SpeakersMerged) SetMessageType(messageType *string) {
	s.MessageType = messageType
}

func (s *SpeakersMerged) SetFromSpeakerId(fromSpeakerId string) {
	s.FromSpeakerId = fromSpeakerId
}

func (s *SpeakersMerged) SetToSpeakerId(toSpeakerId string) {
	s.ToSpeakerId = toSpeakerId
}

func (s *SpeakersMerged) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SpeakersMerged) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SpeakersMerged to string"
	}
	return string(jsonData)
}

func (s *SpeakersMerged) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SpeakersMerged
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SpeakersMerged(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SpeakersMerged) MarshalJSON() ([]byte, error) {
	type embed SpeakersMerged
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
