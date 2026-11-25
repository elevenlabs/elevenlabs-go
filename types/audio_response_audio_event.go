package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AudioResponseAudioEvent struct {
	AudioBase64 string `json:"audio_base_64" xml:"audio_base_64" required:"true"`
	EventId     int    `json:"event_id" xml:"event_id" required:"true"`
	// Character-level timing data for the audio chunk, useful for synchronized text display or lip-syncing.
	Alignment *AudioResponseAudioEventAlignment `json:"alignment,omitempty" xml:"alignment,omitempty"`
	// Whether this is the last audio chunk of the current agent response.
	IsFinal         *bool `json:"is_final,omitempty" xml:"is_final,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AudioResponseAudioEvent) GetAudioBase64() (value string) {
	if a == nil {
		return
	}
	return a.AudioBase64
}

func (a *AudioResponseAudioEvent) GetEventId() (value int) {
	if a == nil {
		return
	}
	return a.EventId
}

func (a *AudioResponseAudioEvent) GetAlignment() (value *AudioResponseAudioEventAlignment) {
	if a == nil {
		return
	}
	return a.Alignment
}

func (a *AudioResponseAudioEvent) GetIsFinal() (value *bool) {
	if a == nil {
		return
	}
	return a.IsFinal
}

func (a *AudioResponseAudioEvent) SetAudioBase64(audioBase64 string) {
	a.AudioBase64 = audioBase64
}

func (a *AudioResponseAudioEvent) SetEventId(eventId int) {
	a.EventId = eventId
}

func (a *AudioResponseAudioEvent) SetAlignment(alignment *AudioResponseAudioEventAlignment) {
	a.Alignment = alignment
}

func (a *AudioResponseAudioEvent) SetIsFinal(isFinal *bool) {
	a.IsFinal = isFinal
}

func (a *AudioResponseAudioEvent) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AudioResponseAudioEvent) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AudioResponseAudioEvent to string"
	}
	return string(jsonData)
}

func (a *AudioResponseAudioEvent) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AudioResponseAudioEvent
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AudioResponseAudioEvent(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AudioResponseAudioEvent) MarshalJSON() ([]byte, error) {
	type embed AudioResponseAudioEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
