package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type InputAudioChunk struct {
	MessageType  *string `json:"message_type,omitempty" xml:"message_type,omitempty"`
	AudioBase64  string  `json:"audio_base_64" xml:"audio_base_64" required:"true"`
	Commit       *bool   `json:"commit,omitempty" xml:"commit,omitempty"`
	PreviousText *string `json:"previous_text,omitempty" xml:"previous_text,omitempty"`
	// A library of terms to recognize more accurately, supported by Scribe v3 Live models. Only set this on the first audio chunk, not as a query parameter. At most 10,000 terms are allowed, each at most 50 characters.
	KeytermLibrary  []string `json:"keyterm_library,omitempty" xml:"keyterm_library,omitempty"`
	extraProperties map[string]interface{}
}

func (i *InputAudioChunk) GetMessageType() (value *string) {
	if i == nil {
		return
	}
	return i.MessageType
}

func (i *InputAudioChunk) GetAudioBase64() (value string) {
	if i == nil {
		return
	}
	return i.AudioBase64
}

func (i *InputAudioChunk) GetCommit() (value *bool) {
	if i == nil {
		return
	}
	return i.Commit
}

func (i *InputAudioChunk) GetPreviousText() (value *string) {
	if i == nil {
		return
	}
	return i.PreviousText
}

func (i *InputAudioChunk) GetKeytermLibrary() (value []string) {
	if i == nil {
		return
	}
	return i.KeytermLibrary
}

func (i *InputAudioChunk) SetMessageType(messageType *string) {
	i.MessageType = messageType
}

func (i *InputAudioChunk) SetAudioBase64(audioBase64 string) {
	i.AudioBase64 = audioBase64
}

func (i *InputAudioChunk) SetCommit(commit *bool) {
	i.Commit = commit
}

func (i *InputAudioChunk) SetPreviousText(previousText *string) {
	i.PreviousText = previousText
}

func (i *InputAudioChunk) SetKeytermLibrary(keytermLibrary []string) {
	i.KeytermLibrary = keytermLibrary
}

func (i *InputAudioChunk) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i InputAudioChunk) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InputAudioChunk to string"
	}
	return string(jsonData)
}

func (i *InputAudioChunk) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, i); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	return nil
}

func (i *InputAudioChunk) MarshalJSON() ([]byte, error) {
	type embed InputAudioChunk
	var marshaler = struct {
		embed
	}{
		embed: embed(*i),
	}
	return json.Marshal(marshaler)
}
