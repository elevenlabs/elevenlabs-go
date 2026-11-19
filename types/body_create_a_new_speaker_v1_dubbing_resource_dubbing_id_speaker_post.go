package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BodyCreateANewSpeakerV1DubbingResourceDubbingIdSpeakerPost struct {
	// Name to attribute to this speaker.
	SpeakerName *string `json:"speaker_name,omitempty" xml:"speaker_name,omitempty"`
	// Either the identifier of a voice from the ElevenLabs voice library, or one of ['track-clone', 'clip-clone'].
	VoiceId *string `json:"voice_id,omitempty" xml:"voice_id,omitempty"`
	// For models that support it, the voice similarity value to use. This will default to 0.65, with a valid range of [0.0, 1.0].
	VoiceStability *float64 `json:"voice_stability,omitempty" xml:"voice_stability,omitempty"`
	// For models that support it, the voice similarity value to use. This will default to 1.0, with a valid range of [0.0, 1.0].
	VoiceSimilarity *float64 `json:"voice_similarity,omitempty" xml:"voice_similarity,omitempty"`
	// For models that support it, the voice style value to use. This will default to 1.0, with a valid range of [0.0, 1.0].
	VoiceStyle      *float64 `json:"voice_style,omitempty" xml:"voice_style,omitempty"`
	extraProperties map[string]interface{}
}

func (b *BodyCreateANewSpeakerV1DubbingResourceDubbingIdSpeakerPost) GetSpeakerName() (value *string) {
	if b == nil {
		return
	}
	return b.SpeakerName
}

func (b *BodyCreateANewSpeakerV1DubbingResourceDubbingIdSpeakerPost) GetVoiceId() (value *string) {
	if b == nil {
		return
	}
	return b.VoiceId
}

func (b *BodyCreateANewSpeakerV1DubbingResourceDubbingIdSpeakerPost) GetVoiceStability() (value *float64) {
	if b == nil {
		return
	}
	return b.VoiceStability
}

func (b *BodyCreateANewSpeakerV1DubbingResourceDubbingIdSpeakerPost) GetVoiceSimilarity() (value *float64) {
	if b == nil {
		return
	}
	return b.VoiceSimilarity
}

func (b *BodyCreateANewSpeakerV1DubbingResourceDubbingIdSpeakerPost) GetVoiceStyle() (value *float64) {
	if b == nil {
		return
	}
	return b.VoiceStyle
}

func (b *BodyCreateANewSpeakerV1DubbingResourceDubbingIdSpeakerPost) SetSpeakerName(speakerName *string) {
	b.SpeakerName = speakerName
}

func (b *BodyCreateANewSpeakerV1DubbingResourceDubbingIdSpeakerPost) SetVoiceId(voiceId *string) {
	b.VoiceId = voiceId
}

func (b *BodyCreateANewSpeakerV1DubbingResourceDubbingIdSpeakerPost) SetVoiceStability(voiceStability *float64) {
	b.VoiceStability = voiceStability
}

func (b *BodyCreateANewSpeakerV1DubbingResourceDubbingIdSpeakerPost) SetVoiceSimilarity(voiceSimilarity *float64) {
	b.VoiceSimilarity = voiceSimilarity
}

func (b *BodyCreateANewSpeakerV1DubbingResourceDubbingIdSpeakerPost) SetVoiceStyle(voiceStyle *float64) {
	b.VoiceStyle = voiceStyle
}

func (b *BodyCreateANewSpeakerV1DubbingResourceDubbingIdSpeakerPost) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b BodyCreateANewSpeakerV1DubbingResourceDubbingIdSpeakerPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyCreateANewSpeakerV1DubbingResourceDubbingIdSpeakerPost to string"
	}
	return string(jsonData)
}

func (b *BodyCreateANewSpeakerV1DubbingResourceDubbingIdSpeakerPost) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, b); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	return nil
}

func (b *BodyCreateANewSpeakerV1DubbingResourceDubbingIdSpeakerPost) MarshalJSON() ([]byte, error) {
	type embed BodyCreateANewSpeakerV1DubbingResourceDubbingIdSpeakerPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
