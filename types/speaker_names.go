package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// With infer_speaker_names, the names inferred so far for diarized speakers, keyed
// by speaker_id. Sent in full each time a speaker is named or a merge changes the
// names.
type SpeakerNames struct {
	MessageType     *string           `json:"message_type,omitempty" xml:"message_type,omitempty"`
	SpeakerNames    map[string]string `json:"speaker_names" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SpeakerNames) GetMessageType() (value *string) {
	if s == nil {
		return
	}
	return s.MessageType
}

func (s *SpeakerNames) GetSpeakerNames() (value map[string]string) {
	if s == nil {
		return
	}
	return s.SpeakerNames
}

func (s *SpeakerNames) SetMessageType(messageType *string) {
	s.MessageType = messageType
}

func (s *SpeakerNames) SetSpeakerNames(speakerNames map[string]string) {
	s.SpeakerNames = speakerNames
}

func (s *SpeakerNames) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SpeakerNames) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SpeakerNames to string"
	}
	return string(jsonData)
}

func (s *SpeakerNames) UnmarshalJSON(data []byte) error {
	type unmarshaler SpeakerNames
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SpeakerNames(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SpeakerNames) MarshalJSON() ([]byte, error) {
	type embed SpeakerNames
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
