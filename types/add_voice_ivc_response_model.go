package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AddVoiceIvcResponseModel struct {
	// The ID of the newly created voice.
	VoiceId string `json:"voice_id" xml:"voice_id" required:"true"`
	// Whether the voice requires verification
	RequiresVerification bool `json:"requires_verification" xml:"requires_verification" required:"true"`
	// Set when the samples match a protected voice that is available in the Voice Library. The new voice requires verification before use; adding the library voice is the alternative.
	LibraryVoiceMatch *LibraryVoiceMatchResponseModel `json:"library_voice_match,omitempty" xml:"library_voice_match,omitempty"`
	extraProperties   map[string]interface{}
}

func (a *AddVoiceIvcResponseModel) GetVoiceId() (value string) {
	if a == nil {
		return
	}
	return a.VoiceId
}

func (a *AddVoiceIvcResponseModel) GetRequiresVerification() (value bool) {
	if a == nil {
		return
	}
	return a.RequiresVerification
}

func (a *AddVoiceIvcResponseModel) GetLibraryVoiceMatch() (value *LibraryVoiceMatchResponseModel) {
	if a == nil {
		return
	}
	return a.LibraryVoiceMatch
}

func (a *AddVoiceIvcResponseModel) SetVoiceId(voiceId string) {
	a.VoiceId = voiceId
}

func (a *AddVoiceIvcResponseModel) SetRequiresVerification(requiresVerification bool) {
	a.RequiresVerification = requiresVerification
}

func (a *AddVoiceIvcResponseModel) SetLibraryVoiceMatch(libraryVoiceMatch *LibraryVoiceMatchResponseModel) {
	a.LibraryVoiceMatch = libraryVoiceMatch
}

func (a *AddVoiceIvcResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AddVoiceIvcResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddVoiceIvcResponseModel to string"
	}
	return string(jsonData)
}

func (a *AddVoiceIvcResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AddVoiceIvcResponseModel) MarshalJSON() ([]byte, error) {
	type embed AddVoiceIvcResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
